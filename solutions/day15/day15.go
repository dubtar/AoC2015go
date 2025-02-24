package day15

import (
	h "go-aoc-template/internal/helpers"

	"regexp"
)

type Disk struct {
	period        int64
	startPosition int64
}

func ParseDisks(lines []string) []Disk {
	r := regexp.MustCompile(`Disc #. has (\d+) positions; at time=0, it is at position (\d+).`)
	disks := []Disk{}
	for i, line := range lines {
		matches := r.FindStringSubmatch(line)
		disk := Disk{period: h.ToInt(matches[1]), startPosition: h.ToInt(matches[2])}
		disk.startPosition = (disk.startPosition + int64(i) + 1) % disk.period
		disks = append(disks, disk)
	}
	return disks
}

func PartOne(lines []string) string {
	disks := ParseDisks(lines)
	best := 0
	for i := disks[0].period - disks[0].startPosition; true; i += disks[0].period {
		found := true
		last := 1
		for _, disk := range disks[1:] {
			if (i+disk.startPosition)%disk.period != 0 {
				found = false
				break
			}
			last += 1
		}
		if found {
			return h.ToString(i)
		}
		if last > best {
			print(i, ": ", last, "\n")
			for _, disk := range disks {
				print((disk.startPosition+int64(i))%disk.period, " ")
			}
			print("\n")
			best = last
		}
	}
	return "Failed"
}

func PartTwo(lines []string) string {
	disks := ParseDisks(lines)
	disks = append(disks, Disk{period: 11, startPosition: int64(len(disks) + 1)})
	best := 0
	for i := disks[0].period - disks[0].startPosition; true; i += disks[0].period {
		found := true
		last := 1
		for _, disk := range disks[1:] {
			if (i+disk.startPosition)%disk.period != 0 {
				found = false
				break
			}
			last += 1
		}
		if found {
			return h.ToString(i)
		}
		if last > best {
			print(i, ": ", last, "\n")
			for _, disk := range disks {
				print((disk.startPosition+int64(i))%disk.period, " ")
			}
			print("\n")
			best = last
		}
	}
	return "Failed"
}

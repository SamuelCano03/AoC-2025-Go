package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func ps(a ...any) {
	fmt.Println(a...)
}

func read() string {
	line, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	return strings.TrimSpace(line)
}

type Interval struct {
	min, max int
}

func toInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func part1() {
	var ranges []Interval
	for {
		line := read()
		if line == "" {
			break
		}
		parts := strings.Split(line, "-")
		ranges = append(ranges, Interval{toInt(parts[0]), toInt(parts[1])})
	}
	freshCount := 0
	for {
		line := read()
		if line == "" {
			break 
		}
		id := toInt(line)
		isFresh := false
		for _, r := range ranges {
			if id >= r.min && id <= r.max {
				isFresh = true
				break
			}
		}
		if isFresh {
			freshCount++
		}
	}
	ps(freshCount)
}

func part2() {
	var ranges []Interval
	for {
		line := read()
		if line == "" {
			break
		}
		parts := strings.Split(line, "-")
		ranges = append(ranges, Interval{toInt(parts[0]), toInt(parts[1])})
	}
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].min < ranges[j].min
	})

	var merged []Interval
	if len(ranges) > 0 {
		curr := ranges[0]
		for i := 1; i < len(ranges); i++ {
			next := ranges[i]
			if next.min <= curr.max+1 { 
				if next.max > curr.max {
					curr.max = next.max
				}
			} else {
				merged = append(merged, curr)
				curr = next
			}
		}
		merged = append(merged, curr)
	}

	totalFresh := 0
	for _, r := range merged {
		totalFresh += (r.max - r.min + 1)
	}

	ps(totalFresh)
}

func main() {
	isPart1 := false; 
	if isPart1 {
		part1()
	} else {
		part2()
	}
}

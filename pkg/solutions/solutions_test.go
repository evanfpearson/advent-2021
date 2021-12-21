package solutions

import (
	"github.com/evanfpearson/advent-2021/pkg/advent"
	"github.com/evanfpearson/advent-2021/pkg/solutions/day1"
	"github.com/evanfpearson/advent-2021/pkg/solutions/day10"
	"github.com/evanfpearson/advent-2021/pkg/solutions/day11"
	"github.com/evanfpearson/advent-2021/pkg/solutions/day12"
	"github.com/evanfpearson/advent-2021/pkg/solutions/day13"
	"github.com/evanfpearson/advent-2021/pkg/solutions/day14"
	"github.com/evanfpearson/advent-2021/pkg/solutions/day2"
	"github.com/evanfpearson/advent-2021/pkg/solutions/day3"
	"github.com/evanfpearson/advent-2021/pkg/solutions/day4"
	"github.com/evanfpearson/advent-2021/pkg/solutions/day5"
	"github.com/evanfpearson/advent-2021/pkg/solutions/day6"
	"github.com/evanfpearson/advent-2021/pkg/solutions/day7"
	"github.com/evanfpearson/advent-2021/pkg/solutions/day8"
	"github.com/evanfpearson/advent-2021/pkg/solutions/day9"
	"testing"
)

func TestSolutions(t *testing.T) {
	tests := []struct {
		day      int
		expected int
		solution advent.Solution
	}{
		{
			1,
			7,
			day1.A,
		},
		{
			1,
			5,
			day1.B,
		},
		{
			2,
			150,
			day2.A,
		},
		{
			2,
			900,
			day2.B,
		},
		{
			3,
			198,
			day3.A,
		},
		{
			3,
			230,
			day3.B,
		},
		{
			4,
			4512,
			day4.A,
		},
		{
			4,
			1924,
			day4.B,
		},
		{
			5,
			5,
			day5.A,
		},
		{
			5,
			12,
			day5.B,
		},
		{
			6,
			5934,
			day6.A,
		},
		{
			6,
			26984457539,
			day6.B,
		},
		{
			7,
			37,
			day7.A,
		},
		{
			7,
			168,
			day7.B,
		},
		{
			8,
			26,
			day8.A,
		},
		{
			8,
			61229,
			day8.B,
		},
		{
			9,
			15,
			day9.A,
		},
		{
			9,
			1134,
			day9.B,
		},
		{
			10,
			26397,
			day10.A,
		},
		{
			10,
			288957,
			day10.B,
		},
		{
			11,
			1656,
			day11.A,
		},
		{
			11,
			195,
			day11.B,
		},
		{
			12,
			226,
			day12.A,
		},
		{
			12,
			3509,
			day12.B,
		},
		{
			13,
			17,
			day13.A,
		},
		{
			13,
			16,
			day13.B,
		},
		{
			14,
			1588,
			day14.A,
		},
	}
	for _, test := range tests {
		out, _ := test.solution(&testCases[test.day-1])
		if out != test.expected {
			t.Errorf("Day %d: expected %d, got %d", test.day, test.expected, out)
		}
	}
}

package main

import (
	"strings"
)

var NormalBlockTimes = [][]string{
	{
		"07:39:00 AM",
		"09:15:00 AM",
	},
	{
		"09:20:00 AM",
		"10:56:00 AM",
	},
	{
		"11:01:00 AM",
		"01:04:00 PM",
	},
	{
		"01:09:00 PM",
		"02:45:00 PM",
	},
}

var Times = map[string][][]string{
	"w/ Assembly": {
		{
			"07:39:00 AM",
			"09:01:00 AM",
			"Block 1/5",
		},
		{
			"09:06:00 AM",
			"10:28:00 AM",
			"Block 2/6",
		},
		{
			"10:33:00 AM",
			"12:23:00 PM",
			"Block 3/7",
		},
		{
			"12:28:00 PM",
			"01:55:00 PM",
			"Block 4/8",
		},
		{
			"02:00:00 PM",
			"02:45:00 PM",
			"Assembly",
		},
	},
	"w/ Anchor": {
		{
			"07:39:00 AM",
			"09:01:00 AM",
			"Block 1/5",
		},
		{
			"09:06:00 AM",
			"10:28:00 AM",
			"Block 2/6",
		},
		{
			"10:33:00 AM",
			"12:23:00 PM",
			"Block 3/7",
		},
		{
			"12:28:00 PM",
			"01:18:00 PM",
			"Anchor",
		},
		{
			"01:23:00 PM",
			"02:45:00 PM",
			"Block 4/8",
		},
	},
	"1-4)": {
		append(NormalBlockTimes[0], "Block 1"),
		append(NormalBlockTimes[1], "Block 2"),
		append(NormalBlockTimes[2], "Block 3"),
		append(NormalBlockTimes[3], "Block 4"),
	},
	"5-8)": {
		append(NormalBlockTimes[0], "Block 5"),
		append(NormalBlockTimes[1], "Block 6"),
		append(NormalBlockTimes[2], "Block 7"),
		append(NormalBlockTimes[3], "Block 8"),
	},
}

func getTimesFromSummary(str string) [][]string {
	for last, times := range Times {
		if strings.HasSuffix(str, last) {
			return times
		}
	}

	return nil
}

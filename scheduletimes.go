package main

import (
	"strings"
)

var NormalBlockTimes = [][]string{
	{
		"07:39:00 AM",
		"09:15:00 AM",
		"Block 1/5",
	},
	{
		"09:20:00 AM",
		"10:56:00 AM",
		"Block 2/6",
	},
	{
		"11:01:00 AM",
		"01:04:00 PM",
		"Block 3/7",
	},
	{
		"01:09:00 PM",
		"02:45:00 PM",
		"Block 4/8",
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
	"1-4)": NormalBlockTimes,
	"5-8)": NormalBlockTimes,
}

func getTimesFromSummary(str string) [][]string {
	for last, times := range Times {
		if strings.HasSuffix(str, last) {
			return times
		}
	}

	return nil
}

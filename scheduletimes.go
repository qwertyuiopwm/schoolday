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
var AnchorTimes = [][]string{
	{
		"07:39:00 AM",
		"09:01:00 AM",
	},
	{
		"09:06:00 AM",
		"10:28:00 AM",
	},
	{
		"10:33:00 AM",
		"12:23:00 PM",
	},
	{
		"12:28:00 PM",
		"01:18:00 PM",
		"Anchor",
	},
	{
		"01:23:00 PM",
		"02:45:00 PM",
	},
}
var AssemblyTimes = [][]string{
	{
		"07:39:00 AM",
		"09:01:00 AM",
	},
	{
		"09:06:00 AM",
		"10:28:00 AM",
	},
	{
		"10:33:00 AM",
		"12:23:00 PM",
	},
	{
		"12:28:00 PM",
		"01:55:00 PM",
	},
	{
		"02:00:00 PM",
		"02:45:00 PM",
		"Assembly",
	},
}

var Times = map[string][][]string{
	"1-4)": {
		append(NormalBlockTimes[0], "Block 1"),
		append(NormalBlockTimes[1], "Block 2"),
		append(NormalBlockTimes[2], "Block 3"),
		append(NormalBlockTimes[3], "Block 4"),
	},
	"1-4) w/ Anchor": {
		append(AnchorTimes[0], "Block 1"),
		append(AnchorTimes[1], "Block 2"),
		AnchorTimes[2],
		append(AnchorTimes[3], "Block 3"),
		append(AnchorTimes[4], "Block 4"),
	},
	"1-4) w/ Assembly": {
		append(AssemblyTimes[0], "Block 1"),
		append(AssemblyTimes[1], "Block 2"),
		append(AssemblyTimes[2], "Block 3"),
		append(AssemblyTimes[3], "Block 4"),
		AssemblyTimes[4],
	},
	"5-8)": {
		append(NormalBlockTimes[0], "Block 5"),
		append(NormalBlockTimes[1], "Block 6"),
		append(NormalBlockTimes[2], "Block 7"),
		append(NormalBlockTimes[3], "Block 8"),
	},
	"5-8) w/ Anchor": {
		append(AnchorTimes[0], "Block 5"),
		append(AnchorTimes[1], "Block 6"),
		AnchorTimes[2],
		append(AnchorTimes[3], "Block 7"),
		append(AnchorTimes[4], "Block 8"),
	},
	"5-8) w/ Assembly": {
		append(AssemblyTimes[0], "Block 5"),
		append(AssemblyTimes[1], "Block 6"),
		append(AssemblyTimes[2], "Block 7"),
		append(AssemblyTimes[3], "Block 8"),
		AssemblyTimes[4],
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

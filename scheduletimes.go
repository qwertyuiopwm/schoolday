package schoolday

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
var AllBlocksTimes = [][]string{
	{
		"07:39:00 AM",
		"08:18:00 AM",
		"Block 5",
	},
	{
		"08:23:00 AM",
		"09:02:00 AM",
		"Block 1",
	},
	{
		"09:07:00 AM",
		"09:45:00 AM",
		"Block 2",
	},
	{
		"09:50:00 AM",
		"10:28:00 AM",
		"Block 6",
	},
	{
		"10:33:00 AM",
		"12:23:00 PM",
		"Block 3",
	},
	{
		"12:28:00 PM",
		"01:10:00 PM",
		"Block 7",
	},
	{
		"01:15:00 PM",
		"01:57:00 PM",
		"Block 8",
	},
	{
		"02:02:00 PM",
		"02:45:00 PM",
		"Block 4",
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
		append(AnchorTimes[2], "Block 3"),
		AnchorTimes[3],
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
		append(AnchorTimes[2], "Block 7"),
		AnchorTimes[3],
		append(AnchorTimes[4], "Block 8"),
	},
	"5-8) w/ Assembly": {
		append(AssemblyTimes[0], "Block 5"),
		append(AssemblyTimes[1], "Block 6"),
		append(AssemblyTimes[2], "Block 7"),
		append(AssemblyTimes[3], "Block 8"),
		AssemblyTimes[4],
	},
	"8 BLOCK DAY": AllBlocksTimes,
}

func getTimesFromSummary(str string) [][]string {
	for last, times := range Times {
		if strings.HasSuffix(str, last) {
			return times
		}
	}

	return nil
}

package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/apognu/gocal"
	"github.com/getlantern/systray"
)

var calURI = "https://bhs.beltonschools.org/calendar/calendar_350_gmt.ics"
var currentEvent *gocal.Event
var TimeLayout = "03:04:05 PM"
var OutputTimeLayout = "15:04:05"
var BDayStart = "B Day"
var ADayStart = "A Day"

var TimeLeftBlockFormat = "Time left in block: %s"
var TimeLeftUntilStart = "Time left until class starts: %s"
var TimeLeftDayFormat = "Time left in day: %s"

// Downloads the BHS calendar and returns a parsed gocal object.
func downloadCalender() (*gocal.Gocal, error) {
	res, err := http.Get(calURI)
	if err != nil {
		return nil, err
	}

	c := gocal.NewParser(res.Body)
	if err := c.Parse(); err != nil {
		return nil, err
	}

	return c, nil
}

// systray main loop function.
func running() {
	closeAppButton := systray.AddMenuItem("Close App", "")
	go func() {
		<-closeAppButton.ClickedCh
		systray.Quit()
	}()
	systray.AddSeparator()
	currentBlock := systray.AddMenuItem("", "")
	tLeftBlock := systray.AddMenuItem(TimeLeftBlockFormat, "")
	tLeftDay := systray.AddMenuItem(TimeLeftDayFormat, "")

	for {
		var title = "No class today!"
		var currentBlockString = "No Block / Passing Period"
		var timeLeftBlockString = ""
		var timeLeftDayString = ""

		var currentEventFrame = currentEvent

		if currentEventFrame == nil {
			currentBlock.Hide()
			tLeftBlock.Hide()
			tLeftDay.Hide()
		}

		currentDateTimeString := time.Now().Format(TimeLayout)
		currentParsed, err := time.Parse(TimeLayout, currentDateTimeString)
		if err != nil {
			fmt.Println(err)
		}

		if currentEventFrame != nil {
			title = currentEventFrame.Summary

			currentTimes := getTimesFromSummary(currentEventFrame.Summary)

			for _, times := range currentTimes {
				startParsed, err := time.Parse(TimeLayout, times[0])
				if err != nil {
					fmt.Println(err)
					return
				}
				endParsed, err := time.Parse(TimeLayout, times[1])
				if err != nil {
					fmt.Println(err)
					return
				}
				if timeLeftBlockString == "" && currentParsed.Before(startParsed) {
					durationLeft := startParsed.Sub(currentParsed)
					parsedDurationLeft := time.Time{}.Add(durationLeft)
					timeLeftBlockString = fmt.Sprintf(TimeLeftUntilStart, parsedDurationLeft.Format(OutputTimeLayout))
				}
				if currentParsed.Before(startParsed) || currentParsed.After(endParsed) {
					continue
				}

				if len(times) > 2 {
					currentBlockString = fmt.Sprintf("Current Block: %s", times[2])

					durationLeft := endParsed.Sub(currentParsed)
					parsedDurationLeft := time.Time{}.Add(durationLeft)
					timeLeftBlockString = fmt.Sprintf(TimeLeftBlockFormat, parsedDurationLeft.Format(OutputTimeLayout))
				}
			}

			if currentTimes != nil {
				endDayParsed, err := time.Parse(TimeLayout, currentTimes[len(currentTimes)-1][1])
				if err != nil {
					fmt.Println(err)
					continue
				}

				durationLeftDay := endDayParsed.Sub(currentParsed)
				parsedDurationLeft := time.Time{}.Add(durationLeftDay)
				timeLeftDayString = fmt.Sprintf(TimeLeftDayFormat, parsedDurationLeft.Format(OutputTimeLayout))
			}
		}

		// Hide items accordingly.
		if timeLeftBlockString == "" {
			tLeftBlock.Hide()
			tLeftDay.Hide()
		} else {
			currentBlock.Show()
			tLeftBlock.Show()
			tLeftDay.Show()
		}

		// Set titles for each item.
		systray.SetTitle(title)
		currentBlock.SetTitle(currentBlockString)
		tLeftBlock.SetTitle(timeLeftBlockString)
		tLeftDay.SetTitle(timeLeftDayString)

		time.Sleep(time.Second * 1)
	}
}

// Set the current event by iterating over calendar events and checking if the current day is the same as the event day.
func setCurrentEvent(cal gocal.Gocal) (retry bool) {
	events := cal.Events
	var newEvent *gocal.Event
	for i := len(events) - 1; i >= 0; i-- {
		e := events[i]

		if e.Start.Day() != time.Now().Day() || e.Start.Month() != time.Now().Month() {
			continue
		}
		
		if strings.HasPrefix(e.Summary, ADayStart) || strings.HasPrefix(e.Summary, BDayStart) {
			newEvent = &e
			break
		}
	}

	currentEvent = newEvent
	return currentEvent == nil
}

func end() {

}

func main() {
	var cal *gocal.Gocal
	cal, err := downloadCalender()
	if err != nil {
		fmt.Println(err)
	}

	go func() {
		for {
			if cal == nil {
				fmt.Println("No calender found")
				newCal, _ := downloadCalender()
				cal = newCal
				continue
			}
			shouldRetry := setCurrentEvent(*cal)

			if shouldRetry {
				cal = nil
				time.Sleep(time.Second * 1)
			}

			continue
		}
	}()

	systray.Run(running, end)
}

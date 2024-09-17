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
		if currentEvent == nil {
			continue
		}

		currentDateTimeString := time.Now().Format(TimeLayout)
		currentParsed, err := time.Parse(TimeLayout, currentDateTimeString)
		if err != nil {
			fmt.Println(err)
		}

		systray.SetTitle(currentEvent.Summary)

		currentTimes := getTimesFromSummary(currentEvent.Summary)
		currentBlockString := "No Block / Passing Period"
		timeLeftBlockString := ""

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

			currentBlockString = fmt.Sprintf("Current Block: %s", times[2])

			durationLeft := endParsed.Sub(currentParsed)
			parsedDurationLeft := time.Time{}.Add(durationLeft)
			timeLeftBlockString = fmt.Sprintf(TimeLeftBlockFormat, parsedDurationLeft.Format(OutputTimeLayout))
		}
		tLeftBlock.SetTitle(timeLeftBlockString)
		if timeLeftBlockString == "" {
			tLeftBlock.Hide()
			tLeftDay.Hide()
		} else {
			tLeftBlock.Show()
			tLeftDay.Show()
		}

		currentBlock.SetTitle(currentBlockString)

		if currentTimes != nil {
			endDayParsed, err := time.Parse(TimeLayout, currentTimes[len(currentTimes)-1][1])
			if err != nil {
				fmt.Println(err)
				continue
			}

			durationLeftDay := endDayParsed.Sub(currentParsed)
			parsedDurationLeft := time.Time{}.Add(durationLeftDay)
			tLeftDay.SetTitle(fmt.Sprintf(TimeLeftDayFormat, parsedDurationLeft.Format(OutputTimeLayout)))
		}

		time.Sleep(time.Second * 1)
	}
}

func setCurrentEvent(cal gocal.Gocal) (retry bool) {
	events := cal.Events
	currentEvent = nil
	for i := len(events) - 1; i >= 0; i-- {
		e := events[i]

		if e.Start.Day() != time.Now().Day() || e.End.Month() != time.Now().Month() {
			continue
		}
		if strings.HasPrefix(e.Summary, ADayStart) || strings.HasPrefix(e.Summary, BDayStart) {
			currentEvent = &e
			break
		}
	}

	if currentEvent == nil {
		downloadCalender()
		return true
	}

	return false
}

func end() {

}

func main() {
	cal, err := downloadCalender()
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		for {
			if setCurrentEvent(*cal) {
				cal, err = downloadCalender()
				if err != nil {
					fmt.Println(err)
				}

				return
			}

			time.Sleep(time.Second * 1)
		}
	}()

	systray.Run(running, end)
}

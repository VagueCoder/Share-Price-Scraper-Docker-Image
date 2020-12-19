package main

import (
	"fmt"
	"time"
	"strings"
)

func main() {
	urls := get_URLs()
	var line string
	var count int
	var dtlayout string = "02-01-2006 3:04:05 PM"
	var start, end time.Time
	var flag bool = true
	var divider string = strings.Repeat("-", 27)

	for round:=0; true; round++ {
		start = time.Now()
		if flag {
			writeStat(flag, fmt.Sprintf("%s>> Scraping Just Started at: %s <<%s\n", divider, start.Format(dtlayout), divider))
			flag = false
		} else {
			writeStat(flag, fmt.Sprintf("Round Count: %d | Useful Records: %d | Round Start: %s | Round End: %s\n", round, count, start.Format(dtlayout), end.Format(dtlayout)))
		}
		count = 0
		for j, url := range urls {
			line = fmt.Sprintf("Rounds Completed: %d | Useful Records: %d | Current Round's Status: (%d/%d)", round, count, j+1, len(urls))
			fmt.Printf(line)
			data, err := getData(url)
			if err == nil{
				count++
				writeJSON(data)
			}
			fmt.Printf("\r%v\r", strings.Repeat(" ", len(line)))
		}
		end = time.Now()
	}
}
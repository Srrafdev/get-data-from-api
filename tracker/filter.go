package box

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func FilterByLocaton(dataArtist []Data_Execute, loca string) []Data_Execute {
	var filter []Data_Execute
	for _, artist := range dataArtist {
		for _, valLoca := range artist.Locations {
			if valLoca == loca {
				filter = append(filter, artist)
			}
		}

	}
	return filter
}

func LenData(loca []Data_Execute) []string {
	uniqueMap := make(map[string]bool)
	var filter []string

	for _, valLoca := range loca {
		for _, location := range valLoca.Locations {
			if !uniqueMap[location] {
				uniqueMap[location] = true
				filter = append(filter, location)
			}
		}
	}
	return filter
}

// Function to parse date strings into time.Time objects
func parseDate(dateStr string) (time.Time, error) {

	return time.Parse(time.DateOnly, dateStr) //yy-mm-dd
}

func convertDate(date string) string {
	a := strings.Split(date, "-")
	date = a[2] + "-" + a[1] + "-" + a[0]
	return date
}

// Function to filter events by a date range
func FilterByFirstAlbum(events []Data_Execute, minDateStr, maxDateStr string) []Data_Execute {
	// Parse the min and max date strings
	minDate, err := parseDate(minDateStr)
	if err != nil {
		fmt.Println("Error parsing min date:", err)
		return nil
	}

	maxDate, err := parseDate(maxDateStr)
	if err != nil {
		fmt.Println("Error parsing max date:", err)
		return nil
	}

	var filteredEvents []Data_Execute
	for _, artis := range events {
		// Parse the artis date string
		eventDate, err := parseDate(convertDate(artis.FirstAlbum))
		if err != nil {
			fmt.Println("Error parsing artis date:", err)
			continue
		}

		// Check if the artis date is within the min and max date range
		if (eventDate.Equal(minDate) || eventDate.After(minDate)) && (eventDate.Equal(maxDate) || eventDate.Before(maxDate)) {
			filteredEvents = append(filteredEvents, artis)
		}
	}
	return filteredEvents
}

func FilterByCreationYear(dataArtists []Data_Execute, min, max string) []Data_Execute {
	var filter []Data_Execute
	minI, _ := strconv.Atoi(min)
	maxI, _ := strconv.Atoi(max)
	for _, artist := range dataArtists {
		if artist.CreationDate >= minI && artist.CreationDate <= maxI{
			filter = append(filter, artist)
		}

	}
	return filter
}

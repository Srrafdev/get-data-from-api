package box

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)
/*Seattle-usa and washington-usa: Seattle is a city in Washington state, both in the USA.
los_angeles-usa and california-usa: Los Angeles is a city in California, both in the USA.
new_york-usa and new_york-usa: New York City is a city in New York state, both in the USA.
san_francisco-usa and california-usa: San Francisco is a city in California, both in the USA.
houston-usa and texas-usa: Houston is a city in Texas, both in the USA.
boston-usa and massachusetts-usa: Boston is a city in Massachusetts, both in the USA.
miami-usa and florida-usa: Miami is a city in Florida, both in the USA.
chicago-usa and illinois-usa: Chicago is a city in Illinois, both in the USA.
atlanta-usa and georgia-usa: Atlanta is a city in Georgia, both in the USA.
seattle-usa and washington-usa: Seattle is a city in Washington state, both in the USA.*/
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

func FilterByNMembers(dataArtist []Data_Execute, memebers []string)[]Data_Execute{
	var filter []Data_Execute
	for _, artist := range dataArtist{
		for _, M := range memebers{
			NM, _ := strconv.Atoi(M)

			if len(artist.Members) == NM{
				filter = append(filter, artist)
			}
		}
	}
	return filter
}

package box

import (
	"strconv"
	"strings"
)

func Search(data []Data_Execute, target string) []Data_Execute {
	var results []Data_Execute
	if target == "" {
		return results
	}

	lowerTarget := strings.ToLower(target)

	for _, artist := range data {
		if containsMatch(artist, lowerTarget) {
			results = append(results, artist)
		}
	}

	return results
}

func containsMatch(artist Data_Execute, target string) bool {
	if strings.Contains(strings.ToLower(artist.Name), target) {
		return true
	}

	if strings.Contains(strings.ToLower(artist.FirstAlbum), target) {
		return true
	}

	if strings.Contains(strconv.Itoa(artist.CreationDate), target) {
		return true
	}

	for _, member := range artist.Members {
		if strings.Contains(strings.ToLower(member), target) {
			return true
		}
	}

	for _, location := range artist.Locations {
		if strings.Contains(strings.ToLower(location), target) {
			return true
		}
	}

	return false
}

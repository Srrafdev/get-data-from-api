package box

import (
	"fmt"
	"strconv"
	"strings"
)

var ArtistSuggest []Data_Execute

func Search(dataArtist []Data_Execute, target string) []Data_Execute {
	var results []Data_Execute
	if target == "" {
		return results
	}

	lowerTarget := strings.ToLower(target)

	for _, artist := range dataArtist {
		match := false

		if strings.Contains(strings.ToLower(artist.Name), lowerTarget) {
			match = true
		} else if strings.Contains(strings.ToLower(artist.FirstAlbum), lowerTarget) {
			match = true
		} else if strings.Contains(strconv.Itoa(artist.CreationDate), lowerTarget) {
			match = true
		} else {
			for _, member := range artist.Members {
				if strings.Contains(strings.ToLower(member), lowerTarget) {
					match = true
					break
				}
			}
		}

		if !match {
			for _, loca := range artist.Locations {
				if strings.Contains(strings.ToLower(loca), lowerTarget) {
					fmt.Println(loca)
					match = true
					break
				}
			}
		}

		if match {
			results = append(results, artist)
		}
	}

	return results
}

type suggest struct {
	name         []string
	member       []string
	Locations    []string
	FirstAlbum   []string
	CreationDate []int
}

func SuggestionSearchAPI(target string) suggest {
	var result suggest
	lowerTarget := strings.ToLower(target)

	for _, artist := range ArtistSuggest {
		if strings.Contains(strings.ToLower(artist.Name), lowerTarget) {
			result.name = append(result.name, artist.Name)

		}else if strings.Contains(strings.ToLower(artist.FirstAlbum), lowerTarget) {
			result.FirstAlbum = append(result.FirstAlbum, artist.FirstAlbum)

		}else if strings.Contains(strconv.Itoa(artist.CreationDate), lowerTarget) {
			result.CreationDate = append(result.CreationDate, artist.CreationDate)

		}else{
			for _, member := range artist.Members {
				if strings.Contains(strings.ToLower(member), lowerTarget) {
					result.member = append(result.member, member)
					break
				}
			}
		}
		
		for _, loca := range artist.Locations {
				if strings.Contains(strings.ToLower(loca), lowerTarget) {
					result.Locations = append(result.Locations, loca)
				break
			}
		}
		
	}
	fmt.Println(result)
	return result
}

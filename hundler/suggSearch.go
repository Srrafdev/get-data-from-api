package box

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	box "box/tracker"
)

var ArtistSuggest []box.Data_Execute

type Suggest struct {
	Name         []string `json:"name"`
	Member       []string `json:"member"`
	Locations    []string `json:"locations"`
	FirstAlbum   []string `json:"first_album"`
	CreationDate []string `json:"creation_date"`
}
type target struct {
	Text string `json:"target"`
}

func SuggestionSearchAPI(res http.ResponseWriter, req *http.Request) {
	var suggestions Suggest
	var target target
	res.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(req.Body).Decode(&target)
	if err != nil {
		fmt.Println("ERROR DECODE :", err)
		return
	}
	if target.Text == "" {
		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode("")
		return
	}

	lowerTarget := strings.ToLower(target.Text)

	for _, artist := range ArtistSuggest {
		if strings.Contains(strings.ToLower(artist.Name), lowerTarget) {
			suggestions.Name = append(suggestions.Name, strconv.Itoa(artist.Id)+"++"+artist.Name)
		}

		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), lowerTarget) {
				suggestions.Member = append(suggestions.Member, strconv.Itoa(artist.Id)+"++"+member)
			}
		}

		for _, location := range artist.Locations {
			if strings.Contains(strings.ToLower(location), lowerTarget) {
				suggestions.Locations = append(suggestions.Locations, strconv.Itoa(artist.Id)+"++"+location+"++"+artist.Name)
			}
		}

		if strings.Contains(strings.ToLower(artist.FirstAlbum), lowerTarget) {
			suggestions.FirstAlbum = append(suggestions.FirstAlbum, strconv.Itoa(artist.Id)+"++"+artist.FirstAlbum+"++"+artist.Name)
		}

		if strings.Contains(strconv.Itoa(artist.CreationDate), lowerTarget) {
			suggestions.CreationDate = append(suggestions.CreationDate, strconv.Itoa(artist.Id)+"++"+strconv.Itoa(artist.CreationDate)+"++"+artist.Name)
		}
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(suggestions)
}

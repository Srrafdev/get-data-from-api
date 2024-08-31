package box

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const URL = "https://groupietrackers.herokuapp.com/api"

func Decode(mystruct interface{}, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to get URL: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(mystruct); err != nil {
		return fmt.Errorf("failed to decode JSON: %v", err)
	}

	return nil
}

func (api *Api)SaveURL() {
	if err := Decode(&api, URL); err != nil {
		fmt.Println(err)
		return
	}
}

func FillMoreDatae(id string) (*Data_Execute, error) {
	var api Api
	api.SaveURL()
	
	lis := make([]string, 4)
	
	// Populate the slice using a for loop
	urls := []string{api.Artists, api.Dates, api.Locations, api.Relation}
	var exeData Data_Execute

	idx, err := strconv.Atoi(id)
	if err != nil || idx > len(api.Artists)+3 || idx < 1 {
		return nil, fmt.Errorf("id is not valide")
	}

	for i := 0; i < len(lis); i++ {
		lis[i] = urls[i] + "/" + id
		Decode(&exeData, lis[i])
	}
	return &exeData, nil
}

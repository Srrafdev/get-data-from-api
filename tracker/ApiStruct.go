package box

type Api struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relation  string `json:"relation"`
}

type Data_Execute struct {
	Id            int                 `json:"id"`
	Image         string              `json:"image"`
	Name          string              `json:"name"`
	Members       []string            `json:"members"`
	CreationDate  int                 `json:"creationDate"`
	FirstAlbum    string              `json:"firstAlbum"`
	Locations     []string            `json:"locations"`
	Dates         []string            `json:"dates"`
	DatesLocation map[string][]string `json:"datesLocations"`
}

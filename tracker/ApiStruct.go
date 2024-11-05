package box

type Api struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relation  string `json:"relation"`
}

type Data_Execute struct {
	Error         string
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
type TempStruct struct {
	Index []Data_Execute `json:"index"`
}

type Execute struct {
	Error string
	//
	LocaAll      []string
	SelectedLoca string
	//
	SaveMinDeta string
	SaveMaxDeta string
	//
	SaveMinCreat string
	SaveMaxCreat string

	SaveNM []string

	DataEX []Data_Execute
}

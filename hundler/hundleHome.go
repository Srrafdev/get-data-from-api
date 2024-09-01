package box

import (
	"bytes"
	"net/http"
	"text/template"

	box "box/tracker"
)

// hundle err in home page
func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not Found 404", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed 405", http.StatusMethodNotAllowed)
		return
	}

	tmp, err := template.ParseFiles("./website/pages/index.html")
	if err != nil {
		http.Error(w, "Internal Server 500", http.StatusInternalServerError)
		return
	}

	var dataArtist []box.Data_Execute
	var api box.Api
	api.SaveURL()
	//
	var tempData box.TempStruct
	box.Decode(&tempData, api.Locations)

	dataArtist = tempData.Index
	box.Decode(&dataArtist, api.Artists)
	locaAll := box.LenData(dataArtist)
	////
	if rr := r.ParseForm(); rr != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}
	// filter locations
	loca := r.FormValue("loca")
	if loca != "" && loca != "all" {
		DT := box.FilterByLocation(dataArtist, loca)
		dataArtist = DT
	}
	// filter first album
	minFL := r.FormValue("mindate")
	maxFL := r.FormValue("maxdeta")
	if minFL != "" && maxFL != "" {
		DTD := box.FilterByFirstAlbum(dataArtist, minFL, maxFL)
		dataArtist = DTD
	}
	// filter Creation Date
	minCD := r.FormValue("minDaCre")
	maxCD := r.FormValue("maxDaCre")
	if minCD != "" && maxCD != "" {
		DTCD := box.FilterByCreationYear(dataArtist, minCD, maxCD)
		dataArtist = DTCD
	}
	if minCD == "" {
		minCD = "1958"
	}
	if maxCD == "" {
		maxCD = "2015"
	}
	// filter N members
	members := []string{r.FormValue("m1"), r.FormValue("m2"), r.FormValue("m3"), r.FormValue("m4"), r.FormValue("m5"), r.FormValue("m6"), r.FormValue("m7"), r.FormValue("m8")}
	if !box.Isnill(members) {
		DTNM := box.FilterByNMembers(dataArtist, members)
		dataArtist = DTNM
	}

	exec := box.Execute{
		LocaAll:      locaAll,
		SelectedLoca: loca,
		SaveMinDeta:  minFL,
		SaveMaxDeta:  maxFL,
		SaveMinCreat: minCD,
		SaveMaxCreat: maxCD,
		SaveNM:       members,
		DataEX:       dataArtist,
	}

	buf := &bytes.Buffer{}
	if err := tmp.Execute(buf, exec); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(buf.Bytes())
}

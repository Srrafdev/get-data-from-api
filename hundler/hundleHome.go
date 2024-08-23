package box

import (
	box "box/tracker"
	"fmt"
	"net/http"
	"text/template"
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

	box.SaveURL(&api)
	fmt.Println(api)
	
	var tempData box.TempStruct
	box.Decode(&tempData, api.Locations)
	
	dataArtist = tempData.Index
	box.Decode(&dataArtist, api.Artists)

	r.ParseForm()
	DT := box.FilterLocaton(dataArtist,r.FormValue("loca"))

	if err := tmp.Execute(w, DT); err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server 500", http.StatusInternalServerError)
		return
	}
}

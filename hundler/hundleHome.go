package box

import (
	box "box/tracker"
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
	box.Decode(&dataArtist, api.Artists)

	var filter []box.Data_Execute
	for _, artist := range dataArtist {
		if artist.CreationDate > 2000 && artist.CreationDate < 2010 {
			filter = append(filter, artist)
		}
	}

	if err := tmp.Execute(w, filter); err != nil {
		http.Error(w, "Internal Server 500", http.StatusInternalServerError)
		return
	}
}

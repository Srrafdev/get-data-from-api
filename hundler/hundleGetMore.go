package box

import (
	box "box/tracker"
	"net/http"
	"text/template"
)

func GetMore(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/GetMore" {
		http.Error(w, "Not Found 404", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed 405", http.StatusMethodNotAllowed)
		return
	}

	tmp, errT := template.ParseFiles("./website/pages/GetMore.html")
	if errT != nil {
		http.Error(w, "Internal Server 500", http.StatusInternalServerError)
		return
	}
	r.ParseForm()
	id := r.FormValue("submit")
	data, errS := box.FillMoreDatae(id)
	if errS != nil {
		http.Error(w, "Bad request 400", http.StatusBadRequest)
		return
	}
	if err := tmp.Execute(w, &data); err != nil {
		http.Error(w, "Internal Server 500", http.StatusInternalServerError)
		return
	}
}

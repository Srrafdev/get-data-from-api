package box

import (
	"net/http"
	"text/template"

	box "box/tracker"
)

func GetMore(w http.ResponseWriter, r *http.Request) {
	tmp, errT := template.ParseFiles("./website/pages/GetMore.html")

	if r.URL.Path != "/GetMore" {
		w.WriteHeader(http.StatusNotFound)
		data := box.Execute{
			Error: "Not Found 404",
		}
		tmp.Execute(w, data)
		return
	}

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		data := box.Execute{
			Error: "Method Not Allowed 405",
		}
		tmp.Execute(w, data)
		return
	}

	if errT != nil {
		w.WriteHeader(http.StatusInternalServerError)
		data := box.Execute{
			Error: "Internal Server Error 500",
		}
		tmp.Execute(w, data)
		return
	}
	r.ParseForm()
	id := r.FormValue("submit")
	data, errS := box.FillMoreDatae(id)
	if errS != nil {
		w.WriteHeader(http.StatusBadRequest)
		data := box.Execute{
			Error: "Bad request 400",
		}
		tmp.Execute(w, data)
		return
	}
	if err := tmp.Execute(w, &data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		data := box.Execute{
			Error: "Internal Server Error 500",
		}
		tmp.Execute(w, data)
		return
	}
}

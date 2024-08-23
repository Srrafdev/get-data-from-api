package main

import (
	box "box/hundler"
	"fmt"
	"net/http"
)

func main() {
	//http.Handle("/styles.css", http.FileServer(http.Dir("./website/style/")))
	//http.Handle("/style.css/", http.StripPrefix("/style/", http.FileServer(http.Dir("website/style"))))
	cssDir := http.FileServer(http.Dir("./website/style"))
	http.Handle("/style/", http.StripPrefix("/style/", cssDir))
	//http.Handle("/styles.css", http.FileServer(http.FS(cssFiles)))

	http.HandleFunc("/", box.Home)
	http.HandleFunc("/GetMore", box.GetMore)

	fmt.Println("Server started on port 8080...")
	if err := (http.ListenAndServe(":8080", nil)); err != nil {
		fmt.Println(err)
		return
	}
}

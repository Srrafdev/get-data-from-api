package main

import (
	box "box/hundler"
	limit "box/ratelimit"
	"fmt"
	"net/http"
)

func main() {
	cssDir := http.FileServer(http.Dir("./website/style"))
	http.Handle("/style/", http.StripPrefix("/style/", cssDir))
	//http.Handle("/styles.css", http.FileServer(http.Dir("./website/style/")))
	//http.Handle("/styles.css", http.FileServer(http.FS(cssFiles)))
	
	limiter := limit.NewLimiter(2)
	http.HandleFunc("/", limit.RateLimitMiddleware(box.Home, limiter))
	http.HandleFunc("/GetMore", limit.RateLimitMiddleware(box.GetMore, limiter))

	fmt.Println("Server started on port 8080...")
	if err := (http.ListenAndServe(":8080", nil)); err != nil {
		fmt.Println(err)
		return
	}
}

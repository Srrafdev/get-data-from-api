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

	limiter := limit.NewLimiter(2)
	http.HandleFunc("/", limit.RateLimitMiddleware(box.Home, limiter))
	http.HandleFunc("/GetMore", limit.RateLimitMiddleware(box.GetMore, limiter))
	http.HandleFunc("/api/search", box.SuggestionSearchAPI)

	fmt.Println("Server started on port 8080...")
	if err := (http.ListenAndServe(":8080", nil)); err != nil {
		fmt.Println(err)
		return
	}
}

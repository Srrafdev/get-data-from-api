package main

import (
	box "box/hundler"
	limit "box/ratelimit"
	"fmt"
	"net/http"
)

func main() {
	limiter := limit.NewLimiter(2)
	//http.Handle("/styles.css", http.FileServer(http.Dir("./website/style/")))
	//http.Handle("/style.css/", http.StripPrefix("/style/", http.FileServer(http.Dir("website/style"))))
	cssDir := http.FileServer(http.Dir("./website/style"))
	http.Handle("/style/", http.StripPrefix("/style/", cssDir))
	//http.Handle("/styles.css", http.FileServer(http.FS(cssFiles)))

	http.HandleFunc("/", limit.RateLimitMiddleware(box.Home, limiter))
	http.HandleFunc("/GetMore", limit.RateLimitMiddleware(box.GetMore, limiter))

	fmt.Println("Server started on port 8090...")
	if err := (http.ListenAndServe(":8090", nil)); err != nil {
		fmt.Println(err)
		return
	}
}

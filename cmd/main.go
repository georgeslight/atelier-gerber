package main

import (
	"fmt"
	"github.com/a-h/templ"
	"kristin-gerber/views/layouts"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	component := layouts.BaseLayout()
	http.Handle("/", templ.Handler(component))
	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}

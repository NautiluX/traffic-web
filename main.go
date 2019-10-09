package main

import (
  "log"
  "github.com/NautiluX/traffic-web/traffic"
	"net/http"
)

func main() {
  log.Println("Hello again!")

  fs := http.FileServer(http.Dir("static"))
  http.HandleFunc("/backend/", serveBackend)
  http.Handle("/", http.StripPrefix("/", fs))

  log.Println("Listening...")
  http.ListenAndServe(":3000", nil)
}


func serveBackend(w http.ResponseWriter, r *http.Request) {
	var color string
	colors, _ := r.URL.Query()["color"]
	if len(colors) == 1 {
		color = colors[0]
	}
	switch color {
	case "red":
		traffic.Red()
	case "yellow":
		traffic.Yellow()
	case "green":
		traffic.Green()
	case "off":
		traffic.Off()
	}
}

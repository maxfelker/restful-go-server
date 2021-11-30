package restfulgoserver

import (
	"log"
	"net/http"
)

type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

type Routes []Route

func Start(PORT string) {
	var port = ":" + PORT
	logMessage := "Listen for requests at http://localhost" + port
	log.Println(logMessage)
	log.Fatal(http.ListenAndServe(port, nil))
}

func RegisterRoutes(routes []Route) {
	log.Println("Registering routes...")
	for _, route := range routes {
		http.HandleFunc(route.Path, route.Handler)
		log.Println(route.Method + " " + route.Path)
	}
}

func Respond(writer http.ResponseWriter, response []byte) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(response)
}

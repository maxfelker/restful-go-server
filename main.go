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
		http.HandleFunc(route.Path, preHandler(route.Method, route.Handler))
		log.Println(route.Method + " " + route.Path)
	}
}

func Respond(writer http.ResponseWriter, response []byte) {
	writer.Write(response)
}

func preHandler(method string, handler http.Handler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		writer.Header().Set("Access-Control-Allow-Credentials", "true")
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Methods", method+", OPTIONS")
		writer.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Endcoding, Content-Type, Content-Length, Authorization, X-CSRF-token")
		if request.Method == "OPTIONS" {
			return
		}
		handler.ServeHTTP(writer, request)
	}
}

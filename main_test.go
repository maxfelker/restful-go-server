package restfulgoserver

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"testing"
)

func getRoutes() []reflect.Value {
	httpMux := reflect.ValueOf(http.DefaultServeMux).Elem()
	return httpMux.FieldByIndex([]int{1}).MapKeys()
}

func verifyRouteCount(routesToTest Routes, registeredRoutes []reflect.Value) {
	var passedRouteCount = len(routesToTest)
	var routesRegisteredCount = len(registeredRoutes)
	if routesRegisteredCount != passedRouteCount {
		log.Panicln("There should be " + fmt.Sprint(passedRouteCount) + " registered but found " + fmt.Sprint(routesRegisteredCount))
	}
}

func TestRegisterRoutes(t *testing.T) {
	fmt.Println("When registering the routes...")
	var routes = Routes{
		{
			Path:   "/user",
			Method: "GET",
			Handler: func(w http.ResponseWriter, r *http.Request) {
			},
		},
		{
			Path:   "/user/1",
			Method: "DELETE",
			Handler: func(w http.ResponseWriter, r *http.Request) {
			},
		},
	}
	RegisterRoutes(routes)
	registeredRoutes := getRoutes()
	fmt.Println("...the number of routes should be equal")
	verifyRouteCount(routes, registeredRoutes)
	fmt.Println("...the supplied route paths match the regisered routes")
	for key, value := range registeredRoutes {
		var pathToTest = routes[key].Path
		var registeredPath = value.String()
		if pathToTest != registeredPath {
			log.Panicln(pathToTest + " was supplied but expected " + registeredPath)
		}
	}
}

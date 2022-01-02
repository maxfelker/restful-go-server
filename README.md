RESTful Go Server
---
A simple HTTP server built in Go

# Using in your project 

The restful server is simple to get setup. In your `main.go` file, use the following example code to get going:

```go
import (
   server "github.com/mw-felker/restful-go-server"
)

func main() {
	const PORT = "8000"
	var routes = server.Routes{
		{
			Path:    "/users",
			Method:  "GET",
			Handler: getUser,
		}
	}
	server.RegisterRoutes(routes)
	server.Start(PORT)
}

// example of a simple get route that hits a database and returns a JSON array
func getUser(writer http.ResponseWriter, request *http.Request) {
	var users []models.Users = retrieveUsers()
	response, e := json.Marshal(users)
	if e != nil {
		log.Panic(e)
	}
	server.Respond(writer, response)
}
```

## Methods

### `Start(PORT string)`

This starts the underlying `http` server and starting listening on the supplied `PORT`. This function should be called after `RegisterRoutes` 

### `RegisterRoutes(routes []Route])`

Registers one or more `Route`s to the server 

### `Respond(writer http.ResponseWriter, response []byte)`

End the request lifecycle and respond to the calling system with a JSON payload.

# TODO

Running list of things that need to be done:

- [ ] Use go internals for http methods
- [ ] Check the server is running on the specified port 
- [ ] Test `Respond` method
- [ ] Allow response `Content-Type` to be configurable per route
- [ ] Test each route's request method
- [ ] Allow response `Access-Control-Allow-Origin` to be configurable per route

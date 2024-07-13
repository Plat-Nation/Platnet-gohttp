package main

import (
	"net/http"

	middleware "github.com/Plat-Nation/BookRecs-Middleware/core"
	complexRoute "github.com/Plat-Nation/Platnet-gohttp/internal/complexRoute"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Example route handler
func IndexHandler(mw *middleware.Middleware) http.Handler{
	// Our handler function can take whatever arguments we want, and then return a HandlerFunc, which can only have (w, r) for arguments
	// We can use a closure (anonymous function) to create the HandlerFunc right here:
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Use the log middleware, http request information will be included automatically
		// Log() takes in the request itself, a message, and then optionally however many zap fields you want to include, in  this case demo data
		mw.Log(r, "Something else happened", zap.String("username", "totallyauser"), zap.Int("numOfInfo", 1))

		// If we don't have much to log, we can keep it really simple too
		mw.Log(r, "Simple log")

		// We can also get the logger directly from the middleware object if we want to do something manually, like log an error:
		exampleErrorMsg := "Failed to complete"
		mw.LogWithLevel(zapcore.ErrorLevel, r, "An error occured", zap.String("errorMessage", exampleErrorMsg), zap.String("someHelpfulInfoLikeAUsername", "example_username"), zap.Bool("LoggedIn", true))

		// We can use a switch statement like an if statement to run different code in cases where the r.Method is different
		// If you want to split these out into their own functions you can use a serveMux and include the route in the url you assign to the handler, like:
		// mux := http.NewServeMux()
		// mux.Handler("POST /", mw.All(someHandlerFunc(mw)))
		switch r.Method {
		case "GET":
			// For simple routes you can include all the code here

			// Set a custom header
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("X-Custom-Header", "MyValue")

			// Set the status code
			w.WriteHeader(http.StatusOK)

			// Write the response body, `` are used to make a multiline string that is a raw string literal, so you can also include characters like \n without escaping them
			responseBody := `{"message": "Hello, world!"}`
			w.Write([]byte(responseBody))
		case "POST":
			// For more complex routes, you can point to a separate function or even another file like this, although a mux like mentioned above is easier
			// We import the handler from complexRoute, and then we pass (mw) to the handler and (w, r) to the inner function
			complexRoute.PostHandler(mw)(w, r)
		default:
		}
	})
}

func HandlePost(mw *middleware.Middleware) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// This function will handle any requests from "POST /some-route/{id}"
		// We can get the value in the id spot like this:
		idString := r.PathValue("id")
		w.Write([]byte(idString))
	})
}

func main() {

	mw, err := middleware.Init(true, true)
	if err != nil {
		panic(err)
	}

	// Create a router and serve routes from routes.go
	mux := http.NewServeMux()
	AddRoutes(mux, mw)

	http.ListenAndServe(":8080", mux)
}


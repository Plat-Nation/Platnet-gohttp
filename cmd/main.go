package main

import (
	"net/http"

	"github.com/Plat-Nation/BookRecs-Middleware/auth"
	"github.com/Plat-Nation/BookRecs-Middleware/log"
	complexRoute "github.com/Plat-Nation/Platnet-gohttp/internal/complexRoute"
	"go.uber.org/zap"
)

// Initialize a sugared zap logger, which is easy to use
func initLogger() *zap.SugaredLogger {
	// We create a Production logger in zap, which writes structured JSON logs
	// If you want to just log something for testing, you can use NewDevelopment() to write the logs in a Human-friendly format
	// All logs go to stderr
	// Must() is used to panic if an zap.newProduction() returns an error, rather than handling it
	logger := zap.Must(zap.NewProduction())
	// We can use Sugar() to wrap our logger in an easier to use API, although it won't be quite as fast
	// This allows us to use logger.Infof("log"), rather than structuring the entire log and all key value pairs
	sugar := logger.Sugar()
	return sugar
}

// Example route handler
func indexHandler(logger *zap.SugaredLogger) http.HandlerFunc{
	// Our handler function can take whatever arguments we want, and then return a HandlerFunc, which can only have (w, r) for arguments
	// We can use a closure (anonymous function) to create the HandlerFunc right here:
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Use the sugared logger
		logger.Infof("Something happened")

		// Use the log middleware, http request information will be included automatically
		// Log() takes in the request itself, a message, and then however many zap fields you want to include, in  this case demo data
		log.Log(r, "Something else happened", zap.String("username", "totallyauser"), zap.Int("numOfInfo", 1))

		// We can use a switch statement like an if statement to run different code in cases where the r.Method is different
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
			// For more complex routes, you can point to a separate function or even another file like this
			// We import the handler from complexRoute, and then we pass (logger) to the handler and (w, r) to the inner function
			complexRoute.PostHandler(logger)(w, r)
		default:
		}
	})
}

func main() {
	logger := initLogger()
	// We can use "defer" here to save this for when we are done using the logger / exiting the program
	// In this case, this just makes sure any buffered logs still get printed
	defer logger.Sync()
	// Use both middleware on the route
	http.Handle("/", log.LogAll(auth.Auth(indexHandler(logger))))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logger.Errorf(err.Error())
	}
}


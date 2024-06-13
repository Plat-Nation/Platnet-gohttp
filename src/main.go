package main

import (
	"net/http"

	auth "github.com/Plat-Nation/BookRecs-Middleware/pkg/auth"
	log "github.com/Plat-Nation/BookRecs-Middleware/pkg/log"
	"go.uber.org/zap"
)

// Initialize a sugared zap logger, which is easy to use
func initLogger() *zap.SugaredLogger {
	// We create a Production logger in zap, which writes structured JSON logs
	// If you want to just log something for testing, you can use NewDevelopment() to write the logs in a Human-friendly format
	// All logs go to stderr
	// Must() is used to panic if an zap.newProduction() returns an error, rather than handling it
	logger := zap.Must(zap.NewProduction())
	// We can use "defer" here to save this for when we are done using the logger / exiting the program
	// In this case, this just makes sure any buffered logs still get printed
	defer logger.Sync()
	// We can use Sugar() to wrap our logger in an easier to use API, although it won't be quite as fast
	// This allows us to use logger.Infof("log"), rather than structuring the entire log and all key value pairs
	sugar := logger.Sugar()
	return sugar
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Use the sugared logger
	logger := initLogger()
	logger.Infof("Something happened")

	// Use the log middleware, http request information will be included automatically
	// Log() takes in the request itself, a message, and then however many zap fields you want to include, in  this case demo data
	log.Log(r, "Something else happened", zap.String("username", "totallyauser"), zap.Int("numOfInfo", 1))

	// Set a custom header
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Custom-Header", "MyValue")

	// Set the status code
	w.WriteHeader(http.StatusOK)

	// Write the response body
	responseBody := `{"message": "Hello, world!"}`
	w.Write([]byte(responseBody))
}

func main() {
	// Use both middleware on the route
	http.Handle("/", log.LogAll(auth.Auth(indexHandler)))
	http.ListenAndServe(":8080", nil)
}


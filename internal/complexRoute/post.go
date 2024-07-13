package complexRoute

import (
	"encoding/json"
	"fmt"
	"net/http"

	middleware "github.com/Plat-Nation/BookRecs-Middleware/core"
	"go.uber.org/zap"
)

// This is a custom type we'll make for the JSON body we expect in the request
// Normally with types you might just have an attribute name and type, but in this case
// we also have this `json:"name"` that we provide. This tells go what JSON keys to use for this so it
// can parse it automatically.
//
// Also note that capital letters are required for the type attributes and type name to be public and used by other parts of our code. If the JSON
// is lowercase, we have it lowercase in the `json:"name"` part on the right. Not a big fan of this decision in Go
type ExpectedRequest struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Pet   string `json:"pet"`
}

// Example of a route handler function that can parse a JSON POST request
// This also needs to be capitalized if it's meant to be public and imported, but other functions can stay lowercase
func PostHandler(mw *middleware.Middleware) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// We initialize a variable called parsedReq of our custom ExpectedRequest type
		var parsedReq ExpectedRequest
		// We can use the "encoding/json" standard library to decode the JSON body
		decoder := json.NewDecoder(r.Body)
		// Then we can attempt to decode it and convert it to an ExpectedRequest type.
		// Note that we use &parsedReq to use the address of the original parsedReq variable, rather than making a copy
		err := decoder.Decode(&parsedReq)
		if err != nil {
			mw.Logger.Error("Failed to decode JSON body", zap.Error(err))
			// In this case we'll just block the request if we can't parse it properly, but
			// you could be more flexible about the type here if that fits your use better
			http.Error(w, "Error: Bad Request", http.StatusBadRequest)
			return
		}

		// Set a custom header
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Custom-Header", "MyValue")

		// Set the status code
		w.WriteHeader(http.StatusOK)

		// Write the response body
		msg := fmt.Sprintf("Hi %s, your age is %d and your pet is %s", parsedReq.Name, parsedReq.Age, parsedReq.Pet)
		responseBody := `{"message": ` + msg + `}`
		w.Write([]byte(responseBody))
	})
}
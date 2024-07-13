package main

import (
	"net/http"

	"github.com/Plat-Nation/BookRecs-Middleware/core"
)

// We can define all of our routes in a separate routes.go file and have them added to the mux in our main.go
// The "net/http" library now allows you to include an HTTP method or route paramters like an Id that might change per user:
// https://github.com/golang/go/discussions/60227
func AddRoutes(
	mux *http.ServeMux,
	mw *core.Middleware,
) {
	mux.Handle("/", mw.AllNoAuth(IndexHandler(mw))) // We'll have this route be unauthenticated
	mux.Handle("POST /some-route/{id}",  mw.All(HandlePost(mw)))
}
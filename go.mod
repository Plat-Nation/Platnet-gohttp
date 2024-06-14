module github.com/Plat-Nation/Platnet-gohttp

go 1.22.4

require (
	github.com/Plat-Nation/BookRecs-Middleware v0.2.4
	go.uber.org/zap v1.27.0
)

require (
	github.com/golang-jwt/jwt/v4 v4.5.0 // indirect
	github.com/stretchr/testify v1.9.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
)

// Can be used to reference the middleware locally if you want to make a modification to the middleware and use it right away
// replace github.com/Plat-Nation/BookRecs-Middleware => ../BookRecs-Middleware

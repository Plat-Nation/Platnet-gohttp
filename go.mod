module github.com/Plat-Nation/Platnet-gohttp

go 1.22.4

require (
	github.com/Plat-Nation/BookRecs-Middleware v0.4.1
	go.uber.org/zap v1.27.0
)

require (
	github.com/aws/aws-sdk-go v1.54.19 // indirect
	github.com/cenkalti/backoff/v4 v4.2.1 // indirect
	github.com/golang-jwt/jwt/v4 v4.5.1 // indirect
	github.com/guregu/dynamo v1.23.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/stretchr/testify v1.9.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/sync v0.5.0 // indirect
)

// Can be used to reference the middleware locally if you want to make a modification to the middleware and use it right away
// replace github.com/Plat-Nation/BookRecs-Middleware => ../BookRecs-Middleware

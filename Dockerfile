FROM alpine:latest

WORKDIR /app
COPY . ./gohttp

EXPOSE 8080

CMD ["./gohttp"]

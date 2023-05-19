FROM golang:1.20 AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=windows go build -o app cmd/server/main.go
FROM alpine:3.14
COPY --from=builder /app .
CMD  ["./app"]
# FROM golang:1.20 
# WORKDIR /Users/User/go/go-rest-api
# COPY . .
# COPY go.mod ./
# RUN go mod download && go mod verify

# COPY . .
# RUN CGO_ENABLED=0 GOOS=windows go build -o app cmd/server/main.go

# CMD ["./app"]
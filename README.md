//
test Postman:
http://localhost:8080/api/v1/comment
raw-body {"slug":"hello",
"body":"body",
"author":"me"}


test:
go test ./... -v
go test -tags=integration ./... -v
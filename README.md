//
test Postman:
http://localhost:8080/api/v1/comment
raw-body {"slug":"hello",
"body":"body",
"author":"me"}


tests:
--=========================integration===============
task integration-test
--==========================acceptance==============
task acceptance-tests

JWT token ==> missionimpossible  (generate from www.jwt.io)
## run with hurl endpoint-test.hurl -v
GET http://localhost:8080/cakes
GET http://localhost:8080/cakes/1
GET http://localhost:8080/cakes/2
POST http://localhost:8080/cakes
{
    "name": "test-dev"
}
PUT http://localhost:8080/cakes/1
{
    "name": "test-dev-changedX"
}
DELETE http://localhost:8080/cakes/2
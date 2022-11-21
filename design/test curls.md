# Test get
`curl http://localhost:8081/read -H "Key: Hello"`
`curl http://localhost:8082/read -H "Key: Hello"`
`curl http://localhost:8083/read -H "Key: Hello"`

# Test create
`curl http://localhost:8081/create -X POST -H "Content-Type: application/json" -d '{"key":"Name", "value": [115, 116, 101, 102, 97, 110]}'`

go build ./src/TinyServer
go install ./src/TinyServer
go build -o ./bin/ServerMain.exe ./src/ServerMain.go
go build -o ./bin/TestClient.exe ./src/TestClient

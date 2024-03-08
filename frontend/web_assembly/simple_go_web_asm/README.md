## This is a demo of using golang to compile a frontend in web-assembly


Run Instructions

    GOOS=js GOARCH=wasm go build -o lib.wasm main.go

    go run server.jo

    Browser : localhost:8000  // this should render the contents in main.go


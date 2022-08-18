tag: cmd/app_run.go cmd/socket_run.go
	rm -rf build/*
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go  build -o build/web-server cmd/app_run.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go  build -o build/socket-server cmd/socket_run.go

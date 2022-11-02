linux: cmd/app_run.go cmd/socket_run.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go  build -o build/web-server cmd/app_run.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go  build -o build/socket-server cmd/socket_run.go

mac: cmd/app_run.go cmd/socket_run.go
	go build -o build/web-server-mac cmd/app_run.go
	go build -o build/socket-server-mac cmd/socket_run.go

clean:
	rm -rf build/*
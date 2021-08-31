hello:
	echo "Hello"

build:
	go build -o build/douglasie v1/main.go

compile:
	echo "Compiling for every OS and Platform"
	GOOS=windows GOARCH=amd64 go build -o build/douglasie-windows-amd64 v1/main.go
	GOOS=linux GOARCH=amd64 go build -o build/douglasie-linux-amd64 v1/main.go
	GOOS=darwin GOARCH=amd64 go build -o build/douglasie-macos-amd64 v1/main.go

	echo "Compiling CLI"
	GOOS=windows GOARCH=amd64 go build -o build/douglasie-cli-windows-amd64 v1/cli/main.go
	GOOS=linux GOARCH=amd64 go build -o build/douglasie-cli-linux-amd64 v1/cli/main.go
	GOOS=darwin GOARCH=amd64 go build -o build/douglasie-cli-macos-amd64 v1/cli/main.go
run:
	go run v1/main.go
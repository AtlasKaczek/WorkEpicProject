gomod = stankryj/WorkEpicProject

run: build
	echo "running..."
	go run main.go

build:
	echo "building..."
	go build -o .

compile:
	echo "Compiling for every OS and Platform"
	GOOS=freebsd GOARCH=386 go build -o WorkEpicProject_freebsd .
	GOOS=linux GOARCH=arm64 go build -o WorkEpicProject_linux64 .
	GOOS=linux GOARCH=arm go build -o WorkEpicProject_linux .

tidy:
	go mod tidy
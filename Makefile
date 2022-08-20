BINARY=go-discord-bot

build:
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY}-darwin main.go
	# GOARCH=amd64 GOOS=linux go build -o ${BINARY}-linux main.go
	# GOARCH=amd64 GOOS=windows go build -o ${BINARY}-windows main.go

run:
	./${BINARY}-darwin

clean:
	go clean
	rm ${BINARY}-darwin
	# rm ${BINARY}-linux
	# rm ${BINARY}-windows

go: build run clean

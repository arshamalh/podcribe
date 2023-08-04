run:
	LIBRARY_PATH=whisper C_INCLUDE_PATH=whisper go run main.go transcribe "http://google.com"

build:
	LIBRARY_PATH=whisper C_INCLUDE_PATH=whisper go build main.go

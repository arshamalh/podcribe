BUILD_NAME=podcribe

run:
	LIBRARY_PATH=whisper C_INCLUDE_PATH=whisper go run main.go transcribe "http://google.com"

build:
	LIBRARY_PATH=${PWD}/whisper.cpp C_INCLUDE_PATH=${PWD}/whisper.cpp go build -o ${BUILD_NAME} .

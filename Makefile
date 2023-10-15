# Default values, but should be overriden like `make run File=data/another.json
version=0.0.1
name=podcribe

help: # Generate list of targets with descriptions
	@grep '^.*\:\s#.*' Makefile | sed 's/\(.*\) # \(.*\)/\1 \2/' | column -t -s ":"

build: # Build the binary for current local system
	LIBRARY_PATH=${PWD}/whisper.cpp C_INCLUDE_PATH=${PWD}/whisper.cpp go build -o podcribe .

build-docker: # Build the docker image
	docker build -t ${name}:${version} .

sample-docker: # Run docker with sample transcribe command
	docker run --rm ${name}:${version} transcribe "http://google.com"

sample-dev: # Run go run command
	LIBRARY_PATH=${PWD}/whisper.cpp C_INCLUDE_PATH=${PWD}/whisper.cpp go run main.go transcribe "http://google.com"

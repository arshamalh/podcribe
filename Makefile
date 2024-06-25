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

sdr: # sample docker remote
	docker run --rm -v ${PWD}/files:/files arshamalh/${name}:latest transcribe files/${file}

sample-dev: # Run go run command
	LIBRARY_PATH=${PWD}/whisper.cpp C_INCLUDE_PATH=${PWD}/whisper.cpp go run main.go transcribe "http://google.com"

runtel:
	go run main.go start --telegram-on

# These commands works but there is no need for the code in that case?
# We expect to be consistent
# cd whisper.cpp/bindings/go
# ./build/go-whisper -model ../../../aimodels/ggml-base.en.bin "../../../files/01 Into You.wav"
# 
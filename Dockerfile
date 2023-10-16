FROM golang:1.20-bullseye as builder
WORKDIR /app
RUN apt-get update && apt-get install -y ffmpeg
RUN git clone https://github.com/ggerganov/whisper.cpp
RUN cd whisper.cpp/bindings/go && make whisper
RUN cd whisper.cpp && make base.en
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go mod edit -replace=github.com/ggerganov/whisper.cpp/bindings/go=./whisper.cpp/bindings/go
RUN make build

FROM debian:bookworm-20210816
COPY --from=builder /app/podcribe /bin/podcribe
COPY --from=builder /usr/bin/ffmpeg /bin/ffmpeg
COPY --from=builder /app/whisper.cpp/models/ggml-base.en.bin aimodels/ggml-base.en.bin
ENTRYPOINT [ "podcribe" ]

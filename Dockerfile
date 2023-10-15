FROM golang:1.20-bullseye as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN git clone https://github.com/ggerganov/whisper.cpp
RUN cd whisper.cpp/bindings/go && make whisper
RUN cd whisper.cpp && make base.en
COPY . ./
RUN go mod edit -replace=github.com/ggerganov/whisper.cpp/bindings/go=./whisper.cpp/bindings/go
RUN make build

FROM debian:bookworm-20210816
COPY --from=builder /app/podcribe /bin/podcribe
COPY --from=builder /app/whisper.cpp/models/ggml-base.en.bin aimodels/ggml-base.en.bin
ENTRYPOINT [ "podcribe" ]

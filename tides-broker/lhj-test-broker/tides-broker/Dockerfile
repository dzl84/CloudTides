FROM golang:latest

WORKDIR /build

# ENV GOPROXY https://goproxy.cn
COPY . .
RUN go mod download
RUN cd cmd && \
    GOARCH=amd64 GOOS=linux go build -o /output/broker .

FROM ubuntu:latest
WORKDIR /build
COPY --from=0 /output/broker /build/
RUN apt-get update \
    && apt-get install -y ca-certificates
ENTRYPOINT ["/build/broker"]

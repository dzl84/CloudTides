FROM golang:latest

WORKDIR /build

# ENV GOPROXY https://goproxy.cn
COPY . .
RUN go env -w GO111MODULE=on
# RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download
RUN cd lhj-test-broker/tides-broker/cmd && \
    GOARCH=amd64 GOOS=linux go build -o /output/broker .
#RUN go mod download
#RUN GOARCH=amd64 GOOS=linux go build -o /output/broker .
#FROM node:7-alpine
#RUN mkdir -p /src/app
#WORKDIR /build
#RUN npm install prometheus-query
#RUN npm install typescript@latest -g
FROM ubuntu:latest
WORKDIR /build
COPY --from=0 /output/broker /build/
# RUN sed -i 's/\/\/.*\/ubuntu/\/\/mirrors.aliyun.com\/ubuntu/g' /etc/apt/sources.list
RUN apt-get update \
    && apt-get install -y ca-certificates
RUN apt update
RUN apt -y install nodejs npm
RUN npm install prometheus-query
RUN npm install typescript@latest -g
#RUN apt-get -y install python-software-properties git build-essential
#RUN add-apt-repository -y ppa:chris-lea/node.js
#RUN apt-get -y install nodejs
#RUN npm install prometheus-query
ENTRYPOINT ["/build/broker"]

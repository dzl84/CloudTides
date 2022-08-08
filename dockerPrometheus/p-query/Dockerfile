
FROM golang:latest

WORKDIR /build

ENV GOPROXY https://goproxy.cn
COPY . .
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download
RUN cd ./ && \
    GOARCH=amd64 GOOS=linux go build -o /output/query_try .

#FROM ubuntu:latest
FROM node:7-alpine
#RUN mkdir -p /src/app
WORKDIR /build
RUN npm install prometheus-query
RUN npm install typescript@latest -g
#COPY . /src/app
#FROM alpine:latest
COPY --from=0 /output/query_try /build/
COPY instant_query.ts /build/
RUN chmod +x /build/instant_query.ts

#RUN apt update
#RUN apt -y install nodejs npm
#RUN npm install prometheus-query

ENTRYPOINT ["/build/query_try"]

# ---- Base Node ----
#FROM alpine:3.5 AS base
# install node
#RUN apk add --no-cache nodejs-current tini


#
# ---- Dependencies ----
#FROM base AS dependencies
# install node packages
#RUN npm set progress=false && npm config set depth 0
#RUN npm install prometheus-query



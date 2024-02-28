# syntax=docker/dockerfile:1

#BUILD GO BACKEND
FROM golang:1.20-alpine AS go-builder

ARG TARGETOS
ARG TARGETARCH

RUN apk add build-base
WORKDIR /usr/local/go/src/vessl
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY main.go ./
COPY ./handlers/ /usr/local/go/src/vessl/handlers
COPY ./modules/ /usr/local/go/src/vessl/modules
RUN CGO_ENABLED=1 GOOS=${TARGETOS} GOARCH=${TARGETARCH} GOFLAGS=-mod=mod go build -ldflags="-w -s" -o /Vessl

#BUILD WEBAPP
FROM node:latest as node-builder

WORKDIR /app
COPY ./webapp/package.json ./
COPY ./webapp/package-lock.json ./
RUN npm config set fetch-retry-mintimeout 20000
RUN npm config set fetch-retry-maxtimeout 120000
RUN npm install --force
COPY ./webapp .
RUN npm install -g @angular/cli
RUN ng build --output-path=/webapp/dist

#BUILD A SMALL FOOTPRINT IMAGE
FROM alpine:latest

COPY --from=go-builder /Vessl /Vessl
COPY --from=node-builder /webapp/dist/ /webapp/dist/
RUN mkdir -p /certs
RUN mkdir -p /database

EXPOSE 443

ENTRYPOINT [ "/Vessl" ]


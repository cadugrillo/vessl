# syntax=docker/dockerfile:1

#BUILD GO BACKEND
FROM golang:1.18-alpine AS go-builder

RUN apk add build-base
WORKDIR /usr/local/go/src/vessl
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY main.go ./
COPY ./handlers/ /usr/local/go/src/vessl/handlers
COPY ./modules/ /usr/local/go/src/vessl/modules
RUN mkdir -p /apps
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 GOFLAGS=-mod=mod go build -ldflags="-w -s" -o /Vessel

#BUILD WEBAPP
FROM node:latest as node-builder

WORKDIR /app
COPY ./webapp/package.json ./
COPY ./webapp/package-lock.json ./
RUN npm install --force
COPY ./webapp .
RUN npm install -g @angular/cli
RUN ng build --output-path=/webapp/dist

#BUILD A SMALL FOOTPRINT IMAGE
FROM nginx:alpine

COPY --from=go-builder /Vessel /Vessel
COPY --from=go-builder /apps /apps
COPY --from=node-builder /webapp/dist/ /webapp/dist/
COPY ./certs/ /certs
RUN mkdir -p /database

EXPOSE 443

CMD [ "/Vessel" ]


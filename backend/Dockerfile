FROM golang:alpine AS build-env
RUN apk --no-cache add build-base git bzr mercurial gcc
COPY . /src
WORKDIR /src
ENV GO111MODULE=on
RUN GOCGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o wirvsvirus

FROM alpine
WORKDIR /app
COPY --from=build-env /src/wirvsvirus /app/
EXPOSE 8080/tcp
ENTRYPOINT ./wirvsvirus --port=8080

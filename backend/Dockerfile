FROM golang:alpine AS build-env
RUN apk --no-cache add build-base git bzr mercurial gcc
ADD . /src
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w' -o wirvsvirus *.go

FROM alpine
WORKDIR /app
COPY --from=build-env /src/wirvsvirus /app/
EXPOSE 8080/tcp
ENTRYPOINT ./wirvsvirus

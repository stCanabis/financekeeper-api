FROM golang:latest AS build-env
ADD . /src
WORKDIR /src
RUN go get -d
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o backend

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/backend /app/
CMD ./backend
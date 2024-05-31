FROM golang:alpine as builder

WORKDIR /app
COPY . .

RUN CGO_ENBALED=0 GOOS=linux GOARCH=arm64 go build -o hellogo -ldflags="-w -s"

ENTRYPOINT [ "./hellogo" ]
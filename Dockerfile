FROM golang

ENV GO111MODULE=on

WORKDIR /app

COPY . .

RUN CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build

EXPOSE 3000

ENTRYPOINT ["/app/go-fiber-docker"]
FROM golang:1.17.7-alpine

RUN apk update && apk add git
RUN go get github.com/cosmtrek/air@v1.29.0
WORKDIR /app

# air -c <tomlファイル名>
CMD ["air", "-c", ".air.toml"]

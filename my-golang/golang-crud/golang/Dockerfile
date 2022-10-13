FROM golang:1.17.7-alpine

RUN apk update &&  apk add git
# RUN go get github.com/cosmtrek/air@v1.29.0
WORKDIR /go/src

# air -c [tomlファイル名] // 設定ファイルを指定してair実行(WORKDIRに.air.tomlを配置しておくこと)
# CMD ["air", "-c", ".air.toml"]
CMD ["go", "run", "main.go"]

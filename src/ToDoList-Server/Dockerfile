# ベースとなるDockerイメージの指定
FROM golang:1.17-alpine as server-build

# コンテナ内の作業ディレクトリを作成し、そこを指定する
WORKDIR /go/app/todo-crud
# ローカルの現在のディレクトリから、コンテナの作業ディレクトリにコピー
COPY . .
# alpineパッケージのアップデート
RUN apk upgrade --update && \
    apk --no-cache add git
RUN go get github.com/cosmtrek/air@v1.29.0

CMD ["air", "-c", ".air.toml"]

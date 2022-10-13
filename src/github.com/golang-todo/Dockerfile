FROM golang:1.16

WORKDIR /app

COPY go.mod .
COPY go.sum .

# go modからパッケージをダウンロード
RUN go mod download

# /app に全てのコードをコピー
COPY . .

# エントリポイント
CMD ["go", "run", "main.go"]

# DockerでGoを動かす
## 1. 要件
1. マルチステージングビルドを使用する
2. 動かすプログラムは「REST API」

## 2. マルチステージングビルド
主参考: [マルチステージビルドの利用](https://matsuand.github.io/docs.docker.jp.onthefly/develop/develop-images/multistage-build/)

#### 2-1. マルチステージングビルドとは
Multi-Staging Build(Multi-Stage Build)
直訳すると、**複数のステージを用いたビルド**

- Docker17.05以上で利用できる新機能
- 可読性・保守性を保ちながらDockerfileを最適化するのに苦労している人の役に立つ

#### 2-2. マルチステージングビルドの目的
- アプリケーションの実行に必要なもののみをビルドすること
そのために複数ステージを用いる

結果、イメージサイズが小さくなり、本番環境の運用のパフォーマンスが向上する

## 3. マルチステージングビルドが生まれた背景
- イメージをビルドする際に取り組む事と言えば、サイズを小さく保ちながらDockerイメージをビルドすること(命題的な意味で)

実際に行っていた事は、**開発用にはアプリケーションのビルドに必要な全てが含まれるDockerfileを使用し**、**プロダクト用にはアプリケーションおよび実行に必要なもののみが含まれるスリム化されたDockerfileを使用する**という事が行われ、**これが非常に一般的だった**。これがいわゆる`ビルダーパターン`(開発パターン)と呼ばれるもの。

しかしこの2つのDockerfilesを保守することは、理想的なやり方ではなかった。

#### 3-1. ビルダーパターン例
上述のビルダーパターンにこだわったやり方の例:

```docker: Dockerfile.build
# syntax=docker/dockerfile:1
FROM golang:1.16
WORKDIR /go/src/github.com/alexellis/href-counter/
COPY app.go ./
RUN go get -d -v golang.org/x/net/html \
  && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .
```
`RUN`コマンド行では、`&&`を使って2つの`RUN`コマンドを連結しています。
イメージ内に不要なレイヤーが生成されるのを防いでいるが、これでは間違いを起こしやすく、保守もやりづらくなる。
例えば: 別のコマンドを挿入した時に行の継続用の`\`を入れ忘れるなどの事態が容易に発生します。

```docker: Dockerfile
# syntax=docker/dockerfile:1
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY app ./
CMD ["./app"]
```

```docker: build.sh
#!/bin/sh
echo Building alexellis2/href-counter:build

docker build --build-arg https_proxy=$https_proxy --build-arg http_proxy=$http_proxy \
    -t alexellis2/href-counter:build . -f Dockerfile.build

docker container create --name extract alexellis2/href-counter:build
docker container cp extract:/go/src/github.com/alexellis/href-counter/app ./app
docker container rm -f extract

echo Building alexellis2/href-counter:latest

docker build --no-cache -t alexellis2/href-counter:latest .
rm ./app
```
`build.sh`を実行する時、最初のイメージをビルドし、成果物をコピーするためにコンテナを生成し、その後に2つ目のイメージをビルドする必要がある。
**2つのイメージは、それなりに容量をとるもので、ローカルディスク上に`app`の成果物も残ったままになってしまう。**

**こういった状況を非常にシンプルに出来る(解決できる)のが、マルチステージビルドです！**

## 4. マルチステージビルドを利用する


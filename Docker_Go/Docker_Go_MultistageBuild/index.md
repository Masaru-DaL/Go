# DockerでGoを動かす
## 1. 要件と道筋
1. マルチステージングビルドを使用する
2. 動かすプログラムは「REST API」

#### 1-1. 参考にするサイト
[Build your Go image](https://matsuand.github.io/docs.docker.jp.onthefly/language/golang/build-images/)を元に環境構築を行っていきます。

- 理由としては以下になります。
1. 公式ドキュメント
2. サンプルアプリケーションがRESTである要素
3. マルチステージへの導線

## 2. Dockerfile
**Dockerイメージを取得する命令を含んだテキストファイルのこと**。
Dockerに対して`docker build`コマンドを実行してイメージビルドを指示すると、Dockerは記述された命令を読み込んで実行し、最終的にDockerイメージを作り出す。

#### 2-1. Dockerfileをもう少し詳しく
- Dockerfileに対して用いられるデフォルトのファイル名は、`Dockerfile`という名前そのままです。
※ファイル拡張子はない(`.go`, `.py`などのこと)
この名前を用いておければ、`docker build`というコマンドを実行する際に、コマンドラインフラグ(オプションのような`-l`などのこと)を追加して指定する必要がない。

- プロジェクトによっては特定の目的のためにDockerfileに別名を与える場合がある。
慣例として、`Dockerfile.<something>`や`<something>.Dockerfile`とする。(somethingにName)このように名前を付けたDockerfileは、`docker build`コマンドの実行に際して`--file`オプション(短縮系`-f`)を用いて指定する。
**基本的にはデフォルト名の`Dockerfile`を用いる事が推奨されている。**

## 3. Dockerfileを書いてみる
実際に書きながら何を意味しているかを確認しながら進めていきます。

まず、プロジェクトのルートディレクトリに`Dockerfile`という名前のファイルを生成し、テキストエディタで開きます。

## 3-1. パーサーディレクティブ
1行目に書くのは`# syntax`パーサーディレクティブです。
以下のように書きます。
```docker: dockerfile
# syntax=docker/dockerfile:1
```
`docker/dockerfile:1`と記述することで、"**常に文法バージョン1の最新リリース版を指し示す**"ことになります。これが推奨されています。

主な特徴は2点
- **任意の記述**ではある。
- 記述する場合は第一に(Dockerfileの一番最初に)記述する事が必要である。

このパーサーディレクティブと呼ばれる記述が何を意味するかというと、Dockerfileの解析にあたってDockerビルダーがどの文法を採用するのかを指示する目的があります。

最初見た何だこれはと思って調べる時に[同じような疑問を持つ人の質問](https://qiita.com/yosaku_ibs/questions/050167327149837d6c5d#answer-884a349811b3ef1cce11)がありました。
Windowsに対応できるように指示したり、ヒアドキュメントというものを使えるようにするために記述したり出来るようです。

**補足**: Buildkit
[参考: Buildkit](https://genzouw.com/entry/2021/07/17/100615/2724/)
Buildkitは、**ビルド処理の前に文法に更新がないかを自動的にチェックし、最新バージョンが用いられていることを確認します**。
最新のDocker環境を利用している限り既に有効になっているようですので、あまり気にしなくても良いですが、パーサーディレクティブと関係はあります。
最新ではないDocker環境の場合、Buildkitが有効になっていない場合があります。その場合はパーサーディレクティブを記述することでビルド前にパーサーをアップグレードするようになります。

#### 3-2. ベースとなるイメージの指定
次に、アプリケーションに使用する基本イメージをDockerに伝える行を追加する必要があります。
Dockerイメージは他のイメージから継承ができます。(既存のイメージを使用できるということ。)
これはオブジェクト指向プログラミングのクラス継承と同じように考える事ができます。

以下は、公式のGoイメージを使用しています。
```docker: dockerfile
FROM golang:1.16-alpine
```
このように`FROM ~`と記述します。
`FROM`はベースとなるイメージを指定し、Dockerfileの先頭(#syntaxの後)に必ず必要です。

この最初の`FROM`の後に続くコマンドはすべて、ここで指定した「ベースイメージ」の上に構築されます。

#### 3-3. ディレクトリの作成
FROM以降に書くコマンドを簡単に実行するために、構築中のイメージ内にディレクトリを作成する。
`WORKDIR </ディレクトリ名>`と記述します。
ここでは`app`と名付けます。

```docker: dockerfile
WORKDIR /app
```
このようにディレクトリを作成すると、このディレクトリを基点としてコマンドを記述する事ができます。この場合、作成したディレクトリに基づく相対パスが使用できます。

#### 3-4. ファイルをイメージにコピー
通常、Goで記述されたプロジェクトをダウンロードして最初に行うことはコンパイルに必要なモジュールをインストールすることです。
しかし、イメージ内でそれを実行するには先にファイルをイメージにコピーする必要があります。

コピーするのは`go.mod`, `go.sum`の2つです。
詳細は[こちら](https://blog.framinal.life/entry/2021/04/11/013819#gomod%E3%81%A8%E3%81%AF)

```docker: dockerfile
COPY go.mod ./
COPY go.sum ./
```
`./`は`WORKDIR`で作成したディレクトリから見た相対パスです。つまり`/app`ディレクトリにコピーすることを指しています。

#### 3-5. ビルド時に実行するコマンドの指定
次に記述するのは`RUN`コマンドです。
`RUN`は**ビルド時に実行するコマンドを指定します**。
- ビルド(build)
> ソースコード上に問題がないかどうかを解析を行った上で、問題がなければオブジェクトコードに変換し、複数のオブジェクトファイルを1つにまとめて実行可能なファイルを作成する作業を指します。
簡単に言うと、プログラムの元ネタから実際のプログラムを作る作業工程のことです。

今回ビルドしたイメージ(Go 公式)にはモジュールファイルがあるので、それを実行するように指定します。

ローカルで実行した場合と全く同じように機能しますが、今回のコマンドは、Goモジュールがイメージ内のディレクトリにインストールされることを意味します。

```docker: dockerfile
RUN go mod download
```

#### 3-6. ソースコードをイメージにコピーする
```docker: dockerfile
COPY *.go ./
```
`COPY *.go`が意味するのは、"ワイルドカード"を用いて、ホスト上の現在のディレクトリ(Dockerfileがあるディレクトリ)にある拡張子が`.go`である全てのファイルを、イメージ内のカレントディレクトリにコピーすることを意味しています。

#### 3-7. アプリケーションのコンパイル
次に、`RUN`コマンドを用いて、アプリケーションのコンパイルを行います。

```docker: dockerfile
RUN go build -o /docker-gs-ping
```
おさらいですが、`RUN`はビルド後に実行するコマンドを指定することを意味します。
つまり、ビルド後に`go build -o /docker-gs-ping`を実行するということです。

`go build -o /docker-gs-ping`の`-o`の意味があやしいですが、[go documentation](https://pkg.go.dev/cmd/go#hdr-Compile_and_run_Go_program)の"go test"の項目に`-o`について書いてあったので引用すると、
> `-o <file>`
> バイナリを指定されたファイルにコンパイルします。
とあります。
つまり、コンパイルした後のファイル名を"docker-gs-ping"として配置するという意味と捉えられます。

**コンパイルされた"docker-gs-ping"というファイルはDockerが実行できる状態になった**ということが重要です。

配置場所は構築中のイメージのファイルシステムのルートです。
ルートディレクトリ(配置場所)に特別な意味はないが、ルートに配置することで読みやすさ、ファイルパスが短くなるため、便利です。

#### 3-8. ポート番号
公式の方には書いてない(意図不明)ですが、完成コードには`EXPOSE 8080`が記述されています。

`EXPOSE`は指定したポート番号をコンテナが公開することをDockerに伝えるという意味があります。

#### 3-9. docker run時に実行するコマンド
最後にコンパイルした"docker-gs-ping"ファイルをコンテナを起動する時に実行するコマンドとしてDockerに指示を出す文を記述します。

```docker: dockerfile
CMD [ "/docker-gs-ping" ]
```

#### 3-10. 完成したDockerfile
完成したDockerfileが以下です。
```docker: dockerfile
# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /docker-gs-ping

EXPOSE 8080

CMD [ "/docker-gs-ping" ]
```

#### 3-11. Comments
Dockerfileには`#`を使用してコメントを書く事ができます。

必ず行頭に`#`を付けて記述します。
コメントはDockerfileを文書化するために便宜的に存在します。

※syntaxディレクティブが存在する場合はこのディレクティブの後に書きましょう。syntaxディレクティブは全てにおいて最優先されます。

- コメント例
```docker: dockerfile(comments)
# syntax=docker/dockerfile:1

# Alpine is chosen for its small footprint
# compared to Ubuntu
FROM golang:1.16-alpine

WORKDIR /app

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# ... the rest of the Dockerfile is ...
# ...   omitted from this example   ...
```

#### 3-12. イメージのビルド
公式の冒頭でサンプルアプリケーションのクローンが促されています。
今回作成したDockerfileも[サンプルアプリケーション](https://github.com/olliefr/docker-gs-ping)にあります。

※自分は理解していなかったのですが、
```docker: dockerfile
COPY go.mod ./
COPY go.sum ./
```
この記述はプロジェクトディレクトリに`go.mod`ファイル、`go.sum`ファイルがあり、それをコピーするという意味です。

なので、公式のDockerfileはサンプルアプリケーションをクローンしてある前提でした。

```code:
$ git clone https://github.com/olliefr/docker-gs-ping
```

- クローンしたファイル内(ディレクトリ内のルート)で以下のコマンドを実行
```code:
$ docker build --tag docker-gs-ping .
```
エラーがなければ`FINISHED`と出ます。

`--tag`は、ビルドしたイメージにラベルを付け、読みやすく認識しやすい文字列値で表示できます。
もし`--tag`を付けない場合は、デフォルト値として`latest`が使用されます。

無事にイメージをビルドできたら、`$ docker image ls`と打ち、作成したイメージを見てみましょう。
REPOSITORY名が`docker-gs-ping`という名前で作成できているのが確認できます。

## 4. マルチステージビルド
[サンプルアプリケーション](https://github.com/olliefr/docker-gs-ping)には、`Dockerfile.multistage`という名前のファイルがあります。
これがマルチステージビルドされた`Dockerfile`です。

```docker: dockerfile(multistage)
# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /docker-gs-ping

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /docker-gs-ping /docker-gs-ping

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/docker-gs-ping"]
```

#### 4-1. マルチステージビルドでイメージのビルド
最初にイメージのビルドを行った階層で行います。
```code:
$ docker build -t docker-gs-ping:multistage -f Dockerfile.multistage .
```
TAGの名前に意味はなく、比較の為に`multistage`と付けています。
`-f`でビルドに用いるファイルを指定します。(Dockerfile.multistageというファイル名のためです。)

- `docker image ls`
![](2022-09-07-17-18-05.png)

注目すべきはSIZEです。
**SIZEが桁違いに違います**。

#### 4-2. Dockerfile(マルチステージ)
マルチステージを行うDockerfileの記述を確認します。

1. `FROM golang:1.16-buster AS build`
`golang:1.16-buster`は、Golangの最新devianパッケージの構成という意味のようです。
どういったバージョンを使うか、を指定するとだけ覚えておけばOKです。
`AS build`で別名を付けてます。ここも重要なポイントです。
この別名は後で効いてきます。

2. `FROM gcr.io/distroless/base-debian10`
`distroless`は、Googleが提供している必要最小限の依存のみが含まれるコンテナイメージのことです。
他詳細 -> [コンテナイメージ使うならdistrolessもいいよねという話](https://zenn.dev/yoshii0110/articles/21ddb58c6f6bfa)

3. `COPY --from=build /docker-gs-ping /docker-gs-ping`
注目すべきは`--from=build`です。ここの`build`という名前は1の別名を指しています。
つまり、1つめのビルドステージのイメージを参照し、実行に必要な`docker-gs-ping`のバイナリだけをピンポイントでコピーしているということです。

4. `ENTRYPOINT ["/docker-gs-ping"]`
`ENTRYPOINT`は、`docker run`時に実行するコマンドを指定します。
> CMD と似てますが、「--entrypoint オプション > ENTRYPOINT > run引数 > CMD」の優先度があります。

#### 4-2. Dockerfileの書き方の違い
- マルチステージビルドの良い所
1. `FROM`を複数回書ける
> `FROM`を2回書く事自体は既にできていたようです。
> 中間イメージを`AS`で名前を付け、それを直接参照できることが新しくできるようになったことです。

以前は複数のDockerfileを組み合わせたりしていた事を1つのDockerfileだけで済むようになった。
`AS`で中間イメージを作り、それを用いて行う。

2. 打ち間違い、可読性の向上
以前は無理に`&&`や`\`を使ってコマンドを繋げていたことをしなくて良くなる。
Dockerfile自体も見やすくなりますし、打ち間違いも減ります。

3. SIZEが大幅に軽量になる
最初にSIZEを比べましたが明らかにマルチステージビルドの方が軽量です。

おおまかにはこの辺りが利点として挙げられます。

## 5. イメージの実行
イメージのビルドがゴールではありません。
次に、ビルドしたイメージを実行します。
コンテナ内でイメージを実行するには`docker run`コマンドを使います。
このコマンドは、1つの"**イメージ名**"を引数に必要とします。

```shell:
$ docker run docker-gs-ping
```

```shell:
   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.2.2
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:8080
```

これが見えたら成功です。
打ち込んだ後、入力状態には戻りません。戻る為にはコンテナを停止させる必要があります。

```code:
$ docker run -d docker-gs-ping
```

`⇨ http server started on [::]:8080`
となっているので`curl`を用いてポート8080にデータを送ってみます。

```shell:
$ curl http://localhost:8080/

# 実行結果
# curl: (7) Failed to connect to localhost port 8080 after 8 ms: Connection refused
```
公式によると、この出力は期待通りのようです。
> **コンテナはネットワークも含めた隔離された環境内で実行されているから**
ということです。

今度は同じ`docker run`をポート8080を公開した上で再起動します。
一度コンテナを停止します。

## 5-1. 起動: フォアグラウンドとバックグラウンドの違い
```shell:
# フォアグラウンドでの起動
# 起動させたターミナルでは起動状態でコマンドを受け付けない
$ docker run docker-gs-ping

# バックグラウンドでの起動
# 起動させたターミナルは起動後も使用できる
$ docker run -d docker-gs-ping
```
この違いは何でしょうか？
自分は便利だからバックグラウンドでいいやーとか思ってたのですが、両者の違いやメリットが存在しました。

1. フォアグラウンドではレスポンスが見れる
`curl`など、通信を行った際にレスポンスコード(200, 300など)が見れます。
開発環境ではこれらを確認しながら開発を進めるのが便利のようです。
起動させた後は別のターミナルを開いて、そちらでターミナルコマンドを実行します。

2. バッググラウンドの用途
どちらかというと開発後の用途になります。
もうサーバーとの状態をフォアグラウンドで確認しなくても良いので、実際のプロダクトはバックグラウンドで起動させます。

#### 5-2. ポートを公開し、通信をやり取りする
```shell:
$ docker run -d -p 8080:8080 docker-gs-ping
# 起動したコンテナIDが表示されます。
# docker ps または、DockerDesktopで起動を確認してみてください。

$ curl http://localhost:8080/

# Hello, Docker! <3 と返ってきます。
```
これは[サンプルアプリケーション](https://github.com/olliefr/docker-gs-ping)の`main.go`の以下が返ってきていました。

```go: main.go
e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello, Docker! <3")
	})
```

#### 5-3. コンテナに名前を付ける
`docker run`後に作成されたコンテナ名(NAMES)はランダムで付けられます。
この名前を用いて再起動や削除を行います。
名前は`docker ps`で確認出来ますが、もし面倒だったり任意の名前を付けたい場合は`--name`オプションを用います。
```shell:
 docker run -d -p 8080:8080 --name rest-server docker-gs-ping
```
このように指定すると、NAMESが"rest-server"という名前でコンテナが作成されます。

**※この章のコンテナの一覧表示、削除、起動、などは割愛します。**

#### 5-4. ふとした疑問
ここで立ち上げたコンテナってマルチステージビルドのイメージなの？という疑問が沸きました。

一度コンテナを全部削除して、イメージもマルチステージを残して削除します。
その後、`docker run`を実行します。

```shell:
$ docker run -d -p 8080:8080 docker-gs-ping

# 以下にエラー文が出力されます。
# Unable to find image 'docker-gs-ping:latest' locally
# docker: Error response from daemon: pull access denied for docker-gs-ping, repository does not exist or may require 'docker login': denied: requested access to the resource is denied.
# See 'docker run --help'.
```

その後付けたタグを指定して実行すると成功します。
```shell:
$ docker run -d -p 8080:8080 docker-gs-ping:multistage

# 実行結果
# コンテナIDの出力
```
マルチステージのタグを付けてイメージを作成していたので、デフォルトの`docker run`は`latest`を参照するのかなという推察です。
タグ名を付けてイメージを作成した場合は明示的に指定する必要があると思われます。

## 6. データベースエンジン
公式の次の手順は、以下です。
1. データベースエンジンを実行し、これをさプルアプリケーションに接続。
2. `Docker Compose`を使用して複数コンテナの管理

#### 6-1. 使用するデータベースエンジン
[CockroachDB](https://www.cockroachlabs.com/product/)と呼ばれ、最新のクラウドネイティブの分散型SQLデータベースです。
CockroachDBのDockerイメージを使用します。

#### 6-2. ストレージ
データベースの重要な点は、**データの永続的な保存を行うこと**と表記してます。

この言い回しは、コンテナのサイクルと関係があります。
コンテナ内で発生したデータは同じコンテナ内のどこかに書き出されるが、コンテナを破棄すると消えてしまいます。
コンテナは生成->削除がある意味1セットな考え方、手軽さがありますので、コンテナにおいてデータを永続的に保存したい場合は"ボリューム"というメカニズムを利用します。

ボリュームの作成には次を実行します。
```shell:
$ docker volume create roach

# 実行結果
# roach
```

ボリュームのリストの表示
```shell:
$ docker volume list

# 実行結果
# DRIVER    VOLUME NAME
# local     roach
```

#### 6-3. ネットワークの構築
サンプルアプリケーションとデータベースエンジンは、ネットワークを介して相互に通信を行います。
さまざまな種類のネットワーク構成が可能で、ユーザー定義ブリッジネットワークと呼ばれるものを使用します。

```shell:
# -dはネットワークを管理するドライバーを指定するオプション
$ docker network create -d bridge mynet

# 実行結果
# NETWORK ID
```
- `docker network create`でブリッジネットワークが作成されます。
- ブリッジネットワークは仮想ブリッジを使用する。
(ブリッジはOSI参照モデルのデータリンク層における通信を制御する)
- ユーザーが作成したブリッジネットワークをユーザー定義ブリッジネットワークと呼ぶ。

ネットワークを一覧表示して確認する
```shell:
$ docker network list

# 実行結果
# NETWORK ID     NAME      DRIVER    SCOPE
# 96bd8ddeb5bb   bridge    bridge    local
# 620d216e0654   host      host      local
# c49f4a66c445   mynet     bridge    local
# e1be2f472332   none      null      local
```
mynet以外に3つありますが、これはDocker自体によって作成されている。
詳細: [ネットワークの概要](https://matsuand.github.io/docs.docker.jp.onthefly/network/)に、今回作成された<NAME>で見ると何であるか確認出来ます。

#### 6-4. 適切な名前付け
- コンピュータサイエンスで難しいと言われている事が2つ
1. キャッシュの無効化と名前付け
2. [Off-by-one-Error](https://ja.wikipedia.org/wiki/Off-by-one%E3%82%A8%E3%83%A9%E3%83%BC)

ネットワークおよび管理ボリュームの名前は、意図した目的を示す名前を付ける事が推奨されている。

#### 6-5. データベースエンジンの起動
ここまでの一通りの作業が終わると、CockroachDBをコンテナで実行し、先ほど作成したボリュームとネットワークに接続できるようになりました。

以下のコマンドを実行すると、DockerがDocker Hubからイメージを取得してローカルで実行してくれる。
```shell:
$ docker run -d \
  --platform linux/x86_64 \
  --name roach \
  --hostname db \
  --network mynet \
  -p 26257:26257 \
  -p 8080:8080 \
  -v roach:/cockroach/cockroach-data \
  cockroachdb/cockroach:latest-v20.1 start-single-node \
  --insecure

# ... output omitted ...
```
- M1Macの人は
公式通りにコマンドを打つとエラーになります。
原因はM1Macによるものです。
明示的にplarformを指定する必要があります。
`--platform linux/x86_64 \`の行です。
※Dockerイメージがplatform(今回の場合M1のarm64)に対応していない場合に起こります。

#### 6-6. データベースエンジンの設定
アプリケーションでの使用を開始する前に行わなければいけない設定が幾つかあります。
コンテナに入り、SQLコマンドを用いて行います。これはCockroachDBの組み込みSQLシェルの機能です。

```shell:
$ docker exec -it roach ./cockroach sql --insecure
# root@~ のようにSQLコマンドを受け付ける状態になる
```

1. 空のデータベースの作成
```sql:
CREATE DATABASE mydb;
```

2. データベースエンジンに新しいユーザーアカウントを登録
```sql:
CREATE USER totoro;
```
"totoro"は任意です。

3. 新しいユーザーにデータベースへのアクセス権の付与
```sql:
GRANT ALL ON DATABASE mydb TO totoro;
```

4. `quit`と入力し、シェルの終了

#### 6-7. この先動かすサンプルアプリケーション
ここから使用するサンプルアプリケーションは、これまでに使用してきた"docker-gs-ping"を拡張したものになります。
- 拡張するには
1. ローカルにコピーしたものを更新する
2. [拡張済みのもの](https://github.com/olliefr/docker-gs-ping-roach)をクローンして使用する
公式では2を推奨していますので、倣ってクローンします。
(docker-gs-pingとは違うディレクトリが良いでしょう。)
```shell:
$ git clone https://github.com/olliefr/docker-gs-ping-roach.git
# ... output omitted ...
```
"docker-gs-ping-roach"というディレクトリがクローンされました。

- 拡張後の変更点
`main.go` ->
データベースの初期化コードと新しいビジネス要件を実装するコードの追加

拡張後のDockerfileを見ると、マルチステージビルドに対応した記述がされていますね。
`FROM gcr.io/distroless/base-debian10`
ここを調べると、ポイントは"distoless"のようです。

ベースとしてディストリビューションにはカーネルを除く基本的な設定ファイルやパッケージが一通り含まれているので、こうした**不要なファイルを削除し、アプリケーションの実行に必要な最小限のファイルのみを含んだコンテナイメージ**をビルドすることを意味しています。

#### 6-8. アプリケーションのビルド
- アプリケーションのビルド
```shell:
$ docker build --tag docker-gs-ping-roach .
```
#### 6-9. アプリケーションの実行
まず、アプリケーションがデータベースへのアクセス方法が認識できるように、いくつかの環境変数を設定する必要があるようです。
`docker run`コマンドを用いて行います。

```shell:
$ docker run -it --rm -d \
  --network mynet \
  --name rest-server \
  -p 80:8080 \
  -e PGUSER=totoro \
  -e PGPASSWORD=myfriend \
  -e PGHOST=db \
  -e PGPORT=26257 \
  -e PGDATABASE=mydb \
  docker-gs-ping-roach
```

以下のコマンドが実行できればOKです。
```shell:
$ curl localhost

# 実行結果
# Hello, Docker! (0)

# または、
$ curl http://localhost/

# 実行結果
# Hello, Docker! (0)
```
DockerDesktopなどで確認すると分かりますが、`-p 80:8080 \`でホストポートを80にしています。
ポート80は**WebサーバがHTTPでWebブラウザなどと通信するために**用いられています。

- 出力された`(0)`はメッセージの合計数です。
アプリケーションにはまだ何も投稿していないので問題ありません。

- `-e PGHOST=db \`
データベースコンテナを起動するときに付けた`--name db`です。

- `-e PGPASSWORD=myfriend \`
サンプルアプリケーションを混乱させないために何か設定しているだけ。ここで付けているパスワードに深い意味はないようです。

- `--name rest-server \`
ここで付けた"rest-server"という名前はコンテナのライフサイクルを管理(起動、削除など)するのに役立ちます。

#### 6-10. アプリケーションのテスト
[curl man page](http://www.mit.edu/afs.new/sipb/user/ssen/src/curl-7.11.1/docs/curl.html)

- メッセージを投稿してみる
```shell:
curl --request POST \
  --url http://localhost/send \
  --header 'content-type: application/json' \
  --data '{"value": "Hello, Docker!"}'
```
`--data`でHTMLフォームでデータを送信したかのように出来ます。
何をやってるかというと、**データを指定urlにJSON形式でPOSTしています**。

`{"value":"Hello, Docker!"}`が出力されます。
これはメッセージがデータベースに保存されたことを意味しています。

- 別のメッセージを投稿してみる
```shell:
$ curl --request POST \
  --url http://localhost/send \
  --header 'content-type: application/json' \
  --data '{"value": "Hello, Oliver!"}'

# 実行結果
# {"value":"Hello, Oliver!"}
```

- メッセージカウンターの確認
```shell:
$ curl localhost

# 実行結果
# Hello, Docker! (2)
```
メッセージを投稿した数だけカウンターが上がっています -> `(2)`

- データベースに保存されているか確認
これは**ボリュームによるデータの保存の永続化**の確認です。
1. コンテナの停止
立ち上がっているコンテナを2つ停止します。
`docker ps`でも確認出来ます。出てきた`NAMES`を指定します。
```shell:
$ docker container stop rest-server roach

# 実行結果
# rest-server
# roach
```

2. コンテナの削除
```shell:
$ docker container rm rest-server roach

# 実行結果
# rest-server
# roach
```

3. 削除されているかの確認
```shell:
$ docker container list --all

# 実行結果
# コンテナが出てこなければOK
# CONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES
```

4. データベースのビルド
```shell:
docker run -d \
  --platform linux/x86_64 \
  --name roach \
  --hostname db \
  --network mynet \
  -p 26257:26257 \
  -p 8080:8080 \
  -v roach:/cockroach/cockroach-data \
  cockroachdb/cockroach:latest-v20.1 start-single-node \
  --insecure
```
※M1Macの方はplatformの指定を忘れずに。

5. アプリケーションのビルド
```shell:
docker run -it --rm -d \
  --network mynet \
  --name rest-server \
  -p 80:8080 \
  -e PGUSER=totoro \
  -e PGPASSWORD=myfriend \
  -e PGHOST=db \
  -e PGPORT=26257 \
  -e PGDATABASE=mydb \
  docker-gs-ping-roach
```

6. クエリの実行
```shell:
curl localhost

# 実行結果
# Hello, Docker! (2)
```
データベースにきちんと保存されている結果が返ってきます。
これは、CockroachDBの管理ボリュームを再利用しているためです。

#### 6-11. 一度全てを停止、削除する
説明していませんでしたが、データベースエンジンをビルドした際のコマンドに`--insecure`というコマンドがありました。
これは"安全ではないモード(状態)"で実行するという意味です。
本番環境ではこのモードで実行してはいけません。
なので、一度全てのコンテナを停止、削除します。

```shell:
# 1. 起動しているコンテナの確認
$ docker container list

# 2. コンテナの停止
$ docker container stop <コンテナ名>

# 3. コンテナの削除
$ docker container rm <コンテナ名>
```

## 7. Docker Compose
ここまで長い`docker`コマンド(引数の長いリスト)を実行してきました。これは打ち間違いも多く、非常に労力がかかります。

これを回避する方法として"Docker Compose"を利用する方法があります。
**1つのDocker Composeファイル(docker-compose.yml)を用いて、"docker-gs-ping-roachアプリケーション"と"CockroachDBデータベース・エンジン"を起動させる事ができます。**

#### 7-1. Docker Composeの構成
アプリケーションのディレクトリに、`docker-compose.yml`という名前のファイルを作り以下のように記述していきます。

```yml:
version: '3.8'

services:
  docker-gs-ping-roach:
    depends_on:
      - roach
    build:
      context: .
    container_name: rest-server
    hostname: rest-server
    networks:
      - mynet
    ports:
      - 80:8080
    environment:
      - PGUSER=${PGUSER:-totoro}
      - PGPASSWORD=${PGPASSWORD:?database password not set}
      - PGHOST=${PGHOST:-db}
      - PGPORT=${PGPORT:-26257}
      - PGDATABASE=${PGDATABASE:-mydb}
    deploy:
      restart_policy:
        condition: on-failure
  roach:
    image: cockroachdb/cockroach:latest-v20.1
    container_name: roach
    hostname: db
    networks:
      - mynet
    ports:
      - 26257:26257
      - 8080:8080
    volumes:
      - roach:/cockroach/cockroach-data
    command: start-single-node --insecure

volumes:
  roach:

networks:
  mynet:
    driver: bridge
```
※`docker-compose.yml`におけるインデントは非常に大事です。これがずれているだけ思ったように動かない場合があります。

このDocker Composeの設定は、`docker run`コマンドに渡すパラメータを全て渡す必要がありません。**超便利です**。

#### 7-2. .envファイル
DockerComposeは、`.env`ファイルがあればそこから自動的に環境変数を読み取ります。
今回のComposeファイルでは`PGPASSWORD`を設定する必要があるため、`.env`ファイルに以下の内容を追加します。

```yml:
# 設定している値はこの値でなくても構いません。
# エラーが発生しないように何らかの値を設定します。
PGPASSWORD=whatever
```

- `.env`ファイルの扱い
見て分かる通り、`.env`ファイルに書く内容はパスワードなど、**他者に知られてはいけない内容を記述します**。
`git`などを用いてパプリックな場所に保管する場合は`.gitignore`にファイルを記載して、セキュアな状態にする必要があります。

#### 7-3. Composeファイル
"docker-compose.yml"というファイル名は、`-f`フラグを指定しない場合に`docker-compose`コマンドで認識されるデフォルトのファイル名です。
これは複数のDockerComposeファイルを持つ事ができることを意味しています。

#### 7-4. DockerComposeの変数置換
- 前提知識
  - 環境変数とは何か？
環境変数とは、開発・テスト・本番などの**環境ごとに変化する値を入れる変数のこと**です。
値を直接的な数値ではなく、変数にすることで**環境ごとに書き換える事なく運用することが出来るのが利用する利点**です。

DockerComposeの非常に優れた機能の1つが**変数置換**です。
`<変数>=${<変数に入れる値>}`のように`docker-compose.yml`に記述します。
"7-1"で作成した`docker-compose.yml`の内容を例にとります。

- `PGUSER=${PGUSER:-totoro}`
環境変数`PGUSER`は、DockerComposeが実行されているホストマシンと同じ値に設定されること意味します。
ホストマシンにこの名前の環境変数がない場合、コンテナ内の変数はデフォルトの`totoro`になります。

- `PGPASSWORD=${PGPASSWORD:?database password not set}`
環境変数`PGPASSWORD`がホスト上に設定されていない場合、DockerComposeがエラーを表示することを意味します。
パスワードは値を設定するのではなく、`.env`から参照することを設定するので、これで問題ありません。
**きちんと`.env`を`.gitignore`に記載してシークレットしましょう。**

#### 7-5. DockerComposeの構成の検証
```shell:
# 以下のコマンドで検証ができる。
$ docker-compose config
```
`.env`の作成を忘れずに行いましょう。
これがないと"PGPASSWORDが設定されていませんよ"というエラーが出力されます。

#### 7-6. DockerComposeを使用してアプリケーションをビルドして実行する
アプリケーションを起動し、正しく動作するかを確認します。
```shell:
$ docker-compose up --build
```
`--build`は、Dockerがイメージをコンパイルして起動するように指定するフラグです。
`--build`を指定した場合、ソースコードが更新された場合にリビルドが発生します。自分のソースコードを編集し、`docker-compose up`を実行する際に`--build`フラグを使い忘れるというのは非常によくある落とし穴、と記述されています。(自分も使い忘れていました。)

DockerComposeによってセットアップが実行され、"プロジェクト名"が割り当てられた為、CockroachDBインスタンス用の新しいボリュームを取得しました。(新しく"docker-gs-ping-roach_roach"というボリュームがあります。)
出力されている内容は、**データベースにこの新しいボリュームが存在しないためにアプリケーションがデータベースへの接続に失敗していることを意味するエラー**のようです。
"docker-compose.yml"に`restart_policy:`を使用してデプロイ設定しているため、失敗したコンテナは20秒ごとに再起動しています。

**これを解決するには、データベースエンジンにログインしてユーザーを作成する必要があります。**
再起動を続けている状態のまま続けます。
また、別のターミナルを立ち上げて以降のコマンドを実行します。

- "6-6"で行ったことと同じ事をします。
1. コンテナに入る
2. データベースの作成 -> `mydb`
3. ユーザーの作成 -> `totoro`
4. 権限の付与
一連の作業を行うと、コンテナが自動的に再起動されます。
そうすると、コンテナの失敗と再起動が停止し、以下の表示が出力されます。

```shell:
rest-server  |
rest-server  |    ____    __
rest-server  |   / __/___/ /  ___
rest-server  |  / _// __/ _ \/ _ \
rest-server  | /___/\__/_//_/\___/ v4.3.0
rest-server  | High performance, minimalist Go web framework
rest-server  | https://echo.labstack.com
rest-server  | ____________________________________O/_______
rest-server  |                                     O\
rest-server  | ⇨ http server started on [::]:8080
```

#### 7-6. アプリケーションのテスト
フォアグラウンドで実行が続いてる場合は別のターミナルで実行します。
```shell:
$ curl http://localhost

# 実行結果
# Hello, Docker! (0)
```

ここまででだいぶ長くなってしまったので、続きはまた別の記事にしたいと思います。

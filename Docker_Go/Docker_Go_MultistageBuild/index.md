# DockerでGoを動かす
## 1. Step.1 要件と道筋
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


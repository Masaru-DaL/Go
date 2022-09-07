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
```docker: #syntax
# syntax=docker/dockerfile:1
```
主な特徴は2点
- **任意の記述**ではある。
- 記述する場合は第一に(Dockerfileの一番最初に)記述する事が必要である。

このパーサーディレクティブと呼ばれる記述が何を意味するかというと、Dockerfileの解析にあたってDockerビルダーがどの文法を採用するのかを指示する目的があります。


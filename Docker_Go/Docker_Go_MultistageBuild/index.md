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

- 質問内容
`docker run`実行時に明示的にplatformを指定する必要があるのはDBだけなのかどうか。
主にDockerでMySQLの環境構築の時にplatformを指定するような記事が多かったためです。

- データベースエンジンの起動
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
`--platform linux/x86_64 \`の記述を追加すると上手くいった。


- アプリケーションの実行
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
`--platform linux/x86_64 \`がなくても上手くいった。

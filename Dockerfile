# ベースイメージを指定
FROM golang:latest

#RUN apk update && apk add --no-cache git

# ワーキングディレクトリを指定
WORKDIR /go/src/app

# ホストのファイルをコンテナのワーキングディレクトリにコピー
COPY go.mod go.sum ./

# ライブラリのダウンロード
RUN go mod download

COPY . .

#CMD ["go", "run", "main.go"]

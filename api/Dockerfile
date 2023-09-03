FROM golang:1.19.9-buster as local

RUN apt-get update && \
    apt-get install -y --no-install-recommends ca-certificates git && \
    rm -rf /var/cache/apt/*

# バイナリファイルにビルド
# RUN GOOS=linux GOARCH=amd64 go build -mod=readonly -v -o server

RUN go install github.com/cosmtrek/air@latest

WORKDIR /api/cmd

# バイナリファイルを実行
CMD ["air"]
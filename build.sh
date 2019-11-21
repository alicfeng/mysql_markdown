#!/usr/bin/env bash
go build -o release/mysql_markdown_mac mysql_markdown.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o release/mysql_markdown_unix mysql_markdown.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o release/mysql_markdown_win mysql_markdown.go
chmod +x release/mysql_markdown_*
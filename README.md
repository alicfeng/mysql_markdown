## mysql_markdown
It can generate markdown structure documents of MySQL succinctly~

#### install
###### unix
```shell
curl -o /usr/local/bin/mysql_markdown -sSL https://raw.githubusercontent.com/alicfeng/mysql_markdown/master/release/mysql_markdown_unix
chmod +x /usr/local/bin/mysql_markdown
```
###### mac
```shell
curl -o /usr/local/bin/mysql_markdown -sSL https://raw.githubusercontent.com/alicfeng/mysql_markdown/master/release/mysql_markdown_mac
chmod +x /usr/local/bin/mysql_markdown
```
###### other
```shell
git clone https://github.com/alicfeng/mysql_markdown.git
cd mysql_markdown
go get "github.com/go-sql-driver/mysql"
go build -o /usr/local/bin/mysql_markdown mysql_markdown.go
chmod +x /usr/local/bin/mysql_markdown
```

#### usage
```shell
mysql_markdown -h
Usage: mysql_markdown [options...]
--help  This help text
-h      host.     default 127.0.0.1
-u      username. default root
-p      password. default root
-d      database. default mysql
-P      port.     default 3306
-c      charset.  default utf8
-o      output.   default current location
```

#### simple
```shell
➜ mysql_markdown -p samego -d samego
mysql connected ...
1/8 the demo table is making ...
2/8 the failed_jobs table is making ...
3/8 the migrations table is making ...
4/8 the password_resets table is making ...
5/8 the roles table is making ...
6/8 the user table is making ...
7/8 the userinfo table is making ...
8/8 the users table is making ...
mysql_markdown finished ...
```

#### md2anyDoc
md转其它类型的文档推荐使用 `typora` 工具 它支持如下转换格式
- md2pdf
- md2html
- md2html(without styles)
- md2word
- md2rtf
- md2openOffice
- md2Epub
- md2latex
- md2MediaWiki
- md2reStructureText
- md2textile
- md2OPML
- md2png

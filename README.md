<h1 align="center">
  <a href="https://github.com/alicfeng/mysql_markdown">
    mysql_markdown
  </a>
</h1>
<p align="center">
  🛠 It can generate markdown structure documents of MySQL succinctly~
</p>
<p align="center">
  <a href="https://github.com/alicfeng/mysql_markdown/releases">
    <img src="https://img.shields.io/github/release/alicfeng/mysql_markdown" alt="release version">
  </a>
  <a href="https://github.com/alicfeng/mysql_markdown">
    <img src="https://travis-ci.org/alicfeng/mysql_markdown.svg?branch=master" alt="build status">
  </a>
  <a href="https://lbesson.mit-license.org">
    <img src="https://img.shields.io/badge/License-MIT-blue.svg" alt="open source">
  </a>
</p>


## 前言背景

在实施软件工程的时候，当要将某一版本归档时，需要汇总的文档要求还是比较高的、各类文档齐全，包括项目架构、项目安装、接口等文档，而数据库表结构说明文档亦属于其一。记得很早之前想找一个可以导出 `MySQL` 数据库表结构说明文档的工具，生态上的这种工具是有的、只不过并没有一个使我比较满足的。当然、看个人所需，我需要一个可以导出 `markdown` 文档的。

于是 `mysql_markdown` 就出现了、应用而生。它是一款基于 `go` 语言编写的一个命令行工具，适用于 `Linux`、`Mac`、`Windows`等平台。**那么它可以做什么？他只有一个功能、就是生成数据库表结构说明文档，格式为`markdown`**。



## 🪤 快速安装

> 哪位老哥协助帮忙发布到包管理器更是方便了 感恩~

到 [releases](https://github.com/alicfeng/mysql_markdown/releases) 页面下载运行，源码编译也行。



## 🚀 快速使用

```shell
# 帮助
➜  mysql_markdown -h
flag needs an argument: -h
Usage: mysql_markdown [options...]
--help  This help text
-h      host.     default 127.0.0.1
-u      username. default root
-p      password. default root
-d      database. default mysql
-P      port.     default 3306
-c      charset.  default utf8
-o      output.   default current location
-t      tables.   default all table and support ',' separator for filter, every item can use regexp
-l      language. default cn,support: cn,en

# 使用
➜ mysql_markdown -p samego -d samego
connecting to MySQL to export markdown document file.
 100% |█████████████████████████████████████████████████████████████████████████████████████| (31/31, 70 it/s)
convert to document finished in mysql_20220429_101148.md
```



# 💾 文档快照

![MySQL 表结构生成 Markdown 文档](https://upload-images.jianshu.io/upload_images/1678789-8050fa3687e575db.png)



## 🏷 特别说明

`markdown` 转其它类型的文档推荐使用 `typora` 工具 它支持如下转换格式

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



## License

mysql_markdown is licensed under the [MIT License](https://raw.githubusercontent.com/alicfeng/mysql_markdown/master/LICENSE).

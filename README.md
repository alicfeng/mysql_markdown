## mysql_markdown
It can generate markdown structure documents of MySQL succinctly~

#### install
```
curl -o /usr/local/bin -sSL https://raw.githubusercontent.com/alicfeng/mysql_markdown/master/mysql_markdown
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

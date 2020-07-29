FROM alpine:3.7
LABEL Maintainer="AlicFeng <a@samego.com>" \
      Description="mysql_markdown based on golang"

COPY release/mysql_markdown_unix /usr/local/sbin/mysql_markdown

RUN chmod a+x /usr/local/sbin/mysql_markdown && \
    mkdir /data/
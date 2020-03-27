#!/usr/bin/env bash

dl_dir="https://github.com/alicfeng/mysql_markdown/releases/"

if [[ "$(uname)" == "Darwin" ]]; then
    # Mac OS X
    echo -e "installing mysql_markdown_mac"
    curl -o /usr/local/bin/mysql_markdown -sSL "${dl_dir}mysql_markdown_mac"
    chmod a+x /usr/local/bin/mysql_markdown
elif [[ "$(expr substr $(uname -s) 1 5)" == "Linux" ]]; then
    # GNU/Linux
    echo "installing mysql_markdown_unix"
    curl -o /usr/local/bin/mysql_markdown -sSL "${dl_dir}mysql_markdown_unix"
    chmod a+x /usr/local/bin/mysql_markdown
elif [[ "$(expr substr $(uname -s) 1 10)" == "MINGW32_NT" ]]; then
    # Windows NT操作系统
    echo "installing mysql_markdown_win"
    curl -o mysql_markdown -sSL "${dl_dir}mysql_markdown_unix"
    #echo -e "sorry! please download source to build as well as install"
fi

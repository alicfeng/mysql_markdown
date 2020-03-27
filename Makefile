GC=go
MAIN_GO_FILE=mysql_markdown.go
RELEASE_DIR=release

build-unix :
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ${GC} build -ldflags -w -o ${RELEASE_DIR}/mysql_markdown_unix ${MAIN_GO_FILE}

build-mac :
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 ${GC} build -ldflags -w -o ${RELEASE_DIR}/mysql_markdown_mac ${MAIN_GO_FILE}

build-win :
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 ${GC} build -ldflags -w -o ${RELEASE_DIR}/mysql_markdown_win ${MAIN_GO_FILE}

build :
	make build-unix
	make build-mac
	make build-win

release :
	make build
	upx release/mysql_markdown_*
	tar -czvf release/mysql_markdown_mac.tar.gz -C release/ mysql_markdown_mac
	tar -czvf release/mysql_markdown_unix.tar.gz -C release/ mysql_markdown_unix
	tar -czvf release/mysql_markdown_win.tar.gz -C release/ mysql_markdown_win

clean :
	@if [ -d ${RELEASE_DIR} ] ; then rm -rf ${RELEASE_DIR}; fi
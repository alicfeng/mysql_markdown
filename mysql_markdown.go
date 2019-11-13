/*
Author   :    AlicFeng
Email    :    a@samego.com
Telegram :    https://t.me/AlicFeng
Github   :    https://github.com/alicfeng
*/

package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"time"
)

/**
Database Configuration
 */
var (
	host     = flag.String("h", "127.0.0.1", "host(127.0.0.1)")
	username = flag.String("u", "root", "username(root)")
	password = flag.String("p", "root", "password(root)")
	database = flag.String("d", "mysql", "database(mysql)")
	port     = flag.Int("P", 3306, "port(3306)")
	charset  = flag.String("c", "utf8", "charset(utf8)")
)

/**
Structured Query Language
 */
const (
	// 查看数据库所有数据表SQL
	SqlTables = "SELECT `table_name`,`table_comment` FROM `information_schema`.`tables` WHERE `table_schema`=?"
	// 查看数据表列信息SQL
	SqlTableColumn = "SELECT `ORDINAL_POSITION`,`COLUMN_NAME`,`COLUMN_TYPE`,`COLUMN_KEY`,`IS_NULLABLE`,`EXTRA`,`COLUMN_COMMENT`,`COLUMN_DEFAULT` FROM `information_schema`.`columns` WHERE `table_schema`=? AND `table_name`=?"
)

/**
struct for table column
 */
type tableColumn struct {
	OrdinalPosition int8   `db:"ORDINAL_POSITION"` // position
	ColumnName      string `db:"COLUMN_NAME"`      // name
	ColumnType      string `db:"COLUMN_TYPE"`      // type
	ColumnKey       string `db:"COLUMN_KEY"`       // key
	IsNullable      string `db:"IS_NULLABLE"`      // nullable
	Extra           string `db:"EXTRA"`            // extra
	ColumnComment   string `db:"COLUMN_COMMENT"`   // comment
	ColumnDefault   string `db:"COLUMN_DEFAULT"`   // default value
}

/**
struct for table message
 */
type tableInfo struct {
	Name    string `db:"table_name"`    // name
	Comment string `db:"table_comment"` // comment
}

/**
connect mysql service
 */
func connect() (*sql.DB, error) {
	// generate dataSourceName
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", *username, *password, *host, *port, *database, *charset)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("mysql connect failed, detail is [%v]", err.Error())
	}

	return db, err
}

/**
query table info about scheme
 */
func queryTables(db *sql.DB, dbName string) ([]tableInfo, error) {
	var tables []tableInfo

	rows, err := db.Query(SqlTables, dbName)
	if err != nil {
		return tables, err
	}

	for rows.Next() {
		var info tableInfo
		//将值存入变量name中
		err = rows.Scan(&info.Name, &info.Comment)
		if err != nil {
			break
		}

		tables = append(tables, info)
	}

	return tables, err
}

/**
查询数据表列信息
 */
func queryTableColumn(db *sql.DB, dbName string, tableName string) ([]tableColumn, error) {
	// 定义承载列信息的切片
	var columns []tableColumn

	rows, err := db.Query(SqlTableColumn, dbName, tableName)
	if err != nil {
		return columns, err
	}
	for rows.Next() {
		var column tableColumn
		rows.Scan(&column.OrdinalPosition, &column.ColumnName, &column.ColumnType, &column.ColumnKey, &column.IsNullable, &column.Extra, &column.ColumnComment, &column.ColumnDefault)
		columns = append(columns, column)
	}

	return columns, err
}

func main() {
	flag.Parse()
	// connect mysql service
	var db, connectErr = connect()
	if connectErr != nil {
		fmt.Println("mysql connect service fail ...")
		return
	}
	fmt.Println("mysql connected ...")

	// query all table name
	var tables, tablesErr = queryTables(db, *database)
	if tablesErr != nil {
		fmt.Println("query tables of database error ...")
		return
	}

	// create and open markdown file
	var mdFileName = *database + "_" + time.Now().Format("20060102_150405") + ".md"
	mdFile, err := os.OpenFile(mdFileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		println(err.Error())
		return
	}

	//
	var tableContent = "## " + *database + " tables message\n"
	for index, table := range tables { //range returns both the index and value
		fmt.Printf("%d/%d the %s table is making ...\n", index+1, len(tables), table.Name)
		tableContent += "###### " + table.Name
		if table.Comment != "" {
			tableContent += "( " + table.Comment + " )"
		}
		tableContent += "\n" +
			"| 序号 | 名称 | 描述 | 类型 | 允许为空 | 默认值 |\n" +
			"| :--: | :--: | :--: | :--: | :--: | :--: |\n"
		var columnInfo, columnInfoErr = queryTableColumn(db, *database, table.Name)
		if columnInfoErr != nil {
			continue
		}
		for _, info := range columnInfo {
			tableContent += fmt.Sprintf(
				"| %d | %s | %s | %s | %s | %s |\n",
				info.OrdinalPosition,
				info.ColumnName,
				info.ColumnComment,
				info.ColumnType,
				info.IsNullable,
				info.ColumnDefault,
			)
		}
	}
	tableContent += "\n\n"
	mdFile.WriteString(tableContent)

	// close database and file handler for release
	db.Close()
	mdFile.Close()
	fmt.Println("mysql_markdown finished ...")
}


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
	"strconv"
	"strings"
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
	output   = flag.String("o", "", "output location")
	tables   = flag.String("t", "", "choose tables")
	version  = flag.Bool("v", false, "show version and exit")
	detail  = flag.Bool("V", false, "show version and exit")
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
	OrdinalPosition int8           `db:"ORDINAL_POSITION"` // position
	ColumnName      string         `db:"COLUMN_NAME"`      // name
	ColumnType      string         `db:"COLUMN_TYPE"`      // type
	ColumnKey       sql.NullString `db:"COLUMN_KEY"`       // key
	IsNullable      string         `db:"IS_NULLABLE"`      // nullable
	Extra           sql.NullString `db:"EXTRA"`            // extra
	ColumnComment   sql.NullString `db:"COLUMN_COMMENT"`   // comment
	ColumnDefault   sql.NullString `db:"COLUMN_DEFAULT"`   // default value
}

/**
struct for table message
*/
type tableInfo struct {
	Name    string         `db:"table_name"`    // name
	Comment sql.NullString `db:"table_comment"` // comment
}

/**
connect mysql service
*/
func connect() (*sql.DB, error) {
	// generate dataSourceName
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", *username, *password, *host, *port, *database, *charset)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("mysql sql open failed, detail is [%v]", err)
		return db, err
	}

	return db, err
}

/**
query table info about scheme
*/
func queryTables(db *sql.DB, dbName string) ([]tableInfo, error) {
	var tableCollect []tableInfo
	var tableArray []string
	var commentArray []sql.NullString

	rows, err := db.Query(SqlTables, dbName)
	if err != nil {
		return tableCollect, err
	}

	for rows.Next() {
		var info tableInfo
		err = rows.Scan(&info.Name, &info.Comment)
		if err != nil {
			fmt.Printf("execute query tables action error,had ignored, detail is [%v]\n", err.Error())
			continue
		}

		tableCollect = append(tableCollect, info)
		tableArray = append(tableArray, info.Name)
		commentArray = append(commentArray, info.Comment)
	}
	// filter tables when specified tables params
	if *tables != "" {
		tableCollect = nil
		chooseTables := strings.Split(*tables, ",")
		for _, item := range chooseTables {
			// handler -t params | table filter
			// not contain
			containIndex := containsString(tableArray, item)
			if -1 == containIndex {
				fmt.Printf("\033[33mthe %s table is not exist\033[0m \n", item)
				continue
			}
			// contain
			var info tableInfo
			info.Name = tableArray[containIndex]
			info.Comment = commentArray[containIndex]
			tableCollect = append(tableCollect, info)
		}
	}

	return tableCollect, err
}

/**
query table column message
*/
func queryTableColumn(db *sql.DB, dbName string, tableName string) ([]tableColumn, error) {
	// 定义承载列信息的切片
	var columns []tableColumn

	rows, err := db.Query(SqlTableColumn, dbName, tableName)
	if err != nil {
		fmt.Printf("execute query table column action error, detail is [%v]\n", err.Error())
		return columns, err
	}
	for rows.Next() {
		var column tableColumn
		err = rows.Scan(&column.OrdinalPosition, &column.ColumnName, &column.ColumnType, &column.ColumnKey, &column.IsNullable, &column.Extra, &column.ColumnComment, &column.ColumnDefault)
		if err != nil {
			fmt.Printf("query table column scan error, detail is [%v]\n", err.Error())
			return columns, err
		}
		columns = append(columns, column)
	}

	return columns, err
}

/**
string array contain string func
 */
func containsString(array []string, val string) (int) {
	for i := 0; i < len(array); i++ {
		if array[i] == val {
			return i
		}
	}
	return -1
}

/**
init func
*/
func init() {
	// init flag for command
	flag.CommandLine.Usage = func() {
		fmt.Println("Usage: mysql_markdown [options...]\n" +
			"--help  This help text" + "\n" +
			"-h      host.     default 127.0.0.1" + "\n" +
			"-u      username. default root" + "\n" +
			"-p      password. default root" + "\n" +
			"-d      database. default mysql" + "\n" +
			"-P      port.     default 3306" + "\n" +
			"-c      charset.  default utf8" + "\n" +
			"-o      output.   default current location\n" +
			"-t      tables.   default all table and support ',' separator for filter" +
			"")
		os.Exit(0)
	}
	flag.Parse()
	if *version {
		fmt.Println("mysql_markdown version: 1.0.1")
		os.Exit(0)
	}
	if *detail {
		fmt.Println(
			"mysql_markdown version: 1.0.1\n" +
				"build by golang 2019.11.27\n" +
				"author		AlicFeng\n" +
				"tutorial	https://github.com/alicfeng/mysql_markdown\n" +
				"价值源于技术,技术源于分享" +
				"")
		os.Exit(0)
	}
}

/**
main func
*/
func main() {
	// connect mysql service
	db, connectErr := connect()
	if connectErr != nil {
		fmt.Printf("\033[31mmysql sql open failed ... \033[0m \n")
		return
	}

	// query all table name
	tables, err := queryTables(db, *database)
	if err != nil {
		fmt.Printf("\033[31mquery tables of database error ... \033[0m \n%v\n", err.Error())
		return
	}

	// create and open markdown file
	if *output == "" {
		// automatically generated if no output file path is specified
		*output = *database + "_" + time.Now().Format("20060102_150405") + ".md"
	}
	mdFile, err := os.OpenFile(*output, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Printf("\033[31mcreate and open markdown file error \033[0m \n%v\n", err.Error())
		return
	}

	// make markdown format content
	var tableContent = "## " + *database + " tables message\n"
	for index, table := range tables {
		// make content process log
		fmt.Printf("%d/%d the %s table is making ...\n", index+1, len(tables), table.Name)

		// markdown header title
		tableContent += "#### " + strconv.Itoa(index+1) + "、 " + table.Name + "\n"
		if table.Comment.String != "" {
			tableContent += table.Comment.String + "\n"
		}

		// markdown table header
		tableContent += "\n" +
			"| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |\n" +
			"| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |\n"
		var columnInfo, columnInfoErr = queryTableColumn(db, *database, table.Name)
		if columnInfoErr != nil {
			continue
		}
		for _, info := range columnInfo {
			tableContent += fmt.Sprintf(
				"| %d | `%s` | %s | %s | %s | %s | %s | %s |\n",
				info.OrdinalPosition,
				info.ColumnName,
				info.ColumnComment.String,
				info.ColumnType,
				info.ColumnKey.String,
				info.IsNullable,
				info.Extra.String,
				info.ColumnDefault.String,
			)
		}
		tableContent += "\n\n"
	}
	mdFile.WriteString(tableContent)

	// close database and file handler for release
	err = db.Close()
	err = mdFile.Close()
	fmt.Printf("\033[32mmysql_markdown finished ... \033[0m \n")
}

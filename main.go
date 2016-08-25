package main

import (
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbString := flag.String("d", "", "mysql instance")
	sqlFile := flag.String("f", "", "output sql to file")
	flag.Parse()

	db, err := sql.Open("mysql", *dbString)
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(os.Stderr, "connect to test db okay")

	rows, err := db.Query("show tables")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	allSql := ""
	for rows.Next() {
		var table string
		err = rows.Scan(&table)
		if err != nil {
			fmt.Fprintf(os.Stderr, "scan test row error: %v\n", err)
			return
		}

		row := db.QueryRow("show create table " + table)
		if err != nil {
			fmt.Fprintf(os.Stderr, "query show create error: %v\n", err)
			return
		}
		var table1, create string
		err = row.Scan(&table1, &create)
		if err != nil {
			fmt.Fprintf(os.Stderr, "scan create table error: %v\n", err)
			return
		}
		create = strings.Replace(create, "CREATE TABLE", "CREATE TABLE IF NOT EXISTS", 1)

		createSql := fmt.Sprintf("DROP TABLE IF EXISTS `%s`;\n%s;", table, create)
		allSql += fmt.Sprintf("-- create table %s\n\n%s\n\n", table, createSql)
	}

	if *sqlFile == "" {
		fmt.Fprintf(os.Stdout, "-- Database\n\n%s\n", allSql)
	} else {
		err = ioutil.WriteFile(*sqlFile, []byte(allSql), 0666)
		if err != nil {
			fmt.Fprintf(os.Stderr, "write file to %s error: %v\n", *sqlFile, err)
		}
	}
}

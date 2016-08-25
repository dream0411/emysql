# emysql - export mysql
---

emysql can connect to mysql server, then use "show tables", "show create table xxx" to export structure of tables, to avoid no enought permission for mysqldump.

## INSTALL

```
go install github.com/dream0411/emysql
```

## USAGE

```
emysql -d 'user:passwd@tcp(127.0.0.1:3306)/gewara?charset=%s&loc=Local' -f 'db.sql'
```

## Something else

If you don't want write password in command line, just "go get" source file and put your connection string into main.go, compile and run.

package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/haifahrul/internal/helper/env"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
	_ "github.com/lib/pq"
)

//PostgreSql Helper for PostgreSQL connect
type PostgreSql struct {
}

//RowCount Just count PostgreSQL result
func (*PostgreSql) RowCount(rows *sql.Rows) (i int) {
	columns, err := rows.Columns()
	if err != nil {
		panic(err)
	}
	columnCount := len(columns)

	values := make([]interface{}, columnCount)
	valuePtrs := make([]interface{}, columnCount)
	for i := range columns {
		valuePtrs[i] = &values[i]
	}

	for rows.Next() {
		err := rows.Scan(valuePtrs...)
		if err != nil {
			panic(err)
		}
		i++
	}

	for i := range values {
		fmt.Print(values[i])
	}

	return
}

//RowParse Parse database result in map rows with count
func (*PostgreSql) RowParse(rows *sql.Rows) (records map[int]map[string]interface{}, count int) {
	columns, err := rows.Columns()
	if err != nil {
		panic(err)
	}
	columnCount := len(columns)

	records = make(map[int]map[string]interface{})

	for rows.Next() {

		values := make([]interface{}, columnCount)
		valuePtrs := make([]interface{}, columnCount)
		for i := range columns {
			valuePtrs[i] = &values[i]
		}

		err := rows.Scan(valuePtrs...)
		if err != nil {
			panic(err)
		}
		row := make(map[string]interface{})

		for i := range columns {
			row[columns[i]] = values[i]
		}
		records[count] = row
		count++
	}

	return

}

//ConnStr Generate Connection string of PostgreSQL from Env vars
func (*PostgreSql) ConnStr() string {

	var (
		host     = os.Getenv("PG_HOST")
		port     = os.Getenv("PG_PORT")
		user     = os.Getenv("PG_USER")
		password = os.Getenv("PG_PASSWORD")
		dbname   = os.Getenv("PG_DBNAME")
		sslmode  = os.Getenv("PG_SSLMODE")
		appname  = os.Getenv("APP_NAME")
	)

	if host == "docker-host" && env.Get.IsDev() {
		host = os.Getenv("DOCKER_HOST")
	}

	if os.Getenv("GO_ENV") == "test" {
		host = os.Getenv("TEST_PG_HOST")
		port = os.Getenv("TEST_PG_PORT")
		user = os.Getenv("TEST_PG_USER")
		password = os.Getenv("TEST_PG_PASSWORD")
		dbname = os.Getenv("TEST_PG_DBNAME")
		sslmode = os.Getenv("TEST_PG_SSLMODE")
	}

	password = url.QueryEscape(password)

	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s&fallback_application_name=%s",
		user, password, host, port, dbname, sslmode, appname,
	)

	return connStr
}

//Connect Generate Database connection handler
func (_p *PostgreSql) Connect() (postgres *sql.DB, err error) {
	connStr := _p.ConnStr()

	postgres, err = sql.Open("postgres", connStr)

	if err != nil {
		return
	}

	err = postgres.Ping()

	if err != nil {
		return
	}

	if env.Get.IsDev() || env.Get.IsTest() {
		log.Println("Connected to PostgreSQL Server: ", connStr)
	}

	return
}

//ConnectSqlx Connecting with Sqlx
func (_p *PostgreSql) ConnectSqlx() (db *sqlx.DB, err error) {
	connStr := _p.ConnStr()
	postgres, err := sql.Open("postgres", connStr)
	db = sqlx.NewDb(postgres, "postgres")
	if err != nil {
		return
	}
	err = db.Ping()
	db.Mapper = reflectx.NewMapperFunc("json", strings.ToLower)
	if err == nil {
		if env.Get.IsDev() || env.Get.IsTest() {
			log.Println("Connected to PostgreSQL Server: ", connStr)
		} else {

			log.Println("Connected to PostgreSQL Server")
		}
	} else {
		if env.Get.IsDev() || env.Get.IsTest() {
			log.Println("Connection failed to PostgreSQL Server: ", connStr)
		} else {
			log.Println("Connection failed to PostgreSQL Server:", os.Getenv("PG_HOST"))
		}
	}
	return
}

//Get Singleton export
var Get = &PostgreSql{}

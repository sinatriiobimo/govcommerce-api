package postgre

import (
	"database/sql"
	"github.com/lib/pq"
	sqltrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/database/sql"
	"log"
	"time"
	config "tlkm-api/configs"
)

var (
	dbTelkomRead  *sql.DB
	dbTelkomWrite *sql.DB
	err           error
)

func NewDBTelkomRead() {
	sqltrace.Register("postgres", &pq.Driver{}, sqltrace.WithServiceName("postgresql"))

	dbTelkomRead, err = sqltrace.Open("postgres", config.Get().Postgre.Telkom.Read)
	if err != nil {
		log.Panic("failed to open postgre client for telkom read: ", err)
	}

	dbTelkomRead.SetMaxOpenConns(350)
	dbTelkomRead.SetMaxIdleConns(10)
	dbTelkomRead.SetConnMaxLifetime(time.Second * 10)
}

func GetDBTelkomRead() *sql.DB {
	if dbTelkomRead == nil {
		NewDBTelkomRead()
	}
	return dbTelkomRead
}

func NewDBTelkomWrite() {
	sqltrace.Register("postgres", &pq.Driver{}, sqltrace.WithServiceName("postgresql"))

	dbTelkomWrite, err = sqltrace.Open("postgres", config.Get().Postgre.Telkom.Write)
	if err != nil {
		log.Panic("failed to open postgre client for telkom write: ", err)
	}
	dbTelkomWrite.SetMaxOpenConns(350)
	dbTelkomWrite.SetMaxIdleConns(10)
	dbTelkomWrite.SetConnMaxLifetime(time.Second * 10)
}

func GetDBTelkomWrite() *sql.DB {
	if dbTelkomWrite == nil {
		NewDBTelkomWrite()
	}
	return dbTelkomWrite
}

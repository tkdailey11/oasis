package db

import (
	"fmt"
)

type DBConfig struct {
	ConnectionString string
	DBName string
}

type DBType int32
const (
	SQL DBType = iota
	Mongo
)
func (t DBType) String() string {
	switch t {
	case SQL:
		return "sql"
	case Mongo:
		return "mongo"
	default:
		return "INVALID TYPE"
	}
}

type DBProvider int32
const (
	Azure DBProvider = iota
	AWS
	Google
)
func (p DBProvider) String() string {
	switch p {
	case Azure:
		return "azure"
	case AWS:
		return "aws"
	case Google:
		return "google"
	default:
		return "INVALID PROVIDER"
	}
}

//var dbConfig DBConfig
// var confMu = &sync.Mutex{}

// func SetConfig(config DBConfig) {
// 	confMu.Lock()
// 	defer confMu.Unlock()

// 	dbConfig = config
// }

func Test(t DBType) {
	fmt.Println(t)
}
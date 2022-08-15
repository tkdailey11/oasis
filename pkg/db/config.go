package db

import "sync"

type DBConfig struct {
	Type DBType
	ConnectionString string
	DBName string
}

type DBType int32
const (
	Mongo DBType = iota
)
func (t DBType) String() string {
	switch t {
	case Mongo:
		return "mongo"
	default:
		return "INVALID TYPE"
	}
}

type DBProvider int32
const (
	Azure DBProvider = iota
)
func (p DBProvider) String() string {
	switch p {
	case Azure:
		return "azure"
	default:
		return "INVALID PROVIDER"
	}
}

var dbConfig DBConfig
var confMu = &sync.Mutex{}

func SetConfig(config DBConfig) {
	confMu.Lock()
	defer confMu.Unlock()

	dbConfig = config
}

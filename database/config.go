package database

import (
	"encoding/json"
	"errors"
)

type ConfigDatabase struct {
	Type           TypeDatabase `json:"type"`
	FileDataSource string       `json:"file_data_source"`
	Debug          bool         `json:"debug"`
	Migration      bool         `json:"migration"`
	Options        map[string]interface{}
}

func convertStruct(src interface{}, des interface{}) error {
	dataBase, err := json.Marshal(src)
	if err != nil {
		return err
	}
	return json.Unmarshal(dataBase, des)
}

func (c ConfigDatabase) Sqlite() (source string, err error) {
	source = "dipper.db:ent?_fk=1"

	raw, ok := c.Options["file_source"]
	if ok {
		source, ok = raw.(string)
		if !ok {
			err = errors.New("file_source not string")
		}
	}

	return
}

func (c ConfigDatabase) Postgrest(name string) (pg *PostgrestConfig, err error) {
	pg = &PostgrestConfig{
		Host:     "127.0.0.1",
		Port:     5432,
		Database: "dipper_engine",
	}

	raw, ok := c.Options[name]
	if ok {
		err = convertStruct(raw, pg)
	}

	return
}

type PostgrestConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string
	Password string `json:"password"`
	Database string `json:"database"`
}

type TypeDatabase string

const (
	SqliteType           TypeDatabase = "sqlite"
	PostgrestType        TypeDatabase = "postgres"
	PostgrestClusterType TypeDatabase = "postgres_cluster"
)

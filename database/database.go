package database

import (
	"context"
	"fmt"
	"github.com/dipper-iot/dipper-engine-server/ent"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Database struct {
	config      *ConfigDatabase
	clientRead  *ent.Client
	clientWrite *ent.Client
}

func NewDatabase(config *ConfigDatabase) *Database {
	return &Database{config: config}
}

func (d *Database) Client() (clientRead *ent.Client, clientWrite *ent.Client) {
	clientRead = d.clientRead
	clientWrite = d.clientWrite
	return
}

func (d *Database) connectPG(name string) (client *ent.Client, err error) {
	c, err := d.config.Postgrest(name)
	if err != nil {
		return nil, err
	}

	client, err = ent.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s", c.Host, c.Port, c.User, c.Database, c.Password))
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
		return nil, err
	}

	return
}

func (d *Database) Connect(ctx context.Context) error {
	var (
		clientRead  *ent.Client
		clientWrite *ent.Client
	)

	switch d.config.Type {
	case PostgrestType:
		{

			client, err := d.connectPG("pg")
			if err != nil {
				log.Fatalf("failed opening connection to postgres: %v", err)
				return err
			}
			clientWrite = client
			clientRead = client

			break
		}
	case PostgrestClusterType:
		{
			var err error
			clientRead, err = d.connectPG("pg_read")
			if err != nil {
				log.Fatalf("failed opening connection to postgres: %v", err)
				return err
			}
			clientWrite, err = d.connectPG("pg_write")
			if err != nil {
				log.Fatalf("failed opening connection to postgres: %v", err)
				return err
			}
			break
		}
	default:
		{
			stringConnect, err := d.config.Sqlite()
			if err != nil {
				return err
			}

			client, err := ent.Open("sqlite3", stringConnect)
			if err != nil {
				log.Fatalf("failed opening connection to sqlite: %v", err)
				return err
			}
			clientWrite = client
			clientRead = client
		}
	}

	if d.config.Debug {
		d.clientRead = clientRead.Debug()
		d.clientWrite = clientWrite.Debug()
	} else {
		d.clientRead = clientRead
		d.clientWrite = clientWrite
	}

	if d.config.Migration {
		// Run the auto migration tool.
		if err := d.clientWrite.Schema.Create(ctx); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
			return err
		}
	}

	return nil
}

func (d *Database) Close() {
	if d.clientWrite != nil {
		d.clientWrite.Close()
	}
	if d.clientRead != nil {
		d.clientRead.Close()
	}
}

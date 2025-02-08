package config

import (
	"github.com/omniful/go_commons/db/sql/postgres"
	"github.com/omniful/go_commons/redis"
)

var Redis_Config = &redis.Config{
	ClusterMode: false,
	Hosts:       []string{"localhost:6379"},
	PoolSize:    1,
}

var Postgres_Config = &postgres.DBConfig{
	Host:     "localhost",
	Port:     "5432",
	Username: "postgres",
	Password: "root",
	Dbname:   "wms_db",
}

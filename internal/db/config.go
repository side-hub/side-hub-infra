package db

import (
	"encoding/json"
	"fmt"
)

type Config struct {
	User     string
	Password string
	Server   string
	Name     string
	Port     int
	MaxConns int
	DBMS     string
}

func (c Config) DSN() string {
	return fmt.Sprintf("%s:%s@(%s:%d)/%s?loc=Local", c.User, c.Password, c.Server, c.Port, c.Name)
}

func (c Config) String() string {
	return func() string { ret, _ := json.Marshal(c); return string(ret) }()
}

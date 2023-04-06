package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Config struct {
	Username  string
	Password  string
	DBName    string
	DSNParams string
}

type Client struct {
	cfg *Config
}

func NewClient(cfg *Config) *Client {
	c := new(Client)
	c.cfg = cfg
	return c
}

func (c *Client) Connect() (*sql.DB, error) {
	auth := ""
	if c.cfg.Username != "" && c.cfg.Password != "" {
		auth = fmt.Sprintf("?_auth&_auth_user=%s&_auth_pass=%s", c.cfg.Username, c.cfg.Password)
	}
	dsn := auth
	if c.cfg.DSNParams != "" {
		if dsn == "" {
			dsn += "?" + c.cfg.DSNParams
		} else {
			dsn += "&" + c.cfg.DSNParams
		}
	}
	db, err := sql.Open("sqlite3", fmt.Sprintf("%s%s", c.cfg.DBName, auth))
	if err != nil {
		return nil, err
	}
	return db, err
}

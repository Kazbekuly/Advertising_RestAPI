package repository

import (
	"Advertising/configs"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresDB(cfg *configs.Config) (*sqlx.DB, error) {
	fmt.Println(cfg.Db.Port)
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Db.Host, cfg.Db.Port, cfg.Db.Username, cfg.Db.Dbname, cfg.Db.Password, cfg.Db.SslMode))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	return db, nil
}

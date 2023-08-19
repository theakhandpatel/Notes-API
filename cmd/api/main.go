package main

import (
	"database/sql"
	"flag"
	"fmt"
	"movieapi/internal/data"
	"net/http"

	_ "github.com/glebarez/go-sqlite"
)

type config struct {
	port int
	env  string
	dsn  string
}

type application struct {
	config config
	models data.Models
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 8080, "API server port")
	flag.StringVar(&cfg.env, "env", "dev", "Environment (dev | prod )")
	flag.StringVar(&cfg.dsn, "dsn", "./database.db", "SQLite DSN | path")

	flag.Parse()

	db, err := openDB(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	app := application{
		config: cfg,
		models: data.NewModels(db),
	}

	router := app.routes()
	addr := fmt.Sprintf(":%d", cfg.port)

	fmt.Println("Server Starting in", cfg.env, "mode at :", cfg.port)
	err = http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Server running in %v mode at :%v", cfg.env, cfg.port)
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("sqlite", cfg.dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS notes(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title VARCHAR(255) NOT NULL,
			body VARCHAR(4000),
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME NOT NULL 
		)
		`)

	if err != nil {
		return nil, err
	}

	return db, nil
}

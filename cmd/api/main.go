package main

import (
	"log"
	"scalfold/internal/db"
	"scalfold/internal/env"
	"scalfold/internal/store"
)

func main() {

	cfg := config{
		addr: env.GetEnvString("ADDR", ":3000"),
		db: dbConfig{
			addr:         env.GetEnvString("DB_CONN_URL", "postgres://randomdev@localhost:5432/scafold?sslmode=disable"),
			maxOpenConns: env.GetEnvInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetEnvInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetEnvString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db, err := db.NewConn(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)

	if err != nil {
		log.Println(err, "Error")
	} else {
		log.Println("Database Connected")
	}

	defer db.Close()

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))

}

package db

import (
	"api/src/config"
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	// connStr := "user=iagopzivitovski password=k0s3emRdHXIN dbname=neondb host=ep-lucky-cell-100773.us-east-2.aws.neon.tech sslmode=verify-full"

	db, erro := sql.Open("postgres", config.StringConexaoBanco)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}
	return db, nil
}

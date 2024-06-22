package dbrepo

import (
	"context"
	"time"

	"github.com/otavio-Pucharelli/filhos-da-luz/internal/models"
)

func (m *PostgresDBRepo) AllUsers() bool {
	return true
}

// InsertResident inserts a new resident into the database
func (m *PostgresDBRepo) InsertResident(res models.Resident) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `INSERT INTO residents 
	(name, cpf, address, city, state, zip, created_at, updated_at)
	VAlUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	_, err := m.DB.Exec(ctx, stmt,
		res.Name,
		res.CPF,
		res.Address,
		res.City,
		res.State,
		res.Zip,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return err
	}

	return nil
}

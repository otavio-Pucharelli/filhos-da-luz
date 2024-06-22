package repository

import "github.com/otavio-Pucharelli/filhos-da-luz/internal/models"

type DatabaseRepo interface {
	AllUsers() bool
	InsertResident(res models.Resident) error
}

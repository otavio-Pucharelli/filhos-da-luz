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
	(name, social_name, cpf, rg, birth_date, sex, marital_status,
	 situation, type_of_residence, current_unit, reception_date,
	 departure_date, address, city, state, zip, father_name,
	 mother_name, payment_responsable_name, payment_responsable_cpf, 
	 responsable_phone, responsable_degree_of_kinship, judiciary, 
	 benefits, suicide, hiv_others_sti, indication, observations created_at, updated_at)
	VAlUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, 
	$11, $12, $13, $14, $15, $16, $17, $18, $19, $20, 
	$21, $22, $23, $24, $25, $26, $27, $28, $29, $30)`

	_, err := m.DB.Exec(ctx, stmt,
		res.Name,
		res.SocialName,
		res.CPF,
		res.RG,
		res.BirthDate,
		res.Sex,
		res.MaritalStatus,
		res.Situation,
		res.TypeOfResidence,
		res.CurrentUnit,
		res.RecptionDate,
		res.DepartureDate,
		res.Address,
		res.City,
		res.State,
		res.Zip,
		res.FatherName,
		res.MotherName,
		res.PaymentResponsableName,
		res.PaymentResponsableCPF,
		res.ResponsablePhone,
		res.ResponsableDegreeOfKinship,
		res.Judiciary,
		res.Benefits,
		res.Suicide,
		res.HivOtherSti,
		res.Indication,
		res.Observations,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return err
	}

	return nil
}

package models

import "time"

// Users is the user model
type User struct {
	ID          int
	Name        string
	Email       string
	Password    string
	AccessLevel string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

//Residents is the resident model
type Resident struct {
	ID                         int
	ResidentPhoto              []byte
	Name                       string
	SocialName                 string
	CPF                        string
	RG                         string
	DocumentPhoto              []byte
	BirthDate                  time.Time
	Sex                        int
	MaritalStatus              int
	Situation                  int
	TypeOfResidence            int
	CurrentUnit                int
	RecptionDate               time.Time
	DepartureDate              time.Time
	Address                    string
	City                       string
	State                      string
	Zip                        string
	FatherName                 string
	MotherName                 string
	PaymentResponsableName     string
	PaymentResponsableCPF      string
	ResponsablePhone           string
	ResponsableDegreeOfKinship string
	Judiciary                  string
	Benefits                   string
	Suicide                    bool
	HivOtherSti                string
	Indication                 string
	Observations               string
	CreatedAt                  time.Time
	UpdatedAt                  time.Time
}

//MultidisciplinaryEvolutions is The MultidisciplinaryEvolution model
type MultidisciplinaryEvolution struct {
	ID         int
	IDResident int
	Date       time.Time
	Message    string
	Resident   Resident
}

package checker

import (
	"log"
	"testing"
	"time"
)

type enumTest int

const (
	enumTest0 enumTest = iota
	enumTest1
	enumTest2
)

type validatorStruct struct {
	Enum            enumTest   `validate:"enum"`
	PhoneUs         string     `validate:"required,phone_us"`
	PhoneBr         string     `validate:"required,phone_br"`
	FullName        string     `validate:"required,full_name"`
	BCrypt          string     `validate:"required,bcrypt"`
	PostalCode      string     `validate:"required,postal_code"`
	Bearer          string     `validate:"required,bearer"`
	BeforeNow       time.Time  `validate:"required,before_now"`
	BeforeToday     time.Time  `validate:"required,before_today"`
	AfterNow        time.Time  `validate:"required,after_now"`
	AfterToday      time.Time  `validate:"required,after_today"`
	Today           time.Time  `validate:"required,today"`
	Now             time.Time  `validate:"required,now"`
	FullNow         time.Time  `validate:"required,full_now"`
	Cpf             string     `validate:"required,cpf"`
	Cnpj            string     `validate:"required,cnpj"`
	CpfCnpj         string     `validate:"required,cpfcnpj"`
	MultipleMongoDb []string   `validate:"required,mongodb"`
	MultipleEnum    []enumTest `validate:"required,enum"`
}

func (e enumTest) IsEnumValid() bool {
	return e == enumTest0 || e == enumTest1 || e == enumTest2
}

func initValidatorStruct() validatorStruct {
	before := time.Date(1999, 1, 1, 0, 0, 0, 0, time.Local)
	after := time.Date(2100, 1, 1, 0, 0, 0, 0, time.Local)
	now := time.Now()
	return validatorStruct{
		Enum:         enumTest0,
		PhoneUs:      "(717) 550-1675",
		PhoneBr:      "47997576131",
		FullName:     "Gabriel Cataldo",
		BCrypt:       "$2a$10$1W70eIOi/iLSPMxRVY9w5OXFalSffXgVzP6u4W/7vmusN4MQd6QL.",
		PostalCode:   "89041-0001",
		Bearer:       "Bearer HN4I6CI4Sbx2zYmv",
		BeforeNow:    after,
		BeforeToday:  after,
		AfterNow:     before,
		AfterToday:   before,
		Today:        now,
		Now:          time.Now(),
		FullNow:      now,
		Cpf:          "11664947051",
		Cnpj:         "52977110000101",
		CpfCnpj:      "11664947051",
		MultipleEnum: []enumTest{enumTest0},
	}
}

func TestValidator(t *testing.T) {
	err := Validate().Struct(initValidatorStruct())
	log.Println("Validate Struct Result:", err)
}

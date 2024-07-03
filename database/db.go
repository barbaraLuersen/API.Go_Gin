package database

import (
	"log"

	"github.com/barbaraLuersen/GoGin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	stringDeConexao := "host=localhost user=root password=rootAdmin dbname=DB_SERIES port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
	DB.AutoMigrate(&models.Aluno{})
}

//O pacote de db realiza a conexão e faz esse DB.AutoMigrate,
//no nosso banco de dados, com base no nosso modelo de aluno.
//E nós inserimos esse gorm.Model, para ele incluir o ID, o
//created at, o updated at e o deleted at.

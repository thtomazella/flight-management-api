package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// StringConexaoBanco é a string de conexão com o MYSQL
	StringConexaoBanco = ""

	// Porta onde a API vai estar rodando
	Porta = 0

	// KEY para acesso ao REDEMET
	SecretKey_Redemet = ""

	// SecretKey é a chave que vai ser usada para assinar os tokens
	SecretKey []byte
)

// Carregar vai inicializar as variáveis de ambiente
func Carregar() {
	var erro error
	// Função para carregar o arquivo .env
	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Porta = 9000
	}

	SecretKey_Redemet = os.Getenv("API_KEY_REDEMET")

	StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"))

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}

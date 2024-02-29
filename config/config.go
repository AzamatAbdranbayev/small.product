package config

import (
	"log"
	"os"
)

type Config struct {
	Http *HttpConfig
	Db   *DbConfig
	Smtp *SmtpConfig
}
type DbConfig struct {
	User   string
	Pass   string
	DbName string
	Host   string
	Port   string
}
type HttpConfig struct {
	Address string
}
type SmtpConfig struct {
	ProducerEmail string
	Token         string
	SmtpHost      string
	SmtpName      string
}

func NewConfig() (*Config, error) {
	log.Println("Loading configuration")

	var config Config
	//TODO: можно читать с конфигурационного файла, лучше чтобы в пайпе он генерился
	// а так пока можно напрямую через переменные окружения
	config.Db = &DbConfig{
		User:   os.Getenv("DB_USER"),
		Pass:   os.Getenv("DB_PASS"),
		DbName: os.Getenv("DB_NAME"),
		Host:   os.Getenv("DB_HOST"),
		Port:   os.Getenv("DB_PORT"),
	}

	config.Http = &HttpConfig{Address: os.Getenv("SERVICE_PORT")}

	config.Smtp = &SmtpConfig{
		ProducerEmail: os.Getenv("SMTP_PRODUCER"),
		Token:         os.Getenv("SMTP_TOKEN"),
		SmtpHost:      os.Getenv("SMTP_HOST"),
		SmtpName:      os.Getenv("SMT_NAME"),
	}

	return &config, nil
}

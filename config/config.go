package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Setup() {
	// Carga el archivo .env si existe
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("No se encontró .env, usando variables de entorno del sistema")
	}

	// Ejemplo de lectura de una variable (puede ser usada más adelante en tu app)
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL no está definida")
	}
}

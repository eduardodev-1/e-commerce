package config

import (
	"flag"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path"
	"strconv"
)

func init() {
	loadFile, err := strconv.ParseBool(os.Getenv("LOAD_FILE"))
	if err != nil {
		loadFile = true
	}
	if loadFile {
		isImageDocker, err := strconv.ParseBool(os.Getenv("IS_IMAGE_DOCKER"))
		if err != nil {
			isImageDocker = false
		}
		if isImageDocker {
			err = os.Setenv("DB_HOST", "host.docker.internal")
			if err != nil {
				log.Println("Error setting DB_HOST")
			}
		}
		log.Print("load env = true")
		loadEnv()
	} else {
		log.Print("load env = false")
	}
}
func loadEnv() {
	// Define the environment flag
	env := flag.String("env", "local", "Environment for the application")
	flag.Parse()

	log.Printf("Running on %s environment\n", *env)

	// Load selected env
	currentEnv := path.Join("env", *env+".env")
	if err := godotenv.Load(currentEnv); err != nil {
		log.Printf("Erro ao carregar arquivo .env: %v", err)
	} else {
		os.Setenv("APP_ENV", *env)
	}

	// Load common env
	commonEnv := path.Join("env", "common.env")
	if err := godotenv.Load(commonEnv); err != nil {
		log.Fatalf("Erro ao carregar arquivo .env: %v", err)
	}
}

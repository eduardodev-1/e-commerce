package config

import (
	"e-commerce/internal/utils"
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
		log.Print("load env = true")
		loadEnv()
	} else {
		log.Print("load env = false")
	}
}
func loadEnv() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local"
	}
	rootDir := utils.GetCurrentRootDir()
	log.Print(env)
	if env == "test" {
		err := godotenv.Load(".env.test")
		if err != nil {
			log.Println(err.Error())
			return
		}
	} else {
		envPath := path.Join(rootDir, "env", "local.env")
		log.Printf("Caminho do arquivo env: %s", envPath)
		err := godotenv.Load(envPath)
		if err != nil {
			log.Println(err.Error())
			return
		}
	}
	envPath := path.Join(rootDir, "env", "common.env")
	log.Printf("Caminho do arquivo env: %s", envPath)
	err := godotenv.Load(envPath)
	if err != nil {
		log.Println(err.Error())
		return
	}

}

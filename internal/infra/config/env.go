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
	if env == "test" {
		err := godotenv.Load(".env.test")
		if err != nil {
			log.Println(err.Error())
			return
		}
	} else {
		envPath := path.Join(rootDir, "env", "local.env")
		err := godotenv.Load(envPath)
		if err != nil {
			log.Println(err.Error())
			return
		}
	}
	envPath := path.Join(rootDir, "env", "common.env")
	err := godotenv.Load(envPath)
	if err != nil {
		log.Println(err.Error())
		return
	}

}

package main

import (
	"log"
	"os"
	"path"
	"tz-go-readme/mdblock"
	_ "tz-go-readme/parsers"

	"github.com/joho/godotenv"
)

const dataFile = "../data.json"

// load .env file
func init() {
	log.SetOutput(os.Stdout)
	godotenv.Load(".env")
}

func main() {
	filename := os.Getenv("MDFILE")
	if filename == "" {
		filename = "README.md"
	}
	cwd,err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	// invoke from root path
	dataFile := path.Join(cwd,"./data.json")
	mdblock.Run("README.md",dataFile)
}

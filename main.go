package main

import (
	"log"
	"os"
	"tz-go-readme/mdblock"
	_ "tz-go-readme/parsers"

	"github.com/joho/godotenv"
)

const dataFile = "../data.json"

// load .env file
func init(){
    log.SetOutput(os.Stdout)
    godotenv.Load(".env")
}

func main()  {
    filename := os.Getenv("MDFILE")
    if filename == "" {
        filename = "README.md"
    }
    mdblock.Run("README.md")
}
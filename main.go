package main

import (
	"fmt"
	"os"

	"github.com/owenso/crypto-portfolio-api/config"
)

func main() {
	a := App{}

	if os.Getenv("ENV") == "" {
		os.Setenv("ENV", "dev")
		fmt.Println("No Environment found. Defaulting to 'dev'")
	}

	configFile, _ := config.LoadConfiguration()
	a.Initialize(configFile.Database.URI)
	a.ConfigureRouting()

	a.Run(":3001")
}

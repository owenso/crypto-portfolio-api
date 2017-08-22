package main

import (
	"fmt"
	"os"
)

func main() {
	a := App{}

	if os.Getenv("ENV") == "" {
		os.Setenv("ENV", "dev")
		fmt.Println("No Environment found. Defaulting to 'dev'")
	}

	config, _ := LoadConfiguration("./config/config." + os.Getenv("ENV") + ".json")
	a.Initialize(config.Database.URI)

	a.Run(":3001")
}

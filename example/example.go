package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Isotere/dotenv"
	"github.com/pkg/errors"
)

func main() {
	err := dotenv.Load()
	if err != nil {
		log.Fatal(errors.Wrap(err, "cannot load enviroment configuration"))
	}

	fmt.Println("String Value: ", dotenv.GetString("ENV_STRING", "dont show"))
	fmt.Println("String Value: ", dotenv.GetString("ENV_STRING_NOT_EXISTS", "default value"))

	fmt.Println("Int Value: ", dotenv.GetString("ENV_INT", "dont show"))
	fmt.Println("Int Value: ", dotenv.GetString("ENV_INT_NOT_EXISTS", "default value int"))

	os.Exit(0)
}

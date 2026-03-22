package main

import (
	"fmt"
	"os"
	dotenv "github.com/joho/godotenv" // import a module with alias
)

func main () {
	programName, questions := os.Args[0], os.Args[1:]
	fmt.Printf("Starting: %s\n", programName)

	err := dotenv.Load()
	if err != nil {
		fmt.Printf("Error Loading the Environment Variable.")
	}

	s3_BUCKET := os.Getenv("S3_BUCKET")
	secretKey := os.Getenv("SECRET_KEY")
	fmt.Printf("Env variable %s %s\n", s3_BUCKET, secretKey)


	if len(questions) == 0 {
		fmt.Printf("Usage:%s <question1> <question2> ...", programName)
		os.Exit(1)
	}else {
		for i , question := range questions {
			fmt.Printf("Questions [%d] > %s\n", i, question)
		}
	}
}
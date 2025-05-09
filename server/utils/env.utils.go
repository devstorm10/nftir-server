/*
	@Author: Logan (Nam) Nguyen
	@Course: SUNY Oswego - CSC 482
	@Instructor: Professor James Early
	@Purpose: env.utils.go provides function that helps load environemnt variables
*/

/* @package initializers */
package utils

// import packages
import (
	"log"

	"github.com/joho/godotenv"
)

/*
@func: LoadEnvVars() - loads environment varables
*/
func LoadEnvVars()  {
	if err := godotenv.Load();err != nil {
    	log.Fatal("Error loading .env file")
  	}
}
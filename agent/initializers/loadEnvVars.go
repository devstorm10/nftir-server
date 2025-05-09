/*
	@Author: Logan (Nam) Nguyen
	@Course: SUNY Oswego - CSC 482
	@Instructor: Professor James Early
	@Purpose: loadEnvVars.go provide function that helps load environemnt variables
*/

/* @package initializers */
package initializers

// import packages
import (
	"log"

	"github.com/joho/godotenv"
)

/*
	@Function LoadEnvVars() - loads environment varables
*/
func LoadEnvVars()  {
	err := godotenv.Load()
  	if err != nil {
    	log.Fatal("Error loading .env file")
  	}
}
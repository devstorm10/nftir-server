/*
	@Author: Logan (Nam) Nguyen
	@Course: SUNY Oswego - CSC 482
	@Instructor: Professor James Early
*/

// @package: main
package main

// Import packages
import (
	"NFTir/agent/initializers"
	"NFTir/agent/services"

	"github.com/jamespearly/loggly"
)

// initialize global variables
var (
	logglyClient *loggly.ClientType
)

/* @function init() - run before main() */
func init()  {
	initializers.LoadEnvVars();
	logglyClient = loggly.New("NFTir")
}


/* @function main() - root function */
func main()  {
	services.PeriodicallyFetchData(logglyClient)
}
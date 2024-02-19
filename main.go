package main

import (
	"fmt"
	"os"
	"scaleX/assignment/cronjobs"
	"scaleX/assignment/initializer"
	"scaleX/assignment/routers"
)

func main() {

	// This function does the initial setups. for now it is just setting up the databases. but we can use for more
	// such as connecting to redis or initializing a logger.
	initialSetups := initializer.Initialize()

	// This function setups and deployes the cron jobs.
	cronjobs.RunCronJobs(initialSetups.DB)

	// gettint the port from .env file
	port := os.Getenv("PORT")

	// This function sets up the required routers for this application
	r := routers.SetupRouter(initialSetups.DB)

	// Finally we are running the application to accept http requests
	r.Run(fmt.Sprintf(":%s", port))
}

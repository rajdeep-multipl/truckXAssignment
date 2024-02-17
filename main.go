package main

import (
	"fmt"
	"os"
	"scaleX/assignment/cronjobs"
	"scaleX/assignment/initializer"
	"scaleX/assignment/routers"
)

func main() {
	initialSetups := initializer.Initialize()

	port := os.Getenv("PORT")

	cronjobs.RunCronJobs(initialSetups.DB)

	r := routers.SetupRouter(initialSetups.DB)
	r.Run(fmt.Sprintf(":%s", port))
}

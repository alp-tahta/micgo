package main

import "fmt"

type ServiceConfig struct {
	RepoHost int
}

func main() {
	serviceConfig := ServiceConfig{}

	fmt.Println("Where is your microservice hosted?\n" +
		"1. Remote repository (GitHub, GitLab, etc.)\n" +
		"2. Local machine\n" +
		"Enter choice (1/2):")
	fmt.Scanln(&serviceConfig.RepoHost)
	fmt.Println(serviceConfig)
}

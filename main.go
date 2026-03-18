package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {

	container := strings.Join(os.Args[1:], " ")

	stopCmd := exec.Command("docker", "stop", container)
	pulledCmd := exec.Command("git pull")
	pulledCmd.Stdout = os.Stdout
	pulledCmd.Stderr = os.Stderr
	deployCmd := exec.Command("docker", "compose", "up", "-d", "--build")
	deployCmd.Stdout = os.Stdout
	deployCmd.Stderr = os.Stderr

	fmt.Println("Stopping the container...")

	err := stopCmd.Run()

	if err != nil {
		fmt.Println("Error stopping the container")
		return
	}

	fmt.Println("Updating the repositories...")

	err = pulledCmd.Run()

	if err != nil {
		fmt.Println("Error updating the project")
		return
	}

	fmt.Println("Deploying the container...")

	err = deployCmd.Run()

	if err != nil {
		fmt.Println("Error deploying the container")
		return
	}

	fmt.Println("Update completed")
}

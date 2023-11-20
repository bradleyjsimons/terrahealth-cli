// Package main provides the entry point to the terrahealth-cli application.
// It handles command-line arguments and executes the appropriate command.
package main

import (
	"fmt"
	"os"

	"github.com/bradleyjsimons/terrahealth-cli/pkg/aws"
)

// App encapsulates the application's dependencies.
// It uses the EC2Service interface to interact with EC2 instances,
// allowing for easy testing with mock implementations.
type App struct {
	ec2Service aws.EC2Service
}

// main is the entry point of the application.
// It calls the run function with command-line arguments and handles any returned error.
func main() {
	app := &App{
		ec2Service: &aws.EC2Handler{},
	}
	if err := app.run(os.Args); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}


// run executes the application logic based on the provided command-line arguments.
// It returns an error if the arguments are invalid or if there is a problem executing the command.
func (a *App) run(args []string) error {
	// Print the application name and arguments
	fmt.Println("TerraHealth", args)

	// Check that at least one argument was provided
	if len(args) < 2 {
		return fmt.Errorf("Usage: terrahealth <command>")
	}

	// Get the command from the arguments
	command := args[1]

	// Execute the command
	switch command {
	case "check-aws":
		// Check EC2 instances
		a.ec2Service.CheckEC2Instances()
	default:
		// Return an error for unknown commands
		return fmt.Errorf("Unknown command: %v", command)
	}

	return nil
}
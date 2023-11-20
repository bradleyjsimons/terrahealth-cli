// Package main provides the entry point to the terrahealth-cli application.
// It handles command-line arguments and executes the appropriate command.
package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch/cloudwatchiface"
	"github.com/bradleyjsimons/terrahealth-cli/pkg/aws"
)

// App encapsulates the application's dependencies.
// It uses the EC2Service interface to interact with EC2 instances,
// and the CloudWatchAPI interface to interact with AWS CloudWatch.
type App struct {
	ec2Service        aws.EC2Service
	cloudWatchService cloudwatchiface.CloudWatchAPI
}

// main is the entry point of the application.
// It initializes the AWS session, sets up the application dependencies,
// and calls the run function with command-line arguments, handling any returned error.
func main() {
	// Initialize a new AWS session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Set up the application dependencies
	app := &App{
		ec2Service:        &aws.EC2Handler{},
		cloudWatchService: aws.NewCloudWatchClient(sess),
	}

	// Run the application with the command-line arguments
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
	case "getInstances":
		// Create a real EC2 service client
		ec2Svc, err := a.ec2Service.NewAWSSession()
		if err != nil {
			return fmt.Errorf("Error creating AWS session: %v", err)
		}

		// Check EC2 instances
		a.ec2Service.CheckEC2Instances(ec2Svc)
	case "fetchCpuUtilization":
		// Check that the instance ID was provided
		if len(args) != 3 {
			return fmt.Errorf("Usage: fetchCpuUtilization <instanceId>")
		}

		// Fetch the CPU utilization metric for the specified instance
		instanceId := args[2]
		metricDataOutput, err := aws.FetchCpuUtilizationMetric(instanceId, a.cloudWatchService)
		if err != nil {
			return fmt.Errorf("Error fetching CPU utilization: %v", err)
		}

		// Print the metric data
		fmt.Println(metricDataOutput)
	default:
		// Return an error for unknown commands
		return fmt.Errorf("Unknown command: %v", command)
	}

	return nil
}
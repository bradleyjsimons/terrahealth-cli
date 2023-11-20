// Package main is the entry point for the TerraHealth CLI application.
// TerraHealth CLI provides functionalities to interact with AWS services
// and check the health of various resources.
package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

// main is the entry point of the TerraHealth CLI application.
// It processes command-line arguments and invokes appropriate functions
// based on the provided commands.
func main() {
	fmt.Println("TerraHealth", os.Args)
	if len(os.Args) < 2 {
		fmt.Println("Usage: terrahealth <command>")
		os.Exit(1)
	}

	command := os.Args[1]
	switch command {
	case "check-aws":
		// Initialize a new AWS session
		sess, err := session.NewSession(&aws.Config{
			Region: aws.String("us-west-2"), 
		})
		if err != nil {
			fmt.Println("Error creating AWS session:", err)
			return
		}

		// Create an EC2 service client
		ec2Svc := ec2.New(sess)

		checkAWSResources(ec2Svc)
	default:
		fmt.Println("Unknown command:", command)
	}
}

// checkAWSResources connects to AWS and performs a health check of AWS resources.
// Currently, it lists the IDs of all active EC2 instances.
// It takes an EC2 service client as a parameter and queries EC2 instances, outputting their IDs.
func checkAWSResources(ec2Svc ec2iface.EC2API) {
	// Describe EC2 instances and handle potential errors
	result, err := ec2Svc.DescribeInstances(nil)
	if err != nil {
		fmt.Println("Error describing EC2 instances:", err)
		return
	}

	// Iterate over the reservations and instances to print instance IDs
	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			fmt.Println("Instance ID:", *instance.InstanceId)
		}
	}
}
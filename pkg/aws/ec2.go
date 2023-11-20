// Package aws contains functionalities related to AWS services.
package aws

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

// EC2Service is an interface that describes the methods needed to interact with EC2 instances.
type EC2Service interface {
	CheckEC2Instances()
}

// EC2Handler is a type that implements the EC2Service interface,
// providing methods to interact with EC2 instances.
type EC2Handler struct{}


// NewAWSSession creates a new AWS session and returns an EC2 service.
// It returns an EC2 service as an ec2iface.EC2API interface, which can be used to interact with EC2.
// If there is a problem creating the AWS session, it returns an error.
func NewAWSSession() (ec2iface.EC2API, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})
	if err != nil {
		return nil, fmt.Errorf("Error creating AWS session: %v", err)
	}

	return ec2.New(sess), nil
}

// CheckEC2Instances connects to AWS, performs a health check of EC2 instances,
// and lists the IDs of all active EC2 instances.
// It takes an EC2 service client as a parameter and queries EC2 instances, outputting their IDs.
func (e *EC2Handler) CheckEC2Instances() {
	ec2Svc, err := NewAWSSession()
	if err != nil {
		fmt.Println("Error creating AWS session:", err)
		return
	}

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
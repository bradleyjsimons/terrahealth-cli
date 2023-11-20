// Package main_test contains tests for the main package.
package main

import (
	"bytes"
	"errors"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

// mockEC2SvcSuccess is a mock EC2 service that successfully returns a single instance.
type mockEC2SvcSuccess struct {
	ec2iface.EC2API
}

// DescribeInstances returns a successful response with a single instance.
func (m *mockEC2SvcSuccess) DescribeInstances(*ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
	return &ec2.DescribeInstancesOutput{
		Reservations: []*ec2.Reservation{
			{
				Instances: []*ec2.Instance{
					{
						InstanceId: aws.String("i-1234567890abcdef0"),
					},
				},
			},
		},
	}, nil
}

// mockEC2SvcFailure is a mock EC2 service that returns an error when describing instances.
type mockEC2SvcFailure struct {
	ec2iface.EC2API
}

// DescribeInstances returns an error.
func (m *mockEC2SvcFailure) DescribeInstances(*ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
	return nil, errors.New("error describing EC2 instances")
}

// TestCheckAWSResources_Success tests that checkAWSResources correctly outputs the ID of a single instance.
func TestCheckAWSResources_Success(t *testing.T) {
	t.Parallel()

	// Create a mock EC2 service with a single instance
	mockSvc := &mockEC2SvcSuccess{}

	// Capture standard output
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call the function under test
	checkAWSResources(mockSvc)

	// Stop capturing standard output
	w.Close()
	os.Stdout = oldStdout

	// Read the captured standard output
	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	// Check that the output includes the expected instance ID
	expectedOutput := "Instance ID: i-1234567890abcdef0\n"
	if output != expectedOutput {
		t.Errorf("Expected output %q, but got %q", expectedOutput, output)
	}
}

// TestCheckAWSResources_Failure tests that checkAWSResources correctly handles an error from the EC2 service.
func TestCheckAWSResources_Failure(t *testing.T) {
	t.Parallel()

	// Create a mock EC2 service that returns an error
	mockSvc := &mockEC2SvcFailure{}

	// Capture standard output
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call the function under test
	checkAWSResources(mockSvc)

	// Stop capturing standard output
	w.Close()
	os.Stdout = oldStdout

	// Read the captured standard output
	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	// Check that the output includes the expected error message
	expectedOutput := "Error describing EC2 instances: error describing EC2 instances\n"
	if output != expectedOutput {
		t.Errorf("Expected output %q, but got %q", expectedOutput, output)
	}
}
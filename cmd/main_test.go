package main

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

// mockEC2Service is a mock implementation of the aws.EC2Service interface.
type mockEC2Service struct {
	checkEC2InstancesCalled bool
	shouldReturnError       bool
}

// NewAWSSession is a mock implementation of aws.EC2Service.NewAWSSession.
func (m *mockEC2Service) NewAWSSession() (ec2iface.EC2API, error) {
	if m.shouldReturnError {
		return nil, fmt.Errorf("mock error")
	}
	return nil, nil
}
// CheckEC2Instances is a mock implementation of aws.EC2Service.CheckEC2Instances.
func (m *mockEC2Service) CheckEC2Instances(ec2Svc ec2iface.EC2API) {
	m.checkEC2InstancesCalled = true
}

func TestRun(t *testing.T) {
	// Create a mock EC2Service.
	mockService := &mockEC2Service{}

	// Create an App with the mock EC2Service.
	app := &App{
		ec2Service: mockService,
	}

	// Test the "check-aws" command.
	err := app.run([]string{"", "check-aws"})
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	if !mockService.checkEC2InstancesCalled {
		t.Errorf("Expected CheckEC2Instances to be called")
	}

	// Test an unknown command.
	err = app.run([]string{"", "unknown-command"})
	if err == nil {
		t.Errorf("Expected error, got none")
	}
	expectedError := "Unknown command: unknown-command"
	if err.Error() != expectedError {
		t.Errorf("Expected error: %v, got: %v", expectedError, err)
	}

	// Test with less than 2 arguments.
	err = app.run([]string{""})
	if err == nil {
		t.Errorf("Expected error, got none")
	}
	expectedError = "Usage: terrahealth <command>"
	if err.Error() != expectedError {
		t.Errorf("Expected error: %v, got: %v", expectedError, err)
	}

	// Test the error case when creating an AWS session.
	mockService.shouldReturnError = true
	err = app.run([]string{"", "check-aws"})
	if err == nil {
		t.Errorf("Expected error, got none")
	}
	expectedError = "Error creating AWS session: mock error"
	if err.Error() != expectedError {
		t.Errorf("Expected error: %v, got: %v", expectedError, err)
	}
}
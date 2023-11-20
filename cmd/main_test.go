// Package main provides the entry point to the terrahealth-cli application.
// This file contains unit tests for the main package.
package main

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/cloudwatch/cloudwatchiface"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

// mockEC2Service is a mock implementation of the aws.EC2Service interface.
// It is used to test the functions that interact with EC2 instances.
type mockEC2Service struct {
	checkEC2InstancesCalled bool
	shouldReturnError       bool
}

// NewAWSSession is a mock implementation of aws.EC2Service.NewAWSSession.
// It returns a mock AWS session or an error based on the shouldReturnError field.
func (m *mockEC2Service) NewAWSSession() (ec2iface.EC2API, error) {
	if m.shouldReturnError {
		return nil, fmt.Errorf("mock error")
	}
	return nil, nil
}

// CheckEC2Instances is a mock implementation of aws.EC2Service.CheckEC2Instances.
// It sets the checkEC2InstancesCalled field to true when called.
func (m *mockEC2Service) CheckEC2Instances(ec2Svc ec2iface.EC2API) {
	m.checkEC2InstancesCalled = true
}

// mockCloudWatchService is a mock implementation of the cloudwatchiface.CloudWatchAPI interface.
// It is used to test the functions that interact with the AWS CloudWatch service.
type mockCloudWatchService struct {
	cloudwatchiface.CloudWatchAPI
	fetchCpuUtilizationCalled bool
}

// GetMetricData is a mock implementation of cloudwatchiface.CloudWatchAPI.GetMetricData.
// It sets the fetchCpuUtilizationCalled field to true when called and returns a mock response.
func (m *mockCloudWatchService) GetMetricData(input *cloudwatch.GetMetricDataInput) (*cloudwatch.GetMetricDataOutput, error) {
	m.fetchCpuUtilizationCalled = true
	return &cloudwatch.GetMetricDataOutput{}, nil
}

// TestRun tests the run function with various command-line arguments.
// It checks that the correct functions are called based on the arguments and that errors are returned when expected.
func TestRun(t *testing.T) {
	// Create a mock EC2Service and a mock CloudWatchService.
	mockEC2Service := &mockEC2Service{}
	mockCloudWatchService := &mockCloudWatchService{}

	// Create an App with the mock services.
	app := &App{
		ec2Service:        mockEC2Service,
		cloudWatchService: mockCloudWatchService,
	}

	// Test the "getInstances" command.
	err := app.run([]string{"", "getInstances"})
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	if !mockEC2Service.checkEC2InstancesCalled {
		t.Errorf("Expected CheckEC2Instances to be called")
	}

	// Test the "fetchCpuUtilization" command.
	err = app.run([]string{"", "fetchCpuUtilization", "test-instance"})
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	if !mockCloudWatchService.fetchCpuUtilizationCalled {
		t.Errorf("Expected FetchCpuUtilizationMetric to be called")
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
	mockEC2Service.shouldReturnError = true
	err = app.run([]string{"", "getInstances"})
	if err == nil {
		t.Errorf("Expected error, got none")
	}
	expectedError = "Error creating AWS session: mock error"
	if err.Error() != expectedError {
		t.Errorf("Expected error: %v, got: %v", expectedError, err)
	}
}
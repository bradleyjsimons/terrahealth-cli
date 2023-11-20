package aws

import (
	"bytes"
	"errors"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

// MockEC2Service is a mock implementation of the EC2Service interface.
type MockEC2Service struct {
	ec2iface.EC2API
	Resp             ec2.DescribeInstancesOutput
	shouldReturnError bool
}

// DescribeInstances is a mock implementation of the DescribeInstances method.
func (m *MockEC2Service) DescribeInstances(input *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
	if m.shouldReturnError {
		return nil, errors.New("mock error")
	}
	return &m.Resp, nil
}

func (m *MockEC2Service) NewAWSSession() (ec2iface.EC2API, error) {
	if m.shouldReturnError {
		return nil, errors.New("mock error")
	}
	return m, nil
}

func TestCheckEC2Instances(t *testing.T) {
	// Create a mock EC2 service with a predefined response
	mockSvc := &MockEC2Service{
		Resp: ec2.DescribeInstancesOutput{
			Reservations: []*ec2.Reservation{
				{
					Instances: []*ec2.Instance{
						{InstanceId: aws.String("i-1234567890")},
					},
				},
			},
		},
	}

	// Create an EC2Handler
	handler := &EC2Handler{}

	// Redirect standard output to a buffer
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call CheckEC2Instances
	handler.CheckEC2Instances(mockSvc)

	// Capture standard output
	w.Close()
	os.Stdout = oldStdout // Restore standard output
	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	// Check the output
	expectedOutput := "Instance ID: i-1234567890\n"
	if output != expectedOutput {
		t.Errorf("Expected output %q, but got %q", expectedOutput, output)
	}
}

func TestNewAWSSession(t *testing.T) {
	handler := &EC2Handler{}

	// Call NewAWSSession
	ec2Svc, err := handler.NewAWSSession()

	// Check that it returns a non-nil EC2API instance and no error
	if ec2Svc == nil {
		t.Errorf("Expected non-nil EC2API instance, but got nil")
	}
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}

func TestCheckEC2Instances_Error(t *testing.T) {
	// Create a mock EC2 service that returns an error
	mockSvc := &MockEC2Service{
		shouldReturnError: true,
	}

	// Create an EC2Handler
	handler := &EC2Handler{}

	// Redirect standard output to a buffer
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call CheckEC2Instances
	handler.CheckEC2Instances(mockSvc)

	// Capture standard output
	w.Close()
	os.Stdout = oldStdout // Restore standard output
	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	// Check the output
	expectedOutput := "Error describing EC2 instances: mock error\n"
	if output != expectedOutput {
		t.Errorf("Expected output %q, but got %q", expectedOutput, output)
	}
}

// func TestNewAWSSession_Error(t *testing.T) {
// 	// Create an EC2Handler
// 	handler := &EC2Handler{}

// 	// Call NewAWSSession
// 	_, err := handler.NewAWSSession()

// 	// Check that it returns an error
// 	if err == nil {
// 		t.Errorf("Expected error, but got nil")
// 	}
// }
// Package aws_test provides tests for the AWS functionality in the aws package.
// It includes tests for initializing a CloudWatch client and fetching CPU utilization metrics for EC2 instances.
package aws_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/cloudwatch/cloudwatchiface"
	"github.com/bradleyjsimons/terrahealth-cli/pkg/aws"
	"github.com/stretchr/testify/assert"
)

// mockCloudWatchClient is a mock implementation of the CloudWatchAPI interface.
// It is used to test the functions that interact with the AWS CloudWatch service.
type mockCloudWatchClient struct {
	cloudwatchiface.CloudWatchAPI
}

// GetMetricData is a mock implementation of the GetMetricData method of the CloudWatchAPI interface.
// It returns a mock response that can be used for testing.
func (m *mockCloudWatchClient) GetMetricData(input *cloudwatch.GetMetricDataInput) (*cloudwatch.GetMetricDataOutput, error) {
	// Mock response
	return &cloudwatch.GetMetricDataOutput{}, nil
}

// TestFetchCpuUtilizationMetric tests the FetchCpuUtilizationMetric function.
// It creates a mock CloudWatch client and calls the function with this client and a test instance ID.
// It asserts that no error is returned.
func TestFetchCpuUtilizationMetric(t *testing.T) {
	// Create a mock CloudWatch client
	mockSvc := &mockCloudWatchClient{}

	// Call the function with the mock client
	_, err := aws.FetchCpuUtilizationMetric("test-instance", mockSvc)

	// Assert no error was returned
	assert.NoError(t, err)
}

// TestNewCloudWatchClient tests the NewCloudWatchClient function.
// It initializes a new AWS session and calls the function with this session.
// It asserts that the returned client is not nil.
func TestNewCloudWatchClient(t *testing.T) {
	// Initialize a new AWS session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Call the function
	client := aws.NewCloudWatchClient(sess)

	// Assert the client was created
	assert.NotNil(t, client)
}
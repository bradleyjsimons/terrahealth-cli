// Package aws provides functionality to interact with AWS services such as EC2 and CloudWatch.
// It includes functions to initialize a CloudWatch client and fetch CPU utilization metrics for EC2 instances.
package aws

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/cloudwatch/cloudwatchiface"
)

// NewCloudWatchClient initializes a new CloudWatch client using the provided AWS session.
// It returns an interface to the CloudWatch API, which can be used to interact with the service.
func NewCloudWatchClient(sess *session.Session) cloudwatchiface.CloudWatchAPI {
    return cloudwatch.New(sess)
}

// FetchCpuUtilizationMetric fetches the CPU utilization metric for the specified EC2 instance.
// It uses the provided CloudWatch client to fetch the metric data from the last 24 hours, with a period of 1 hour.
// The function returns the metric data output and any error encountered.
func FetchCpuUtilizationMetric(instanceId string, cwClient cloudwatchiface.CloudWatchAPI) (*cloudwatch.GetMetricDataOutput, error) {
    // Construct the input for the metric query
    metricInput := &cloudwatch.GetMetricDataInput{
        StartTime: aws.Time(time.Now().Add(-time.Hour * 24)), // start time 24 hours ago
        EndTime:   aws.Time(time.Now()),                      // end time now
        MetricDataQueries: []*cloudwatch.MetricDataQuery{
            {
                Id: aws.String("cpuUtilization"),
                MetricStat: &cloudwatch.MetricStat{
                    Metric: &cloudwatch.Metric{
                        Namespace:  aws.String("AWS/EC2"),
                        MetricName: aws.String("CPUUtilization"),
                        Dimensions: []*cloudwatch.Dimension{
                            {
                                Name:  aws.String("InstanceId"),
                                Value: aws.String(instanceId),
                            },
                        },
                    },
                    Period: aws.Int64(3600),             // period in seconds (3600s = 1 hour)
                    Stat:   aws.String("Average"),       // statistic to return
                    Unit:   aws.String("Percent"),       // unit
                },
                ReturnData: aws.Bool(true),
            },
        },
    }

    // Fetch the metric data
    return cwClient.GetMetricData(metricInput)
}
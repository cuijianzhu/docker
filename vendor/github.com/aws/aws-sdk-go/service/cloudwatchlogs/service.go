// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package cloudwatchlogs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/aws/aws-sdk-go/private/protocol/jsonrpc"
)

// You can use Amazon CloudWatch Logs to monitor, store, and access your log
// files from Amazon Elastic Compute Cloud (Amazon EC2) instances, Amazon CloudTrail,
// or other sources. You can then retrieve the associated log data from CloudWatch
// Logs using the Amazon CloudWatch console, the CloudWatch Logs commands in
// the AWS CLI, the CloudWatch Logs API, or the CloudWatch Logs SDK.
//
// You can use CloudWatch Logs to:
//
//    * Monitor Logs from Amazon EC2 Instances in Real-time: You can use CloudWatch
//    Logs to monitor applications and systems using log data. For example,
//    CloudWatch Logs can track the number of errors that occur in your application
//    logs and send you a notification whenever the rate of errors exceeds a
//    threshold you specify. CloudWatch Logs uses your log data for monitoring;
//    so, no code changes are required. For example, you can monitor application
//    logs for specific literal terms (such as "NullReferenceException") or
//    count the number of occurrences of a literal term at a particular position
//    in log data (such as "404" status codes in an Apache access log). When
//    the term you are searching for is found, CloudWatch Logs reports the data
//    to a Amazon CloudWatch metric that you specify.
//
//    * Monitor Amazon CloudTrail Logged Events: You can create alarms in Amazon
//    CloudWatch and receive notifications of particular API activity as captured
//    by CloudTrail and use the notification to perform troubleshooting.
//
//    * Archive Log Data: You can use CloudWatch Logs to store your log data
//    in highly durable storage. You can change the log retention setting so
//    that any log events older than this setting are automatically deleted.
//    The CloudWatch Logs agent makes it easy to quickly send both rotated and
//    non-rotated log data off of a host and into the log service. You can then
//    access the raw log data when you need it.
//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type CloudWatchLogs struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// A ServiceName is the name of the service the client will make API calls to.
const ServiceName = "logs"

// New creates a new instance of the CloudWatchLogs client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a CloudWatchLogs client from just a session.
//     svc := cloudwatchlogs.New(mySession)
//
//     // Create a CloudWatchLogs client with additional configuration
//     svc := cloudwatchlogs.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *CloudWatchLogs {
	c := p.ClientConfig(ServiceName, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion string) *CloudWatchLogs {
	svc := &CloudWatchLogs{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2014-03-28",
				JSONVersion:   "1.1",
				TargetPrefix:  "Logs_20140328",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(jsonrpc.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a CloudWatchLogs operation and runs any
// custom request initialization.
func (c *CloudWatchLogs) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
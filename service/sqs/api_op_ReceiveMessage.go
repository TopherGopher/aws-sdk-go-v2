// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sqs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/ReceiveMessageRequest
type ReceiveMessageInput struct {
	_ struct{} `type:"structure"`

	// A list of attributes that need to be returned along with each message. These
	// attributes include:
	//
	//    * All - Returns all values.
	//
	//    * ApproximateFirstReceiveTimestamp - Returns the time the message was
	//    first received from the queue (epoch time (http://en.wikipedia.org/wiki/Unix_time)
	//    in milliseconds).
	//
	//    * ApproximateReceiveCount - Returns the number of times a message has
	//    been received from the queue but not deleted.
	//
	//    * SenderId For an IAM user, returns the IAM user ID, for example ABCDEFGHI1JKLMNOPQ23R.
	//    For an IAM role, returns the IAM role ID, for example ABCDE1F2GH3I4JK5LMNOP:i-a123b456.
	//
	//    * SentTimestamp - Returns the time the message was sent to the queue (epoch
	//    time (http://en.wikipedia.org/wiki/Unix_time) in milliseconds).
	//
	//    * MessageDeduplicationId - Returns the value provided by the producer
	//    that calls the SendMessage action.
	//
	//    * MessageGroupId - Returns the value provided by the producer that calls
	//    the SendMessage action. Messages with the same MessageGroupId are returned
	//    in sequence.
	//
	//    * SequenceNumber - Returns the value provided by Amazon SQS.
	AttributeNames []QueueAttributeName `locationNameList:"AttributeName" type:"list" flattened:"true"`

	// The maximum number of messages to return. Amazon SQS never returns more messages
	// than this value (however, fewer messages might be returned). Valid values:
	// 1 to 10. Default: 1.
	MaxNumberOfMessages *int64 `type:"integer"`

	// The name of the message attribute, where N is the index.
	//
	//    * The name can contain alphanumeric characters and the underscore (_),
	//    hyphen (-), and period (.).
	//
	//    * The name is case-sensitive and must be unique among all attribute names
	//    for the message.
	//
	//    * The name must not start with AWS-reserved prefixes such as AWS. or Amazon.
	//    (or any casing variants).
	//
	//    * The name must not start or end with a period (.), and it should not
	//    have periods in succession (..).
	//
	//    * The name can be up to 256 characters long.
	//
	// When using ReceiveMessage, you can send a list of attribute names to receive,
	// or you can return all of the attributes by specifying All or .* in your request.
	// You can also use all message attributes starting with a prefix, for example
	// bar.*.
	MessageAttributeNames []string `locationNameList:"MessageAttributeName" type:"list" flattened:"true"`

	// The URL of the Amazon SQS queue from which messages are received.
	//
	// Queue URLs and names are case-sensitive.
	//
	// QueueUrl is a required field
	QueueUrl *string `type:"string" required:"true"`

	// This parameter applies only to FIFO (first-in-first-out) queues.
	//
	// The token used for deduplication of ReceiveMessage calls. If a networking
	// issue occurs after a ReceiveMessage action, and instead of a response you
	// receive a generic error, you can retry the same action with an identical
	// ReceiveRequestAttemptId to retrieve the same set of messages, even if their
	// visibility timeout has not yet expired.
	//
	//    * You can use ReceiveRequestAttemptId only for 5 minutes after a ReceiveMessage
	//    action.
	//
	//    * When you set FifoQueue, a caller of the ReceiveMessage action can provide
	//    a ReceiveRequestAttemptId explicitly.
	//
	//    * If a caller of the ReceiveMessage action doesn't provide a ReceiveRequestAttemptId,
	//    Amazon SQS generates a ReceiveRequestAttemptId.
	//
	//    * You can retry the ReceiveMessage action with the same ReceiveRequestAttemptId
	//    if none of the messages have been modified (deleted or had their visibility
	//    changes).
	//
	//    * During a visibility timeout, subsequent calls with the same ReceiveRequestAttemptId
	//    return the same messages and receipt handles. If a retry occurs within
	//    the deduplication interval, it resets the visibility timeout. For more
	//    information, see Visibility Timeout (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-visibility-timeout.html)
	//    in the Amazon Simple Queue Service Developer Guide. If a caller of the
	//    ReceiveMessage action still processes messages when the visibility timeout
	//    expires and messages become visible, another worker consuming from the
	//    same queue can receive the same messages and therefore process duplicates.
	//    Also, if a consumer whose message processing time is longer than the visibility
	//    timeout tries to delete the processed messages, the action fails with
	//    an error. To mitigate this effect, ensure that your application observes
	//    a safe threshold before the visibility timeout expires and extend the
	//    visibility timeout as necessary.
	//
	//    * While messages with a particular MessageGroupId are invisible, no more
	//    messages belonging to the same MessageGroupId are returned until the visibility
	//    timeout expires. You can still receive messages with another MessageGroupId
	//    as long as it is also visible.
	//
	//    * If a caller of ReceiveMessage can't track the ReceiveRequestAttemptId,
	//    no retries work until the original visibility timeout expires. As a result,
	//    delays might occur but the messages in the queue remain in a strict order.
	//
	// The length of ReceiveRequestAttemptId is 128 characters. ReceiveRequestAttemptId
	// can contain alphanumeric characters (a-z, A-Z, 0-9) and punctuation (!"#$%&'()*+,-./:;<=>?@[\]^_`{|}~).
	//
	// For best practices of using ReceiveRequestAttemptId, see Using the ReceiveRequestAttemptId
	// Request Parameter (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/using-receiverequestattemptid-request-parameter.html)
	// in the Amazon Simple Queue Service Developer Guide.
	ReceiveRequestAttemptId *string `type:"string"`

	// The duration (in seconds) that the received messages are hidden from subsequent
	// retrieve requests after being retrieved by a ReceiveMessage request.
	VisibilityTimeout *int64 `type:"integer"`

	// The duration (in seconds) for which the call waits for a message to arrive
	// in the queue before returning. If a message is available, the call returns
	// sooner than WaitTimeSeconds. If no messages are available and the wait time
	// expires, the call returns successfully with an empty list of messages.
	WaitTimeSeconds *int64 `type:"integer"`
}

// String returns the string representation
func (s ReceiveMessageInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ReceiveMessageInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ReceiveMessageInput"}

	if s.QueueUrl == nil {
		invalidParams.Add(aws.NewErrParamRequired("QueueUrl"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A list of received messages.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/ReceiveMessageResult
type ReceiveMessageOutput struct {
	_ struct{} `type:"structure"`

	// A list of messages.
	Messages []Message `locationNameList:"Message" type:"list" flattened:"true"`
}

// String returns the string representation
func (s ReceiveMessageOutput) String() string {
	return awsutil.Prettify(s)
}

const opReceiveMessage = "ReceiveMessage"

// ReceiveMessageRequest returns a request value for making API operation for
// Amazon Simple Queue Service.
//
// Retrieves one or more messages (up to 10), from the specified queue. Using
// the WaitTimeSeconds parameter enables long-poll support. For more information,
// see Amazon SQS Long Polling (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-long-polling.html)
// in the Amazon Simple Queue Service Developer Guide.
//
// Short poll is the default behavior where a weighted random set of machines
// is sampled on a ReceiveMessage call. Thus, only the messages on the sampled
// machines are returned. If the number of messages in the queue is small (fewer
// than 1,000), you most likely get fewer messages than you requested per ReceiveMessage
// call. If the number of messages in the queue is extremely small, you might
// not receive any messages in a particular ReceiveMessage response. If this
// happens, repeat the request.
//
// For each message returned, the response includes the following:
//
//    * The message body.
//
//    * An MD5 digest of the message body. For information about MD5, see RFC1321
//    (https://www.ietf.org/rfc/rfc1321.txt).
//
//    * The MessageId you received when you sent the message to the queue.
//
//    * The receipt handle.
//
//    * The message attributes.
//
//    * An MD5 digest of the message attributes.
//
// The receipt handle is the identifier you must provide when deleting the message.
// For more information, see Queue and Message Identifiers (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-queue-message-identifiers.html)
// in the Amazon Simple Queue Service Developer Guide.
//
// You can provide the VisibilityTimeout parameter in your request. The parameter
// is applied to the messages that Amazon SQS returns in the response. If you
// don't include the parameter, the overall visibility timeout for the queue
// is used for the returned messages. For more information, see Visibility Timeout
// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-visibility-timeout.html)
// in the Amazon Simple Queue Service Developer Guide.
//
// A message that isn't deleted or a message whose visibility isn't extended
// before the visibility timeout expires counts as a failed receive. Depending
// on the configuration of the queue, the message might be sent to the dead-letter
// queue.
//
// In the future, new attributes might be added. If you write code that calls
// this action, we recommend that you structure your code so that it can handle
// new attributes gracefully.
//
//    // Example sending a request using ReceiveMessageRequest.
//    req := client.ReceiveMessageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/ReceiveMessage
func (c *Client) ReceiveMessageRequest(input *ReceiveMessageInput) ReceiveMessageRequest {
	op := &aws.Operation{
		Name:       opReceiveMessage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ReceiveMessageInput{}
	}

	req := c.newRequest(op, input, &ReceiveMessageOutput{})
	return ReceiveMessageRequest{Request: req, Input: input, Copy: c.ReceiveMessageRequest}
}

// ReceiveMessageRequest is the request type for the
// ReceiveMessage API operation.
type ReceiveMessageRequest struct {
	*aws.Request
	Input *ReceiveMessageInput
	Copy  func(*ReceiveMessageInput) ReceiveMessageRequest
}

// Send marshals and sends the ReceiveMessage API request.
func (r ReceiveMessageRequest) Send(ctx context.Context) (*ReceiveMessageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ReceiveMessageResponse{
		ReceiveMessageOutput: r.Request.Data.(*ReceiveMessageOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ReceiveMessageResponse is the response type for the
// ReceiveMessage API operation.
type ReceiveMessageResponse struct {
	*ReceiveMessageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ReceiveMessage request.
func (r *ReceiveMessageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

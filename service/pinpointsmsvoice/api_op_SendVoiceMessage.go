// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpointsmsvoice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A request to create and send a new voice message.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-sms-voice-2018-09-05/SendVoiceMessageRequest
type SendVoiceMessageInput struct {
	_ struct{} `type:"structure"`

	// The phone number that appears on recipients' devices when they receive the
	// message.
	CallerId *string `type:"string"`

	// The name of the configuration set that you want to use to send the message.
	ConfigurationSetName *string `type:"string"`

	// An object that contains a voice message and information about the recipient
	// that you want to send it to.
	Content *VoiceMessageContent `type:"structure"`

	// The phone number that you want to send the voice message to.
	DestinationPhoneNumber *string `type:"string"`

	// The phone number that Amazon Pinpoint should use to send the voice message.
	// This isn't necessarily the phone number that appears on recipients' devices
	// when they receive the message, because you can specify a CallerId parameter
	// in the request.
	OriginationPhoneNumber *string `type:"string"`
}

// String returns the string representation
func (s SendVoiceMessageInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SendVoiceMessageInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.CallerId != nil {
		v := *s.CallerId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CallerId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ConfigurationSetName != nil {
		v := *s.ConfigurationSetName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ConfigurationSetName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Content != nil {
		v := s.Content

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Content", v, metadata)
	}
	if s.DestinationPhoneNumber != nil {
		v := *s.DestinationPhoneNumber

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DestinationPhoneNumber", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.OriginationPhoneNumber != nil {
		v := *s.OriginationPhoneNumber

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "OriginationPhoneNumber", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// An object that that contains the Message ID of a Voice message that was sent
// successfully.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-sms-voice-2018-09-05/SendVoiceMessageResponse
type SendVoiceMessageOutput struct {
	_ struct{} `type:"structure"`

	// A unique identifier for the voice message.
	MessageId *string `type:"string"`
}

// String returns the string representation
func (s SendVoiceMessageOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SendVoiceMessageOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.MessageId != nil {
		v := *s.MessageId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MessageId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opSendVoiceMessage = "SendVoiceMessage"

// SendVoiceMessageRequest returns a request value for making API operation for
// Amazon Pinpoint SMS and Voice Service.
//
// Create a new voice message and send it to a recipient's phone number.
//
//    // Example sending a request using SendVoiceMessageRequest.
//    req := client.SendVoiceMessageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-sms-voice-2018-09-05/SendVoiceMessage
func (c *Client) SendVoiceMessageRequest(input *SendVoiceMessageInput) SendVoiceMessageRequest {
	op := &aws.Operation{
		Name:       opSendVoiceMessage,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/sms-voice/voice/message",
	}

	if input == nil {
		input = &SendVoiceMessageInput{}
	}

	req := c.newRequest(op, input, &SendVoiceMessageOutput{})
	return SendVoiceMessageRequest{Request: req, Input: input, Copy: c.SendVoiceMessageRequest}
}

// SendVoiceMessageRequest is the request type for the
// SendVoiceMessage API operation.
type SendVoiceMessageRequest struct {
	*aws.Request
	Input *SendVoiceMessageInput
	Copy  func(*SendVoiceMessageInput) SendVoiceMessageRequest
}

// Send marshals and sends the SendVoiceMessage API request.
func (r SendVoiceMessageRequest) Send(ctx context.Context) (*SendVoiceMessageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SendVoiceMessageResponse{
		SendVoiceMessageOutput: r.Request.Data.(*SendVoiceMessageOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SendVoiceMessageResponse is the response type for the
// SendVoiceMessage API operation.
type SendVoiceMessageResponse struct {
	*SendVoiceMessageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SendVoiceMessage request.
func (r *SendVoiceMessageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

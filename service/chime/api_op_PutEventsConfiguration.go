// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/PutEventsConfigurationRequest
type PutEventsConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime account ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// The bot ID.
	//
	// BotId is a required field
	BotId *string `location:"uri" locationName:"botId" type:"string" required:"true"`

	// Lambda function ARN that allows the bot to receive outgoing events.
	LambdaFunctionArn *string `type:"string"`

	// HTTPS endpoint that allows the bot to receive outgoing events.
	OutboundEventsHTTPSEndpoint *string `type:"string"`
}

// String returns the string representation
func (s PutEventsConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutEventsConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutEventsConfigurationInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.BotId == nil {
		invalidParams.Add(aws.NewErrParamRequired("BotId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutEventsConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.LambdaFunctionArn != nil {
		v := *s.LambdaFunctionArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LambdaFunctionArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.OutboundEventsHTTPSEndpoint != nil {
		v := *s.OutboundEventsHTTPSEndpoint

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "OutboundEventsHTTPSEndpoint", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "accountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BotId != nil {
		v := *s.BotId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "botId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/PutEventsConfigurationResponse
type PutEventsConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// The configuration that allows a bot to receive outgoing events. Can be either
	// an HTTPS endpoint or a Lambda function ARN.
	EventsConfiguration *EventsConfiguration `type:"structure"`
}

// String returns the string representation
func (s PutEventsConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutEventsConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.EventsConfiguration != nil {
		v := s.EventsConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "EventsConfiguration", v, metadata)
	}
	return nil
}

const opPutEventsConfiguration = "PutEventsConfiguration"

// PutEventsConfigurationRequest returns a request value for making API operation for
// Amazon Chime.
//
// Creates an events configuration that allows a bot to receive outgoing events
// sent by Amazon Chime. Choose either an HTTPS endpoint or a Lambda function
// ARN. For more information, see Bot.
//
//    // Example sending a request using PutEventsConfigurationRequest.
//    req := client.PutEventsConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/PutEventsConfiguration
func (c *Client) PutEventsConfigurationRequest(input *PutEventsConfigurationInput) PutEventsConfigurationRequest {
	op := &aws.Operation{
		Name:       opPutEventsConfiguration,
		HTTPMethod: "PUT",
		HTTPPath:   "/accounts/{accountId}/bots/{botId}/events-configuration",
	}

	if input == nil {
		input = &PutEventsConfigurationInput{}
	}

	req := c.newRequest(op, input, &PutEventsConfigurationOutput{})
	return PutEventsConfigurationRequest{Request: req, Input: input, Copy: c.PutEventsConfigurationRequest}
}

// PutEventsConfigurationRequest is the request type for the
// PutEventsConfiguration API operation.
type PutEventsConfigurationRequest struct {
	*aws.Request
	Input *PutEventsConfigurationInput
	Copy  func(*PutEventsConfigurationInput) PutEventsConfigurationRequest
}

// Send marshals and sends the PutEventsConfiguration API request.
func (r PutEventsConfigurationRequest) Send(ctx context.Context) (*PutEventsConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutEventsConfigurationResponse{
		PutEventsConfigurationOutput: r.Request.Data.(*PutEventsConfigurationOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutEventsConfigurationResponse is the response type for the
// PutEventsConfiguration API operation.
type PutEventsConfigurationResponse struct {
	*PutEventsConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutEventsConfiguration request.
func (r *PutEventsConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/GetSubscriptionDefinitionVersionRequest
type GetSubscriptionDefinitionVersionInput struct {
	_ struct{} `type:"structure"`

	NextToken *string `location:"querystring" locationName:"NextToken" type:"string"`

	// SubscriptionDefinitionId is a required field
	SubscriptionDefinitionId *string `location:"uri" locationName:"SubscriptionDefinitionId" type:"string" required:"true"`

	// SubscriptionDefinitionVersionId is a required field
	SubscriptionDefinitionVersionId *string `location:"uri" locationName:"SubscriptionDefinitionVersionId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetSubscriptionDefinitionVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetSubscriptionDefinitionVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetSubscriptionDefinitionVersionInput"}

	if s.SubscriptionDefinitionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SubscriptionDefinitionId"))
	}

	if s.SubscriptionDefinitionVersionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SubscriptionDefinitionVersionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetSubscriptionDefinitionVersionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.SubscriptionDefinitionId != nil {
		v := *s.SubscriptionDefinitionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "SubscriptionDefinitionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SubscriptionDefinitionVersionId != nil {
		v := *s.SubscriptionDefinitionVersionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "SubscriptionDefinitionVersionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Information about a subscription definition version.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/GetSubscriptionDefinitionVersionResponse
type GetSubscriptionDefinitionVersionOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the subscription definition version.
	Arn *string `type:"string"`

	// The time, in milliseconds since the epoch, when the subscription definition
	// version was created.
	CreationTimestamp *string `type:"string"`

	// Information about the subscription definition version.
	Definition *SubscriptionDefinitionVersion `type:"structure"`

	// The ID of the subscription definition version.
	Id *string `type:"string"`

	// The token for the next set of results, or ''null'' if there are no additional
	// results.
	NextToken *string `type:"string"`

	// The version of the subscription definition version.
	Version *string `type:"string"`
}

// String returns the string representation
func (s GetSubscriptionDefinitionVersionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetSubscriptionDefinitionVersionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CreationTimestamp != nil {
		v := *s.CreationTimestamp

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CreationTimestamp", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Definition != nil {
		v := s.Definition

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Definition", v, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Version", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetSubscriptionDefinitionVersion = "GetSubscriptionDefinitionVersion"

// GetSubscriptionDefinitionVersionRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Retrieves information about a subscription definition version.
//
//    // Example sending a request using GetSubscriptionDefinitionVersionRequest.
//    req := client.GetSubscriptionDefinitionVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/GetSubscriptionDefinitionVersion
func (c *Client) GetSubscriptionDefinitionVersionRequest(input *GetSubscriptionDefinitionVersionInput) GetSubscriptionDefinitionVersionRequest {
	op := &aws.Operation{
		Name:       opGetSubscriptionDefinitionVersion,
		HTTPMethod: "GET",
		HTTPPath:   "/greengrass/definition/subscriptions/{SubscriptionDefinitionId}/versions/{SubscriptionDefinitionVersionId}",
	}

	if input == nil {
		input = &GetSubscriptionDefinitionVersionInput{}
	}

	req := c.newRequest(op, input, &GetSubscriptionDefinitionVersionOutput{})
	return GetSubscriptionDefinitionVersionRequest{Request: req, Input: input, Copy: c.GetSubscriptionDefinitionVersionRequest}
}

// GetSubscriptionDefinitionVersionRequest is the request type for the
// GetSubscriptionDefinitionVersion API operation.
type GetSubscriptionDefinitionVersionRequest struct {
	*aws.Request
	Input *GetSubscriptionDefinitionVersionInput
	Copy  func(*GetSubscriptionDefinitionVersionInput) GetSubscriptionDefinitionVersionRequest
}

// Send marshals and sends the GetSubscriptionDefinitionVersion API request.
func (r GetSubscriptionDefinitionVersionRequest) Send(ctx context.Context) (*GetSubscriptionDefinitionVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSubscriptionDefinitionVersionResponse{
		GetSubscriptionDefinitionVersionOutput: r.Request.Data.(*GetSubscriptionDefinitionVersionOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetSubscriptionDefinitionVersionResponse is the response type for the
// GetSubscriptionDefinitionVersion API operation.
type GetSubscriptionDefinitionVersionResponse struct {
	*GetSubscriptionDefinitionVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSubscriptionDefinitionVersion request.
func (r *GetSubscriptionDefinitionVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

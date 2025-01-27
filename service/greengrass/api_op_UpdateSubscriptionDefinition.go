// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/UpdateSubscriptionDefinitionRequest
type UpdateSubscriptionDefinitionInput struct {
	_ struct{} `type:"structure"`

	Name *string `type:"string"`

	// SubscriptionDefinitionId is a required field
	SubscriptionDefinitionId *string `location:"uri" locationName:"SubscriptionDefinitionId" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateSubscriptionDefinitionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateSubscriptionDefinitionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateSubscriptionDefinitionInput"}

	if s.SubscriptionDefinitionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SubscriptionDefinitionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateSubscriptionDefinitionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SubscriptionDefinitionId != nil {
		v := *s.SubscriptionDefinitionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "SubscriptionDefinitionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/UpdateSubscriptionDefinitionResponse
type UpdateSubscriptionDefinitionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateSubscriptionDefinitionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateSubscriptionDefinitionOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opUpdateSubscriptionDefinition = "UpdateSubscriptionDefinition"

// UpdateSubscriptionDefinitionRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Updates a subscription definition.
//
//    // Example sending a request using UpdateSubscriptionDefinitionRequest.
//    req := client.UpdateSubscriptionDefinitionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/UpdateSubscriptionDefinition
func (c *Client) UpdateSubscriptionDefinitionRequest(input *UpdateSubscriptionDefinitionInput) UpdateSubscriptionDefinitionRequest {
	op := &aws.Operation{
		Name:       opUpdateSubscriptionDefinition,
		HTTPMethod: "PUT",
		HTTPPath:   "/greengrass/definition/subscriptions/{SubscriptionDefinitionId}",
	}

	if input == nil {
		input = &UpdateSubscriptionDefinitionInput{}
	}

	req := c.newRequest(op, input, &UpdateSubscriptionDefinitionOutput{})
	return UpdateSubscriptionDefinitionRequest{Request: req, Input: input, Copy: c.UpdateSubscriptionDefinitionRequest}
}

// UpdateSubscriptionDefinitionRequest is the request type for the
// UpdateSubscriptionDefinition API operation.
type UpdateSubscriptionDefinitionRequest struct {
	*aws.Request
	Input *UpdateSubscriptionDefinitionInput
	Copy  func(*UpdateSubscriptionDefinitionInput) UpdateSubscriptionDefinitionRequest
}

// Send marshals and sends the UpdateSubscriptionDefinition API request.
func (r UpdateSubscriptionDefinitionRequest) Send(ctx context.Context) (*UpdateSubscriptionDefinitionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateSubscriptionDefinitionResponse{
		UpdateSubscriptionDefinitionOutput: r.Request.Data.(*UpdateSubscriptionDefinitionOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateSubscriptionDefinitionResponse is the response type for the
// UpdateSubscriptionDefinition API operation.
type UpdateSubscriptionDefinitionResponse struct {
	*UpdateSubscriptionDefinitionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateSubscriptionDefinition request.
func (r *UpdateSubscriptionDefinitionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

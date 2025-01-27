// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/DeleteSubscriptionDefinitionRequest
type DeleteSubscriptionDefinitionInput struct {
	_ struct{} `type:"structure"`

	// SubscriptionDefinitionId is a required field
	SubscriptionDefinitionId *string `location:"uri" locationName:"SubscriptionDefinitionId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteSubscriptionDefinitionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteSubscriptionDefinitionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteSubscriptionDefinitionInput"}

	if s.SubscriptionDefinitionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SubscriptionDefinitionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteSubscriptionDefinitionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.SubscriptionDefinitionId != nil {
		v := *s.SubscriptionDefinitionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "SubscriptionDefinitionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/DeleteSubscriptionDefinitionResponse
type DeleteSubscriptionDefinitionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteSubscriptionDefinitionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteSubscriptionDefinitionOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteSubscriptionDefinition = "DeleteSubscriptionDefinition"

// DeleteSubscriptionDefinitionRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Deletes a subscription definition.
//
//    // Example sending a request using DeleteSubscriptionDefinitionRequest.
//    req := client.DeleteSubscriptionDefinitionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/DeleteSubscriptionDefinition
func (c *Client) DeleteSubscriptionDefinitionRequest(input *DeleteSubscriptionDefinitionInput) DeleteSubscriptionDefinitionRequest {
	op := &aws.Operation{
		Name:       opDeleteSubscriptionDefinition,
		HTTPMethod: "DELETE",
		HTTPPath:   "/greengrass/definition/subscriptions/{SubscriptionDefinitionId}",
	}

	if input == nil {
		input = &DeleteSubscriptionDefinitionInput{}
	}

	req := c.newRequest(op, input, &DeleteSubscriptionDefinitionOutput{})
	return DeleteSubscriptionDefinitionRequest{Request: req, Input: input, Copy: c.DeleteSubscriptionDefinitionRequest}
}

// DeleteSubscriptionDefinitionRequest is the request type for the
// DeleteSubscriptionDefinition API operation.
type DeleteSubscriptionDefinitionRequest struct {
	*aws.Request
	Input *DeleteSubscriptionDefinitionInput
	Copy  func(*DeleteSubscriptionDefinitionInput) DeleteSubscriptionDefinitionRequest
}

// Send marshals and sends the DeleteSubscriptionDefinition API request.
func (r DeleteSubscriptionDefinitionRequest) Send(ctx context.Context) (*DeleteSubscriptionDefinitionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteSubscriptionDefinitionResponse{
		DeleteSubscriptionDefinitionOutput: r.Request.Data.(*DeleteSubscriptionDefinitionOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteSubscriptionDefinitionResponse is the response type for the
// DeleteSubscriptionDefinition API operation.
type DeleteSubscriptionDefinitionResponse struct {
	*DeleteSubscriptionDefinitionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteSubscriptionDefinition request.
func (r *DeleteSubscriptionDefinitionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

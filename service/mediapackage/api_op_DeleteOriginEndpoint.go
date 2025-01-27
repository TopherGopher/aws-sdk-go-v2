// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediapackage

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-2017-10-12/DeleteOriginEndpointRequest
type DeleteOriginEndpointInput struct {
	_ struct{} `type:"structure"`

	// Id is a required field
	Id *string `location:"uri" locationName:"id" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteOriginEndpointInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteOriginEndpointInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteOriginEndpointInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteOriginEndpointInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-2017-10-12/DeleteOriginEndpointResponse
type DeleteOriginEndpointOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteOriginEndpointOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteOriginEndpointOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteOriginEndpoint = "DeleteOriginEndpoint"

// DeleteOriginEndpointRequest returns a request value for making API operation for
// AWS Elemental MediaPackage.
//
// Deletes an existing OriginEndpoint.
//
//    // Example sending a request using DeleteOriginEndpointRequest.
//    req := client.DeleteOriginEndpointRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-2017-10-12/DeleteOriginEndpoint
func (c *Client) DeleteOriginEndpointRequest(input *DeleteOriginEndpointInput) DeleteOriginEndpointRequest {
	op := &aws.Operation{
		Name:       opDeleteOriginEndpoint,
		HTTPMethod: "DELETE",
		HTTPPath:   "/origin_endpoints/{id}",
	}

	if input == nil {
		input = &DeleteOriginEndpointInput{}
	}

	req := c.newRequest(op, input, &DeleteOriginEndpointOutput{})
	return DeleteOriginEndpointRequest{Request: req, Input: input, Copy: c.DeleteOriginEndpointRequest}
}

// DeleteOriginEndpointRequest is the request type for the
// DeleteOriginEndpoint API operation.
type DeleteOriginEndpointRequest struct {
	*aws.Request
	Input *DeleteOriginEndpointInput
	Copy  func(*DeleteOriginEndpointInput) DeleteOriginEndpointRequest
}

// Send marshals and sends the DeleteOriginEndpoint API request.
func (r DeleteOriginEndpointRequest) Send(ctx context.Context) (*DeleteOriginEndpointResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteOriginEndpointResponse{
		DeleteOriginEndpointOutput: r.Request.Data.(*DeleteOriginEndpointOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteOriginEndpointResponse is the response type for the
// DeleteOriginEndpoint API operation.
type DeleteOriginEndpointResponse struct {
	*DeleteOriginEndpointOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteOriginEndpoint request.
func (r *DeleteOriginEndpointResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

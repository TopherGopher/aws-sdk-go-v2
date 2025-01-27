// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package worklink

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/worklink-2018-09-25/RevokeDomainAccessRequest
type RevokeDomainAccessInput struct {
	_ struct{} `type:"structure"`

	// The name of the domain.
	//
	// DomainName is a required field
	DomainName *string `min:"1" type:"string" required:"true"`

	// The ARN of the fleet.
	//
	// FleetArn is a required field
	FleetArn *string `min:"20" type:"string" required:"true"`
}

// String returns the string representation
func (s RevokeDomainAccessInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RevokeDomainAccessInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RevokeDomainAccessInput"}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}
	if s.DomainName != nil && len(*s.DomainName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DomainName", 1))
	}

	if s.FleetArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("FleetArn"))
	}
	if s.FleetArn != nil && len(*s.FleetArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("FleetArn", 20))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RevokeDomainAccessInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.DomainName != nil {
		v := *s.DomainName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DomainName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.FleetArn != nil {
		v := *s.FleetArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FleetArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/worklink-2018-09-25/RevokeDomainAccessResponse
type RevokeDomainAccessOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RevokeDomainAccessOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RevokeDomainAccessOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opRevokeDomainAccess = "RevokeDomainAccess"

// RevokeDomainAccessRequest returns a request value for making API operation for
// Amazon WorkLink.
//
// Moves a domain to INACTIVE status if it was in the ACTIVE status.
//
//    // Example sending a request using RevokeDomainAccessRequest.
//    req := client.RevokeDomainAccessRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/worklink-2018-09-25/RevokeDomainAccess
func (c *Client) RevokeDomainAccessRequest(input *RevokeDomainAccessInput) RevokeDomainAccessRequest {
	op := &aws.Operation{
		Name:       opRevokeDomainAccess,
		HTTPMethod: "POST",
		HTTPPath:   "/revokeDomainAccess",
	}

	if input == nil {
		input = &RevokeDomainAccessInput{}
	}

	req := c.newRequest(op, input, &RevokeDomainAccessOutput{})
	return RevokeDomainAccessRequest{Request: req, Input: input, Copy: c.RevokeDomainAccessRequest}
}

// RevokeDomainAccessRequest is the request type for the
// RevokeDomainAccess API operation.
type RevokeDomainAccessRequest struct {
	*aws.Request
	Input *RevokeDomainAccessInput
	Copy  func(*RevokeDomainAccessInput) RevokeDomainAccessRequest
}

// Send marshals and sends the RevokeDomainAccess API request.
func (r RevokeDomainAccessRequest) Send(ctx context.Context) (*RevokeDomainAccessResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RevokeDomainAccessResponse{
		RevokeDomainAccessOutput: r.Request.Data.(*RevokeDomainAccessOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RevokeDomainAccessResponse is the response type for the
// RevokeDomainAccess API operation.
type RevokeDomainAccessResponse struct {
	*RevokeDomainAccessOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RevokeDomainAccess request.
func (r *RevokeDomainAccessResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

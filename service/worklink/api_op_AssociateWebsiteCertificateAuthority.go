// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package worklink

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/worklink-2018-09-25/AssociateWebsiteCertificateAuthorityRequest
type AssociateWebsiteCertificateAuthorityInput struct {
	_ struct{} `type:"structure"`

	// The root certificate of the CA.
	//
	// Certificate is a required field
	Certificate *string `min:"1" type:"string" required:"true"`

	// The certificate name to display.
	DisplayName *string `type:"string"`

	// The ARN of the fleet.
	//
	// FleetArn is a required field
	FleetArn *string `min:"20" type:"string" required:"true"`
}

// String returns the string representation
func (s AssociateWebsiteCertificateAuthorityInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociateWebsiteCertificateAuthorityInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssociateWebsiteCertificateAuthorityInput"}

	if s.Certificate == nil {
		invalidParams.Add(aws.NewErrParamRequired("Certificate"))
	}
	if s.Certificate != nil && len(*s.Certificate) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Certificate", 1))
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
func (s AssociateWebsiteCertificateAuthorityInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Certificate != nil {
		v := *s.Certificate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Certificate", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DisplayName != nil {
		v := *s.DisplayName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DisplayName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.FleetArn != nil {
		v := *s.FleetArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FleetArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/worklink-2018-09-25/AssociateWebsiteCertificateAuthorityResponse
type AssociateWebsiteCertificateAuthorityOutput struct {
	_ struct{} `type:"structure"`

	// A unique identifier for the CA.
	WebsiteCaId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s AssociateWebsiteCertificateAuthorityOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AssociateWebsiteCertificateAuthorityOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.WebsiteCaId != nil {
		v := *s.WebsiteCaId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "WebsiteCaId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opAssociateWebsiteCertificateAuthority = "AssociateWebsiteCertificateAuthority"

// AssociateWebsiteCertificateAuthorityRequest returns a request value for making API operation for
// Amazon WorkLink.
//
// Imports the root certificate of a certificate authority (CA) used to obtain
// TLS certificates used by associated websites within the company network.
//
//    // Example sending a request using AssociateWebsiteCertificateAuthorityRequest.
//    req := client.AssociateWebsiteCertificateAuthorityRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/worklink-2018-09-25/AssociateWebsiteCertificateAuthority
func (c *Client) AssociateWebsiteCertificateAuthorityRequest(input *AssociateWebsiteCertificateAuthorityInput) AssociateWebsiteCertificateAuthorityRequest {
	op := &aws.Operation{
		Name:       opAssociateWebsiteCertificateAuthority,
		HTTPMethod: "POST",
		HTTPPath:   "/associateWebsiteCertificateAuthority",
	}

	if input == nil {
		input = &AssociateWebsiteCertificateAuthorityInput{}
	}

	req := c.newRequest(op, input, &AssociateWebsiteCertificateAuthorityOutput{})
	return AssociateWebsiteCertificateAuthorityRequest{Request: req, Input: input, Copy: c.AssociateWebsiteCertificateAuthorityRequest}
}

// AssociateWebsiteCertificateAuthorityRequest is the request type for the
// AssociateWebsiteCertificateAuthority API operation.
type AssociateWebsiteCertificateAuthorityRequest struct {
	*aws.Request
	Input *AssociateWebsiteCertificateAuthorityInput
	Copy  func(*AssociateWebsiteCertificateAuthorityInput) AssociateWebsiteCertificateAuthorityRequest
}

// Send marshals and sends the AssociateWebsiteCertificateAuthority API request.
func (r AssociateWebsiteCertificateAuthorityRequest) Send(ctx context.Context) (*AssociateWebsiteCertificateAuthorityResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateWebsiteCertificateAuthorityResponse{
		AssociateWebsiteCertificateAuthorityOutput: r.Request.Data.(*AssociateWebsiteCertificateAuthorityOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateWebsiteCertificateAuthorityResponse is the response type for the
// AssociateWebsiteCertificateAuthority API operation.
type AssociateWebsiteCertificateAuthorityResponse struct {
	*AssociateWebsiteCertificateAuthorityOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateWebsiteCertificateAuthority request.
func (r *AssociateWebsiteCertificateAuthorityResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

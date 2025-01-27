// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package amplify

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Request structure for the delete Domain Association request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/DeleteDomainAssociationRequest
type DeleteDomainAssociationInput struct {
	_ struct{} `type:"structure"`

	// Unique Id for an Amplify App.
	//
	// AppId is a required field
	AppId *string `location:"uri" locationName:"appId" min:"1" type:"string" required:"true"`

	// Name of the domain.
	//
	// DomainName is a required field
	DomainName *string `location:"uri" locationName:"domainName" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteDomainAssociationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDomainAssociationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDomainAssociationInput"}

	if s.AppId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AppId"))
	}
	if s.AppId != nil && len(*s.AppId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AppId", 1))
	}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteDomainAssociationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AppId != nil {
		v := *s.AppId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "appId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DomainName != nil {
		v := *s.DomainName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "domainName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/DeleteDomainAssociationResult
type DeleteDomainAssociationOutput struct {
	_ struct{} `type:"structure"`

	// Structure for Domain Association, which associates a custom domain with an
	// Amplify App.
	//
	// DomainAssociation is a required field
	DomainAssociation *DomainAssociation `locationName:"domainAssociation" type:"structure" required:"true"`
}

// String returns the string representation
func (s DeleteDomainAssociationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteDomainAssociationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DomainAssociation != nil {
		v := s.DomainAssociation

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "domainAssociation", v, metadata)
	}
	return nil
}

const opDeleteDomainAssociation = "DeleteDomainAssociation"

// DeleteDomainAssociationRequest returns a request value for making API operation for
// AWS Amplify.
//
// Deletes a DomainAssociation.
//
//    // Example sending a request using DeleteDomainAssociationRequest.
//    req := client.DeleteDomainAssociationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/DeleteDomainAssociation
func (c *Client) DeleteDomainAssociationRequest(input *DeleteDomainAssociationInput) DeleteDomainAssociationRequest {
	op := &aws.Operation{
		Name:       opDeleteDomainAssociation,
		HTTPMethod: "DELETE",
		HTTPPath:   "/apps/{appId}/domains/{domainName}",
	}

	if input == nil {
		input = &DeleteDomainAssociationInput{}
	}

	req := c.newRequest(op, input, &DeleteDomainAssociationOutput{})
	return DeleteDomainAssociationRequest{Request: req, Input: input, Copy: c.DeleteDomainAssociationRequest}
}

// DeleteDomainAssociationRequest is the request type for the
// DeleteDomainAssociation API operation.
type DeleteDomainAssociationRequest struct {
	*aws.Request
	Input *DeleteDomainAssociationInput
	Copy  func(*DeleteDomainAssociationInput) DeleteDomainAssociationRequest
}

// Send marshals and sends the DeleteDomainAssociation API request.
func (r DeleteDomainAssociationRequest) Send(ctx context.Context) (*DeleteDomainAssociationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDomainAssociationResponse{
		DeleteDomainAssociationOutput: r.Request.Data.(*DeleteDomainAssociationOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDomainAssociationResponse is the response type for the
// DeleteDomainAssociation API operation.
type DeleteDomainAssociationResponse struct {
	*DeleteDomainAssociationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDomainAssociation request.
func (r *DeleteDomainAssociationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

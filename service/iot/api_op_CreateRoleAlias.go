// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateRoleAliasInput struct {
	_ struct{} `type:"structure"`

	// How long (in seconds) the credentials will be valid.
	CredentialDurationSeconds *int64 `locationName:"credentialDurationSeconds" min:"900" type:"integer"`

	// The role alias that points to a role ARN. This allows you to change the role
	// without having to update the device.
	//
	// RoleAlias is a required field
	RoleAlias *string `location:"uri" locationName:"roleAlias" min:"1" type:"string" required:"true"`

	// The role ARN.
	//
	// RoleArn is a required field
	RoleArn *string `locationName:"roleArn" min:"20" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateRoleAliasInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateRoleAliasInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateRoleAliasInput"}
	if s.CredentialDurationSeconds != nil && *s.CredentialDurationSeconds < 900 {
		invalidParams.Add(aws.NewErrParamMinValue("CredentialDurationSeconds", 900))
	}

	if s.RoleAlias == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoleAlias"))
	}
	if s.RoleAlias != nil && len(*s.RoleAlias) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RoleAlias", 1))
	}

	if s.RoleArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoleArn"))
	}
	if s.RoleArn != nil && len(*s.RoleArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("RoleArn", 20))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateRoleAliasInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.CredentialDurationSeconds != nil {
		v := *s.CredentialDurationSeconds

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "credentialDurationSeconds", protocol.Int64Value(v), metadata)
	}
	if s.RoleArn != nil {
		v := *s.RoleArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "roleArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RoleAlias != nil {
		v := *s.RoleAlias

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "roleAlias", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CreateRoleAliasOutput struct {
	_ struct{} `type:"structure"`

	// The role alias.
	RoleAlias *string `locationName:"roleAlias" min:"1" type:"string"`

	// The role alias ARN.
	RoleAliasArn *string `locationName:"roleAliasArn" type:"string"`
}

// String returns the string representation
func (s CreateRoleAliasOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateRoleAliasOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.RoleAlias != nil {
		v := *s.RoleAlias

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "roleAlias", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RoleAliasArn != nil {
		v := *s.RoleAliasArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "roleAliasArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateRoleAlias = "CreateRoleAlias"

// CreateRoleAliasRequest returns a request value for making API operation for
// AWS IoT.
//
// Creates a role alias.
//
//    // Example sending a request using CreateRoleAliasRequest.
//    req := client.CreateRoleAliasRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CreateRoleAliasRequest(input *CreateRoleAliasInput) CreateRoleAliasRequest {
	op := &aws.Operation{
		Name:       opCreateRoleAlias,
		HTTPMethod: "POST",
		HTTPPath:   "/role-aliases/{roleAlias}",
	}

	if input == nil {
		input = &CreateRoleAliasInput{}
	}

	req := c.newRequest(op, input, &CreateRoleAliasOutput{})
	return CreateRoleAliasRequest{Request: req, Input: input, Copy: c.CreateRoleAliasRequest}
}

// CreateRoleAliasRequest is the request type for the
// CreateRoleAlias API operation.
type CreateRoleAliasRequest struct {
	*aws.Request
	Input *CreateRoleAliasInput
	Copy  func(*CreateRoleAliasInput) CreateRoleAliasRequest
}

// Send marshals and sends the CreateRoleAlias API request.
func (r CreateRoleAliasRequest) Send(ctx context.Context) (*CreateRoleAliasResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateRoleAliasResponse{
		CreateRoleAliasOutput: r.Request.Data.(*CreateRoleAliasOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateRoleAliasResponse is the response type for the
// CreateRoleAlias API operation.
type CreateRoleAliasResponse struct {
	*CreateRoleAliasOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateRoleAlias request.
func (r *CreateRoleAliasResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

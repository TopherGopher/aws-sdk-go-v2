// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworkscm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworkscm-2016-11-01/ExportServerEngineAttributeRequest
type ExportServerEngineAttributeInput struct {
	_ struct{} `type:"structure"`

	// The name of the export attribute. Currently, the supported export attribute
	// is Userdata. This exports a user data script that includes parameters and
	// values provided in the InputAttributes list.
	//
	// ExportAttributeName is a required field
	ExportAttributeName *string `type:"string" required:"true"`

	// The list of engine attributes. The list type is EngineAttribute. An EngineAttribute
	// list item is a pair that includes an attribute name and its value. For the
	// Userdata ExportAttributeName, the following are supported engine attribute
	// names.
	//
	//    * RunList In Chef, a list of roles or recipes that are run in the specified
	//    order. In Puppet, this parameter is ignored.
	//
	//    * OrganizationName In Chef, an organization name. AWS OpsWorks for Chef
	//    Automate always creates the organization default. In Puppet, this parameter
	//    is ignored.
	//
	//    * NodeEnvironment In Chef, a node environment (for example, development,
	//    staging, or one-box). In Puppet, this parameter is ignored.
	//
	//    * NodeClientVersion In Chef, the version of the Chef engine (three numbers
	//    separated by dots, such as 13.8.5). If this attribute is empty, OpsWorks
	//    for Chef Automate uses the most current version. In Puppet, this parameter
	//    is ignored.
	InputAttributes []EngineAttribute `type:"list"`

	// The name of the server from which you are exporting the attribute.
	//
	// ServerName is a required field
	ServerName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ExportServerEngineAttributeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ExportServerEngineAttributeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ExportServerEngineAttributeInput"}

	if s.ExportAttributeName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ExportAttributeName"))
	}

	if s.ServerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServerName"))
	}
	if s.ServerName != nil && len(*s.ServerName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ServerName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworkscm-2016-11-01/ExportServerEngineAttributeResponse
type ExportServerEngineAttributeOutput struct {
	_ struct{} `type:"structure"`

	// The requested engine attribute pair with attribute name and value.
	EngineAttribute *EngineAttribute `type:"structure"`

	// The server name used in the request.
	ServerName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ExportServerEngineAttributeOutput) String() string {
	return awsutil.Prettify(s)
}

const opExportServerEngineAttribute = "ExportServerEngineAttribute"

// ExportServerEngineAttributeRequest returns a request value for making API operation for
// AWS OpsWorks CM.
//
// Exports a specified server engine attribute as a base64-encoded string. For
// example, you can export user data that you can use in EC2 to associate nodes
// with a server.
//
// This operation is synchronous.
//
// A ValidationException is raised when parameters of the request are not valid.
// A ResourceNotFoundException is thrown when the server does not exist. An
// InvalidStateException is thrown when the server is in any of the following
// states: CREATING, TERMINATED, FAILED or DELETING.
//
//    // Example sending a request using ExportServerEngineAttributeRequest.
//    req := client.ExportServerEngineAttributeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworkscm-2016-11-01/ExportServerEngineAttribute
func (c *Client) ExportServerEngineAttributeRequest(input *ExportServerEngineAttributeInput) ExportServerEngineAttributeRequest {
	op := &aws.Operation{
		Name:       opExportServerEngineAttribute,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ExportServerEngineAttributeInput{}
	}

	req := c.newRequest(op, input, &ExportServerEngineAttributeOutput{})
	return ExportServerEngineAttributeRequest{Request: req, Input: input, Copy: c.ExportServerEngineAttributeRequest}
}

// ExportServerEngineAttributeRequest is the request type for the
// ExportServerEngineAttribute API operation.
type ExportServerEngineAttributeRequest struct {
	*aws.Request
	Input *ExportServerEngineAttributeInput
	Copy  func(*ExportServerEngineAttributeInput) ExportServerEngineAttributeRequest
}

// Send marshals and sends the ExportServerEngineAttribute API request.
func (r ExportServerEngineAttributeRequest) Send(ctx context.Context) (*ExportServerEngineAttributeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ExportServerEngineAttributeResponse{
		ExportServerEngineAttributeOutput: r.Request.Data.(*ExportServerEngineAttributeOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ExportServerEngineAttributeResponse is the response type for the
// ExportServerEngineAttribute API operation.
type ExportServerEngineAttributeResponse struct {
	*ExportServerEngineAttributeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ExportServerEngineAttribute request.
func (r *ExportServerEngineAttributeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

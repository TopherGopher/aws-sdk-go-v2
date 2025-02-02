// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot1clickprojects

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot1click-projects-2018-05-14/UpdateProjectRequest
type UpdateProjectInput struct {
	_ struct{} `type:"structure"`

	// An optional user-defined description for the project.
	Description *string `locationName:"description" type:"string"`

	// An object defining the project update. Once a project has been created, you
	// cannot add device template names to the project. However, for a given placementTemplate,
	// you can update the associated callbackOverrides for the device definition
	// using this API.
	PlacementTemplate *PlacementTemplate `locationName:"placementTemplate" type:"structure"`

	// The name of the project to be updated.
	//
	// ProjectName is a required field
	ProjectName *string `location:"uri" locationName:"projectName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateProjectInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateProjectInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateProjectInput"}

	if s.ProjectName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProjectName"))
	}
	if s.ProjectName != nil && len(*s.ProjectName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProjectName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateProjectInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PlacementTemplate != nil {
		v := s.PlacementTemplate

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "placementTemplate", v, metadata)
	}
	if s.ProjectName != nil {
		v := *s.ProjectName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "projectName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot1click-projects-2018-05-14/UpdateProjectResponse
type UpdateProjectOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateProjectOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateProjectOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opUpdateProject = "UpdateProject"

// UpdateProjectRequest returns a request value for making API operation for
// AWS IoT 1-Click Projects Service.
//
// Updates a project associated with your AWS account and region. With the exception
// of device template names, you can pass just the values that need to be updated
// because the update request will change only the values that are provided.
// To clear a value, pass the empty string (i.e., "").
//
//    // Example sending a request using UpdateProjectRequest.
//    req := client.UpdateProjectRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot1click-projects-2018-05-14/UpdateProject
func (c *Client) UpdateProjectRequest(input *UpdateProjectInput) UpdateProjectRequest {
	op := &aws.Operation{
		Name:       opUpdateProject,
		HTTPMethod: "PUT",
		HTTPPath:   "/projects/{projectName}",
	}

	if input == nil {
		input = &UpdateProjectInput{}
	}

	req := c.newRequest(op, input, &UpdateProjectOutput{})
	return UpdateProjectRequest{Request: req, Input: input, Copy: c.UpdateProjectRequest}
}

// UpdateProjectRequest is the request type for the
// UpdateProject API operation.
type UpdateProjectRequest struct {
	*aws.Request
	Input *UpdateProjectInput
	Copy  func(*UpdateProjectInput) UpdateProjectRequest
}

// Send marshals and sends the UpdateProject API request.
func (r UpdateProjectRequest) Send(ctx context.Context) (*UpdateProjectResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateProjectResponse{
		UpdateProjectOutput: r.Request.Data.(*UpdateProjectOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateProjectResponse is the response type for the
// UpdateProject API operation.
type UpdateProjectResponse struct {
	*UpdateProjectOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateProject request.
func (r *UpdateProjectResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

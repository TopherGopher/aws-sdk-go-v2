// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package robomaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/DeleteRobotApplicationRequest
type DeleteRobotApplicationInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the the robot application.
	//
	// Application is a required field
	Application *string `locationName:"application" min:"1" type:"string" required:"true"`

	// The version of the robot application to delete.
	ApplicationVersion *string `locationName:"applicationVersion" min:"1" type:"string"`
}

// String returns the string representation
func (s DeleteRobotApplicationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteRobotApplicationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteRobotApplicationInput"}

	if s.Application == nil {
		invalidParams.Add(aws.NewErrParamRequired("Application"))
	}
	if s.Application != nil && len(*s.Application) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Application", 1))
	}
	if s.ApplicationVersion != nil && len(*s.ApplicationVersion) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApplicationVersion", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteRobotApplicationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Application != nil {
		v := *s.Application

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "application", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ApplicationVersion != nil {
		v := *s.ApplicationVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "applicationVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/DeleteRobotApplicationResponse
type DeleteRobotApplicationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteRobotApplicationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteRobotApplicationOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteRobotApplication = "DeleteRobotApplication"

// DeleteRobotApplicationRequest returns a request value for making API operation for
// AWS RoboMaker.
//
// Deletes a robot application.
//
//    // Example sending a request using DeleteRobotApplicationRequest.
//    req := client.DeleteRobotApplicationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/DeleteRobotApplication
func (c *Client) DeleteRobotApplicationRequest(input *DeleteRobotApplicationInput) DeleteRobotApplicationRequest {
	op := &aws.Operation{
		Name:       opDeleteRobotApplication,
		HTTPMethod: "POST",
		HTTPPath:   "/deleteRobotApplication",
	}

	if input == nil {
		input = &DeleteRobotApplicationInput{}
	}

	req := c.newRequest(op, input, &DeleteRobotApplicationOutput{})
	return DeleteRobotApplicationRequest{Request: req, Input: input, Copy: c.DeleteRobotApplicationRequest}
}

// DeleteRobotApplicationRequest is the request type for the
// DeleteRobotApplication API operation.
type DeleteRobotApplicationRequest struct {
	*aws.Request
	Input *DeleteRobotApplicationInput
	Copy  func(*DeleteRobotApplicationInput) DeleteRobotApplicationRequest
}

// Send marshals and sends the DeleteRobotApplication API request.
func (r DeleteRobotApplicationRequest) Send(ctx context.Context) (*DeleteRobotApplicationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteRobotApplicationResponse{
		DeleteRobotApplicationOutput: r.Request.Data.(*DeleteRobotApplicationOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteRobotApplicationResponse is the response type for the
// DeleteRobotApplication API operation.
type DeleteRobotApplicationResponse struct {
	*DeleteRobotApplicationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteRobotApplication request.
func (r *DeleteRobotApplicationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

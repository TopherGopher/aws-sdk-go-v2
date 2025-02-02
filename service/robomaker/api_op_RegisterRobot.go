// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package robomaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/RegisterRobotRequest
type RegisterRobotInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the fleet.
	//
	// Fleet is a required field
	Fleet *string `locationName:"fleet" min:"1" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the robot.
	//
	// Robot is a required field
	Robot *string `locationName:"robot" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s RegisterRobotInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RegisterRobotInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RegisterRobotInput"}

	if s.Fleet == nil {
		invalidParams.Add(aws.NewErrParamRequired("Fleet"))
	}
	if s.Fleet != nil && len(*s.Fleet) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Fleet", 1))
	}

	if s.Robot == nil {
		invalidParams.Add(aws.NewErrParamRequired("Robot"))
	}
	if s.Robot != nil && len(*s.Robot) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Robot", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RegisterRobotInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Fleet != nil {
		v := *s.Fleet

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "fleet", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Robot != nil {
		v := *s.Robot

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "robot", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/RegisterRobotResponse
type RegisterRobotOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the fleet that the robot will join.
	Fleet *string `locationName:"fleet" min:"1" type:"string"`

	// Information about the robot registration.
	Robot *string `locationName:"robot" min:"1" type:"string"`
}

// String returns the string representation
func (s RegisterRobotOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RegisterRobotOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Fleet != nil {
		v := *s.Fleet

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "fleet", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Robot != nil {
		v := *s.Robot

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "robot", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opRegisterRobot = "RegisterRobot"

// RegisterRobotRequest returns a request value for making API operation for
// AWS RoboMaker.
//
// Registers a robot with a fleet.
//
//    // Example sending a request using RegisterRobotRequest.
//    req := client.RegisterRobotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/RegisterRobot
func (c *Client) RegisterRobotRequest(input *RegisterRobotInput) RegisterRobotRequest {
	op := &aws.Operation{
		Name:       opRegisterRobot,
		HTTPMethod: "POST",
		HTTPPath:   "/registerRobot",
	}

	if input == nil {
		input = &RegisterRobotInput{}
	}

	req := c.newRequest(op, input, &RegisterRobotOutput{})
	return RegisterRobotRequest{Request: req, Input: input, Copy: c.RegisterRobotRequest}
}

// RegisterRobotRequest is the request type for the
// RegisterRobot API operation.
type RegisterRobotRequest struct {
	*aws.Request
	Input *RegisterRobotInput
	Copy  func(*RegisterRobotInput) RegisterRobotRequest
}

// Send marshals and sends the RegisterRobot API request.
func (r RegisterRobotRequest) Send(ctx context.Context) (*RegisterRobotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RegisterRobotResponse{
		RegisterRobotOutput: r.Request.Data.(*RegisterRobotOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RegisterRobotResponse is the response type for the
// RegisterRobot API operation.
type RegisterRobotResponse struct {
	*RegisterRobotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RegisterRobot request.
func (r *RegisterRobotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

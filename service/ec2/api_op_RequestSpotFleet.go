// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for RequestSpotFleet.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/RequestSpotFleetRequest
type RequestSpotFleetInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The configuration for the Spot Fleet request.
	//
	// SpotFleetRequestConfig is a required field
	SpotFleetRequestConfig *SpotFleetRequestConfigData `locationName:"spotFleetRequestConfig" type:"structure" required:"true"`
}

// String returns the string representation
func (s RequestSpotFleetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RequestSpotFleetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RequestSpotFleetInput"}

	if s.SpotFleetRequestConfig == nil {
		invalidParams.Add(aws.NewErrParamRequired("SpotFleetRequestConfig"))
	}
	if s.SpotFleetRequestConfig != nil {
		if err := s.SpotFleetRequestConfig.Validate(); err != nil {
			invalidParams.AddNested("SpotFleetRequestConfig", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of RequestSpotFleet.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/RequestSpotFleetResponse
type RequestSpotFleetOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the Spot Fleet request.
	SpotFleetRequestId *string `locationName:"spotFleetRequestId" type:"string"`
}

// String returns the string representation
func (s RequestSpotFleetOutput) String() string {
	return awsutil.Prettify(s)
}

const opRequestSpotFleet = "RequestSpotFleet"

// RequestSpotFleetRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Creates a Spot Fleet request.
//
// The Spot Fleet request specifies the total target capacity and the On-Demand
// target capacity. Amazon EC2 calculates the difference between the total capacity
// and On-Demand capacity, and launches the difference as Spot capacity.
//
// You can submit a single request that includes multiple launch specifications
// that vary by instance type, AMI, Availability Zone, or subnet.
//
// By default, the Spot Fleet requests Spot Instances in the Spot Instance pool
// where the price per unit is the lowest. Each launch specification can include
// its own instance weighting that reflects the value of the instance type to
// your application workload.
//
// Alternatively, you can specify that the Spot Fleet distribute the target
// capacity across the Spot pools included in its launch specifications. By
// ensuring that the Spot Instances in your Spot Fleet are in different Spot
// pools, you can improve the availability of your fleet.
//
// You can specify tags for the Spot Instances. You cannot tag other resource
// types in a Spot Fleet request because only the instance resource type is
// supported.
//
// For more information, see Spot Fleet Requests (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-fleet-requests.html)
// in the Amazon EC2 User Guide for Linux Instances.
//
//    // Example sending a request using RequestSpotFleetRequest.
//    req := client.RequestSpotFleetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/RequestSpotFleet
func (c *Client) RequestSpotFleetRequest(input *RequestSpotFleetInput) RequestSpotFleetRequest {
	op := &aws.Operation{
		Name:       opRequestSpotFleet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RequestSpotFleetInput{}
	}

	req := c.newRequest(op, input, &RequestSpotFleetOutput{})
	return RequestSpotFleetRequest{Request: req, Input: input, Copy: c.RequestSpotFleetRequest}
}

// RequestSpotFleetRequest is the request type for the
// RequestSpotFleet API operation.
type RequestSpotFleetRequest struct {
	*aws.Request
	Input *RequestSpotFleetInput
	Copy  func(*RequestSpotFleetInput) RequestSpotFleetRequest
}

// Send marshals and sends the RequestSpotFleet API request.
func (r RequestSpotFleetRequest) Send(ctx context.Context) (*RequestSpotFleetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RequestSpotFleetResponse{
		RequestSpotFleetOutput: r.Request.Data.(*RequestSpotFleetOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RequestSpotFleetResponse is the response type for the
// RequestSpotFleet API operation.
type RequestSpotFleetResponse struct {
	*RequestSpotFleetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RequestSpotFleet request.
func (r *RequestSpotFleetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

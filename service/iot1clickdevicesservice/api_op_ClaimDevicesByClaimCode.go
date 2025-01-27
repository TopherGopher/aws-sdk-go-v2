// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot1clickdevicesservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/devices-2018-05-14/ClaimDevicesByClaimCodeRequest
type ClaimDevicesByClaimCodeInput struct {
	_ struct{} `type:"structure"`

	// ClaimCode is a required field
	ClaimCode *string `location:"uri" locationName:"claimCode" type:"string" required:"true"`
}

// String returns the string representation
func (s ClaimDevicesByClaimCodeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ClaimDevicesByClaimCodeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ClaimDevicesByClaimCodeInput"}

	if s.ClaimCode == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClaimCode"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ClaimDevicesByClaimCodeInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ClaimCode != nil {
		v := *s.ClaimCode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "claimCode", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/devices-2018-05-14/ClaimDevicesByClaimCodeResponse
type ClaimDevicesByClaimCodeOutput struct {
	_ struct{} `type:"structure"`

	// The claim code provided by the device manufacturer.
	ClaimCode *string `locationName:"claimCode" min:"12" type:"string"`

	// The total number of devices associated with the claim code that has been
	// processed in the claim request.
	Total *int64 `locationName:"total" type:"integer"`
}

// String returns the string representation
func (s ClaimDevicesByClaimCodeOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ClaimDevicesByClaimCodeOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ClaimCode != nil {
		v := *s.ClaimCode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "claimCode", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Total != nil {
		v := *s.Total

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "total", protocol.Int64Value(v), metadata)
	}
	return nil
}

const opClaimDevicesByClaimCode = "ClaimDevicesByClaimCode"

// ClaimDevicesByClaimCodeRequest returns a request value for making API operation for
// AWS IoT 1-Click Devices Service.
//
// Adds device(s) to your account (i.e., claim one or more devices) if and only
// if you received a claim code with the device(s).
//
//    // Example sending a request using ClaimDevicesByClaimCodeRequest.
//    req := client.ClaimDevicesByClaimCodeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devices-2018-05-14/ClaimDevicesByClaimCode
func (c *Client) ClaimDevicesByClaimCodeRequest(input *ClaimDevicesByClaimCodeInput) ClaimDevicesByClaimCodeRequest {
	op := &aws.Operation{
		Name:       opClaimDevicesByClaimCode,
		HTTPMethod: "PUT",
		HTTPPath:   "/claims/{claimCode}",
	}

	if input == nil {
		input = &ClaimDevicesByClaimCodeInput{}
	}

	req := c.newRequest(op, input, &ClaimDevicesByClaimCodeOutput{})
	return ClaimDevicesByClaimCodeRequest{Request: req, Input: input, Copy: c.ClaimDevicesByClaimCodeRequest}
}

// ClaimDevicesByClaimCodeRequest is the request type for the
// ClaimDevicesByClaimCode API operation.
type ClaimDevicesByClaimCodeRequest struct {
	*aws.Request
	Input *ClaimDevicesByClaimCodeInput
	Copy  func(*ClaimDevicesByClaimCodeInput) ClaimDevicesByClaimCodeRequest
}

// Send marshals and sends the ClaimDevicesByClaimCode API request.
func (r ClaimDevicesByClaimCodeRequest) Send(ctx context.Context) (*ClaimDevicesByClaimCodeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ClaimDevicesByClaimCodeResponse{
		ClaimDevicesByClaimCodeOutput: r.Request.Data.(*ClaimDevicesByClaimCodeOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ClaimDevicesByClaimCodeResponse is the response type for the
// ClaimDevicesByClaimCode API operation.
type ClaimDevicesByClaimCodeResponse struct {
	*ClaimDevicesByClaimCodeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ClaimDevicesByClaimCode request.
func (r *ClaimDevicesByClaimCodeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

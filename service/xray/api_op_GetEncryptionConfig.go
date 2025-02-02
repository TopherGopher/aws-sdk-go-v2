// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package xray

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/xray-2016-04-12/GetEncryptionConfigRequest
type GetEncryptionConfigInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s GetEncryptionConfigInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetEncryptionConfigInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/xray-2016-04-12/GetEncryptionConfigResult
type GetEncryptionConfigOutput struct {
	_ struct{} `type:"structure"`

	// The encryption configuration document.
	EncryptionConfig *EncryptionConfig `type:"structure"`
}

// String returns the string representation
func (s GetEncryptionConfigOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetEncryptionConfigOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.EncryptionConfig != nil {
		v := s.EncryptionConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "EncryptionConfig", v, metadata)
	}
	return nil
}

const opGetEncryptionConfig = "GetEncryptionConfig"

// GetEncryptionConfigRequest returns a request value for making API operation for
// AWS X-Ray.
//
// Retrieves the current encryption configuration for X-Ray data.
//
//    // Example sending a request using GetEncryptionConfigRequest.
//    req := client.GetEncryptionConfigRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/xray-2016-04-12/GetEncryptionConfig
func (c *Client) GetEncryptionConfigRequest(input *GetEncryptionConfigInput) GetEncryptionConfigRequest {
	op := &aws.Operation{
		Name:       opGetEncryptionConfig,
		HTTPMethod: "POST",
		HTTPPath:   "/EncryptionConfig",
	}

	if input == nil {
		input = &GetEncryptionConfigInput{}
	}

	req := c.newRequest(op, input, &GetEncryptionConfigOutput{})
	return GetEncryptionConfigRequest{Request: req, Input: input, Copy: c.GetEncryptionConfigRequest}
}

// GetEncryptionConfigRequest is the request type for the
// GetEncryptionConfig API operation.
type GetEncryptionConfigRequest struct {
	*aws.Request
	Input *GetEncryptionConfigInput
	Copy  func(*GetEncryptionConfigInput) GetEncryptionConfigRequest
}

// Send marshals and sends the GetEncryptionConfig API request.
func (r GetEncryptionConfigRequest) Send(ctx context.Context) (*GetEncryptionConfigResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetEncryptionConfigResponse{
		GetEncryptionConfigOutput: r.Request.Data.(*GetEncryptionConfigOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetEncryptionConfigResponse is the response type for the
// GetEncryptionConfig API operation.
type GetEncryptionConfigResponse struct {
	*GetEncryptionConfigOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetEncryptionConfig request.
func (r *GetEncryptionConfigResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

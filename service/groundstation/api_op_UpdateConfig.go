// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package groundstation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/groundstation-2019-05-23/UpdateConfigRequest
type UpdateConfigInput struct {
	_ struct{} `type:"structure"`

	// Object containing the parameters for a Config.
	//
	// See the subtype definitions for what each type of Config contains.
	//
	// ConfigData is a required field
	ConfigData *ConfigTypeData `locationName:"configData" type:"structure" required:"true"`

	// ConfigId is a required field
	ConfigId *string `location:"uri" locationName:"configId" type:"string" required:"true"`

	// ConfigType is a required field
	ConfigType ConfigCapabilityType `location:"uri" locationName:"configType" type:"string" required:"true" enum:"true"`

	// Name is a required field
	Name *string `locationName:"name" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateConfigInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateConfigInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateConfigInput"}

	if s.ConfigData == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigData"))
	}

	if s.ConfigId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigId"))
	}
	if len(s.ConfigType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ConfigType"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if s.ConfigData != nil {
		if err := s.ConfigData.Validate(); err != nil {
			invalidParams.AddNested("ConfigData", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateConfigInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ConfigData != nil {
		v := s.ConfigData

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "configData", v, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ConfigId != nil {
		v := *s.ConfigId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "configId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.ConfigType) > 0 {
		v := s.ConfigType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "configType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/groundstation-2019-05-23/ConfigIdResponse
type UpdateConfigOutput struct {
	_ struct{} `type:"structure"`

	ConfigArn *string `locationName:"configArn" type:"string"`

	ConfigId *string `locationName:"configId" type:"string"`

	ConfigType ConfigCapabilityType `locationName:"configType" type:"string" enum:"true"`
}

// String returns the string representation
func (s UpdateConfigOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateConfigOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ConfigArn != nil {
		v := *s.ConfigArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "configArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ConfigId != nil {
		v := *s.ConfigId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "configId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.ConfigType) > 0 {
		v := s.ConfigType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "configType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

const opUpdateConfig = "UpdateConfig"

// UpdateConfigRequest returns a request value for making API operation for
// AWS Ground Station.
//
// Updates the Config used when scheduling contacts.
//
// Updating a Config will not update the execution parameters for existing future
// contacts scheduled with this Config.
//
//    // Example sending a request using UpdateConfigRequest.
//    req := client.UpdateConfigRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/groundstation-2019-05-23/UpdateConfig
func (c *Client) UpdateConfigRequest(input *UpdateConfigInput) UpdateConfigRequest {
	op := &aws.Operation{
		Name:       opUpdateConfig,
		HTTPMethod: "PUT",
		HTTPPath:   "/config/{configType}/{configId}",
	}

	if input == nil {
		input = &UpdateConfigInput{}
	}

	req := c.newRequest(op, input, &UpdateConfigOutput{})
	return UpdateConfigRequest{Request: req, Input: input, Copy: c.UpdateConfigRequest}
}

// UpdateConfigRequest is the request type for the
// UpdateConfig API operation.
type UpdateConfigRequest struct {
	*aws.Request
	Input *UpdateConfigInput
	Copy  func(*UpdateConfigInput) UpdateConfigRequest
}

// Send marshals and sends the UpdateConfig API request.
func (r UpdateConfigRequest) Send(ctx context.Context) (*UpdateConfigResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateConfigResponse{
		UpdateConfigOutput: r.Request.Data.(*UpdateConfigOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateConfigResponse is the response type for the
// UpdateConfig API operation.
type UpdateConfigResponse struct {
	*UpdateConfigOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateConfig request.
func (r *UpdateConfigResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

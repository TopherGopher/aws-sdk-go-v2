// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/PutVoiceConnectorOriginationRequest
type PutVoiceConnectorOriginationInput struct {
	_ struct{} `type:"structure"`

	// The origination setting details to add.
	//
	// Origination is a required field
	Origination *Origination `type:"structure" required:"true"`

	// The Amazon Chime Voice Connector ID.
	//
	// VoiceConnectorId is a required field
	VoiceConnectorId *string `location:"uri" locationName:"voiceConnectorId" type:"string" required:"true"`
}

// String returns the string representation
func (s PutVoiceConnectorOriginationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutVoiceConnectorOriginationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutVoiceConnectorOriginationInput"}

	if s.Origination == nil {
		invalidParams.Add(aws.NewErrParamRequired("Origination"))
	}

	if s.VoiceConnectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VoiceConnectorId"))
	}
	if s.Origination != nil {
		if err := s.Origination.Validate(); err != nil {
			invalidParams.AddNested("Origination", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutVoiceConnectorOriginationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Origination != nil {
		v := s.Origination

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Origination", v, metadata)
	}
	if s.VoiceConnectorId != nil {
		v := *s.VoiceConnectorId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "voiceConnectorId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/PutVoiceConnectorOriginationResponse
type PutVoiceConnectorOriginationOutput struct {
	_ struct{} `type:"structure"`

	// The updated origination setting details.
	Origination *Origination `type:"structure"`
}

// String returns the string representation
func (s PutVoiceConnectorOriginationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutVoiceConnectorOriginationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Origination != nil {
		v := s.Origination

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Origination", v, metadata)
	}
	return nil
}

const opPutVoiceConnectorOrigination = "PutVoiceConnectorOrigination"

// PutVoiceConnectorOriginationRequest returns a request value for making API operation for
// Amazon Chime.
//
// Adds origination settings for the specified Amazon Chime Voice Connector.
//
//    // Example sending a request using PutVoiceConnectorOriginationRequest.
//    req := client.PutVoiceConnectorOriginationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/PutVoiceConnectorOrigination
func (c *Client) PutVoiceConnectorOriginationRequest(input *PutVoiceConnectorOriginationInput) PutVoiceConnectorOriginationRequest {
	op := &aws.Operation{
		Name:       opPutVoiceConnectorOrigination,
		HTTPMethod: "PUT",
		HTTPPath:   "/voice-connectors/{voiceConnectorId}/origination",
	}

	if input == nil {
		input = &PutVoiceConnectorOriginationInput{}
	}

	req := c.newRequest(op, input, &PutVoiceConnectorOriginationOutput{})
	return PutVoiceConnectorOriginationRequest{Request: req, Input: input, Copy: c.PutVoiceConnectorOriginationRequest}
}

// PutVoiceConnectorOriginationRequest is the request type for the
// PutVoiceConnectorOrigination API operation.
type PutVoiceConnectorOriginationRequest struct {
	*aws.Request
	Input *PutVoiceConnectorOriginationInput
	Copy  func(*PutVoiceConnectorOriginationInput) PutVoiceConnectorOriginationRequest
}

// Send marshals and sends the PutVoiceConnectorOrigination API request.
func (r PutVoiceConnectorOriginationRequest) Send(ctx context.Context) (*PutVoiceConnectorOriginationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutVoiceConnectorOriginationResponse{
		PutVoiceConnectorOriginationOutput: r.Request.Data.(*PutVoiceConnectorOriginationOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutVoiceConnectorOriginationResponse is the response type for the
// PutVoiceConnectorOrigination API operation.
type PutVoiceConnectorOriginationResponse struct {
	*PutVoiceConnectorOriginationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutVoiceConnectorOrigination request.
func (r *PutVoiceConnectorOriginationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

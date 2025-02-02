// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/CreateFieldLevelEncryptionProfileRequest
type CreateFieldLevelEncryptionProfileInput struct {
	_ struct{} `type:"structure" payload:"FieldLevelEncryptionProfileConfig"`

	// The request to create a field-level encryption profile.
	//
	// FieldLevelEncryptionProfileConfig is a required field
	FieldLevelEncryptionProfileConfig *FieldLevelEncryptionProfileConfig `locationName:"FieldLevelEncryptionProfileConfig" type:"structure" required:"true" xmlURI:"http://cloudfront.amazonaws.com/doc/2019-03-26/"`
}

// String returns the string representation
func (s CreateFieldLevelEncryptionProfileInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateFieldLevelEncryptionProfileInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateFieldLevelEncryptionProfileInput"}

	if s.FieldLevelEncryptionProfileConfig == nil {
		invalidParams.Add(aws.NewErrParamRequired("FieldLevelEncryptionProfileConfig"))
	}
	if s.FieldLevelEncryptionProfileConfig != nil {
		if err := s.FieldLevelEncryptionProfileConfig.Validate(); err != nil {
			invalidParams.AddNested("FieldLevelEncryptionProfileConfig", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateFieldLevelEncryptionProfileInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.FieldLevelEncryptionProfileConfig != nil {
		v := s.FieldLevelEncryptionProfileConfig

		metadata := protocol.Metadata{XMLNamespaceURI: "http://cloudfront.amazonaws.com/doc/2019-03-26/"}
		e.SetFields(protocol.PayloadTarget, "FieldLevelEncryptionProfileConfig", v, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/CreateFieldLevelEncryptionProfileResult
type CreateFieldLevelEncryptionProfileOutput struct {
	_ struct{} `type:"structure" payload:"FieldLevelEncryptionProfile"`

	// The current version of the field level encryption profile. For example: E2QWRUHAPOMQZL.
	ETag *string `location:"header" locationName:"ETag" type:"string"`

	// Returned when you create a new field-level encryption profile.
	FieldLevelEncryptionProfile *FieldLevelEncryptionProfile `type:"structure"`

	// The fully qualified URI of the new profile resource just created. For example:
	// https://cloudfront.amazonaws.com/2010-11-01/field-level-encryption-profile/EDFDVBD632BHDS5.
	Location *string `location:"header" locationName:"Location" type:"string"`
}

// String returns the string representation
func (s CreateFieldLevelEncryptionProfileOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateFieldLevelEncryptionProfileOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ETag != nil {
		v := *s.ETag

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "ETag", protocol.StringValue(v), metadata)
	}
	if s.Location != nil {
		v := *s.Location

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Location", protocol.StringValue(v), metadata)
	}
	if s.FieldLevelEncryptionProfile != nil {
		v := s.FieldLevelEncryptionProfile

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "FieldLevelEncryptionProfile", v, metadata)
	}
	return nil
}

const opCreateFieldLevelEncryptionProfile = "CreateFieldLevelEncryptionProfile2019_03_26"

// CreateFieldLevelEncryptionProfileRequest returns a request value for making API operation for
// Amazon CloudFront.
//
// Create a field-level encryption profile.
//
//    // Example sending a request using CreateFieldLevelEncryptionProfileRequest.
//    req := client.CreateFieldLevelEncryptionProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/CreateFieldLevelEncryptionProfile
func (c *Client) CreateFieldLevelEncryptionProfileRequest(input *CreateFieldLevelEncryptionProfileInput) CreateFieldLevelEncryptionProfileRequest {
	op := &aws.Operation{
		Name:       opCreateFieldLevelEncryptionProfile,
		HTTPMethod: "POST",
		HTTPPath:   "/2019-03-26/field-level-encryption-profile",
	}

	if input == nil {
		input = &CreateFieldLevelEncryptionProfileInput{}
	}

	req := c.newRequest(op, input, &CreateFieldLevelEncryptionProfileOutput{})
	return CreateFieldLevelEncryptionProfileRequest{Request: req, Input: input, Copy: c.CreateFieldLevelEncryptionProfileRequest}
}

// CreateFieldLevelEncryptionProfileRequest is the request type for the
// CreateFieldLevelEncryptionProfile API operation.
type CreateFieldLevelEncryptionProfileRequest struct {
	*aws.Request
	Input *CreateFieldLevelEncryptionProfileInput
	Copy  func(*CreateFieldLevelEncryptionProfileInput) CreateFieldLevelEncryptionProfileRequest
}

// Send marshals and sends the CreateFieldLevelEncryptionProfile API request.
func (r CreateFieldLevelEncryptionProfileRequest) Send(ctx context.Context) (*CreateFieldLevelEncryptionProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateFieldLevelEncryptionProfileResponse{
		CreateFieldLevelEncryptionProfileOutput: r.Request.Data.(*CreateFieldLevelEncryptionProfileOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateFieldLevelEncryptionProfileResponse is the response type for the
// CreateFieldLevelEncryptionProfile API operation.
type CreateFieldLevelEncryptionProfileResponse struct {
	*CreateFieldLevelEncryptionProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateFieldLevelEncryptionProfile request.
func (r *CreateFieldLevelEncryptionProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

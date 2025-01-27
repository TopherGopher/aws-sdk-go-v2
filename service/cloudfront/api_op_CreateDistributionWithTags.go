// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The request to create a new distribution with tags.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/CreateDistributionWithTagsRequest
type CreateDistributionWithTagsInput struct {
	_ struct{} `type:"structure" payload:"DistributionConfigWithTags"`

	// The distribution's configuration information.
	//
	// DistributionConfigWithTags is a required field
	DistributionConfigWithTags *DistributionConfigWithTags `locationName:"DistributionConfigWithTags" type:"structure" required:"true" xmlURI:"http://cloudfront.amazonaws.com/doc/2019-03-26/"`
}

// String returns the string representation
func (s CreateDistributionWithTagsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDistributionWithTagsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDistributionWithTagsInput"}

	if s.DistributionConfigWithTags == nil {
		invalidParams.Add(aws.NewErrParamRequired("DistributionConfigWithTags"))
	}
	if s.DistributionConfigWithTags != nil {
		if err := s.DistributionConfigWithTags.Validate(); err != nil {
			invalidParams.AddNested("DistributionConfigWithTags", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateDistributionWithTagsInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.DistributionConfigWithTags != nil {
		v := s.DistributionConfigWithTags

		metadata := protocol.Metadata{XMLNamespaceURI: "http://cloudfront.amazonaws.com/doc/2019-03-26/"}
		e.SetFields(protocol.PayloadTarget, "DistributionConfigWithTags", v, metadata)
	}
	return nil
}

// The returned result of the corresponding request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/CreateDistributionWithTagsResult
type CreateDistributionWithTagsOutput struct {
	_ struct{} `type:"structure" payload:"Distribution"`

	// The distribution's information.
	Distribution *Distribution `type:"structure"`

	// The current version of the distribution created.
	ETag *string `location:"header" locationName:"ETag" type:"string"`

	// The fully qualified URI of the new distribution resource just created. For
	// example: https://cloudfront.amazonaws.com/2010-11-01/distribution/EDFDVBD632BHDS5.
	Location *string `location:"header" locationName:"Location" type:"string"`
}

// String returns the string representation
func (s CreateDistributionWithTagsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateDistributionWithTagsOutput) MarshalFields(e protocol.FieldEncoder) error {
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
	if s.Distribution != nil {
		v := s.Distribution

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "Distribution", v, metadata)
	}
	return nil
}

const opCreateDistributionWithTags = "CreateDistributionWithTags2019_03_26"

// CreateDistributionWithTagsRequest returns a request value for making API operation for
// Amazon CloudFront.
//
// Create a new distribution with tags.
//
//    // Example sending a request using CreateDistributionWithTagsRequest.
//    req := client.CreateDistributionWithTagsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/CreateDistributionWithTags
func (c *Client) CreateDistributionWithTagsRequest(input *CreateDistributionWithTagsInput) CreateDistributionWithTagsRequest {
	op := &aws.Operation{
		Name:       opCreateDistributionWithTags,
		HTTPMethod: "POST",
		HTTPPath:   "/2019-03-26/distribution?WithTags",
	}

	if input == nil {
		input = &CreateDistributionWithTagsInput{}
	}

	req := c.newRequest(op, input, &CreateDistributionWithTagsOutput{})
	return CreateDistributionWithTagsRequest{Request: req, Input: input, Copy: c.CreateDistributionWithTagsRequest}
}

// CreateDistributionWithTagsRequest is the request type for the
// CreateDistributionWithTags API operation.
type CreateDistributionWithTagsRequest struct {
	*aws.Request
	Input *CreateDistributionWithTagsInput
	Copy  func(*CreateDistributionWithTagsInput) CreateDistributionWithTagsRequest
}

// Send marshals and sends the CreateDistributionWithTags API request.
func (r CreateDistributionWithTagsRequest) Send(ctx context.Context) (*CreateDistributionWithTagsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDistributionWithTagsResponse{
		CreateDistributionWithTagsOutput: r.Request.Data.(*CreateDistributionWithTagsOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDistributionWithTagsResponse is the response type for the
// CreateDistributionWithTags API operation.
type CreateDistributionWithTagsResponse struct {
	*CreateDistributionWithTagsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDistributionWithTags request.
func (r *CreateDistributionWithTagsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

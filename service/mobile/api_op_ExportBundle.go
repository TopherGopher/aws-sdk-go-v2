// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mobile

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Request structure used to request generation of custom SDK and tool packages
// required to integrate mobile web or app clients with backed AWS resources.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mobile-2017-07-01/ExportBundleRequest
type ExportBundleInput struct {
	_ struct{} `type:"structure"`

	// Unique bundle identifier.
	//
	// BundleId is a required field
	BundleId *string `location:"uri" locationName:"bundleId" type:"string" required:"true"`

	// Developer desktop or target application platform.
	Platform Platform `location:"querystring" locationName:"platform" type:"string" enum:"true"`

	// Unique project identifier.
	ProjectId *string `location:"querystring" locationName:"projectId" type:"string"`
}

// String returns the string representation
func (s ExportBundleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ExportBundleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ExportBundleInput"}

	if s.BundleId == nil {
		invalidParams.Add(aws.NewErrParamRequired("BundleId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ExportBundleInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.BundleId != nil {
		v := *s.BundleId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "bundleId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Platform) > 0 {
		v := s.Platform

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "platform", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.ProjectId != nil {
		v := *s.ProjectId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "projectId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Result structure which contains link to download custom-generated SDK and
// tool packages used to integrate mobile web or app clients with backed AWS
// resources.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mobile-2017-07-01/ExportBundleResult
type ExportBundleOutput struct {
	_ struct{} `type:"structure"`

	// URL which contains the custom-generated SDK and tool packages used to integrate
	// the client mobile app or web app with the AWS resources created by the AWS
	// Mobile Hub project.
	DownloadUrl *string `locationName:"downloadUrl" type:"string"`
}

// String returns the string representation
func (s ExportBundleOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ExportBundleOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DownloadUrl != nil {
		v := *s.DownloadUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "downloadUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opExportBundle = "ExportBundle"

// ExportBundleRequest returns a request value for making API operation for
// AWS Mobile.
//
// Generates customized software development kit (SDK) and or tool packages
// used to integrate mobile web or mobile app clients with backend AWS resources.
//
//    // Example sending a request using ExportBundleRequest.
//    req := client.ExportBundleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mobile-2017-07-01/ExportBundle
func (c *Client) ExportBundleRequest(input *ExportBundleInput) ExportBundleRequest {
	op := &aws.Operation{
		Name:       opExportBundle,
		HTTPMethod: "POST",
		HTTPPath:   "/bundles/{bundleId}",
	}

	if input == nil {
		input = &ExportBundleInput{}
	}

	req := c.newRequest(op, input, &ExportBundleOutput{})
	return ExportBundleRequest{Request: req, Input: input, Copy: c.ExportBundleRequest}
}

// ExportBundleRequest is the request type for the
// ExportBundle API operation.
type ExportBundleRequest struct {
	*aws.Request
	Input *ExportBundleInput
	Copy  func(*ExportBundleInput) ExportBundleRequest
}

// Send marshals and sends the ExportBundle API request.
func (r ExportBundleRequest) Send(ctx context.Context) (*ExportBundleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ExportBundleResponse{
		ExportBundleOutput: r.Request.Data.(*ExportBundleOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ExportBundleResponse is the response type for the
// ExportBundle API operation.
type ExportBundleResponse struct {
	*ExportBundleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ExportBundle request.
func (r *ExportBundleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

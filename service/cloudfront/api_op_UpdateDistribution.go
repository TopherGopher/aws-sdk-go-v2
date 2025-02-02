// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The request to update a distribution.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/UpdateDistributionRequest
type UpdateDistributionInput struct {
	_ struct{} `type:"structure" payload:"DistributionConfig"`

	// The distribution's configuration information.
	//
	// DistributionConfig is a required field
	DistributionConfig *DistributionConfig `locationName:"DistributionConfig" type:"structure" required:"true" xmlURI:"http://cloudfront.amazonaws.com/doc/2019-03-26/"`

	// The distribution's id.
	//
	// Id is a required field
	Id *string `location:"uri" locationName:"Id" type:"string" required:"true"`

	// The value of the ETag header that you received when retrieving the distribution's
	// configuration. For example: E2QWRUHAPOMQZL.
	IfMatch *string `location:"header" locationName:"If-Match" type:"string"`
}

// String returns the string representation
func (s UpdateDistributionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateDistributionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateDistributionInput"}

	if s.DistributionConfig == nil {
		invalidParams.Add(aws.NewErrParamRequired("DistributionConfig"))
	}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.DistributionConfig != nil {
		if err := s.DistributionConfig.Validate(); err != nil {
			invalidParams.AddNested("DistributionConfig", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateDistributionInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.IfMatch != nil {
		v := *s.IfMatch

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "If-Match", protocol.StringValue(v), metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Id", protocol.StringValue(v), metadata)
	}
	if s.DistributionConfig != nil {
		v := s.DistributionConfig

		metadata := protocol.Metadata{XMLNamespaceURI: "http://cloudfront.amazonaws.com/doc/2019-03-26/"}
		e.SetFields(protocol.PayloadTarget, "DistributionConfig", v, metadata)
	}
	return nil
}

// The returned result of the corresponding request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/UpdateDistributionResult
type UpdateDistributionOutput struct {
	_ struct{} `type:"structure" payload:"Distribution"`

	// The distribution's information.
	Distribution *Distribution `type:"structure"`

	// The current version of the configuration. For example: E2QWRUHAPOMQZL.
	ETag *string `location:"header" locationName:"ETag" type:"string"`
}

// String returns the string representation
func (s UpdateDistributionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateDistributionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ETag != nil {
		v := *s.ETag

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "ETag", protocol.StringValue(v), metadata)
	}
	if s.Distribution != nil {
		v := s.Distribution

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "Distribution", v, metadata)
	}
	return nil
}

const opUpdateDistribution = "UpdateDistribution2019_03_26"

// UpdateDistributionRequest returns a request value for making API operation for
// Amazon CloudFront.
//
// Updates the configuration for a web distribution.
//
// When you update a distribution, there are more required fields than when
// you create a distribution. When you update your distribution by using this
// API action, follow the steps here to get the current configuration and then
// make your updates, to make sure that you include all of the required fields.
// To view a summary, see Required Fields for Create Distribution and Update
// Distribution (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-overview-required-fields.html)
// in the Amazon CloudFront Developer Guide.
//
// The update process includes getting the current distribution configuration,
// updating the XML document that is returned to make your changes, and then
// submitting an UpdateDistribution request to make the updates.
//
// For information about updating a distribution using the CloudFront console
// instead, see Creating a Distribution (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-creating-console.html)
// in the Amazon CloudFront Developer Guide.
//
// To update a web distribution using the CloudFront API
//
// Submit a GetDistributionConfig (https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_GetDistributionConfig.html)
// request to get the current configuration and an Etag header for the distribution.
//
// If you update the distribution again, you must get a new Etag header.
//
// Update the XML document that was returned in the response to your GetDistributionConfig
// request to include your changes.
//
// When you edit the XML file, be aware of the following:
//
//    * You must strip out the ETag parameter that is returned.
//
//    * Additional fields are required when you update a distribution. There
//    may be fields included in the XML file for features that you haven't configured
//    for your distribution. This is expected and required to successfully update
//    the distribution.
//
//    * You can't change the value of CallerReference. If you try to change
//    this value, CloudFront returns an IllegalUpdate error.
//
//    * The new configuration replaces the existing configuration; the values
//    that you specify in an UpdateDistribution request are not merged into
//    your existing configuration. When you add, delete, or replace values in
//    an element that allows multiple values (for example, CNAME), you must
//    specify all of the values that you want to appear in the updated distribution.
//    In addition, you must update the corresponding Quantity element.
//
// Submit an UpdateDistribution request to update the configuration for your
// distribution:
//
//    * In the request body, include the XML document that you updated in Step
//    2. The request body must include an XML document with a DistributionConfig
//    element.
//
//    * Set the value of the HTTP If-Match header to the value of the ETag header
//    that CloudFront returned when you submitted the GetDistributionConfig
//    request in Step 1.
//
// Review the response to the UpdateDistribution request to confirm that the
// configuration was successfully updated.
//
// Optional: Submit a GetDistribution (https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_GetDistribution.html)
// request to confirm that your changes have propagated. When propagation is
// complete, the value of Status is Deployed.
//
//    // Example sending a request using UpdateDistributionRequest.
//    req := client.UpdateDistributionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/UpdateDistribution
func (c *Client) UpdateDistributionRequest(input *UpdateDistributionInput) UpdateDistributionRequest {
	op := &aws.Operation{
		Name:       opUpdateDistribution,
		HTTPMethod: "PUT",
		HTTPPath:   "/2019-03-26/distribution/{Id}/config",
	}

	if input == nil {
		input = &UpdateDistributionInput{}
	}

	req := c.newRequest(op, input, &UpdateDistributionOutput{})
	return UpdateDistributionRequest{Request: req, Input: input, Copy: c.UpdateDistributionRequest}
}

// UpdateDistributionRequest is the request type for the
// UpdateDistribution API operation.
type UpdateDistributionRequest struct {
	*aws.Request
	Input *UpdateDistributionInput
	Copy  func(*UpdateDistributionInput) UpdateDistributionRequest
}

// Send marshals and sends the UpdateDistribution API request.
func (r UpdateDistributionRequest) Send(ctx context.Context) (*UpdateDistributionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateDistributionResponse{
		UpdateDistributionOutput: r.Request.Data.(*UpdateDistributionOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateDistributionResponse is the response type for the
// UpdateDistribution API operation.
type UpdateDistributionResponse struct {
	*UpdateDistributionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateDistribution request.
func (r *UpdateDistributionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to delete open and click tracking options in a configuration
// set.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/DeleteConfigurationSetTrackingOptionsRequest
type DeleteConfigurationSetTrackingOptionsInput struct {
	_ struct{} `type:"structure"`

	// The name of the configuration set from which you want to delete the tracking
	// options.
	//
	// ConfigurationSetName is a required field
	ConfigurationSetName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteConfigurationSetTrackingOptionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteConfigurationSetTrackingOptionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteConfigurationSetTrackingOptionsInput"}

	if s.ConfigurationSetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigurationSetName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// An empty element returned on a successful request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/DeleteConfigurationSetTrackingOptionsResponse
type DeleteConfigurationSetTrackingOptionsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteConfigurationSetTrackingOptionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteConfigurationSetTrackingOptions = "DeleteConfigurationSetTrackingOptions"

// DeleteConfigurationSetTrackingOptionsRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Deletes an association between a configuration set and a custom domain for
// open and click event tracking.
//
// By default, images and links used for tracking open and click events are
// hosted on domains operated by Amazon SES. You can configure a subdomain of
// your own to handle these events. For information about using custom domains,
// see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/configure-custom-open-click-domains.html).
//
// Deleting this kind of association will result in emails sent using the specified
// configuration set to capture open and click events using the standard, Amazon
// SES-operated domains.
//
//    // Example sending a request using DeleteConfigurationSetTrackingOptionsRequest.
//    req := client.DeleteConfigurationSetTrackingOptionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/DeleteConfigurationSetTrackingOptions
func (c *Client) DeleteConfigurationSetTrackingOptionsRequest(input *DeleteConfigurationSetTrackingOptionsInput) DeleteConfigurationSetTrackingOptionsRequest {
	op := &aws.Operation{
		Name:       opDeleteConfigurationSetTrackingOptions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteConfigurationSetTrackingOptionsInput{}
	}

	req := c.newRequest(op, input, &DeleteConfigurationSetTrackingOptionsOutput{})
	return DeleteConfigurationSetTrackingOptionsRequest{Request: req, Input: input, Copy: c.DeleteConfigurationSetTrackingOptionsRequest}
}

// DeleteConfigurationSetTrackingOptionsRequest is the request type for the
// DeleteConfigurationSetTrackingOptions API operation.
type DeleteConfigurationSetTrackingOptionsRequest struct {
	*aws.Request
	Input *DeleteConfigurationSetTrackingOptionsInput
	Copy  func(*DeleteConfigurationSetTrackingOptionsInput) DeleteConfigurationSetTrackingOptionsRequest
}

// Send marshals and sends the DeleteConfigurationSetTrackingOptions API request.
func (r DeleteConfigurationSetTrackingOptionsRequest) Send(ctx context.Context) (*DeleteConfigurationSetTrackingOptionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteConfigurationSetTrackingOptionsResponse{
		DeleteConfigurationSetTrackingOptionsOutput: r.Request.Data.(*DeleteConfigurationSetTrackingOptionsOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteConfigurationSetTrackingOptionsResponse is the response type for the
// DeleteConfigurationSetTrackingOptions API operation.
type DeleteConfigurationSetTrackingOptionsResponse struct {
	*DeleteConfigurationSetTrackingOptionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteConfigurationSetTrackingOptions request.
func (r *DeleteConfigurationSetTrackingOptionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

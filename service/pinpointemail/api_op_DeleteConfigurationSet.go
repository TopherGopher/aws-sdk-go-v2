// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpointemail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A request to delete a configuration set.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/DeleteConfigurationSetRequest
type DeleteConfigurationSetInput struct {
	_ struct{} `type:"structure"`

	// The name of the configuration set that you want to delete.
	//
	// ConfigurationSetName is a required field
	ConfigurationSetName *string `location:"uri" locationName:"ConfigurationSetName" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteConfigurationSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteConfigurationSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteConfigurationSetInput"}

	if s.ConfigurationSetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigurationSetName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteConfigurationSetInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ConfigurationSetName != nil {
		v := *s.ConfigurationSetName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ConfigurationSetName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// An HTTP 200 response if the request succeeds, or an error message if the
// request fails.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/DeleteConfigurationSetResponse
type DeleteConfigurationSetOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteConfigurationSetOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteConfigurationSetOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteConfigurationSet = "DeleteConfigurationSet"

// DeleteConfigurationSetRequest returns a request value for making API operation for
// Amazon Pinpoint Email Service.
//
// Delete an existing configuration set.
//
// In Amazon Pinpoint, configuration sets are groups of rules that you can apply
// to the emails you send. You apply a configuration set to an email by including
// a reference to the configuration set in the headers of the email. When you
// apply a configuration set to an email, all of the rules in that configuration
// set are applied to the email.
//
//    // Example sending a request using DeleteConfigurationSetRequest.
//    req := client.DeleteConfigurationSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/DeleteConfigurationSet
func (c *Client) DeleteConfigurationSetRequest(input *DeleteConfigurationSetInput) DeleteConfigurationSetRequest {
	op := &aws.Operation{
		Name:       opDeleteConfigurationSet,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v1/email/configuration-sets/{ConfigurationSetName}",
	}

	if input == nil {
		input = &DeleteConfigurationSetInput{}
	}

	req := c.newRequest(op, input, &DeleteConfigurationSetOutput{})
	return DeleteConfigurationSetRequest{Request: req, Input: input, Copy: c.DeleteConfigurationSetRequest}
}

// DeleteConfigurationSetRequest is the request type for the
// DeleteConfigurationSet API operation.
type DeleteConfigurationSetRequest struct {
	*aws.Request
	Input *DeleteConfigurationSetInput
	Copy  func(*DeleteConfigurationSetInput) DeleteConfigurationSetRequest
}

// Send marshals and sends the DeleteConfigurationSet API request.
func (r DeleteConfigurationSetRequest) Send(ctx context.Context) (*DeleteConfigurationSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteConfigurationSetResponse{
		DeleteConfigurationSetOutput: r.Request.Data.(*DeleteConfigurationSetOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteConfigurationSetResponse is the response type for the
// DeleteConfigurationSet API operation.
type DeleteConfigurationSetResponse struct {
	*DeleteConfigurationSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteConfigurationSet request.
func (r *DeleteConfigurationSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

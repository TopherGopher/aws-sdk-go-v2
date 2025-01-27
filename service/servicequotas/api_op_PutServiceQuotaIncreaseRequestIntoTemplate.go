// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicequotas

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/service-quotas-2019-06-24/PutServiceQuotaIncreaseRequestIntoTemplateRequest
type PutServiceQuotaIncreaseRequestIntoTemplateInput struct {
	_ struct{} `type:"structure"`

	// Specifies the AWS Region for the quota.
	//
	// AwsRegion is a required field
	AwsRegion *string `min:"1" type:"string" required:"true"`

	// Specifies the new, increased value for the quota.
	//
	// DesiredValue is a required field
	DesiredValue *float64 `type:"double" required:"true"`

	// Specifies the service quota that you want to use.
	//
	// QuotaCode is a required field
	QuotaCode *string `min:"1" type:"string" required:"true"`

	// Specifies the service that you want to use.
	//
	// ServiceCode is a required field
	ServiceCode *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s PutServiceQuotaIncreaseRequestIntoTemplateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutServiceQuotaIncreaseRequestIntoTemplateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutServiceQuotaIncreaseRequestIntoTemplateInput"}

	if s.AwsRegion == nil {
		invalidParams.Add(aws.NewErrParamRequired("AwsRegion"))
	}
	if s.AwsRegion != nil && len(*s.AwsRegion) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AwsRegion", 1))
	}

	if s.DesiredValue == nil {
		invalidParams.Add(aws.NewErrParamRequired("DesiredValue"))
	}

	if s.QuotaCode == nil {
		invalidParams.Add(aws.NewErrParamRequired("QuotaCode"))
	}
	if s.QuotaCode != nil && len(*s.QuotaCode) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("QuotaCode", 1))
	}

	if s.ServiceCode == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServiceCode"))
	}
	if s.ServiceCode != nil && len(*s.ServiceCode) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ServiceCode", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/service-quotas-2019-06-24/PutServiceQuotaIncreaseRequestIntoTemplateResponse
type PutServiceQuotaIncreaseRequestIntoTemplateOutput struct {
	_ struct{} `type:"structure"`

	// A structure that contains information about one service quota increase request.
	ServiceQuotaIncreaseRequestInTemplate *ServiceQuotaIncreaseRequestInTemplate `type:"structure"`
}

// String returns the string representation
func (s PutServiceQuotaIncreaseRequestIntoTemplateOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutServiceQuotaIncreaseRequestIntoTemplate = "PutServiceQuotaIncreaseRequestIntoTemplate"

// PutServiceQuotaIncreaseRequestIntoTemplateRequest returns a request value for making API operation for
// Service Quotas.
//
// Defines and adds a quota to the service quota template. To add a quota to
// the template, you must provide the ServiceCode, QuotaCode, AwsRegion, and
// DesiredValue. Once you add a quota to the template, use ListServiceQuotaIncreaseRequestsInTemplate
// to see the list of quotas in the template.
//
//    // Example sending a request using PutServiceQuotaIncreaseRequestIntoTemplateRequest.
//    req := client.PutServiceQuotaIncreaseRequestIntoTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/service-quotas-2019-06-24/PutServiceQuotaIncreaseRequestIntoTemplate
func (c *Client) PutServiceQuotaIncreaseRequestIntoTemplateRequest(input *PutServiceQuotaIncreaseRequestIntoTemplateInput) PutServiceQuotaIncreaseRequestIntoTemplateRequest {
	op := &aws.Operation{
		Name:       opPutServiceQuotaIncreaseRequestIntoTemplate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutServiceQuotaIncreaseRequestIntoTemplateInput{}
	}

	req := c.newRequest(op, input, &PutServiceQuotaIncreaseRequestIntoTemplateOutput{})
	return PutServiceQuotaIncreaseRequestIntoTemplateRequest{Request: req, Input: input, Copy: c.PutServiceQuotaIncreaseRequestIntoTemplateRequest}
}

// PutServiceQuotaIncreaseRequestIntoTemplateRequest is the request type for the
// PutServiceQuotaIncreaseRequestIntoTemplate API operation.
type PutServiceQuotaIncreaseRequestIntoTemplateRequest struct {
	*aws.Request
	Input *PutServiceQuotaIncreaseRequestIntoTemplateInput
	Copy  func(*PutServiceQuotaIncreaseRequestIntoTemplateInput) PutServiceQuotaIncreaseRequestIntoTemplateRequest
}

// Send marshals and sends the PutServiceQuotaIncreaseRequestIntoTemplate API request.
func (r PutServiceQuotaIncreaseRequestIntoTemplateRequest) Send(ctx context.Context) (*PutServiceQuotaIncreaseRequestIntoTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutServiceQuotaIncreaseRequestIntoTemplateResponse{
		PutServiceQuotaIncreaseRequestIntoTemplateOutput: r.Request.Data.(*PutServiceQuotaIncreaseRequestIntoTemplateOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutServiceQuotaIncreaseRequestIntoTemplateResponse is the response type for the
// PutServiceQuotaIncreaseRequestIntoTemplate API operation.
type PutServiceQuotaIncreaseRequestIntoTemplateResponse struct {
	*PutServiceQuotaIncreaseRequestIntoTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutServiceQuotaIncreaseRequestIntoTemplate request.
func (r *PutServiceQuotaIncreaseRequestIntoTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

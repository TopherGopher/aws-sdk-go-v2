// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/PutDataCatalogEncryptionSettingsRequest
type PutDataCatalogEncryptionSettingsInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Data Catalog to set the security configuration for. If none
	// is provided, the AWS account ID is used by default.
	CatalogId *string `min:"1" type:"string"`

	// The security configuration to set.
	//
	// DataCatalogEncryptionSettings is a required field
	DataCatalogEncryptionSettings *DataCatalogEncryptionSettings `type:"structure" required:"true"`
}

// String returns the string representation
func (s PutDataCatalogEncryptionSettingsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutDataCatalogEncryptionSettingsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutDataCatalogEncryptionSettingsInput"}
	if s.CatalogId != nil && len(*s.CatalogId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CatalogId", 1))
	}

	if s.DataCatalogEncryptionSettings == nil {
		invalidParams.Add(aws.NewErrParamRequired("DataCatalogEncryptionSettings"))
	}
	if s.DataCatalogEncryptionSettings != nil {
		if err := s.DataCatalogEncryptionSettings.Validate(); err != nil {
			invalidParams.AddNested("DataCatalogEncryptionSettings", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/PutDataCatalogEncryptionSettingsResponse
type PutDataCatalogEncryptionSettingsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutDataCatalogEncryptionSettingsOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutDataCatalogEncryptionSettings = "PutDataCatalogEncryptionSettings"

// PutDataCatalogEncryptionSettingsRequest returns a request value for making API operation for
// AWS Glue.
//
// Sets the security configuration for a specified catalog. After the configuration
// has been set, the specified encryption is applied to every catalog write
// thereafter.
//
//    // Example sending a request using PutDataCatalogEncryptionSettingsRequest.
//    req := client.PutDataCatalogEncryptionSettingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/PutDataCatalogEncryptionSettings
func (c *Client) PutDataCatalogEncryptionSettingsRequest(input *PutDataCatalogEncryptionSettingsInput) PutDataCatalogEncryptionSettingsRequest {
	op := &aws.Operation{
		Name:       opPutDataCatalogEncryptionSettings,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutDataCatalogEncryptionSettingsInput{}
	}

	req := c.newRequest(op, input, &PutDataCatalogEncryptionSettingsOutput{})
	return PutDataCatalogEncryptionSettingsRequest{Request: req, Input: input, Copy: c.PutDataCatalogEncryptionSettingsRequest}
}

// PutDataCatalogEncryptionSettingsRequest is the request type for the
// PutDataCatalogEncryptionSettings API operation.
type PutDataCatalogEncryptionSettingsRequest struct {
	*aws.Request
	Input *PutDataCatalogEncryptionSettingsInput
	Copy  func(*PutDataCatalogEncryptionSettingsInput) PutDataCatalogEncryptionSettingsRequest
}

// Send marshals and sends the PutDataCatalogEncryptionSettings API request.
func (r PutDataCatalogEncryptionSettingsRequest) Send(ctx context.Context) (*PutDataCatalogEncryptionSettingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutDataCatalogEncryptionSettingsResponse{
		PutDataCatalogEncryptionSettingsOutput: r.Request.Data.(*PutDataCatalogEncryptionSettingsOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutDataCatalogEncryptionSettingsResponse is the response type for the
// PutDataCatalogEncryptionSettings API operation.
type PutDataCatalogEncryptionSettingsResponse struct {
	*PutDataCatalogEncryptionSettingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutDataCatalogEncryptionSettings request.
func (r *PutDataCatalogEncryptionSettingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

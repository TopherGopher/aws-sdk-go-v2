// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecr-2015-09-21/PutImageTagMutabilityRequest
type PutImageTagMutabilityInput struct {
	_ struct{} `type:"structure"`

	// The tag mutability setting for the repository. If MUTABLE is specified, image
	// tags can be overwritten. If IMMUTABLE is specified, all image tags within
	// the repository will be immutable which will prevent them from being overwritten.
	//
	// ImageTagMutability is a required field
	ImageTagMutability ImageTagMutability `locationName:"imageTagMutability" type:"string" required:"true" enum:"true"`

	// The AWS account ID associated with the registry that contains the repository
	// in which to update the image tag mutability settings. If you do not specify
	// a registry, the default registry is assumed.
	RegistryId *string `locationName:"registryId" type:"string"`

	// The name of the repository in which to update the image tag mutability settings.
	//
	// RepositoryName is a required field
	RepositoryName *string `locationName:"repositoryName" min:"2" type:"string" required:"true"`
}

// String returns the string representation
func (s PutImageTagMutabilityInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutImageTagMutabilityInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutImageTagMutabilityInput"}
	if len(s.ImageTagMutability) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ImageTagMutability"))
	}

	if s.RepositoryName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RepositoryName"))
	}
	if s.RepositoryName != nil && len(*s.RepositoryName) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("RepositoryName", 2))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecr-2015-09-21/PutImageTagMutabilityResponse
type PutImageTagMutabilityOutput struct {
	_ struct{} `type:"structure"`

	// The image tag mutability setting for the repository.
	ImageTagMutability ImageTagMutability `locationName:"imageTagMutability" type:"string" enum:"true"`

	// The registry ID associated with the request.
	RegistryId *string `locationName:"registryId" type:"string"`

	// The repository name associated with the request.
	RepositoryName *string `locationName:"repositoryName" min:"2" type:"string"`
}

// String returns the string representation
func (s PutImageTagMutabilityOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutImageTagMutability = "PutImageTagMutability"

// PutImageTagMutabilityRequest returns a request value for making API operation for
// Amazon EC2 Container Registry.
//
// Updates the image tag mutability settings for a repository.
//
//    // Example sending a request using PutImageTagMutabilityRequest.
//    req := client.PutImageTagMutabilityRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecr-2015-09-21/PutImageTagMutability
func (c *Client) PutImageTagMutabilityRequest(input *PutImageTagMutabilityInput) PutImageTagMutabilityRequest {
	op := &aws.Operation{
		Name:       opPutImageTagMutability,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutImageTagMutabilityInput{}
	}

	req := c.newRequest(op, input, &PutImageTagMutabilityOutput{})
	return PutImageTagMutabilityRequest{Request: req, Input: input, Copy: c.PutImageTagMutabilityRequest}
}

// PutImageTagMutabilityRequest is the request type for the
// PutImageTagMutability API operation.
type PutImageTagMutabilityRequest struct {
	*aws.Request
	Input *PutImageTagMutabilityInput
	Copy  func(*PutImageTagMutabilityInput) PutImageTagMutabilityRequest
}

// Send marshals and sends the PutImageTagMutability API request.
func (r PutImageTagMutabilityRequest) Send(ctx context.Context) (*PutImageTagMutabilityResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutImageTagMutabilityResponse{
		PutImageTagMutabilityOutput: r.Request.Data.(*PutImageTagMutabilityOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutImageTagMutabilityResponse is the response type for the
// PutImageTagMutability API operation.
type PutImageTagMutabilityResponse struct {
	*PutImageTagMutabilityOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutImageTagMutability request.
func (r *PutImageTagMutabilityResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

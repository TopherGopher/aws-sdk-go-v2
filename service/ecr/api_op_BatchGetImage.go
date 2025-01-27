// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecr

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecr-2015-09-21/BatchGetImageRequest
type BatchGetImageInput struct {
	_ struct{} `type:"structure"`

	// The accepted media types for the request.
	//
	// Valid values: application/vnd.docker.distribution.manifest.v1+json | application/vnd.docker.distribution.manifest.v2+json
	// | application/vnd.oci.image.manifest.v1+json
	AcceptedMediaTypes []string `locationName:"acceptedMediaTypes" min:"1" type:"list"`

	// A list of image ID references that correspond to images to describe. The
	// format of the imageIds reference is imageTag=tag or imageDigest=digest.
	//
	// ImageIds is a required field
	ImageIds []ImageIdentifier `locationName:"imageIds" min:"1" type:"list" required:"true"`

	// The AWS account ID associated with the registry that contains the images
	// to describe. If you do not specify a registry, the default registry is assumed.
	RegistryId *string `locationName:"registryId" type:"string"`

	// The repository that contains the images to describe.
	//
	// RepositoryName is a required field
	RepositoryName *string `locationName:"repositoryName" min:"2" type:"string" required:"true"`
}

// String returns the string representation
func (s BatchGetImageInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchGetImageInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchGetImageInput"}
	if s.AcceptedMediaTypes != nil && len(s.AcceptedMediaTypes) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AcceptedMediaTypes", 1))
	}

	if s.ImageIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("ImageIds"))
	}
	if s.ImageIds != nil && len(s.ImageIds) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ImageIds", 1))
	}

	if s.RepositoryName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RepositoryName"))
	}
	if s.RepositoryName != nil && len(*s.RepositoryName) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("RepositoryName", 2))
	}
	if s.ImageIds != nil {
		for i, v := range s.ImageIds {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ImageIds", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecr-2015-09-21/BatchGetImageResponse
type BatchGetImageOutput struct {
	_ struct{} `type:"structure"`

	// Any failures associated with the call.
	Failures []ImageFailure `locationName:"failures" type:"list"`

	// A list of image objects corresponding to the image references in the request.
	Images []Image `locationName:"images" type:"list"`
}

// String returns the string representation
func (s BatchGetImageOutput) String() string {
	return awsutil.Prettify(s)
}

const opBatchGetImage = "BatchGetImage"

// BatchGetImageRequest returns a request value for making API operation for
// Amazon EC2 Container Registry.
//
// Gets detailed information for specified images within a specified repository.
// Images are specified with either imageTag or imageDigest.
//
//    // Example sending a request using BatchGetImageRequest.
//    req := client.BatchGetImageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecr-2015-09-21/BatchGetImage
func (c *Client) BatchGetImageRequest(input *BatchGetImageInput) BatchGetImageRequest {
	op := &aws.Operation{
		Name:       opBatchGetImage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &BatchGetImageInput{}
	}

	req := c.newRequest(op, input, &BatchGetImageOutput{})
	return BatchGetImageRequest{Request: req, Input: input, Copy: c.BatchGetImageRequest}
}

// BatchGetImageRequest is the request type for the
// BatchGetImage API operation.
type BatchGetImageRequest struct {
	*aws.Request
	Input *BatchGetImageInput
	Copy  func(*BatchGetImageInput) BatchGetImageRequest
}

// Send marshals and sends the BatchGetImage API request.
func (r BatchGetImageRequest) Send(ctx context.Context) (*BatchGetImageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchGetImageResponse{
		BatchGetImageOutput: r.Request.Data.(*BatchGetImageOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchGetImageResponse is the response type for the
// BatchGetImage API operation.
type BatchGetImageResponse struct {
	*BatchGetImageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchGetImage request.
func (r *BatchGetImageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

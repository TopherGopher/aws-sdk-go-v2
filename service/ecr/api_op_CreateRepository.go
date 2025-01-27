// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecr-2015-09-21/CreateRepositoryRequest
type CreateRepositoryInput struct {
	_ struct{} `type:"structure"`

	// The tag mutability setting for the repository. If this parameter is omitted,
	// the default setting of MUTABLE will be used which will allow image tags to
	// be overwritten. If IMMUTABLE is specified, all image tags within the repository
	// will be immutable which will prevent them from being overwritten.
	ImageTagMutability ImageTagMutability `locationName:"imageTagMutability" type:"string" enum:"true"`

	// The name to use for the repository. The repository name may be specified
	// on its own (such as nginx-web-app) or it can be prepended with a namespace
	// to group the repository into a category (such as project-a/nginx-web-app).
	//
	// RepositoryName is a required field
	RepositoryName *string `locationName:"repositoryName" min:"2" type:"string" required:"true"`

	// The metadata that you apply to the repository to help you categorize and
	// organize them. Each tag consists of a key and an optional value, both of
	// which you define. Tag keys can have a maximum character length of 128 characters,
	// and tag values can have a maximum length of 256 characters.
	Tags []Tag `locationName:"tags" type:"list"`
}

// String returns the string representation
func (s CreateRepositoryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateRepositoryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateRepositoryInput"}

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

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecr-2015-09-21/CreateRepositoryResponse
type CreateRepositoryOutput struct {
	_ struct{} `type:"structure"`

	// The repository that was created.
	Repository *Repository `locationName:"repository" type:"structure"`
}

// String returns the string representation
func (s CreateRepositoryOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateRepository = "CreateRepository"

// CreateRepositoryRequest returns a request value for making API operation for
// Amazon EC2 Container Registry.
//
// Creates an image repository.
//
//    // Example sending a request using CreateRepositoryRequest.
//    req := client.CreateRepositoryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecr-2015-09-21/CreateRepository
func (c *Client) CreateRepositoryRequest(input *CreateRepositoryInput) CreateRepositoryRequest {
	op := &aws.Operation{
		Name:       opCreateRepository,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateRepositoryInput{}
	}

	req := c.newRequest(op, input, &CreateRepositoryOutput{})
	return CreateRepositoryRequest{Request: req, Input: input, Copy: c.CreateRepositoryRequest}
}

// CreateRepositoryRequest is the request type for the
// CreateRepository API operation.
type CreateRepositoryRequest struct {
	*aws.Request
	Input *CreateRepositoryInput
	Copy  func(*CreateRepositoryInput) CreateRepositoryRequest
}

// Send marshals and sends the CreateRepository API request.
func (r CreateRepositoryRequest) Send(ctx context.Context) (*CreateRepositoryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateRepositoryResponse{
		CreateRepositoryOutput: r.Request.Data.(*CreateRepositoryOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateRepositoryResponse is the response type for the
// CreateRepository API operation.
type CreateRepositoryResponse struct {
	*CreateRepositoryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateRepository request.
func (r *CreateRepositoryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

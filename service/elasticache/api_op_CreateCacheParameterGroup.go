// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a CreateCacheParameterGroup operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/CreateCacheParameterGroupMessage
type CreateCacheParameterGroupInput struct {
	_ struct{} `type:"structure"`

	// The name of the cache parameter group family that the cache parameter group
	// can be used with.
	//
	// Valid values are: memcached1.4 | memcached1.5 | redis2.6 | redis2.8 | redis3.2
	// | redis4.0 | redis5.0 |
	//
	// CacheParameterGroupFamily is a required field
	CacheParameterGroupFamily *string `type:"string" required:"true"`

	// A user-specified name for the cache parameter group.
	//
	// CacheParameterGroupName is a required field
	CacheParameterGroupName *string `type:"string" required:"true"`

	// A user-specified description for the cache parameter group.
	//
	// Description is a required field
	Description *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateCacheParameterGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateCacheParameterGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateCacheParameterGroupInput"}

	if s.CacheParameterGroupFamily == nil {
		invalidParams.Add(aws.NewErrParamRequired("CacheParameterGroupFamily"))
	}

	if s.CacheParameterGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("CacheParameterGroupName"))
	}

	if s.Description == nil {
		invalidParams.Add(aws.NewErrParamRequired("Description"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/CreateCacheParameterGroupResult
type CreateCacheParameterGroupOutput struct {
	_ struct{} `type:"structure"`

	// Represents the output of a CreateCacheParameterGroup operation.
	CacheParameterGroup *CacheParameterGroup `type:"structure"`
}

// String returns the string representation
func (s CreateCacheParameterGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateCacheParameterGroup = "CreateCacheParameterGroup"

// CreateCacheParameterGroupRequest returns a request value for making API operation for
// Amazon ElastiCache.
//
// Creates a new Amazon ElastiCache cache parameter group. An ElastiCache cache
// parameter group is a collection of parameters and their values that are applied
// to all of the nodes in any cluster or replication group using the CacheParameterGroup.
//
// A newly created CacheParameterGroup is an exact duplicate of the default
// parameter group for the CacheParameterGroupFamily. To customize the newly
// created CacheParameterGroup you can change the values of specific parameters.
// For more information, see:
//
//    * ModifyCacheParameterGroup (https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_ModifyCacheParameterGroup.html)
//    in the ElastiCache API Reference.
//
//    * Parameters and Parameter Groups (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/ParameterGroups.html)
//    in the ElastiCache User Guide.
//
//    // Example sending a request using CreateCacheParameterGroupRequest.
//    req := client.CreateCacheParameterGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/CreateCacheParameterGroup
func (c *Client) CreateCacheParameterGroupRequest(input *CreateCacheParameterGroupInput) CreateCacheParameterGroupRequest {
	op := &aws.Operation{
		Name:       opCreateCacheParameterGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateCacheParameterGroupInput{}
	}

	req := c.newRequest(op, input, &CreateCacheParameterGroupOutput{})
	return CreateCacheParameterGroupRequest{Request: req, Input: input, Copy: c.CreateCacheParameterGroupRequest}
}

// CreateCacheParameterGroupRequest is the request type for the
// CreateCacheParameterGroup API operation.
type CreateCacheParameterGroupRequest struct {
	*aws.Request
	Input *CreateCacheParameterGroupInput
	Copy  func(*CreateCacheParameterGroupInput) CreateCacheParameterGroupRequest
}

// Send marshals and sends the CreateCacheParameterGroup API request.
func (r CreateCacheParameterGroupRequest) Send(ctx context.Context) (*CreateCacheParameterGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateCacheParameterGroupResponse{
		CreateCacheParameterGroupOutput: r.Request.Data.(*CreateCacheParameterGroupOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateCacheParameterGroupResponse is the response type for the
// CreateCacheParameterGroup API operation.
type CreateCacheParameterGroupResponse struct {
	*CreateCacheParameterGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateCacheParameterGroup request.
func (r *CreateCacheParameterGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/ModifyGlobalClusterMessage
type ModifyGlobalClusterInput struct {
	_ struct{} `type:"structure"`

	// Indicates if the global database cluster has deletion protection enabled.
	// The global database cluster can't be deleted when deletion protection is
	// enabled.
	DeletionProtection *bool `type:"boolean"`

	// The DB cluster identifier for the global cluster being modified. This parameter
	// is not case-sensitive.
	//
	// Constraints:
	//
	//    * Must match the identifier of an existing global database cluster.
	GlobalClusterIdentifier *string `type:"string"`

	// The new cluster identifier for the global database cluster when modifying
	// a global database cluster. This value is stored as a lowercase string.
	//
	// Constraints:
	//
	//    * Must contain from 1 to 63 letters, numbers, or hyphens
	//
	//    * The first character must be a letter
	//
	//    * Can't end with a hyphen or contain two consecutive hyphens
	//
	// Example: my-cluster2
	NewGlobalClusterIdentifier *string `type:"string"`
}

// String returns the string representation
func (s ModifyGlobalClusterInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/ModifyGlobalClusterResult
type ModifyGlobalClusterOutput struct {
	_ struct{} `type:"structure"`

	// A data type representing an Aurora global database.
	GlobalCluster *GlobalCluster `type:"structure"`
}

// String returns the string representation
func (s ModifyGlobalClusterOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyGlobalCluster = "ModifyGlobalCluster"

// ModifyGlobalClusterRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Modify a setting for an Amazon Aurora global cluster. You can change one
// or more database configuration parameters by specifying these parameters
// and the new values in the request. For more information on Amazon Aurora,
// see What Is Amazon Aurora? (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
// in the Amazon Aurora User Guide.
//
// This action only applies to Aurora DB clusters.
//
//    // Example sending a request using ModifyGlobalClusterRequest.
//    req := client.ModifyGlobalClusterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/ModifyGlobalCluster
func (c *Client) ModifyGlobalClusterRequest(input *ModifyGlobalClusterInput) ModifyGlobalClusterRequest {
	op := &aws.Operation{
		Name:       opModifyGlobalCluster,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyGlobalClusterInput{}
	}

	req := c.newRequest(op, input, &ModifyGlobalClusterOutput{})
	return ModifyGlobalClusterRequest{Request: req, Input: input, Copy: c.ModifyGlobalClusterRequest}
}

// ModifyGlobalClusterRequest is the request type for the
// ModifyGlobalCluster API operation.
type ModifyGlobalClusterRequest struct {
	*aws.Request
	Input *ModifyGlobalClusterInput
	Copy  func(*ModifyGlobalClusterInput) ModifyGlobalClusterRequest
}

// Send marshals and sends the ModifyGlobalCluster API request.
func (r ModifyGlobalClusterRequest) Send(ctx context.Context) (*ModifyGlobalClusterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyGlobalClusterResponse{
		ModifyGlobalClusterOutput: r.Request.Data.(*ModifyGlobalClusterOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyGlobalClusterResponse is the response type for the
// ModifyGlobalCluster API operation.
type ModifyGlobalClusterResponse struct {
	*ModifyGlobalClusterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyGlobalCluster request.
func (r *ModifyGlobalClusterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

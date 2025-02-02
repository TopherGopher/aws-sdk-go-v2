// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kafka

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/kafka-2018-11-14/DescribeClusterRequest
type DescribeClusterInput struct {
	_ struct{} `type:"structure"`

	// ClusterArn is a required field
	ClusterArn *string `location:"uri" locationName:"clusterArn" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeClusterInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeClusterInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeClusterInput"}

	if s.ClusterArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeClusterInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ClusterArn != nil {
		v := *s.ClusterArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "clusterArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Returns information about a cluster.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kafka-2018-11-14/DescribeClusterResponse
type DescribeClusterOutput struct {
	_ struct{} `type:"structure"`

	// The cluster information.
	ClusterInfo *ClusterInfo `locationName:"clusterInfo" type:"structure"`
}

// String returns the string representation
func (s DescribeClusterOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeClusterOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ClusterInfo != nil {
		v := s.ClusterInfo

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "clusterInfo", v, metadata)
	}
	return nil
}

const opDescribeCluster = "DescribeCluster"

// DescribeClusterRequest returns a request value for making API operation for
// Managed Streaming for Kafka.
//
// Returns a description of the MSK cluster whose Amazon Resource Name (ARN)
// is specified in the request.
//
//    // Example sending a request using DescribeClusterRequest.
//    req := client.DescribeClusterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kafka-2018-11-14/DescribeCluster
func (c *Client) DescribeClusterRequest(input *DescribeClusterInput) DescribeClusterRequest {
	op := &aws.Operation{
		Name:       opDescribeCluster,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/clusters/{clusterArn}",
	}

	if input == nil {
		input = &DescribeClusterInput{}
	}

	req := c.newRequest(op, input, &DescribeClusterOutput{})
	return DescribeClusterRequest{Request: req, Input: input, Copy: c.DescribeClusterRequest}
}

// DescribeClusterRequest is the request type for the
// DescribeCluster API operation.
type DescribeClusterRequest struct {
	*aws.Request
	Input *DescribeClusterInput
	Copy  func(*DescribeClusterInput) DescribeClusterRequest
}

// Send marshals and sends the DescribeCluster API request.
func (r DescribeClusterRequest) Send(ctx context.Context) (*DescribeClusterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeClusterResponse{
		DescribeClusterOutput: r.Request.Data.(*DescribeClusterOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeClusterResponse is the response type for the
// DescribeCluster API operation.
type DescribeClusterResponse struct {
	*DescribeClusterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeCluster request.
func (r *DescribeClusterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

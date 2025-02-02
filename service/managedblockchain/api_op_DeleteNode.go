// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package managedblockchain

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/managedblockchain-2018-09-24/DeleteNodeInput
type DeleteNodeInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier of the member that owns this node.
	//
	// MemberId is a required field
	MemberId *string `location:"uri" locationName:"memberId" min:"1" type:"string" required:"true"`

	// The unique identifier of the network that the node belongs to.
	//
	// NetworkId is a required field
	NetworkId *string `location:"uri" locationName:"networkId" min:"1" type:"string" required:"true"`

	// The unique identifier of the node.
	//
	// NodeId is a required field
	NodeId *string `location:"uri" locationName:"nodeId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteNodeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteNodeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteNodeInput"}

	if s.MemberId == nil {
		invalidParams.Add(aws.NewErrParamRequired("MemberId"))
	}
	if s.MemberId != nil && len(*s.MemberId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MemberId", 1))
	}

	if s.NetworkId == nil {
		invalidParams.Add(aws.NewErrParamRequired("NetworkId"))
	}
	if s.NetworkId != nil && len(*s.NetworkId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NetworkId", 1))
	}

	if s.NodeId == nil {
		invalidParams.Add(aws.NewErrParamRequired("NodeId"))
	}
	if s.NodeId != nil && len(*s.NodeId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NodeId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteNodeInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MemberId != nil {
		v := *s.MemberId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "memberId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NetworkId != nil {
		v := *s.NetworkId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "networkId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NodeId != nil {
		v := *s.NodeId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "nodeId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/managedblockchain-2018-09-24/DeleteNodeOutput
type DeleteNodeOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteNodeOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteNodeOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteNode = "DeleteNode"

// DeleteNodeRequest returns a request value for making API operation for
// Amazon Managed Blockchain.
//
// Deletes a peer node from a member that your AWS account owns. All data on
// the node is lost and cannot be recovered.
//
//    // Example sending a request using DeleteNodeRequest.
//    req := client.DeleteNodeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/managedblockchain-2018-09-24/DeleteNode
func (c *Client) DeleteNodeRequest(input *DeleteNodeInput) DeleteNodeRequest {
	op := &aws.Operation{
		Name:       opDeleteNode,
		HTTPMethod: "DELETE",
		HTTPPath:   "/networks/{networkId}/members/{memberId}/nodes/{nodeId}",
	}

	if input == nil {
		input = &DeleteNodeInput{}
	}

	req := c.newRequest(op, input, &DeleteNodeOutput{})
	return DeleteNodeRequest{Request: req, Input: input, Copy: c.DeleteNodeRequest}
}

// DeleteNodeRequest is the request type for the
// DeleteNode API operation.
type DeleteNodeRequest struct {
	*aws.Request
	Input *DeleteNodeInput
	Copy  func(*DeleteNodeInput) DeleteNodeRequest
}

// Send marshals and sends the DeleteNode API request.
func (r DeleteNodeRequest) Send(ctx context.Context) (*DeleteNodeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteNodeResponse{
		DeleteNodeOutput: r.Request.Data.(*DeleteNodeOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteNodeResponse is the response type for the
// DeleteNode API operation.
type DeleteNodeResponse struct {
	*DeleteNodeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteNode request.
func (r *DeleteNodeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

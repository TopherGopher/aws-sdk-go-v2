// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package managedblockchain

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/managedblockchain-2018-09-24/CreateProposalInput
type CreateProposalInput struct {
	_ struct{} `type:"structure"`

	// The type of actions proposed, such as inviting a member or removing a member.
	// The types of Actions in a proposal are mutually exclusive. For example, a
	// proposal with Invitations actions cannot also contain Removals actions.
	//
	// Actions is a required field
	Actions *ProposalActions `type:"structure" required:"true"`

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the operation. An idempotent operation completes no more than one time.
	// This identifier is required only if you make a service request directly using
	// an HTTP client. It is generated automatically if you use an AWS SDK or the
	// AWS CLI.
	//
	// ClientRequestToken is a required field
	ClientRequestToken *string `min:"1" type:"string" required:"true" idempotencyToken:"true"`

	// A description for the proposal that is visible to voting members, for example,
	// "Proposal to add Example Corp. as member."
	Description *string `type:"string"`

	// The unique identifier of the member that is creating the proposal. This identifier
	// is especially useful for identifying the member making the proposal when
	// multiple members exist in a single AWS account.
	//
	// MemberId is a required field
	MemberId *string `min:"1" type:"string" required:"true"`

	// The unique identifier of the network for which the proposal is made.
	//
	// NetworkId is a required field
	NetworkId *string `location:"uri" locationName:"networkId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateProposalInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateProposalInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateProposalInput"}

	if s.Actions == nil {
		invalidParams.Add(aws.NewErrParamRequired("Actions"))
	}

	if s.ClientRequestToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientRequestToken"))
	}
	if s.ClientRequestToken != nil && len(*s.ClientRequestToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientRequestToken", 1))
	}

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
	if s.Actions != nil {
		if err := s.Actions.Validate(); err != nil {
			invalidParams.AddNested("Actions", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateProposalInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Actions != nil {
		v := s.Actions

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Actions", v, metadata)
	}
	var ClientRequestToken string
	if s.ClientRequestToken != nil {
		ClientRequestToken = *s.ClientRequestToken
	} else {
		ClientRequestToken = protocol.GetIdempotencyToken()
	}
	{
		v := ClientRequestToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ClientRequestToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MemberId != nil {
		v := *s.MemberId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MemberId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NetworkId != nil {
		v := *s.NetworkId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "networkId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/managedblockchain-2018-09-24/CreateProposalOutput
type CreateProposalOutput struct {
	_ struct{} `type:"structure"`

	// The unique identifier of the proposal.
	ProposalId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s CreateProposalOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateProposalOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ProposalId != nil {
		v := *s.ProposalId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ProposalId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateProposal = "CreateProposal"

// CreateProposalRequest returns a request value for making API operation for
// Amazon Managed Blockchain.
//
// Creates a proposal for a change to the network that other members of the
// network can vote on, for example, a proposal to add a new member to the network.
// Any member can create a proposal.
//
//    // Example sending a request using CreateProposalRequest.
//    req := client.CreateProposalRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/managedblockchain-2018-09-24/CreateProposal
func (c *Client) CreateProposalRequest(input *CreateProposalInput) CreateProposalRequest {
	op := &aws.Operation{
		Name:       opCreateProposal,
		HTTPMethod: "POST",
		HTTPPath:   "/networks/{networkId}/proposals",
	}

	if input == nil {
		input = &CreateProposalInput{}
	}

	req := c.newRequest(op, input, &CreateProposalOutput{})
	return CreateProposalRequest{Request: req, Input: input, Copy: c.CreateProposalRequest}
}

// CreateProposalRequest is the request type for the
// CreateProposal API operation.
type CreateProposalRequest struct {
	*aws.Request
	Input *CreateProposalInput
	Copy  func(*CreateProposalInput) CreateProposalRequest
}

// Send marshals and sends the CreateProposal API request.
func (r CreateProposalRequest) Send(ctx context.Context) (*CreateProposalResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateProposalResponse{
		CreateProposalOutput: r.Request.Data.(*CreateProposalOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateProposalResponse is the response type for the
// CreateProposal API operation.
type CreateProposalResponse struct {
	*CreateProposalOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateProposal request.
func (r *CreateProposalResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

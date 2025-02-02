// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package medialive

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A request to create actions (add actions to the schedule), delete actions
// (remove actions from the schedule), or both create and delete actions.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/medialive-2017-10-14/BatchUpdateScheduleRequest
type BatchUpdateScheduleInput struct {
	_ struct{} `type:"structure"`

	// ChannelId is a required field
	ChannelId *string `location:"uri" locationName:"channelId" type:"string" required:"true"`

	// Schedule actions to create in the schedule.
	Creates *BatchScheduleActionCreateRequest `locationName:"creates" type:"structure"`

	// Schedule actions to delete from the schedule.
	Deletes *BatchScheduleActionDeleteRequest `locationName:"deletes" type:"structure"`
}

// String returns the string representation
func (s BatchUpdateScheduleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchUpdateScheduleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchUpdateScheduleInput"}

	if s.ChannelId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ChannelId"))
	}
	if s.Creates != nil {
		if err := s.Creates.Validate(); err != nil {
			invalidParams.AddNested("Creates", err.(aws.ErrInvalidParams))
		}
	}
	if s.Deletes != nil {
		if err := s.Deletes.Validate(); err != nil {
			invalidParams.AddNested("Deletes", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s BatchUpdateScheduleInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Creates != nil {
		v := s.Creates

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "creates", v, metadata)
	}
	if s.Deletes != nil {
		v := s.Deletes

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "deletes", v, metadata)
	}
	if s.ChannelId != nil {
		v := *s.ChannelId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "channelId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/medialive-2017-10-14/BatchUpdateScheduleResponse
type BatchUpdateScheduleOutput struct {
	_ struct{} `type:"structure"`

	// List of actions that have been created in the schedule.
	Creates *BatchScheduleActionCreateResult `locationName:"creates" type:"structure"`

	// List of actions that have been deleted from the schedule.
	Deletes *BatchScheduleActionDeleteResult `locationName:"deletes" type:"structure"`
}

// String returns the string representation
func (s BatchUpdateScheduleOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s BatchUpdateScheduleOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Creates != nil {
		v := s.Creates

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "creates", v, metadata)
	}
	if s.Deletes != nil {
		v := s.Deletes

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "deletes", v, metadata)
	}
	return nil
}

const opBatchUpdateSchedule = "BatchUpdateSchedule"

// BatchUpdateScheduleRequest returns a request value for making API operation for
// AWS Elemental MediaLive.
//
// Update a channel schedule
//
//    // Example sending a request using BatchUpdateScheduleRequest.
//    req := client.BatchUpdateScheduleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/medialive-2017-10-14/BatchUpdateSchedule
func (c *Client) BatchUpdateScheduleRequest(input *BatchUpdateScheduleInput) BatchUpdateScheduleRequest {
	op := &aws.Operation{
		Name:       opBatchUpdateSchedule,
		HTTPMethod: "PUT",
		HTTPPath:   "/prod/channels/{channelId}/schedule",
	}

	if input == nil {
		input = &BatchUpdateScheduleInput{}
	}

	req := c.newRequest(op, input, &BatchUpdateScheduleOutput{})
	return BatchUpdateScheduleRequest{Request: req, Input: input, Copy: c.BatchUpdateScheduleRequest}
}

// BatchUpdateScheduleRequest is the request type for the
// BatchUpdateSchedule API operation.
type BatchUpdateScheduleRequest struct {
	*aws.Request
	Input *BatchUpdateScheduleInput
	Copy  func(*BatchUpdateScheduleInput) BatchUpdateScheduleRequest
}

// Send marshals and sends the BatchUpdateSchedule API request.
func (r BatchUpdateScheduleRequest) Send(ctx context.Context) (*BatchUpdateScheduleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchUpdateScheduleResponse{
		BatchUpdateScheduleOutput: r.Request.Data.(*BatchUpdateScheduleOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchUpdateScheduleResponse is the response type for the
// BatchUpdateSchedule API operation.
type BatchUpdateScheduleResponse struct {
	*BatchUpdateScheduleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchUpdateSchedule request.
func (r *BatchUpdateScheduleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

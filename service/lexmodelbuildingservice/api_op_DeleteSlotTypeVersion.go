// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lex-models-2017-04-19/DeleteSlotTypeVersionRequest
type DeleteSlotTypeVersionInput struct {
	_ struct{} `type:"structure"`

	// The name of the slot type.
	//
	// Name is a required field
	Name *string `location:"uri" locationName:"name" min:"1" type:"string" required:"true"`

	// The version of the slot type to delete. You cannot delete the $LATEST version
	// of the slot type. To delete the $LATEST version, use the DeleteSlotType operation.
	//
	// Version is a required field
	Version *string `location:"uri" locationName:"version" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteSlotTypeVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteSlotTypeVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteSlotTypeVersionInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if s.Version == nil {
		invalidParams.Add(aws.NewErrParamRequired("Version"))
	}
	if s.Version != nil && len(*s.Version) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Version", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteSlotTypeVersionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "version", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lex-models-2017-04-19/DeleteSlotTypeVersionOutput
type DeleteSlotTypeVersionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteSlotTypeVersionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteSlotTypeVersionOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteSlotTypeVersion = "DeleteSlotTypeVersion"

// DeleteSlotTypeVersionRequest returns a request value for making API operation for
// Amazon Lex Model Building Service.
//
// Deletes a specific version of a slot type. To delete all versions of a slot
// type, use the DeleteSlotType operation.
//
// This operation requires permissions for the lex:DeleteSlotTypeVersion action.
//
//    // Example sending a request using DeleteSlotTypeVersionRequest.
//    req := client.DeleteSlotTypeVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lex-models-2017-04-19/DeleteSlotTypeVersion
func (c *Client) DeleteSlotTypeVersionRequest(input *DeleteSlotTypeVersionInput) DeleteSlotTypeVersionRequest {
	op := &aws.Operation{
		Name:       opDeleteSlotTypeVersion,
		HTTPMethod: "DELETE",
		HTTPPath:   "/slottypes/{name}/version/{version}",
	}

	if input == nil {
		input = &DeleteSlotTypeVersionInput{}
	}

	req := c.newRequest(op, input, &DeleteSlotTypeVersionOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteSlotTypeVersionRequest{Request: req, Input: input, Copy: c.DeleteSlotTypeVersionRequest}
}

// DeleteSlotTypeVersionRequest is the request type for the
// DeleteSlotTypeVersion API operation.
type DeleteSlotTypeVersionRequest struct {
	*aws.Request
	Input *DeleteSlotTypeVersionInput
	Copy  func(*DeleteSlotTypeVersionInput) DeleteSlotTypeVersionRequest
}

// Send marshals and sends the DeleteSlotTypeVersion API request.
func (r DeleteSlotTypeVersionRequest) Send(ctx context.Context) (*DeleteSlotTypeVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteSlotTypeVersionResponse{
		DeleteSlotTypeVersionOutput: r.Request.Data.(*DeleteSlotTypeVersionOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteSlotTypeVersionResponse is the response type for the
// DeleteSlotTypeVersion API operation.
type DeleteSlotTypeVersionResponse struct {
	*DeleteSlotTypeVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteSlotTypeVersion request.
func (r *DeleteSlotTypeVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/DeleteLayerVersionRequest
type DeleteLayerVersionInput struct {
	_ struct{} `type:"structure"`

	// The name or Amazon Resource Name (ARN) of the layer.
	//
	// LayerName is a required field
	LayerName *string `location:"uri" locationName:"LayerName" min:"1" type:"string" required:"true"`

	// The version number.
	//
	// VersionNumber is a required field
	VersionNumber *int64 `location:"uri" locationName:"VersionNumber" type:"long" required:"true"`
}

// String returns the string representation
func (s DeleteLayerVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteLayerVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteLayerVersionInput"}

	if s.LayerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LayerName"))
	}
	if s.LayerName != nil && len(*s.LayerName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LayerName", 1))
	}

	if s.VersionNumber == nil {
		invalidParams.Add(aws.NewErrParamRequired("VersionNumber"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteLayerVersionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.LayerName != nil {
		v := *s.LayerName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "LayerName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VersionNumber != nil {
		v := *s.VersionNumber

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "VersionNumber", protocol.Int64Value(v), metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/DeleteLayerVersionOutput
type DeleteLayerVersionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteLayerVersionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteLayerVersionOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteLayerVersion = "DeleteLayerVersion"

// DeleteLayerVersionRequest returns a request value for making API operation for
// AWS Lambda.
//
// Deletes a version of an AWS Lambda layer (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html).
// Deleted versions can no longer be viewed or added to functions. To avoid
// breaking functions, a copy of the version remains in Lambda until no functions
// refer to it.
//
//    // Example sending a request using DeleteLayerVersionRequest.
//    req := client.DeleteLayerVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/DeleteLayerVersion
func (c *Client) DeleteLayerVersionRequest(input *DeleteLayerVersionInput) DeleteLayerVersionRequest {
	op := &aws.Operation{
		Name:       opDeleteLayerVersion,
		HTTPMethod: "DELETE",
		HTTPPath:   "/2018-10-31/layers/{LayerName}/versions/{VersionNumber}",
	}

	if input == nil {
		input = &DeleteLayerVersionInput{}
	}

	req := c.newRequest(op, input, &DeleteLayerVersionOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteLayerVersionRequest{Request: req, Input: input, Copy: c.DeleteLayerVersionRequest}
}

// DeleteLayerVersionRequest is the request type for the
// DeleteLayerVersion API operation.
type DeleteLayerVersionRequest struct {
	*aws.Request
	Input *DeleteLayerVersionInput
	Copy  func(*DeleteLayerVersionInput) DeleteLayerVersionRequest
}

// Send marshals and sends the DeleteLayerVersion API request.
func (r DeleteLayerVersionRequest) Send(ctx context.Context) (*DeleteLayerVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteLayerVersionResponse{
		DeleteLayerVersionOutput: r.Request.Data.(*DeleteLayerVersionOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteLayerVersionResponse is the response type for the
// DeleteLayerVersion API operation.
type DeleteLayerVersionResponse struct {
	*DeleteLayerVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteLayerVersion request.
func (r *DeleteLayerVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

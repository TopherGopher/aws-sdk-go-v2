// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/PublishSchemaRequest
type PublishSchemaInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) that is associated with the development schema.
	// For more information, see arns.
	//
	// DevelopmentSchemaArn is a required field
	DevelopmentSchemaArn *string `location:"header" locationName:"x-amz-data-partition" type:"string" required:"true"`

	// The minor version under which the schema will be published. This parameter
	// is recommended. Schemas have both a major and minor version associated with
	// them.
	MinorVersion *string `min:"1" type:"string"`

	// The new name under which the schema will be published. If this is not provided,
	// the development schema is considered.
	Name *string `min:"1" type:"string"`

	// The major version under which the schema will be published. Schemas have
	// both a major and minor version associated with them.
	//
	// Version is a required field
	Version *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s PublishSchemaInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PublishSchemaInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PublishSchemaInput"}

	if s.DevelopmentSchemaArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DevelopmentSchemaArn"))
	}
	if s.MinorVersion != nil && len(*s.MinorVersion) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MinorVersion", 1))
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
func (s PublishSchemaInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MinorVersion != nil {
		v := *s.MinorVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MinorVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Version", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DevelopmentSchemaArn != nil {
		v := *s.DevelopmentSchemaArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-data-partition", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/PublishSchemaResponse
type PublishSchemaOutput struct {
	_ struct{} `type:"structure"`

	// The ARN that is associated with the published schema. For more information,
	// see arns.
	PublishedSchemaArn *string `type:"string"`
}

// String returns the string representation
func (s PublishSchemaOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PublishSchemaOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.PublishedSchemaArn != nil {
		v := *s.PublishedSchemaArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "PublishedSchemaArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opPublishSchema = "PublishSchema"

// PublishSchemaRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Publishes a development schema with a major version and a recommended minor
// version.
//
//    // Example sending a request using PublishSchemaRequest.
//    req := client.PublishSchemaRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/PublishSchema
func (c *Client) PublishSchemaRequest(input *PublishSchemaInput) PublishSchemaRequest {
	op := &aws.Operation{
		Name:       opPublishSchema,
		HTTPMethod: "PUT",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/schema/publish",
	}

	if input == nil {
		input = &PublishSchemaInput{}
	}

	req := c.newRequest(op, input, &PublishSchemaOutput{})
	return PublishSchemaRequest{Request: req, Input: input, Copy: c.PublishSchemaRequest}
}

// PublishSchemaRequest is the request type for the
// PublishSchema API operation.
type PublishSchemaRequest struct {
	*aws.Request
	Input *PublishSchemaInput
	Copy  func(*PublishSchemaInput) PublishSchemaRequest
}

// Send marshals and sends the PublishSchema API request.
func (r PublishSchemaRequest) Send(ctx context.Context) (*PublishSchemaResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PublishSchemaResponse{
		PublishSchemaOutput: r.Request.Data.(*PublishSchemaOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PublishSchemaResponse is the response type for the
// PublishSchema API operation.
type PublishSchemaResponse struct {
	*PublishSchemaOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PublishSchema request.
func (r *PublishSchemaResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

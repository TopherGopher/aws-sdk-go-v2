// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigatewayv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/DeleteRouteResponseRequest
type DeleteRouteResponseInput struct {
	_ struct{} `type:"structure"`

	// ApiId is a required field
	ApiId *string `location:"uri" locationName:"apiId" type:"string" required:"true"`

	// RouteId is a required field
	RouteId *string `location:"uri" locationName:"routeId" type:"string" required:"true"`

	// RouteResponseId is a required field
	RouteResponseId *string `location:"uri" locationName:"routeResponseId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteRouteResponseInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteRouteResponseInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteRouteResponseInput"}

	if s.ApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiId"))
	}

	if s.RouteId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RouteId"))
	}

	if s.RouteResponseId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RouteResponseId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteRouteResponseInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApiId != nil {
		v := *s.ApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "apiId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RouteId != nil {
		v := *s.RouteId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "routeId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RouteResponseId != nil {
		v := *s.RouteResponseId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "routeResponseId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/DeleteRouteResponseOutput
type DeleteRouteResponseOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteRouteResponseOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteRouteResponseOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteRouteResponse = "DeleteRouteResponse"

// DeleteRouteResponseRequest returns a request value for making API operation for
// AmazonApiGatewayV2.
//
// Deletes a RouteResponse.
//
//    // Example sending a request using DeleteRouteResponseRequest.
//    req := client.DeleteRouteResponseRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/DeleteRouteResponse
func (c *Client) DeleteRouteResponseRequest(input *DeleteRouteResponseInput) DeleteRouteResponseRequest {
	op := &aws.Operation{
		Name:       opDeleteRouteResponse,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v2/apis/{apiId}/routes/{routeId}/routeresponses/{routeResponseId}",
	}

	if input == nil {
		input = &DeleteRouteResponseInput{}
	}

	req := c.newRequest(op, input, &DeleteRouteResponseOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteRouteResponseRequest{Request: req, Input: input, Copy: c.DeleteRouteResponseRequest}
}

// DeleteRouteResponseRequest is the request type for the
// DeleteRouteResponse API operation.
type DeleteRouteResponseRequest struct {
	*aws.Request
	Input *DeleteRouteResponseInput
	Copy  func(*DeleteRouteResponseInput) DeleteRouteResponseRequest
}

// Send marshals and sends the DeleteRouteResponse API request.
func (r DeleteRouteResponseRequest) Send(ctx context.Context) (*DeleteRouteResponseResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteRouteResponseResponse{
		DeleteRouteResponseOutput: r.Request.Data.(*DeleteRouteResponseOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteRouteResponseResponse is the response type for the
// DeleteRouteResponse API operation.
type DeleteRouteResponseResponse struct {
	*DeleteRouteResponseOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteRouteResponse request.
func (r *DeleteRouteResponseResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/DescribeEndpointTypesMessage
type DescribeEndpointTypesInput struct {
	_ struct{} `type:"structure"`

	// Filters applied to the describe action.
	//
	// Valid filter names: engine-name | endpoint-type
	Filters []Filter `type:"list"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to
	// the value specified by MaxRecords.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a pagination token called a marker
	// is included in the response so that the remaining results can be retrieved.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeEndpointTypesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeEndpointTypesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeEndpointTypesInput"}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/DescribeEndpointTypesResponse
type DescribeEndpointTypesOutput struct {
	_ struct{} `type:"structure"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to
	// the value specified by MaxRecords.
	Marker *string `type:"string"`

	// The types of endpoints that are supported.
	SupportedEndpointTypes []SupportedEndpointType `type:"list"`
}

// String returns the string representation
func (s DescribeEndpointTypesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeEndpointTypes = "DescribeEndpointTypes"

// DescribeEndpointTypesRequest returns a request value for making API operation for
// AWS Database Migration Service.
//
// Returns information about the type of endpoints available.
//
//    // Example sending a request using DescribeEndpointTypesRequest.
//    req := client.DescribeEndpointTypesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/DescribeEndpointTypes
func (c *Client) DescribeEndpointTypesRequest(input *DescribeEndpointTypesInput) DescribeEndpointTypesRequest {
	op := &aws.Operation{
		Name:       opDescribeEndpointTypes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxRecords",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeEndpointTypesInput{}
	}

	req := c.newRequest(op, input, &DescribeEndpointTypesOutput{})
	return DescribeEndpointTypesRequest{Request: req, Input: input, Copy: c.DescribeEndpointTypesRequest}
}

// DescribeEndpointTypesRequest is the request type for the
// DescribeEndpointTypes API operation.
type DescribeEndpointTypesRequest struct {
	*aws.Request
	Input *DescribeEndpointTypesInput
	Copy  func(*DescribeEndpointTypesInput) DescribeEndpointTypesRequest
}

// Send marshals and sends the DescribeEndpointTypes API request.
func (r DescribeEndpointTypesRequest) Send(ctx context.Context) (*DescribeEndpointTypesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeEndpointTypesResponse{
		DescribeEndpointTypesOutput: r.Request.Data.(*DescribeEndpointTypesOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeEndpointTypesRequestPaginator returns a paginator for DescribeEndpointTypes.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeEndpointTypesRequest(input)
//   p := databasemigrationservice.NewDescribeEndpointTypesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeEndpointTypesPaginator(req DescribeEndpointTypesRequest) DescribeEndpointTypesPaginator {
	return DescribeEndpointTypesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeEndpointTypesInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeEndpointTypesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeEndpointTypesPaginator struct {
	aws.Pager
}

func (p *DescribeEndpointTypesPaginator) CurrentPage() *DescribeEndpointTypesOutput {
	return p.Pager.CurrentPage().(*DescribeEndpointTypesOutput)
}

// DescribeEndpointTypesResponse is the response type for the
// DescribeEndpointTypes API operation.
type DescribeEndpointTypesResponse struct {
	*DescribeEndpointTypesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeEndpointTypes request.
func (r *DescribeEndpointTypesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/DescribeServiceUpdatesMessage
type DescribeServiceUpdatesInput struct {
	_ struct{} `type:"structure"`

	// An optional marker returned from a prior request. Use this marker for pagination
	// of results from this operation. If this parameter is specified, the response
	// includes only records beyond the marker, up to the value specified by MaxRecords.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response
	MaxRecords *int64 `type:"integer"`

	// The unique ID of the service update
	ServiceUpdateName *string `type:"string"`

	// The status of the service update
	ServiceUpdateStatus []ServiceUpdateStatus `type:"list"`
}

// String returns the string representation
func (s DescribeServiceUpdatesInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/ServiceUpdatesMessage
type DescribeServiceUpdatesOutput struct {
	_ struct{} `type:"structure"`

	// An optional marker returned from a prior request. Use this marker for pagination
	// of results from this operation. If this parameter is specified, the response
	// includes only records beyond the marker, up to the value specified by MaxRecords.
	Marker *string `type:"string"`

	// A list of service updates
	ServiceUpdates []ServiceUpdate `locationNameList:"ServiceUpdate" type:"list"`
}

// String returns the string representation
func (s DescribeServiceUpdatesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeServiceUpdates = "DescribeServiceUpdates"

// DescribeServiceUpdatesRequest returns a request value for making API operation for
// Amazon ElastiCache.
//
// Returns details of the service updates
//
//    // Example sending a request using DescribeServiceUpdatesRequest.
//    req := client.DescribeServiceUpdatesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/DescribeServiceUpdates
func (c *Client) DescribeServiceUpdatesRequest(input *DescribeServiceUpdatesInput) DescribeServiceUpdatesRequest {
	op := &aws.Operation{
		Name:       opDescribeServiceUpdates,
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
		input = &DescribeServiceUpdatesInput{}
	}

	req := c.newRequest(op, input, &DescribeServiceUpdatesOutput{})
	return DescribeServiceUpdatesRequest{Request: req, Input: input, Copy: c.DescribeServiceUpdatesRequest}
}

// DescribeServiceUpdatesRequest is the request type for the
// DescribeServiceUpdates API operation.
type DescribeServiceUpdatesRequest struct {
	*aws.Request
	Input *DescribeServiceUpdatesInput
	Copy  func(*DescribeServiceUpdatesInput) DescribeServiceUpdatesRequest
}

// Send marshals and sends the DescribeServiceUpdates API request.
func (r DescribeServiceUpdatesRequest) Send(ctx context.Context) (*DescribeServiceUpdatesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeServiceUpdatesResponse{
		DescribeServiceUpdatesOutput: r.Request.Data.(*DescribeServiceUpdatesOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeServiceUpdatesRequestPaginator returns a paginator for DescribeServiceUpdates.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeServiceUpdatesRequest(input)
//   p := elasticache.NewDescribeServiceUpdatesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeServiceUpdatesPaginator(req DescribeServiceUpdatesRequest) DescribeServiceUpdatesPaginator {
	return DescribeServiceUpdatesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeServiceUpdatesInput
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

// DescribeServiceUpdatesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeServiceUpdatesPaginator struct {
	aws.Pager
}

func (p *DescribeServiceUpdatesPaginator) CurrentPage() *DescribeServiceUpdatesOutput {
	return p.Pager.CurrentPage().(*DescribeServiceUpdatesOutput)
}

// DescribeServiceUpdatesResponse is the response type for the
// DescribeServiceUpdates API operation.
type DescribeServiceUpdatesResponse struct {
	*DescribeServiceUpdatesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeServiceUpdates request.
func (r *DescribeServiceUpdatesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

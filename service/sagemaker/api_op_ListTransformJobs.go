// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/ListTransformJobsRequest
type ListTransformJobsInput struct {
	_ struct{} `type:"structure"`

	// A filter that returns only transform jobs created after the specified time.
	CreationTimeAfter *time.Time `type:"timestamp"`

	// A filter that returns only transform jobs created before the specified time.
	CreationTimeBefore *time.Time `type:"timestamp"`

	// A filter that returns only transform jobs modified after the specified time.
	LastModifiedTimeAfter *time.Time `type:"timestamp"`

	// A filter that returns only transform jobs modified before the specified time.
	LastModifiedTimeBefore *time.Time `type:"timestamp"`

	// The maximum number of transform jobs to return in the response. The default
	// value is 10.
	MaxResults *int64 `min:"1" type:"integer"`

	// A string in the transform job name. This filter returns only transform jobs
	// whose name contains the specified string.
	NameContains *string `type:"string"`

	// If the result of the previous ListTransformJobs request was truncated, the
	// response includes a NextToken. To retrieve the next set of transform jobs,
	// use the token in the next request.
	NextToken *string `type:"string"`

	// The field to sort results by. The default is CreationTime.
	SortBy SortBy `type:"string" enum:"true"`

	// The sort order for results. The default is Descending.
	SortOrder SortOrder `type:"string" enum:"true"`

	// A filter that retrieves only transform jobs with a specific status.
	StatusEquals TransformJobStatus `type:"string" enum:"true"`
}

// String returns the string representation
func (s ListTransformJobsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTransformJobsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListTransformJobsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/ListTransformJobsResponse
type ListTransformJobsOutput struct {
	_ struct{} `type:"structure"`

	// If the response is truncated, Amazon SageMaker returns this token. To retrieve
	// the next set of transform jobs, use it in the next request.
	NextToken *string `type:"string"`

	// An array of TransformJobSummary objects.
	//
	// TransformJobSummaries is a required field
	TransformJobSummaries []TransformJobSummary `type:"list" required:"true"`
}

// String returns the string representation
func (s ListTransformJobsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListTransformJobs = "ListTransformJobs"

// ListTransformJobsRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Lists transform jobs.
//
//    // Example sending a request using ListTransformJobsRequest.
//    req := client.ListTransformJobsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/ListTransformJobs
func (c *Client) ListTransformJobsRequest(input *ListTransformJobsInput) ListTransformJobsRequest {
	op := &aws.Operation{
		Name:       opListTransformJobs,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListTransformJobsInput{}
	}

	req := c.newRequest(op, input, &ListTransformJobsOutput{})
	return ListTransformJobsRequest{Request: req, Input: input, Copy: c.ListTransformJobsRequest}
}

// ListTransformJobsRequest is the request type for the
// ListTransformJobs API operation.
type ListTransformJobsRequest struct {
	*aws.Request
	Input *ListTransformJobsInput
	Copy  func(*ListTransformJobsInput) ListTransformJobsRequest
}

// Send marshals and sends the ListTransformJobs API request.
func (r ListTransformJobsRequest) Send(ctx context.Context) (*ListTransformJobsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListTransformJobsResponse{
		ListTransformJobsOutput: r.Request.Data.(*ListTransformJobsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListTransformJobsRequestPaginator returns a paginator for ListTransformJobs.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListTransformJobsRequest(input)
//   p := sagemaker.NewListTransformJobsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListTransformJobsPaginator(req ListTransformJobsRequest) ListTransformJobsPaginator {
	return ListTransformJobsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListTransformJobsInput
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

// ListTransformJobsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListTransformJobsPaginator struct {
	aws.Pager
}

func (p *ListTransformJobsPaginator) CurrentPage() *ListTransformJobsOutput {
	return p.Pager.CurrentPage().(*ListTransformJobsOutput)
}

// ListTransformJobsResponse is the response type for the
// ListTransformJobs API operation.
type ListTransformJobsResponse struct {
	*ListTransformJobsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListTransformJobs request.
func (r *ListTransformJobsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

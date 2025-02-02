// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/ListCompilationJobsRequest
type ListCompilationJobsInput struct {
	_ struct{} `type:"structure"`

	// A filter that returns the model compilation jobs that were created after
	// a specified time.
	CreationTimeAfter *time.Time `type:"timestamp"`

	// A filter that returns the model compilation jobs that were created before
	// a specified time.
	CreationTimeBefore *time.Time `type:"timestamp"`

	// A filter that returns the model compilation jobs that were modified after
	// a specified time.
	LastModifiedTimeAfter *time.Time `type:"timestamp"`

	// A filter that returns the model compilation jobs that were modified before
	// a specified time.
	LastModifiedTimeBefore *time.Time `type:"timestamp"`

	// The maximum number of model compilation jobs to return in the response.
	MaxResults *int64 `min:"1" type:"integer"`

	// A filter that returns the model compilation jobs whose name contains a specified
	// string.
	NameContains *string `type:"string"`

	// If the result of the previous ListCompilationJobs request was truncated,
	// the response includes a NextToken. To retrieve the next set of model compilation
	// jobs, use the token in the next request.
	NextToken *string `type:"string"`

	// The field by which to sort results. The default is CreationTime.
	SortBy ListCompilationJobsSortBy `type:"string" enum:"true"`

	// The sort order for results. The default is Ascending.
	SortOrder SortOrder `type:"string" enum:"true"`

	// A filter that retrieves model compilation jobs with a specific DescribeCompilationJobResponse$CompilationJobStatus
	// status.
	StatusEquals CompilationJobStatus `type:"string" enum:"true"`
}

// String returns the string representation
func (s ListCompilationJobsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListCompilationJobsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListCompilationJobsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/ListCompilationJobsResponse
type ListCompilationJobsOutput struct {
	_ struct{} `type:"structure"`

	// An array of CompilationJobSummary objects, each describing a model compilation
	// job.
	//
	// CompilationJobSummaries is a required field
	CompilationJobSummaries []CompilationJobSummary `type:"list" required:"true"`

	// If the response is truncated, Amazon SageMaker returns this NextToken. To
	// retrieve the next set of model compilation jobs, use this token in the next
	// request.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListCompilationJobsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListCompilationJobs = "ListCompilationJobs"

// ListCompilationJobsRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Lists model compilation jobs that satisfy various filters.
//
// To create a model compilation job, use CreateCompilationJob. To get information
// about a particular model compilation job you have created, use DescribeCompilationJob.
//
//    // Example sending a request using ListCompilationJobsRequest.
//    req := client.ListCompilationJobsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/ListCompilationJobs
func (c *Client) ListCompilationJobsRequest(input *ListCompilationJobsInput) ListCompilationJobsRequest {
	op := &aws.Operation{
		Name:       opListCompilationJobs,
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
		input = &ListCompilationJobsInput{}
	}

	req := c.newRequest(op, input, &ListCompilationJobsOutput{})
	return ListCompilationJobsRequest{Request: req, Input: input, Copy: c.ListCompilationJobsRequest}
}

// ListCompilationJobsRequest is the request type for the
// ListCompilationJobs API operation.
type ListCompilationJobsRequest struct {
	*aws.Request
	Input *ListCompilationJobsInput
	Copy  func(*ListCompilationJobsInput) ListCompilationJobsRequest
}

// Send marshals and sends the ListCompilationJobs API request.
func (r ListCompilationJobsRequest) Send(ctx context.Context) (*ListCompilationJobsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListCompilationJobsResponse{
		ListCompilationJobsOutput: r.Request.Data.(*ListCompilationJobsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListCompilationJobsRequestPaginator returns a paginator for ListCompilationJobs.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListCompilationJobsRequest(input)
//   p := sagemaker.NewListCompilationJobsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListCompilationJobsPaginator(req ListCompilationJobsRequest) ListCompilationJobsPaginator {
	return ListCompilationJobsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListCompilationJobsInput
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

// ListCompilationJobsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListCompilationJobsPaginator struct {
	aws.Pager
}

func (p *ListCompilationJobsPaginator) CurrentPage() *ListCompilationJobsOutput {
	return p.Pager.CurrentPage().(*ListCompilationJobsOutput)
}

// ListCompilationJobsResponse is the response type for the
// ListCompilationJobs API operation.
type ListCompilationJobsResponse struct {
	*ListCompilationJobsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListCompilationJobs request.
func (r *ListCompilationJobsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

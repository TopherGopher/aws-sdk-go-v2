// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mturk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mturk-requester-2017-01-17/ListWorkersWithQualificationTypeRequest
type ListWorkersWithQualificationTypeInput struct {
	_ struct{} `type:"structure"`

	// Limit the number of results returned.
	MaxResults *int64 `min:"1" type:"integer"`

	// Pagination Token
	NextToken *string `min:"1" type:"string"`

	// The ID of the Qualification type of the Qualifications to return.
	//
	// QualificationTypeId is a required field
	QualificationTypeId *string `min:"1" type:"string" required:"true"`

	// The status of the Qualifications to return. Can be Granted | Revoked.
	Status QualificationStatus `type:"string" enum:"true"`
}

// String returns the string representation
func (s ListWorkersWithQualificationTypeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListWorkersWithQualificationTypeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListWorkersWithQualificationTypeInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if s.QualificationTypeId == nil {
		invalidParams.Add(aws.NewErrParamRequired("QualificationTypeId"))
	}
	if s.QualificationTypeId != nil && len(*s.QualificationTypeId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("QualificationTypeId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mturk-requester-2017-01-17/ListWorkersWithQualificationTypeResponse
type ListWorkersWithQualificationTypeOutput struct {
	_ struct{} `type:"structure"`

	// If the previous response was incomplete (because there is more data to retrieve),
	// Amazon Mechanical Turk returns a pagination token in the response. You can
	// use this pagination token to retrieve the next set of results.
	NextToken *string `min:"1" type:"string"`

	// The number of Qualifications on this page in the filtered results list, equivalent
	// to the number of Qualifications being returned by this call.
	NumResults *int64 `type:"integer"`

	// The list of Qualification elements returned by this call.
	Qualifications []Qualification `type:"list"`
}

// String returns the string representation
func (s ListWorkersWithQualificationTypeOutput) String() string {
	return awsutil.Prettify(s)
}

const opListWorkersWithQualificationType = "ListWorkersWithQualificationType"

// ListWorkersWithQualificationTypeRequest returns a request value for making API operation for
// Amazon Mechanical Turk.
//
// The ListWorkersWithQualificationType operation returns all of the Workers
// that have been associated with a given Qualification type.
//
//    // Example sending a request using ListWorkersWithQualificationTypeRequest.
//    req := client.ListWorkersWithQualificationTypeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mturk-requester-2017-01-17/ListWorkersWithQualificationType
func (c *Client) ListWorkersWithQualificationTypeRequest(input *ListWorkersWithQualificationTypeInput) ListWorkersWithQualificationTypeRequest {
	op := &aws.Operation{
		Name:       opListWorkersWithQualificationType,
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
		input = &ListWorkersWithQualificationTypeInput{}
	}

	req := c.newRequest(op, input, &ListWorkersWithQualificationTypeOutput{})
	return ListWorkersWithQualificationTypeRequest{Request: req, Input: input, Copy: c.ListWorkersWithQualificationTypeRequest}
}

// ListWorkersWithQualificationTypeRequest is the request type for the
// ListWorkersWithQualificationType API operation.
type ListWorkersWithQualificationTypeRequest struct {
	*aws.Request
	Input *ListWorkersWithQualificationTypeInput
	Copy  func(*ListWorkersWithQualificationTypeInput) ListWorkersWithQualificationTypeRequest
}

// Send marshals and sends the ListWorkersWithQualificationType API request.
func (r ListWorkersWithQualificationTypeRequest) Send(ctx context.Context) (*ListWorkersWithQualificationTypeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListWorkersWithQualificationTypeResponse{
		ListWorkersWithQualificationTypeOutput: r.Request.Data.(*ListWorkersWithQualificationTypeOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListWorkersWithQualificationTypeRequestPaginator returns a paginator for ListWorkersWithQualificationType.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListWorkersWithQualificationTypeRequest(input)
//   p := mturk.NewListWorkersWithQualificationTypeRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListWorkersWithQualificationTypePaginator(req ListWorkersWithQualificationTypeRequest) ListWorkersWithQualificationTypePaginator {
	return ListWorkersWithQualificationTypePaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListWorkersWithQualificationTypeInput
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

// ListWorkersWithQualificationTypePaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListWorkersWithQualificationTypePaginator struct {
	aws.Pager
}

func (p *ListWorkersWithQualificationTypePaginator) CurrentPage() *ListWorkersWithQualificationTypeOutput {
	return p.Pager.CurrentPage().(*ListWorkersWithQualificationTypeOutput)
}

// ListWorkersWithQualificationTypeResponse is the response type for the
// ListWorkersWithQualificationType API operation.
type ListWorkersWithQualificationTypeResponse struct {
	*ListWorkersWithQualificationTypeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListWorkersWithQualificationType request.
func (r *ListWorkersWithQualificationTypeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
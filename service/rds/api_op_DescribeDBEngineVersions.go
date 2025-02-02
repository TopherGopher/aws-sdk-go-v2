// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DescribeDBEngineVersionsMessage
type DescribeDBEngineVersionsInput struct {
	_ struct{} `type:"structure"`

	// The name of a specific DB parameter group family to return details for.
	//
	// Constraints:
	//
	//    * If supplied, must match an existing DBParameterGroupFamily.
	DBParameterGroupFamily *string `type:"string"`

	// A value that indicates whether only the default version of the specified
	// engine or engine and major version combination is returned.
	DefaultOnly *bool `type:"boolean"`

	// The database engine to return.
	Engine *string `type:"string"`

	// The database engine version to return.
	//
	// Example: 5.1.49
	EngineVersion *string `type:"string"`

	// This parameter is not currently supported.
	Filters []Filter `locationNameList:"Filter" type:"list"`

	// A value that indicates whether to include engine versions that aren't available
	// in the list. The default is to list only available engine versions.
	IncludeAll *bool `type:"boolean"`

	// A value that indicates whether to list the supported character sets for each
	// engine version.
	//
	// If this parameter is enabled and the requested engine supports the CharacterSetName
	// parameter for CreateDBInstance, the response includes a list of supported
	// character sets for each engine version.
	ListSupportedCharacterSets *bool `type:"boolean"`

	// A value that indicates whether to list the supported time zones for each
	// engine version.
	//
	// If this parameter is enabled and the requested engine supports the TimeZone
	// parameter for CreateDBInstance, the response includes a list of supported
	// time zones for each engine version.
	ListSupportedTimezones *bool `type:"boolean"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to
	// the value specified by MaxRecords.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more than the
	// MaxRecords value is available, a pagination token called a marker is included
	// in the response so that the following results can be retrieved.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeDBEngineVersionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDBEngineVersionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeDBEngineVersionsInput"}
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

// Contains the result of a successful invocation of the DescribeDBEngineVersions
// action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DBEngineVersionMessage
type DescribeDBEngineVersionsOutput struct {
	_ struct{} `type:"structure"`

	// A list of DBEngineVersion elements.
	DBEngineVersions []DBEngineVersion `locationNameList:"DBEngineVersion" type:"list"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to
	// the value specified by MaxRecords.
	Marker *string `type:"string"`
}

// String returns the string representation
func (s DescribeDBEngineVersionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeDBEngineVersions = "DescribeDBEngineVersions"

// DescribeDBEngineVersionsRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Returns a list of the available DB engines.
//
//    // Example sending a request using DescribeDBEngineVersionsRequest.
//    req := client.DescribeDBEngineVersionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DescribeDBEngineVersions
func (c *Client) DescribeDBEngineVersionsRequest(input *DescribeDBEngineVersionsInput) DescribeDBEngineVersionsRequest {
	op := &aws.Operation{
		Name:       opDescribeDBEngineVersions,
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
		input = &DescribeDBEngineVersionsInput{}
	}

	req := c.newRequest(op, input, &DescribeDBEngineVersionsOutput{})
	return DescribeDBEngineVersionsRequest{Request: req, Input: input, Copy: c.DescribeDBEngineVersionsRequest}
}

// DescribeDBEngineVersionsRequest is the request type for the
// DescribeDBEngineVersions API operation.
type DescribeDBEngineVersionsRequest struct {
	*aws.Request
	Input *DescribeDBEngineVersionsInput
	Copy  func(*DescribeDBEngineVersionsInput) DescribeDBEngineVersionsRequest
}

// Send marshals and sends the DescribeDBEngineVersions API request.
func (r DescribeDBEngineVersionsRequest) Send(ctx context.Context) (*DescribeDBEngineVersionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDBEngineVersionsResponse{
		DescribeDBEngineVersionsOutput: r.Request.Data.(*DescribeDBEngineVersionsOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeDBEngineVersionsRequestPaginator returns a paginator for DescribeDBEngineVersions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeDBEngineVersionsRequest(input)
//   p := rds.NewDescribeDBEngineVersionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeDBEngineVersionsPaginator(req DescribeDBEngineVersionsRequest) DescribeDBEngineVersionsPaginator {
	return DescribeDBEngineVersionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeDBEngineVersionsInput
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

// DescribeDBEngineVersionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeDBEngineVersionsPaginator struct {
	aws.Pager
}

func (p *DescribeDBEngineVersionsPaginator) CurrentPage() *DescribeDBEngineVersionsOutput {
	return p.Pager.CurrentPage().(*DescribeDBEngineVersionsOutput)
}

// DescribeDBEngineVersionsResponse is the response type for the
// DescribeDBEngineVersions API operation.
type DescribeDBEngineVersionsResponse struct {
	*DescribeDBEngineVersionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDBEngineVersions request.
func (r *DescribeDBEngineVersionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

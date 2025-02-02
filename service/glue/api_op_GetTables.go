// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetTablesRequest
type GetTablesInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Data Catalog where the tables reside. If none is provided,
	// the AWS account ID is used by default.
	CatalogId *string `min:"1" type:"string"`

	// The database in the catalog whose tables to list. For Hive compatibility,
	// this name is entirely lowercase.
	//
	// DatabaseName is a required field
	DatabaseName *string `min:"1" type:"string" required:"true"`

	// A regular expression pattern. If present, only those tables whose names match
	// the pattern are returned.
	Expression *string `type:"string"`

	// The maximum number of tables to return in a single response.
	MaxResults *int64 `min:"1" type:"integer"`

	// A continuation token, included if this is a continuation call.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s GetTablesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetTablesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetTablesInput"}
	if s.CatalogId != nil && len(*s.CatalogId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CatalogId", 1))
	}

	if s.DatabaseName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DatabaseName"))
	}
	if s.DatabaseName != nil && len(*s.DatabaseName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DatabaseName", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetTablesResponse
type GetTablesOutput struct {
	_ struct{} `type:"structure"`

	// A continuation token, present if the current list segment is not the last.
	NextToken *string `type:"string"`

	// A list of the requested Table objects.
	TableList []Table `type:"list"`
}

// String returns the string representation
func (s GetTablesOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetTables = "GetTables"

// GetTablesRequest returns a request value for making API operation for
// AWS Glue.
//
// Retrieves the definitions of some or all of the tables in a given Database.
//
//    // Example sending a request using GetTablesRequest.
//    req := client.GetTablesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetTables
func (c *Client) GetTablesRequest(input *GetTablesInput) GetTablesRequest {
	op := &aws.Operation{
		Name:       opGetTables,
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
		input = &GetTablesInput{}
	}

	req := c.newRequest(op, input, &GetTablesOutput{})
	return GetTablesRequest{Request: req, Input: input, Copy: c.GetTablesRequest}
}

// GetTablesRequest is the request type for the
// GetTables API operation.
type GetTablesRequest struct {
	*aws.Request
	Input *GetTablesInput
	Copy  func(*GetTablesInput) GetTablesRequest
}

// Send marshals and sends the GetTables API request.
func (r GetTablesRequest) Send(ctx context.Context) (*GetTablesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetTablesResponse{
		GetTablesOutput: r.Request.Data.(*GetTablesOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetTablesRequestPaginator returns a paginator for GetTables.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetTablesRequest(input)
//   p := glue.NewGetTablesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetTablesPaginator(req GetTablesRequest) GetTablesPaginator {
	return GetTablesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetTablesInput
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

// GetTablesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetTablesPaginator struct {
	aws.Pager
}

func (p *GetTablesPaginator) CurrentPage() *GetTablesOutput {
	return p.Pager.CurrentPage().(*GetTablesOutput)
}

// GetTablesResponse is the response type for the
// GetTables API operation.
type GetTablesResponse struct {
	*GetTablesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetTables request.
func (r *GetTablesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

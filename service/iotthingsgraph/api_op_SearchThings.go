// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotthingsgraph

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/SearchThingsRequest
type SearchThingsInput struct {
	_ struct{} `type:"structure"`

	// The ID of the entity to which the things are associated.
	//
	// The IDs should be in the following format.
	//
	// urn:tdm:REGION/ACCOUNT ID/default:device:DEVICENAME
	//
	// EntityId is a required field
	EntityId *string `locationName:"entityId" type:"string" required:"true"`

	// The maximum number of results to return in the response.
	MaxResults *int64 `locationName:"maxResults" min:"1" type:"integer"`

	// The version of the user's namespace. Defaults to the latest version of the
	// user's namespace.
	NamespaceVersion *int64 `locationName:"namespaceVersion" type:"long"`

	// The string that specifies the next page of results. Use this when you're
	// paginating results.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s SearchThingsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SearchThingsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SearchThingsInput"}

	if s.EntityId == nil {
		invalidParams.Add(aws.NewErrParamRequired("EntityId"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/SearchThingsResponse
type SearchThingsOutput struct {
	_ struct{} `type:"structure"`

	// The string to specify as nextToken when you request the next page of results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// An array of things in the result set.
	Things []Thing `locationName:"things" type:"list"`
}

// String returns the string representation
func (s SearchThingsOutput) String() string {
	return awsutil.Prettify(s)
}

const opSearchThings = "SearchThings"

// SearchThingsRequest returns a request value for making API operation for
// AWS IoT Things Graph.
//
// Searches for things associated with the specified entity. You can search
// by both device and device model.
//
// For example, if two different devices, camera1 and camera2, implement the
// camera device model, the user can associate thing1 to camera1 and thing2
// to camera2. SearchThings(camera2) will return only thing2, but SearchThings(camera)
// will return both thing1 and thing2.
//
// This action searches for exact matches and doesn't perform partial text matching.
//
//    // Example sending a request using SearchThingsRequest.
//    req := client.SearchThingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/SearchThings
func (c *Client) SearchThingsRequest(input *SearchThingsInput) SearchThingsRequest {
	op := &aws.Operation{
		Name:       opSearchThings,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &SearchThingsInput{}
	}

	req := c.newRequest(op, input, &SearchThingsOutput{})
	return SearchThingsRequest{Request: req, Input: input, Copy: c.SearchThingsRequest}
}

// SearchThingsRequest is the request type for the
// SearchThings API operation.
type SearchThingsRequest struct {
	*aws.Request
	Input *SearchThingsInput
	Copy  func(*SearchThingsInput) SearchThingsRequest
}

// Send marshals and sends the SearchThings API request.
func (r SearchThingsRequest) Send(ctx context.Context) (*SearchThingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SearchThingsResponse{
		SearchThingsOutput: r.Request.Data.(*SearchThingsOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewSearchThingsRequestPaginator returns a paginator for SearchThings.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.SearchThingsRequest(input)
//   p := iotthingsgraph.NewSearchThingsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewSearchThingsPaginator(req SearchThingsRequest) SearchThingsPaginator {
	return SearchThingsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *SearchThingsInput
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

// SearchThingsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type SearchThingsPaginator struct {
	aws.Pager
}

func (p *SearchThingsPaginator) CurrentPage() *SearchThingsOutput {
	return p.Pager.CurrentPage().(*SearchThingsOutput)
}

// SearchThingsResponse is the response type for the
// SearchThings API operation.
type SearchThingsResponse struct {
	*SearchThingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SearchThings request.
func (r *SearchThingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

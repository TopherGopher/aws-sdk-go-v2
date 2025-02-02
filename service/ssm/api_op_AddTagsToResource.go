// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/AddTagsToResourceRequest
type AddTagsToResourceInput struct {
	_ struct{} `type:"structure"`

	// The resource ID you want to tag.
	//
	// Use the ID of the resource. Here are some examples:
	//
	// ManagedInstance: mi-012345abcde
	//
	// MaintenanceWindow: mw-012345abcde
	//
	// PatchBaseline: pb-012345abcde
	//
	// For the Document and Parameter values, use the name of the resource.
	//
	// The ManagedInstance type for this API action is only for on-premises managed
	// instances. You must specify the name of the managed instance in the following
	// format: mi-ID_number. For example, mi-1a2b3c4d5e6f.
	//
	// ResourceId is a required field
	ResourceId *string `type:"string" required:"true"`

	// Specifies the type of resource you are tagging.
	//
	// The ManagedInstance type for this API action is for on-premises managed instances.
	// You must specify the name of the managed instance in the following format:
	// mi-ID_number. For example, mi-1a2b3c4d5e6f.
	//
	// ResourceType is a required field
	ResourceType ResourceTypeForTagging `type:"string" required:"true" enum:"true"`

	// One or more tags. The value parameter is required, but if you don't want
	// the tag to have a value, specify the parameter with no value, and we set
	// the value to an empty string.
	//
	// Do not enter personally identifiable information in this field.
	//
	// Tags is a required field
	Tags []Tag `type:"list" required:"true"`
}

// String returns the string representation
func (s AddTagsToResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddTagsToResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AddTagsToResourceInput"}

	if s.ResourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceId"))
	}
	if len(s.ResourceType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ResourceType"))
	}

	if s.Tags == nil {
		invalidParams.Add(aws.NewErrParamRequired("Tags"))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/AddTagsToResourceResult
type AddTagsToResourceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AddTagsToResourceOutput) String() string {
	return awsutil.Prettify(s)
}

const opAddTagsToResource = "AddTagsToResource"

// AddTagsToResourceRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Adds or overwrites one or more tags for the specified resource. Tags are
// metadata that you can assign to your documents, managed instances, maintenance
// windows, Parameter Store parameters, and patch baselines. Tags enable you
// to categorize your resources in different ways, for example, by purpose,
// owner, or environment. Each tag consists of a key and an optional value,
// both of which you define. For example, you could define a set of tags for
// your account's managed instances that helps you track each instance's owner
// and stack level. For example: Key=Owner and Value=DbAdmin, SysAdmin, or Dev.
// Or Key=Stack and Value=Production, Pre-Production, or Test.
//
// Each resource can have a maximum of 50 tags.
//
// We recommend that you devise a set of tag keys that meets your needs for
// each resource type. Using a consistent set of tag keys makes it easier for
// you to manage your resources. You can search and filter the resources based
// on the tags you add. Tags don't have any semantic meaning to Amazon EC2 and
// are interpreted strictly as a string of characters.
//
// For more information about tags, see Tagging Your Amazon EC2 Resources (http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html)
// in the Amazon EC2 User Guide.
//
//    // Example sending a request using AddTagsToResourceRequest.
//    req := client.AddTagsToResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/AddTagsToResource
func (c *Client) AddTagsToResourceRequest(input *AddTagsToResourceInput) AddTagsToResourceRequest {
	op := &aws.Operation{
		Name:       opAddTagsToResource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AddTagsToResourceInput{}
	}

	req := c.newRequest(op, input, &AddTagsToResourceOutput{})
	return AddTagsToResourceRequest{Request: req, Input: input, Copy: c.AddTagsToResourceRequest}
}

// AddTagsToResourceRequest is the request type for the
// AddTagsToResource API operation.
type AddTagsToResourceRequest struct {
	*aws.Request
	Input *AddTagsToResourceInput
	Copy  func(*AddTagsToResourceInput) AddTagsToResourceRequest
}

// Send marshals and sends the AddTagsToResource API request.
func (r AddTagsToResourceRequest) Send(ctx context.Context) (*AddTagsToResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AddTagsToResourceResponse{
		AddTagsToResourceOutput: r.Request.Data.(*AddTagsToResourceOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AddTagsToResourceResponse is the response type for the
// AddTagsToResource API operation.
type AddTagsToResourceResponse struct {
	*AddTagsToResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AddTagsToResource request.
func (r *AddTagsToResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

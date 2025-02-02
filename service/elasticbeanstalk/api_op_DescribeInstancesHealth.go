// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Parameters for a call to DescribeInstancesHealth.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticbeanstalk-2010-12-01/DescribeInstancesHealthRequest
type DescribeInstancesHealthInput struct {
	_ struct{} `type:"structure"`

	// Specifies the response elements you wish to receive. To retrieve all attributes,
	// set to All. If no attribute names are specified, returns a list of instances.
	AttributeNames []InstancesHealthAttribute `type:"list"`

	// Specify the AWS Elastic Beanstalk environment by ID.
	EnvironmentId *string `type:"string"`

	// Specify the AWS Elastic Beanstalk environment by name.
	EnvironmentName *string `min:"4" type:"string"`

	// Specify the pagination token returned by a previous call.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeInstancesHealthInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeInstancesHealthInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeInstancesHealthInput"}
	if s.EnvironmentName != nil && len(*s.EnvironmentName) < 4 {
		invalidParams.Add(aws.NewErrParamMinLen("EnvironmentName", 4))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Detailed health information about the Amazon EC2 instances in an AWS Elastic
// Beanstalk environment.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticbeanstalk-2010-12-01/DescribeInstancesHealthResult
type DescribeInstancesHealthOutput struct {
	_ struct{} `type:"structure"`

	// Detailed health information about each instance.
	//
	// The output differs slightly between Linux and Windows environments. There
	// is a difference in the members that are supported under the <CPUUtilization>
	// type.
	InstanceHealthList []SingleInstanceHealth `type:"list"`

	// Pagination token for the next page of results, if available.
	NextToken *string `min:"1" type:"string"`

	// The date and time that the health information was retrieved.
	RefreshedAt *time.Time `type:"timestamp"`
}

// String returns the string representation
func (s DescribeInstancesHealthOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeInstancesHealth = "DescribeInstancesHealth"

// DescribeInstancesHealthRequest returns a request value for making API operation for
// AWS Elastic Beanstalk.
//
// Retrieves detailed information about the health of instances in your AWS
// Elastic Beanstalk. This operation requires enhanced health reporting (https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/health-enhanced.html).
//
//    // Example sending a request using DescribeInstancesHealthRequest.
//    req := client.DescribeInstancesHealthRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticbeanstalk-2010-12-01/DescribeInstancesHealth
func (c *Client) DescribeInstancesHealthRequest(input *DescribeInstancesHealthInput) DescribeInstancesHealthRequest {
	op := &aws.Operation{
		Name:       opDescribeInstancesHealth,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeInstancesHealthInput{}
	}

	req := c.newRequest(op, input, &DescribeInstancesHealthOutput{})
	return DescribeInstancesHealthRequest{Request: req, Input: input, Copy: c.DescribeInstancesHealthRequest}
}

// DescribeInstancesHealthRequest is the request type for the
// DescribeInstancesHealth API operation.
type DescribeInstancesHealthRequest struct {
	*aws.Request
	Input *DescribeInstancesHealthInput
	Copy  func(*DescribeInstancesHealthInput) DescribeInstancesHealthRequest
}

// Send marshals and sends the DescribeInstancesHealth API request.
func (r DescribeInstancesHealthRequest) Send(ctx context.Context) (*DescribeInstancesHealthResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeInstancesHealthResponse{
		DescribeInstancesHealthOutput: r.Request.Data.(*DescribeInstancesHealthOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeInstancesHealthResponse is the response type for the
// DescribeInstancesHealth API operation.
type DescribeInstancesHealthResponse struct {
	*DescribeInstancesHealthOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeInstancesHealth request.
func (r *DescribeInstancesHealthResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

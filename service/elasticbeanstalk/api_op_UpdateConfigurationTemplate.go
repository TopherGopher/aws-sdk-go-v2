// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The result message containing the options for the specified solution stack.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticbeanstalk-2010-12-01/UpdateConfigurationTemplateMessage
type UpdateConfigurationTemplateInput struct {
	_ struct{} `type:"structure"`

	// The name of the application associated with the configuration template to
	// update.
	//
	// If no application is found with this name, UpdateConfigurationTemplate returns
	// an InvalidParameterValue error.
	//
	// ApplicationName is a required field
	ApplicationName *string `min:"1" type:"string" required:"true"`

	// A new description for the configuration.
	Description *string `type:"string"`

	// A list of configuration option settings to update with the new specified
	// option value.
	OptionSettings []ConfigurationOptionSetting `type:"list"`

	// A list of configuration options to remove from the configuration set.
	//
	// Constraint: You can remove only UserDefined configuration options.
	OptionsToRemove []OptionSpecification `type:"list"`

	// The name of the configuration template to update.
	//
	// If no configuration template is found with this name, UpdateConfigurationTemplate
	// returns an InvalidParameterValue error.
	//
	// TemplateName is a required field
	TemplateName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateConfigurationTemplateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateConfigurationTemplateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateConfigurationTemplateInput"}

	if s.ApplicationName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationName"))
	}
	if s.ApplicationName != nil && len(*s.ApplicationName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApplicationName", 1))
	}

	if s.TemplateName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TemplateName"))
	}
	if s.TemplateName != nil && len(*s.TemplateName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TemplateName", 1))
	}
	if s.OptionSettings != nil {
		for i, v := range s.OptionSettings {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "OptionSettings", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.OptionsToRemove != nil {
		for i, v := range s.OptionsToRemove {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "OptionsToRemove", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Describes the settings for a configuration set.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticbeanstalk-2010-12-01/ConfigurationSettingsDescription
type UpdateConfigurationTemplateOutput struct {
	_ struct{} `type:"structure"`

	// The name of the application associated with this configuration set.
	ApplicationName *string `min:"1" type:"string"`

	// The date (in UTC time) when this configuration set was created.
	DateCreated *time.Time `type:"timestamp"`

	// The date (in UTC time) when this configuration set was last modified.
	DateUpdated *time.Time `type:"timestamp"`

	// If this configuration set is associated with an environment, the DeploymentStatus
	// parameter indicates the deployment status of this configuration set:
	//
	//    * null: This configuration is not associated with a running environment.
	//
	//    * pending: This is a draft configuration that is not deployed to the associated
	//    environment but is in the process of deploying.
	//
	//    * deployed: This is the configuration that is currently deployed to the
	//    associated running environment.
	//
	//    * failed: This is a draft configuration that failed to successfully deploy.
	DeploymentStatus ConfigurationDeploymentStatus `type:"string" enum:"true"`

	// Describes this configuration set.
	Description *string `type:"string"`

	// If not null, the name of the environment for this configuration set.
	EnvironmentName *string `min:"4" type:"string"`

	// A list of the configuration options and their values in this configuration
	// set.
	OptionSettings []ConfigurationOptionSetting `type:"list"`

	// The ARN of the platform.
	PlatformArn *string `type:"string"`

	// The name of the solution stack this configuration set uses.
	SolutionStackName *string `type:"string"`

	// If not null, the name of the configuration template for this configuration
	// set.
	TemplateName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateConfigurationTemplateOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateConfigurationTemplate = "UpdateConfigurationTemplate"

// UpdateConfigurationTemplateRequest returns a request value for making API operation for
// AWS Elastic Beanstalk.
//
// Updates the specified configuration template to have the specified properties
// or configuration option values.
//
// If a property (for example, ApplicationName) is not provided, its value remains
// unchanged. To clear such properties, specify an empty string.
//
// Related Topics
//
//    * DescribeConfigurationOptions
//
//    // Example sending a request using UpdateConfigurationTemplateRequest.
//    req := client.UpdateConfigurationTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticbeanstalk-2010-12-01/UpdateConfigurationTemplate
func (c *Client) UpdateConfigurationTemplateRequest(input *UpdateConfigurationTemplateInput) UpdateConfigurationTemplateRequest {
	op := &aws.Operation{
		Name:       opUpdateConfigurationTemplate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateConfigurationTemplateInput{}
	}

	req := c.newRequest(op, input, &UpdateConfigurationTemplateOutput{})
	return UpdateConfigurationTemplateRequest{Request: req, Input: input, Copy: c.UpdateConfigurationTemplateRequest}
}

// UpdateConfigurationTemplateRequest is the request type for the
// UpdateConfigurationTemplate API operation.
type UpdateConfigurationTemplateRequest struct {
	*aws.Request
	Input *UpdateConfigurationTemplateInput
	Copy  func(*UpdateConfigurationTemplateInput) UpdateConfigurationTemplateRequest
}

// Send marshals and sends the UpdateConfigurationTemplate API request.
func (r UpdateConfigurationTemplateRequest) Send(ctx context.Context) (*UpdateConfigurationTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateConfigurationTemplateResponse{
		UpdateConfigurationTemplateOutput: r.Request.Data.(*UpdateConfigurationTemplateOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateConfigurationTemplateResponse is the response type for the
// UpdateConfigurationTemplate API operation.
type UpdateConfigurationTemplateResponse struct {
	*UpdateConfigurationTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateConfigurationTemplate request.
func (r *UpdateConfigurationTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

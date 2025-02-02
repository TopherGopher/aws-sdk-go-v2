// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package databasemigrationservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/ApplyPendingMaintenanceActionMessage
type ApplyPendingMaintenanceActionInput struct {
	_ struct{} `type:"structure"`

	// The pending maintenance action to apply to this resource.
	//
	// ApplyAction is a required field
	ApplyAction *string `type:"string" required:"true"`

	// A value that specifies the type of opt-in request, or undoes an opt-in request.
	// You can't undo an opt-in request of type immediate.
	//
	// Valid values:
	//
	//    * immediate - Apply the maintenance action immediately.
	//
	//    * next-maintenance - Apply the maintenance action during the next maintenance
	//    window for the resource.
	//
	//    * undo-opt-in - Cancel any existing next-maintenance opt-in requests.
	//
	// OptInType is a required field
	OptInType *string `type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the AWS DMS resource that the pending maintenance
	// action applies to.
	//
	// ReplicationInstanceArn is a required field
	ReplicationInstanceArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ApplyPendingMaintenanceActionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ApplyPendingMaintenanceActionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ApplyPendingMaintenanceActionInput"}

	if s.ApplyAction == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplyAction"))
	}

	if s.OptInType == nil {
		invalidParams.Add(aws.NewErrParamRequired("OptInType"))
	}

	if s.ReplicationInstanceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReplicationInstanceArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/ApplyPendingMaintenanceActionResponse
type ApplyPendingMaintenanceActionOutput struct {
	_ struct{} `type:"structure"`

	// The AWS DMS resource that the pending maintenance action will be applied
	// to.
	ResourcePendingMaintenanceActions *ResourcePendingMaintenanceActions `type:"structure"`
}

// String returns the string representation
func (s ApplyPendingMaintenanceActionOutput) String() string {
	return awsutil.Prettify(s)
}

const opApplyPendingMaintenanceAction = "ApplyPendingMaintenanceAction"

// ApplyPendingMaintenanceActionRequest returns a request value for making API operation for
// AWS Database Migration Service.
//
// Applies a pending maintenance action to a resource (for example, to a replication
// instance).
//
//    // Example sending a request using ApplyPendingMaintenanceActionRequest.
//    req := client.ApplyPendingMaintenanceActionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/ApplyPendingMaintenanceAction
func (c *Client) ApplyPendingMaintenanceActionRequest(input *ApplyPendingMaintenanceActionInput) ApplyPendingMaintenanceActionRequest {
	op := &aws.Operation{
		Name:       opApplyPendingMaintenanceAction,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ApplyPendingMaintenanceActionInput{}
	}

	req := c.newRequest(op, input, &ApplyPendingMaintenanceActionOutput{})
	return ApplyPendingMaintenanceActionRequest{Request: req, Input: input, Copy: c.ApplyPendingMaintenanceActionRequest}
}

// ApplyPendingMaintenanceActionRequest is the request type for the
// ApplyPendingMaintenanceAction API operation.
type ApplyPendingMaintenanceActionRequest struct {
	*aws.Request
	Input *ApplyPendingMaintenanceActionInput
	Copy  func(*ApplyPendingMaintenanceActionInput) ApplyPendingMaintenanceActionRequest
}

// Send marshals and sends the ApplyPendingMaintenanceAction API request.
func (r ApplyPendingMaintenanceActionRequest) Send(ctx context.Context) (*ApplyPendingMaintenanceActionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ApplyPendingMaintenanceActionResponse{
		ApplyPendingMaintenanceActionOutput: r.Request.Data.(*ApplyPendingMaintenanceActionOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ApplyPendingMaintenanceActionResponse is the response type for the
// ApplyPendingMaintenanceAction API operation.
type ApplyPendingMaintenanceActionResponse struct {
	*ApplyPendingMaintenanceActionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ApplyPendingMaintenanceAction request.
func (r *ApplyPendingMaintenanceActionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

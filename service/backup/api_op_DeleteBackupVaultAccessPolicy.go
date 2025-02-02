// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package backup

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/DeleteBackupVaultAccessPolicyInput
type DeleteBackupVaultAccessPolicyInput struct {
	_ struct{} `type:"structure"`

	// The name of a logical container where backups are stored. Backup vaults are
	// identified by names that are unique to the account used to create them and
	// the AWS Region where they are created. They consist of lowercase letters,
	// numbers, and hyphens.
	//
	// BackupVaultName is a required field
	BackupVaultName *string `location:"uri" locationName:"backupVaultName" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteBackupVaultAccessPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteBackupVaultAccessPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteBackupVaultAccessPolicyInput"}

	if s.BackupVaultName == nil {
		invalidParams.Add(aws.NewErrParamRequired("BackupVaultName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteBackupVaultAccessPolicyInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.BackupVaultName != nil {
		v := *s.BackupVaultName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "backupVaultName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/DeleteBackupVaultAccessPolicyOutput
type DeleteBackupVaultAccessPolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteBackupVaultAccessPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteBackupVaultAccessPolicyOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteBackupVaultAccessPolicy = "DeleteBackupVaultAccessPolicy"

// DeleteBackupVaultAccessPolicyRequest returns a request value for making API operation for
// AWS Backup.
//
// Deletes the policy document that manages permissions on a backup vault.
//
//    // Example sending a request using DeleteBackupVaultAccessPolicyRequest.
//    req := client.DeleteBackupVaultAccessPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/DeleteBackupVaultAccessPolicy
func (c *Client) DeleteBackupVaultAccessPolicyRequest(input *DeleteBackupVaultAccessPolicyInput) DeleteBackupVaultAccessPolicyRequest {
	op := &aws.Operation{
		Name:       opDeleteBackupVaultAccessPolicy,
		HTTPMethod: "DELETE",
		HTTPPath:   "/backup-vaults/{backupVaultName}/access-policy",
	}

	if input == nil {
		input = &DeleteBackupVaultAccessPolicyInput{}
	}

	req := c.newRequest(op, input, &DeleteBackupVaultAccessPolicyOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteBackupVaultAccessPolicyRequest{Request: req, Input: input, Copy: c.DeleteBackupVaultAccessPolicyRequest}
}

// DeleteBackupVaultAccessPolicyRequest is the request type for the
// DeleteBackupVaultAccessPolicy API operation.
type DeleteBackupVaultAccessPolicyRequest struct {
	*aws.Request
	Input *DeleteBackupVaultAccessPolicyInput
	Copy  func(*DeleteBackupVaultAccessPolicyInput) DeleteBackupVaultAccessPolicyRequest
}

// Send marshals and sends the DeleteBackupVaultAccessPolicy API request.
func (r DeleteBackupVaultAccessPolicyRequest) Send(ctx context.Context) (*DeleteBackupVaultAccessPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteBackupVaultAccessPolicyResponse{
		DeleteBackupVaultAccessPolicyOutput: r.Request.Data.(*DeleteBackupVaultAccessPolicyOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteBackupVaultAccessPolicyResponse is the response type for the
// DeleteBackupVaultAccessPolicy API operation.
type DeleteBackupVaultAccessPolicyResponse struct {
	*DeleteBackupVaultAccessPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteBackupVaultAccessPolicy request.
func (r *DeleteBackupVaultAccessPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

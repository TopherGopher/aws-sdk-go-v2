// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package backup

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/DescribeBackupJobInput
type DescribeBackupJobInput struct {
	_ struct{} `type:"structure"`

	// Uniquely identifies a request to AWS Backup to back up a resource.
	//
	// BackupJobId is a required field
	BackupJobId *string `location:"uri" locationName:"backupJobId" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeBackupJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeBackupJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeBackupJobInput"}

	if s.BackupJobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("BackupJobId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeBackupJobInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.BackupJobId != nil {
		v := *s.BackupJobId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "backupJobId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/DescribeBackupJobOutput
type DescribeBackupJobOutput struct {
	_ struct{} `type:"structure"`

	// Uniquely identifies a request to AWS Backup to back up a resource.
	BackupJobId *string `type:"string"`

	// The size, in bytes, of a backup.
	BackupSizeInBytes *int64 `type:"long"`

	// An Amazon Resource Name (ARN) that uniquely identifies a backup vault; for
	// example, arn:aws:backup:us-east-1:123456789012:vault:aBackupVault.
	BackupVaultArn *string `type:"string"`

	// The name of a logical container where backups are stored. Backup vaults are
	// identified by names that are unique to the account used to create them and
	// the AWS Region where they are created. They consist of lowercase letters,
	// numbers, and hyphens.
	BackupVaultName *string `type:"string"`

	// The size in bytes transferred to a backup vault at the time that the job
	// status was queried.
	BytesTransferred *int64 `type:"long"`

	// The date and time that a job to create a backup job is completed, in Unix
	// format and Coordinated Universal Time (UTC). The value of CreationDate is
	// accurate to milliseconds. For example, the value 1516925490.087 represents
	// Friday, January 26, 2018 12:11:30.087 AM.
	CompletionDate *time.Time `type:"timestamp"`

	// Contains identifying information about the creation of a backup job, including
	// the BackupPlanArn, BackupPlanId, BackupPlanVersion, and BackupRuleId of the
	// backup plan that is used to create it.
	CreatedBy *RecoveryPointCreator `type:"structure"`

	// The date and time that a backup job is created, in Unix format and Coordinated
	// Universal Time (UTC). The value of CreationDate is accurate to milliseconds.
	// For example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	CreationDate *time.Time `type:"timestamp"`

	// The date and time that a job to back up resources is expected to be completed,
	// in Unix format and Coordinated Universal Time (UTC). The value of ExpectedCompletionDate
	// is accurate to milliseconds. For example, the value 1516925490.087 represents
	// Friday, January 26, 2018 12:11:30.087 AM.
	ExpectedCompletionDate *time.Time `type:"timestamp"`

	// Specifies the IAM role ARN used to create the target recovery point; for
	// example, arn:aws:iam::123456789012:role/S3Access.
	IamRoleArn *string `type:"string"`

	// Contains an estimated percentage that is complete of a job at the time the
	// job status was queried.
	PercentDone *string `type:"string"`

	// An ARN that uniquely identifies a recovery point; for example, arn:aws:backup:us-east-1:123456789012:recovery-point:1EB3B5E7-9EB0-435A-A80B-108B488B0D45.
	RecoveryPointArn *string `type:"string"`

	// An ARN that uniquely identifies a saved resource. The format of the ARN depends
	// on the resource type.
	ResourceArn *string `type:"string"`

	// The type of AWS resource to be backed-up; for example, an Amazon Elastic
	// Block Store (Amazon EBS) volume or an Amazon Relational Database Service
	// (Amazon RDS) database.
	ResourceType *string `type:"string"`

	// Specifies the time in Unix format and Coordinated Universal Time (UTC) when
	// a backup job must be started before it is canceled. The value is calculated
	// by adding the start window to the scheduled time. So if the scheduled time
	// were 6:00 PM and the start window is 2 hours, the StartBy time would be 8:00
	// PM on the date specified. The value of StartBy is accurate to milliseconds.
	// For example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	StartBy *time.Time `type:"timestamp"`

	// The current state of a resource recovery point.
	State BackupJobState `type:"string" enum:"true"`

	// A detailed message explaining the status of the job to back up a resource.
	StatusMessage *string `type:"string"`
}

// String returns the string representation
func (s DescribeBackupJobOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeBackupJobOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.BackupJobId != nil {
		v := *s.BackupJobId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "BackupJobId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BackupSizeInBytes != nil {
		v := *s.BackupSizeInBytes

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "BackupSizeInBytes", protocol.Int64Value(v), metadata)
	}
	if s.BackupVaultArn != nil {
		v := *s.BackupVaultArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "BackupVaultArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BackupVaultName != nil {
		v := *s.BackupVaultName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "BackupVaultName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BytesTransferred != nil {
		v := *s.BytesTransferred

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "BytesTransferred", protocol.Int64Value(v), metadata)
	}
	if s.CompletionDate != nil {
		v := *s.CompletionDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CompletionDate",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.CreatedBy != nil {
		v := s.CreatedBy

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "CreatedBy", v, metadata)
	}
	if s.CreationDate != nil {
		v := *s.CreationDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CreationDate",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.ExpectedCompletionDate != nil {
		v := *s.ExpectedCompletionDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ExpectedCompletionDate",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.IamRoleArn != nil {
		v := *s.IamRoleArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "IamRoleArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PercentDone != nil {
		v := *s.PercentDone

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "PercentDone", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RecoveryPointArn != nil {
		v := *s.RecoveryPointArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RecoveryPointArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResourceArn != nil {
		v := *s.ResourceArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ResourceArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResourceType != nil {
		v := *s.ResourceType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ResourceType", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StartBy != nil {
		v := *s.StartBy

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "StartBy",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if len(s.State) > 0 {
		v := s.State

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "State", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.StatusMessage != nil {
		v := *s.StatusMessage

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "StatusMessage", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opDescribeBackupJob = "DescribeBackupJob"

// DescribeBackupJobRequest returns a request value for making API operation for
// AWS Backup.
//
// Returns metadata associated with creating a backup of a resource.
//
//    // Example sending a request using DescribeBackupJobRequest.
//    req := client.DescribeBackupJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/DescribeBackupJob
func (c *Client) DescribeBackupJobRequest(input *DescribeBackupJobInput) DescribeBackupJobRequest {
	op := &aws.Operation{
		Name:       opDescribeBackupJob,
		HTTPMethod: "GET",
		HTTPPath:   "/backup-jobs/{backupJobId}",
	}

	if input == nil {
		input = &DescribeBackupJobInput{}
	}

	req := c.newRequest(op, input, &DescribeBackupJobOutput{})
	return DescribeBackupJobRequest{Request: req, Input: input, Copy: c.DescribeBackupJobRequest}
}

// DescribeBackupJobRequest is the request type for the
// DescribeBackupJob API operation.
type DescribeBackupJobRequest struct {
	*aws.Request
	Input *DescribeBackupJobInput
	Copy  func(*DescribeBackupJobInput) DescribeBackupJobRequest
}

// Send marshals and sends the DescribeBackupJob API request.
func (r DescribeBackupJobRequest) Send(ctx context.Context) (*DescribeBackupJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeBackupJobResponse{
		DescribeBackupJobOutput: r.Request.Data.(*DescribeBackupJobOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeBackupJobResponse is the response type for the
// DescribeBackupJob API operation.
type DescribeBackupJobResponse struct {
	*DescribeBackupJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeBackupJob request.
func (r *DescribeBackupJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

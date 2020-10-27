// Code generated by smithy-go-codegen DO NOT EDIT.

package ssoadmin

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Attaches an IAM inline policy to a permission set. If the permission set is
// already referenced by one or more account assignments, you will need to call
// ProvisionPermissionSet after this action to apply the corresponding IAM policy
// updates to all assigned accounts.
func (c *Client) PutInlinePolicyToPermissionSet(ctx context.Context, params *PutInlinePolicyToPermissionSetInput, optFns ...func(*Options)) (*PutInlinePolicyToPermissionSetOutput, error) {
	if params == nil {
		params = &PutInlinePolicyToPermissionSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutInlinePolicyToPermissionSet", params, optFns, addOperationPutInlinePolicyToPermissionSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutInlinePolicyToPermissionSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutInlinePolicyToPermissionSetInput struct {

	// The IAM inline policy to attach to a PermissionSet.
	//
	// This member is required.
	InlinePolicy *string

	// The ARN of the SSO instance under which the operation will be executed. For more
	// information about ARNs, see Amazon Resource Names (ARNs) and AWS Service
	// Namespaces in the AWS General Reference.
	//
	// This member is required.
	InstanceArn *string

	// The ARN of the permission set.
	//
	// This member is required.
	PermissionSetArn *string
}

type PutInlinePolicyToPermissionSetOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutInlinePolicyToPermissionSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutInlinePolicyToPermissionSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutInlinePolicyToPermissionSet{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpPutInlinePolicyToPermissionSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutInlinePolicyToPermissionSet(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opPutInlinePolicyToPermissionSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sso",
		OperationName: "PutInlinePolicyToPermissionSet",
	}
}
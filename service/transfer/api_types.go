// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package transfer

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

// Describes the properties of the server that was specified. Information returned
// includes the following: the server Amazon Resource Name (ARN), the authentication
// configuration and type, the logging role, the server ID and state, and assigned
// tags or metadata.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/transfer-2018-11-05/DescribedServer
type DescribedServer struct {
	_ struct{} `type:"structure"`

	// Specifies the unique Amazon Resource Name (ARN) for the server to be described.
	//
	// Arn is a required field
	Arn *string `min:"20" type:"string" required:"true"`

	// The virtual private cloud (VPC) endpoint settings that you configured for
	// your SFTP server.
	EndpointDetails *EndpointDetails `type:"structure"`

	// The type of endpoint that your SFTP server is connected to. If your SFTP
	// server is connected to a VPC endpoint, your server isn't accessible over
	// the public internet.
	EndpointType EndpointType `type:"string" enum:"true"`

	// This value contains the message-digest algorithm (MD5) hash of the server's
	// host key. This value is equivalent to the output of the ssh-keygen -l -E
	// md5 -f my-new-server-key command.
	HostKeyFingerprint *string `type:"string"`

	// Specifies information to call a customer-supplied authentication API. This
	// field is not populated when the IdentityProviderType of the server is SERVICE_MANAGED>.
	IdentityProviderDetails *IdentityProviderDetails `type:"structure"`

	// This property defines the mode of authentication method enabled for this
	// service. A value of SERVICE_MANAGED means that you are using this server
	// to store and access SFTP user credentials within the service. A value of
	// API_GATEWAY indicates that you have integrated an API Gateway endpoint that
	// will be invoked for authenticating your user into the service.
	IdentityProviderType IdentityProviderType `type:"string" enum:"true"`

	// This property is an AWS Identity and Access Management (IAM) entity that
	// allows the server to turn on Amazon CloudWatch logging for Amazon S3 events.
	// When set, user activity can be viewed in your CloudWatch logs.
	LoggingRole *string `type:"string"`

	// This property is a unique system-assigned identifier for the SFTP server
	// that you instantiate.
	ServerId *string `type:"string"`

	// The condition of the SFTP server for the server that was described. A value
	// of ONLINE indicates that the server can accept jobs and transfer files. A
	// State value of OFFLINE means that the server cannot perform file transfer
	// operations.
	//
	// The states of STARTING and STOPPING indicate that the server is in an intermediate
	// state, either not fully able to respond, or not fully offline. The values
	// of START_FAILED or STOP_FAILED can indicate an error condition.
	State State `type:"string" enum:"true"`

	// This property contains the key-value pairs that you can use to search for
	// and group servers that were assigned to the server that was described.
	Tags []Tag `min:"1" type:"list"`

	// The number of users that are assigned to the SFTP server you specified with
	// the ServerId.
	UserCount *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribedServer) String() string {
	return awsutil.Prettify(s)
}

// Returns properties of the user that you want to describe.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/transfer-2018-11-05/DescribedUser
type DescribedUser struct {
	_ struct{} `type:"structure"`

	// This property contains the unique Amazon Resource Name (ARN) for the user
	// that was requested to be described.
	//
	// Arn is a required field
	Arn *string `min:"20" type:"string" required:"true"`

	// This property specifies the landing directory (or folder), which is the location
	// that files are written to or read from in an Amazon S3 bucket for the described
	// user. An example is /bucket_name/home/username .
	HomeDirectory *string `type:"string"`

	// Specifies the name of the policy in use for the described user.
	Policy *string `type:"string"`

	// This property specifies the IAM role that controls your user's access to
	// your Amazon S3 bucket. The policies attached to this role will determine
	// the level of access you want to provide your users when transferring files
	// into and out of your Amazon S3 bucket or buckets. The IAM role should also
	// contain a trust relationship that allows the SFTP server to access your resources
	// when servicing your SFTP user's transfer requests.
	Role *string `type:"string"`

	// This property contains the public key portion of the Secure Shell (SSH) keys
	// stored for the described user.
	SshPublicKeys []SshPublicKey `type:"list"`

	// This property contains the key-value pairs for the user requested. Tag can
	// be used to search for and group users for a variety of purposes.
	Tags []Tag `min:"1" type:"list"`

	// This property is the name of the user that was requested to be described.
	// User names are used for authentication purposes. This is the string that
	// will be used by your user when they log in to your SFTP server.
	UserName *string `type:"string"`
}

// String returns the string representation
func (s DescribedUser) String() string {
	return awsutil.Prettify(s)
}

// The configuration settings for the virtual private cloud (VPC) endpoint for
// your SFTP server.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/transfer-2018-11-05/EndpointDetails
type EndpointDetails struct {
	_ struct{} `type:"structure"`

	// The ID of the VPC endpoint.
	VpcEndpointId *string `type:"string"`
}

// String returns the string representation
func (s EndpointDetails) String() string {
	return awsutil.Prettify(s)
}

// Returns information related to the type of user authentication that is in
// use for a server's users. A server can have only one method of authentication.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/transfer-2018-11-05/IdentityProviderDetails
type IdentityProviderDetails struct {
	_ struct{} `type:"structure"`

	// The InvocationRole parameter provides the type of InvocationRole used to
	// authenticate the user account.
	InvocationRole *string `type:"string"`

	// The Url parameter provides contains the location of the service endpoint
	// used to authenticate users.
	Url *string `type:"string"`
}

// String returns the string representation
func (s IdentityProviderDetails) String() string {
	return awsutil.Prettify(s)
}

// Returns properties of the server that was specified.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/transfer-2018-11-05/ListedServer
type ListedServer struct {
	_ struct{} `type:"structure"`

	// The unique Amazon Resource Name (ARN) for the server to be listed.
	//
	// Arn is a required field
	Arn *string `min:"20" type:"string" required:"true"`

	// The type of VPC endpoint that your SFTP server is connected to. If your SFTP
	// server is connected to a VPC endpoint, your server isn't accessible over
	// the public internet.
	EndpointType EndpointType `type:"string" enum:"true"`

	// The authentication method used to validate a user for the server that was
	// specified. This can include Secure Shell (SSH), user name and password combinations,
	// or your own custom authentication method. Valid values include SERVICE_MANAGED
	// or API_GATEWAY.
	IdentityProviderType IdentityProviderType `type:"string" enum:"true"`

	// The AWS Identity and Access Management entity that allows the server to turn
	// on Amazon CloudWatch logging.
	LoggingRole *string `type:"string"`

	// This value is the unique system assigned identifier for the SFTP servers
	// that were listed.
	ServerId *string `type:"string"`

	// This property describes the condition of the SFTP server for the server that
	// was described. A value of ONLINE> indicates that the server can accept jobs
	// and transfer files. A State value of OFFLINE means that the server cannot
	// perform file transfer operations.
	//
	// The states of STARTING and STOPPING indicate that the server is in an intermediate
	// state, either not fully able to respond, or not fully offline. The values
	// of START_FAILED or STOP_FAILED can indicate an error condition.
	State State `type:"string" enum:"true"`

	// This property is a numeric value that indicates the number of users that
	// are assigned to the SFTP server you specified with the ServerId.
	UserCount *int64 `type:"integer"`
}

// String returns the string representation
func (s ListedServer) String() string {
	return awsutil.Prettify(s)
}

// Returns properties of the user that you specify.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/transfer-2018-11-05/ListedUser
type ListedUser struct {
	_ struct{} `type:"structure"`

	// This property is the unique Amazon Resource Name (ARN) for the user that
	// you want to learn about.
	//
	// Arn is a required field
	Arn *string `min:"20" type:"string" required:"true"`

	// This value specifies the location that files are written to or read from
	// an Amazon S3 bucket for the user you specify by their ARN.
	HomeDirectory *string `type:"string"`

	// The role in use by this user. A role is an AWS Identity and Access Management
	// (IAM) entity that, in this case, allows the SFTP server to act on a user's
	// behalf. It allows the server to inherit the trust relationship that enables
	// that user to perform file operations to their Amazon S3 bucket.
	Role *string `type:"string"`

	// This value is the number of SSH public keys stored for the user you specified.
	SshPublicKeyCount *int64 `type:"integer"`

	// The name of the user whose ARN was specified. User names are used for authentication
	// purposes.
	UserName *string `type:"string"`
}

// String returns the string representation
func (s ListedUser) String() string {
	return awsutil.Prettify(s)
}

// Provides information about the public Secure Shell (SSH) key that is associated
// with a user account for a specific server (as identified by ServerId). The
// information returned includes the date the key was imported, the public key
// contents, and the public key ID. A user can store more than one SSH public
// key associated with their user name on a specific SFTP server.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/transfer-2018-11-05/SshPublicKey
type SshPublicKey struct {
	_ struct{} `type:"structure"`

	// The date that the public key was added to the user account.
	//
	// DateImported is a required field
	DateImported *time.Time `type:"timestamp" required:"true"`

	// The content of the SSH public key as specified by the PublicKeyId.
	//
	// SshPublicKeyBody is a required field
	SshPublicKeyBody *string `type:"string" required:"true"`

	// The SshPublicKeyId parameter contains the identifier of the public key.
	//
	// SshPublicKeyId is a required field
	SshPublicKeyId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s SshPublicKey) String() string {
	return awsutil.Prettify(s)
}

// Creates a key-value pair for a specific resource. Tags are metadata that
// you can use to search for and group a resource for various purposes. You
// can apply tags to servers, users, and roles. A tag key can take more than
// one value. For example, to group servers for accounting purposes, you might
// create a tag called Group and assign the values Research and Accounting to
// that group.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/transfer-2018-11-05/Tag
type Tag struct {
	_ struct{} `type:"structure"`

	// The name assigned to the tag that you create.
	//
	// Key is a required field
	Key *string `type:"string" required:"true"`

	// This property contains one or more values that you assigned to the key name
	// you create.
	//
	// Value is a required field
	Value *string `type:"string" required:"true"`
}

// String returns the string representation
func (s Tag) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Tag) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Tag"}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}

	if s.Value == nil {
		invalidParams.Add(aws.NewErrParamRequired("Value"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

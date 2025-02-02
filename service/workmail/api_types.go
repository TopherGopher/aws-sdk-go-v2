// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workmail

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

// At least one delegate must be associated to the resource to disable automatic
// replies from the resource.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/BookingOptions
type BookingOptions struct {
	_ struct{} `type:"structure"`

	// The resource's ability to automatically reply to requests. If disabled, delegates
	// must be associated to the resource.
	AutoAcceptRequests *bool `type:"boolean"`

	// The resource's ability to automatically decline any conflicting requests.
	AutoDeclineConflictingRequests *bool `type:"boolean"`

	// The resource's ability to automatically decline any recurring requests.
	AutoDeclineRecurringRequests *bool `type:"boolean"`
}

// String returns the string representation
func (s BookingOptions) String() string {
	return awsutil.Prettify(s)
}

// The name of the attribute, which is one of the values defined in the UserAttribute
// enumeration.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/Delegate
type Delegate struct {
	_ struct{} `type:"structure"`

	// The identifier for the user or group associated as the resource's delegate.
	//
	// Id is a required field
	Id *string `type:"string" required:"true"`

	// The type of the delegate: user or group.
	//
	// Type is a required field
	Type MemberType `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s Delegate) String() string {
	return awsutil.Prettify(s)
}

// The representation of an Amazon WorkMail group.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/Group
type Group struct {
	_ struct{} `type:"structure"`

	// The date indicating when the group was disabled from Amazon WorkMail use.
	DisabledDate *time.Time `type:"timestamp"`

	// The email of the group.
	Email *string `min:"1" type:"string"`

	// The date indicating when the group was enabled for Amazon WorkMail use.
	EnabledDate *time.Time `type:"timestamp"`

	// The identifier of the group.
	Id *string `min:"12" type:"string"`

	// The name of the group.
	Name *string `min:"1" type:"string"`

	// The state of the group, which can be ENABLED, DISABLED, or DELETED.
	State EntityState `type:"string" enum:"true"`
}

// String returns the string representation
func (s Group) String() string {
	return awsutil.Prettify(s)
}

// The representation of a user or group.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/Member
type Member struct {
	_ struct{} `type:"structure"`

	// The date indicating when the member was disabled from Amazon WorkMail use.
	DisabledDate *time.Time `type:"timestamp"`

	// The date indicating when the member was enabled for Amazon WorkMail use.
	EnabledDate *time.Time `type:"timestamp"`

	// The identifier of the member.
	Id *string `type:"string"`

	// The name of the member.
	Name *string `type:"string"`

	// The state of the member, which can be ENABLED, DISABLED, or DELETED.
	State EntityState `type:"string" enum:"true"`

	// A member can be a user or group.
	Type MemberType `type:"string" enum:"true"`
}

// String returns the string representation
func (s Member) String() string {
	return awsutil.Prettify(s)
}

// The representation of an organization.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/OrganizationSummary
type OrganizationSummary struct {
	_ struct{} `type:"structure"`

	// The alias associated with the organization.
	Alias *string `min:"1" type:"string"`

	// The error message associated with the organization. It is only present if
	// unexpected behavior has occurred with regards to the organization. It provides
	// insight or solutions regarding unexpected behavior.
	ErrorMessage *string `type:"string"`

	// The identifier associated with the organization.
	OrganizationId *string `type:"string"`

	// The state associated with the organization.
	State *string `type:"string"`
}

// String returns the string representation
func (s OrganizationSummary) String() string {
	return awsutil.Prettify(s)
}

// Permission granted to a user, group, or resource to access a certain aspect
// of another user, group, or resource mailbox.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/Permission
type Permission struct {
	_ struct{} `type:"structure"`

	// The identifier of the user, group, or resource to which the permissions are
	// granted.
	//
	// GranteeId is a required field
	GranteeId *string `min:"12" type:"string" required:"true"`

	// The type of user, group, or resource referred to in GranteeId.
	//
	// GranteeType is a required field
	GranteeType MemberType `type:"string" required:"true" enum:"true"`

	// The permissions granted to the grantee. SEND_AS allows the grantee to send
	// email as the owner of the mailbox (the grantee is not mentioned on these
	// emails). SEND_ON_BEHALF allows the grantee to send email on behalf of the
	// owner of the mailbox (the grantee is not mentioned as the physical sender
	// of these emails). FULL_ACCESS allows the grantee full access to the mailbox,
	// irrespective of other folder-level permissions set on the mailbox.
	//
	// PermissionValues is a required field
	PermissionValues []PermissionType `type:"list" required:"true"`
}

// String returns the string representation
func (s Permission) String() string {
	return awsutil.Prettify(s)
}

// The representation of a resource.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/Resource
type Resource struct {
	_ struct{} `type:"structure"`

	// The date indicating when the resource was disabled from Amazon WorkMail use.
	DisabledDate *time.Time `type:"timestamp"`

	// The email of the resource.
	Email *string `min:"1" type:"string"`

	// The date indicating when the resource was enabled for Amazon WorkMail use.
	EnabledDate *time.Time `type:"timestamp"`

	// The identifier of the resource.
	Id *string `min:"12" type:"string"`

	// The name of the resource.
	Name *string `min:"1" type:"string"`

	// The state of the resource, which can be ENABLED, DISABLED, or DELETED.
	State EntityState `type:"string" enum:"true"`

	// The type of the resource: equipment or room.
	Type ResourceType `type:"string" enum:"true"`
}

// String returns the string representation
func (s Resource) String() string {
	return awsutil.Prettify(s)
}

// The representation of an Amazon WorkMail user.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/User
type User struct {
	_ struct{} `type:"structure"`

	// The date indicating when the user was disabled from Amazon WorkMail use.
	DisabledDate *time.Time `type:"timestamp"`

	// The display name of the user.
	DisplayName *string `type:"string"`

	// The email of the user.
	Email *string `min:"1" type:"string"`

	// The date indicating when the user was enabled for Amazon WorkMail use.
	EnabledDate *time.Time `type:"timestamp"`

	// The identifier of the user.
	Id *string `min:"12" type:"string"`

	// The name of the user.
	Name *string `min:"1" type:"string"`

	// The state of the user, which can be ENABLED, DISABLED, or DELETED.
	State EntityState `type:"string" enum:"true"`

	// The role of the user.
	UserRole UserRole `type:"string" enum:"true"`
}

// String returns the string representation
func (s User) String() string {
	return awsutil.Prettify(s)
}

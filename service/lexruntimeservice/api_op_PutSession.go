// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lexruntimeservice

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/runtime.lex-2016-11-28/PutSessionRequest
type PutSessionInput struct {
	_ struct{} `type:"structure"`

	// The message that Amazon Lex returns in the response can be either text or
	// speech based depending on the value of this field.
	//
	//    * If the value is text/plain; charset=utf-8, Amazon Lex returns text in
	//    the response.
	//
	//    * If the value begins with audio/, Amazon Lex returns speech in the response.
	//    Amazon Lex uses Amazon Polly to generate the speech in the configuration
	//    that you specify. For example, if you specify audio/mpeg as the value,
	//    Amazon Lex returns speech in the MPEG format.
	//
	//    * If the value is audio/pcm, the speech is returned as audio/pcm in 16-bit,
	//    little endian format.
	//
	//    * The following are the accepted values: audio/mpeg audio/ogg audio/pcm
	//    audio/* (defaults to mpeg) text/plain; charset=utf-8
	Accept *string `location:"header" locationName:"Accept" type:"string"`

	// The alias in use for the bot that contains the session data.
	//
	// BotAlias is a required field
	BotAlias *string `location:"uri" locationName:"botAlias" type:"string" required:"true"`

	// The name of the bot that contains the session data.
	//
	// BotName is a required field
	BotName *string `location:"uri" locationName:"botName" type:"string" required:"true"`

	// Sets the next action that the bot should take to fulfill the conversation.
	DialogAction *DialogAction `locationName:"dialogAction" type:"structure"`

	// Map of key/value pairs representing the session-specific context information.
	// It contains application information passed between Amazon Lex and a client
	// application.
	SessionAttributes map[string]string `locationName:"sessionAttributes" type:"map"`

	// The ID of the client application user. Amazon Lex uses this to identify a
	// user's conversation with your bot.
	//
	// UserId is a required field
	UserId *string `location:"uri" locationName:"userId" min:"2" type:"string" required:"true"`
}

// String returns the string representation
func (s PutSessionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutSessionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutSessionInput"}

	if s.BotAlias == nil {
		invalidParams.Add(aws.NewErrParamRequired("BotAlias"))
	}

	if s.BotName == nil {
		invalidParams.Add(aws.NewErrParamRequired("BotName"))
	}

	if s.UserId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserId"))
	}
	if s.UserId != nil && len(*s.UserId) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("UserId", 2))
	}
	if s.DialogAction != nil {
		if err := s.DialogAction.Validate(); err != nil {
			invalidParams.AddNested("DialogAction", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutSessionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.DialogAction != nil {
		v := s.DialogAction

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "dialogAction", v, metadata)
	}
	if s.SessionAttributes != nil {
		v := s.SessionAttributes

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "sessionAttributes", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.Accept != nil {
		v := *s.Accept

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Accept", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BotAlias != nil {
		v := *s.BotAlias

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "botAlias", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BotName != nil {
		v := *s.BotName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "botName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.UserId != nil {
		v := *s.UserId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "userId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/runtime.lex-2016-11-28/PutSessionResponse
type PutSessionOutput struct {
	_ struct{} `type:"structure" payload:"AudioStream"`

	// The audio version of the message to convey to the user.
	AudioStream io.ReadCloser `locationName:"audioStream" type:"blob"`

	// Content type as specified in the Accept HTTP header in the request.
	ContentType *string `location:"header" locationName:"Content-Type" type:"string"`

	//    * ConfirmIntent - Amazon Lex is expecting a "yes" or "no" response to
	//    confirm the intent before fulfilling an intent.
	//
	//    * ElicitIntent - Amazon Lex wants to elicit the user's intent.
	//
	//    * ElicitSlot - Amazon Lex is expecting the value of a slot for the current
	//    intent.
	//
	//    * Failed - Conveys that the conversation with the user has failed. This
	//    can happen for various reasons, including the user does not provide an
	//    appropriate response to prompts from the service, or if the Lambda function
	//    fails to fulfill the intent.
	//
	//    * Fulfilled - Conveys that the Lambda function has sucessfully fulfilled
	//    the intent.
	//
	//    * ReadyForFulfillment - Conveys that the client has to fulfill the intent.
	DialogState DialogState `location:"header" locationName:"x-amz-lex-dialog-state" type:"string" enum:"true"`

	// The name of the current intent.
	IntentName *string `location:"header" locationName:"x-amz-lex-intent-name" type:"string"`

	// The next message that should be presented to the user.
	Message *string `location:"header" locationName:"x-amz-lex-message" min:"1" type:"string"`

	// The format of the response message. One of the following values:
	//
	//    * PlainText - The message contains plain UTF-8 text.
	//
	//    * CustomPayload - The message is a custom format for the client.
	//
	//    * SSML - The message contains text formatted for voice output.
	//
	//    * Composite - The message contains an escaped JSON object containing one
	//    or more messages from the groups that messages were assigned to when the
	//    intent was created.
	MessageFormat MessageFormatType `location:"header" locationName:"x-amz-lex-message-format" type:"string" enum:"true"`

	// Map of key/value pairs representing session-specific context information.
	SessionAttributes aws.JSONValue `location:"header" locationName:"x-amz-lex-session-attributes" type:"jsonvalue"`

	// A unique identifier for the session.
	SessionId *string `location:"header" locationName:"x-amz-lex-session-id" type:"string"`

	// If the dialogState is ElicitSlot, returns the name of the slot for which
	// Amazon Lex is eliciting a value.
	SlotToElicit *string `location:"header" locationName:"x-amz-lex-slot-to-elicit" type:"string"`

	// Map of zero or more intent slots Amazon Lex detected from the user input
	// during the conversation.
	//
	// Amazon Lex creates a resolution list containing likely values for a slot.
	// The value that it returns is determined by the valueSelectionStrategy selected
	// when the slot type was created or updated. If valueSelectionStrategy is set
	// to ORIGINAL_VALUE, the value provided by the user is returned, if the user
	// value is similar to the slot values. If valueSelectionStrategy is set to
	// TOP_RESOLUTION Amazon Lex returns the first value in the resolution list
	// or, if there is no resolution list, null. If you don't specify a valueSelectionStrategy
	// the default is ORIGINAL_VALUE.
	Slots aws.JSONValue `location:"header" locationName:"x-amz-lex-slots" type:"jsonvalue"`
}

// String returns the string representation
func (s PutSessionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutSessionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ContentType != nil {
		v := *s.ContentType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.DialogState) > 0 {
		v := s.DialogState

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-lex-dialog-state", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.IntentName != nil {
		v := *s.IntentName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-lex-intent-name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Message != nil {
		v := *s.Message

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-lex-message", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.MessageFormat) > 0 {
		v := s.MessageFormat

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-lex-message-format", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.SessionAttributes != nil {
		v := s.SessionAttributes

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-lex-session-attributes", protocol.JSONValue{V: v, EscapeMode: protocol.Base64Escape}, metadata)
	}
	if s.SessionId != nil {
		v := *s.SessionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-lex-session-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SlotToElicit != nil {
		v := *s.SlotToElicit

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-lex-slot-to-elicit", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Slots != nil {
		v := s.Slots

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-lex-slots", protocol.JSONValue{V: v, EscapeMode: protocol.Base64Escape}, metadata)
	}
	// Skipping AudioStream Output type's body not valid.
	return nil
}

const opPutSession = "PutSession"

// PutSessionRequest returns a request value for making API operation for
// Amazon Lex Runtime Service.
//
// Creates a new session or modifies an existing session with an Amazon Lex
// bot. Use this operation to enable your application to set the state of the
// bot.
//
// For more information, see Managing Sessions (https://docs.aws.amazon.com/lex/latest/dg/how-session-api.html).
//
//    // Example sending a request using PutSessionRequest.
//    req := client.PutSessionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/runtime.lex-2016-11-28/PutSession
func (c *Client) PutSessionRequest(input *PutSessionInput) PutSessionRequest {
	op := &aws.Operation{
		Name:       opPutSession,
		HTTPMethod: "POST",
		HTTPPath:   "/bot/{botName}/alias/{botAlias}/user/{userId}/session",
	}

	if input == nil {
		input = &PutSessionInput{}
	}

	req := c.newRequest(op, input, &PutSessionOutput{})
	return PutSessionRequest{Request: req, Input: input, Copy: c.PutSessionRequest}
}

// PutSessionRequest is the request type for the
// PutSession API operation.
type PutSessionRequest struct {
	*aws.Request
	Input *PutSessionInput
	Copy  func(*PutSessionInput) PutSessionRequest
}

// Send marshals and sends the PutSession API request.
func (r PutSessionRequest) Send(ctx context.Context) (*PutSessionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutSessionResponse{
		PutSessionOutput: r.Request.Data.(*PutSessionOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutSessionResponse is the response type for the
// PutSession API operation.
type PutSessionResponse struct {
	*PutSessionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutSession request.
func (r *PutSessionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

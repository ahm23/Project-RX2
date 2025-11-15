package types

// storage module event types
const (
	EventTypeSignContract   = "post_file"
	EventTypeCancelContract = "delete_file"

	AttributeValueCategory = ModuleName

	AttributeKeySigner   = "signer" // sign storage deal
	AttributeKeyContract = "file"
	AttributeKeyPayOnce  = "pay_once"
	AttributeKeyStart    = "start"
)

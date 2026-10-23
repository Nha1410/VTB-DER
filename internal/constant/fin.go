// Package constant holds the fixed values of the VSD FIN wire format (Phụ lục 06, QĐ 115/QĐ-VSD).
// It imports nothing but the standard library.
package constant

// FIN message framing, outbound direction (member → VSD). See docs/02-cau-truc-dien-fin.md.
const (
	// VSDBIC is the 8-character BICCODE of VSD acting as CCP; the default receiver of every message.
	VSDBIC = "VSDSVN06"

	// BICLength is the length of a BICCODE as issued by VSD (`VSD` + 3 letters of the member + `XX`).
	BICLength = 8

	// Block 1 (basic header): `{1:` F 01 <BIC8>A XXX <session:4><isn:6> `}`.
	basicHeaderApplicationID = "F"
	basicHeaderServiceID     = "01"

	// SenderTerminalSuffix completes the 12-character sender address: BIC8 + "A" + "XXX".
	SenderTerminalSuffix = "AXXX"
	// ReceiverTerminalSuffix completes the 12-character receiver address: BIC8 + "XXXX" ("always X").
	ReceiverTerminalSuffix = "XXXX"

	// Block 2 (application header) input identifier.
	AppHeaderInputID = "I"

	// MaxSessionNumber and MaxSequenceNumber bound the 4-digit session and 6-digit ISN fields.
	MaxSessionNumber  = 9999
	MaxSequenceNumber = 999999

	// MessageTypeLength is the digit count of an MT code such as "598".
	MessageTypeLength = 3

	// ObsolescencePeriod is the only value block 2 accepts for the optional obsolescence field.
	ObsolescencePeriod = "020"

	// TextBlockEnd closes block 4 and is the only line allowed to start with "-".
	TextBlockEnd = "-}"

	// TrailerBlock is block 5; the spec fixes no content and says to copy the sample.
	TrailerBlock = "{5:{MAC:00000000}{CHK:F1DBCA886BBF}{TNG:}}"

	// LineBreak separates the lines of block 4. The spec does not state it; CRLF is the SWIFT FIN rule.
	LineBreak = "\r\n"
)

// Message priority (block 2 field e).
const (
	PriorityUrgent = "U"
	PriorityNormal = "N"
)

// BasicHeaderPrefix is the constant start of block 1 up to the sender address.
const BasicHeaderPrefix = "{1:" + basicHeaderApplicationID + basicHeaderServiceID

package model

// FINMessage is one outbound message (member → VSD) before framing. It carries what blocks 1 and 2
// need plus the already-assembled lines of block 4.
type FINMessage struct {
	// SenderBIC is the 8-character BICCODE of the sending member, e.g. "VSDSSIXX".
	SenderBIC string
	// SessionNumber is the 4-digit session (0-9999).
	SessionNumber int
	// SequenceNumber is the input sequence number, ISN (0-999999). It must differ for every message
	// sent within one session; the caller owns that guarantee.
	SequenceNumber int
	// MessageType is the 3-digit MT code, e.g. "598".
	MessageType string
	// ReceiverBIC is the 8-character BICCODE of the receiver. Empty means VSD (constant.VSDBIC).
	ReceiverBIC string
	// Priority is constant.PriorityUrgent or constant.PriorityNormal.
	Priority string
	// DeliveryMonitoring is "1" or "3" for urgent messages, "2" or empty for normal ones.
	DeliveryMonitoring string
	// ObsolescencePeriod is empty or constant.ObsolescencePeriod.
	ObsolescencePeriod string
	// Text holds the lines of block 4 in order, e.g. ":20:REF0001". Free text inside a line (names,
	// addresses) must already be converted with iso15022.Encode: the encoder only validates it.
	Text []string
}

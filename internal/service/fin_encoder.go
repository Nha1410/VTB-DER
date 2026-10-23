// Package service holds the use cases. It knows nothing about transports (HTTP, files, TCP): it
// turns domain types into wire bytes and reports problems as *model.CustomError.
package service

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Nha1410/VTB-DER/internal/common/iso15022"
	"github.com/Nha1410/VTB-DER/internal/constant"
	"github.com/Nha1410/VTB-DER/internal/model"
)

var (
	bicPattern         = regexp.MustCompile(`^[A-Z0-9]{8}$`)
	messageTypePattern = regexp.MustCompile(`^[0-9]{3}$`)
)

// EncodeFIN frames a message the way docs/02-cau-truc-dien-fin.md section 1.1 prescribes -
// block 1 (basic header), block 2 (application header), block 4 (text) and block 5 (trailer) - and
// returns the bytes of the .fin file to drop into the GateWay Client `send` folder.
//
// It validates everything it can know about the frame and the character set of block 4, and returns
// a *model.CustomError with code ErrBadRequest when the message cannot be sent. It does not convert
// Vietnamese: free text must already be converted with iso15022.Encode, otherwise validation fails.
func EncodeFIN(msg model.FINMessage) (fin []byte, err error) {
	receiverBIC := msg.ReceiverBIC
	if receiverBIC == "" {
		receiverBIC = constant.VSDBIC
	}
	if err = validateFrame(msg, receiverBIC); err != nil {
		return nil, err
	}
	if err = validateText(msg.Text); err != nil {
		return nil, err
	}

	var builder strings.Builder
	builder.WriteString(basicHeader(msg))
	builder.WriteString(applicationHeader(msg, receiverBIC))
	builder.WriteString("{4:" + constant.LineBreak)
	for _, line := range msg.Text {
		builder.WriteString(line + constant.LineBreak)
	}
	builder.WriteString(constant.TextBlockEnd)
	builder.WriteString(constant.TrailerBlock)
	return []byte(builder.String()), nil
}

// basicHeader builds block 1: {1:F01<BIC8>AXXX<session:4><isn:6>}.
func basicHeader(msg model.FINMessage) (block string) {
	return fmt.Sprintf("%s%s%s%04d%06d}", constant.BasicHeaderPrefix, msg.SenderBIC, constant.SenderTerminalSuffix,
		msg.SessionNumber, msg.SequenceNumber)
}

// applicationHeader builds block 2: {2:I<MT><BIC8>XXXX<priority>[<monitoring>][<obsolescence>]}.
func applicationHeader(msg model.FINMessage, receiverBIC string) (block string) {
	return "{2:" + constant.AppHeaderInputID + msg.MessageType + receiverBIC + constant.ReceiverTerminalSuffix +
		msg.Priority + msg.DeliveryMonitoring + msg.ObsolescencePeriod + "}"
}

// validateFrame checks the fields of blocks 1 and 2.
func validateFrame(msg model.FINMessage, receiverBIC string) (err error) {
	switch {
	case !bicPattern.MatchString(msg.SenderBIC):
		return badRequest("sender BIC must be 8 upper-case letters or digits, got %q", msg.SenderBIC)
	case !bicPattern.MatchString(receiverBIC):
		return badRequest("receiver BIC must be 8 upper-case letters or digits, got %q", receiverBIC)
	case !messageTypePattern.MatchString(msg.MessageType):
		return badRequest("message type must be %d digits, got %q", constant.MessageTypeLength, msg.MessageType)
	case msg.SessionNumber < 0 || msg.SessionNumber > constant.MaxSessionNumber:
		return badRequest("session number must be 0-%d, got %d", constant.MaxSessionNumber, msg.SessionNumber)
	case msg.SequenceNumber < 0 || msg.SequenceNumber > constant.MaxSequenceNumber:
		return badRequest("sequence number must be 0-%d, got %d", constant.MaxSequenceNumber, msg.SequenceNumber)
	case msg.ObsolescencePeriod != "" && msg.ObsolescencePeriod != constant.ObsolescencePeriod:
		return badRequest("obsolescence period must be empty or %q, got %q", constant.ObsolescencePeriod, msg.ObsolescencePeriod)
	}
	return validatePriority(msg.Priority, msg.DeliveryMonitoring)
}

// validatePriority applies the rule of block 2 fields (e) and (f): urgent takes monitoring 1 or 3,
// normal takes 2 or nothing.
func validatePriority(priority string, monitoring string) (err error) {
	switch priority {
	case constant.PriorityUrgent:
		if monitoring != "1" && monitoring != "3" {
			return badRequest("urgent priority needs delivery monitoring 1 or 3, got %q", monitoring)
		}
	case constant.PriorityNormal:
		if monitoring != "" && monitoring != "2" {
			return badRequest("normal priority needs delivery monitoring 2 or none, got %q", monitoring)
		}
	default:
		return badRequest("priority must be %q or %q, got %q", constant.PriorityUrgent, constant.PriorityNormal, priority)
	}
	return nil
}

// validateText checks block 4: at least one line, every character valid. A line that is still
// Vietnamese fails here, which is the point - converting is the caller's explicit step.
func validateText(lines []string) (err error) {
	if len(lines) == 0 {
		return badRequest("text block must have at least one line")
	}
	for index, line := range lines {
		if err = iso15022.ValidateCharset(line); err != nil {
			return badRequest("text line %d: %v", index+1, err)
		}
	}
	return nil
}

// badRequest builds the ErrBadRequest error every validation failure returns.
func badRequest(format string, args ...any) (err error) {
	return model.NewError(model.ErrBadRequest, fmt.Sprintf(format, args...))
}

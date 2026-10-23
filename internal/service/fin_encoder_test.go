package service

import (
	"errors"
	"strings"
	"testing"

	"github.com/Nha1410/VTB-DER/internal/common/iso15022"
	"github.com/Nha1410/VTB-DER/internal/model"
)

// validMessage returns a message that encodes cleanly; each case changes one field.
func validMessage() (msg model.FINMessage) {
	return model.FINMessage{
		SenderBIC:      "VSDCTBVN",
		SessionNumber:  20,
		SequenceNumber: 21,
		MessageType:    "598",
		Priority:       "N",
		Text:           []string{":20:REF0001", ":12:600"},
	}
}

func TestEncodeFIN_FramesLikeTheSpecExample(t *testing.T) {
	// Spec example: {1:F01VSDCTBVNAXXX0020000021}{2:I500VSDSVN06XXXXN}{4:...-}{5:{MAC:00000000}{CHK:F1DBCA886BBF}{TNG:}}
	inputMsg := validMessage()
	inputMsg.MessageType = "500"

	gotFIN, err := EncodeFIN(inputMsg)
	if err != nil {
		t.Fatalf("EncodeFIN() error: %v", err)
	}

	wantFIN := "{1:F01VSDCTBVNAXXX0020000021}{2:I500VSDSVN06XXXXN}{4:\r\n:20:REF0001\r\n:12:600\r\n-}" +
		"{5:{MAC:00000000}{CHK:F1DBCA886BBF}{TNG:}}"
	if string(gotFIN) != wantFIN {
		t.Errorf("EncodeFIN() =\n%q\nwant\n%q", gotFIN, wantFIN)
	}
}

func TestEncodeFIN_Block2Options(t *testing.T) {
	tests := []struct {
		name           string
		priority       string
		monitoring     string
		obsolescence   string
		receiverBIC    string
		wantApplicHead string
	}{
		{name: "urgent with monitoring 1", priority: "U", monitoring: "1", wantApplicHead: "{2:I598VSDSVN06XXXXU1}"},
		{name: "urgent with monitoring 3 and obsolescence", priority: "U", monitoring: "3", obsolescence: "020", wantApplicHead: "{2:I598VSDSVN06XXXXU3020}"},
		{name: "normal with monitoring 2", priority: "N", monitoring: "2", wantApplicHead: "{2:I598VSDSVN06XXXXN2}"},
		{name: "explicit receiver", priority: "N", receiverBIC: "VSDHSCXX", wantApplicHead: "{2:I598VSDHSCXXXXXXN}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputMsg := validMessage()
			inputMsg.Priority, inputMsg.DeliveryMonitoring = tt.priority, tt.monitoring
			inputMsg.ObsolescencePeriod, inputMsg.ReceiverBIC = tt.obsolescence, tt.receiverBIC

			gotFIN, err := EncodeFIN(inputMsg)
			if err != nil {
				t.Fatalf("EncodeFIN() error: %v", err)
			}
			if !strings.Contains(string(gotFIN), tt.wantApplicHead) {
				t.Errorf("EncodeFIN() = %q, want it to contain %q", gotFIN, tt.wantApplicHead)
			}
		})
	}
}

func TestEncodeFIN_RejectsInvalidMessage(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(msg *model.FINMessage)
	}{
		{name: "short sender BIC", mutate: func(msg *model.FINMessage) { msg.SenderBIC = "VSD" }},
		{name: "lower-case sender BIC", mutate: func(msg *model.FINMessage) { msg.SenderBIC = "vsdctbvn" }},
		{name: "bad receiver BIC", mutate: func(msg *model.FINMessage) { msg.ReceiverBIC = "X" }},
		{name: "message type not 3 digits", mutate: func(msg *model.FINMessage) { msg.MessageType = "59" }},
		{name: "session too large", mutate: func(msg *model.FINMessage) { msg.SessionNumber = 10000 }},
		{name: "negative sequence", mutate: func(msg *model.FINMessage) { msg.SequenceNumber = -1 }},
		{name: "sequence too large", mutate: func(msg *model.FINMessage) { msg.SequenceNumber = 1000000 }},
		{name: "unknown priority", mutate: func(msg *model.FINMessage) { msg.Priority = "X" }},
		{name: "urgent without monitoring", mutate: func(msg *model.FINMessage) { msg.Priority = "U" }},
		{name: "normal with monitoring 1", mutate: func(msg *model.FINMessage) { msg.DeliveryMonitoring = "1" }},
		{name: "bad obsolescence", mutate: func(msg *model.FINMessage) { msg.ObsolescencePeriod = "030" }},
		{name: "empty text", mutate: func(msg *model.FINMessage) { msg.Text = nil }},
		{name: "unconverted vietnamese", mutate: func(msg *model.FINMessage) { msg.Text = []string{":95Q::INVE//Nguyễn"} }},
		{name: "brace in text", mutate: func(msg *model.FINMessage) { msg.Text = []string{":20:A}B"} }},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputMsg := validMessage()
			tt.mutate(&inputMsg)

			gotFIN, err := EncodeFIN(inputMsg)

			var customErr *model.CustomError
			if !errors.As(err, &customErr) || customErr.Code != model.ErrBadRequest {
				t.Fatalf("EncodeFIN() error = %v, want ErrBadRequest", err)
			}
			if gotFIN != nil {
				t.Errorf("EncodeFIN() returned bytes %q alongside an error", gotFIN)
			}
		})
	}
}

func TestEncodeFIN_AcceptsConvertedVietnamese(t *testing.T) {
	convertedName, err := iso15022.Encode("NGUYỄN VĂN A / B")
	if err != nil {
		t.Fatalf("Encode() error: %v", err)
	}
	inputMsg := validMessage()
	inputMsg.Text = []string{":95Q::INVE//" + convertedName}

	gotFIN, err := EncodeFIN(inputMsg)
	if err != nil {
		t.Fatalf("EncodeFIN() error: %v", err)
	}
	if !strings.Contains(string(gotFIN), ":95Q::INVE//NGUY?EEX?N V?AW?N A ?_? B") {
		t.Errorf("EncodeFIN() = %q, want the converted name in block 4", gotFIN)
	}
}

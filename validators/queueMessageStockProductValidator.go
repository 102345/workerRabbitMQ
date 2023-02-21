package validators

import (
	"fmt"
	"strings"

	valid "github.com/asaskevich/govalidator"
)

// Const referente as regras da mensageria
const (
	MessageMaximumMessageSizeInvalid = "Message size of one of the invalid bits"
	MessageBitLetterInvalid          = "Message contains one of the bits with letter(s)"
	MessageFieldNotNumeric           = "Message contains one of the non-numeric fields"
	MaximumMessageSize               = 31
)

func ValidateMessageStockProduct(message string) (string, bool) {

	messageRet := ""

	if len(message) != MaximumMessageSize {
		messageRet = MessageMaximumMessageSizeInvalid
		return messageRet, false
	}

	if strings.ContainsAny(message, "abcdefghijklmoqrstuvxzABCDEFGHIJKLMOQRTUVXZ") {
		messageRet = MessageBitLetterInvalid
		return messageRet, false
	}

	fields := strings.SplitN(message, ":", -1)
	cont := 0
	for _, field := range fields {
		fmt.Println(field)
		if cont <= 2 {
			check := valid.IsFloat(field)
			if !check {
				messageRet = MessageFieldNotNumeric
				return messageRet, false
			}
		}
		cont++
	}

	return messageRet, true

}

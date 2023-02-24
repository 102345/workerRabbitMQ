package validators

import (
	"fmt"
	"strconv"
	"strings"

	valid "github.com/asaskevich/govalidator"
	"github.com/marc/workerRabbitMQ-example/core/domain"
	"github.com/marc/workerRabbitMQ-example/core/dto"
)

// Const referente as regras da mensageria
const (
	MessageMaximumMessageSizeInvalid = "Message size of one of the invalid bits"
	MessageBitLetterInvalid          = "Message contains one of the bits with letter(s)"
	MessageFieldNotNumeric           = "Message contains one of the non-numeric fields"
	MessageCodeProductInvalid        = "Message contains invalid product code"
	MaximumMessageSize               = 31
)

func ValidateMessageStockProduct(message string, productUseCase domain.IProductUseCase) (string, dto.StockProductDTO) {

	messageRet := ""

	if len(message) != MaximumMessageSize {
		messageRet = MessageMaximumMessageSizeInvalid
		return messageRet, dto.StockProductDTO{}
	}

	if strings.ContainsAny(message, "abcdefghijklmoqrstuvxzABCDEFGHIJKLMOQRTUVXZ") {
		messageRet = MessageBitLetterInvalid
		return messageRet, dto.StockProductDTO{}
	}

	fields := strings.SplitN(message, ":", -1)
	cont := 0
	stockProductDTO := dto.StockProductDTO{}

	for _, field := range fields {
		fmt.Println(field)
		if cont <= 2 {
			check := valid.IsFloat(field)
			if !check {
				messageRet = MessageFieldNotNumeric
				return messageRet, dto.StockProductDTO{}
			}
		}
		if cont == 0 {
			productId, _ := strconv.ParseInt(field, 10, 32)
			stockProductDTO.ProductID = int32(productId)

			product, _ := productUseCase.FindById(productId)

			if product.ID == 0 {
				messageRet = MessageCodeProductInvalid
				return messageRet, dto.StockProductDTO{}
			}

		} else if cont == 1 {
			quantity, _ := strconv.ParseInt(field, 10, 32)
			stockProductDTO.Quantity = int32(quantity)
		} else if cont == 2 {
			balance, _ := strconv.ParseInt(field, 10, 32)
			stockProductDTO.Balance = int32(balance)
		} else {
			if field == "N" {
				stockProductDTO.Balance = stockProductDTO.Balance * (-1)
			}
		}
		cont++
	}

	return "", stockProductDTO

}

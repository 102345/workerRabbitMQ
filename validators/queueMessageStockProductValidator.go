package validators

import (
	"strconv"
	"strings"

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

	var messageRet string
	var stockProductDTO dto.StockProductDTO

	if len(message) != MaximumMessageSize {
		messageRet = MessageMaximumMessageSizeInvalid
		return messageRet, dto.StockProductDTO{}
	}

	if strings.ContainsAny(message, "abcdefghijklmoqrstuvxzABCDEFGHIJKLMOQRTUVXZ") {
		messageRet = MessageBitLetterInvalid
		return messageRet, stockProductDTO
	}

	fields := strings.SplitN(message, ":", -1)

	productID, err := strconv.ParseInt(fields[0], 10, 32)
	if err != nil {
		return MessageFieldNotNumeric, stockProductDTO
	}

	product, err := productUseCase.FindById(productID)
	if err != nil {
		return MessageCodeProductInvalid, stockProductDTO
	}

	if product.ID == 0 {
		return MessageCodeProductInvalid, stockProductDTO
	}

	quantity, err := strconv.Atoi(fields[1])
	if err != nil {
		return MessageFieldNotNumeric, stockProductDTO
	}

	balance, err := strconv.Atoi(fields[2])
	if err != nil {
		return MessageFieldNotNumeric, stockProductDTO
	}

	if len(fields) > 3 && strings.ToUpper(fields[3]) == "N" {
		balance = -balance
	}

	stockProductDTO.ProductID = int32(productID)
	stockProductDTO.Quantity = int32(quantity)
	stockProductDTO.Balance = int32(balance)

	return messageRet, stockProductDTO

}

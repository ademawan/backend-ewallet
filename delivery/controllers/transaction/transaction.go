package transaction

import (
	"backend-ewallet/delivery/controllers/common"
	"backend-ewallet/entities"
	"backend-ewallet/middlewares"
	"backend-ewallet/repository/transaction"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TransactionController struct {
	repo transaction.Transaction
}

func New(repository transaction.Transaction) *TransactionController {
	return &TransactionController{
		repo: repository,
	}
}

func (ac *TransactionController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		transaction := CreateTransactionRequestFormat{}

		c.Bind(&transaction)
		err := c.Validate(&transaction)

		if err != nil {
			return c.JSON(http.StatusBadRequest, common.ResponseUser(http.StatusBadRequest, "There is some problem from input", nil))
		}

		res, err_repo := ac.repo.Create(entities.Transaction{
			SenderID:        transaction.SenderID,
			RecipientID:     transaction.RecipientID,
			RecievedAmount:  transaction.RecievedAmount,
			SentAmount:      transaction.SentAmount,
			TransactionType: transaction.TransactionType,
		})

		if err_repo != nil {
			return c.JSON(http.StatusConflict, common.ResponseUser(http.StatusConflict, err_repo.Error(), nil))
		}

		response := TransactionCreateResponse{}
		response.TransactionID = res.TransactionID
		response.SenderID = res.SenderID
		response.RecipientID = res.RecipientID
		response.RecievedAmount = res.RecievedAmount
		response.SentAmount = res.SentAmount
		response.TransactionType = res.TransactionType

		return c.JSON(http.StatusCreated, common.ResponseUser(http.StatusCreated, "Success create transaction", response))

	}
}

func (ac *TransactionController) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		transactionID := middlewares.ExtractTokenUserID(c)

		res, err := ac.repo.GetByID(transactionID)

		if err != nil {
			statusCode := http.StatusInternalServerError
			errorMessage := "There is some problem from the server"
			if err.Error() == "record not found" {
				statusCode = http.StatusNotFound
				errorMessage = err.Error()
			}
			return c.JSON(statusCode, common.ResponseUser(http.StatusNotFound, errorMessage, nil))
		}

		return c.JSON(http.StatusOK, common.ResponseUser(http.StatusOK, "Success get transaction", res))
	}
}

func (ac *TransactionController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		transactionID := middlewares.ExtractTokenUserID(c)
		var newTransaction = UpdateTransactionRequestFormat{}
		c.Bind(&newTransaction)

		err := c.Validate(&newTransaction)
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.ResponseUser(http.StatusBadRequest, "There is some problem from input", nil))
		}

		res, err_repo := ac.repo.Update(transactionID, entities.Transaction{
			SenderID:        newTransaction.SenderID,
			RecipientID:     newTransaction.RecipientID,
			RecievedAmount:  newTransaction.RecievedAmount,
			SentAmount:      newTransaction.SentAmount,
			TransactionType: newTransaction.TransactionType,
		})

		if err_repo != nil {
			return c.JSON(http.StatusInternalServerError, common.ResponseUser(http.StatusInternalServerError, "There is some error on server", nil))
		}

		response := TransactionUpdateResponse{}
		response.TransactionID = res.TransactionID
		response.SenderID = res.SenderID
		response.RecipientID = res.RecipientID
		response.RecievedAmount = res.RecievedAmount
		response.SentAmount = res.SentAmount
		response.TransactionType = res.TransactionType

		return c.JSON(http.StatusOK, common.ResponseUser(http.StatusOK, "Success update transaction", response))
	}
}

func (ac *TransactionController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		transactionID := middlewares.ExtractTokenUserID(c)
		err := ac.repo.Delete(transactionID)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.ResponseUser(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusOK, common.ResponseUser(http.StatusOK, "Success delete transaction", nil))
	}
}

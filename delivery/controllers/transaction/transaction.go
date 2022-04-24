package transaction

import (
	"backend-ewallet/delivery/controllers/common"
	"backend-ewallet/entities"
	"backend-ewallet/middlewares"
	"backend-ewallet/repository/transaction"
	"backend-ewallet/repository/user"
	"backend-ewallet/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lithammer/shortuuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type TransactionController struct {
	repo     transaction.Transaction
	mt       coreapi.Client
	userRepo user.User
}

func New(repository transaction.Transaction, mt coreapi.Client, userRepo user.User) *TransactionController {
	return &TransactionController{
		repo:     repository,
		mt:       mt,
		userRepo: userRepo,
	}
}

func (ac *TransactionController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		transaction := CreateTransactionRequestFormat{}
		userID := middlewares.ExtractTokenUserID(c)

		c.Bind(&transaction)
		err := c.Validate(&transaction)

		if err != nil {
			return c.JSON(http.StatusBadRequest, common.ResponseUser(http.StatusBadRequest, "There is some problem from input", nil))
		}
		transaction.SenderID = userID

		res, err_repo := ac.repo.Create(entities.Transaction{
			SenderID:        transaction.SenderID,
			RecipientID:     transaction.RecipientID,
			TransactionType: transaction.TransactionType,
		})

		if err_repo != nil {
			return c.JSON(http.StatusConflict, common.ResponseUser(http.StatusConflict, err_repo.Error(), nil))
		}

		response := TransactionCreateResponse{}
		response.TransactionID = res.TransactionID
		response.SenderID = res.SenderID
		response.RecipientID = res.RecipientID
		response.TransactionType = res.TransactionType

		return c.JSON(http.StatusCreated, common.ResponseUser(http.StatusCreated, "Success create transaction", response))

	}
}

func (ac *TransactionController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := middlewares.ExtractTokenUserID(c)

		res, err := ac.repo.Get(userID)

		if err != nil {
			statusCode := http.StatusInternalServerError
			errorMessage := "There is some problem from the server"
			if err.Error() == "record not found" {
				statusCode = http.StatusNotFound
				errorMessage = err.Error()
			}
			return c.JSON(statusCode, common.ResponseUser(http.StatusNotFound, errorMessage, res))
		}

		return c.JSON(http.StatusOK, common.ResponseUser(http.StatusOK, "Success get transaction", res))
	}
}

func (ac *TransactionController) GetTransactionSend() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := middlewares.ExtractTokenUserID(c)

		res, err := ac.repo.GetTransactionSend(userID)

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
func (ac *TransactionController) GetTransactionReceived() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := middlewares.ExtractTokenUserID(c)

		res, err := ac.repo.GetTransactionReceived(userID)

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

func (ac *TransactionController) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := middlewares.ExtractTokenUserID(c)
		transactionID := c.Param("transaction_id")

		res, err := ac.repo.GetByID(userID, transactionID)

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
func (cont *TransactionController) CreatePayment() echo.HandlerFunc {
	return func(c echo.Context) error {
		var payment PaymentTypeRequest
		userID := middlewares.ExtractTokenUserID(c)

		var result *coreapi.ChargeReq
		resUser, errUser := cont.userRepo.GetByID(userID)
		if errUser != nil {
			return c.JSON(http.StatusInternalServerError, common.ResponseUser(http.StatusInternalServerError, "Your booking is not found", nil))
		}

		uid := shortuuid.New()

		switch payment.Payment_method {
		case "gopay":
			result = &coreapi.ChargeReq{
				PaymentType: coreapi.PaymentTypeGopay,

				TransactionDetails: midtrans.TransactionDetails{
					OrderID:  uid,
					GrossAmt: int64(payment.Amount),
				},
				Items: &[]midtrans.ItemDetails{
					{
						ID:    userID,
						Name:  resUser.Name,
						Price: int64(payment.Amount),
						Qty:   1,
					},
				},
			}

		case "shopeepay":
			result = &coreapi.ChargeReq{
				PaymentType: coreapi.PaymentTypeShopeepay,

				TransactionDetails: midtrans.TransactionDetails{
					OrderID:  uid,
					GrossAmt: int64(payment.Amount),
				},
				Items: &[]midtrans.ItemDetails{
					{
						ID:    userID,
						Name:  resUser.Name,
						Price: int64(payment.Amount),
						Qty:   1,
					},
				},
				CustomerDetails: &midtrans.CustomerDetails{
					FName: "roger",
					LName: "san",
					Email: "dani@gmail.com",
					Phone: "089876543210",
				},
				ShopeePay: &coreapi.ShopeePayDetails{
					CallbackUrl: "https://plastic-cougar-32.loca.lt/booking/payment/callback",
				},
			}
		case "qris":
			result = &coreapi.ChargeReq{
				PaymentType: coreapi.PaymentTypeQris,

				TransactionDetails: midtrans.TransactionDetails{
					OrderID:  uid,
					GrossAmt: int64(payment.Amount),
				},
				Items: &[]midtrans.ItemDetails{
					{
						ID:    userID,
						Name:  resUser.Name,
						Price: int64(payment.Amount),
						Qty:   1,
					},
				},
			}

		}

		apiRes, err := utils.CreateTransaction(cont.mt, result)
		// log.Info(apiRes)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.ResponseUser(http.StatusInternalServerError, "failed topup", nil))

		}

		var data PaymentResponse
		data.OrderID = apiRes.OrderID
		data.GrossAmount = apiRes.GrossAmount
		data.PaymentType = apiRes.PaymentType
		// for i := range apiRes.Actions {
		data.Url = /*  append(data.Url,  */ apiRes.Actions[1].URL /* ) */
		// }

		return c.JSON(http.StatusOK, common.ResponseUser(http.StatusOK, "Success create payment", data))

	}
}
func (cont *TransactionController) CallBack() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request RequestCallBackMidtrans
		// user_uid := middlewares.ExtractTokenId(c)

		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusInternalServerError, common.ResponseUser(http.StatusInternalServerError, "Failed to create payment", nil))
		}
		amount, _ := strconv.Atoi(request.Gross_amount)
		response := entities.Transaction{}
		switch request.Transaction_status {
		case "settlement":
			res, err := cont.repo.Create(entities.Transaction{SenderID: request.Transaction_id, RecipientID: request.Order_id, TransactionType: "topup", Amount: uint(amount)})
			if err != nil {
				return c.JSON(http.StatusInternalServerError, common.ResponseUser(http.StatusInternalServerError, "try again later", nil))

			}
			response = res
		}

		// var strDebug string
		// strDebug = spew.Sdump(request)
		// ZapLogger.Info(`request: ` + strDebug)

		return c.JSON(http.StatusOK, common.ResponseUser(http.StatusOK, "Success topup", response))

	}
}

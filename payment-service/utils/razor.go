package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"errors"

	"github.com/razorpay/razorpay-go"
	"github.com/spf13/viper"
)

func Executerazorpay(Type string, amount int) (string, error) {

	client := razorpay.NewClient(viper.GetString("RAZOR_KEY"), viper.GetString("RAZOR_SECRET"))

	data := map[string]interface{}{
		"amount":   amount * 100,
		"currency": "INR",
		"receipt":  Type,
	}

	body, err := client.Order.Create(data, nil)
	if err != nil {
		return "", errors.New("payment not initiated")
	}
	razorId, _ := body["id"].(string)
	return razorId, nil
}

func RazorPaymentVerification(sign, orderId, paymentId string) error {
	secret := viper.GetString("RAZOR_SECRET")
	data := orderId + "|" + paymentId

	h := hmac.New(sha256.New, []byte(secret))

	_, err := h.Write([]byte(data))
	if err != nil {
		panic(err)
	}

	sha := hex.EncodeToString(h.Sum(nil))
	if subtle.ConstantTimeCompare([]byte(sha), []byte(sign)) != 1 {
		return errors.New("payment failed")
	} else {
		return nil
	}
}

package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	payment_service "github.com/rishad004/Gv_protofiles/payment"
	"github.com/rishad004/project-gv/apiGateway/inertnal/domain"
	"github.com/rishad004/project-gv/apiGateway/utils"
	"github.com/spf13/viper"
)

func (h *ApiHanlder) PaymentRendering(w http.ResponseWriter, r *http.Request) {

	orderId := r.URL.Query().Get("id")

	utils.RenderTemplate(w, "razor.html", map[string]any{
		"Order": orderId,
		"Key":   viper.GetString("RAZOR_KEY"),
	})
}

func (h *ApiHanlder) PaymentVerifying(w http.ResponseWriter, r *http.Request) {
	var razor domain.Razor

	if err := json.NewDecoder(r.Body).Decode(&razor); err != nil {
		utils.SendJSONResponse(w, "Invalid payload!", http.StatusBadRequest, r)
		return
	}

	fmt.Println(razor.Signature, "--", razor.Order, "--", razor.Payment)

	if _, err := h.PaymentPb.PaymentVerify(context.Background(), &payment_service.PaymentVerifyRequest{
		Id:        int32(r.Context().Value("Id").(uint)),
		Signature: razor.Signature,
		Orderid:   razor.Order,
		Paymentid: razor.Payment,
	}); err != nil {
		utils.SendJSONResponse(w, err.Error(), http.StatusBadRequest, r)
		return
	}

	utils.SendJSONResponse(w, "Subscribed the channel successfully!", http.StatusOK, r)
}

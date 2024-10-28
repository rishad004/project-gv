package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"

	"github.com/rishad004/project-gv/apiGateway/inertnal/domain"
)

func SendJSONResponse(w http.ResponseWriter, message any, statusCode int, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)

	response := domain.Response{
		Message:    message,
		StatusCode: statusCode,
	}
	json.NewEncoder(w).Encode(response)
}

func RenderTemplate(w http.ResponseWriter, tmpl string, data map[string]any) {

	parsedTemplate, err := template.ParseFiles("pkg/templates/" + tmpl)
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}
	err = parsedTemplate.Execute(w, data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}

}

func PublicIpFinder() (error, string) {
	url := "https://api64.ipify.org"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return err, ""
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return err, ""
	}

	publicIP := string(body)
	return nil, publicIP
}

func SetCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

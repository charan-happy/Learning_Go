// building stripe payment backend

package main // calling the main package
//import the necessary packages
import (
	"encoding/json"
	"fmt"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/paymentintent"
	"io"
	"bytes"
	"log"
	"net/http"
	

)
// main function declaration
type Struct struct {
		ProductId string `json:"product_id"`
		FirstName string `json:"first_name"`
		LastName string `json:"last_name"`
		Address1 string `json:"address_1"`
		Address2 string `json:"address_2"`
		City string `json:"city"`
		State string `json:"state"`
		Zip string `json:"zip"`
		Country string `json:"country"`
}
func main() {
	stripe.Key = "sk_test_51Pe9I42NHGmeKCM0l0TYeDEPBXQ1TxnJsC3rPqvKPbef4lvZ14DNySNB8ml2Y2EYWCxO9yhd9lhY24sgChab9Amo00E47BX6f0"
    http.HandleFunc("/create-payment-intent", handleCreatePaymentIntent)
    http.HandleFunc("/health", handleHealth)


	log.Println("Listening on port 4242...") 
	var err = http.ListenAndServe("localhost:4242", nil)
	if err != nil { // handling error
		log.Fatal(err)
	}
}

func handleCreatePaymentIntent(writer http.ResponseWriter, request *http.Request) {
    // validating request method
	if request.Method != "POST" {
     http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	 return
  }

  var req Struct
  err := json.NewDecoder(request.Body).Decode(&req)
  if err != nil {
	log.Println(err)
	return
  }

  params := &stripe.PaymentIntentParams{
	Amount: stripe.Int64(calculateOrderAmount(req.ProductId)),
	Currency: stripe.String(string(stripe.CurrencyUSD)),
	AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
		Enabled: stripe.Bool(true),
	},
  }
  paymentIntent, err := paymentintent.New(params)
  if err != nil {
	http.Error(writer, err.Error(), http.StatusInternalServerError)
  }
  fmt.Println(paymentIntent.ClientSecret)

  var response struct {
      ClientSecret string `json:"clientSecret"`
  }

  response.ClientSecret = paymentIntent.ClientSecret
 
  var buf bytes.Buffer 
  err = json.NewEncoder(&buf).Encode(response)
  if err != nil {
	http.Error(writer, err.Error(), http.StatusInternalServerError)
  }
  writer.Header().Set("Content-Type", "application/json")
  
  _, err = io.Copy(writer, &buf)
  if err != nil {
	fmt.Println(err)
  }

}

func handleHealth(writer http.ResponseWriter, r *http.Request) {
	response := []byte("Server is up and running .....")

	_, err := writer.Write(response)
	if err != nil {
		fmt.Println(err)
	}
}

func calculateOrderAmount(ProductId string) int64 {
	 switch ProductId {
	 case "Forever Pants":
		return 2600
	 case "Forever Shirts":
        return 1550
	 case "Forever Tshirts":
	    return 3000
	 }
	 return 0
}

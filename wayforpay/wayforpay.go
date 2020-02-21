package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type params struct {
	MerchantAccount               string   `json:"merchantAccount"`
	MerchantDomainName            string   `json:"merchantDomainName"`
	MerchantTransactionSecureType string   `json:"merchantTransactionSecureType"`
	OrderReference                string   `json:"orderReference"`
	OrderDate                     string   `json:"orderDate"`
	Amount                        string   `json:"amount"`
	Currency                      string   `json:"currency"`
	ProductName                   []string `json:"productName"`
	ProductPrice                  []string `json:"productPrice"`
	ProductCount                  []string `json:"productCount"`
	ApiVersion                    string   `json:"apiVersion"`
	MerchantSignature             string   `json:"merchantSignature"`
	TransactionType               string   `json:"transactionType"`
}

func ValidMAC(message, messageMAC, key []byte) bool {
	mac := hmac.New(md5.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(messageMAC, expectedMAC)
}

type response struct {
	MerchantAccount   string `json:"merchantAccount"`
	OrderReference    string `json:"orderReference"`
	MerchantSignature string `json:"merchantSignature"`
	Amount            string `json:"amount"`
	Currency          string `json:"currency"`
	AuthCode          string `json:"authCode"`
	Email             string `json:"email"`
	Phone             string `json:"phone"`
	CreatedDate       string `json:"createdDate"`
	ProcessingDate    string `json:"processingDate"`
	CardPan           string `json:"cardPan"`
	CardType          string `json:"cardType"`
	IssuerBankCountry string `json:"issuerBankCountry"`
	IssuerBankName    string `json:"issuerBankName"`
	RecToken          string `json:"recToken"`
	TransactionStatus string `json:"transactionStatus"`
	Reason            string `json:"reason"`
	ReasonCode        string `json:"reasonCode"`
	Fee               string `json:"fee"`
	PaymentSystem     string `json:"paymentSystem"`
}

type requestStruct struct {
	MerchantAccount    string   `json:"merchantAccount"`
	MerchantDomainName string   `json:"merchantDomainName"`
	OrderReference     string   `json:"orderReference"`
	OrderDate          int64    `json:"orderDate"`
	Amount             string   `json:"amount"`
	Currency           string   `json:"currency"`
	ProductName        []string `json:"productName"`
	ProductCount       []string `json:"productCount"`
	ProductPrice       []string `json:"productPrice"`
	MerchantSignature  string   `json:"merchantSignature"`
	TransactionType    string   `json:"transactionType"`
	APIVersion         string   `json:"apiVersion"`
	ServiceURL         string   `json:"serviceUrl"`
}

func (a requestStruct) buildString() string {
	values := []string{}
	values = append(values, a.MerchantAccount)
	values = append(values, a.MerchantDomainName)
	values = append(values, a.OrderReference)
	values = append(values, strconv.FormatInt(a.OrderDate, 10))
	values = append(values, a.Amount)
	values = append(values, a.Currency)
	values = append(values, a.ProductName...)
	values = append(values, a.ProductCount...)
	values = append(values, a.ProductPrice...)
	result := strings.Join(values, ";")
	fmt.Println(result)
	return result
}

type apiSResponseStruct struct {
	InvoiceURL string `json:"invoiceUrl"`
	Reason     string `json:"reason"`
	ReasonCode string `json:"reasonCode"`
	QRCode     string `json:"qrCode"`
}

type serviceURLRequest struct {
	MerchantAccount   string `json:"merchantAccount"`
	OrderReference    string `json:"orderReference"`
	MerchantSignature string `json:"merchantSignature"`
	Amount            string `json:"amount"`
	Currency          string `json:"currency"`
	AuthCode          string `json:"authCode"`
	Email             string `json:"email"`
	Phone             string `json:"phone"`
	CreatedDate       string `json:"createdDate"`
	ProcessingDate    string `json:"processingDate"`
	CardPan           string `json:"cardPan"`
	CardType          string `json:"cardType"`
	IssuerBankCountry string `json:"issuerBankCountry"`
	IssuerBankName    string `json:"issuerBankName"`
	RecToken          string `json:"recToken"`
	TransactionStatus string `json:"transactionStatus"`
	Reason            string `json:"reason"`
	ReasonCode        string `json:"reasonCode"`
	Fee               string `json:"fee"`
	PaymentSystem     string `json:"paymentSystem"`
}

func (a requestStruct) generateSignature(key string) string {
	message := a.buildString()
	mac := hmac.New(md5.New, []byte(key))
	mac.Write([]byte(message))

	return hex.EncodeToString(mac.Sum(nil))
}

func makeTimestamp() int64 {
	now := time.Now().Local().Add(time.Hour * time.Duration(1))
	nanos := now.UnixNano()
	secs := nanos / 1000000000
	return secs
}
func main() {
	key := "flk3409refn54t54t*FNJRET"
	products := []string{"Процессор Intel Core i5-4670 3.4GHz"}
	as := requestStruct{
		MerchantAccount:    "test_merch_n1",
		MerchantDomainName: "marketplace.horobets.me",
		OrderReference:     "DH783023dsxsazdsa1sfd",
		OrderDate:          makeTimestamp(),
		Amount:             "1547.36",
		Currency:           "UAH",
		ProductName:        products,
		ProductCount:       []string{"1"},
		ProductPrice:       []string{"1000"},
		TransactionType:    "CREATE_INVOICE",
		APIVersion:         "1",
		ServiceURL:         "http://dream.market/",
	}
	sign := as.generateSignature(key)
	fmt.Println(sign)
	fmt.Println(as.OrderDate)
	as.MerchantSignature = sign
	APIURL := "https://api.wayforpay.com/api"
	jsonStr, err := json.Marshal(as)

	if err != nil {
		fmt.Println("error:", err)
	}
	req, err := http.NewRequest("POST", APIURL, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)

	var response apiSResponseStruct
	err = json.Unmarshal(body, &response)
	fmt.Println(response)
	fmt.Println("response Body:", string(body))
}

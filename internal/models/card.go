package models

type UpdatePriorityRequest struct {
	CardNumber     string `json:"cardNumber"`
	CardExpiry     string `json:"cardExpiry"`
	CardCvc        string `json:"cardCVC"`
	DonationAmount string `json:"donationAmount"`
}

type Rating struct {
	Rating int64 `json:"rating"`
}

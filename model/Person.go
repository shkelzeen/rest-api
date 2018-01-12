package model

type Person struct {
	Name       string `json:"name,omitempty"`
	Email      string `json:"email,omitempty"`
	City       string `json:"city,omitempty"`
	Mac        string `json:"mac,omitempty"`
	Timestamp  string `json:"timestamp,omitempty"`
	Creditcard string `json:"creditcard,omitempty"`
}

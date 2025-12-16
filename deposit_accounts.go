package gotropipay

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// CountryDestination represents destination country details
type CountryDestination struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	SepaZone    bool   `json:"sepaZone"`
	Slug        string `json:"slug"`
	CallingCode int    `json:"callingCode"`
}

// AllowedAccount represents an account allowed for the beneficiary
type AllowedAccount struct {
	ID       int    `json:"id"`
	Alias    string `json:"alias"`
	Currency string `json:"currency"`
	Type     int    `json:"type"`
}

// DepositAccount represents a beneficiary (deposit account)
type DepositAccount struct {
	ID                   int                 `json:"id"`
	AccountNumber        string              `json:"accountNumber"`
	FirstName            string              `json:"firstName"`
	LastName             string              `json:"lastName"`
	Alias                string              `json:"alias"`
	Swift                string              `json:"swift"`
	Type                 int                 `json:"type"`
	PersonType           int                 `json:"personType"`
	State                interface{}         `json:"state"` // Can be string "active" or int 0
	CountryDestinationID int                 `json:"countryDestinationId"`
	DocumentNumber       string              `json:"documentNumber"`
	Address              string              `json:"address"`
	Phone                string              `json:"phone"`
	Email                string              `json:"email"`
	CreatedAt            string              `json:"createdAt"`
	UpdatedAt            string              `json:"updatedAt"`
	CountryDestination   *CountryDestination `json:"countryDestination,omitempty"`
	PaymentMethods       []string            `json:"paymentMethods,omitempty"`
	AllowedAccounts      []AllowedAccount    `json:"allowedAccounts,omitempty"`
	Allowed              bool                `json:"allowed"`
}

// CreateDepositAccountRequest represents payload to create a beneficiary
type CreateDepositAccountRequest struct {
	AccountNumber        string `json:"accountNumber"`
	FirstName            string `json:"firstName"`
	LastName             string `json:"lastName"`
	CountryDestinationID int    `json:"countryDestinationId"`
	Type                 int    `json:"type"`
	Alias                string `json:"alias,omitempty"`
	Email                string `json:"email,omitempty"`
	Phone                string `json:"phone,omitempty"`
	Address              string `json:"address,omitempty"`
	Swift                string `json:"swift,omitempty"`
}

// UpdateDepositAccountRequest represents payload to update a beneficiary
type UpdateDepositAccountRequest struct {
	ID    int    `json:"id"`
	Alias string `json:"alias"`
}

// DeleteDepositAccountRequest represents payload to delete a beneficiary
type DeleteDepositAccountRequest struct {
	SecurityCode string `json:"securityCode"`
}

// ValidateAccountNumberRequest represents payload to validate account
type ValidateAccountNumberRequest struct {
	AccountNumber        string `json:"accountNumber"`
	CountryDestinationID int    `json:"countryDestinationId"`
	Type                 int    `json:"type"`
	Currency             string `json:"currency"`
	PaymentType          int    `json:"paymentType"`
}

// ValidateAccountNumberResponse represents response for account validation
type ValidateAccountNumberResponse struct {
	Valid        bool        `json:"valid"`
	Type         interface{} `json:"type"`         // May be null
	ErrorCode    interface{} `json:"errorCode"`    // May be null
	ErrorMessage interface{} `json:"errorMessage"` // May be null
}

// listDepositAccountsResponse used for unmarshalling the list response
type listDepositAccountsResponse struct {
	Items []DepositAccount `json:"items"`
}

// CreateDepositAccount creates a new beneficiary record
func (c *Client) CreateDepositAccount(ctx context.Context, req CreateDepositAccountRequest) (*DepositAccount, error) {
	var resp DepositAccount
	err := c.Request(ctx, "POST", "/depositaccounts/", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListDepositAccounts retrieves a list of beneficiaries
func (c *Client) ListDepositAccounts(ctx context.Context, limit, offset int, search string) ([]DepositAccount, error) {
	params := url.Values{}
	if limit > 0 {
		params.Add("limit", strconv.Itoa(limit))
	}
	if offset > 0 {
		params.Add("offset", strconv.Itoa(offset))
	}
	if search != "" {
		params.Add("search", search)
	}

	path := "/depositaccounts/"
	if len(params) > 0 {
		path += "?" + params.Encode()
	}

	var resp listDepositAccountsResponse
	err := c.Request(ctx, "GET", path, nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Items, nil
}

// GetDepositAccount retrieves details of a single beneficiary
func (c *Client) GetDepositAccount(ctx context.Context, id int) (*DepositAccount, error) {
	var resp DepositAccount
	path := fmt.Sprintf("/depositaccounts/%d", id)
	err := c.Request(ctx, "GET", path, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateDepositAccount updates a beneficiary alias
func (c *Client) UpdateDepositAccount(ctx context.Context, req UpdateDepositAccountRequest) (*DepositAccount, error) {
	var resp DepositAccount
	err := c.Request(ctx, "PUT", "/depositaccounts/", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteDepositAccount deletes a beneficiary
func (c *Client) DeleteDepositAccount(ctx context.Context, id int, securityCode string) error {
	path := fmt.Sprintf("/depositaccounts/%d", id)
	req := DeleteDepositAccountRequest{SecurityCode: securityCode}
	return c.Request(ctx, "DELETE", path, req, nil)
}

// ValidateAccountNumber checks account format and existence
func (c *Client) ValidateAccountNumber(ctx context.Context, req ValidateAccountNumberRequest) (*ValidateAccountNumberResponse, error) {
	var resp ValidateAccountNumberResponse
	err := c.Request(ctx, "POST", "/depositaccounts/validateaccountnumber", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

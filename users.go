package gotropipay

import (
	"context"
)

// User represents a user profile in Tropipay
type User struct {
	ID         string                 `json:"id"`
	Name       string                 `json:"name"`
	Surname    string                 `json:"surname"`
	Email      string                 `json:"email"`
	Phone      string                 `json:"phone"`
	State      int                    `json:"state"`
	KycLevel   int                    `json:"kycLevel"`
	Balance    int64                  `json:"balance"` // In cents
	PendingIn  int64                  `json:"pendingIn"`
	PendingOut int64                  `json:"pendingOut"`
	TwoFaMode  int                    `json:"twoFaMode"`
	Logo       string                 `json:"logo"`
	CreatedAt  string                 `json:"createdAt"`
	UpdatedAt  string                 `json:"updatedAt"`
	Group      map[string]interface{} `json:"group,omitempty"`
	UserDetail map[string]interface{} `json:"userDetail,omitempty"`
	Options    map[string]interface{} `json:"options,omitempty"`
}

// SendSecurityCodeRequest represents the payload to send a security code
type SendSecurityCodeRequest struct {
	Type        string `json:"type"`                  // "sms" or "email"
	CallingCode string `json:"callingCode,omitempty"` // Required when type is sms
	Phone       string `json:"phone,omitempty"`       // Required when type is sms
	Email       string `json:"email,omitempty"`       // Required when type is email
}

// ValidateSecurityTokenRequest represents the payload to validate a token
type ValidateSecurityTokenRequest struct {
	SecurityCode string `json:"securityCode"`
	Type         string `json:"type"` // "sms", "email", "totp"
}

// ValidateSecurityTokenResponse is the response from validating a token
type ValidateSecurityTokenResponse struct {
	IsValid bool   `json:"isValid"`
	User    User   `json:"user"`
	Token   string `json:"token"`
}

// Configure2FARequest represents the payload to configure 2FA
type Configure2FARequest struct {
	Enabled      bool   `json:"enabled"`
	Type         string `json:"type"` // "totp", "sms"
	SecurityCode string `json:"securityCode"`
}

// Get2FASecretResponse contains the secret and QR code for TOTP
type Get2FASecretResponse struct {
	Secret    string `json:"secret"`
	QRCodeURL string `json:"qrCodeUrl"`
}

// ChangePasswordRequest represents the payload to change password
type ChangePasswordRequest struct {
	OldPass string `json:"oldPass"`
	NewPass string `json:"newPass"`
}

// DisableUserResponse represents the response when disabling a user
type DisableUserResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// GetUserProfile retrieves the details of the authenticated user.
func (c *Client) GetUserProfile(ctx context.Context) (*User, error) {
	var user User
	err := c.Request(ctx, "GET", "/users/profile", nil, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// SendSecurityCode sends a security code to the user's phone or email.
func (c *Client) SendSecurityCode(ctx context.Context, req SendSecurityCodeRequest) error {
	return c.Request(ctx, "POST", "/users/sendSecurityCode", req, nil)
}

// ValidateSecurityToken validates a security code sent to the user.
func (c *Client) ValidateSecurityToken(ctx context.Context, req ValidateSecurityTokenRequest) (*ValidateSecurityTokenResponse, error) {
	var resp ValidateSecurityTokenResponse
	err := c.Request(ctx, "POST", "/users/validateToken", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Configure2FA enables or disables two-factor authentication.
func (c *Client) Configure2FA(ctx context.Context, req Configure2FARequest) error {
	return c.Request(ctx, "POST", "/users/2fa", req, nil)
}

// Get2FASecret generates a new TOTP secret for setting up 2FA.
func (c *Client) Get2FASecret(ctx context.Context) (*Get2FASecretResponse, error) {
	var resp Get2FASecretResponse
	err := c.Request(ctx, "POST", "/users/2fa/secret", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangePassword changes the user's account password.
func (c *Client) ChangePassword(ctx context.Context, req ChangePasswordRequest) error {
	return c.Request(ctx, "POST", "/users/pass", req, nil)
}

// DisableUserAccount disables the user account.
func (c *Client) DisableUserAccount(ctx context.Context) (*DisableUserResponse, error) {
	var resp DisableUserResponse
	err := c.Request(ctx, "POST", "/users/disable", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

package webauthn

type InitialLoginResponse struct {
	Challenge string `json:"challenge"`
}
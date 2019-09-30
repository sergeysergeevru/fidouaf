package webauthn


type PublicKeyCredentialCreationOptions struct {
	Challenge Challenge
	Rp RP
	PubKeyCredParams []PubKeyCredParams
}

type RP struct {
	Id string
	Icon string
	Name string
}

type Challenge string

type User struct {
	DisplayName string
	Icon string
	Id []uint8
	Name string
}

type PubKeyCredParams struct {
	Type string // "public-key"
	Alg string //
}
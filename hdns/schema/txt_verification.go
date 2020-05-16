package schema

// TXT verification defines the schema of TXT Verification object.
type TxtVerification struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}

type TxtVerificationResponse struct {
	TxtVerification TxtVerification `json:"txt_verification"`
}

package hdns

// TxtVerification represents TXT Verification of zone.
type TxtVerification struct {
	Name  string
	Token string
}

// TxtVerificationClient is a client for the TxtVerification API.
type TxtVerificationClient struct {
	client *Client
}

package tokens

// TokenCreateResult contains the document structure returned from a Create call.
type TokenCreateResult map[string]interface{}

// TokenID retrieves a token generated by a Create call from an token creation response.
func (r TokenCreateResult) TokenID() (string, error) {
	return "", nil
}
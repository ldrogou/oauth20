package model

import "fmt"

type Oauth struct {
	ID           int64   `db:"id"`
	AccessToken  string  `db:"access_token"`
	TokenType    string  `db:"token_type"`
	ExpiresIN    float64 `db:"expires_in"`
	RefreshToken string  `db:"refresh_token"`
	Param        Param   `db:"param_id"`
}

func (o Oauth) String() string {
	return fmt.Sprintf("id=%v, accessToken=%v, tokenType=%v, expiresIN=%v, refreshToken=%v",
		o.ID, o.AccessToken, o.TokenType, o.ExpiresIN, o.RefreshToken)
}

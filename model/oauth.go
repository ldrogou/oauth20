package model

import "fmt"

type Oauth struct {
	ID           int64  `db:"id"`
	AccessToken  string `db:"access_token"`
	ExpireIN     int    `db:"expire_in"`
	RefreshToken string `db:"refreh_token"`
}

func (o Oauth) String() string {
	return fmt.Sprintf("id=%v, accessToken=%v, expireIN=%v, refreshToken=%v",
		o.ID, o.AccessToken, o.ExpireIN, o.RefreshToken)
}

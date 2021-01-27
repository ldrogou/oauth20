package model

import "fmt"

type Param struct {
	ID           int64  `db:"id"`
	State        string `db:"state"`
	Domaine      string `db:"domaine"`
	ClientID     string `db:"client_id"`
	ClientSecret string `db:"client_secret"`
	GrantType    string `db:"grant_type"`
}

func (p Param) String() string {
	return fmt.Sprintf("id=%v, state=%v, domaine=%v, clientId=%v, clientSecret=%v, grantType=%v",
		p.ID, p.State, p.Domaine, p.ClientID, p.ClientSecret, p.GrantType)
}

package model

import "fmt"

type Param struct {
	ID           int64  `db:"id"`
	Domaine      string `db:"domaine"`
	ClientID     string `db:"client_id"`
	ClientSecret string `db:"client_secret"`
	GrantType    string `db:"grant_type"`
}

func (p Param) String() string {
	return fmt.Sprintf("id=%v, title=%v, releaseDate=%v, duration=%v, trailerURL=%v",
		p.ID, p.Domaine, p.ClientID, p.ClientSecret, p.GrantType)
}

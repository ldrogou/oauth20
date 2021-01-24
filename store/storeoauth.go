package store

import (
	_ "github.com/mattn/go-sqlite3"
)

/**
type StoreOauth interface {
	GetOauth() (*model.Oauth, error)
	CreateOauth(m *model.Oauth) error
	DeleteOauth() error
}

func (store *DbStore) GetOauth() (*model.Oauth, error) {
	var oauth = &model.Oauth{}
	err := store.db.Get(oauth, "SELECT * FROM oauth")
	if err != nil {
		return oauth, err
	}
	return oauth, nil
}

func (store *DbStore) CreateOauth(o *model.Oauth) error {
	log.Println("je suis ici")
	log.Println(store.db.Ping())
	log.Println("je suis ici -------")
	log.Printf("la valeur de o %v \n", o)
	log.Println("je suis ici =======")
	res, err := store.db.Exec("INSERT INTO oauth (access_token, expire_in, refresh_token) VALUES (?, ?, ?)",
		o.AccessToken, o.ExpireIN, o.RefreshToken)
	log.Println("je suis ici @@@@@@@@")

	if err != nil {
		return err
	}

	o.ID, err = res.LastInsertId()
	return err

}

func (store *DbStore) DeleteOauth() error {
	_, err := store.db.Exec("DELETE TABLE oauth", nil)
	if err != nil {
		return err
	}

	return err

}
**/

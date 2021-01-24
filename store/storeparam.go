package store

import (
	_ "github.com/mattn/go-sqlite3"
)

/**
type StoreParam interface {
	GetParam() (*model.Param, error)
	CreateParam(m *model.Param) error
	DeleteParam() error
}

func (store *DbStore) GetParam() (*model.Param, error) {
	var param = &model.Param{}
	err := store.db.Get(param, "SELECT * FROM param")
	if err != nil {
		return param, err
	}
	return param, nil
}

func (store *DbStore) CreateParam(p *model.Param) error {
	res, err := store.db.Exec("INSERT INTO param (domaine, client_id, client_secret, grant_type) VALUES (?, ?, ?, ?)",
		p.Domaine, p.ClientID, p.ClientSecret, p.GrantType)
	if err != nil {
		return err
	}

	p.ID, err = res.LastInsertId()
	return err

}

func (store *DbStore) DeleteParam() error {
	_, err := store.db.Exec("DELETE TABLE param", nil)
	if err != nil {
		return err
	}

	return err

}
**/

package store

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/ldrogou/goauth20/model"
	_ "github.com/mattn/go-sqlite3"
)

type Store interface {
	Open() error
	Close() error

	GetOauth() (*model.Oauth, error)
	CreateOauth(m *model.Oauth) error
	DeleteOauth() error

	GetParam() (*model.Param, error)
	CreateParam(m *model.Param) error
	DeleteParam() error
}

type DbStore struct {
	db *sqlx.DB
}

var schemaAuth = `
CREATE TABLE IF NOT EXISTS oauth
(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	access_token TEXT,
	expire_in INTEGER,
	refresh_token TEXT
)
`

var schemaParam = `
CREATE TABLE IF NOT EXISTS param
(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	domaine TEXT,
	client_id TEXT,
	client_secret TEXT,
	grant_type TEXT
)
`

func (store *DbStore) Open() error {
	db, err := sqlx.Connect("sqlite3", "./oauth.db")
	if err != nil {
		return err
	}
	log.Println("Connected db")
	db.MustExec(schemaAuth)
	db.MustExec(schemaParam)

	store.db = db
	return nil
}

func (store *DbStore) Close() error {
	return store.db.Close()
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
	res, err := store.db.Exec("INSERT INTO oauth (access_token, expire_in, refresh_token) VALUES (?, ?, ?)",
		o.AccessToken, o.ExpireIN, o.RefreshToken)

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

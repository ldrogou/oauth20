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

	GetOauth(id int64) (*model.Oauth, error)
	CreateOauth(m *model.Oauth) error
	DeleteOauth(id int64) error

	GetParam(state string) (*model.Param, error)
	CreateParam(m *model.Param) error
	DeleteParam(state string) error
}

type DbStore struct {
	db *sqlx.DB
}

var schemaAuth = `
CREATE TABLE IF NOT EXISTS oauth
(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	access_token TEXT,
	token_type TEXT,
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

func (store *DbStore) GetOauth(id int64) (*model.Oauth, error) {
	var oauth = &model.Oauth{}
	err := store.db.Get(oauth, "SELECT * FROM oauth where id=$1", id)
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

func (store *DbStore) DeleteOauth(id int64) error {
	_, err := store.db.Exec("DELETE FROM oauth where id=?", id)
	if err != nil {
		return err
	}

	return err

}
func (store *DbStore) GetParam(state string) (*model.Param, error) {
	var param = &model.Param{}
	err := store.db.Get(param, "SELECT * FROM param where state=$1", state)
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

func (store *DbStore) DeleteParam(state string) error {
	_, err := store.db.Exec("DELETE FROM param where state=?", state)
	if err != nil {
		return err
	}

	return err

}

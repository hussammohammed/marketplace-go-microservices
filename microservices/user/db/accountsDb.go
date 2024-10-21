package db

import (
	"gopkg.in/mgo.v2"
)

type DB struct {
	session *mgo.Session
}

// Dial mongo server.
func Dial(uri string) (*DB, error) {
	d := &DB{}
	var err error
	d.session, err = mgo.Dial(uri)
	if err != nil {
		return nil, err
	}
	return d, nil
}

// Clone session from current connection of mongo.
func (d *DB) Clone() *mgo.Session {
	return d.session.Clone()
}

// Close current open session of mongodb.
func (d *DB) Close() {
	d.session.Close()
}

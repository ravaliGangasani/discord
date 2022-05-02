package models

import (
	"time"

	"gopkg.in/mgo.v2"
)

// nolint:gochecknoglobals
var (
	session *mgo.Session
	info    = &mgo.DialInfo{
		Addrs: []string{
			"127.0.0.1:27017",
		},
		Database: "Autonomy",
		Timeout:  15 * time.Second,
	}

	_ = func() (err error) {
		session, err = mgo.DialWithInfo(info)
		if err != nil {
			return err
		}

		return nil
	}()
)

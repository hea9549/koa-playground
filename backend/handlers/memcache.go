package handlers

import (
	"github.com/tidwall/buntdb"
	"time"
)

var defaultExpireHours = 4

func Save(key, value string, db *buntdb.DB) error {
	return SaveWithTTL(key, value, db, 1)
}

func SaveWithTTL(key, value string, db *buntdb.DB, hours int) error {
	if hours <= 0 {
		hours = defaultExpireHours
	}
	return db.Update(func(tx *buntdb.Tx) error {
		tx.Set(key, value, &buntdb.SetOptions{Expires: true, TTL: time.Second * time.Duration(hours)})
		return nil
	})
}

func Delete(keys []string, db *buntdb.DB) error {
	return db.Update(func(tx *buntdb.Tx) error {
		for _, key := range keys {
			if _, err := tx.Delete(key); err != nil {
				return err
			}
		}
		return nil
	})
}

func Update(key, newValue string, db *buntdb.DB) error {
	return UpdateWithTTL(key, newValue, db, 1)
}

func UpdateWithTTL(key, newValue string, db *buntdb.DB, hours int) error {
	if hours <= 0 {
		hours = defaultExpireHours
	}
	return db.Update(func(tx *buntdb.Tx) error {
		tx.Set(key, newValue, &buntdb.SetOptions{Expires: true, TTL: time.Second * time.Duration(hours)})
		return nil
	})
}
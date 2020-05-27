package db

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

var world = []byte("world")
var db *bolt.DB
var err error

// OpenDB ...
func OpenDB() {
	db, err = bolt.Open("bolt.db", 0644, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// CloseDB ...
func CloseDB() {
	db.Close()
}

// Create ...
func Create(key []byte, value []byte) (string, error) {
	OpenDB()
	defer CloseDB()

	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(world)
		if err != nil {
			return err
		}

		val := bucket.Get(key)
		if val != nil && len(val) > 0 {
			return fmt.Errorf("Key is already existed")
		}

		err = bucket.Put(key, value)
		if err != nil {
			return err
		}
		return nil
	})

	return string(value), err
}

// Read ...
func Read(key []byte) (string, error) {
	OpenDB()
	defer CloseDB()

	var val []byte
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(world)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found", world)
		}

		val = bucket.Get(key)
		return nil
	})

	return string(val), err
}

// Update ...
func Update(key []byte, value []byte) (string, error) {
	OpenDB()
	defer CloseDB()

	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(world)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found", world)
		}

		val := bucket.Get(key)
		if val == nil || len(val) == 0 {
			return fmt.Errorf("Key is not existed")
		}

		err = bucket.Put(key, value)
		if err != nil {
			return err
		}
		return nil
	})

	return string(value), err
}

// Delete ...
func Delete(key []byte) (string, error) {
	OpenDB()
	defer CloseDB()

	var val []byte
	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(world)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found", world)
		}

		val = bucket.Get(key)
		if val == nil || len(val) == 0 {
			return fmt.Errorf("Key is not existed")
		}

		err = bucket.Delete(key)
		if err != nil {
			return err
		}

		return nil
	})

	return string(val), err
}

// GetNextSequence ...
func GetNextSequence() uint64 {
	OpenDB()
	defer CloseDB()

	var nextID uint64
	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(world)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found", world)
		}

		nextID, _ = bucket.NextSequence()
		return nil
	})

	return nextID
}

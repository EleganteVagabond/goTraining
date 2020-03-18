package urlshort

import (
	"encoding/json"
	"log"

	"github.com/boltdb/bolt"
)

var (
	dbName       = "my.db"
	dbBucketName = "urlmap"
)

func prepareDB() (*bolt.DB, error) {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open(dbName, 0600, nil)
	if err == nil {
		createDBScaffolding(db)
	}
	// defer db.Close()
	return db, err
}

// vars must be exported for json marshaling to work
type dbRecord struct {
	Urlref string
	Remap  string
}

func createDBScaffolding(db *bolt.DB) error {
	err := db.Update(func(tx *bolt.Tx) error {

		bucket, err := tx.CreateBucketIfNotExists([]byte(dbBucketName))
		if err != nil {
			return err
		}
		dbrs := []dbRecord{
			{"/urlshortdb", "https://golangbot.com/learn-golang-series/"},
			{"/urlshortdb2", "https://www.facebook.com"},
			{"/urlshort-godoc", "https://www.google.com"}, // this will be overriden later by maphandler
		}
		for i, rec := range dbrs {
			dbk := []byte{byte(i)}
			log.Println("inserting key", dbk, "val", rec)
			err = insertDBRecord(bucket, dbk, rec)
			if err != nil {
				break
			}
		}
		return err
	})
	return err
}
func insertDBRecord(bucket *bolt.Bucket, dbk []byte, dbr dbRecord) error {
	// uncomment to refresh value for this key
	//bucket.Delete(dbk)
	// check if the key exists
	exval := bucket.Get(dbk)
	if exval == nil {
		// add the record
		jsonData, err := json.Marshal(dbr)
		if err != nil {
			return err
		}
		err = bucket.Put(dbk, jsonData)
		if err != nil {
			return err
		}
	}
	return nil
}

// load the values from the db using the json unmarshal function, storing in a map
func getDBValues(db *bolt.DB) (map[string]string, error) {
	ret := make(map[string]string)
	geterr := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(dbBucketName))
		var err error
		if b != nil {
			err = b.ForEach(func(k, v []byte) error {
				var dbr dbRecord
				err = json.Unmarshal(v, &dbr)
				if err != nil {
					return err
				}
				log.Println("gdbv adding key", k, "val", dbr)
				ret[dbr.Urlref] = dbr.Remap
				return nil
			})
		} else {
			log.Fatal("bucket not found")
		}
		return err
	})
	return ret, geterr
}

// LoadURLMapFromDB loads the data stored in the default bolt database
func LoadURLMapFromDB() (map[string]string, error) {
	var pathsToUrls map[string]string
	// try to open the db
	db, err := prepareDB()
	if err == nil {
		pathsToUrls, err = getDBValues(db)
	}
	defer db.Close()
	return pathsToUrls, err
}

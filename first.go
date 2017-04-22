package main

import (
	"log"
        _ "github.com/boltdb/bolt"
	"github.com/hashicorp/raft-boltdb"
	"github.com/hashicorp/raft"
)

type Hack struct {
     raftStore     *raftboltdb.BoltStore
}

func main() {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	/*
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	db.View(func(tx *bolt.Tx) error {
			//b := tx.Bucket([]byte("Suffrage"))

			c := tx.Cursor()
			for k,v := c.First(); k != nil; k,v = c.Next() {
			  log.Printf("key=%s, value=%s\n", k, v)
			}

			b := tx.Bucket([]byte("conf"))
			c = b.Cursor()
			for k,v := c.First(); k != nil; k,v = c.Next() {
			  log.Printf("key=%s, value=%s\n", k, v)
			}
			return nil
	})
	log.Print("worked")
	defer db.Close()
*/
	// Create the backend raft store for logs and stable storage.
	store, err := raftboltdb.NewBoltStore("raft.db")
	if err != nil {
	    log.Fatal(err)
	}
	if store != nil {
	   log.Print("opened the bolt store");
	}
	//lastIndex, err := store.LastIndex
	
	raft, err := raft.NewRaft(nil, nil, store, store, nil, nil, nil)
	if err != nil {
	    log.Fatal(err)
	}
	if raft != nil { log.Print("got a raft")}
}

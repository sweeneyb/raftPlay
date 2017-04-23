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
	store, err := raftboltdb.NewBoltStore("raft.db")
	if err != nil {
	    log.Fatal(err)
	}
	if store != nil {
	   log.Print("opened the bolt store");
	}
	lastIndex, err := store.LastIndex()
	log.Printf("%d",lastIndex)
	raftLog := &raft.Log{}	
	store.GetLog(lastIndex, raftLog)
	log.Printf("%s", raftLog)
	log.Printf("%s", raftLog.Data)
	
}

package main

import (
	"log"
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
	//log.Printf("%s", raftLog)
	//log.Printf("%s", raftLog.Data)
	var i, term uint64
   	for i = 0; i<=lastIndex; i++ {
	    store.GetLog(i, raftLog)
	    log.Printf("%d", i)
	    log.Printf("%s", raftLog)
	}
	term = raftLog.Term
	var removeLog = Log{Index: i, Term: term, Type: raft.LogRemovePeer, peer: "foo"}
	log.Printf("%s", removeLog)
	
}

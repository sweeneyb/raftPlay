package main

import (
	"bytes"
	"github.com/hashicorp/go-msgpack/codec"
	"github.com/hashicorp/raft"
	"github.com/hashicorp/raft-boltdb"
	"log"
)

type Hack struct {
	raftStore *raftboltdb.BoltStore
}

func main() {
	store, err := raftboltdb.NewBoltStore("raft.db")
	if err != nil {
		log.Fatal(err)
	}
	if store != nil {
		log.Print("opened the bolt store")
	}
	lastIndex, err := store.LastIndex()
	log.Printf("%d", lastIndex)
	raftLog := &raft.Log{}
	store.GetLog(lastIndex, raftLog)
	//log.Printf("%s", raftLog)
	//log.Printf("%s", raftLog.Data)
	var i, term uint64
	for i = 0; i <= lastIndex; i++ {
		store.GetLog(i, raftLog)
		log.Printf("%d", i)
		log.Printf("%s", raftLog)
	}
	term = raftLog.Term
	var removeLog = Log{Index: i, Term: term, Type: raft.LogRemovePeer}
	removeLog.Data = encodePeers([]string{"192.168.0.1:8300"})
	log.Printf("%s", removeLog)

}

func encodePeers(peers []string) []byte {
	// Encode each peer
	var encPeers [][]byte
	for _, p := range peers {
		// the net_trasport.go impl of EncodePeer is simple
		encPeers = append(encPeers, []byte(p))
	}

	// Encode the entire array
	buf, err := encodeMsgPack(encPeers)
	if err != nil {
		log.Fatalf("failed to encode peers: %v", err)
	}

	return buf.Bytes()
}

// Encode writes an encoded object to a new bytes buffer.
func encodeMsgPack(in interface{}) (*bytes.Buffer, error) {
	buf := bytes.NewBuffer(nil)
	hd := codec.MsgpackHandle{}
	enc := codec.NewEncoder(buf, &hd)
	err := enc.Encode(in)
	return buf, err
}

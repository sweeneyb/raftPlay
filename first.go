package main

import (
	"bytes"
	"fmt"
	"github.com/hashicorp/go-msgpack/codec"
	//"github.com/hashicorp/raft"
	//"github.com/hashicorp/raft-boltdb"
	"github.com/mitchellh/cli"
	"log"
	"os"
)

type Hack struct {
	//	raftStore *raftboltdb.BoltStore
}

func main() {

	ui := &cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}

	c := cli.NewCLI("raftTool", "0.0.1")
	c.Args = os.Args[1:]

	c.Commands = map[string]cli.CommandFactory{
		"logs": func() (cli.Command, error) {
			return &LogCommand{Ui: ui}, nil
		},
		"conf": func() (cli.Command, error) {
			return &ConfCommand{Ui: ui}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	os.Exit(exitStatus)

	/*
	   	dbPtr := flag.String("db-file", "raft.db", "filename of the raft db to alter")
	   	flag.Usage = func() {
	   		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	   		flag.PrintDefaults()
	   		fmt.Print("The first argument should be the node to remove. ie, 192.168.0.2:8300\n")
	   	}
	   	flag.Parse()
	   	if len(flag.Args()) < 1 {
	   		flag.Usage()
	   		os.Exit(1)
	   	}
	   	log.Print(flag.Args()[0])

	   	// This doesn't error if the file doesn't exist.  Which isn't ideal for this usage
	   	store, err := raftboltdb.NewBoltStore(*dbPtr)
	   	if err != nil {
	   		log.Fatal(err)
	   	}
	   	if store != nil {
	   		log.Print("opened the bolt store")
	   	}
	   	lastIndex, err := store.LastIndex()
	   	log.Printf("last transaction log index: %d", lastIndex)
	   	raftLog := &raft.Log{}
	   	store.GetLog(lastIndex, raftLog)
	   	var i uint64
	   	var term uint64
	   	i, err = store.FirstIndex()
	   	log.Printf("first index: %s", i)
	   	for ; i <= lastIndex; i++ {
	   		err = store.GetLog(i, raftLog)
	   		if err != nil {
	   			log.Print(err)
	   			break
	   		}
	   		fmt.Printf("index: %d\n", i)
	   		fmt.Printf("%s\n", raftLog)
	   		if(raftLog.Type == 2) {
	   		    fmt.Printf("%s\n", decodePeers(raftLog.Data))
	   		    if(len(decodePeers(raftLog.Data)) > 1) {
	                         fmt.Printf("would delete\n")
	   		      //err = store.DeleteRange(1, i)
	   		      if err != nil {
	   		        log.Print(err)
	   	              }
	   		    }
	   		}
	   	}
	   	term = raftLog.Term
	   	if i == 0 {
	   		log.Fatal("no transaction logs. Is this a real raft store?")
	   		os.Exit(2)
	   	}

	   	//var removeLog = &raft.Log{Index: i, Term: term, Type: raft.LogAddPeer}
	   	var removeLog = &raft.Log{Index: i, Term: term, Type: raft.LogRemovePeer}
	   	removeLog.Data = encodePeers([]string{flag.Args()[0]})
	   	log.Printf("to be appended: %s", removeLog)
	*/
	/*
		err = store.StoreLog(removeLog)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Print("message appended")
		}
	*/
	/*
		err = store.Close()
		if err != nil {
		  log.Fatal(err)
		}
	*/

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

//from raft's util.go
// decodePeers is used to deserialize a list of peers.
func decodePeers(buf []byte) []string {
	// Decode the buffer first
	var encPeers [][]byte
	if err := decodeMsgPack(buf, &encPeers); err != nil {
		panic(fmt.Errorf("failed to decode peers: %v", err))
	}

	// Deserialize each peer
	var peers []string
	for _, enc := range encPeers {
		peers = append(peers, string(enc))
	}
	return peers
}

// Decode reverses the encode operation on a byte slice input.
func decodeMsgPack(buf []byte, out interface{}) error {
	r := bytes.NewBuffer(buf)
	hd := codec.MsgpackHandle{}
	dec := codec.NewDecoder(r, &hd)
	return dec.Decode(out)
}

package main

import (
	"github.com/filecoin-project/specs-storage/storage"
	gen "github.com/whyrusleeping/cbor-gen"
)

func main() {
	// Common types
	if err := gen.WriteTupleEncodersToFile("./pkg/types_gen.go", "storage",
		storage.PoStCandidateWithTicket{},
	); err != nil {
		panic(err)
	}
}

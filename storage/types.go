package storage

import (
	"github.com/filecoin-project/specs-actors/actors/abi"
)

type PoStCandidateWithTicket struct {
	Candidate abi.PoStCandidate
	Ticket    [32]byte
}

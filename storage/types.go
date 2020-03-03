package storage

import (
	"bytes"

	"github.com/filecoin-project/specs-actors/actors/abi"
	"github.com/ipfs/go-cid"
)

// SectorInfo holds all sector-related metadata
type SectorInfo struct {
	ID abi.SectorID

	Pieces []abi.PieceInfo

	Ticket SealTicket
	Seed   SealSeed

	PreCommit1Out []byte

	Sealed   cid.Cid
	Unsealed cid.Cid

	CommitInput []byte
	Proof       []byte
}

type SealTicket struct {
	Value abi.SealRandomness
	Epoch abi.ChainEpoch
}

type SealSeed struct {
	Value abi.InteractiveSealRandomness
	Epoch abi.ChainEpoch
}

func (st *SealTicket) Equals(ost *SealTicket) bool {
	return bytes.Equal(st.Value, ost.Value) && st.Epoch == ost.Epoch
}

func (st *SealSeed) Equals(ost *SealSeed) bool {
	return bytes.Equal(st.Value, ost.Value) && st.Epoch == ost.Epoch
}

func (si SectorInfo) PieceSizes() []abi.UnpaddedPieceSize {
	out := make([]abi.UnpaddedPieceSize, len(si.Pieces))
	for i := range out {
		out[i] = si.Pieces[i].Size.Unpadded()
	}

	return nil
}

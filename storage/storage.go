package storage

import (
	"context"
	"io"

	ffi "github.com/filecoin-project/filecoin-ffi"
	"github.com/filecoin-project/specs-actors/actors/abi"
	"github.com/ipfs/go-cid"
)

type Data = io.Reader

type Storage interface {
	// Creates a new empty sector
	NewSector(ctx context.Context) (SectorInfo, error)
	// Add a piece to an existing *unsealed* sector
	AddPiece(ctx context.Context, sector abi.SectorID, pieces []abi.PieceInfo, pieceSize abi.UnpaddedPieceSize, r Data) (cid.Cid, SectorInfo, error)
}

type Verifier interface {
	GenerateEPostCandidates(sectorInfo []abi.SectorInfo, challengeSeed abi.PoStRandomness, faults []abi.SectorNumber) ([]ffi.PoStCandidateWithTicket, error)
	GenerateFallbackPoSt(sectorInfo []abi.SectorInfo, challengeSeed abi.PoStRandomness, faults []abi.SectorNumber) ([]ffi.PoStCandidateWithTicket, []abi.PoStProof, error)
	ComputeElectionPoSt(sectorInfo []abi.SectorInfo, challengeSeed abi.PoStRandomness, winners []abi.PoStCandidate) ([]abi.PoStProof, error)
}

// Review: can we reduce eachg of this, which arg goes where
type Sealer interface {
	SealPrecommit1(ctx context.Context, si SectorInfo) (SectorInfo, error)
	SealPrecommit2(ctx context.Context, si SectorInfo) (SectorInfo, error)
	SealCommit1(ctx context.Context, si SectorInfo) (SectorInfo, error)
	SealCommit2(ctx context.Context, si SectorInfo) (SectorInfo, error)
}
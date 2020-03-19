package storage

import (
	"context"
	"io"

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/specs-actors/actors/abi"
)

type Data = io.Reader

type Storage interface {
	// Creates a new empty sector (only allocate on disk. Layers above
	//  storage are responsible for assigning sector IDs)
	NewSector(ctx context.Context, sector abi.SectorID) error
	// Add a piece to an existing *unsealed* sector
	AddPiece(ctx context.Context, sector abi.SectorID, pieceSizes []abi.UnpaddedPieceSize, newPieceSize abi.UnpaddedPieceSize, pieceData Data) (abi.PieceInfo, error)
}

type FallbackPostOut struct {
	PoStInputs []PoStCandidateWithTicket
	Proof      []abi.PoStProof
}

type Prover interface {
	GenerateEPostCandidates(ctx context.Context, miner abi.ActorID, sectorInfo []abi.SectorInfo, challengeSeed abi.PoStRandomness, faults []abi.SectorNumber) ([]PoStCandidateWithTicket, error)
	GenerateFallbackPoSt(ctx context.Context, miner abi.ActorID, sectorInfo []abi.SectorInfo, challengeSeed abi.PoStRandomness, faults []abi.SectorNumber) (FallbackPostOut, error)
	ComputeElectionPoSt(ctx context.Context, miner abi.ActorID, sectorInfo []abi.SectorInfo, challengeSeed abi.PoStRandomness, winners []abi.PoStCandidate) ([]abi.PoStProof, error)
}

type PreCommit1Out []byte

type Commit1Out []byte

type Proof []byte

type SectorCids struct {
	Unsealed cid.Cid
	Sealed   cid.Cid
}

type Sealer interface {
	SealPreCommit1(ctx context.Context, sector abi.SectorID, ticket abi.SealRandomness, pieces []abi.PieceInfo) (PreCommit1Out, error)
	SealPreCommit2(ctx context.Context, sector abi.SectorID, pc1o PreCommit1Out) (SectorCids, error)

	// MakeProvable makes sure that sector is ready to receive challenges from
	//  proving calls
	MakeProvable(ctx context.Context, sector abi.SectorID) error

	SealCommit1(ctx context.Context, sector abi.SectorID, ticket abi.SealRandomness, seed abi.InteractiveSealRandomness, pieces []abi.PieceInfo, cids SectorCids) (Commit1Out, error)
	SealCommit2(ctx context.Context, sector abi.SectorID, c1o Commit1Out) (Proof, error)
	FinalizeSector(ctx context.Context, sector abi.SectorID) error
}

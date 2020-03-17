package storage

import (
	"context"
	"io"

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/specs-actors/actors/abi"
)

type Data = io.Reader

type Storage interface {
	// Creates a new empty sector
	NewSector(ctx context.Context, miner abi.ActorID) (abi.SectorID, error)
	// Add a piece to an existing *unsealed* sector
	AddPiece(ctx context.Context, sector abi.SectorID, pieceSizes []abi.UnpaddedPieceSize, newPieceSize abi.UnpaddedPieceSize, pieceData Data) (abi.PieceInfo, error)
}

type Prover interface {
	GenerateEPostCandidates(miner abi.ActorID, sectorInfo []abi.SectorInfo, challengeSeed abi.PoStRandomness, faults []abi.SectorNumber) ([]PoStCandidateWithTicket, error)
	GenerateFallbackPoSt(miner abi.ActorID, sectorInfo []abi.SectorInfo, challengeSeed abi.PoStRandomness, faults []abi.SectorNumber) ([]PoStCandidateWithTicket, []abi.PoStProof, error)
	ComputeElectionPoSt(miner abi.ActorID, sectorInfo []abi.SectorInfo, challengeSeed abi.PoStRandomness, winners []abi.PoStCandidate) ([]abi.PoStProof, error)
}

type PreCommit1Out []byte

type Commit1Out []byte

type Proof []byte

type Sealer interface {
	SealPreCommit1(ctx context.Context, sector abi.SectorID, ticket abi.SealRandomness, pieces []abi.PieceInfo) (PreCommit1Out, error)
	SealPreCommit2(ctx context.Context, sector abi.SectorID, pc1o PreCommit1Out) (sealedCID cid.Cid, unsealedCID cid.Cid, err error)
	SealCommit1(ctx context.Context, sector abi.SectorID, ticket abi.SealRandomness, seed abi.InteractiveSealRandomness, pieces []abi.PieceInfo, sealedCID cid.Cid, unsealedCID cid.Cid) (Commit1Out, error)
	SealCommit2(ctx context.Context, sector abi.SectorID, c1o Commit1Out) (Proof, error)
	FinalizeSector(ctx context.Context, sector abi.SectorID) error
}

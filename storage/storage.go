package storage

import (
	"context"
	"io"

	"github.com/filecoin-project/specs-actors/actors/abi"
	"github.com/ipfs/go-cid"
)

type Data = io.Reader

type Storage interface {
	// Creates a new empty sector
	NewSector(ctx context.Context) (abi.SectorNumber, error)
	// Add a piece to an existing *unsealed* sector
	AddPiece(ctx context.Context, sector abi.SectorNumber, pieceSizes []abi.UnpaddedPieceSize, newPieceSize abi.UnpaddedPieceSize, pieceData Data) (abi.PieceInfo, error)
}

type Prover interface {
	GenerateEPostCandidates(sectorInfo []abi.SectorInfo, challengeSeed abi.PoStRandomness, faults []abi.SectorNumber) ([]PoStCandidateWithTicket, error)
	GenerateFallbackPoSt(sectorInfo []abi.SectorInfo, challengeSeed abi.PoStRandomness, faults []abi.SectorNumber) ([]PoStCandidateWithTicket, []abi.PoStProof, error)
	ComputeElectionPoSt(sectorInfo []abi.SectorInfo, challengeSeed abi.PoStRandomness, winners []abi.PoStCandidate) ([]abi.PoStProof, error)
}

type PreCommit1Out []byte

type Commit1Out []byte

type Proof []byte

type Sealer interface {
	SealPreCommit1(ctx context.Context, sectorNum abi.SectorNumber, ticket abi.SealRandomness, pieces []abi.PieceInfo) (PreCommit1Out, error)
	SealPreCommit2(context.Context, abi.SectorNumber, PreCommit1Out) (sealedCID cid.Cid, unsealedCID cid.Cid, err error)
	SealCommit1(ctx context.Context, sectorNum abi.SectorNumber, ticket abi.SealRandomness, seed abi.InteractiveSealRandomness, pieces []abi.PieceInfo, sealedCID cid.Cid, unsealedCID cid.Cid) (Commit1Out, error)
	SealCommit2(context.Context, abi.SectorNumber, Commit1Out) (Proof, error)
	FinalizeSector(context.Context, abi.SectorNumber) error
}

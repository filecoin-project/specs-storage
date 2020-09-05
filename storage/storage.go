package storage

import (
	"context"
	"io"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"

	proof "github.com/filecoin-project/specs-actors/actors/runtime/proof"
)

type Data = io.Reader

type Storage interface {
	// Creates a new empty sector (only allocate on disk. Layers above
	//  storage are responsible for assigning sector IDs)
	NewSector(ctx context.Context, sector abi.SectorID) error
	// Add a piece to an existing *unsealed* sector
	AddPiece(ctx context.Context, sector abi.SectorID, pieceSizes []abi.UnpaddedPieceSize, newPieceSize abi.UnpaddedPieceSize, pieceData Data) (abi.PieceInfo, error)
}

type Prover interface {
	GenerateWinningPoSt(ctx context.Context, minerID abi.ActorID, sectorInfo []proof.SectorInfo, randomness abi.PoStRandomness) ([]proof.PoStProof, error)
	GenerateWindowPoSt(ctx context.Context, minerID abi.ActorID, sectorInfo []proof.SectorInfo, randomness abi.PoStRandomness) (proof []proof.PoStProof, skipped []abi.SectorID, err error)
}

type PreCommit1Out []byte

type Commit1Out []byte

type Proof []byte

type SectorCids struct {
	Unsealed cid.Cid
	Sealed   cid.Cid
}

type Range struct {
	Offset abi.UnpaddedPieceSize
	Size   abi.UnpaddedPieceSize
}

type Sealer interface {
	SealPreCommit1(ctx context.Context, sector abi.SectorID, ticket abi.SealRandomness, pieces []abi.PieceInfo) (PreCommit1Out, error)
	SealPreCommit2(ctx context.Context, sector abi.SectorID, pc1o PreCommit1Out) (SectorCids, error)

	SealCommit1(ctx context.Context, sector abi.SectorID, ticket abi.SealRandomness, seed abi.InteractiveSealRandomness, pieces []abi.PieceInfo, cids SectorCids) (Commit1Out, error)
	SealCommit2(ctx context.Context, sector abi.SectorID, c1o Commit1Out) (Proof, error)

	FinalizeSector(ctx context.Context, sector abi.SectorID, keepUnsealed []Range) error

	// ReleaseUnsealed marks parts of the unsealed sector file as safe to drop
	//  (called by the fsm on restart, allows storage to keep no persistent
	//   state about unsealed fast-retrieval copies)
	ReleaseUnsealed(ctx context.Context, sector abi.SectorID, safeToFree []Range) error

	// Removes all data associated with the specified sector
	Remove(ctx context.Context, sector abi.SectorID) error
}

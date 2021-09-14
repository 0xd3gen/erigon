package stagedsync

import (
	"context"

	"github.com/ledgerwatch/erigon-lib/kv"
	"github.com/ledgerwatch/erigon/eth/ethconfig"
	"github.com/ledgerwatch/erigon/turbo/snapshotsync"
)

type SnapshotBodiesCfg struct {
	enabled          bool
	db               kv.RwDB
	epochSize        uint64
	snapshotDir      string
	tmpDir           string
	client           *snapshotsync.Client
	snapshotMigrator *snapshotsync.SnapshotMigrator
}

func StageSnapshotBodiesCfg(db kv.RwDB, snapshot ethconfig.Snapshot, client *snapshotsync.Client, snapshotMigrator *snapshotsync.SnapshotMigrator, tmpDir string) SnapshotBodiesCfg {
	return SnapshotBodiesCfg{
		enabled:          snapshot.Enabled && snapshot.Mode.Bodies,
		db:               db,
		snapshotDir:      snapshot.Dir,
		client:           client,
		snapshotMigrator: snapshotMigrator,
		tmpDir:           tmpDir,
	}
}

func SpawnBodiesSnapshotGenerationStage(s *StageState, tx kv.RwTx, cfg SnapshotBodiesCfg, initialSync bool, ctx context.Context) error {
	if !initialSync || cfg.epochSize == 0 {
		return nil
	}
	return nil
}

func UnwindBodiesSnapshotGenerationStage(s *UnwindState, tx kv.RwTx, cfg SnapshotBodiesCfg, ctx context.Context) (err error) {
	useExternalTx := tx != nil
	if !useExternalTx {
		tx, err = cfg.db.BeginRw(ctx)
		if err != nil {
			return err
		}
		defer tx.Rollback()
	}

	if err := s.Done(tx); err != nil {
		return err
	}
	if !useExternalTx {
		if err := tx.Commit(); err != nil {
			return err
		}
	}
	return nil
}

func PruneBodiesSnapshotGenerationStage(s *PruneState, tx kv.RwTx, cfg SnapshotBodiesCfg, ctx context.Context) (err error) {
	useExternalTx := tx != nil
	if !useExternalTx {
		tx, err = cfg.db.BeginRw(ctx)
		if err != nil {
			return err
		}
		defer tx.Rollback()
	}

	if !useExternalTx {
		if err := tx.Commit(); err != nil {
			return err
		}
	}
	return nil
}

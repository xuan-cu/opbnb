package node

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

type abnormalAPI struct {
	ConfigPersistence ConfigPersistence
}

func NewAbnormalAPI(ConfigPersistence ConfigPersistence) *abnormalAPI {
	return &abnormalAPI{
		ConfigPersistence: ConfigPersistence,
	}
}

func (n *abnormalAPI) ResetDerivationPipeline(ctx context.Context) error {
	return nil
}

func (n *abnormalAPI) StartSequencer(ctx context.Context, blockHash common.Hash) error {
	return errors.New("sequencer not running")
}

func (n *abnormalAPI) StopSequencer(ctx context.Context) (common.Hash, error) {
	if err := n.ConfigPersistence.SequencerStopped(); err != nil {
		return common.Hash{}, errors.New("SequencerStopped error")
	}
	return common.Hash{}, nil
}

func (n *abnormalAPI) SequencerActive(ctx context.Context) (bool, error) {
	if state, err := n.ConfigPersistence.SequencerState(); err != nil {
		return false, errors.New("SequencerStopped error")
	} else if state != StateStarted {
		return false, nil
	}
	return true, nil
}

func (n *abnormalAPI) AbnormalAPI(ctx context.Context) (bool, error) {
	return true, nil
}

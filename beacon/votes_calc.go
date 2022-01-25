package beacon

import (
	"context"
	"fmt"
	"math/big"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/log"
)

func (pd *ProtocolDriver) calcVotes(ctx context.Context, epoch types.EpochID, round types.RoundID) (allVotes, []string, error) {
	logger := pd.logger.WithContext(ctx).WithFields(epoch, round)
	pd.mu.Lock()
	defer pd.mu.Unlock()

	logger.With().Debug("calculating votes", log.String("vote_margins", fmt.Sprint(pd.votesMargin)))

	ownCurrentRoundVotes, undecided, err := pd.calcOwnCurrentRoundVotes()
	if err != nil {
		return allVotes{}, nil, fmt.Errorf("calc own current round votes: %w", err)
	}

	logger.With().Debug("calculated votes for one round",
		log.String("for_votes", fmt.Sprint(ownCurrentRoundVotes.valid)),
		log.String("against_votes", fmt.Sprint(ownCurrentRoundVotes.invalid)))

	return ownCurrentRoundVotes, undecided, nil
}

func (pd *ProtocolDriver) calcOwnCurrentRoundVotes() (allVotes, []string, error) {
	ownCurrentRoundVotes := allVotes{
		valid:   make(proposalSet),
		invalid: make(proposalSet),
	}

	positiveVotingThreshold := pd.votingThreshold(pd.epochWeight)
	negativeThreshold := new(big.Int).Neg(positiveVotingThreshold)

	var undecided []string
	for vote, weightCount := range pd.votesMargin {
		switch {
		case weightCount.Cmp(positiveVotingThreshold) >= 0:
			ownCurrentRoundVotes.valid[vote] = struct{}{}
		case weightCount.Cmp(negativeThreshold) <= 0:
			ownCurrentRoundVotes.invalid[vote] = struct{}{}
		default:
			undecided = append(undecided, vote)
		}
	}
	return ownCurrentRoundVotes, undecided, nil
}

func tallyUndecided(votes *allVotes, undecided []string, coinFlip bool) {
	for _, vote := range undecided {
		if coinFlip {
			votes.valid[vote] = struct{}{}
		} else {
			votes.invalid[vote] = struct{}{}
		}
	}
}
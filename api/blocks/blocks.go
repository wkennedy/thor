// Copyright (c) 2018 The VeChainThor developers

// Distributed under the GNU Lesser General Public License v3.0 software license, see the accompanying
// file LICENSE or <https://www.gnu.org/licenses/lgpl-3.0.html>

package blocks

import (
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/common/math"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/vechain/thor/api/utils"
	"github.com/vechain/thor/chain"
	"github.com/vechain/thor/thor"
)

type Blocks struct {
	repo *chain.Repository
}

func New(repo *chain.Repository) *Blocks {
	return &Blocks{
		repo,
	}
}

func (b *Blocks) handleGetBlock(w http.ResponseWriter, req *http.Request) error {
	revision, err := b.parseRevision(mux.Vars(req)["revision"])
	if err != nil {
		return utils.BadRequest(errors.WithMessage(err, "revision"))
	}
	expanded := req.URL.Query().Get("expanded")
	if expanded != "" && expanded != "false" && expanded != "true" {
		return utils.BadRequest(errors.WithMessage(errors.New("should be boolean"), "expanded"))
	}

	summary, err := b.getBlockSummary(revision)
	if err != nil {
		if b.repo.IsNotFound(err) {
			return utils.WriteJSON(w, nil)
		}
		return err
	}
	isTrunk, err := b.isTrunk(summary.Header.ID(), summary.Header.Number())
	if err != nil {
		return err
	}

	jSummary := buildJSONBlockSummary(summary, isTrunk)
	if expanded == "true" {
		txs, err := b.repo.GetBlockTransactions(summary.Header.ID())
		if err != nil {
			return err
		}
		receipts, err := b.repo.GetBlockReceipts(summary.Header.ID())
		if err != nil {
			return err
		}

		return utils.WriteJSON(w, &JSONExpandedBlock{
			jSummary,
			buildJSONEmbeddedTxs(txs, receipts),
		})
	}

	return utils.WriteJSON(w, &JSONCollapsedBlock{
		jSummary,
		summary.Txs,
	})
}

func (b *Blocks) parseRevision(revision string) (interface{}, error) {
	if revision == "" || revision == "best" {
		return nil, nil
	}
	if len(revision) == 66 || len(revision) == 64 {
		blockID, err := thor.ParseBytes32(revision)
		if err != nil {
			return nil, err
		}
		return blockID, nil
	}
	n, err := strconv.ParseUint(revision, 0, 0)
	if err != nil {
		return nil, err
	}
	if n > math.MaxUint32 {
		return nil, errors.New("block number out of max uint32")
	}
	return uint32(n), err
}

func (b *Blocks) getBlockSummary(revision interface{}) (s *chain.BlockSummary, err error) {
	var id thor.Bytes32
	switch revision.(type) {
	case thor.Bytes32:
		id = revision.(thor.Bytes32)
	case uint32:
		id, err = b.repo.NewBestChain().GetBlockID(revision.(uint32))
		if err != nil {
			return
		}
	default:
		id = b.repo.BestBlock().Header().ID()
	}
	return b.repo.GetBlockSummary(id)
}

func (b *Blocks) isTrunk(blkID thor.Bytes32, blkNum uint32) (bool, error) {
	idByNum, err := b.repo.NewBestChain().GetBlockID(blkNum)
	if err != nil {
		return false, err
	}
	return blkID == idByNum, nil
}

func (b *Blocks) Mount(root *mux.Router, pathPrefix string) {
	sub := root.PathPrefix(pathPrefix).Subrouter()
	sub.Path("/{revision}").Methods("GET").HandlerFunc(utils.WrapHandlerFunc(b.handleGetBlock))

}

package mempool

import (
	"context"
	"cosmossdk.io/log"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
)

type txName struct {
	sender sdk.AccAddress
}

func TestThresholdMempool(t *testing.T) {
	accounts := simtypes.RandomAccounts(rand.New(rand.NewSource(0)), 5)
	alice := accounts[0].Address
	bob := accounts[1].Address
	barb := accounts[2].Address
	cindy := accounts[3].Address

	tests := []struct {
		txs   []txName
		order []int
	}{
		{
			txs: []txName{
				{sender: alice},
				{sender: bob},
				{sender: cindy},
				{sender: alice},
				{sender: barb},
			},
			// Fifo order should be maintained
			order: []int{0, 1, 2, 3, 4},
		},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			logger := log.NewTestLogger(t)
			pool := NewThresholdMempool(logger)
			for j, tt := range tc.txs {
				tx := testTx{
					id:      j,
					address: tt.sender,
					nonce:   uint64(j),
				}
				err := pool.Insert(context.Background(), tx)
				require.NoError(t, err)
				require.Equal(t, 1+j, pool.CountTx())
			}

			var orderedTxs []sdk.Tx
			itr := pool.SelectPending(context.Background(), nil)
			for itr != nil {
				current := itr.Tx()
				orderedTxs = append(orderedTxs, current)
				itr = itr.Next()
			}

			var txOrder []int
			for _, tx := range orderedTxs {
				txOrder = append(txOrder, tx.(testTx).id)
			}
			require.Equal(t, tc.order, txOrder)

			// Pool should be empty
			require.Equal(t, len(pool.pool.txs), 0)

			last := len(pool.pendingPool.txs) - 1
			lastTx := pool.pendingPool.txs[last]
			err := pool.Update(context.Background(), lastTx.tx)
			require.NoError(t, err)

			newPendingLen := len(pool.pendingPool.txs)
			require.Equal(t, newPendingLen, 4)
			newPoolLen := len(pool.pool.txs)
			require.Equal(t, newPoolLen, 1)
			// Update select to only return txs with priority 1
			// Create function that updates priority
			var readyTxs []sdk.Tx
			itr2 := pool.Select(context.Background(), nil)
			for itr2 != nil {
				current := itr2.Tx()
				readyTxs = append(readyTxs, current)
				itr2 = itr2.Next()
			}

			require.Equal(t, len(readyTxs), 1)

			err = pool.Remove(readyTxs[0])
			require.NoError(t, err)
			require.Equal(t, 0, len(pool.pool.txs))
		})
	}
}

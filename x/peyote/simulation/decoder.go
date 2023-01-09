package simulation

import (
	"bytes"
	"fmt"

	tmkv "github.com/tendermint/tendermint/libs/kv"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/warmage-sports/peyote/x/peyote/internal/types"
)

// DecodeStore unmarshals the KVPair's Value to the corresponding type
func DecodeStore(cdc *codec.Codec, kvA, kvB tmkv.Pair) string {
	switch {
	case bytes.Equal(kvA.Key[:1], types.BondsKeyPrefix):
		var bondA, bondB types.Bond
		cdc.MustUnmarshalBinaryBare(kvA.Value, &bondA)
		cdc.MustUnmarshalBinaryBare(kvB.Value, &bondB)
		return fmt.Sprintf("%v\n%v", bondA, bondB)

	case bytes.Equal(kvA.Key[:1], types.BatchesKeyPrefix):
		var batchA, batchB types.Batch
		cdc.MustUnmarshalBinaryBare(kvA.Value, &batchA)
		cdc.MustUnmarshalBinaryBare(kvB.Value, &batchB)
		return fmt.Sprintf("%v\n%v", batchA, batchB)

	case bytes.Equal(kvA.Key[:1], types.LastBatchesKeyPrefix):
		var batchA, batchB types.Batch
		cdc.MustUnmarshalBinaryBare(kvA.Value, &batchA)
		cdc.MustUnmarshalBinaryBare(kvB.Value, &batchB)
		return fmt.Sprintf("%v\n%v", batchA, batchB)

	default:
		panic(fmt.Sprintf("invalid %s key prefix %X", types.ModuleName, kvA.Key[:1]))
	}
}

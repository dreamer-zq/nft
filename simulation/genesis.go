package simulation

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/irismod/nft/types"
)

const (
	kitties = "kitties"
	doggos  = "doggos"
)

// RandomizedGenState generates a random GenesisState for nft
func RandomizedGenState(simState *module.SimulationState) {
	collections := types.NewCollections(
		types.NewCollection(types.Denom{
			Name:    doggos,
			Schema:  "",
			Creator: nil,
		}, types.NFTs{}),
		types.NewCollection(types.Denom{
			Name:    kitties,
			Schema:  "",
			Creator: nil,
		}, types.NFTs{}))
	for _, acc := range simState.Accounts {
		// 10% of accounts own an NFT
		if simState.Rand.Intn(100) < 10 {
			baseNFT := types.NewBaseNFT(
				simulation.RandStringOfLength(simState.Rand, 5), // id
				acc.Address,
				simulation.RandStringOfLength(simState.Rand, 45), // tokenURI
				simulation.RandStringOfLength(simState.Rand, 10),
			)

			// 50% doggos and 50% kitties
			if simState.Rand.Intn(100) < 50 {
				collections[0].Denom.Creator = baseNFT.Owner
				collections[0] = collections[0].AddNFT(&baseNFT)
			} else {
				collections[1].Denom.Creator = baseNFT.Owner
				collections[1] = collections[1].AddNFT(&baseNFT)
			}
		}
	}

	nftGenesis := types.NewGenesisState(collections)

	fmt.Printf("Selected randomly generated NFT genesis state:\n%s\n", codec.MustMarshalJSONIndent(simState.Cdc, nftGenesis))
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(nftGenesis)
}

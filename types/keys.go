package types

import (
	"bytes"
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName is the name of the module
	ModuleName = "nft"

	// StoreKey is the default store key for NFT
	StoreKey = ModuleName

	// QuerierRoute is the querier route for the NFT store.
	QuerierRoute = ModuleName

	// RouterKey is the message route for the NFT module
	RouterKey = ModuleName
)

var (
	PrefixNFT        = []byte{0x01}
	PrefixOwners     = []byte{0x02} // key for balance of NFTs held by an address
	PrefixCollection = []byte{0x03} // key for balance of NFTs held by an address

	delimiter = []byte("/")
)

func SplitKeyOwner(key []byte) (address sdk.AccAddress, denom, id string, err error) {
	key = key[len(PrefixOwners)+len(delimiter):]
	keys := bytes.Split(key, delimiter)

	switch len(keys) {
	case 3:
		address, _ = sdk.AccAddressFromBech32(string(keys[0]))
		denom = string(keys[1])
		id = string(keys[2])
		return
	case 2:
		address, _ = sdk.AccAddressFromBech32(string(keys[0]))
		denom = string(keys[1])
		return
	case 1:
		address, _ = sdk.AccAddressFromBech32(string(keys[0]))
		return
	default:
		return address, denom, id, errors.New("wrong KeyOwner")
	}
}

// KeyOwner gets the key of a collection owned by an account address
func KeyOwner(address sdk.AccAddress, denom, id string) []byte {
	key := append(PrefixOwners, delimiter...)
	if address != nil {
		key = append(key, []byte(address.String())...)
		key = append(key, delimiter...)
	}

	if address != nil && len(denom) > 0 {
		key = append(key, []byte(denom)...)
		key = append(key, delimiter...)
	}

	if address != nil && len(denom) > 0 && len(id) > 0 {

		key = append(key, []byte(id)...)
	}
	return key
}

// KeyNFT gets the NFT by an ID
func KeyNFT(denom, id string) []byte {
	key := append(PrefixNFT, delimiter...)
	if len(denom) > 0 {
		key = append(key, []byte(denom)...)
		key = append(key, delimiter...)
	}

	if len(denom) > 0 && len(id) > 0 {
		key = append(key, []byte(id)...)
	}
	return key
}

// KeyNFT gets the NFT by an ID
func KeyCollection(denom string) []byte {
	key := append(PrefixCollection, delimiter...)
	return append(key, []byte(denom)...)
}

func SplitKeyCollection(key []byte) (denom string) {
	key = key[len(PrefixCollection)+len(delimiter):]
	return string(key)
}

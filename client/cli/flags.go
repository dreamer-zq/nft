package cli

import (
	flag "github.com/spf13/pflag"
)

const (
	FlagTokenURI  = "token-uri"
	FlagMetadata  = "metadata"
	FlagRecipient = "recipient"
	FlagOwner     = "owner"
	FlagDenom     = "denom"
	FlagSchema    = "schema"
)

var (
	FsIssueDenom  = flag.NewFlagSet("", flag.ContinueOnError)
	FsMintNFT     = flag.NewFlagSet("", flag.ContinueOnError)
	FsEditNFT     = flag.NewFlagSet("", flag.ContinueOnError)
	FsTransferNFT = flag.NewFlagSet("", flag.ContinueOnError)
	FsQuerySupply = flag.NewFlagSet("", flag.ContinueOnError)
	FsQueryOwner  = flag.NewFlagSet("", flag.ContinueOnError)
)

func init() {
	FsIssueDenom.String(FlagSchema, "[do-not-modify]", "denom data structure definition")

	FsMintNFT.String(FlagTokenURI, "[do-not-modify]", "URI for supplemental off-chain metadata (should return a JSON object)")
	FsMintNFT.String(FlagRecipient, "", "Receiver of the nft, if not filled, the default is the sender of the transaction")
	FsMintNFT.String(FlagMetadata, "", "the metadata of nft")

	FsEditNFT.String(FlagTokenURI, "[do-not-modify]", "URI for supplemental off-chain metadata (should return a JSON object)")
	FsEditNFT.String(FlagMetadata, "[do-not-modify]", "the metadata of nft")

	FsTransferNFT.String(FlagTokenURI, "[do-not-modify]", "URI for supplemental off-chain metadata (should return a JSON object)")
	FsTransferNFT.String(FlagMetadata, "[do-not-modify]", "the metadata of nft")

	FsQuerySupply.String(FlagOwner, "", "the owner of a nft")

	FsQueryOwner.String(FlagDenom, "", "the name of a collection")
}

package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/warmage-sports/peyote/x/peyote/internal/keeper"
	"github.com/warmage-sports/peyote/x/peyote/internal/types"
	"strings"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	peyoteQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Bonds querying subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	peyoteQueryCmd.AddCommand(flags.GetCommands(
		GetCmdBonds(storeKey, cdc),
		GetCmdBond(storeKey, cdc),
		GetCmdBatch(storeKey, cdc),
		GetCmdLastBatch(storeKey, cdc),
		GetCmdCurrentPrice(storeKey, cdc),
		GetCmdCurrentReserve(storeKey, cdc),
		GetCmdCustomPrice(storeKey, cdc),
		GetCmdBuyPrice(storeKey, cdc),
		GetCmdSellReturn(storeKey, cdc),
		GetCmdSwapReturn(storeKey, cdc),
		GetCmdQueryParams(cdc),
	)...)

	return peyoteQueryCmd
}

func GetCmdBonds(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "peyote-list",
		Short: "List of all peyote",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, _, err := cliCtx.QueryWithData(
				fmt.Sprintf("custom/%s/peyote",
					queryRoute), nil)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			var out types.QueryBonds
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdBond(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "bond [bond-token]",
		Short: "Query info of a bond",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			bondToken := args[0]

			res, _, err := cliCtx.QueryWithData(
				fmt.Sprintf("custom/%s/bond/%s",
					queryRoute, bondToken), nil)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			var out types.Bond
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdBatch(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "batch [bond-token]",
		Short: "Query info of a bond's current batch",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			bondToken := args[0]

			res, _, err := cliCtx.QueryWithData(
				fmt.Sprintf("custom/%s/batch/%s",
					queryRoute, bondToken), nil)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			var out types.Batch
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdLastBatch(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "last-batch [bond-token]",
		Short: "Query info of a bond's last batch",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			bondToken := args[0]

			res, _, err := cliCtx.QueryWithData(
				fmt.Sprintf("custom/%s/last_batch/%s",
					queryRoute, bondToken), nil)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			var out types.Batch
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdCurrentPrice(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "current-price [bond-token]",
		Short: "Query current price(s) of the bond",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			bondToken := args[0]

			res, _, err := cliCtx.QueryWithData(
				fmt.Sprintf("custom/%s/current_price/%s",
					queryRoute, bondToken), nil)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			var out sdk.DecCoins
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdCurrentReserve(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:     "current-reserve [bond-token]",
		Example: "current-reserve abc",
		Short:   "Query current balance(s) of the reserve pool",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			bondToken := args[0]

			res, _, err := cliCtx.QueryWithData(
				fmt.Sprintf("custom/%s/current_reserve/%s",
					queryRoute, bondToken), nil)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			var out sdk.Coins
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdCustomPrice(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:     "price [bond-token-with-amount]",
		Example: "price 10abc",
		Short:   "Query price(s) of the bond at a specific supply",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			bondTokenWithAmount := args[0]

			bondCoinWithAmount, err := sdk.ParseCoin(bondTokenWithAmount)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			res, _, err := cliCtx.QueryWithData(
				fmt.Sprintf("custom/%s/custom_price/%s/%s",
					queryRoute, bondCoinWithAmount.Denom,
					bondCoinWithAmount.Amount.String()), nil)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			var out sdk.DecCoins
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdBuyPrice(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:     "buy-price [bond-token-with-amount]",
		Example: "buy-price 10abc",
		Short:   "Query price(s) of buying an amount of tokens of the bond",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			bondTokenWithAmount := args[0]

			bondCoinWithAmount, err := sdk.ParseCoin(bondTokenWithAmount)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			res, _, err := cliCtx.QueryWithData(
				fmt.Sprintf("custom/%s/buy_price/%s/%s",
					queryRoute, bondCoinWithAmount.Denom,
					bondCoinWithAmount.Amount.String()), nil)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			var out types.QueryBuyPrice
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdSellReturn(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:     "sell-return [bond-token-with-amount]",
		Example: "sell-return 10abc",
		Short:   "Query return(s) on selling an amount of tokens of the bond",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			bondTokenWithAmount := args[0]

			bondCoinWithAmount, err := sdk.ParseCoin(bondTokenWithAmount)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			res, _, err := cliCtx.QueryWithData(
				fmt.Sprintf("custom/%s/sell_return/%s/%s",
					queryRoute, bondCoinWithAmount.Denom,
					bondCoinWithAmount.Amount.String()), nil)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			var out types.QuerySellReturn
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdSwapReturn(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:     "swap-return [bond-token] [from-token-with-amount] [to-token]",
		Example: "swap-return abc 10res1 res2",
		Short:   "Query return(s) on swapping an amount of tokens to another token",
		Args:    cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			bondToken := args[0]
			fromTokenWithAmount := args[1]
			toToken := args[2]

			fromCoinWithAmount, err := sdk.ParseCoin(fromTokenWithAmount)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			res, _, err := cliCtx.QueryWithData(
				fmt.Sprintf("custom/%s/swap_return/%s/%s/%s/%s",
					queryRoute, bondToken, fromCoinWithAmount.Denom,
					fromCoinWithAmount.Amount.String(), toToken), nil)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return nil
			}

			var out types.QuerySwapReturn
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

// GetCmdQueryParams implements a command to fetch peyote parameters.
func GetCmdQueryParams(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "params",
		Short: "Query the current peyote parameters",
		Args:  cobra.NoArgs,
		Long: strings.TrimSpace(`Query genesis parameters for the peyote module:

$ <appcli> query peyote params
`),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			route := fmt.Sprintf("custom/%s/%s",
				types.QuerierRoute, keeper.QueryParams)
			res, _, err := cliCtx.QueryWithData(route, nil)
			if err != nil {
				return err
			}

			var params types.Params
			cdc.MustUnmarshalJSON(res, &params)
			return cliCtx.PrintOutput(params)
		},
	}
}

package rest

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"github.com/warmage-sports/peyote/x/peyote/internal/types"
	"net/http"
)

func registerQueryRoutes(cliCtx context.CLIContext, r *mux.Router, queryRoute string) {
	r.HandleFunc(
		"/peyote", queryBondsHandler(cliCtx, queryRoute),
	).Methods("GET")

	r.HandleFunc(
		fmt.Sprintf("/peyote/{%s}", RestBondToken),
		queryBondHandler(cliCtx, queryRoute),
	).Methods("GET")

	r.HandleFunc(
		fmt.Sprintf("/peyote/{%s}/batch", RestBondToken),
		queryBatchHandler(cliCtx, queryRoute),
	).Methods("GET")

	r.HandleFunc(
		fmt.Sprintf("/peyote/{%s}/last_batch", RestBondToken),
		queryLastBatchHandler(cliCtx, queryRoute),
	).Methods("GET")

	r.HandleFunc(
		fmt.Sprintf("/peyote/{%s}/current_price", RestBondToken),
		queryCurrentPriceHandler(cliCtx, queryRoute),
	).Methods("GET")

	r.HandleFunc(
		fmt.Sprintf("/peyote/{%s}/current_reserve", RestBondToken),
		queryCurrentReserveHandler(cliCtx, queryRoute),
	).Methods("GET")

	r.HandleFunc(
		fmt.Sprintf("/peyote/{%s}/price/{%s}", RestBondToken, RestBondAmount),
		queryCustomPriceHandler(cliCtx, queryRoute),
	).Methods("GET")

	r.HandleFunc(
		fmt.Sprintf("/peyote/{%s}/buy_price/{%s}", RestBondToken, RestBondAmount),
		queryBuyPriceHandler(cliCtx, queryRoute),
	).Methods("GET")

	r.HandleFunc(
		fmt.Sprintf("/peyote/{%s}/sell_return/{%s}", RestBondToken, RestBondAmount),
		querySellReturnHandler(cliCtx, queryRoute),
	).Methods("GET")

	r.HandleFunc(
		fmt.Sprintf("/peyote/{%s}/swap_return/{%s}/{%s}", RestBondToken, RestFromTokenWithAmount, RestToToken),
		querySwapReturnHandler(cliCtx, queryRoute),
	).Methods("GET")

	r.HandleFunc(
		"/peyote/params",
		queryParamsRequestHandler(cliCtx),
	).Methods("GET")
}

func queryBondsHandler(cliCtx context.CLIContext, queryRoute string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _, err := cliCtx.QueryWithData(
			fmt.Sprintf("custom/%s/peyote", queryRoute), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func queryBondHandler(cliCtx context.CLIContext, queryRoute string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bondToken := vars[RestBondToken]

		res, _, err := cliCtx.QueryWithData(
			fmt.Sprintf("custom/%s/bond/%s",
				queryRoute, bondToken), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func queryBatchHandler(cliCtx context.CLIContext, queryRoute string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bondToken := vars[RestBondToken]

		res, _, err := cliCtx.QueryWithData(
			fmt.Sprintf("custom/%s/batch/%s",
				queryRoute, bondToken), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func queryLastBatchHandler(cliCtx context.CLIContext, queryRoute string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bondToken := vars[RestBondToken]

		res, _, err := cliCtx.QueryWithData(
			fmt.Sprintf("custom/%s/last_batch/%s",
				queryRoute, bondToken), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func queryCurrentPriceHandler(cliCtx context.CLIContext, queryRoute string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bondToken := vars[RestBondToken]

		res, _, err := cliCtx.QueryWithData(
			fmt.Sprintf("custom/%s/current_price/%s",
				queryRoute, bondToken), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func queryCurrentReserveHandler(cliCtx context.CLIContext, queryRoute string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bondToken := vars[RestBondToken]

		res, _, err := cliCtx.QueryWithData(
			fmt.Sprintf("custom/%s/current_reserve/%s",
				queryRoute, bondToken), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func queryCustomPriceHandler(cliCtx context.CLIContext, queryRoute string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bondToken := vars[RestBondToken]
		bondAmount := vars[RestBondAmount]

		res, _, err := cliCtx.QueryWithData(
			fmt.Sprintf("custom/%s/custom_price/%s/%s",
				queryRoute, bondToken, bondAmount), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func queryBuyPriceHandler(cliCtx context.CLIContext, queryRoute string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bondToken := vars[RestBondToken]
		bondAmount := vars[RestBondAmount]

		res, _, err := cliCtx.QueryWithData(
			fmt.Sprintf("custom/%s/buy_price/%s/%s",
				queryRoute, bondToken, bondAmount), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func querySellReturnHandler(cliCtx context.CLIContext, queryRoute string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bondToken := vars[RestBondToken]
		bondAmount := vars[RestBondAmount]

		res, _, err := cliCtx.QueryWithData(
			fmt.Sprintf("custom/%s/sell_return/%s/%s",
				queryRoute, bondToken, bondAmount), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func querySwapReturnHandler(cliCtx context.CLIContext, queryRoute string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bondToken := vars[RestBondToken]
		fromTokenWithAmount := vars[RestFromTokenWithAmount]
		toToken := vars[RestToToken]

		reserveCoinWithAmount, err := sdk.ParseCoin(fromTokenWithAmount)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		res, _, err := cliCtx.QueryWithData(
			fmt.Sprintf("custom/%s/swap_return/%s/%s/%s/%s",
				queryRoute, bondToken, reserveCoinWithAmount.Denom,
				reserveCoinWithAmount.Amount.String(), toToken), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func queryParamsRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		route := fmt.Sprintf("custom/%s/parameters", types.QuerierRoute)

		res, height, err := cliCtx.QueryWithData(route, nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

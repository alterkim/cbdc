package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers cbdc-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	// TODO: Design API for Contract(payment, transfer, exchange)
	r.HandleFunc("/cbdc/whois", buyNameHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/cbdc/whois", listWhoisHandler(cliCtx, "cbdc")).Methods("GET")
	r.HandleFunc("/cbdc/whois/{key}", getWhoisHandler(cliCtx, "cbdc")).Methods("GET")
	r.HandleFunc("/cbdc/whois/{key}/resolve", resolveNameHandler(cliCtx, "cbdc")).Methods("GET")
	r.HandleFunc("/cbdc/whois", setWhoisHandler(cliCtx)).Methods("PUT")
	r.HandleFunc("/cbdc/whois", deleteWhoisHandler(cliCtx)).Methods("DELETE")

}

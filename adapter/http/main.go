package http

import (
	"net/http"

	"github.com/marcobagdal/finance-go/adapter/http/actuator"
	"github.com/marcobagdal/finance-go/adapter/http/transaction"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Init() {
	http.HandleFunc("/", transaction.GetTransactions)
	http.HandleFunc("/transaction", transaction.CreateTransaction)

	http.HandleFunc("/health", actuator.Health)

	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":8080", nil)
}

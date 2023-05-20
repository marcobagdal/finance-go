package transaction

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/marcobagdal/finance-go/model/transaction"
	"github.com/marcobagdal/finance-go/util"
)

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	parsedTime, err := util.StringToDateTime(time.Now())

	if err != nil {
		log.Println(err)
		http.Error(w, "Erro ao converter data", http.StatusInternalServerError)
		return
	}

	transactions := transaction.Transactions{
		{
			Title:     "Salário",
			Amount:    1200.00,
			Type:      0,
			CreatedAt: parsedTime,
		},
	}

	err = json.NewEncoder(w).Encode(transactions)

	if err != nil {
		http.Error(w, "Erro ao codificar transação em JSON", http.StatusInternalServerError)
	}
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var response transaction.Transactions

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erro ao ler requisição", http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		http.Error(w, "Erro ao codificar transação em JSON", http.StatusInternalServerError)
	}

	log.Println("Corpo da requisicao: ", response)

}

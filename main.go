package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	. "github.com/ubiq/spectrum-api/config"
	. "github.com/ubiq/spectrum-api/dao"
	. "github.com/ubiq/spectrum-api/models"
)

var config_ = Config{}
var dao_ = SpectrumDAO{}
var port string

type AccountTxn struct {
	Txns  []Transaction `bson:"txns" json:"txns"`
	Total int           `bson:"total" json:"total"`
}

type AccountTokenTransfer struct {
	Txns  []TokenTransfer `bson:"txns" json:"txns"`
	Total int             `bson:"total" json:"total"`
}

type BlockRes struct {
	Blocks []Block `bson:"blocks" json:"blocks"`
	Total  int     `bson:"total" json:"total"`
}

type UncleRes struct {
	Uncles []Uncle `bson:"uncles" json:"uncles"`
	Total  int     `bson:"total" json:"total"`
}

func getBlockByHash(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	block, err := dao_.BlockByHash(params["hash"])
	if err != nil {
		respondWithError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJson(w, r, http.StatusOK, block)
}

func getBlockByNumber(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	number, uerr := strconv.ParseUint(params["number"], 10, 64)
	if uerr != nil {
		respondWithError(w, r, http.StatusBadRequest, uerr.Error())
		return
	}
	block, err := dao_.BlockByNumber(number)
	if err != nil {
		respondWithError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJson(w, r, http.StatusOK, block)
}

func getLatestBlock(w http.ResponseWriter, r *http.Request) {
	blocks, err := dao_.LatestBlock()
	if err != nil {
		respondWithError(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, r, http.StatusOK, blocks)
}

func getLatestBlocks(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	limit, err := strconv.Atoi(params["limit"])
	if err != nil {
		respondWithError(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	if limit > 1000 {
		limit = 1000
	}
	blocks, err := dao_.LatestBlocks(limit)
	if err != nil {
		respondWithError(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	count, err := dao_.TotalBlockCount()
	if err != nil {
		respondWithError(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	var res BlockRes
	res.Blocks = blocks
	res.Total = count

	respondWithJson(w, r, http.StatusOK, res)
}

func getLatestForkedBlocks(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	limit, err := strconv.Atoi(params["limit"])
	if err != nil {
		respondWithError(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	if limit > 1000 {
		limit = 1000
	}
	blocks, err := dao_.LatestForkedBlocks(limit)
	if err != nil {
		respondWithError(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, r, http.StatusOK, blocks)
}

func getLatestTransactions(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	limit, err := strconv.Atoi(params["limit"])
	if err != nil {
		respondWithError(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	if limit > 1000 {
		limit = 1000
	}
	txns, err := dao_.LatestTransactions(limit)
	if err != nil {
		respondWithError(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	count, err := dao_.TotalTxnCount()
	if err != nil {
		respondWithError(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	var res AccountTxn
	res.Txns = txns
	res.Total = count

	respondWithJson(w, r, http.StatusOK, res)
}

func getLatestTransactionsByAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	txns, err := dao_.LatestTransactionsByAccount(params["hash"])
	if err != nil {
		respondWithError(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	count, err := dao_.TxnCount(params["hash"])
	if err != nil {
		respondWithError(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	var res AccountTxn
	res.Txns = txns
	res.Total = count

	respondWithJson(w, r, http.StatusOK, res)
}

func getTransactionsByBlockNumber(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	number, uerr := strconv.ParseUint(params["number"], 10, 64)
	if uerr != nil {
		respondWithError(w, r, http.StatusBadRequest, uerr.Error())
		return
	}
	txns, err := dao_.TransactionsByBlockNumber(number)
	if err != nil {
		respondWithError(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, r, http.StatusOK, txns)
}

func getLatestTokenTransfersByAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	txns, err := dao_.LatestTokenTransfersByAccount(params["hash"])
	if err != nil {
		respondWithError(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	count, err := dao_.TokenTransferCount(params["hash"])
	if err != nil {
		respondWithError(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	var res AccountTokenTransfer
	res.Txns = txns
	res.Total = count

	respondWithJson(w, r, http.StatusOK, res)
}

func getLatestTokenTransfers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	limit, err := strconv.Atoi(params["limit"])
	if err != nil {
		respondWithError(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	if limit > 1000 {
		limit = 1000
	}

	transfers, err := dao_.LatestTokenTransfers(limit)
	if err != nil {
		respondWithError(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	count, err := dao_.TotalTokenTransferCount()
	if err != nil {
		respondWithError(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	var res AccountTokenTransfer
	res.Txns = transfers
	res.Total = count

	respondWithJson(w, r, http.StatusOK, res)
}

func getLatestUncles(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	limit, err := strconv.Atoi(params["limit"])
	if err != nil {
		respondWithError(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	if limit > 1000 {
		limit = 1000
	}
	uncles, err := dao_.LatestUncles(limit)
	if err != nil {
		respondWithError(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	count, err := dao_.TotalUncleCount()
	if err != nil {
		respondWithError(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	var res UncleRes
	res.Uncles = uncles
	res.Total = count

	respondWithJson(w, r, http.StatusOK, res)
}

func getTransactionByHash(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	txn, err := dao_.TransactionByHash(params["hash"])
	if err != nil {
		respondWithError(w, r, http.StatusOK, err.Error())
		return
	}
	respondWithJson(w, r, http.StatusOK, txn)
}

func getTransactionByContractAddress(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	txn, err := dao_.TransactionByContractAddress(params["hash"])
	if err != nil {
		respondWithError(w, r, http.StatusOK, err.Error())
		return
	}
	respondWithJson(w, r, http.StatusOK, txn)
}

func getUncleByHash(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	uncle, err := dao_.UncleByHash(params["hash"])
	if err != nil {
		respondWithError(w, r, http.StatusOK, err.Error())
		return
	}
	respondWithJson(w, r, http.StatusOK, uncle)
}

func getStore(w http.ResponseWriter, r *http.Request) {
	store, err := dao_.Store()
	if err != nil {
		respondWithError(w, r, http.StatusOK, err.Error())
		return
	}
	respondWithJson(w, r, http.StatusOK, store)
}

func respondWithError(w http.ResponseWriter, r *http.Request, code int, msg string) {

	log.WithFields(log.Fields{
		"path":  r.URL,
		"ip":    r.RemoteAddr,
		"error": msg,
	}).Error()

	respondWithJson(w, r, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {

	response, err := json.Marshal(payload)

	if err != nil {
		log.Errorf("Could not marshal response: %v", err)
	}

	log.WithFields(log.Fields{
		"path": r.URL,
		"ip":   r.RemoteAddr,
	}).Info()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func init() {
	config_.Read()

	port = config_.Port
	dao_.Server = config_.Server
	dao_.Database = config_.Database
	dao_.Connect()
}

func init() {
	log.SetFormatter(&log.TextFormatter{DisableLevelTruncation: true, FullTimestamp: true, TimestampFormat: time.RFC822})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	log.Info("Api started on port ", port)

	r := mux.NewRouter()
	r.HandleFunc("/status", getStore).Methods("GET")
	r.HandleFunc("/block/{number}", getBlockByNumber).Methods("GET")
	r.HandleFunc("/blockbyhash/{hash}", getBlockByHash).Methods("GET")
	r.HandleFunc("/blocktransactions/{number}", getTransactionsByBlockNumber).Methods("GET")
	r.HandleFunc("/latest", getLatestBlock).Methods("GET")
	r.HandleFunc("/latestblocks/{limit}", getLatestBlocks).Methods("GET")
	r.HandleFunc("/latestforkedblocks/{limit}", getLatestForkedBlocks).Methods("GET")
	r.HandleFunc("/latesttransactions/{limit}", getLatestTransactions).Methods("GET")
	r.HandleFunc("/latestaccounttxns/{hash}", getLatestTransactionsByAccount).Methods("GET")
	r.HandleFunc("/latestaccounttokentxns/{hash}", getLatestTokenTransfersByAccount).Methods("GET")
	r.HandleFunc("/latesttokentransfers/{limit}", getLatestTokenTransfers).Methods("GET")
	r.HandleFunc("/latestuncles/{limit}", getLatestUncles).Methods("GET")
	r.HandleFunc("/transaction/{hash}", getTransactionByHash).Methods("GET")
	r.HandleFunc("/transactionbycontract/{hash}", getTransactionByContractAddress).Methods("GET")
	r.HandleFunc("/uncle/{hash}", getUncleByHash).Methods("GET")

	handler := cors.Default().Handler(r)
	if err := http.ListenAndServe(port, handler); err != nil {
		log.Fatal(err)
	}
}

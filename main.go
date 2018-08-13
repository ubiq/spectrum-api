package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/rs/cors"
	"github.com/gorilla/mux"
	. "github.com/ubiq/spectrum-api/config"
	. "github.com/ubiq/spectrum-api/dao"
)

var config_ = Config{}
var dao_ = SpectrumDAO{}
var port string

func getBlockByHash(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	block, err := dao_.BlockByHash(params["hash"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, block)
}

func getBlockByNumber(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	number, uerr := strconv.ParseUint(params["number"], 10, 64)
	if uerr != nil {
		respondWithError(w, http.StatusBadRequest, uerr.Error())
		return
	}
	block, err := dao_.BlockByNumber(number)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, block)
}

func getLatestBlock(w http.ResponseWriter, r *http.Request) {
	blocks, err := dao_.LatestBlock()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, blocks)
}

func getLatestBlocks(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	limit, err := strconv.Atoi(params["limit"])
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if limit > 1000 {
		limit = 1000
	}
	blocks, err := dao_.LatestBlocks(limit)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, blocks)
}

func getLatestTransactions(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	limit, err := strconv.Atoi(params["limit"])
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if limit > 1000 {
		limit = 1000
	}
	txns, err := dao_.LatestTransactions(limit)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, txns)
}

func getLatestUncles(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	limit, err := strconv.Atoi(params["limit"])
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if limit > 1000 {
		limit = 1000
	}
	uncles, err := dao_.LatestUncles(limit)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, uncles)
}

func getTransactionByHash(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	txn, err := dao_.TransactionByHash(params["hash"])
	if err != nil {
		respondWithError(w, http.StatusOK, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, txn)
}

func getUncleByHash(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	uncle, err := dao_.UncleByHash(params["hash"])
	if err != nil {
		respondWithError(w, http.StatusOK, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, uncle)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
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

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/block/{number}", getBlockByNumber).Methods("GET")
	r.HandleFunc("/blockbyhash/{hash}", getBlockByHash).Methods("GET")
	r.HandleFunc("/latest", getLatestBlock).Methods("GET")
	r.HandleFunc("/latestblocks/{limit}", getLatestBlocks).Methods("GET")
	r.HandleFunc("/latesttransactions/{limit}", getLatestTransactions).Methods("GET")
	r.HandleFunc("/latestuncles/{limit}", getLatestUncles).Methods("GET")
	r.HandleFunc("/transaction/{hash}", getTransactionByHash).Methods("GET")
	r.HandleFunc("/uncle/{hash}", getUncleByHash).Methods("GET")

	handler := cors.Default().Handler(r)
	if err := http.ListenAndServe(port, handler); err != nil {
		log.Fatal(err)
	}
}

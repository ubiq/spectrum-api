package dao

import (
  "log"

  . "github.com/ubiq/spectrum-api/models"
  mgo "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)

type SpectrumDAO struct {
  Server    string
  Database  string
}

var db *mgo.Database

const (
  BLOCKS = "blocks"
  TXNS = "transactions"
  UNCLES = "uncles"
  TRANSFERS = "tokentransfers"
)

func (e *SpectrumDAO) Connect() {
  session, err := mgo.Dial(e.Server)
  if err != nil {
    log.Fatal(err)
  }
  db = session.DB(e.Database)
}

func (e *SpectrumDAO) BlockByNumber(number uint64) (Block, error) {
  var block Block
  err := db.C(BLOCKS).Find(bson.M{"number": number}).One(&block)
  return block, err
}

func (e *SpectrumDAO) BlockByHash(hash string) (Block, error) {
  var block Block
  err := db.C(BLOCKS).Find(bson.M{"hash": hash}).One(&block)
  return block, err
}

func (e *SpectrumDAO) LatestBlock() (Block, error) {
  var block Block
  err := db.C(BLOCKS).Find(bson.M{}).Sort("-number").Limit(1).One(&block)
  return block, err
}

func (e *SpectrumDAO) LatestBlocks(limit int) ([]Block, error) {
  var blocks []Block
  err := db.C(BLOCKS).Find(bson.M{}).Sort("-number").Limit(limit).All(&blocks)
  return blocks, err
}

func (e *SpectrumDAO) LatestUncles(limit int) ([]Uncle, error) {
  var uncles []Uncle
  err := db.C(UNCLES).Find(bson.M{}).Sort("-blockNumber").Limit(limit).All(&uncles)
  return uncles, err
}

func (e *SpectrumDAO) TransactionByHash(hash string) (Transaction, error) {
  var txn Transaction
  err := db.C(TXNS).Find(bson.M{"hash": hash}).One(&txn)
  return txn, err
}

func (e *SpectrumDAO) UncleByHash(hash string) (Uncle, error) {
  var uncle Uncle
  err := db.C(UNCLES).Find(bson.M{"hash": hash}).One(&uncle)
  return uncle, err
}

func (e *SpectrumDAO) LatestTransactions(limit int) ([]Transaction, error) {
  var txns []Transaction
  err := db.C(TXNS).Find(bson.M{}).Sort("-blockNumber").Limit(limit).All(&txns)
  return txns, err
}

func (e *SpectrumDAO) LatestTokenTransfers(limit int) ([]TokenTransfer, error) {
  var transfers []TokenTransfer
  err := db.C(TRANSFERS).Find(bson.M{}).Sort("-blockNumber").Limit(limit).All(&transfers)
  return transfers, err
}

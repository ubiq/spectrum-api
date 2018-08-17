package models

type Block struct {
	Number           uint64  `bson:"number" json:"number"`
	Timestamp        uint64  `bson:"timestamp" json:"timestamp"`
	Transactions     uint64  `bson:"transactions" json:"transactions"`
	Hash             string  `bson:"hash" json:"hash"`
	ParentHash       string  `bson:"parentHash" json:"parentHash"`
	Sha3Uncles       string  `bson:"sha3Uncles" json:"sha3Uncles"`
	Miner            string  `bson:"miner" json:"miner"`
	Difficulty       string  `bson:"difficulty" json:"difficulty"`
	TotalDifficulty  string  `bson:"totalDifficulty" json:"totalDifficulty"`
	Size             uint64  `bson:"size" json:"size"`
	GasUsed          uint64  `bson:"gasUsed" json:"gasUsed"`
	GasLimit         uint64  `bson:"gasLimit" json:"gasLimit"`
	Nonce            string  `bson:"nonce" json:"nonce"`
	Uncles           uint64  `bson:"uncles" json:"uncles"`
	BlockReward      string  `bson:"blockReward" json:"blockReward"`
	UnclesReward     string  `bson:"unclesReward" json:"unclesReward"`
	AvgGasPrice      string  `bson:"avgGasPrice" json:"avgGasPrice"`
	TxFees           string  `bson:"txFees" json:"txFees"`
	ExtraData        string  `bson:"extraData" json:"extraData"`
}

type TxLog struct {
	Address          string    `bson:"address" json:"address"`
	Topics           []string  `bson:"topics" json:"topics"`
	Data             string    `bson:"data" json:"data"`
	BlockNumber      uint64    `bson:"blockNumber" json:"blockNumber"`
	TransactionIndex uint64    `bson:"transactionIndex" json:"transactionIndex"`
	TransactionHash  string    `bson:"transactionHash" json:"transactionHash"`
	BlockHash        string    `bson:"blockHash" json:"blockHash"`
	LogIndex         uint64    `bson:"logIndex" json:"logIndex"`
	Removed          bool      `bson:"removed" json:"removed"`
}

type Transaction struct {
	BlockHash        string  `bson:"blockHash" json:"blockHash"`
	BlockNumber      uint64  `bson:"blockNumber" json:"blockNumber"`
	Hash             string  `bson:"hash" json:"hash"`
	Timestamp        uint64  `bson:"timestamp" json:"timestamp"`
	Input            string  `bson:"input" json:"input"`
	Value            string  `bson:"value" json:"value"`
	Gas              uint64  `bson:"gas" json:"gas"`
	GasUsed          uint64  `bson:"gasUsed" json:"gasUsed"`
	GasPrice         string  `bson:"gasPrice" json:"gasPrice"`
	Nonce            uint64  `bson:"nonce" json:"nonce"`
	TransactionIndex uint64  `bson:"transactionIndex" json:"transactionIndex"`
	From             string  `bson:"from" json:"from"`
	To               string  `bson:"to" json:"to"`
	ContractAddress  string  `bson:"contractAddress" json:"contractAddress"`
	Logs             []TxLog `bson:"logs" json:"logs"`
}

type TokenTransfer struct {
	BlockNumber      uint64  `bson:"blockNumber" json:"blockNumber"`
	Hash             string  `bson:"hash" json:"hash"`
	Timestamp        uint64  `bson:"timestamp" json:"timestamp"`
	From             string  `bson:"from" json:"from"`
	To               string  `bson:"to" json:"to"`
	Value            string  `bson:"value" json:"value"`
	Contract         string  `bson:"contract" json:"contract"`
	Method           string  `bson:"method" json:"method"`
}

type Uncle struct {
	Number           uint64  `bson:"number" json:"number"`
	Position         uint64  `bson:"position" json:"position"`
	BlockNumber      uint64  `bson:"blockNumber" json:"blockNumber"`
	Hash             string  `bson:"hash" json:"hash"`
	ParentHash       string  `bson:"parentHash" json:"parentHash"`
	Sha3Uncles       string  `bson:"sha3Uncles" json:"sha3Uncles"`
	Miner            string  `bson:"miner" json:"miner"`
	Difficulty       string  `bson:"difficulty" json:"difficulty"`
	GasUsed          uint64  `bson:"gasUsed" json:"gasUsed"`
	GasLimit         uint64  `bson:"gasLimit" json:"gasLimit"`
	Timestamp        uint64  `bson:"timestamp" json:"timestamp"`
	Reward           string  `bson:"reward" json:"reward"`
}

type Store struct {
	Timestamp        uint64  `bson:"timestamp" json:"timestamp"`
	Symbol           string  `bson:"symbol" json:"symbol"`
	Supply           string  `bson:"supply" json:"supply"`
	LatestBlock      Block   `bson:"latestBlock" json:"latestBlock"`
	Price            string  `bson:"price" json:"price"`
}

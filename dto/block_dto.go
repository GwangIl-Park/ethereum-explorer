package dto

type Block struct {
	BlockHeight string
	//Status      bool
	Timestamp  uint64 
	Receipient string
	//Reward
	Size       uint64
	GasUsed    uint64
	GasLimit   uint64
	BaseFee    string
	ExtraData  string
	Hash       string
	ParentHash string
	TransactionCount uint64
}

type GetBlocksDTO struct {
	Blocks []Block
}

type GetBlockHeightsDTO struct {
	Heights []string
}

type GetBlockByHeightDTO struct {
	Block Block
}

func (gbDTO GetBlocksDTO) GetDTO() interface{} {
	return gbDTO
}

func (gbhDTO GetBlockHeightsDTO) GetDTO() interface{} {
	return gbhDTO
}

func (gbbhDTO GetBlockByHeightDTO) GetDTO() interface{} {
	return gbbhDTO
}
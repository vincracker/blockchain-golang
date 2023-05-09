package blockchain

import "math/big"

const difficulty = 12

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

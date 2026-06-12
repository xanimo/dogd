// Copyright (c) 2014-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package chaincfg

import (
	"time"

	"github.com/btcsuite/btcd/chainhash/v2"
	"github.com/btcsuite/btcd/wire/v2"
)

// genesisCoinbaseTx is the coinbase transaction for the genesis blocks for
// the main network, regression test network, and test network (version 3).
// The pszTimestamp embedded is "Nintondo".
var genesisCoinbaseTx = wire.MsgTx{
	Version: 1,
	TxIn: []*wire.TxIn{
		{
			PreviousOutPoint: wire.OutPoint{
				Hash:  chainhash.Hash{},
				Index: 0xffffffff,
			},
			SignatureScript: []byte{
				0x04, 0xff, 0xff, 0x00, 0x1d, 0x01, 0x04, 0x08, /* |........| */
				0x4e, 0x69, 0x6e, 0x74, 0x6f, 0x6e, 0x64, 0x6f, /* |Nintondo| */
			},
			Sequence: 0xffffffff,
		},
	},
	TxOut: []*wire.TxOut{
		{
			Value: 0x20C855800, // 88 DOGE
			PkScript: []byte{
				0x41, 0x04, 0x01, 0x84, 0x71, 0x0f, 0xa6, 0x89, /* |A...q...| */
				0xad, 0x50, 0x23, 0x69, 0x0c, 0x80, 0xf3, 0xa4, /* |.P#i....| */
				0x9c, 0x8f, 0x13, 0xf8, 0xd4, 0x5b, 0x8c, 0x85, /* |.....[..| */
				0x7f, 0xbc, 0xbc, 0x8b, 0xc4, 0xa8, 0xe4, 0xd3, /* |........| */
				0xeb, 0x4b, 0x10, 0xf4, 0xd4, 0x60, 0x4f, 0xa0, /* |.K...`O.| */
				0x8d, 0xce, 0x60, 0x1a, 0xaf, 0x0f, 0x47, 0x02, /* |..`...G.| */
				0x16, 0xfe, 0x1b, 0x51, 0x85, 0x0b, 0x4a, 0xcf, /* |...Q..J.| */
				0x21, 0xb1, 0x79, 0xc4, 0x50, 0x70, 0xac, 0x7b, /* |!.y.Pp.{| */
				0x03, 0xa9, 0xac,                                /* |...| */
			},
		},
	},
	LockTime: 0,
}

// genesisHash is the hash of the first block in the block chain for the main
// network (genesis block).
// 1a91e3dace36e2be3bf030a65679fe821aa1d6ef92e7c9902eb318182c355691
var genesisHash = chainhash.Hash([chainhash.HashSize]byte{
	0x91, 0x56, 0x35, 0x2c, 0x18, 0x18, 0xb3, 0x2e,
	0x90, 0xc9, 0xe7, 0x92, 0xef, 0xd6, 0xa1, 0x1a,
	0x82, 0xfe, 0x79, 0x56, 0xa6, 0x30, 0xf0, 0x3b,
	0xbe, 0xe2, 0x36, 0xce, 0xda, 0xe3, 0x91, 0x1a,
})

// genesisMerkleRoot is the hash of the first transaction in the genesis block
// for the main network.
// 5b2a3f53f605d62c53e62932dac6925e3d74afa5a4b459745c36d42d0ed26a69
var genesisMerkleRoot = chainhash.Hash([chainhash.HashSize]byte{
	0x69, 0x6a, 0xd2, 0x0e, 0x2d, 0xd4, 0x36, 0x5c,
	0x74, 0x59, 0xb4, 0xa4, 0xa5, 0xaf, 0x74, 0x3d,
	0x5e, 0x92, 0xc6, 0xda, 0x32, 0x29, 0xe6, 0x53,
	0x2c, 0xd6, 0x05, 0xf6, 0x53, 0x3f, 0x2a, 0x5b,
})

// genesisBlock defines the genesis block of the block chain which serves as the
// public transaction ledger for the main network.
var genesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  chainhash.Hash{},
		MerkleRoot: genesisMerkleRoot,
		Timestamp:  time.Unix(1386325540, 0), // 2013-12-06 10:25:40 +0000 UTC
		Bits:       0x1e0ffff0,
		Nonce:      99943,
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// regTestGenesisHash is the hash of the first block in the block chain for the
// regression test network (genesis block).
// 3d2160a3b5dc4a9d62e7e66a295f70313ac808440ef7400d6c0772171ce973a5
var regTestGenesisHash = chainhash.Hash([chainhash.HashSize]byte{
	0xa5, 0x73, 0xe9, 0x1c, 0x17, 0x72, 0x07, 0x6c,
	0x0d, 0x40, 0xf7, 0x0e, 0x44, 0x08, 0xc8, 0x3a,
	0x31, 0x70, 0x5f, 0x29, 0x6a, 0xe6, 0xe7, 0x62,
	0x9d, 0x4a, 0xdc, 0xb5, 0xa3, 0x60, 0x21, 0x3d,
})

// regTestGenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the regression test network. It is the same as the merkle root for
// the main network.
var regTestGenesisMerkleRoot = genesisMerkleRoot

// regTestGenesisBlock defines the genesis block of the block chain which serves
// as the public transaction ledger for the regression test network.
var regTestGenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  chainhash.Hash{},
		MerkleRoot: regTestGenesisMerkleRoot,
		Timestamp:  time.Unix(1296688602, 0), // 2011-02-02 23:16:42 +0000 UTC
		Bits:       0x207fffff,
		Nonce:      2,
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// testNet3GenesisHash is the hash of the first block in the block chain for the
// Dogecoin test network.
// bb0a78264637406b6360aad926284d544d7049f45189db5664f3c4d07350559e
var testNet3GenesisHash = chainhash.Hash([chainhash.HashSize]byte{
	0x9e, 0x55, 0x50, 0x73, 0xd0, 0xc4, 0xf3, 0x64,
	0x56, 0xdb, 0x89, 0x51, 0xf4, 0x49, 0x70, 0x4d,
	0x54, 0x4d, 0x28, 0x26, 0xd9, 0xaa, 0x60, 0x63,
	0x6b, 0x40, 0x37, 0x46, 0x26, 0x78, 0x0a, 0xbb,
})

// testNet3GenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the Dogecoin test network. It is the same as the merkle root for
// the main network.
var testNet3GenesisMerkleRoot = genesisMerkleRoot

// testNet3GenesisBlock defines the genesis block of the block chain which
// serves as the public transaction ledger for the Dogecoin test network.
var testNet3GenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  chainhash.Hash{},
		MerkleRoot: testNet3GenesisMerkleRoot,
		Timestamp:  time.Unix(1391503289, 0), // 2014-02-04 09:28:09 +0000 UTC
		Bits:       0x1e0ffff0,
		Nonce:      997879,
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// simNetGenesisHash is the hash of the first block in the block chain for the
// simulation test network (computed from Dogecoin coinbase, timestamp=1401292357,
// bits=0x207fffff, nonce=2).
// a6e6f89bdde5b3829d363402ca1d38929a1d363beb32aeca6bf3b906749c741a
var simNetGenesisHash = chainhash.Hash([chainhash.HashSize]byte{
	0x1a, 0x74, 0x9c, 0x74, 0x06, 0xb9, 0xf3, 0x6b,
	0xca, 0xae, 0x32, 0xeb, 0x3b, 0x36, 0x1d, 0x9a,
	0x92, 0x38, 0x1d, 0xca, 0x02, 0x34, 0x36, 0x9d,
	0x82, 0xb3, 0xe5, 0xdd, 0x9b, 0xf8, 0xe6, 0xa6,
})

// simNetGenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the simulation test network. It is the same as the merkle root for
// the main network.
var simNetGenesisMerkleRoot = genesisMerkleRoot

// simNetGenesisBlock defines the genesis block of the block chain which serves
// as the public transaction ledger for the simulation test network.
var simNetGenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  chainhash.Hash{},
		MerkleRoot: simNetGenesisMerkleRoot,
		Timestamp:  time.Unix(1401292357, 0), // 2014-05-28 15:52:37 +0000 UTC
		Bits:       0x207fffff,
		Nonce:      2,
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

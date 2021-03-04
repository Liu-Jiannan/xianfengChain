package consensus

import (
	"../utils"
	"bytes"
	"crypto/sha256"
	"math/big"
)

//目的：拿到区块的属性数据(属性值)
	//1.通过结构体引用，引用block结构体,然后访问其属性,比如block.hash
	//2、接口


const DIFFICULTY = 20 //挖矿难度值系数

type PoW struct {
	Block  BlockInterface
	Target *big.Int
}

func (pow PoW) FindNonce() ([32]byte,int64) {
	//1、给定一个nonce,计算hash
	var nonce int64
	nonce = 0

	//无限循环
	hashBig := new(big.Int)
	for {
		hash:=CalculateHash(pow.Block,nonce)

		//32字节->256位
		//2、拿到系统目标值
		target := pow.Target

		//3、比较大小
		hashBig = hashBig.SetBytes(hash[:])
		//result := bytes.Compare(hash[:], target.Bytes())
		result := hashBig.Cmp(target)
		if result == -1 {
			return hash,nonce
		}
		nonce++ //再计算一次
	}
}
/**
	根据区块已有的信息和当前nonce的赋值，计算区块的hash
 */
func CalculateHash(block BlockInterface,nonce int64) [32]byte {
	heightByte,_:= utils.Int2Byte(block.GetHeight())
	versionByte,_:=utils.Int2Byte(block.GetVersion())
	timeByte,_:=utils.Int2Byte(block.GetTimeStamp())
	nonceByte,_:=utils.Int2Byte(nonce)

	prev := block.GetPrevHash()

	blockByte:=bytes.Join([][]byte{heightByte,versionByte,prev[:], timeByte, nonceByte, block.GetData()},[]byte{})

	//计算区块的hash
	hash := sha256.Sum256(blockByte)
	return hash
}

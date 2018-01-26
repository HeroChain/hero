package main

import "bytes"

//TxInput represents a transaction input 
type TXInput struct {
    Txid    []byte
    Vout    int
    Signature []byte
    PubKey  []byte
}

// UseKey checks whether the address initiated the transaction
func (in *TXInput) UseKey(pubKeyHash []byte) bool {
    lockingHash := HashPubKey(in.PubKey)
    return bytes.Compare(lockingHash, pubKeyHash) == 0
}
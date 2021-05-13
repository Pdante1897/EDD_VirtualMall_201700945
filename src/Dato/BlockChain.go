package Dato

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	Indice   int
	Fecha    string
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

type BlockChain struct {
	Blocks []*Block
}

func (this *Block) DerivarHash() {
	ind := strconv.Itoa(this.Indice)
	nonce := strconv.Itoa(int(this.Nonce))

	inf := bytes.Join([][]byte{[]byte(ind), []byte(this.Fecha), this.PrevHash, this.Data, []byte(nonce)}, []byte{})
	hash := sha256.Sum256(inf)
	this.Hash = hash[:]
}

func (this *BlockChain) AgregarBlock(data string) {
	var nuevo *Block
	if len(this.Blocks) != 0 {
		prevBlock := this.Blocks[len(this.Blocks)-1]
		nuevo = CrearBlock(len(this.Blocks)+1, data, prevBlock.Hash)

	} else {
		nuevo = CrearBlock(len(this.Blocks)+1, data, []byte{})
	}
	this.Blocks = append(this.Blocks, nuevo)
}

func Genesis() *Block {
	return CrearBlock(0, "Inicial", []byte{})
}

func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{}}
}

func CrearBlock(indice int, data string, prevHash []byte) *Block {
	var today time.Time
	today = time.Now()
	var fecha strings.Builder
	fmt.Fprintf(&fecha, "%v-%v-%v::%v:%v:%v", today.Day(), today.Month(), today.Year(), today.Hour(), today.Minute(), today.Second())

	block := Block{indice, fecha.String(), []byte{}, []byte(data), prevHash, 0}
	pow := NewProof(&block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	block.DirBlock()
	return &block
}

func (this *Block) DirBlock() {
	var cadena strings.Builder
	fmt.Fprintf(&cadena, "Indice: %v \n", this.Indice)
	fmt.Fprintf(&cadena, "Fecha: %s \n", this.Fecha)
	fmt.Fprintf(&cadena, "Data: %s \n", this.Data)
	fmt.Fprintf(&cadena, "PrevHash: %x \n", this.PrevHash)
	fmt.Fprintf(&cadena, "Hash: %x \n", this.Hash)
	fmt.Fprintf(&cadena, "Nonce: %v \n", this.Nonce)
	dir, err := filepath.Abs(filepath.Dir("./"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	err = os.MkdirAll(dir+"/bloques", 0777)
	fil, err := os.Create(dir + "\\bloques\\" + strconv.Itoa(this.Indice) + ".txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	bytes, err := fil.WriteString(cadena.String())
	if err != nil {
		fmt.Println(err)
		fil.Close()
		return
	}
	fmt.Println(bytes, "bytes escritos satisfactoriamente! :D")
	err = fil.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

}

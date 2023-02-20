package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Block struct {
	Pos       int
	Data      BookCheckout
	Timestamp string
	Hash      string
	PrevHash  string
}

type Book struct {
	ID          string `json:"id"`
	ISBN        string `json:"isbn"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	PublishDate string `json:"publish_date"`
}

type BookCheckout struct {
	BookID       string `json:"book_id"`
	User         string `json:"user"`
	CheckoutDate string `json:"checkout_date"`
	IsGenesis    bool   `json:"is_genesis"`
}

type Blockchain struct {
	blocks []*Block
}

func (b *Block) generateHash() {
	bytes, _ := json.Marshal(b.Data)

	data := strconv.Itoa(b.Pos) + b.Timestamp + string(bytes)

	hash := sha256.New()
	hash.Write([]byte(data))
	b.Hash = hex.EncodeToString(hash.Sum(nil))
}

func (b *Block) validateHash(hash string) bool {
	b.generateHash()
	return b.Hash == hash
}

func (bc *Blockchain) AddBlock(data BookCheckout) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	block := CreateBlock(prevBlock, data)
	if validBlock(block, prevBlock) {
		bc.blocks = append(bc.blocks, block)
	}
}

func CreateBlock(prevBlock *Block, data BookCheckout) *Block {
	block := new(Block)
	block.Pos = prevBlock.Pos + 1
	block.Data = data
	block.Timestamp = time.Now().String()
	block.PrevHash = prevBlock.Hash
	block.generateHash()

	return block
}

func validBlock(block *Block, prevBlock *Block) bool {
	if block.PrevHash != prevBlock.Hash {
		return false
	}
	if !block.validateHash(block.Hash) {
		return false
	}
	if prevBlock.Pos+1 != block.Pos {
		return false
	}
	return true
}

func newBook(w http.ResponseWriter, r *http.Request) {
	var book Book

	// fmt.Println(r)
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("could not create: %v\n", err)
		w.Write([]byte("could not create new book"))
		return
	}

	h := md5.New()
	io.WriteString(h, book.ISBN+book.PublishDate)
	book.ID = fmt.Sprintf("%x", h.Sum(nil))
	response, err := json.MarshalIndent(book, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("could not marshall payload: %v\n", err)
		w.Write([]byte("could not savde book data"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func writeBlock(w http.ResponseWriter, r *http.Request) {
	var checkoutItem BookCheckout

	if err := json.NewDecoder(r.Body).Decode(&checkoutItem); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Could not write block: %v", err)
		w.Write([]byte("Could not write block"))
		return
	}

	blockchain.AddBlock(checkoutItem)
}

func getBlockchain(w http.ResponseWriter, r *http.Request) {
	jbytes, err := json.MarshalIndent(blockchain.blocks, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	io.WriteString(w, string(jbytes))
}

var blockchain *Blockchain

func GenesisBlock() *Block {
	return CreateBlock(new(Block), BookCheckout{IsGenesis: true})
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{GenesisBlock()}}
}

func main() {
	blockchain = NewBlockchain()
	r := mux.NewRouter()
	r.HandleFunc("/", getBlockchain).Methods("GET")
	r.HandleFunc("/", writeBlock).Methods("POST")
	r.HandleFunc("/new", newBook).Methods("POST")

	go func() {
		for _, block := range blockchain.blocks {
			fmt.Printf("Prev.hash: %x\n", block.PrevHash)
			bytes, _ := json.MarshalIndent(block.Data, "", " ")
			fmt.Printf("Data:%v\n", string(bytes))
			fmt.Printf("Hash:%x\n", block.Hash)
			fmt.Println()
		}
	}()

	log.Println("Listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}

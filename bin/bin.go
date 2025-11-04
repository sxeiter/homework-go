package bin

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"
)

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
}

type BinList struct {
	Bins []Bin
}

func NewBin(name string, private bool) *Bin {
	return &Bin{
		Id:        generateId(),
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}
}

func generateId() string {
	bytes := make([]byte, 8)
	if _, err := rand.Read(bytes); err != nil {
		return fmt.Sprintf("bin_%d", time.Now().UnixNano())
	}
	return "bin_" + hex.EncodeToString(bytes)
}

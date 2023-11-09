package controllers

import (
	"time"

	"github.com/astaxie/beego"
)

type PeerController struct {
	beego.Controller
}

type PeerChain struct {
	Chain      []Block
	PeerBlocks map[string][]int
	Difficulty int
}

type Peer struct {
	Index      int
	PrevHash   string
	Data       []string
	Timestamp  time.Time
	Difficulty int
}

func NewPeerchain() *PeerChain {
	pc := &PeerChain{}
	pc.Difficulty = 0x1e11ffff
	pc.CreateGenesisPeer()
	return pc
}

func (pc *PeerChain) CreateGenesisPeer() {
	genesisBlock := Block{
		Index:      len(pc.Chain),
		PrevHash:   "pc.Chain[len(pc.Chain)-1].PrevHash",
		Data:       []string{"hello world!"},
		Timestamp:  time.Now().String(),
		Difficulty: pc.Difficulty,
	}
	pc.Chain = append(pc.Chain, genesisBlock)
}

func (c *PeerController) AddBlock() {
	pc := NewPeerchain()
	newBlock := Block{
		Index:      len(pc.Chain),
		PrevHash:   pc.Chain[len(pc.Chain)-1].PrevHash,
		Data:       []string{"new data"},
		Timestamp:  time.Now().String(),
		Difficulty: pc.Difficulty,
	}
	pc.Chain = append(pc.Chain, newBlock)

	c.Data["json"] = map[string]string{"message": "Block added successfully"}
	c.ServeJSON()
}

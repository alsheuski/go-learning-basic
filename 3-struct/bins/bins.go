// Package bins contains Bin and BinList structures
package bins

import "time"

type Bin struct {
	ID        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
}

func NewBin(id string, private bool, createdAt time.Time, name string) *Bin {
	return &Bin{
		id,
		private,
		createdAt,
		name,
	}
}

type BinList struct {
	List []Bin `json:"list"`
}

func NewBinList(list []Bin) *BinList {
	return &BinList{list}
}

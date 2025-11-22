package main

import "time"

type Bin struct {
	ID        string
	Private   bool
	CreatedAt time.Time
	Name      string
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
	List []Bin
}

func NewBinList(list []Bin) *BinList {
	return &BinList{list}
}

func main() {
}

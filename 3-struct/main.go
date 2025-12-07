package main

import (
	"fmt"
	"time"

	"purple.learn/3-struct/bins"
	"purple.learn/3-struct/storage"
)

func main() {
	storage := storage.NewJSONStorage()
	bin := bins.NewBin("test-1", false, time.Now(), "My First Bin")
	binList := bins.NewBinList([]bins.Bin{*bin})
	storage.SaveToJSONFile(*binList)

	fmt.Println("Saved successfully")
}

package main

import (
	"exercises/7-phnorm/norm"
	"exercises/7-phnorm/norm/db"
	"fmt"
)

func main() {
	pnDB := db.SetupDB()
	pns := db.GetAllPNs(pnDB)
	fmt.Println(len(pns), "numbers in the DB")
	pns = norm.NormalizeAll(pns)
	db.UpdatePNs(pns, pnDB)
	db.RemoveDuplicatePNs(pns, pnDB)
	newPNs := db.GetAllPNs(pnDB)
	fmt.Println("After removing duplicates & normalizing", len(newPNs), "numbers in the db")
	defer pnDB.Close()
}

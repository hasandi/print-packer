package main

import (
	"fmt"
	"github.com/hasandi/print-packer/internal/print"
)

func main() {
	machines := []*print.Machine{
		print.NewMachine(360, 520),
		print.NewMachine(520, 740),
		print.NewMachine(740, 1020),
	}
	papers := []*print.Paper{
		print.NewPaper(650, 900),
		print.NewPaper(650, 1000),
	}
	pr := print.NewPaper(210, 297)

	var res []*print.Paper

	for _, m := range machines {
		for _, pl := range papers {
			for _, r := range pr.PackTo(m) {
				res = append(res, r.PackTo(pl)...)
			}
		}
	}

	fmt.Printf("Generated %d combinations\n", len(res))
}

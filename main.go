package main

import (
	"RollingHash/fileio"
	"RollingHash/hash"
	"fmt"
)

func main() {

	r, err := fileio.Open("test/text.txt")
	if err != nil {
		fmt.Errorf("ERROR: ", err)
	}

	hash.Split(r)
}

// ako chunka mi e po baluk ot 64 togava direkto savveni kontenta
//1.  idva mi 1 file - cheta go i mu nasmqtam hashovete po blokove i go slagam v vector
//2. idva mi vtori file - cheta go i mu nasmqtam hashovete po blokove i go slagam v vector
//3.(-- Tova ti e rolling hash) smqtash levenstein distance i operacii koito da postignesh minimum
//4. nakraq pravish diff na mestata kudeto si reshil da rpavish change
// ako byte-ovete sa ti poveche ot poluvinata golemina na bloka smeni change s remove/add

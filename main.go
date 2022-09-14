package main

import (
	"RollingHash/fileio"
	"RollingHash/hash"
	"fmt"
	"strings"
)

func main() {
	oFile := "test_files/old.txt"
	oldFileContent, err := fileio.Open(oFile)
	if err != nil {
		fmt.Errorf("ERROR: ", err)
	}
	oBlocks, err := hash.Split(oldFileContent)

	nFile := "test_files/new.txt"
	nFileContent, err := fileio.Open(nFile)
	if err != nil {
		fmt.Errorf("ERROR", err)
	}
	nBlocks, err := hash.Split(nFileContent)

	patternAsString := strings.Join(oBlocks, "")
	pattern := hash.RabinKarp(patternAsString)

	patternAsString := strings.Join(nFileContent[:], "")
	window := hash.RabinKarp(patternAsString)
	for i := 0; i < len(nBlocks); i++ {

	}

}

// ako chunka mi e po baluk ot 64 togava direkto savveni kontenta
//1.  idva mi 1 file - cheta go i mu nasmqtam hachetashovete po blokove i go slagam v vector
//2. idva mi vtori file -  go i mu nasmqtam hashovete po blokove i go slagam v vector
//3.(-- Tova ti e rolling hash) smqtash levenstein distance i operacii koito da postignesh minimum
//4. nakraq pravish diff na mestata kudeto si reshil da rpavish change
// ako byte-ovete sa ti poveche ot poluvinata golemina na bloka smeni change s remove/add

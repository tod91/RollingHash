package main

import (
	"RollingHash/helper"
	"fmt"
)

func main() {
	//oFile := "test_files/old.txt"
	//oldFileContent, err := fileio.Open(oFile)
	//if err != nil {
	//	fmt.Errorf("ERROR: ", err)
	//}
	//oBlocks, err := hash.Split(oldFileContent)     // string slice ot heshirani blokove
	//patternAsString := strings.Join(oBlocks, "")   // pravi ti go na 1 golqm string
	//patternHash := hash.RabinKarp(patternAsString) //izkarva ti hesha na tozi string
	//
	//nFile := "test_files/new.txt"
	//nFileContent, err := fileio.Open(nFile)
	//if err != nil {
	//	fmt.Errorf("ERROR", err)
	//}
	//nBlocks, err := hash.Split(nFileContent)                   // string slice ot heshirani blokove
	//windowAsString := strings.Join(nBlocks[:len(oBlocks)], "") // prozoreca ti go pravi na string
	//windowHash := hash.RabinKarp(windowAsString)               // izkarvati hesha na prozoreca
	//
	//delta := helper.NewDelta()
	//for i := 0; i < len(nBlocks); i++ {
	//	if patternHash == windowHash {
	//		windowAsString = strings.Join(nBlocks[i:len(oBlocks)], "")
	//		if helper.SanityCheck(patternAsString, windowAsString) {
	//			//delta.Add()
	//		}
	//	}
	//
	//	//windowHash = hash.SlideWindow(windowHash)
	//
	//	//strings.Join(nBlocks[:len(patternAsString)]
	//
	//}
	//
	//if len(delta.Changes) == 0 {
	//	fmt.Println("Files are identical")
	//} else {
	//	fmt.Println(delta.Changes)
	//}

	fmt.Println(helper.LevenshteinDistance("cereal", "saturdzz"))
}

// ako chunka mi e po baluk ot 64 togava direkto savveni kontenta
//1.  idva mi 1 file - cheta go i mu nasmqtam hachetashovete po blokove i go slagam v vector
//2. idva mi vtori file -  go i mu nasmqtam hashovete po blokove i go slagam v vector
//3.(-- Tova ti e rolling hash) smqtash levenstein distance i operacii koito da postignesh minimum
//4. nakraq pravish diff na mestata kudeto si reshil da rpavish change
// ako byte-ovete sa ti poveche ot poluvinata golemina na bloka smeni change s remove/add

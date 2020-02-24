
package main

import (
	"fmt"
)



type Auto struct {
	Marka		string
	GodVipusca		int
	ObyemBagajnica 	int
	ZapuskDvigatelya	bool
	PolojenieOcon		bool
	BagajnicFull		float32
}
	
func AutoN() Auto {
	nissan := Auto {
		Marka:				"Nissan",
		GodVipusca:			2020,
		ObyemBagajnica: 	80,
		ZapuskDvigatelya:	true,
		PolojenieOcon:		true,
		BagajnicFull:		0.5,
			}
			return nissan
	}

func AutoR() Auto {
		lada := Auto {
			Marka:				"LADA",
			GodVipusca:			2015,
			ObyemBagajnica: 	40,
			ZapuskDvigatelya:	true,
			PolojenieOcon:		true,
			BagajnicFull:		0.6,
			}
			return lada
	}


func checkNewer(nissan Auto, lada Auto) {
	if nissan.GodVipusca < lada.GodVipusca {
		fmt.Printf("Авто %v старше %v", nissan.Marka, lada.Marka)
	} else if nissan.GodVipusca > lada.GodVipusca {
		fmt.Printf("Авто %v младше %v", nissan.Marka, lada.Marka)
	} else {
		fmt.Println("Оба авто одинакового возраста")
	}
} 

func main() {
		Ni := AutoN()
		La := AutoR()
		checkNewer(Ni, La)
}
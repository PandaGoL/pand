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
	type Gruz struct {
		Marka		string
		GodVipusca		int
		ObyemKuzova 	int
		ZapuskDvigatelya	bool
		PolojenieOcon		bool
		BagajnicFull		float32

}
func AutoCar() {
	nissan := Auto{
		Marka:			"Nissan",
		GodVipusca:		2020,
		ObyemBagajnica: 	80,
		ZapuskDvigatelya:	true,
		PolojenieOcon:		true,
		BagajnicFull:		0.5,
			}
		honda := Auto{
		Marka:			"Honda",
		GodVipusca:		1978,
		ObyemBagajnica: 	40,
		ZapuskDvigatelya:	false,
		PolojenieOcon:		false,
		BagajnicFull:		0.9,
			}
	fmt.Println(nissan, honda)
	fmt.Printf("%+v", nissan)
}

func GrozAuto(){
		man := Gruz{
		Marka:				"MAN",
		GodVipusca:			2010,
		ObyemKuzova: 		500,
		ZapuskDvigatelya:	true,
		PolojenieOcon:		false,
		BagajnicFull:		0.1,
			}
		Kamaz := Gruz{
		Marka:				"Kamaz",
		GodVipusca:			1978,
		ObyemKuzova: 		600,
		ZapuskDvigatelya:	false,
		PolojenieOcon:		false,
		BagajnicFull:		1.0,
			}
	fmt.Println(man, Kamaz)
	fmt.Printf("%+v", Kamaz)
}



func main() {
		AutoCar()
		GrozAuto()
}
package receiver

import "fmt"

type Fatura struct {
	Isim   string
	Bahsis float32
	Items  map[string]float32
}

// type Yemek struct{
// 	Ismi string
// 	Birimfiyat int32
// 	Ulke

// }

// type Ulke struct{
// 	Ismi string
// }

func (f *Fatura) Hesap() string {
	fmt.Println(f.Isim)
	var toplamtutar float32 = 0
	text := "------Faturaniz---------\n"
	for index,value := range f.Items{
		text += fmt.Sprintf("%v ....$%v\n",index,value)
		toplamtutar += value
	}
	text += fmt.Sprintf("Bahsis ....$%v\n",f.Bahsis)
	text += fmt.Sprintf("-------------------------\ntoplam tutariniz : %v$\n",toplamtutar+f.Bahsis)
	return text
}
package receiver

import "fmt"

type Fatura struct {
	Isim   string
	Bahsis float32
	Items  map[string]float32
}

func (f *Fatura) Hesap() string {
	fmt.Println(f.Isim) //map elemani yazdirmak
	var toplamtutar float32 = 0
	text := "------Faturaniz---------\n"
	for index,value := range f.Items{
		text += fmt.Sprintf("%-10v:$%v\n",index,value)
		toplamtutar += value
	}
	text += fmt.Sprintf("Bahsis    :%v$\n",f.Bahsis)
	text += fmt.Sprintf("-------------------------\ntoplam tutariniz : %v$\n",toplamtutar+f.Bahsis)
	return text
}


func (f *Fatura) Update(name string, price float32){
	f.Items[name] = price //mape eleman eklemek
}



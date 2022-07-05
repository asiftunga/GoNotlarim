package romantoint

func RomanToInt(s string) int {
    var toplam int = 0
    var Sayilar = map[string]int{"I":1,"V":5,"X":10,"L":50,"C":100,"D":500,"M":1000}
	for i:=0; i<len(s); i++{
		switch string(s[i]) {
		case "I":
			if i != len(s)-1{
				if string(s[i+1]) == "V" || string(s[i+1]) == "X" {
					toplam -= Sayilar[string(s[i])]
					continue
				}else{
					toplam += Sayilar[string(s[i])]
					continue
				}
			}
			toplam += Sayilar[string(s[i])]
		case "X":
			if i != len(s)-1{
			if string(s[i+1]) == "L" || string(s[i+1]) == "C" {
				toplam -= Sayilar[string(s[i])]
				continue
			}else{
				toplam += Sayilar[string(s[i])]
				continue			
			}}
			toplam += Sayilar[string(s[i])]
		case "C":
			if i != len(s)-1{
			if string(s[i+1]) == "D" || string(s[i+1]) == "M" {
				toplam -= Sayilar[string(s[i])]
				continue
			}else{
				toplam += Sayilar[string(s[i])]	
				continue		
			}}
			toplam += Sayilar[string(s[i])]
		default:
			toplam += Sayilar[string(s[i])]
		}
	}
    return toplam
}
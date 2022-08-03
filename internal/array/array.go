package array
import "github.com/Pallinder/go-randomdata"


//diyelim ki 10 adet odamiz var ve her odanin 2 adet misafiri olabilsin
//boyle bir durumda 2 dimensional array kullanmak mantikli olur

func Generate() [10][2] string { //random bir sekilde [10][2] dizi olusturduk ve bunu asagidaki fonksiyonu kullanarak randomdatalar ile doldurduk 
	guestList := [10][2]string{} //bos bir string array tanimladik burda
	for index,_ := range guestList{
		guestList[index] = RoomGuests(index)
	}
	return guestList
}


func RoomGuests(roomId int) [2]string{
	guests := [2]string{} //burada da bos bir string array tanimladik
	guests[0] = randomdata.SillyName()
	guests[1] = randomdata.SillyName()
	return guests
}
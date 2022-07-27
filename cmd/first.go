/*
go dilinde variablelar nasil tanimlanir?
degisken  isimlerine identifier denir degerlerine ise expression denir

var degiskenIsmi degiskenturu = deger   seklinde degisken tanimlamasi yapilabilir

var sayi int32 = 84 seklinde bir tanimlama yapilabilir




var degisken1,degisken2,degisken3 degiskenlerinturu = deger1,deger2,deger3

var geek1, geek2, geek3 int = 232, 784, 854 boyle bir tanimlamada sadece ayni turde degiskenler tanimlanabilir

var geek4, geek5, geek6 = 100, "GFG", 7896.46 fakat tip (type) kismini (int) kaldirirsak istedigim sekilde degisken tanimlamasi yapabilirim

once degiskeni tanimlayip sonra deger atamasi gerceklestirebilirim. //@ onemli bir not olarak degiskeni tanimlayip deger atamasi gerceklestirmemis olsam bile go dilinde uninitialized variablelar yoktur bu nedenle bool lar false degeri sayisal degerler 0 ve string ifadeler "" seklinde ilk deger atamasi yapilir (initialize edilir)


    var geek1 int
    var geek2 string
    var geek3 float64
    var geek4 bool

seklindedir
=============kisa degisken tanimlamalari==================================
kisa degisken tanimlamalari su sekilde yapilir
degiskenismi := deger

!   kisa degisken tanimlamalari fonksiyonlarin disarisinda kullanilamazlar

============sabit degisken tanimlamak=====================================

const degiskenIsmi degiskenturu = deger

tipsiz bir const tanimlamasi yapilabilir
    const degiskenismi = deger  seklinde



degisken tanimlamalarinda su yontem de kullanilabilir

var(
     geek1 = 100
     geek2 = 200.57
     geek3 bool
     geek4 string = "GeeksforGeeks"
)
bu sekilde birden fazla degiskeni tek seferde tanimlamis oldum

While using type during declaration you are only allowed to declare multiple variables of the same type. But removing type during declarations you are allowed to declare multiple variables of different types.


==============================================================================
                    if else if else ve switch case yapisi
==============================================================================
if kosul{

}else if kosul{

}else{

}
! bu kullanim seklinden baska bir kullanim sekli yoktur. Kosullar parantezler icerisine alinamaz else ve else if ayri bir satirda yazilamaz

======================= switch case yapisi ===================================

switch ageJohn {
case 10:
   fmt.Println("John is 10 years old")
case 20:
   fmt.Println("John is 20 years old")
case 100:
   fmt.Println("John is 100 years old")
default:
   fmt.Println("John has neither 10,20 nor 100 years old")
}


switch ageSum := ageJohn + agePaul; ageSum {
case 10:
   fmt.Println("ageJohn + agePaul = 10")
case 20, 30, 40: //*\label{switchMulti}
   fmt.Println("ageJohn + agePaul = 20 or 30 or 40")
case 2 * agePaul:
   fmt.Println("ageJohn + agePaul = 2 times agePaul")
}

switch {
case agePaul > ageJohn:
   fmt.Println("agePaul > ageJohn")
case agePaul == ageJohn:
   fmt.Println("agePaul == ageJohn")
case agePaul < ageJohn:
   fmt.Println("agePaul < ageJohn")
}

switch {
case agePaul > ageJohn:
   ...
}

is equivalent to

switch true {
case agePaul > ageJohn:
    ...
}


==================================for kullanimi========================================
package main

import "fmt"

func main() {

    i := 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
        !   forun while seklinde kullanilmasi
    }

    for j := 7; j <= 9; j++ {
        fmt.Println(j)
    }

    for {
        fmt.Println("loop")
        break
    }

    for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
}
//ayni if yapisinda oldugu sekilde for kosullari da parantezler icerisine alinamaz. Alindigi takdirde hata mesaji ile karsilasilir

*/

// ==========================================================================================
// FONKSIYON KULLANIMI
// ==========================================================================================

/*
func computePrice(rate float32, nights int) (price float32) {
   price = rate * float32(nights)
   return
}
boyle bir kullanimda fonksiyonun return kisminda bir sey yazmaya gerek yoktur cunku fonksiyon taniminda geri donus degerine isim verdik, eger isim vermemis olsaydim mutlaka return float32 turunde geri donus degeri koymam gerekirdi



*/
//! goruldugu uzere go dilinde de ayni sekilde dongulere isim verilebilmekte.
/* package main

import "fmt"

func main() {

    for2: for i:=0;i<20;i++{
        fmt.Println("hello")
		for m:=0; m<10;m++{
            fmt.Println("world")
			if m==5{
				break for2
			}
		}
	}

} */
//=============================================================================================
//* _ kullanimi bu kismi es gec anlamina gelmektedir. Ornekte 3 adet return degeri var, diyelim ki benim sadece iclerinden bir tanesine ihtiyacim var bu durumda yapmam gereken _ kullanmaktir bu kisaca diger degiskenleri ignore et anlamina gelir ayni zamanda import ettigim bir seyi su an kullanmayacak isem bas kismina _ ekleyebilirim. Ayni sekilde degisken tanimlayip o degiskeni kullanmadigim durumlarda da bunu uygulayabilirim ..... asagilarda da aciklamistim aslinda ancak _ varsa bu ayni zamanda init fonksiyonu icerisinde onun initialized edilmesi anlamini tasir. Daha detayli bilgi icin asagidaki notlarimi okuyabilirim

/* package main

import (
	"fmt"
	new "newgroupproject/internal/new"   //yap dikkat edersem kullandigim dosya yoluna takma isim atamasi yapmisim. Bu ~new~ ismini kullanarak modulume direkt erisim saglayabilirim
    _ "newgroupproject/internal/email"

)

func main(){
	l1:=23
	new.Deneme() //yap burada bu erisimin bir ornegi gorulmekte
	fmt.Println("hellooo ",l1)
} */

//boyle bir kullanimda l1:=23 fonc disinda kullanildigi icin func icinden buna ulasamiyoruz ancak soyle bir durum var eger ben bu kullanim yerine su sekilde bir kullanim yaparsam -> var l1 = 23 or var l1 int = 23 seklinde bu durumda func icinden bu degiskene erisebilirim herhangi bir sikinti yasamam
//============================================================================
// package-imports/application-refactor/problem/main.go

/* package main

import (
    "fmt"
    "newgroupproject/internal/email"
    "newgroupproject/internal/invoice"
)
func main() {

    // first reservation
    customerName := "Doe"
    customerEmail := "john.doe@example.com"
    var nights uint = 12
    emailContents := email.GetEmailContents("M", customerName, nights) //@ GetEmailContents isimli fonksiyonun geri donus degeri Sprintf oldugundan dolayi burada bir degisken kullanrak geri donus degerini bir degiskene atamis bulunmaktayiz
    fmt.Println(emailContents)
    email.SendEmail(emailContents, customerEmail)
    invoice.CreateAndSaveInvoice(customerName, nights, 145.32)
} */
/*
yukaridaki ornege benzer sekilde eger import kisminda takma isim verirsem bu durumda verdigim takma isim uzerinden go modullerine ulasabilirim

soyle bir structure oldugunu varsayalim


├── vehicle
│   ├── go.mod
│   ├── main.go
│   └── car
│       └── car.go

bu vehicle bir go moduludur

Go inside the vehicle directory and run the following command to create a go module named vehicle.

! go mod init vehicle

The above command will create a file named go.mod. The following will be the contents of the file.

module vehicle

go 1.14

bu sekilde bir go module'u olusturmus oldum

vehicle\main.go su sekilde

package main

import (
	car1 "vehicle/car"
	car2 "vehicle/car"

	"fmt"
)

func main() {
	c1 := new(car1.Car) //! new keywordu oop bir dil olmayan go nun struct yapisi sayesinde sanki oop bir dilmis gibi davranmasini sagliyor. new keywordu aslinda objenin memory adresinin geri dondurulmesini sagliyor boylece yeni yaratilmis olan obje diger objeyi isaret edebiliyor. Aslinda gerceklesmis olan sey yeni yaratilan obje diger obje icin bir pointerdir.
	c1.Single(10)
	fmt.Println(c1.Price)

	c2 := new(car2.Car)
	c2.Double(10)
	fmt.Println(c2.Price)
}

@ goruldugu uzere burada takma isimler uzerinden paketlere ulasim saglanmis

vehicle\car\car.go dosyasi ise su sekilde

package car

type Car struct {
	Price int
}

func (obj *Car) Single(price int) int {
	obj.Price = price
	return obj.Price
}

func (obj *Car) Double(price int) int {
	obj.Price = price * 2
	return obj.Price
}

vehicle>go run main.go bu komutu calistirdigimiz zaman ise su output karsimiza cikmakta:
    10
    20
*/

//! ayni temada olan fonksiyonlari birbirleri ile birlestirip onlari birer paket haline getirmek iyi bir aliskanliktir

//* https://www.practical-go-lessons.com/chap-11-packages-and-imports#answers bu adreste go paketinin dogru bir sekilde github uzerinden nasil diger kisilerle paylasilacagi anlatilmis. (go.mod file buna gore yapiliyor)

//* https://www.practical-go-lessons.com/chap-11-packages-and-imports#answers bu adreste key takeaways kisminda paketlerin genel bir ozeti bulunmakta

//@ paket ayni directories icerisinde bulunan bir veya birden cok source files'a verilen isimdir. Source files : su anda icinde yazdigim first.go file'i source filedir. Source file icerisinde package blabla ardindan import edilen seyler ardindan degisken tanimlamalari ardindan ise fonksiyon tanimlamalarinin oldugu bir yerdir
//@ bir isim bir paketi tanimlar.
//@ buyuk harf ile baslayan bir isimlendirme exported diye isimlendirilir ve public anlami tasir
//! bir modul kokunde (root kisminda) bir go.mod file dosyasi bulunan dosya agacinda depolanan bir go paketleri toplulugudur
/*
.
├── cmd
│   └── first.go
├── go.mod
├── go.mod.save
└── internal
    ├── email
    │   └── email.go
    ├── invoice
    │   └── invoice.go
    └── new
        └── new.go

mesela boyle bir yapida bu yapinin tamami bir moduldur ancak icerisindeki email, invoice, new ler birer pakettir
*/

//bir paketi _ blank identifier ile cagirdiginda eger paketin icerisinde init fonksiyonu var ise init fonksiyonunu calistirmis oluruz

/*
   ------initializion sirasi--------
   oncelikle import ettigimiz paketler initialized edilir
   oncelikle paketimize import edilmis olan paketlerin degiskenleri (variablelari initialized edilir)
   ardindan import ettigimiz paketlerin init fonksiyonlari calisir
   sonra paketin kendisi initialized edilir
   paket icerisindeki degiskenler initialized edilir
   son olarak paketin icersindeki init fonksiyonu calistirilir
   onemli : init fonksiyonlari sira ile calisir yani ayni anda calismaz
*/

/*
notlarim: bir paketin oncelikle variablelari ardindan ise init fonksiyonu calisir
eger import edilmis bir paket varsa oncelikle onlarin variablelari sonrasinda ise onlarin init fonksiyonlari calisir
init fonksiyonlari zorunlu fonksiyonlar degillerdir ve birden cok init fonksiyonu bulunabilir
init fonksiyonlari variablelaardan sonra initialize edilirler
init fonksiyonlari sira ile calisirlar
bir variable initialize edilecekse ve birden cok variable varsa oncelikle bagimsiz olan (digerlerine bagimli olmayan variable) initialize edilir ardindan en az bagimli olan ve bu sekilde devam eder.
*/

/*
 var nameDisplayer func(name, firstname string) string bu sekilde bir kullanim mevcut bu kullanimin amaci ise sudur fonksiyonu henuz tamamlamadan tanimlama kismidir


===============================================================================

type Country struct {
    Name        string
    CapitalCity string
}
struct yapisi bu sekilde
kullanimi ise bu sekilde

france := Country{
    Name:        "France",
    CapitalCity: "Paris",
}

usa := Country{
    Name: "United Sates of America",
}

bu kisimda bos biraktigimiz capitalcity kismi bos kalir. (bos kalirdan kastim initialize edilir ancak baslangic degeri atanmaz default deger atanir) String bir ifade icin bu "" seklindedir



empty := Country{}

ayni sekilde bos bir struct objesi de olusturabiliriz, bu structin tum field kisimlari default sekilde initialize edilirler



japan := Country{
    "Tokyo",
    "Japan",
}
su sekilde bir kullanim yapilabilir ancak bu sekilde bir kullanim yapildigi zaman herhangi bir field alani bos gecilemez ve bir kural daha var
bu sekilde bir kullanim yapiliyorsa iki yontemi birbiri ile birlestiremezsin. Iki yontemden birisini kullanman gerekir*/

/*

 usa := Country{
    Name: "United Sates of America",
}
usa.CapitalCity = "Washington DC"

erismek icin bu sekilde bir kullanim vardir

//@ su sekilde de bir kullanim yapilabilir

usa := Country {}
usa.Name = "United States of America"
usa.CapitalCity = "Washington DC"




if usa.Name == "France" {
  fmt.Println("we have an error !")
}

bu sekilde bir kullanim da yapilabilir

*/

// bir structurein icinde baska bir structure field olarak kullanilabilir

/*
type Hotel struct {
    Name     string
    Capacity uint8
    Rooms    uint8
    Smoking  bool
    Country
}

type Country struct {
    Name        string
    CapitalCity string
}

*/

//------Ornek bir kullanim sekli

/* package main

import "fmt"

type Hotel struct {
    Name string
    Country
}

type Country struct {
    Name        string
    CapitalCity string
}

func main() {
    hotel := Hotel{
        Name:    "Hotel super luxe",
        Country: Country{
            "France",
            "blabladenenbiryer",
        },
    }
    fmt.Println(hotel.Name)
    fmt.Println(hotel.Country.Name)
    fmt.Println(hotel.Country.CapitalCity)

    /*  ==ciktisi==
    Hotel super luxe
    France
    blabladenenbiryer
    /*
}
*/

//@ bu sekilde de bir kullanim mevcuttur
/* package main

import "fmt"

//degiskenler main disinda tanimlanirlar

type Hotel struct {
    Name string
    Country
}

type Country struct {
    Name        string
    CapitalCity string
}

func main(){
    hotel := Hotel{}
    hotel.Name = "ATM otelleri"
    hotel.Country.Name = "Turkey"
    hotel.Country.CapitalCity = "Antalya"
    fmt.Println(hotel.Name)
    fmt.Println(hotel.Country)
}
===========output=============
ATM otelleri
{Turkey Antalya}
*/

// methodlar receiveri olan fonksiyonlara verilen isimlerdir

// package main

// import "fmt"

// func main() {
/*
   ! var DegiskenIsmi *DegiskenTuru = &AdresiGosterilecekDegisken
   ! var po *int = &i seklinde pointer tanimlamasi yapabilecegim gibi
   ! po := &i seklinde de pointer tanimlamasi yapabilirim
*/
/*

   i:=1
   po := &i        //@ bunun ile var po *int = &i ayni anlama geliyor karistirmamaliyim :)
   po2 := &po
   fmt.Println("i nin degeri: ",i)
   fmt.Println("i nin adresi: ",&i)
   fmt.Println("pointer degeri: ",po)
   fmt.Println("Pointerin adresi: ", &po)
   fmt.Println("Pointerin isaret ettigi adresteki deger: ", *po)
   fmt.Println("Pointerin adresini tutan pointerin degeri: ",po2)
   fmt.Println("Pointerin adresini tutan pointerin adresi: ",&po2)
   fmt.Println("Pointerin adresini tutan pointerin isaret ettigi adresteki degeri: ",*po2)
   fmt.Println("Pointerin adresini tutan pointerin isaret ettigi adresteki pointerin isaret ettigi adresteki degeri: ",**po2)
   ========cikti============
   i nin degeri:  1
   i nin adresi:  0xc00008a040
   pointer degeri:  0xc00008a040
   Pointerin adresi:  0xc00009e038
   Pointerin isaret ettigi adresteki deger:  1
   Pointerin adresini tutan pointerin degeri:  0xc00009e038
   Pointerin adresini tutan pointerin adresi:  0xc00009e040
   Pointerin adresini tutan pointerin isaret ettigi adresteki deger:  0xc00008a040
   Pointerin adresini tutan pointerin isaret ettigi adresteki pointerin isaret ettigi adresteki deger:  1

*/

/*     i := 1

fmt.Println("initial:", i)

zeroval(i)
fmt.Println("zeroval:", i)

zeroptr(&i)
fmt.Println("zeroptr:", i) */

//}

/* func zeroval(ival int) {
    ival = 0
}

func zeroptr(iptr *int) {
    *iptr = 0
} */

/* package main

import (
	"fmt"
	_ "fmt"
	"newgroupproject/internal/cart"
	_ "time"
)

func main(){
    newcart := cart.Cart{}
    newcart.Lock()
    totalPrice := newcart.TotalPrice()
    err := newcart.Lock()
}
! receiverli metotlari gormek icin internal/cart/cart.go kismina git
bu kodun soyle bir mantigi var aslinda. internal klasoru icersinde cart isimli bir klasorum var o klasorun icerisinde cart.go isimli bir paketim (source code'a paket ismi veriliyordu) var. Bu paketi import kisminda import etmisim. Daha sonra bu paketimin icerisinde Cart isimli bir struct'im ve iki adet metodum var. Metodlarimin receiveleri Cart tipinde oldugundan dolayi cart tipinde bir sey olmadan kullanamam bundan dolayi metotlarimi kullanabilmek icin newcart isimli Cart tipindeki cart.Cart{} olusturdum, daha sonra bunu kullanarak metotlarima erisebildim
=====================================peki value receiver mi yoksa pointer receiver mi kullanmaliyim?=======================
eger value receiver olusturursam gonderdigim degisken method icerisinde kopyasi (internal) olusturulur ve o sekilde kullanilir  yani gonderdigim degisken degistirilmez.
! ayrica ustteki kullanim cok az da olsa memory kisminda fazlalik olusturur. Bu cogu zaman goz ardi edilebilir ancak memory onemli olan kullanimlarda dikkat edilmelidir
ancak eger pointer tipinde bir degisken yollarsam yolladigim degiskenin adresi gideceginden dolayi yollayacagim degisken degistirilir
@receiver kullanirken genel kullanim type in first letteridir. c *Cart gibi
*/

// methods/example-project/main.go
/*package main

import (
	"fmt"
	_"github.com/Rhymond/go-money"
	"newgroupproject/internal/cart"
	_"newgroupproject/internal/product"
)

func main() {

    var newcart = cart.Cart{}
    total,error := newcart.TotalPrice()
    fmt.Printf("toplam tutar : %d hata kodu : %d",total,error)

	/*

	   totalPrice, err := newCart.TotalPrice()
	   if err != nil {
	       fmt.Printf("impossible to compute price of the cart: %s", err)
	       return
	   }
	   fmt.Println("Total Price", totalPrice.Display())

	   err = newCart.Lock()
	   if err != nil {
	       fmt.Printf("impossible to lock the cart: %s", err)
	       return
	   } */

//?================================================================================================================
//? bu kod blogu sadece receiver fonksiyonlari ve struct yapisini daha iyi anlamak icin var
//?================================================================================================================

/* package main

import (
	"fmt"
	"newgroupproject/internal/receiver"
    _"strconv"
)

func main(){
    //bunu pointer seklinde kullanmadigimda da oluyor peki pointerin buradaki amaci nedir amk
    //? burdaki b yi su sekilde kullandigimda neden bir fark olusmuyor bunu sonradan tekrar incelemem gerekli cok ONEMLI (b := receiver.Fatura)
    var b *receiver.Fatura = &receiver.Fatura{
        Isim: "Market Faturasi",
        Bahsis: 40.00,
        Items: map[string]float32{"patates": 5.00, "elma":4.00, "biber":10.00, "salatalik":12.00},
    }
    fmt.Printf("%+v\n", receiver.Fatura{})
    fmt.Println(b)
    b.Update("patlican",20.00)
    b.Update("domates",35.00)
    b.Update("lahana",17.00)
    fmt.Println(b.Hesap())
}
*/

//?===========================================================================================================
//leetcode sorusu : roman to int
/* package main

import (
    "fmt"
    "newgroupproject/internal/romantoint"
)

func main(){
    fmt.Println(romantoint.RomanToInt("MCMXCIV"))
} */

//============================================================================================================
//*                            Pointerlar icin notlarim
/*
type User struct {
    ID string
    Username string
}

! bu sekilde bir structim oldugu zaman *User demek User turune ait tum degiskenleri belirtir anlamindadir

pointer tanimlamasi su sekilde yapilir ->

var p *int
var answer int = 42
p = &answer

veya ayni sekilde
var answer int = 42
p := &answer

seklinde pointer tanimlamasi yapabilirim

simdi pointerlar aslinda adres tutarlar peki ben pointerin isaret ettigi o adresteki degeri kullanmak istersem ne yapmaliyim? * dereference operatorunu kullanmaliyim

bunu su sekilde yapabilirim. Ustte tanimlamis oldugum bir struct var
bu struct elemanlarina erismek icin sunlari yapabilirim

cart := Cart{
    ID: "12345",
    Paid: true,
}

cartPtr := &cart //! cartPtr su anda sadece 0x385v854 seklinde bir adres tutuyor (yani deger degil degerlerin saklandigi adresi tutuyor (isaret ediyor))

cartDeref := *cartPtr  //! iste bu sekilde su an pointerin isaret ettigi yerdeki value lara ulasmis oldumm

@ & -> adres bilgisini al (reference)
@ * -> adres bilgisine ulas (takip et) (dereference)
* operatoru ile *card seklindeki kullanim birbirinin aynisi degildir. *card bir isaretci tipi olabilecegi gibi ayni zamanda referansi kaldirilmis bir isaretci degiskenini de gosteriyor olabilir. Kullanimini yakindan inceleyerek hangisi olduguna karar vermem gerekir


!     var myPointerVar *int
//@============================
!    fmt.Println(*myPointerVar)  mesela soyle bir kullanim yanlistir bunun nedenini detayli aciklamak istiyorum

@ oncelikle tipi(adresini gosterecegi) *int olan bir pointer degiskeni tanimladik. Daha sonra bunu dereference etmeye calistik ancak oncesinde referans etmemistik ki.. yani bu sekilde sanki gostermedigi bir yer oldugu halde o yere gitmeye calisiyoruz aslinda o yer nil dir. Yani yoktur bu nedenle program hata verir. Bu programin calismasinin tek yolu degisken tanimlamasindan (kirmizi yerden) sonra reference atamasi yapilmasidir. mypointer = &blabla seklinde. Sonra buradaki kullanimi yapabilirim


map ve slice lar reference tipleridir. Yani kendi ic yapilarina reference lilardir. Bu soyle bir kullanim saglar, bir method parametre olarak bir maps aliyor ise ve bu parametre pointer seklinde olmasa bile yapilan degisiklikler direkt olarak mapse yansir.
*/

// package main

// import (
//     "fmt"
//     _"strings"

// )

//slicelar ayni turdeki elementlerin birer koleksiyonudur.

/* func main(){
    EUcountries := []string{"Austria", "Belgium", "Bulgaria"} //slice tanimlamasi yaptik
    addCountries(EUcountries)
    fmt.Println(EUcountries)
    upper(EUcountries)
    fmt.Println(EUcountries)
} */

//slice eleman eklemek icin

/* func addCountries(countries []string) {
    countries = append(countries, []string{"Croatia", "Republic of Cyprus", "Czech Republic", "Denmark", "Estonia", "Finland", "France", "Germany", "Greece", "Hungary", "Ireland", "Italy", "Latvia", "Lithuania", "Luxembourg", "Malta", "Netherlands", "Poland", "Portugal", "Romania", "Slovakia", "Slovenia", "Spain", "Sweden"}...)
} */

//bu sekilde bu fonksiyonu main icerisinde calistirdigimizda [Austria Belgium Bulgaria] sonucuna ulasiyoruz. Bunun sebebi parametre gecisi sirasinda internal bir kopya olusturulmasidir. Eger addcountries fonksiyonu icerisinde bu print isini gerceklestirmis olsaydik ekledigimiz tum countryleri gorurduk ([Austria Belgium Bulgaria] bunlar dahil)

/* func upper(countries []string) {
    for k, _ := range countries {
        countries[k] = strings.ToUpper(countries[k])
    }
} */

//peki bu sekilde bir kullanim yapsak ve main icersinde bu fonksiyonu calistirsam ne olurdu? Daha deminki gibi bir sonuc bekliyorsam bu tamamen yanlis cevap su olurdu [AUSTRIA BELGIUM BULGARIA]. Peki ama daha demin degismeyen slice ne oldu da su anda degisti?  adim adim aciklayayim
/*
ilk olarak upper fonksiyonu eucontries sliceinin bir local kopyasini olusturdu
bu fonksiyonun icerisinde slice elementinin valuelarini degistirdik
bu kopya hala temel olan slice i referans ediyor ancak soyle bir durum var //! sadece ama sadece var olan degerleri degistririsek bu olur demek istedigim slice icerisinde var olan elemanlar degistirilebilir bunun nedeni ise:
@ aklimda daha kolay kalmasi icin sanki icerisindeki her bir value orijinal olan slice in valuelarina referans gosteriyormuslar gibi dusunebilirim
//! ancak yeni eleman eklersem bu sonucumuza yansimaz.
*/

/* func main() {
    EUcountries := []string{"Austria", "Belgium", "Bulgaria"}
    addCountries2(&EUcountries) //slice adresini yolladik
    fmt.Println(EUcountries)
}

func addCountries2(countriesPointer *[]string) { //@ parametre olarak slice pointeri veriyoruz
    *countriesPointer = append(*countriesPointer, []string{"Croatia", "Republic of Cyprus", "Czech Republic", "Denmark", "Estonia", "Finland", "France", "Germany", "Greece", "Hungary", "Ireland", "Italy", "Latvia", "Lithuania", "Luxembourg", "Malta", "Netherlands", "Poland", "Portugal", "Romania", "Slovakia", "Slovenia", "Spain", "Sweden"}...)
} */

// kisaca addcountries2 fonksiyonuna bir adres bilgisi yollandi ve fonksiyon kendi icerisinde olusturdugu degisken ile bu adres bilgisinin degerine  * operatoru ile ulasti

//simdi bir kullanim gosterecegim. dereference operatorunu kullanmadan sadece adres bilgisi ile struct kullanabiliriz
/* func main(){

    type Cart struct {
        ID   string
        Paid bool
    }

    cart := Cart{
        ID:          "115552221",
        Paid:   true,
    }
    cartPtr := &cart
    cartPtr.ID = "tunga"
    cartPtr.Paid = false

    fmt.Println(cart)

    //{tunga false} seklinde bir ciktisi olur
}
*/

//burdaki kullanim ile su kullanim ayni seydir

/* func main(){

    type Cart struct {
        ID   string
        Paid bool
    }

    cart := Cart{
        ID:          "115552221",
        Paid:   true,
    }
    cartPtr := &cart
    (*cartPtr).ID = "asif tunga"
    (*cartPtr).Paid = false

    fmt.Println(cart)

    //{asif tunga false} seklinde bir ciktisi olur
} */

/*
type Cat struct{
    Color string
    Age int16
    Name string
}


func (cat *Cat) Rename(newname string){
    cat.Name = newname
}

func (cat Cat) Rename2(newname string){
    cat.Name = newname
}

func Rename3(adress *Cat,newname string){
    (*adress).Name = newname
}

func main(){
    cat := Cat{Color: "blue", Age: 8, Name: "Milow"}
    cat.Rename("Bob")
    fmt.Println(cat.Name)
    //! Bob

    cat.Rename2("Ben")
    fmt.Println(cat.Name)
    //! Bob yazdirir cunku fonksiyon icerisindeki local degiskeni degistirdik

    adres := &cat
    Rename3(adres,"hasan")
    fmt.Println(cat.Name)
    //! hasan yazdirir :)
}
*/

//@=========================ne zaman pointer receiver ne zaman value receiver kullanmam gerekli====================================================
/*
                            //yap Pointer receiver kullan
struct yapim cok buyuk ve agir ise
receiveri degistirmek istiyorsam (ornegin struct alaninda bir degiskenin isim alanini degistirmeye calisiyorsam)
synchronization primitive (like sync.Mutex) gibi bir sey kullaniyorsak

                            //yap Value receiver kullan
struct yapim small ise
receiveri degistirmeye gerek yok ise
eger receiver map, func, chan, slice, string, interface ise bunlarla value receiver kullanmam gerekir cunku bu yapilar yapisal olarak hali hazirda pointerlar ile calisirlar
*/
//@=========================================interface kavrami==================================================

//arayuz dedigimiz sey aslinda bir dizi ortak davranisi tanimlayan bir kontraktir.
//herhangi bir implementation olmadan (ornek olarak metotlarin davranislarini tanimlar)
//kisaca interface herhangi bir implementation olmadan methodlarin nasil bir tipte oldugunu tanimlar

//ornek bir interface implementationu bu sekilde gerceklesir

/* package main

import "fmt"

type Human struct{
    Firstname string
    Lastname string
    Age int
    Country string
}

type Cat struct{
    Name string
}

type Dog struct{
    Name string
}

func (c Cat) ReceiverAffection(from Human){
    fmt.Printf("the cat name %s receive affection from human who name is %s\n",c.Name,from.Firstname)
}

func (c Cat) GiveAffection(to Human){
    fmt.Printf("the cat name %s give affection to human who name is %s\n",c.Name,to.Firstname)
}

func (d Dog) ReceiverAffection(from Human){
    fmt.Printf("the dog name %s receive affection form human who name is %s\n",d.Name,from.Firstname)
}


func (d Dog) GiveAffection(to Human){
    fmt.Printf("the dog name %s give affection to human who name is %s \n",d.Name,to.Firstname)

}

type DomasticAnimal interface{
    ReceiverAffection(from Human)
    GiveAffection(to Human)
}

func Pet(animal DomasticAnimal, human Human){
    animal.GiveAffection(human)
    animal.ReceiverAffection(human)
}


func main(){
    Pet(Dog{"MISHA"},Human{"asif","tunga",23,"turkiye"})
    Pet(Cat{"TOM"},Human{"hediye","ozgul",20,"bulgaristan"})
    }  */

// ?================================ anlamama yardimci olacak baska bir ornek olabilir ===================================================

// type Article struct {
// 	Title string
// 	Author string
// }

// type Book struct{
//     Title string
//     Author string
//     Pages int
// }

// func (b Book) String() string{
//     return fmt.Sprintf("The %q book was written by %s",b.Title,b.Author)
// }

// func (a Article) String() string {
// 	return fmt.Sprintf("The %q article was written by %s.", a.Title, a.Author)
// }

// type Stringer interface{
//     String() string //@ burada yaptigim tek sey Stringer isminde String() isimli string geri donuslu bir metot koydum. Interface tanimlamasi bu kadar
// }

// func main() {
//     a := Article{
//         Title: "Understanding Interfaces in Go",
//         Author: "Sammy Shark",
//     }
//     b:=Article{"Umutlu Bir Bilgisayar Muhendisinin Yasami","Asif Tunga Mubarek"} // ayni zamanda bu sekilde de kullaniliyor (daha fazla bilgi icin yukaridaki struct yapisina bakabilirim)
//     c:=Book{"Yasamim","Asif Tunga Mubarek",23}
//     Print(a)
//     Print(b)
//     Print(c)
// }

// func Print(s Stringer) {
//     fmt.Println(s.String())
// }
// cok onemli bir not : //! bir method fonksiyondan farkli olarak sadece tanimlandigi turun orneginden cagrilabilir
//A method is a special function that is scoped to a specific type in Go. Unlike a function, a method can only be called from the instance of the type it was defined on.
//simdi bu yapi beni kod tekrarindan kurtardi. Bunu bir ornekle gostermek istiyorum
/*
diyelimki elimde iki tane daha struct daha var

type Yazar struct{
    Name string
    Age int
    Country string
}

type Basimevi struct{
    Locasyon string
    Country string
}

func (y Yazar) String() string {
	return fmt.Sprintf("The %q article was written by %s.", y.Name, y.Country)
}

func (b Basimevi) String() string {
	return fmt.Sprintf("The %q article was written by %s.", b.Locasyon, b.Country)
}

//@ asagidaki print fonksiyonunu da degistirmemis oldugumu varsayalim (bundan kastim aslinda su: degistirmemisim derken interface ile kullanmamisim anlaminda soyleyebilirim)

func Print(a Article){
    fmt.Println(a.String())
}

func Print(b Book){
    fmt.Println(b.String())
}

func Print(y Yazar){
    fmt.Println(y.String())
}

func Print(b Basimevi){
    fmt.Println(b.String())
}

*/
//@================================================== peki bunu interface ile yapmis olsaydim nasil mi olurdu? ========================================= neler neler

// package main

// import "fmt"

// type Yazar struct{
//     Name string
//     Age int
//     Country string
// }

// type Basimevi struct{
//     Locasion string
//     Country string
// }

// type Article struct {
// 	Title string
// 	Author string
// }

// type Book struct{
//     Title string
//     Author string
//     Pages int
// }

// func (b Book) String() string{
//     return fmt.Sprintf("The %q book was written by %s",b.Title,b.Author)
// }

// func (a Article) String() string {
// 	return fmt.Sprintf("The %q article was written by %s.", a.Title, a.Author)
// }

// func (y Yazar) String() string {
// 	return fmt.Sprintf("The %q sehrinde yasayan kisi: %s.", y.Country, y.Name)
// }

// func (b Basimevi) String() string {
// 	return fmt.Sprintf("The %q lokasyonu su ulkede bulunur: %s.", b.Locasion, b.Country)
// }

// func main(){
//    a:=Article{"baslik","asif tunga mubarek"}
//    b:=Book{"kitap basligi","asif tunga mubarek",300}
//    c:=Yazar{"asif tunga mubarek",26,"antakya"}
//    d:=Basimevi{"mersin","turkey"}
//    Print(a)
//    Print(b)
//    Print(c)
//    Print(d)
// }

// type Stringer interface{
//     String() string
// }

// func Print(s Stringer) {
//     fmt.Println(s.String())
// }
/*
KONU ILE ALAKASIZ VIM NOTU : eger bir sonraki karaktere gitmek istiyor isem yapmam gereken tek sey insert moda a ile gecis yapmak. veya bunun yerine direkt olarak cumlenin sonuna gitmek istiyor isem A seklinde kullanabilirim ama bu direkt insert moda sokar haberim olsun
$ seklinde de kullanabilirim bu insert moda sokmaz direkt cumlenin sonuna giderim bu sekilde
*/

//@ interfaceleri anlamak icin kendime alacagim not kisimlari buralar
//? burada alinan notlar sadece interface konusunu daha iyi anlamak icindir. Her bir not acik bir sekilde yazilacaktir.
// burada anlatilan kodlar https://go.kaanksc.com/boeluem-3/arayuez-interface adresinden alinmistir. Detayli bilgi icin bu siteyi inceleyebilirim

/* package main

import "fmt"

type domatesSalçası struct {
    marka string
}

func (s domatesSalçası) Ye() {
    fmt.Sprintf("Domates salçası yenildi")
}

type BiberSalcasi struct {
    marka string
}

func (s BiberSalcasi) Ye() (dondur string) {
	fmt.Println("Biber salçası yenildi")
    return
}


type Salca interface{
    Ye() string    //bu kisim aslinda Ye() isimli bir fonksiyonun string bir ifade dondurdugunu anlatir
} */

//* deneme amacli alan ------------------------------------------------------------------------------------

// func Besle(salcam Salca){
//     salcam.Ye()
// }

// func main(){
//     Besle(BiberSalcasi{"tamek"}) //burada biber salcasi kismini parametre seklinde gonderdik. ASlinda ust kisimdaki besle fonksiyonunda interface i parametre olarak kullanmamiz sayesinde interface uzerinden fonksiyona erisiyorum gibi dusunebilirim
// }

//*------------------------------------------------- bu kisim ise ust kisim yerine yapilabilir--------------------------------------------------------
/* func main(){
    biber := BiberSalcasi{"tamek"}
    var salcam Salca = &biber
    fmt.Print(salcam.Ye())
} */

//@ ustteki ile bu ayni sonucu verir. Burada yaptigim sey ise aslinda deneme amacli olarak yazdigim kisimdan cok da farkli degil. Orda bir fonksiyon olusturdum ve o fonksiyon vasitasi ile interface yapisini parametre olarak gonderip interface yapisi uzerinden kontratta yani interface uzerinde belirlenmis olan fonksiyona interface uzerinden ulasmis oldum. (struct yapimi interface uzerinden gonderdim)
//@bu kisimda yaptigim sey ise oncelikle bir struct yapimi olustrudum ardindan interface type li bir variable a bu structimin tutuldugu adresi yolladim bir nevi interface imi pointer gibi kullanmis oldum ardindan bu interface variable im sayesinde, interface uzerinden gerekli fonksiyonuma ulasmis oldum

//bak-> fonksiyonlarin struct yapilari icerisinde nasil kullanildiklarini ogrenecegim

//! interfacelerin kullanim alanlarina ornek olarak -> bir suru database mevcut (postrgresql, mysql) ve bunlarin hepsinin Open metotlari ayri ayri. Bu durumda bu databaselere baglanmak icin ayni kodu initialize(baslatma) yapmayacagiz
//@ bir interface tanimlayarak bircok implemantationun kullanabilecegi bir kontrat hazirlamis oluruz aslinda
/*
type Driver interface {
    Open(name string) (Conn, error)
}
*/

/*
interface icerisine interface implement edebilirim ancak bunu yapmanin elbette dezavantaji vardir.
simdi bunun bir ornegini denemek istiyorum
*/
//bak-> alt kisimda bak dedigim yer burasi

// package main

// import "fmt"

// type BiberSalcasi struct{
//     marka string
// }

// func (b BiberSalcasi) Salca(){
//     fmt.Println("salca yapimi baslamis bulunmakta")
// }

// func (b BiberSalcasi) Ye(){
//     fmt.Println("salca yenildi")
// }

// type Yenier interface{
//     Ye()
// }

// type Salcaer interface{
//     Salca()
//     Yenier
// }

// func Olustur(s Salcaer){
//     s.Salca()
//     s.Ye()
// }

// func main(){
//     Olustur(BiberSalcasi{"tunga salcalari"})
// }
// bu sekilde gerceklestiremedim diger yol ile deneyecegim su anda
// tahmin ettigim gibi bu yol ile yapilabildi

//bak-> ustte yaptigim yol ile neden olmadigini mutlaka ogrenmeliyim
/* package main

import "fmt"

type BiberSalcasi struct{
    marka string
}

func (b BiberSalcasi) Salca(){
    fmt.Println("salca yapimi baslamis bulunmakta")
}

func (b BiberSalcasi) Ye(){
    fmt.Println("salca yenildi")
}

type Yenier interface{
    Ye()
}

type Salcaer interface{
    Salca()
    Yenier
}

func main(){
    b := BiberSalcasi{"tamek"}
    var s Salcaer = b
    s.Salca()
    s.Ye()
} */

//STANDART LIBRARYDE BULUNAN BIRCOK INTERFACE VARDIR BUNLARDAN EN UNLULERI ISE SUNLARDIR : Error interface, Stringer interface, Interface interface bu interfaceleri daha detayli incelemek istersem ; https://www.practical-go-lessons.com/chap-16-interfaces adresinden gorebilirim

//Go da bir diger interface tipi ise bos interfacedir. Bu yazabilecegim en kucuk ve en kisa interfacedir. Bu su sekilde yazilir interface{}
//peki ben bunu nerede kullanacagim? tanimlama olarak bos bir interface deger her tipte degiskeni tutabilir. Eger her tipi kabul eden bir metot yazmak istersem interface kullanabilirim

// func (l *Logger) Fatal(v...interface{}) {} bu sekilde bir kullanim butun tipleri kabul eder
//bos bir interface' i parametre olarak alan fonksiyonlarin genel olarak aldiklari bu interface in efektif tipini bilmeleri gerekmektedir. Bunu yapmak icin ise fonksiyon type switch yapisini kullanabilir. Bu normal switchte oldugu gibi valuelari degil typelari karsilastirir

/*
 printany prints an argument passed to panic.
 If panic is called with a value that has a String or Error method,
 it has already been converted into a string by preprintpanics.

func printany(i interface{}) {
    switch v := i.(type) {
    case nil:
        print("nil")
    case bool:
        print(v)
    case int:
        print(v)
    case int8:
        print(v)
    case int16:
        print(v)
    case int32:
        print(v)
    case int64:
        print(v)
    case uint:
        print(v)
    case uint8:
        print(v)
    case uint16:
        print(v)
    case uint32:
        print(v)
    case uint64:
        print(v)
    case uintptr:
        print(v)
    case float32:
        print(v)
    case float64:
        print(v)
    case complex64:
        print(v)
    case complex128:
        print(v)
    case string:
        print(v)
    default:
        printanycustomtype(i)
    }
}
*/

//peki hangi durumlarda bos interface kullanmaliyim? Hatta daha iyi bir soru kullanmali miyim? Bu soruyu soyle yanitlayabilirim kullanirken cok dikkatli olunmasi gerekli. Eger baska bir secenegim yok ise empty interface kullanmaliyim. Empty interface benim fonksiyon veya methodlarimi kullanacak kisiye hicbir ipucu vermez ayrica bir fonksiyonda empty interface kabul edersem kodum daha karmasik bir hale gelir bunun nedeni bu fonksiyonun oncesinde interface tipini check etmesidir.

/*
func (c Cart) ApplyCoupon(coupon Coupon) error  {
    //...
}

func (c Cart) ApplyCoupon2(coupon interface{}) (interface{},interface{}) {
    //...
}

hangisini tercih ederdim? elbette en usttekini.. Goruldugu uzere ApplyCoupon fonksiyonu strictly hangi parametre kullanilacagini ve geri donus degerini belirtir.
*/

/* package main

import (
    "newgroupproject/internal/cart"
    "errors"
)

type Al interface{
    GetById() (error)
    Put()
}

type CartStorePostgres struct{}

func (c *CartStorePostgres) GetById(ID string) (*cart.Cart, error) {
    // implement me
    return
}

func (c *CartStorePostgres) Put(cart *cart.Cart) (*cart.Cart, error) {
    // implement me
    return
}


func main(){

} */



/* package main

import (
	"fmt"
)


type Database interface{
    Baglan(d *Db)
}


type Db struct{
    dbisim string
    connectionstring string
}

func (d Db) Baglan(){
    fmt.Println(d.connectionstring +" db baglanma islemi gerceklestirildi")
}


func Bagla(Dber Database){
    deneme := Dber.Baglan()
    deneme.Baglan(deneme)
}

func main(){
    a:=&Db{"tamek","pstgresql"}
    Bagla(a)

}

*/

//! kendime not : interfaceler icin pointer kullanmaya gerek yoktur. Program ilk calistigi sirada zaten kendisi initialize edildigi icin tekrardan bir isaretci kullanip yeniden ram kullanmak gereksizdir.


/*  //@ API nedir sorusunun cevabi niteliginde
api lari aslinda birer contract gibi dusunebilirim. Tipki ustteki interfaceler gibi bir yapisi var. Bir api iki farkli seyin bir arada calismasina ve iletisim kurmalarina (communication) izin verir (genellikle json formatinda dosya yollayarak). API kisaca sunlardan olusur  Bir yazılım parçası ile etkileşime maruz kalan bir dizi yapıdır (sabitler, değişkenler, işlevler ...).

Bu tanımla, go paketi fmt'nin go programcısına onunla etkileşim kurması için bir API sunduğunu söyleyebiliriz. API'si, örneğin Println gibi kullanabileceğiniz bir dizi işlevi temsil eder. Ayrıca paketin dışa aktarılan tanımlayıcılarını da (sabitler, değişkenler, türler) kapsar.

Go Modülleri, modülün oluştuğu paket(ler)in dışa aktarılan tüm tanımlayıcılarından oluşan bir API'yi kullanıma sunar.
*/


/* 
A version number must have the following format : X.Y.Z
Where :

the major version

the minor version

the patch version (bug fix)

How does it work? X, Y, and Z are positive numbers (without leading zeroes). Those numbers are incremented following a specific norm.

When you create new features that breaks the existing API of your software, you increment the major version number.

old version : 1.0.0 / new version : 2.0.0
When you create new features or make performance improvement that do not break the existing API you increment the minor version number.

old version : 1.0.0 / new version 1.1.0
When you fix a bug in your code, you just increment the patch version

old version : 1.0.0 / new version 1.0.1
When you create a major version, you set to zero the minor and the patch version number. When you release a new feature, you set to zero the patch version.

eger programlamanin baslarinda versiyonlar yaratiyor isem major versiyonuna 0 demeliyim bu sayede diger developerlar benim kodlarimi kendi programlarina dahil ettiklerinde stabil olmadigini ve api in degisebilecegini anlayabileceklerdir
*/


//=== alttaki paket icin ve versiyonlama ile ilgili biligileri almak icin time paketine bakabilirim
/* 
package main

import (
	"newgroupproject/internal/time"
	"fmt"
    )

func main(){
	fmt.Println(time.WhatTimeIsIt())
} */

//dependency graph paketlerin calismasi icin gerekli olan bagimli oldugu paketleri gosteren bir graphictir. Versiyonlar ve bagimli olduklari kisimlar nodelar seklinde yazilirlar (detayli bilgi icin https://www.practical-go-lessons.com/chap-17-go-modules)

//go dilini yaratan kisiler MVS seklinde bir sey gelistirdiler. Bu minimal version selection anlamina gelmektedir. Go.mod dosyasini olusturur ve projede kullanilan butun dependenciesleri gosterir.


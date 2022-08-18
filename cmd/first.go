/*
go dilinde variablelar nasil tanimlanir?
degisken  isimlerine identifier denir degerlerine ise expression denir

var degiskenIsmi degiskenturu = deger(expression)   seklinde degisken tanimlamasi yapilabilir

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

degiskenismi := degiskentipi(degiskendegeri) seklinde de bir tanimlama yapilabilirmis

ornek => expected := uint(112)

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


//tum degisken tanimlamalari tek bir cati altinda

var degisken int32 = 23
var degisken1,degisken2 int = 23,34
var degisken1,degisken2 = 23 , "neden"      boyle bir tanimlamada istedigim turde degisken kullanabilirim ancak usttekinde sadece ayni turdeki degiskenleri kullanabilirdim
var sonradandegereklenecekdegisken int
sonradandegereklenecekdegisken = 23
degisken1 := 23
degisken1 := int(23)
const degiskenismi int = 23
const degiskenismi = "deneme"
var(
     geek1 = 100
     geek2 = 200.57
     geek3 bool
     geek4 string = "GeeksforGeeks"
)
bu sekilde butun degisken tanimlamalarini tek bir cati altinda toplamis oldum


==============================================================================
                    if else if else ve switch case yapisi
==============================================================================
if kosul{

}else if kosul{

}else{

}
! bu kullanim seklinden baska bir kullanim sekli yoktur. Kosullar parantezler icerisine alinamaz else ve else if ayri bir satirda yazilamaz

	if v := math.Pow(x, n); v < lim {
		return v
	}

    if kullaniminda short declaration denen bir sey vardir. Bununla birlikte kisa bir tanimlama if statement icerisinde kosuldan once yazilabilir

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
   onemli-> : init fonksiyonlari sira ile calisir yani ayni anda calismaz
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
 var nameDisplayer func(name, firstname string) string 
 
 bu sekilde bir kullanim mevcut bu kullanimin amaci ise sudur fonksiyonu henuz tamamlamadan tanimlama kismidir


===============================STRUCT YAPILARI================================================

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
//================methodlar ve pointerlar hakkinda kisa bilgiler===========================

// methodlar receiveri olan fonksiyonlara verilen isimlerdir

// package main

// import "fmt"

// func main() {
/*
   ! var DegiskenIsmi *DegiskenTuru = &AdresiGosterilecekDegisken
   ! var po *int = &i   seklinde pointer tanimlamasi yapabilecegim gibi
   ! po := &i   seklinde de pointer tanimlamasi yapabilirim
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
   ===================cikti================================
   i nin degeri:  1
   i nin adresi:  0xc00008a040
   pointer degeri:  0xc00008a040    pointerin degeri kisaca i nin adresidir
   Pointerin adresi:  0xc00009e038  pointerin kendisine has bir adresi vardir
   Pointerin isaret ettigi adresteki deger:  1
   Pointerin adresini tutan pointerin degeri:  0xc00009e038
   Pointerin adresini tutan pointerin adresi:  0xc00009e040
   Pointerin adresini tutan pointerin isaret ettigi adresteki deger:  0xc00008a040
   Pointerin adresini tutan pointerin isaret ettigi adresteki pointerin isaret ettigi adresteki deger:  1
*/


/* 
package main

import "fmt"

func main(){
    i:=1
    fmt.Println("initial:", i)
    zeroval(i)
    fmt.Println("zeroval:", i)
    zeroptr(&i)
    fmt.Println("zeroptr:", i)
}

func zeroval(ival int) {
    ival = 0
}

func zeroptr(iptr *int) {
    *iptr = 0
} */
// =============output========================
// initial: 1
// zeroval: 1
// zeroptr: 0

//kisaca burada anlatilmak istenen eger pointer seklinde kullanmazsam fonk icinde bir kopyasi olusturulur kendi orijinal degeri degistirilmez
// =========================================================================

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
bu kodun soyle bir mantigi var aslinda. internal klasoru icersinde cart isimli bir klasorum var o klasorun icerisinde cart.go isimli bir paketim (source code'a paket ismi veriliyordu) var. Bu paketi ustteki kodda import kisminda import etmisim. Daha sonra bu paketimin icerisinde Cart isimli bir struct'im ve iki adet metodum var. Metodlarimin receiveleri Cart tipinde oldugundan dolayi cart tipinde bir degisken olmadan metodlarimi kullanamam, bundan dolayi metotlarimi kullanabilmek icin newcart isimli Cart tipindeki cart.Cart{} olusturdum, daha sonra bunu kullanarak metotlarima erisebildim
=====================================//@ peki value receiver mi yoksa pointer receiver mi kullanmaliyim?=======================
eger value receiver olusturursam gonderdigim degisken method icerisinde kopyasi (internal) olusturulur ve o sekilde kullanilir  yani gonderdigim degisken degistirilmez.
! ayrica ustteki kullanim cok az da olsa memory kisminda fazlalik olusturur. Bu cogu zaman goz ardi edilebilir ancak memory onemli olan kullanimlarda dikkat edilmelidir
ancak eger pointer tipinde bir degisken yollarsam yolladigim degiskenin adresi gideceginden dolayi yollayacagim degisken degistirilir
@ receiver kullanirken genel kullanim type in first letteridir. c *Cart gibi
*/

// methods/example-project/main.go

/* package main

import (
	"fmt"
	_"github.com/Rhymond/go-money"
	"newgroupproject/internal/cart"
	_"newgroupproject/internal/product"
)

func main() {

    var newCart = cart.Cart{}   //(same as newcart:=cart.Cart{})
    total,error := newCart.TotalPrice()
    fmt.Printf("toplam tutar : %d hata kodu : %d",total,error)

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
	}
} */

//?================================================================================================================
//? bu kod blogu sadece receiver fonksiyonlari ve struct yapisini daha iyi anlamak icin var
//?================================================================================================================

/* package main

import (
	"fmt"
	"newgroupproject/internal/receiver"
    _"strconv"
) */

//alttaki koddaki receiver ismi yukarida import ettigim newgroupproject/internal/receiver isminden geliyor. (10 saat bunu dusundum amk)
/* 
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
//! ancak yeni eleman eklersem bu sonucumuza yansimaz. Pointer kullanmadan orijinal degere yeni eleman eklenmesi yapilamaz
*/

// package main

// import "fmt"


// func main() {
//     EUcountries := []string{"Austria", "Belgium", "Bulgaria"}
//     addCountries2(&EUcountries) //slice adresini yolladik
//     fmt.Println(EUcountries)
// }
// func addCountries2(countriesPointer *[]string) { //@ parametre olarak slice pointeri veriyoruz
//     *countriesPointer = append(*countriesPointer, []string{"Croatia", "Republic of Cyprus", "Czech Republic", "Denmark", "Estonia", "Finland", "France", "Germany", "Greece", "Hungary", "Ireland", "Italy", "Latvia", "Lithuania", "Luxembourg", "Malta", "Netherlands", "Poland", "Portugal", "Romania", "Slovakia", "Slovenia", "Spain", "Sweden"}...)
// }

// kisaca addcountries2 fonksiyonuna bir adres bilgisi yollandi ve fonksiyon kendi icerisinde olusturdugu degisken ile bu adres bilgisinin degerine  * operatoru ile ulasti

//? ============================================= receiver fonksiyon kullanarak bu sekilde yapilir ===============================================
/* type EUcountries []string

func main() {
    hey := EUcountries{"austuria,belgium,danmark"}
    hey.addCountries2()
    fmt.Println(hey)
} */

// func (countriesPointer *EUcountries) addCountries2() { //@  receiver olarak slice pointeri veriyoruz
//     *countriesPointer = append(*countriesPointer, []string{"Croatia", "Republic of Cyprus", "Czech Republic", "Denmark", "Estonia", "Finland", "France", "Germany", "Greece", "Hungary", "Ireland", "Italy", "Latvia", "Lithuania", "Luxembourg", "Malta", "Netherlands", "Poland", "Portugal", "Romania", "Slovakia", "Slovenia", "Spain", "Sweden"}...)
// }
//? ===============================================================================================================================================




//! simdi bir kullanim gosterecegim. adres bilgisi ile struct kullanabilir ve fieldlarina ulasabiliriz

/* func main(){

    type Cart struct {
        ID   string
        Paid bool
    }

    cart := Cart{
        ID:          "115552221",
        Paid:   true,
    }
    cartPtr := &cart    //! bu kisimda cartPtr isimli degiskene cart structurenin adres bilgisini atadik
    cartPtr.ID = "tunga"    //@ goruldugu gibi bu degisken uzerinden fieldlara ulasim saglayabildik
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
    (*cartPtr).ID = "asif tunga"    // onemli-> bir kullanim ornegi => burda cartPtr isimli pointerimizin isaret ettigi adresteki degerlere dereference operatoru sayesinde ulasiyoruz
    (*cartPtr).Paid = false

    fmt.Println(cart)

    //{asif tunga false} seklinde bir ciktisi olur
} */


/* 
//* bu alttaki kodlar ustte yamis oldugum durumun bir ornegi aslinda ozellikle Rename3 ve Rename4 func lari

package main

import "fmt"


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

func Rename4(address *Cat, newname string){
    address.Name = newname
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

    Rename4(adres, "TOM")
    fmt.Println(cat.Name)
    //! TOM yazdirir
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

func (c Cat) ReceiverAffection(from Human){     //! sadece Cat tipindeki structlar ile kullanilabilen (receiver olarak almis) ve Human tipindeki from degiskenini parametre olarak alan bir fonksiyon var bu kisimda
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

func Pet(animal DomasticAnimal, human Human){   //! Pet fonksiyonu animal isminda DomesticAnimal(interface) tipinde bir degisken parametresi aliyor
    animal.GiveAffection(human)
    animal.ReceiverAffection(human)
}
//@ ust kisimda yaptigim kisaca ikinci bir fonksiyondan interface'e ulasip o interface icersinde tanimlanmis olan fonksiyonlara erismek.

func main(){
    Pet(Dog{"MISHA"},Human{"asif","tunga",23,"turkiye"})
    Pet(Cat{"TOM"},Human{"hediye","ozgul",20,"bulgaristan"})
    }  */

// ?================================ anlamama yardimci olacak baska bir ornek olabilir ===================================================
/* 

package main

import "fmt"

type Article struct {
	Title string
	Author string
}

type Book struct{
    Title string
    Author string
    Pages int
}

func (b Book) Strings() string{
    return fmt.Sprintf("The %q book was written by %s",b.Title,b.Author)
}

func (a Article) Strings() string {
	return fmt.Sprintf("The %q article was written by %s.", a.Title, a.Author)
}

type Stringer interface{
    Strings() string //@ burada yaptigim tek sey Stringer isminde bir interfacein icerisine String() isimli string geri donuslu bir metot koydum. Interface tanimlamasi bu kadar
}

func main() {
    a := Article{
        Title: "Understanding Interfaces in Go",
        Author: "Sammy Shark",
    }
    b:=Article{"Umutlu Bir Bilgisayar Muhendisinin Yasami","Asif Tunga Mubarek"} // ayni zamanda bu sekilde de kullaniliyor (daha fazla bilgi icin yukaridaki struct yapisina bakabilirim)
    c:=Book{"Yasamim","Asif Tunga Mubarek",23}

    Prints(a) //bak-> dikkat ettiysem interface'i parametre olarak alan fonksiyona parametre olarak gonderdigim sey interfacein kendisi degil hatta icerisindeki fonksiyonu bile degil. O fonksiyonu (merhodu) kullanabilen, metot yaziminda tanimlanmis receiverin tipindeki structtur.
    Prints(b)
    Prints(c)
}

func Prints(s Stringer) {
    fmt.Println(s.Strings())
}

//burada yaptigim sey yine farkli degil. Elimde Stringer isminde bir interface var. Bu interface icerisinde Strings() isminde bir fonksiyonum var. Bu method birden cok olabilir (hepsinde farkli receiverlari olabilir mesela). Ayrica elimde Prints(s Stringer) seklinde bir adet methodum daha var (parametre olarak interface aliyor). Bu methodu kullanirken parametre olarak interfacei yollamak yerine interface icerisindeki metodun receiverinda type olarak bulunan bir struct yollamam gerekir.
*/



/* 
-----------------bu kod blogunun ciktisi ise su sekilde olur---------------------------------
The "Understanding Interfaces in Go" article was written by Sammy Shark.
The "Umutlu Bir Bilgisayar Muhendisinin Yasami" article was written by Asif Tunga Mubarek.
The "Yasamim" book was written by Asif Tunga Mubarek
*/

//==========================================================================================================================================
// onemli-> bir not : //! bir method fonksiyondan farkli olarak sadece tanimlandigi turun orneginden cagrilabilir
//A method is a special function that is scoped to a specific type in Go. Unlike a function, a method can only be called from the instance of the type it was defined on.

//simdi bu yapi beni kod tekrarindan kurtardi. Bunu bir ornekle gostermek istiyorum
/* 
package main

import "fmt"

type Yazar struct{
    Name string
    Age int
    Country string
}

type Basimevi struct{
    Locasyon string
    Country string
}

type Article struct {
	Title string
	Author string
}

type Book struct{
    Title string
    Author string
    Pages int
}

func (b Book) String() string{
    return fmt.Sprintf("The %q book was written by %s",b.Title,b.Author)
}

func (a Article) String() string {
	return fmt.Sprintf("The %q article was written by %s.", a.Title, a.Author)
}

func (y Yazar) String() string {
	return fmt.Sprintf("The %q article was written by %s.", y.Name, y.Country)
}

func (b Basimevi) String() string {
	return fmt.Sprintf("The %q article was written by %s.", b.Locasyon, b.Country)
}

//@ asagidaki print fonksiyonunu da degistirmemis oldugumu varsayalim (bundan kastim aslinda su: degistirmemisim derken interface ile kullanmamisim anlaminda soyleyebilirim)

func (a Article)Print(){
    fmt.Println(a.String())
}

func (b Book)Print(){
    fmt.Println(b.String())
}

func (c Yazar)Print(){
    fmt.Println(c.String())
}

func (d Basimevi)Print(){
    fmt.Println(d.String())
}

//not-> farkli parametrelerle olsa dahi ayni isimle birden fazla fonksiyon yazilamaz. Ancak farkli receiverlar ile ayni isimde birden fazla method yazilabilir 

func main(){
   a:=Article{"baslik","asif tunga mubarek"}
   b:=Book{"kitap basligi","asif tunga mubarek",300}
   c:=Yazar{"asif tunga mubarek",26,"antakya"}
   d:=Basimevi{"mersin","turkey"}
//onemli-> oncelikle bunun kullanim sekli Print(a) seklinde degildir. Print(a) kullanimi sadece parametreler icin gecerlidir. Receiver oldugundan alttaki sekillerde kullanilir
    a.Print()
    b.Print()
    c.Print()
    d.Print()
}
  */
//*==================================== peki bunu interface ile yapmis olsaydim nasil mi olurdu? ===================
/* 
package main

import "fmt"

type Yazar struct{
    Name string
    Age int
    Country string
}

type Basimevi struct{
    Locasion string
    Country string
}

type Article struct {
	Title string
	Author string
}

type Book struct{
    Title string
    Author string
    Pages int
}

func (b Book) String() string{
    return fmt.Sprintf("The %q book was written by %s",b.Title,b.Author)
}

func (a Article) String() string {
	return fmt.Sprintf("The %q article was written by %s.", a.Title, a.Author)
}

func (y Yazar) String() string {
	return fmt.Sprintf("The %q sehrinde yasayan kisi: %s.", y.Country, y.Name)
}

func (b Basimevi) String() string {
	return fmt.Sprintf("The %q lokasyonu su ulkede bulunur: %s.", b.Locasion, b.Country)
}

type Stringer interface{
    String() string
}

func Print(s Stringer) {
    fmt.Println(s.String())
}


func main(){
   a:=Article{"baslik","asif tunga mubarek"}
   b:=Book{"kitap basligi","asif tunga mubarek",300}
   c:=Yazar{"asif tunga mubarek",26,"antakya"}
   d:=Basimevi{"mersin","turkey"}
   Print(a)
   Print(b)
   Print(c)
   Print(d)
}
 */
/*
KONU ILE ALAKASIZ VIM NOTU : eger bir sonraki karaktere gitmek istiyor isem yapmam gereken tek sey insert moda a ile gecis yapmak. veya bunun yerine direkt olarak cumlenin sonuna gitmek istiyor isem A seklinde kullanabilirim ama bu direkt insert moda sokar haberim olsun
$ seklinde de kullanabilirim bu insert moda sokmaz direkt cumlenin sonuna giderim bu sekilde
0 ile de cumlenin basina giderim
*/

//@ interfaceleri anlamak icin kendime alacagim not kisimlari buralar
//? burada alinan notlar sadece interface konusunu daha iyi anlamak icindir. Her bir not acik bir sekilde yazilacaktir.
// burada anlatilan kodlar https://go.kaanksc.com/boeluem-3/arayuez-interface adresinden alinmistir. Detayli bilgi icin bu siteyi inceleyebilirim

/* package main

import "fmt"

type domatesSalçasi struct {
    marka string
}

type BiberSalcasi struct {
    marka string
}

func (s domatesSalçasi) Ye() (geridonus string){
    fmt.Println("Domates salçasi yenildi")
    return
}
func (s BiberSalcasi) Ye() (dondur string) {
	fmt.Println("Biber salçasi yenildi")
    return
}

type Salca interface{
    Ye() string    //bu kisim aslinda Ye() isimli bir fonksiyonun string bir ifade dondurdugunu anlatir
}

func Besle(salcam Salca){
    salcam.Ye()
}

func main(){
    Besle(BiberSalcasi{"tamek"}) //burada biber salcasi struct ornegi kismini parametre seklinde gonderdik. Aslinda ust kisimdaki besle fonksiyonunda interface i parametre olarak kullanmamiz sayesinde (ancak interface gondermiyoruz parametre olarak buna dikkat etmeliyim) interface uzerinden fonksiyona erisiyorum gibi dusunebilirim
}
 */
//*----------------------------------- bu kisim ise ust kisim yerine yapilabilir----------------------------------------------
/* func main(){
    biber := BiberSalcasi{"tamek"}  //! biber salcasi tipinde bir struct olusturduk
    var salcam Salca = &biber   //! daha sonra bu struct icin bir pointer olusturduk
    fmt.Print(salcam.Ye())
}*/
//@ ustteki ile bu ayni sonucu verir. Burada yaptigim sey ise aslinda deneme amacli olarak yazdigim kisimdan cok da farkli degil. yesil kismin ust tarafinda bir fonksiyon olusturdum ve o fonksiyon vasitasi ile interface yapisini parametre olarak gonderip interface yapisi uzerinden kontratta yani interface uzerinde belirlenmis olan fonksiyona interface uzerinden ulasmis oldum. (struct yapimi interface uzerinden gonderdim)
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
//     s.Ye()   //not-> altta kirmizi kisimda buradan bahsediyorum
// }

//not-> benim buradaki hatam aslinda tam bu kisimda. Dusunceme gore bir interface'e ulasmak diger interface icerisindeki fonksiyonu calistirir ancak bu tamamen yanlis bir dusunce. Bir interface uzerinden diger fonksiyona ulastigim dogru ancak diger fonksiyonu cagirmadan nasil calismasini bekleyebilirim ki 


// func main(){
//     Olustur(BiberSalcasi{"tunga salcalari"})
// }
//========================================================================= bir diger yol ile yapimi

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
//onemli-> kisaca kendime interfaceler icin su notu diyebilirim. Iki farkli ulasim soz konusu aslinda. Istersem hemen ustte oldugu gibi main fonksiyonu icerisinden ulasabilirim, bir diger yontem bunun bir ustunde oldugu gibi yeni bir fonksiyon tanimlayip parametre olarak interface verip onun uzerinden ulasabilirim

//STANDART LIBRARYDE BULUNAN BIRCOK INTERFACE VARDIR BUNLARDAN EN UNLULERI ISE SUNLARDIR : Error interface, Stringer interface, Interface interface bu interfaceleri daha detayli incelemek istersem ; https://www.practical-go-lessons.com/chap-16-interfaces adresinden gorebilirim

//Go da bir diger interface tipi ise bos interfacedir. Bu yazabilecegim en kucuk ve en kisa interfacedir. Bu su sekilde yazilir interface{}
//peki ben bunu nerede kullanacagim? tanimlama olarak bos bir interface deger her tipte degiskeni tutabilir. Eger her tipi kabul eden bir metot yazmak istersem interface kullanabilirim

// func (l *Logger) Fatal(v...interface{}) {}   bu sekilde bir kullanim butun tipleri kabul eder
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

/*  
//@ API nedir sorusunun cevabi niteliginde
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

//dependency graph paketlerin calismasi icin gerekli olan bagimli oldugu paketleri gosteren bir graphtir. Versiyonlar ve bagimli olduklari kisimlar nodelar seklinde yazilirlar (detayli bilgi icin https://www.practical-go-lessons.com/chap-17-go-modules)

//go dilini yaratan kisiler MVS seklinde bir sey gelistirdiler. Bu minimal version selection anlamina gelmektedir. Go.mod dosyasini olusturur ve projede kullanilan butun dependenciesleri gosterir.

//her modul kendisinin kullandigi gerekli modullerin listesini vermelidir. Moduller modul pathi ile tanimlanirlar. Her modulun minimum uyumlu versiyonu belirtilmelidir. Bu  dosya go.mod dosyasindadir.
//Her modul import compatibility rule kuralina uymak zorundadir. Bu kural; ayni module pathine sahip moduller geriye uyumlu olmalidir der
// yeni bir dependency (bagimlilik) nasil eklenir? go get komutu ile dependecy eklenebilir. Bu komut gerekli modulu indirir ve go.mod dosyasina kaydeder
//peki go hangi versiyonlari secer? => 1) en son etiketlenmis kararli surum 2) son etiketlenmis pre-release surum 3) son etiketlenmemis versiyon (son eklenmis olan commit ayrica pseudo-version) diye adlandirilir

//@ Yapi listesinin olusturulmasi=> built list is Go programinin build edilmesi icin gerekli olan modullerdir
//bu list elemanlarinin hepsi iki seyden olusurlar 1) modul path module' u tanimlar 2) revision tanimlayicisi (tag veya commit id)
//! belirli bir modul icin build list olusturmak icin gerekli olan seyler sunlardir
/*
initialize empty list L
current module icin gerekli modul listelerini al
gerekli olan her modul icin :
    bu modul icin gerekli olan modul listesini al
    bu elemanlari L listesine ekle
    bu operasyonu eklenen elemanlar icin tekrar ettir
    sonunda liste ayni modul pathina sahip birden fazla modul girisi icerebilir eger durum boyleyse her module path icin en yeni versiyonu tut (demek istedigi aslinda su : eger modul path ayni ise en yeni versiyon olani listede tut digerini sil)
//@ son build list of modulu ekrana yazdirmak icin go list -m all seklinde bir komut kullanabilirim
//@ dependencyleri son minor veya patchlerine guncellemek istersem eger go get -u komutunu kullanabilirim. -u newer minor veya patch indirir
//* diyelimki benim projem su modulu kullaniyor olsun gitlab.com/loir402/bluesodium
//* sonrasinda kullanilan bu modulun yeni bir surumu cikti diyelim (v2.0.1)
//* go get -u gitlab.com/loir402/bluesodium kullanarak  yeni versiyonu indiremem bunun nedeni yeni olan modulun different bir path i olmasidir
//* bu da import compatibility kuralina gore moduller backward compatible olmalidir. Bu modul bu kurali bozmaktadir
//@ bir modulun son surumu son degisiklikleri getirir bu da onceki surumlerin kullanicilarini etkileyecektir
//@ bundan dolayi ayri bir module pathine sahip olmalidir
go.mod dosyasina bakarsask sununla karsilasiriz

module gitlab.com/loir402/bluesodium/v2

go 1.15

module path degisti artik gitlab.com/loir402/bluesodium yerine gitlab.com/loir402/bluesodium/v2 oldu

onemli-> bir modul v0 dan v2 ye veya v1 den v2 ye guncelleniyorsa import compatibiliry rule kuralina uymak icin pathini degistirmelidir

!  To specifically require the major version 2, you need to launch the command. :

!  go get -u gitlab.com/loir402/bluesodium/v2

*/
/*
her gereksinim en son modul listesi gerekiyormus gibi okunur
import compatibility rule sayesinde hicbir kirilma yasanmamalidir (bu operation sadece new patches ve minor versionlari gerekli kilar)
Eger yeni surumler bulunduysa build list ve go.mod dosyasi modified edilir

bunu duzgun bir sekilde yapmak icin yapilmasi gereken komutlar sirasi ile

go get -u ./...
go: github.com/andybalholm/cascadia upgrade => v1.2.0
go: golang.org/x/net upgrade => v0.0.0-20210119194325-5f4716e94777

bu command output the upgraded dependencies

go.mod dosyasina bir bakalim

module thisIsATest

go 1.15

require (
    github.com/andybalholm/cascadia v1.2.0 // indirect
    gitlab.com/loir402/bluesodium/v2 v2.1.1
    golang.org/x/net v0.0.0-20210119194325-5f4716e94777 // indirect
)

burada yazan indirect kelimeleri ; Bu satırlar, derleme listesini oluştururken bu belirli yükseltilmiş sürümleri kullandığımızdan emin olmak için eklenir. Aksi takdirde, yapı listesi aynı kalacaktır.


//@ belirli bir modulu yeni bir versiyona guncellemek icin neler yapilir

Hiçbir yükseltme yapılmamış gibi bir ilk yapı listesi oluşturulur.
İstenen yükseltme ile ikinci bir tane inşa edilir.
Daha sonra iki liste birleştirilir.
Bir modül listede iki kez listeleniyorsa, en yeni sürüm seçilir.

//@ belirli bir modulu downgrade nasil yapilabilir
//! tabi bunu yapmak icin nedenlerimiz olabilir : modulun yeni surumu bir bug yaratiyorsa, bir guvenlik acigi varsa, application performansini dusuruyorsa ve bunun gibi bircok nedenden dolayi downgrade yapmak isteyebiliriz

Diyelim ki v1.1.0 sürümünde E modülünü kullandık ve E sürümünü v1.0.0'a düşürmek istiyoruz.
Go, uygulamanın her gereksinimini alacaktır:
Yapı listesi, bu özel gereksinim için oluşturulmuştur.
Derleme listesi yasaklanmış bir E sürümü içeriyorsa (yani, v1.1.0 veya üzeri sürümlerde E)
Daha sonra, bu gereksinimin sürümü düşürülür. Düşürülmüş sürüm, v1.1.0 (veya üzeri) sürümünde hala E içeriyorsa, önceki sürümde düşürülür
Görev, yasaklı sürümler derleme listesinde olmayana kadar tekrarlanır.
Bu işlem, artık gerekmeyen gereksinimleri potansiyel olarak kaldıracaktır.

//encode is the process of converting a piece data from a format to another format
from the output we can go back to the input


! go.sum dosyasi module direct ve indirect kriptografig hashleri tutar

//@ ornek olarak bir go.sum dosyasi olusturalim

go mod init

boylelikle bir go.mod dosyasi olusturduk ancak bu bos??? gerekli dependency ler go.mod dosyasina eklenmemis bunu elle yapmaktan kacinmak icin go install komutunu calistiririz bu komut sayesinde go.mod dosyasi guncellenir ve go.sum dosyasi olusturulmus olur

go.sum dosyasi her bir modul icin iki satir icerir
gitlab.com/loir402/foo v1.0.0 h1:sIEfKULMonD3L9COiI2GyGN7SdzXLw0rbT5lcW60t84=
gitlab.com/loir402/foo v1.0.0/go.mod h1:+IP28RhAFM6FlBl5iSYCGAJoG5GtZpUH4Mteu0ekyDY=

ilk satir modulun dosyalarinin tamaminin sha256 hashidir

ikinci satir modulun go.mod dosyasinin hashidir

bu hashlar daha sonra base 64 e convert edilir
h1 kelimesi sabittir bunun anlami go librarysinde kullanilan hash1 fonksiyonun calismis oldugudur
buradaki checksum bağımlı modüllerin indirilen sürümlerinin ilk indirme ile aynı olduğundan emin olmak için buradadır


patch aslinda sudur :

package corge

func Corge() string {
    return "Corge"
}

=========== bu alttaki kisim ile aslinda patch etmis oluyoruz kisaca api degistirmemis oluyor (api demek aslinda modul tarafindan disa aktarilan her sey demektir)

package corge

func Corge() string {
    return fmt.Sprintf("Corge")
}

*/
// ======================

/* package main

import (
    "fmt"
    "newgroupproject/internal/bar"
    "newgroupproject/internal/foo"
)
func main(){
    fmt.Println(foo.Foo())
    fmt.Println(bar.Bar())
} */

//! diyelim ki bir tane fonksiyon yazdik ve onu api gibi kullandik simdi burada onemli birkac notum var
//@ oncelikle kullandigimiz bu api kismini degistirip daha sonra tekrar guncellemek istersem diye asagidaki notlar cok onemli

//! butun dependencyleri son surumlerine update etmek icin go get -u ./... bunu yaptigimizda go yapilan patchleri anlar ve onlari otomatik olarak indirir. go.mod dosyasina yeni bir satir eklenir
//@ ancak go.sum dosyasi degismez bunun nedeni ileride olurda downgrade etmek istersek diye bir guvenlik onlemidir (elbette yeni versiyon da eklenir sadece eski versiyonu silinmez demek istiyorum)
//! bir dependencyi son surumune update etmek go get blablalinki
//@ bir dependencyi spesifik bir surume guncellemek icin ise go get module_path@X
//burada x commit hashi olabilecegi gibi versiyon numarasi da olabilir
/*
ornek olarak bir sey vermek istiyorum. Onceden su sekilde olan public api miza su sekilde bir fonksiyon eklemesi yaptim bu fonksiyon eklemesinden sonra artik api degisecegi icin onceden v1.0.1 olan surumumuzu v1.1.0 sekline getirdim
// bar/bar.go
package bar

import (
    "fmt"

    "gitlab.com/loir402/baz"
    "gitlab.com/loir402/qux"
)

func Bar() string {
    return fmt.Sprintf("Bar %s %s", baz.Baz(), qux.Qux())
}

//@ func Bar2() string {
    //@ return fmt.Sprintf("Bar2 %s %s", baz.Baz(), qux.Qux())
//@ }          yeni eklenen fonksiyon

yapilacak islemler ve degisen dosyalar ise su sekildedir

$ go get gitlab.com/loir402/bar@v1.1.0
----------------------------------------
go: finding gitlab.com/loir402/bar v1.1.0
go: downloading gitlab.com/loir402/bar v1.1.0

go.mod dosyasinda ise gerceklesen degisiklikler su sekildedir

module gitlab.com/loir402/myApp

require (
    gitlab.com/loir402/bar v1.1.0
    gitlab.com/loir402/corge v1.0.1 // indirect
    gitlab.com/loir402/foo v1.0.1
)


simdi ise go.sum dosyasinda nasil degisiklikler olduguna bir goz atalim

gitlab.com/loir402/bar v1.0.0 h1:l8z9pDvQfdWLOG4HNaEPHdd1FMaceFfIUv7nucKDr/E=
gitlab.com/loir402/bar v1.0.0/go.mod h1:i/AbOUnjwb8HUqxgi4yndsuKTdcZ/ztfO7lLbu5P/2E=
gitlab.com/loir402/bar v1.1.0 h1:VntceKGOvGEiCGeyyaik5NwU+4APgyS86IZ5/hm6uEc=
gitlab.com/loir402/bar v1.1.0/go.mod h1:i/AbOUnjwb8HUqxgi4yndsuKTdcZ/ztfO7lLbu5P/2E=

goruldugu gibi eski olan versiyonlari safety olmasi acisindan hala tutulmaktadir (downgrade yapilabilme ihtimaline karsi)

//* downgrade a dependecy ise su sekilde yapiyoruz

$ go get gitlab.com/loir402/bar@v1.0.0

bu geri dondurur ve v1.0.0 versiyonuna geri donmus oluruz

//? son olarak ise bir seyden daha bahsetmek istiyorum. Soyle bir senaryo hayal edelim. Yeni bir paket olusturdum ve bu paketi main paketimin icinde api seklinde kullandim. Daha sonra go mod install yaparak bu paketimi go.sum ve go.mod dosyalarinin icerisine kaydedilmesini sagladim. Daha sonra fikir degistirdim ve bu fonksiyonu disari api seklinde aktaran bu paketimi sildim. Ancak go.sum ve go.mod dosyalarinda silinmemis sekilde durmaya devam etmekte. Bunun onune gecebilmek icin go mod tidy -v komutu kullanilir
//go aracı tarafından indirilen bağımlılıkların farklı sürümleri $GOPATH/pkg/mod dizininde durmaktadir

//bak-> ======================================================================================================================

oncelikle lock file in ne oldugundan bahsetmek istiyorum. lock file kisaca list of dependency that your project use along with a specific revision(tag or commit hash)
Peki lock dosyasinin go.sum dosyasindan farki tam olarak nedir > go.sum dosyasi projede direkt olarak kullanilan modullerin veya dolayli olarak kullanilan modullerin kriptographic sum larini ve versionlarini saklar
go.sum dosyasının amacı, modüllerin bir sonraki indirmede değiştirilmemesini sağlamaktır. Bir kilit dosyasının amacı, tekrarlanabilir yapılara izin vermektir.

Go use a deterministic approach to version selection.

The build list will remain stable with time (if revisions used are not deleted by maintainer)

It means that the build list generated in January will not be different from the one generated in December.

Thus the introduction of a lock file is not necessary.

Dependency management systems that have introduced lock files generally don’t have such a deterministic approach

The lock file is needed to ensure that the builds of the application are reproducible by listing all the dependencies used and at which version.


go.sum dosyalarini ve go.mod dosyalarini mutlaka commit etmeliyim bunun nedeni => Others need the go.mod file to construct their build list.

Others will use the go.sum to ensure that the module downloaded have not been altered.


kullanilan diger komutlar :
    go mod graph -> standart output seklinde dependency graph yazdirir
    go mod vendor -> dependencylerin sourcelari ile birlikte bir vendor dosyasi olusturur (bunu su sekilde dusunebilirim aslinda sanki githubtan kullandigim bir projeyi klonlamisim gibi direkt olarak benim bilgisayarima indiriyor)
    go mod verify -> Adından da anlaşılacağı gibi, mod doğrulaması yerel olarak depolanan bağımlılıklarınızı kontrol edecektir. Go, yerel olarak depolanan bağımlılıklarınızın değiştirilmediğini kontrol edecektir. Bu denetim, bağımlılıklarınızın değiştirilmiş bir sürümünü değil, doğru sürümünü kullandığınızdan emin olmak için çok kullanışlıdır. Bu değişiklikler buildin başarısız olmasına neden olabilir.


==========================================================================================================================
A minor version introduces breaking changes. True or False?

    False

    Major versions introduce breaking changes

What is the command to update a module to its latest version?

    $ go get -u path/of/the/module

What is the name of the set of algorithms used in Go to manage dependencies?

    MVS: Minimum Version Selection

Create a sentence describing a Module with the following words: packages, go.mod, go.sum.

    A module is a group of packages and two files, a go.mod & a go.sum.

What is the purpose of the go.mod file?

    The go.mod file will define the module path of the current module

    The go.mod file lists minimum versions of direct dependencies.

    It also lists the minimum versions of indirect dependencies

        It happens when the developer upgraded or removed some dependencies manually

    Each dependency is identified by a module path.

    It also gives the expected language version for the module

What is the purpose of the go.sum file?

    The go.sum file is to ensure that the source code of dependencies downloaded are the same as the one downloaded by the original developer.

What is the command to initialize a module?

    In your module directory :
    $ go mod init path/of/the/module

What is the command to display the build list of a module?

    In your module directory :
    $ go list -m all

When the major version 2 is released, the module path is not modified. True or False?

    False

    Because a new major version introduces breaking changes, the module path should change (to respect the import compatibility rule)

    The string “v2” should be added to the module path


============================== kisaca ozetlemek gerekirse eger =====================================================
A Go module is a set of packages that are versioned together with a version control system (for instance, Git)

Go modules are identified by a module path.

    Module paths describe what the module does and where we can find it

A version is identified by a tag that describes the version changes.

To describe what changes are added in a version, we usually use a versioning scheme which is a set of rules enforced by developers

Semantic Versioning is a versioning scheme that Go uses

    In this scheme, a version number is a string formatted this way => “vX.Y.Z” (v is optional)

        X, Y, Z are unsigned integers.

        X is the major version number. When you increase this number, it means that you introduce breaking changes

        Y is the minor number. This number should be increased when new non-breaking features are added.

        Z is the patch number. This number is increased when a patch is created (a bug fix, for instance)

When a module publish a new major version greater or equal than 2, it should add a major version suffix to the module path

    gitlab.com/loir402/bluesodium becomes gitlab.com/loir402/bluesodium/v2

We can initialize a Go module in an existing project by executing the go mod init my/new/module/path command.

To add a new direct dependency to a program, use the go get command:

$ go get my/new/module/to/add

To upgrade all your dependencies, use the go get command:

$ go get -u ./...

To upgrade one dependency, use the go get command:

$ go get -u the/module/path/you/want/to/upgrade

In the go.mod file, you can replace the code of a module by another one (stored on a code-sharing website or locally)

replace gitlab.com/loir402/corge => ./corgeforked

You can also exclude a specific version from your builds.

exclude gitlab.com/loir402/bluesodium v2.0.1

===============================================================================================================================

go get komutu modullerin en son surumlerini getirmek icin bir proxy kullanir
serverlari kisaca aciklamak gerekirse server dedigimiz sey clientlara functionality provide eden applicationlardir.
client ornek olarak bir web browser olabilir

eger bu functionality yi internet uzerinden saglayan bir uygulama (application) var ise biz buna web server diyoruz

proxy ise client ile server arasindadir ve bir nevi araci sekilde calisir

proxy server resourcelar icin gelen requestleri alir
    requestleri kendisi karsilayabilir (direkt olarak)
    arkasindaki servera aktarabilir ardindan arkasindaki serverdan gelen yanitlari clienta geri dondurebilir (bu tarz proxylere forward proxy diyoruz)
    requesti durdurabilir

Proxy serverlar genel olarak su amaclar icin kullanilir :
    internet trafigini izleyebilirler (genel olarak sirketler kullanicilarinin internet trafiklerini bu yontemler ile izlerler)
    bazi siteleri yasaklayabilirler
    Caching yapabilirler
        proxyleri resourcelari caching yapmak icin kullaniriz
        hedeflenen sunucuyu aramak yerine clienta cache alinmis bir response gonderirler


go module proxyler son zamanlarda dile eklenen bir seydir, go get komutu kullanilmadan once moduller dogrudan web sayfasi uzerinden indirilirdi
bunun kotu yanlari modul paylasilan web sitesinden silinebilirdi veya tagler versiyon numaralari silinebilirdi. Eger bizim applicationumuz bu module bagimli ise build sirasinda hatalarla karsilasabilirdik

======================================= UNIT TESTLER ========================================================================

*/

// package main

// import "fmt"

// func main(){
//     price := TotalPrice(3,100,500)
//     fmt.Println(price) //alttaki kod icin bu sekilde bir deneme yapabiliriz
// }

// func TotalPrice(nigths, rate, citytax uint) uint{
//     return nigths*rate+citytax //bu fonksiyonu yazdik peki biz bunun dogrulugundan nasil emin olacagiz bunun icin unit testler yazmamiz gerekir
// }

//dogrulugundan emin olmak icin direkt olarak sonucu if degerleri icerisinde main fonksiyonunda test edebiliriz. Tabi oyle bir durum ne kadar efektif olabilir? Unit testte tum sistem test edilmez -> Sistemin kucuk ve belli parcalarini alir test ederiz

//Bazi kisiler unit testlerin gereksiz oldugunu dusunebilir bazilari ise yaptiklari her degisiklikte test ederler. Yapilan bu testleri otomatik yapabildigimizi dusunelim? hatta bu testleri her build aldigimizda calistigini dusunelim hatta daha da iyisi bu testlerin pakette herhangi bir degisiklik yapildiginda calistigini dusunelim

/*
single unit testlere test case denir. Bir grup test case'e ise test set (veya test suite) denir
Soyle bir durum hayal edelim. Bir stringi alip buyuk harflere ceviren bir fonksiyonumuz olsun
test case sunlardan olusur :
    bir test inputu (ornek olarak coffee)
    bir expected output (in our case it will be COFFEE)
    the actual output of our function
    Bu string testini gerceklestirmek icin go paketleri kullanabilir veya go nun bize saglamis oldugu string karsilastirma islemlerini yapabiliriz. Test caselerin bu kismina ise assertion denir.

Unit testler bir fonksiyonun veya metodun duzgun calisip calismadigini kontrol ederler. Unit testler olmadan developerlar kodlarini development fazinda test edebilirler. Bu tarz testler manualdir ve tekrarlanamaz. Development fazi bittikten sonra bu manual testler calistirilamaz.

eger testleri proje dosyasina koyarsak bu testleri sonrasinda da kullanabiliriz. Bu sayede projeyi nasty regressionslara karsi korumus oluruz (yeni featurelarin onceden varolan bir seyleri break etmesi)

API design genellikle unit testler yazarken gelisir bunun nedeni gelistirdigimiz fonksiyonu ayni zamanda cagiracak olmamizdir. TDD yontemlerini kullanirsak bu gelisme daha da bariz bir hale gelir.

Unit testler ayni zamanda code documentation islemi de gorur. Specific bir fonksiyonun nasil calistigini merak eden kullanicilar unit testlere bakarak kolayca anlayabilir.

Test dosyalari nereye konmali? Cogu programlama dilinde testler "test" isminde ayri bir klasorde bulunur ancak GO programlama dilinde testler direkt olarak pakette bulunur

compare.go
compare_test.go
replace.go
replace_test.go

seklinde bir hierrarchi vardir
Burada bir kullanim dikkatimi cekti. xxx.go seklindeki bir paket icin xxx_test.go seklinde bir test case ismi vardir. xxx_test.go seklinde ismi olan bir paket dosyasi kod derlendigi zaman derleyici tarafindan gormezden gelinir.
*/
//@ foo klasorune bak
//test fonksiyonlari func TestFonksiyonIsmi(t *testing.T) seklinde olmalidir. Testten sonra genlen fonksiyon ismi mutalaka capital ile baslamalidir.

//testlerdeki basariyi olcecek bir T tipinde tanimlanmis bir yontem yoktur. If unit test return without calling a failure method its considered success

/*
To signal a failure, you can use the following methods :

Error : will log and marks the test function as failed. Execution will continue.

Errorf : will log (with the specified format) and marks the test function as failed. Execution will continue.

Fail : will mark the function as failing. Execution will continue.

FailNow : this one will mark the test as failed and stop the execution of the current test function (if you have other assertions, they will not be tested).

You also have the methods Fatal and Fatalf that will log and call internally FailNow.


@ bazen unit testleri destekleyecek test dosyalarini saklamamiz gerekir. Ornegin bir ornek configuration file veya a model CSV file (for application generate file). Bu tarz dosyalari testdata folderinda saklamaliyiz

kisacasi paket icerisindeki testdata folder unit testlerin kullanacagi filelari tutar


Go standart kitaplığı, birim testinizi harici kitaplıklar olmadan oluşturmak için gerekli tüm araçları sağlar. Bu gerçeğe rağmen, harici "assertion kitaplıkları" kullanan projeler görmek yaygındır. Bir assertion kitaplığı, assertion oluşturmak için birçok işlev ve yöntem sunar. Çok popüler bir modül github.com/stretchr/testify'dır.
*/

/* To add it to your project, type the following command in your terminal :

$ go get github.com/stretchr/testify

As an example, this is the previous unit test written with the help of the package assert from the module github.com/stretchr/testify :

package foo

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestFoo(t *testing.T) {
    assert.Equal(t, "Foo", Foo(), "they should be equal")
} */

//unit testleri calistirmak icin go test komutu kullanilir ancak bu komutun kullanilabilmesi icin testin yazildigi dizinde olunmalidir

//projedeki butun testleri calistirmak icin ise go test ./... komutu kullanilir

/*
bunu kullandigimiz zaman su sekilde bi sonucla karsilasiriz
$ go test ./...
?       newgroupproject/cmd     [no test files]
?       newgroupproject/internal/bar    [no test files]
?       newgroupproject/internal/baz    [no test files]
?       newgroupproject/internal/cart   [no test files]
?       newgroupproject/internal/corge  [no test files]
?       newgroupproject/internal/email  [no test files]
--- FAIL: TestFoo (0.00s)
    foo_test.go:9: expected Foo do not match actual Foor
FAIL
FAIL    newgroupproject/internal/foo    0.242s
?       newgroupproject/internal/invoice        [no test files]
?       newgroupproject/internal/new    [no test files]
?       newgroupproject/internal/product        [no test files]
?       newgroupproject/internal/qux    [no test files]
?       newgroupproject/internal/receiver       [no test files]
?       newgroupproject/internal/romantoint     [no test files]
?       newgroupproject/internal/time   [no test files]
?       newgroupproject/internal/user   [no test files]
FAIL
*/

/*
peki fail olan bir unit testin outputu nedir?

Tunga@ATM MINGW64 ~/Desktop/newgroupproject/internal/foo (master)
$ go test
--- FAIL: TestFoo (0.00s)
    foo_test.go:9: expected Foo do not match actual Foor
FAIL
exit status 1
FAIL    newgroupproject/internal/foo    0.239s
*/
//Program, 1 durum koduyla çıkar ve sürekli entegrasyon araçları oluşturmak istiyorsanız bunu otomatik olarak algılamanıza olanak tanır.
//0 dan farkli bir sinyal kodu hata mesaji donduruldugu anlamina gelir
//0 ise herhangi bir hata olmadigi mesajini dondurur
//go test komutunu çalıştırdığınızda, go aracı test edilen paketlerde (paketlerin ve test dosyalarının kaynağı) otomatik olarak go vet'i çalıştıracaktır.

//go vet komutu, Go araç zincirinin bir parçasıdır. Olası hataları tespit etmek için kaynak kodunda syntax doğrulaması gerçekleştirir.

//Birim testlerini başlatmadan önce bir dizi go vet komutunu otomatik olarak çalıştırmak harika bir fikir. Programınıza zarar vermeden önce hataları keşfetmenizi sağlayabilir!

//@ testleri calistirmadan compile etmek icin go test -c komutunu yazmamiz yeterlidir. Bunun sonucunda packageName.test seklinde bir adet dosya olusur.
/// denemem sonucunda foo.test seklinde bir dosya olusmasi yerine foo.test.exe seklinde bir dosya olusmus bulunmakta. Bunun nedenini tam anlamadim ancak bana onemsiz bir detay gibi gozuktu.

//yazdigimiz test fonksiyonunda fonksiyonumuzu yanlizca bir duruma karsi test ettik ama gercek hayatta birden cok test senaryosu ile test etmek isteyebiliriz //@ bunun ornegi icin price folderina bakabilirim

//test table kullanma yontemi birden cok test icin daha uygun bir cozum olabilir Price_test.go paketine bakabilirim

//@ cok onemli bir not : su anda ornek uzerinden ogrendigime gore go dilinde if icerisinde initilalized edebiliyorsun. Ornek kullanim olarak =>
/*
func main() {
    if x, y := 5, 38; x == 5 { //@ goruldugu gibi ilk once init yapiliyor ondan sonra statement kullaniliyor
        fmt.Printf("Whee! %d\n", y)
    }
}
*/

/*
go test komutunu 2 modda calistirmak mumkundur.
-----------local directory mode--------------
$go test seklinde kullanilir yani hicbir sey eklenmez. Go paketi aktif olan directory icerisinde build alir. Modulde bulunan butun unit testler calistirilmaz sadece o pakete ozgu olan testler calistirilir (directory icersinde bulunan)
----------Package List mode -----------------
$go test modulepath/packageName
Bu komut ile birlikte go'ya istedigim yerdeki paket teslerini yapmasini soyleyebilirim hatta path duzgun girilir ise projedeki butun testleri de calistirabilirim. Bu modda iken go otomatik olarak successful olan test sonuclarini cache eder boylelikle multiple defa test etmenin onune gecmis olunur.
bunu test etmek icin go test string seklinde string paketini test edebiliriz.

ilk basta bize verilen cikti su sekilde olur
ok     strings    4.256s

tekrar ayni komutu calistirdigimizda ise bize verilen yeni cikti ise su sekilde olur
ok     strings    (cached)

digeri 4 sn surdu bu aslinda oldukca buyuk bir sure bu ise aninda calisti bunun nedeni cache edilmis olmasindan dolayidir.

eger paketin herhangi bir test dosyasini veya test dosyasini degistirdigimiz zaman durumlarin degisecegini ve eski testin gecersiz kaldigini unutmamak gerekir. Bu nedenle cache uzerine alinmis olan success mesaji gecersiz sayilir bunun onune gecmek icin ise cache alinmasini iptal edebiliriz.

go test strings -count= 1
seklinde yaparsak cache iptal edilmis olur.

go source file icersinde eger environment var kullaniyorsak go sonuclari cache eder (if the env var are not changing)

ornek olarak su anda foo_test.go dosyasina ekledigim yeni TestFoo2 functionu bir kere calisirsa ilk kisimdaki gibi ikinci kere calisirsa ise  su sekilde bir cikti verir
------------------------------------------------------------------------
Running tool: C:\Program Files\Go\bin\go.exe test -timeout 30s -run ^TestFoo2$ newgroupproject/internal/foo

=== RUN   TestFoo2

--- PASS: TestFoo2 (0.00s)
PASS
ok      newgroupproject/internal/foo    0.265s


> Test run finished at 8/2/2022, 11:02:05 AM <

Running tool: C:\Program Files\Go\bin\go.exe test -timeout 30s -run ^TestFoo2$ newgroupproject/internal/foo

=== RUN   TestFoo2

--- PASS: TestFoo2 (0.00s)
PASS
ok      newgroupproject/internal/foo    (cached)
------------------------------------------------------------------
elbette env var degistirilirse go cache edilmis test sonucunu kullanmaz

bu yukarida anlattigim ayni mekanizma open files da da gecerlidir.
*/

//! bazen test sayilari o kadar cok artar ki bunun onune gecebilmek icin testleri concurrent sekilde beraber calistirmak gerekebilir. Bunu nasil yaptigimi gormek icin corge_test.go paketine bakabilirim

//---------------advanced usage of go test command---------------------

/*
command line argumentleri alan bir test yazabiliriz.
//bak-> foo_test.go icerisinde bulunan testargs func neden command line argument almiyor bunu arastir.

=================================FLAGS===================================
You can pass all the existing build flags to the go test command line.

Bilinmesi gereken en onemli flagslardan birisi -cover flagi dir
Bunun nedeni kodumuzun yuzde kacinin test tarafindan covered edildigini ogrenebiliriz

$ go test -cover

PASS
coverage: 100.0% of statements

//not-> sonra bu kisma tekrar geri donecegim basimi asiri derecede agirtti


*/

//@========================================SONUNDA ARRAYLAR KONUSUNA GIRIS YAPTIM BE================

//array ayni ture sahip elementlerin collectionlaridir

//array icerisindeki elements typelarina element type adi verilir

//bu elementler derleme zamaninda bilinen sabit sayida elemana sahiptir (ornek olarak diyelim ki ay isimlerini tutmak istiyorsunuz bu ay isimleri icin string tipinde bir array kullanabiliriz. Ay sayisi ise compile time da belli olur)

/* package main

import "fmt"

func main(){
    var myArray [2]int
    myArray[0] = 156
    myArray[1] = 147
    fmt.Println(myArray)
} */

//! ustteki yonteme ek olarak bu sekilde bir kullanim daha uygundur diyebiliriz

/* package main

import "fmt"

func main(){
    myArray := [2]int{156,147}
    fmt.Println(myArray)
}

*/
//onemli-> ARRAY TANIMLAMALARI
/*
 //not-> long way
var a [2]int
a[0] = 156
a[1] = 147
fmt.Println(a)

//not-> more concise
b := [2]string{"FR", "US"}
fmt.Println(b)

//not-> size computed by the compiler
c := [...]float64{13.2, 37.2}
fmt.Println(c)

//not-> values not set (yet)
d := [2]int{}
fmt.Println(d)
*/

//? bu sekildeki bir kullanimda bos bir array tanimladik ancak dikkat! bu tanimladigimiz array bos olmasina karsin tanimsiz anlamina gelmez cunku bilindigi uzere go dilinde tanimsiz degiskenler tanimlanamaz
/* package main

import "log"

func main() {
    var myA [3]int
    log.Printf("%v", myA)
} */

//ciktisi bu sekildedir : 2022/08/02 14:16:18 [0 0 0]

/*
myEmptyArray2 := [2]string{}
fmt.Println(myEmptyArray2)
// output : [] seklinde bir ciktiya sahip olunur
*/

//len(v) v variablein uzunlugunu return eder.
//cap temelinde len ile ayni islevi gorur

//ACCESS ELEMENT OF ARRAY

/*
package main

import "fmt"

func main() {
    b := [2]string{"FR", "US"}
    firstElement := b[0]
    secondElement := b[1]
    fmt.Println(firstElement, secondElement)
}
*/

//for loop kullanarak array icinde gezmek su sekilde olmaktadir

/* package main

import "fmt"

func main(){
myArray := [13]int{156, 147, 5454, 154,21,2,45,23,154,32,654,32,454}
for index, element := range myArray {
    fmt.Printf("element at index %d is %d\n", index, element)
}
}

*/

//! string uzerinde gezmek icin su sekilde bir yontem uygulayabiliriz

/* package main

import "fmt"

func main(){
myArray := "merhabalar canlarim"
for index, element := range myArray {
    fmt.Printf("element at index %d is %s\n", index, string(element))
}
} */

//bazi durumlarda sadece indexlere veya sadece elementlere ihtiyacimiz vardir. Bunu su sekilde blank identifier kullanarak yapabilirim

/*
for _, element := range myArray {
    fmt.Printf("element is %s\n", string(element))
}


for index, _ := range myArray {
    fmt.Printf("element at index %d\n", index)
}

*/

//bu ustte kullanilan bastan baslayan iteration icin gecerlidir. Eger spesifik olarak bir indexten baslayip digerinde durmak istiyorsam normal for loop kullanmam gerekir.

/*
myArray2 := [2]int{156, 147}
for i := 0; i < len(myArray2); i++ {
    fmt.Printf("element at index %d is %d\n", i, myArray2[i])
}
// output :
//element at index 0 is 156
//element at index 1 is 147
*/

//eger azalan sirada bir seyler yapmak istiyorsam bu sekilde bir kullanim yapabilirm

/*
for i := len(a) - 1; i >= 0; i-- {
    fmt.Println(i, a[i])
}
*/

//@ bir dizinin icerisindeki elementi for loop yardimi ile nasil bulabiliriz

/* func getIndex(haystack [10]int, needle int) int {
    for index, element := range haystack {
        if element == needle {
            return index
        }
    }
    return -1
} */

// bu kullanim aslinda sinirli bir kullanimdir cunku hadi diyelim aradigimiz element arrayin sonunda yer aliyor hepsini tek tek gezmis olacak ayrica bu fonksiyon sadece 10 elemanlik bir arrayi parameter olarak alabiliyor bu nedenlerden dolayi bu fonksiyon yetersizdir.

//arraylari karsilastirmak icin oncelikle su iki ozelligin ayni olmasi gerekmektedir. Bunlardan ilki : ayni tip olmalidirlar ikincisi ise : ayni uzunluga sahip olmalidirlar. Bu karsilastirmada kullanilabilecek operatorler sadece == ve != dur.

/*
a := [2]int{1, 2}
b := [2]int{1, 2}
if a == b {
    print("equal")
} else {
    print("not equal")
}

// output : equal
*/

//@ PEKI BIR ARRAYI FONKSIYONA NASIL PASS EDERIZ? ==> bu onemli bir hatadir bunun nedeni ise bir arrayi fonksiyona gonderdigimiz zaman aslinda o fonksiyonda o arrayin bir kopyasini olusturmus oluruz bu da bazi performans kayiplarina neden olur
//! eger gercekten o arrayi degistirmek istiyor isek pointer kullanmaliyiz

/*

package demo

const NewValue = "changedValue"

func UpdateArray1(array [2]string) {
    array[0] = NewValue
}

-------------------------------
package demo

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestUpdateArray1(t *testing.T) {
    testArray := [2]string{"Value1", "Value2"}
    UpdateArray1(testArray)
    assert.Equal(t, NewValue, testArray[0])
}


------------------test sonucu --------------------

$ go test ./...
--- FAIL: TestUpdateArray1 (0.00s)
    demo_test.go:11:
                Error Trace:    demo_test.go:11
                Error:          Not equal:
                                expected: "changedValue"
                                actual  : "Value1"

                                Diff:
                                --- Expected
                                +++ Actual
                                @@ -1 +1 @@
                                -changedValue
                                +Value1
                Test:           TestUpdateArray1
FAIL
FAIL    maximilien-andile.com/array/copy/demo   0.016s
FAIL


==========================SOLUTIONS======================================
package demo

const NewValue = "changedValue"

func UpdateArray2(array *[2]string) { //@ burada nasil pointer seklinde bir array parametresi alindiginin kullanimi gosterilmis
    array[0] = NewValue
}


----------------------

// array/demo/demo_test.go
package demo

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestUpdateArray2(t *testing.T) {
    testArray := [2]string{"Value1", "Value2"}
    UpdateArray2(&testArray)
    assert.Equal(t, NewValue, testArray[0]) //@ uzun sekilde kullanmak yerine bu sekilde de test yazilabilir
}


===========test ciktisi ==================
$ go test ./...
ok      maximilien-andile.com/array/copy/demo   0.015s

*/

//su sekilde bir kullanim ile array copy edilebilir

/*
   testArray := [2]string{"Value1", "Value2"}
   newCopy := testArray
   -----
   testArray[1] = "updated"
   assert.Equal(t, "updated", newCopy[1])
   bu alt kisim hata verir bunun nedeni sudur : testArray kopyalandiginda onun bir kopyasi olusur sonradan testArrayi degistirmek olusan kopyada herhangi bir degisiklik yasanmasina sebep olmaz
*/

//@ soyle bir kullanimda pointer seklinde array adresini alip kopyaladigimiz icin bir degisiklik yapildiginda o kisimda da yapilmis olur

/*
func TestArrayReference(t *testing.T) {
    testArray := [2]string{"Value1", "Value2"}
    reference := &testArray
    ----------------
    testArray[1] = "updated"
    assert.Equal(t, "updated", reference[1])
}   burada herhangi bir hata meydana gelmeden calisir
*/

//multidimensional arrays
/*
[0,0
 1,3
 4,9]

[3][2] int seklinde tanimlanan bir array yukaridaki gibi bir buyukluge sahiptir

bu arrayde herhangi bir elemente ulasmak icin su sekilde bir kullanim yapilabilir

value := a[2][1]   //9 degeridir


/// ASAGIDAKI DOSYA ICIN INTERNAL/ARRAY KLASORUNE BAKABILIRIM
*/
/* package main
import (
    "newgroupproject/internal/array"
    "log"
)

func main(){
    NewArray:= array.Generate()
    log.Println(NewArray)
} */

//go da multi dimentional arrayler olusturmak mumkundur
/*
package main

import "fmt"

func main() {
    // [[2,4],[5,3]]
    // [[1,2],[9,6]]
    // threeD := [2][2][2]string{} dedigimizde ilk [2] olan yukaridan asagiya array buyuklugu ondan sonra gelen [2] soldan saga olan array ondan sonra gelen [2] ise her bir kutucugun icerisindeki 2 li degiskenler
    threeD := [2][2][2]string{}
    threeD[0][0][0] = "element 0,0,0"
    threeD[0][1][0] = "element 0,1,0"
    threeD[0][0][1] = "element 0,1,1"
    fmt.Println(threeD)
}
*/

//! onemli : oncelikle birden cok dimension olan arraylar visualize etmesi oldukca zor olan seylerdir boyle bir durumda baska bir cozum yolu olabilir mi diye kendi kendime sormam gerekir. Bir arrayin indexindeki elemana ulasmak oldukca hizli bir islemdir ancak bir arrayin icerisindeki bir elemani bulmak oldukca zahmetli bir is olabilir cunku muhtemelen arrayin butun elemanlarini tek tek dolasmam gerekecek. Boyle bir durum varsa map kullanmak daha mantikli bir davranis olacaktir
//arraylar cok iyidir ancak onlarim limitationlari sizelarinin fixed olmasidir.
/*
The length of an array is known at compile time.

In other words, an array once created cannot grow

There is no built-in function to find an element in an array.
*/

//@======================================SLICELAR==============================================
/*
slicelar buyuyebilen arrayler gibi dusunebilirim. Yine ayni tipte olmalari gerekmekte. Buyuyebilmesinin asil nedeni derleme zamaninda fix bir boyut olmamasi ve calisma zamaninda buyutulebilmesinden kaynaklanir.

kullanimi ise su sekildedir []T    T burada slice tipini belirtmek icin kullanilir
/the zero value of slice is nil

*/

//new slice yaratimi

// package main

/* import "fmt"
func main(){
    s:= make([]int,3) //bu sekilde 3 uzunlugunda bir slice olusturdum
    s[0] = 12
    s[2] = 45
    //!su sekilde de bir kullanim yapilabilir
    s2 := []int{10,12}
    fmt.Println(s)
    fmt.Println(s2)
}
*/
//slicing islemi : go dilinde slicing islemi bir seyin bir parcasi anlami tasimaktadir. Ornek olarak bir peynir slice'i tamamen bir peynir degildir ancak yine de peynir parcasidir. bir arrayi bir arrayi isaret eden pointeri ve bir slice i slicing islemine tabi tutabiliriz elbette bunun sonucu yine bir slice olacaktir

// bu islemin su sekilde bir syntaxi vardir

// s := e[low:high]

//ornek olarak
/*
package main

import "fmt"

func main(){
    customers := [4]string{"john haack","asif tunga mubarek","yagmur akan","mavis ve misha"}
    customerSlice:= customers[0:2]
    fmt.Println(customerSlice) //[john haack asif tunga mubarek]
    //goruldugu gibi burada 2 derken cikan sonucun uzunlugu seklinde dusunmem gerekir bunu da soyle aklimda tutabilirim high her zaman bir cikarilacak veya saymaya sifirdan degil de birden baslayacagim
}
*/

//@bir slice kopyalaniyor mu? hayir aslinda orjinal olani pointerlar ile gostermis oluyoruz

/* package main

import "fmt"

func main() {
    customers := [4]string{"John Doe", "Helmuth Verein", "Dany Beril", "Oliver Lump"}
    customersSlice := customers[0:1]
    fmt.Println(customersSlice)
    // modify original array
    customers[0] = "John Doe degistirildi"
    fmt.Println("orijinal arrayin degistirilmesinden sonra bu sekilde gozukur")
    fmt.Println(customersSlice)
    =======
    john doe
    orijinal arrayin degistirilmesinden sonra bu sekilde gozukur
    john doe degistirildi
}
*/

/*
package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
*/
/* output
[John Paul George Ringo]
[John Paul] [Paul George]
[John XXX] [XXX George]
[John XXX George Ringo] */

//! slicing islemi stringlere de uygulanabilir bir islemdir. Ancak unutmamak gerekir bu islemden sonra ortaya cikan sey de bir stringdir.
/*
package main

import "fmt"

func main(){
    otelismi:="ramada mersin"
    s:= otelismi[0:4]
    fmt.Println(s) //output: rama //saymaya birden baslarsam dogru sonuc alirim
}
*/
//go'da bulunan stringler degistirilemezlerdir yani bir kere olusturulduktan sonra orijinal olani degistirmek sonradan slicing islemini atadigimiz degiskenin degerinin degistirilmesine neden olmaz bunun nedeni stringlerin referans seklinde olmamasindan kaynaklanir
//ornek olarak
/*
package main

import "fmt"

func main(){
    otelismi := "ramada mersin"
    s := otelismi[0:4]
    fmt.Println(s)
    otelismi = "mersin ramada"
    fmt.Println(s)
    //--------output
    //rama
    //rama
}

*/

//len hazir bir fonksiyondur ve slice icerisindeki elemanlarin uzunlugunu dondurur.

/* package main

import "fmt"

func main(){
    vatRates := []float32{4.65,4,14,53}
    fmt.Println("slice uzunlugu : ",len(vatRates))  //output => slice uzunlugu :  4
}
*/
/*
slicelar icten olarak aslinda struct yapilaridir ve internally slice bir adet arraya sahiptir ve o arraye bir adet pointer barindirir. Kullanici olarak bu arraya asla erisim gerceklestiremeyiz cunku bu array internaldir.
biz bir arraya slicing islemi uyguladigimizda go o arraya bir pointer atar

bir slice'in internal represantionu su sekilde olabilir
    pointer to array
    length
    capacity



capacity su demektir: underlying arrayde allocated edilen element sayisini tarif eder. Bunu anlamak icin bir ornege bakalim
*/

/* package main

import "fmt"


func main(){
    names := [4]string{"tunga","ozgul","sila","sena"}
    mySlice := names[1:3]
    fmt.Println(mySlice)
    fmt.Println("length: ",len(mySlice))
    fmt.Println("capacity: ",cap(mySlice))
    fmt.append
    /*
    output
    [ozgul sila]
    length:  2
    capacity: 3
*/
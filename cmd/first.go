/*
go dilinde variablelar nasil tanimlanir?
degiskenisimlerine identifier denir degerlerine ise expression denir

var degiskenIsmi degiskenturu = deger   seklinde degisken tanimlamasi yapilabilir
var degisken1,degisken2,degisken3 degiskenlerinturu = deger1,deger2,deger3
=============kisa degisken tanimlamalari==================================
kisa degisken tanimlamalari su sekilde yapilir
degiskenismi := deger

!   kisa degisken tanimlamalari fonksiyonlarin disarisinda kullanilamazlar

============sabit degisken tanimlamak=====================================

const degiskenIsmi degiskenturu = deger

tipsiz bir const tanimlamasi yapilabilir
    const degiskenismi = deger  seklinde

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

//* _ kullanimi bu kismi es gec anlamina gelmektedir. Ornekte 3 adet return degeri var, diyelim ki benim sadece iclerinden bir tanesine ihtiyacim var bu durumda yapmam gereken _ kullanmaktir bu kisaca diger degiskenleri ignore et anlamina gelir ayni zamanda import ettigim bir seyi su an kullanmayacak isem bas kismina _ ekleyebilirim. Ayni sekilde degisken tanimlayip o degiskeni kullanmadigim durumlarda da bunu uygulayabilirim

/* package main

import (
	"fmt"
	new "newgroupproject/internal/new"
    _ "newgroupproject/internal/email"

)

func main(){
	l1:=23
	new.Deneme()
	fmt.Println("hellooo ",l1)
} */

//boyle bir kullanimda l1:=23 fonc disinda kullanildigi icin func icinden buna ulasamiyoruz ancak soyle bir durum var eger ben bu kullanim yerine su sekilde bir kullanim yaparsam var l1 = 23 or var l1 int = 23 seklinde bu durumda func icinden bu degiskene erisebilirim herhangi bir sikinti yasamam

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
    emailContents := email.GetEmailContents("M", customerName, nights)
    fmt.Println(emailContents)
    email.SendEmail(emailContents, customerEmail)
    invoice.CreateAndSaveInvoice(customerName, nights, 145.32)
} */

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
 var nameDisplayer func(name, firstname string) string bu sekilde bir kullanim mevcut bu kullanimin amaci ise sudur

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

bu kisimda bos biraktigimiz capitalcity kismi bos kalir. String bir ifade icin bu "" seklindedir



empty := Country{}

ayni sekilde bos bir belge de olusturabiliriz



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

// methodlar receiveri olan fonksiyonlara verilen isimlerdir

// package main

// import "fmt"

// func main() {
/*
   ! var po *int = &i seklinde pointer tanimlamasi yapabilecegim gibi
   ! po := &i seklinde de pointer tanimlamasi yapabilirim
*/
/*

   i:=1
   po := &i
   po2 := &po
   fmt.Println("i nin degeri: ",i)
   fmt.Println("i nin adresi: ",&i)
   fmt.Println("pointer degeri: ",po)
   fmt.Println("Pointerin adresi: ", &po)
   fmt.Println("Pointerin isaret ettigi adresteki degeri: ", *po)
   fmt.Println("Pointerin adresini tutan pointerin degeri: ",po2)
   fmt.Println("Pointerin adresini tutan pointerin adresi: ",&po2)
   fmt.Println("Pointerin adresini tutan pointerin isaret ettigi adresteki degeri: ",*po2)
   fmt.Println("Pointerin adresini tutan pointerin isaret ettigi adresteki pointerin isaret ettigi adresteki degeri: ",**po2)
   ========cikti============
   i nin degeri:  1
   i nin adresi:  0xc00008a040
   pointer degeri:  0xc00008a040
   Pointerin adresi:  0xc00009e038
   Pointerin isaret ettigi adresteki degeri:  1
   Pointerin adresini tutan pointerin degeri:  0xc00009e038
   Pointerin adresini tutan pointerin adresi:  0xc00009e040
   Pointerin adresini tutan pointerin isaret ettigi adresteki degeri:  0xc00008a040
   Pointerin adresini tutan pointerin isaret ettigi adresteki pointerin isaret ettigi adresteki degeri:  1

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
bu kodun soyle bir mantigi var aslinda. internal klasoru icersinde cart isimli bir klasorum var o klasorun icerisinde cart.go isimli bir paketim var. Bu paketi import kisminda import etmisim. Daha sonra bu paketimin icerisinde Cart isimli bir struct'im ve iki adet metodum var. Metodlarimin receiveleri cart tipinde oldugundan dolayi cart tipinde bir sey olmadan kullanamam bundan dolayi metotlarimi kullanabilmek icin newcart isimli cart tipindeki cart.Cart{} olusturdum, daha sonra bunu kullanarak metotlarima erisebildim
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
    b := &receiver.Fatura{
        Isim: "Market Faturasi",
        Bahsis: 40.00,
        Items: map[string]float32{"patates": 5.00, "elma":4.00, "biber":10.00, "salatalik":12.00},
    }

    b.Update("patlican",20.00)
    b.Update("domates",35.00)
    b.Update("lahana",17.00)
    fmt.Println(b.Hesap())
} */


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

cartDeref := *cartPtr  //! iste bu sekilde su an pointerin isaret ettigi yerdeki value lara ulasis oldumm

@ & -> adres bilgisini al (reference)
@ * -> adres bilgisine ulas (takip et) (dereference)
* operatoru ile *card seklindeki kullanim birbirinin aynisi degildir. *card bir isaretci tipi olabilecegi gibi ayni zamanda referansi kaldirilmis bir isaretci degiskenini de gosteriyor olabilir. Kullanimini yakindan inceleyerek hangisi olduguna karar vermem gerekir


!     var myPointerVar *int
!    fmt.Println(*myPointerVar)  mesela soyle bir kullanim yanlistir bunun nedenini detayli aciklamak istiyorum

@ oncelikle tipi(adresini gosterecegi) *int olan bir pointer degiskeni tanimladik. Daha sonra bunu dereference etmeye calistik ancak oncesinde referans etmemistik ki.. yani bu sekilde sanki gostermedigi bir yer oldugu halde o yere gitmeye calisiyoruz aslinda o yer nil dir. Yani yoktur bu nedenle program hata verir. Bu programin calismasinin tek yolu degisken tanimlamasindan sonra reference atamasi yapilmasidir. mypointer = &blabla seklinde. Sonra buradaki kullanimi yapabilirim


map ve slice lar reference tipleridir. Yani kendi ic yapilarina reference lilardir. Bu soyle bir kullanim saglar, bir method parametre olarak bir maps aliyor ise ve bu parametre pointer seklinde olmasa bile yapilan degisiklikler direkt olarak mapse yansir.
*/




package main

import (
    "fmt"
    _"strings"

)

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
bu kopya hala temel olan slice i referans ediyor ancak soyle bir durum var //! sadece ama sadece var olan degerleri degistririsek bu olur
@ aklimda daha kolay kalmasi icin sanki icerisindeki her bir value orijinal olan slice in valuelarina referans gosteriyormuslar gibi dusunebilirim
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
                            Pointer receiver kullan
struct yapim cok buyuk ve agir ise
receiveri degistirmek istiyorsam (ornegin struct alaninda bir degiskenin isim alanini degistirmeye calisiyorsam)
synchronization primitive (like sync.Mutex) gibi bir sey kullaniyorsak

                            Value receiver kullan
struct yapim small ise
receiveri degistirmeye gerek yok ise
eger receiver map, func, chan, slice, string, interface ise bunlarla value receiver kullanmam gerekir cunku bu yapilar yapisal olarak hali hazirda pointerlar ile calisirlar
*/
//@=========================================interface kavrami==================================================

/* 
arayuz dedigimiz sey aslinda bir dizi ortak davranisi tanimlayan bir seydir. 
herhangi bir implementation olmadan (ornek olarak metotlarin davranislarini tanimlar)
kisaca interface herhangi bir implementation olmadan methodlarin nasil bir tipte oldugunu tanimlar
*/

//ornek bir interface imlementationu bu sekilde gerceklesir 

/* 
type DomasticAnimal interface{
    ReceiverAffection(from Human)
    GiveAffection(to Human)
}

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


func Pet(animal DomasticAnimal, human Human){
    animal.GiveAffection(human)
    animal.ReceiverAffection(human)
}


func main(){
    Pet(Dog{"TOM"},Human{"asif","tunga",23,"turkiye"})
} 

*/

// ================================ anlamama yardimci olacak baska bir ornek olabilir ===================================================


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



type Stringer interface{
    String() string             //@ burada yaptigim tek sey Stringer isminde String() isimli string geri donuslu bir metot koydum. Interface tanimlamasi bu kadar
}

func main() {
    a := Article{
        Title: "Understanding Interfaces in Go",
        Author: "Sammy Shark",
    }
    b:=Article{"Umutlu Bir Bilgisayar Muhendisinin Yasami","Asif Tunga Mubarek"} // ayni zamanda bu sekilde de kullaniliyor (daha fazla bilgi icin yukaridaki struct yapisina bakabilirim)
    c:=Book{"Yasamim","Asif Tunga Mubarek",23}
    Print(a)
    Print(b)
    Print(c)
}


func Print(s Stringer) {
    fmt.Println(s.String())
}
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

asagidaki print fonksiyonunu da degistirmemis oldugumu varsayalim

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


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















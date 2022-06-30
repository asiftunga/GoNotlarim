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



package main

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
}

/* // send an email
func sendEmail(contents string, to string) {
    // ...
    // ...
}

// prepare email template
func getEmailContents(title string, name string, nights uint) string {
    text := "Dear %s %s,\n your room reservation for %d night(s) is confirmed. Have a nice day !"
    return fmt.Sprintf(text,
        title,
        name,
        nights)
} */

/* // create the invoice for the reservation
func createAndSaveInvoice(name string, nights uint, price float32) {
    // ...
} */

//! ayni temada olan fonksiyonlari birbirleri ile birlestirip onlari birer paket haline getirmek iyi bir aliskanliktir
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


package main

import (
	"fmt"
	deneme "newgroupproject/new"
)

func main(){
	l1:=23
	deneme.Deneme()
	fmt.Println("hellooo ",l1)
}

//boyle bir kullanimda l1:=23 fonc disinda kullanildigi icin func icinden buna ulasamiyoruz ancak soyle bir durum var eger ben bu kullanim yerine su sekilde bir kullanim yaparsam var l1 = 23 or var l1 int = 23 seklinde bu durumda func icinden bu degiskene erisebilirim herhangi bir sikinti yasamam




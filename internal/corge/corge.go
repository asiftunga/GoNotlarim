// corge/corge.go
package corge

import "fmt"

// func Corge() string {
//     return "Corge"
// }

//simdi bu kodumuzu patch edelim => patch demek aslinda su demektir ; API da herhangi bir degisiklik yapmaz. API (modul tarafindan disariya aktarilan her sey demektir)

func Corge() string {
	return fmt.Sprintf("Corge")
}

//su anda biz aslinda kodumuzu patch etmis olduk
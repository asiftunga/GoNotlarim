package time

import "time"

func WhatTimeIsIt() string{
    return time.Now().Format(time.RFC3339)
}


/* 
package gomodule

import "time"

func WhatTimeIsIt(format string) string {
    return time.Now().Format(format)
}
*/
//simdi diyelim ki bizim bu fonksiyonumuzu kullanan developerlar time formatini kendileri secmek istediler ve bu gelen istekler dogrultusunda biz de fonksiyonumuzu revize ettik ve yeni halini ortaya cikarttik. Simdi bu durumdan dolayi bizim aslinda fonksiyonumuzun imzasi degismis oldu, sasirtmayacak sekilde fonksiyonumuzu onceden implement etmis olan kodlar calismayacaktir. Bu nedenle yeni bir major versiyon yaratilir
//Versiyon bir kere yaratildi mi artik degistirilemez yani versiyonu ortaya cikartip sonra bir hata ile karsilasilinca hatayi giderip sonra o versiyonu silip ayni versiyon ismi ile cikartilamaz


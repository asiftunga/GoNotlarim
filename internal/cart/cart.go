package cart

import "time"

type Cart struct{
	ID        string
    CreatedAt time.Time
    UpdatedAt time.Time
    lockedAt  time.Time
    CurrencyCode string
    isLocked     bool
}

func (c *Cart) TotalPrice() (error){
    //buradaki c *Cart olan kisim bu methodun receiveridir.
    //receiverlarin normal parametrelerden farki aslinda sunlardir; 
    //bu fonksiyonu kullanirken sadece cart tipinde kullanabiliriz baska bir sekilde kullanamayiz
    return nil
}


func (c *Cart) Lock() error {
    // ...
    return nil
}
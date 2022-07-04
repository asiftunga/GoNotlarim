// methods/example-project/cart/cart.go
package cart

import (
	"errors"
	"github.com/Rhymond/go-money"
	"newgroupproject/internal/product"
	"os/user"
	"time"
)

type Cart struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	lockedAt  time.Time
	user.User
	Items        []Item
	CurrencyCode string
	isLocked     bool
}

type Item struct {
	product.Product
	Quantity uint8
}

func (c *Cart) TotalPrice() (*money.Money, error) {
	//buradaki c *Cart olan kisim bu methodun receiveridir.
	//receiverlarin normal parametrelerden farki aslinda sunlardir;
	//! bu fonksiyonu kullanirken sadece cart tipinde kullanabiliriz baska bir sekilde kullanamayiz
    total := money.New(0, c.CurrencyCode)
	var err error
    for _, v := range c.Items {
        itemSubtotal := v.Product.Price.Multiply(int64(v.Quantity))
        total, err = total.Add(itemSubtotal)
        if err != nil {
            return nil, err
        }
    }
    return total, nil
}

func (c *Cart) Lock() error {
	if c.isLocked {
		return errors.New("kart zaten coktan kilitlenmis")
	}
	c.isLocked = true       //* iste bu sekilde kullaniyorum
	c.lockedAt = time.Now() //kilitlenme zamanini girdim
	return nil
}

func (c *Cart) delete() error {
	// to implement
	return nil
}

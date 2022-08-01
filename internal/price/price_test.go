package price

import "testing"

func Test_totalPrice1(t *testing.T) {
    // test case 1
    var expected uint = 0
    actual := totalPrice(0, 150, 12)
    if expected != actual {
        t.Errorf("Expected %d does not match actual %d", expected, actual)
    }
    // test case 2
    expected = uint(112)
    actual = totalPrice(1, 100, 12)
    if expected != actual {
        t.Errorf("Expected %d does not match actual %d", expected, actual)
    }

    // test case 3
    expected = uint(224)
    actual = totalPrice(2, 100, 12)
    if expected != actual {
        t.Errorf("Expected %d does not match actual %d", expected, actual)
    }

}
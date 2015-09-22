// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9.
// The sum of these multiples is 23.
// FIND THE SUM OF ALL THE MULTIPLES OF 3 OR 5 BELOW 1000.

package projeuler

import (
    "fmt"
    "errors"
    "math"
)

func findIntsLessThanAndDivisibleBy(lessThan int, multiples... int) (int, error) {
    encountered := map[int]bool{}
    list := []int{}
    for _, num := range multiples {
        if num <= 0 {
            return int(math.NaN()), errors.New("Argument out of bounds")
        }

        for i :=0; i < lessThan; i += num {
            if encountered[i] == true {
            } else {
                encountered[i] = true
                list = append(list, i)
            }
        }
    }

    sum := 0
    for _, n := range list {
        sum += n
    }
    return sum, nil
}

func main() {
    r, e := findIntsLessThanAndDivisibleBy(1000, 3, 5)
    if e != nil {
        fmt.Println(e)
    } else {
        fmt.Println(r)
    }
}


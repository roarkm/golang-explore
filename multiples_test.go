package euler

import (
    "testing"
)


func TestShouldSolveVerbatim(t *testing.T) {
    // If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9.
    // The sum of these multiples is 23.

    // FIND THE SUM OF ALL THE MULTIPLES OF 3 OR 5 BELOW 1000.
    n, e := sumMultiplesLessThan(1000, 3, 5)
    r := 233168
    if n != r {
        t.Errorf("Result should equal %d but was %d", r, n)
    }
    if e != nil {
        t.Errorf("Error raised when there should be none.")
    }
}

func TestShouldSetErrorForArsOutOfBounds(t * testing.T) {
    _, e := sumMultiplesLessThan(1000, -3, 5)
    if e == nil {
        t.Errorf("No error was set when a multiple was set to a negative.")
    }

    _, e = sumMultiplesLessThan(-1000, 3, 5)
    if e == nil {
        t.Errorf("No error was set when the cap was set to a negative.")
    }

    _, e = sumMultiplesLessThan(100, 101, 5)
    if e == nil {
        t.Errorf("No error was set when a multiple was greater than the cap.")
    }
}


func TestShouldFindOccurencesOfMultiples(t * testing.T) {
    n, _ := sumMultiplesLessThan(10, 3)
    r := 18
    if n != r {
        t.Errorf("Result should equal %d but was %d", r, n)
    }

    n, _ = sumMultiplesLessThan(100, 9)
    r = 594
    if n != r {
        t.Errorf("Result should equal %d but was %d", r, n)
    }

    n, _ = sumMultiplesLessThan(99, 9)
    r = 495
    if n != r {
        t.Errorf("Result should equal %d but was %d", r, n)
    }

    n, _ = sumMultiplesLessThan(100, 9, 4, 23)
    r = 1824
    if n != r {
        t.Errorf("Result should equal %d but was %d", r, n)
    }
}

package main

func CountV1(pile []int, plafond int) []int {
    slice := make([]int, plafond)
    for _, nom := range pile  {
        if nom <= plafond {
            slice[nom-1]++}
    }
    return slice
}
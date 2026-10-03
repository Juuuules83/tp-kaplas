package main

import (
	"fmt"
	"testing"
)

var sink int // garde le résultat pour que le compilateur ne supprime pas l'appel
func BenchmarkSmallest(b *testing.B) {
	for _, n := range []int{1_000, 10_000, 100_000} {
		pile := Shuffled(n) // préparation hors de la mesure
		b.Run(fmt.Sprintf("V1/n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sink = SmallestV1(pile)
			}
		})
	}
}
func BenchmarkDuplicateV1(b *testing.B) {
	for _, n := range []int {1_000, 10_000, 100_000} {
		pile := WithDuplicate(n, 3) // préparation hors de la mesure
		b.Run(fmt.Sprintf("V1/n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
						sink = DuplicateV1(pile)
			}
		})
	}
}
func BenchmarkDuplicateV2(b *testing.B) {
	for _, n := range []int {1_000, 10_000, 100_000} {
		pile := WithDuplicate(n, 3) // préparation hors de la mesure
		b.Run(fmt.Sprintf("V2/n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
						sink = DuplicateV2(pile)
			}
		})
	}
}
func BenchmarkTowerHeightV1(b *testing.B){
	for _, n := range []int{1_000, 10_000, 100_000}{
		b.Run(fmt.Sprintf("V1/n=%d", n),func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sink = TowerHeightV1(n)
			}
		})
	}
}
		func BenchmarkTowerHeightV2(b *testing.B){
	for _, n := range []int{1_000, 10_000, 100_000}{
		b.Run(fmt.Sprintf("V2/n=%d", n),func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sink = TowerHeightV2(n)
			}
		})
	}
}

func BenchmarkCountV1(b *testing.B) {
    for _, n := range []int{1_000, 10_000, 100_000} {
        pile := Random(n, n)
        b.Run(fmt.Sprintf("V1/n=%d", n), func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                sink = len (CountV1(pile, n)) 
            }
        })
    }
}
package main

func DuplicateV1(pile []int) int {

mapbool := make(map[int]bool)

for _, nom := range pile {
	mapbool[nom] = true
	}
return len(mapbool)
}
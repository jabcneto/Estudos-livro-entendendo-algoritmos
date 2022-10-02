package main

func main() {
	escolhido := 32
	list := []int{1, 2, 3, 9, 11, 32, 59, 102, 302, 888}
	min := 0
	max := len(list) - 1
	enquanto := min != max
	for enquanto {
		enquanto = min != max
		meio := (max + min) / 2
		
		if list[meio] == escolhido {
			print(list[meio])
			break
		} else if escolhido > list[meio] {
			min = meio + 1
		} else if escolhido < list[meio] {
			max = meio - 1
		}

		if min == max {
			print("nao encontrado")
		}
	}

}

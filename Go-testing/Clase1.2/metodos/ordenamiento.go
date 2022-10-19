package metodos

func insercionSort(numbers *[]int) {
	var j, aux int
	n := *numbers
	for i := 1; i < len(n); i++ {
		aux = n[i]
		j = i - 1
		for (j >= 0) && (aux < n[j]) {
			n[j+1] = n[j]
			j--
		}
		n[j+1] = aux
	}
}

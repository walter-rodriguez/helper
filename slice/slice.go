package slice

// Busca la posicion de un slice de cualquier tipo con la función de comparación `validate`
// Retorna la posicion del slice si existe, si no existe o no se encuentra devuelve -1
func SliceIndex(limit int, validate func(int) bool) int {
	for i := 0; i < limit; i++ {
		if validate(i) {
			return i
		}
	}

	return -1
}

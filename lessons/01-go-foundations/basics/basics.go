package basics

// Sum recebe dois inteiros e devolve a soma entre eles.
func Sum(left int, right int) int {
	count := 0
	count = left + right

	return count
}

// Greeting recebe um nome e devolve uma saudação.
func Greeting(name string) string {
	return "Olá, " + name
}

// RectangleArea recebe as dimensões de um retângulo e devolve sua área.
func RectangleArea(width float64, height float64) float64 {
	area := 0.0

	area = width * height

	return area
}

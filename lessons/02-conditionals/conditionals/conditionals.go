package conditionals

// IsAdult informa se age atingiu a maioridade definida pelo exercício.
func IsAdult(age int) bool {
	return age >= 18
}

// Larger devolve o maior entre left e right.
func Larger(left int, right int) int {
	if left > right {
		return left
	}

	return right
}

// CanAccess exige simultaneamente maioridade e uma conta ativa.
func CanAccess(age int, accountActive bool) bool {

	isAdult := IsAdult(age)

	return isAdult && accountActive
}

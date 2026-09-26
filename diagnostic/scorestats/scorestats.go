package scorestats

// Summary reúne as estatísticas calculadas para um conjunto válido de notas.
type Summary struct {
	Count   int
	Total   int
	Average float64
	Min     int
	Max     int
}

// Analyze valida scores e devolve suas estatísticas.
// A implementação faz parte do diagnóstico inicial.
func Analyze(scores []int) (Summary, error) {
	panic("TODO: implemente Analyze")
}

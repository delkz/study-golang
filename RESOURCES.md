# Fontes de aprendizagem

Este arquivo registra somente fontes efetivamente usadas nas sessões. Para conceitos de Go, priorize documentação, especificação, pacotes e publicações oficiais. Cada inclusão deve informar o tema sustentado e a data de consulta, evitando listas extensas sem relação com a prática realizada.

## Fontes utilizadas

### Preparação do ambiente — 2026-09-26

- [Download and install](https://go.dev/doc/install) — instalação e verificação do toolchain de Go.
- [Tutorial: Get started with Go](https://go.dev/doc/tutorial/getting-started) — referência inicial sobre módulos e comandos básicos do toolchain.

Essas fontes sustentam apenas a preparação descrita no `README.md`; ainda não constituem evidência de aprendizagem.

### Diagnóstico e lição 1 — 2026-09-26

- [Go Modules Reference](https://go.dev/ref/mod) — relação entre módulos, pacotes e caminhos.
- [The Go Programming Language Specification](https://go.dev/ref/spec) — pacotes, valores zero e execução de programas.
- [Add a test](https://go.dev/doc/tutorial/add-a-test) — convenções de arquivos e funções de teste.
- [Errors are values](https://go.dev/blog/errors-are-values) — modelo de erros como valores programáveis.

As fontes sustentam o material preparado. O estado das competências só será atualizado depois da prática e da revisão.

### Lição 1 adaptada — 2026-09-26

- [Slice types — Go specification](https://go.dev/ref/spec#Slice_types) — definição de slice, comprimento, capacidade e armazenamento compartilhado.
- [Appending to and copying slices — Go specification](https://go.dev/ref/spec#Appending_and_copying_slices) — comportamento normativo de `append` e `copy`.
- [Go slices: usage and internals](https://go.dev/blog/slices-intro) — modelo visual do descritor e do array subjacente.
- [Package builtin](https://pkg.go.dev/builtin) — contratos das funções predefinidas usadas na prática.

Esse material sobre slices foi adiado para uma lição futura após a revisão do ponto de partida.

### Lição 1 — fundamentos de funções e variáveis — 2026-09-26

- [Function declarations — Go specification](https://go.dev/ref/spec#Function_declarations) — estrutura de assinaturas, parâmetros, resultados e corpo.
- [Short variable declarations — Go specification](https://go.dev/ref/spec#Short_variable_declarations) — regras de `:=` dentro de funções.
- [Variables with initializers — A Tour of Go](https://go.dev/tour/basics/9) — declaração com `var` e inferência por inicializador.
- [Type inference — A Tour of Go](https://go.dev/tour/basics/14) — inferência de tipos em declarações locais.

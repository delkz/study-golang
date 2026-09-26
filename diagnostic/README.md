# Diagnóstico inicial de Go

Este diagnóstico serve para escolher o ponto de partida da trilha. Ele não vale nota e não pressupõe conhecimento prévio de Go. Responder `não sei` é mais útil do que pesquisar uma resposta, porque o objetivo é registrar o que você consegue mobilizar agora.

Tempo sugerido: 30 a 45 minutos. Não consulte a primeira lição antes de terminar esta atividade.

## Parte 1 — modelo mental

Responda com suas próprias palavras, de forma curta:

1. O que você entende por linguagem compilada e estaticamente tipada? Como isso difere do fluxo normal de TypeScript executado com Node.js?
2. O que você espera que sejam um módulo e um pacote em Go? Se não souber, explique como imagina que se relacionem.
3. Como você imagina que Go representa e trata falhas sem depender do fluxo tradicional de `try/catch`?
4. Qual seria, na sua opinião, a diferença entre um array e uma slice em Go?
5. Antes desta sessão, o que você já fez em Go? Inclua cursos, exemplos, projetos, leitura ou informe que está começando do zero.

Registre as respostas em [RESPONSES.md](RESPONSES.md).

## Parte 2 — leitura de código

Sem executar o trecho abaixo, escreva em `RESPONSES.md`:

- o que você acredita que será impresso;
- por que a alteração feita por `alias[0] = 10` pode ou não aparecer em `numbers`;
- por que o `append` pode ser relevante para essa relação.

```go
package main

import "fmt"

func main() {
	numbers := []int{2, 4}
	alias := numbers

	alias[0] = 10
	alias = append(alias, 8)

	fmt.Println(numbers)
	fmt.Println(alias)
}
```

Depois de registrar sua previsão, você pode executar uma cópia do trecho para conferir. Não altere sua resposta original; acrescente o resultado observado e explique qualquer diferença.

## Parte 3 — implementação curta

Implemente `Analyze` em [scorestats/scorestats.go](scorestats/scorestats.go).

Se ainda não souber como iniciar, não pesquise uma solução: deixe o scaffolding intacto e registre essa dificuldade ao entregar. A ausência de implementação também é uma evidência válida para escolher o ponto de partida.

### Cenário

Um serviço recebe notas inteiras de `0` a `100` e precisa produzir um resumo para outra camada da aplicação.

### Requisitos

- Para uma entrada válida, retorne quantidade, soma, média, menor nota e maior nota.
- A média deve usar divisão decimal, não divisão inteira.
- Uma lista vazia deve retornar erro.
- Qualquer nota menor que `0` ou maior que `100` deve retornar erro.
- Em caso de erro, retorne o valor zero de `Summary`.
- Não altere a slice recebida.
- Use apenas a biblioteca padrão.

### Critérios de aceitação

Na raiz do repositório, execute:

```powershell
go test ./diagnostic/scorestats -v
go fmt ./...
go vet ./...
```

Os testes devem passar e as ferramentas não devem reportar problemas. Antes da implementação, `go test` falhará de propósito porque a função contém um `panic` de scaffolding.

### O que será observado

- leitura de assinatura e tipos;
- controle de fluxo e iteração;
- uso de slices e structs;
- conversão necessária para calcular a média;
- tratamento explícito de entradas inválidas;
- capacidade de ler falhas de teste e ajustar a solução;
- clareza do código e dos nomes.

Não adicione abstrações, dependências ou otimizações antecipadas. O objetivo é observar os fundamentos.

## Entrega

Quando terminar, diga que concluiu o diagnóstico. Eu revisarei `RESPONSES.md`, a implementação e os resultados das ferramentas antes de atualizar seu mapa de competências e liberar ou adaptar a primeira lição.

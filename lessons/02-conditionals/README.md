# Lição 2 — booleanos, comparações e condicionais

Status: concluída.

## Objetivo

Ao final desta lição, você deverá conseguir:

- declarar e retornar valores `bool`;
- produzir booleanos com comparações;
- combinar condições com `&&`, `||` e `!`;
- controlar o fluxo de uma função com `if` e `else`;
- reconhecer casos de fronteira, como uma idade exatamente igual ao limite.

Continuaremos usando apenas funções, parâmetros, variáveis e retornos. Ainda não entraremos em loops, slices, structs ou erros.

## 1. O tipo `bool`

Um `bool` possui somente dois valores possíveis: `true` ou `false`.

Uma comparação já produz um booleano:

```go
temperature := 30
isHot := temperature > 25
```

Nesse exemplo, `isHot` recebe `true`. Não é necessário escrever um `if` apenas para transformar uma comparação em `true` ou `false`.

Os principais operadores de comparação são:

| Operador | Significado |
| --- | --- |
| `==` | igual |
| `!=` | diferente |
| `<` | menor |
| `<=` | menor ou igual |
| `>` | maior |
| `>=` | maior ou igual |

Assim como em TypeScript, comparação usa `==`, mas Go não possui a coerção implícita de tipos associada ao `==` do JavaScript. Os operandos precisam ter tipos compatíveis.

## 2. Combinando condições

Operadores lógicos combinam ou invertem valores booleanos:

- `left && right` é verdadeiro somente quando os dois lados são verdadeiros;
- `left || right` é verdadeiro quando pelo menos um lado é verdadeiro;
- `!value` inverte o booleano.

Exemplo isolado:

```go
hasAccess := accountActive && age >= 18
```

Não escreva `accountActive == true` quando o próprio booleano já expressa a condição.

## 3. Decisões com `if`

`if` executa um bloco quando sua condição é verdadeira:

```go
func Absolute(value int) int {
	if value < 0 {
		return -value
	}

	return value
}
```

Go não usa parênteses em volta da condição, mas exige chaves no bloco.

Como o primeiro ramo termina com `return`, não precisamos escrever `else`. Se a condição for verdadeira, a função já termina; caso contrário, a execução segue naturalmente.

`else` continua disponível quando as duas ramificações precisam realizar trabalho antes de a função continuar:

```go
if condition {
	// caminho verdadeiro
} else {
	// caminho falso
}
```

## Exercício

Implemente as três funções em [conditionals.go](conditionals/conditionals.go):

1. `IsAdult`: devolve `true` para idade igual ou superior a `18`.
2. `Larger`: devolve o maior entre dois inteiros; se forem iguais, pode devolver qualquer um deles.
3. `CanAccess`: devolve `true` somente quando a conta estiver ativa e a idade for igual ou superior a `18`.

Restrições:

- mantenha as assinaturas existentes;
- use uma comparação direta em `IsAdult`;
- use `if` em `Larger`;
- use o operador `&&` em `CanAccess`;
- não altere os testes.

## Executando

Trabalhe em uma função por vez:

```powershell
go test ./lessons/02-conditionals/conditionals -v
```

Quando todos os testes passarem:

```powershell
go fmt ./lessons/02-conditionals/conditionals
go vet ./lessons/02-conditionals/conditionals
```

## Checkpoint

Depois de concluir, responda:

1. Qual é o resultado de `18 >= 18` e por que essa fronteira importa em `IsAdult`?
true, importante pois 18 anos tambem é adulto, nao só apartir dos 19, seria a mesma coisa de um age > 17
2. Por que `return age >= 18` já é suficiente para devolver um `bool`?
pois o retorno é uma booleana
3. Qual é a diferença entre `&&` e `||`?
and & or
4. Por que um `else` é desnecessário quando o bloco do `if` já termina com `return`?
se a condição for verdadeira a função ja vai terminar, entao o else nao faz sentido aqui

## Fontes primárias

- [Operators — Go specification](https://go.dev/ref/spec#Operators)
- [Comparison operators — Go specification](https://go.dev/ref/spec#Comparison_operators)
- [Logical operators — Go specification](https://go.dev/ref/spec#Logical_operators)
- [If statements — Go specification](https://go.dev/ref/spec#If_statements)

# Lição 1 — funções, parâmetros e variáveis

Status: liberada.

## Objetivo

Ao final desta lição, você deverá conseguir:

- reconhecer as partes de uma função em Go;
- receber valores por parâmetros;
- declarar e atribuir variáveis locais;
- entender inferência de tipo com `:=`;
- retornar um valor;
- chamar funções por meio de testes prontos.

Nesta etapa não usaremos slices, loops, structs nem tratamento de erros. Esses assuntos voltarão quando a base necessária estiver firme.

## 1. Estrutura de uma função

Observe este exemplo isolado:

```go
func Double(value int) int {
	result := value * 2
	return result
}
```

Leia a primeira linha da esquerda para a direita:

- `func` declara uma função;
- `Double` é o nome da função;
- `value` é o nome do parâmetro recebido;
- `int` dentro dos parênteses é o tipo do parâmetro;
- o último `int` é o tipo do valor devolvido.

O corpo da função fica entre `{` e `}`. Como a assinatura promete devolver um `int`, o fluxo normal da função precisa terminar devolvendo um inteiro com `return`.

Em TypeScript, uma assinatura semelhante seria `function double(value: number): number`. A diferença importante é que `int` é um tipo inteiro específico de Go; ele não representa todos os números como o tipo `number` do JavaScript.

## 2. Parâmetros são variáveis locais

Quando alguém chama `Double(5)`, o valor `5` é associado ao parâmetro local `value`. A função usa esse valor durante sua execução.

Uma função pode receber mais de um parâmetro:

```go
func Subtract(left int, right int) int {
	return left - right
}
```

Cada chamada recebe seus próprios valores. `Subtract(10, 3)` e `Subtract(8, 2)` executam o mesmo comportamento com entradas diferentes.

## 3. Declarando variáveis

Dentro de uma função, estas formas são válidas:

```go
var count int
var name string = "David"
message := "Olá"
```

- `var count int` cria um `int` sem valor explícito; ele começa com o valor zero de seu tipo, que é `0`.
- `var name string = "David"` informa o tipo e o valor inicial.
- `message := "Olá"` pede que Go infira o tipo a partir do valor; nesse caso, `string`.

`:=` declara uma variável nova e só pode ser usado dentro de funções. Depois que ela existe, uma nova atribuição usa apenas `=`:

```go
message := "Olá"
message = "Bem-vindo"
```

Go não permite declarar uma variável local e simplesmente nunca usá-la. Essa regra ajuda a manter o código livre de sobras acidentais.

## 4. Expressões e retorno

Uma expressão produz um valor. Alguns exemplos:

```go
left + right
width * height
"Olá, " + name
```

Você pode guardar o resultado em uma variável e retorná-la, como no exemplo `Double`, ou retornar diretamente. Nesta lição, prefira criar uma variável local antes do `return`; isso deixa visível o ciclo receber → calcular → devolver que estamos praticando.

## Exercício

O arquivo [basics.go](basics/basics.go) contém três funções incompletas:

1. `Sum`: recebe dois inteiros e devolve a soma.
2. `Greeting`: recebe um nome e devolve `"Olá, "` seguido do nome.
3. `RectangleArea`: recebe largura e altura e devolve a área.

Para cada função:

1. remova o `panic`;
2. crie uma variável local com `:=` para guardar o resultado;
3. devolva essa variável com `return`.

Não altere nomes, parâmetros, tipos ou testes.

## Executando

Na raiz do repositório:

```powershell
go test ./lessons/01-go-foundations/basics -v
```

No início, o primeiro teste falhará por causa do `panic` intencional. Trabalhe em uma função por vez e execute novamente. Quando todos passarem:

```powershell
go fmt ./lessons/01-go-foundations/basics
go vet ./lessons/01-go-foundations/basics
```

## Checkpoint

Depois de concluir, responda com suas palavras:

1. Qual é a diferença entre um parâmetro e o valor enviado na chamada?
2. O que o `int` depois dos parênteses representa?
3. Qual é a diferença entre `:=` e `=`?
4. O que acontece com `var count int` quando nenhum valor inicial é informado?

## Fontes primárias

- [Function declarations — Go specification](https://go.dev/ref/spec#Function_declarations)
- [Short variable declarations — Go specification](https://go.dev/ref/spec#Short_variable_declarations)
- [Variables with initializers — A Tour of Go](https://go.dev/tour/basics/9)
- [Type inference — A Tour of Go](https://go.dev/tour/basics/14)

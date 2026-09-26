# Lição 1 — do módulo ao feedback dos testes

Status: preparada, aguardando revisão do diagnóstico.

Esta lição será ajustada caso o diagnóstico mostre que o conteúdo está abaixo ou acima do seu ponto de partida. Não a use antes de registrar as respostas e a primeira tentativa do diagnóstico.

## Resultado esperado

Ao final, você deverá conseguir explicar como módulo, pacote e arquivo se relacionam; ler uma assinatura com múltiplos retornos; raciocinar sobre valores e slices; tratar uma falha como valor; e usar testes para orientar uma implementação pequena.

## 1. Módulo, pacote e arquivo

Um módulo é a unidade identificada pelo `go.mod`: ele declara o caminho do módulo e registra informações necessárias ao gerenciamento de dependências. Um módulo pode conter vários pacotes.

Um pacote reúne arquivos Go do mesmo diretório que usam o mesmo nome de pacote e são compilados juntos. Um arquivo é apenas uma parte física desse pacote; ele não equivale a um módulo JavaScript.

Comparação útil com Node.js:

- `go.mod` tem parte do papel de identidade e dependências exercido por `package.json`;
- um pacote Go é uma unidade de compilação e organização, não um arquivo importável isolado;
- a analogia termina aí: resolução, visibilidade e inicialização seguem regras próprias de Go.

No diagnóstico, `github.com/delkz/study-golang` é o módulo e `diagnostic/scorestats` é um pacote dentro dele.

## 2. Valores, cópia e slices

Uma atribuição comum em Go copia o valor atribuído. Isso não significa que todo dado subjacente seja profundamente copiado.

Uma slice é um pequeno descritor que referencia uma região de um array subjacente. Copiar uma slice copia esse descritor; duas slices podem, portanto, observar o mesmo armazenamento. `append` pode reutilizar esse armazenamento ou alocar outro, dependendo da capacidade disponível. É por isso que código com slices deve ser analisado em termos de comprimento, capacidade e compartilhamento do array subjacente.

Comparação com TypeScript: existe uma semelhança superficial com duas variáveis apontando para o mesmo array, mas slices possuem comprimento e capacidade próprios e `append` pode mudar a região de armazenamento usada por apenas uma delas.

## 3. Múltiplos retornos e erros como valores

Go permite que uma função retorne mais de um valor. É comum que o último seja um `error`:

```text
resultado, err := operação()
```

`error` participa do fluxo como qualquer outro valor. Quem chama a função decide se deve retornar o erro, acrescentar contexto, tentar outra estratégia ou convertê-lo em uma resposta externa. Isso difere de uma exceção lançada, que transfere o controle implicitamente até algum ponto de captura.

No diagnóstico, `Analyze` devolve `Summary` e `error`. O contrato determina que uma falha deve vir acompanhada do valor zero de `Summary`, evitando que um resultado parcial pareça válido.

## 4. Testes como contrato executável

Arquivos terminados em `_test.go` são reconhecidos pelo comando `go test`. As funções de teste exercitam comportamento observável: entradas, saídas, erros e efeitos colaterais.

Os testes fornecidos verificam:

- casos comuns e valores de fronteira;
- divisão decimal no cálculo da média;
- rejeição de entradas vazias ou inválidas;
- retorno do valor zero quando ocorre erro;
- preservação da slice recebida.

Leia uma falha de teste como evidência específica sobre o contrato, não como instrução para inserir valores especiais que façam apenas aquele caso passar.

## Prática guiada

1. Explique por que `Summary{}` representa o valor zero dessa struct.
2. Explique por que inicializar mínimo e máximo com `0` pode produzir um algoritmo frágil em outros domínios, mesmo que notas válidas incluam `0`.
3. Implemente `Analyze` sem alterar os testes.
4. Execute `go test ./diagnostic/scorestats -v` após cada mudança pequena.
5. Quando os testes passarem, execute `go fmt ./...` e `go vet ./...`.
6. Explique a complexidade de tempo e de memória adicional da solução.

## Checkpoint

Ao concluir, responda:

- O que é copiado quando uma slice é atribuída a outra variável?
- Por que `float64(total / count)` não corrige uma divisão inteira já realizada?
- Quem é responsável por decidir o que fazer com o `error` retornado por `Analyze`?
- O que os testes provaram e o que eles ainda não provam sobre a implementação?

## Fontes primárias

- [Go Modules Reference](https://go.dev/ref/mod)
- [The Go Programming Language Specification](https://go.dev/ref/spec)
- [Add a test](https://go.dev/doc/tutorial/add-a-test)
- [Errors are values](https://go.dev/blog/errors-are-values)

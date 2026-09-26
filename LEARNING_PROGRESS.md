# Progresso de aprendizagem em Go

Última atualização: 2026-09-26

## Objetivo

Desenvolver domínio prático de Go, do básico ao avançado, suficiente para projetar, implementar, testar e explicar aplicações de back-end com segurança profissional, complementando a experiência avançada em front-end e os conhecimentos de TypeScript e Node.js.

## Situação atual

- Etapa: fundamentos iniciais.
- Exercício atual: Lição 1 concluída; próxima lição ainda não iniciada.
- Evidências registradas: implementou e validou funções com parâmetros, variáveis locais, inferência, atribuição, expressões e retorno; explicou os conceitos no checkpoint.
- Restrições pedagógicas: a implementação principal deve ser escrita pelo estudante.

O estado de uma competência representa a evidência disponível no momento, não uma classificação permanente. Uma competência pode voltar para `revisar` quando surgirem lacunas ou quando não houver retenção em uma aplicação posterior.

## Mapa de competências

| Área | Estado | Evidência |
| --- | --- | --- |
| Fundamentos e modelo da linguagem | em prática | Implementou três funções tipadas; testes, formatação e `go vet` passaram; explicou parâmetros, retorno, `:=`, `=` e valor zero. |
| Módulos, pacotes e organização | não avaliado | — |
| Tratamento de erros | não avaliado | — |
| Testes | não avaliado | — |
| Interfaces e composição | não avaliado | — |
| Contexto e concorrência | não avaliado | — |
| HTTP e APIs | não avaliado | — |
| Persistência de dados | não avaliado | — |
| Segurança e observabilidade | não avaliado | — |
| Desempenho e operação | não avaliado | — |
| Arquitetura e manutenção | não avaliado | — |

Estados permitidos: `não avaliado`, `apresentado`, `em prática`, `demonstrado` e `revisar`.

## Próximo passo recomendado

Praticar funções e variáveis em uma pequena variação que introduza `bool` e condicionais sem misturar ainda loops, slices, structs ou erros. Retomar `scorestats.Analyze` somente depois desses fundamentos.

## Histórico de sessões

### 2026-09-26 — preparação do ambiente de aprendizagem

- Definidas as regras pedagógicas e a skill local de mentoria.
- Criados o acompanhamento persistente, o catálogo de fontes e a documentação inicial do repositório.
- Nenhuma competência técnica de Go foi avaliada nesta preparação.
- Próxima ação: realizar o diagnóstico inicial.

### 2026-09-26 — diagnóstico e primeira lição preparados

- Criado um diagnóstico com perguntas de modelo mental, leitura de código e implementação curta.
- Adicionados testes que especificam o comportamento esperado sem implementar a solução.
- Preparada uma primeira lição sobre módulos, pacotes, slices, erros e testes.
- Nenhuma competência foi reclassificada: a avaliação depende da tentativa do estudante.
- Próxima ação: concluir e enviar o diagnóstico para revisão.

### 2026-09-26 — ponto de partida ajustado

- As respostas conceituais e a dificuldade de iniciar `scorestats.Analyze` mostraram que o diagnóstico prático reunia conceitos demais para o ponto de partida.
- A implementação de `Analyze` foi adiada; o arquivo permanece como desafio futuro, sem contar como pendência da primeira lição.
- A Lição 1 foi reduzida a funções, parâmetros, variáveis locais, expressões e retorno.
- Nenhuma competência foi marcada como demonstrada.
- Próxima ação: implementar as funções básicas da Lição 1.

### 2026-09-26 — Lição 1 concluída

- Implementadas `Sum`, `Greeting` e `RectangleArea` com parâmetros e retornos tipados.
- `go test`, `gofmt` e `go vet` foram executados sem problemas.
- O checkpoint confirmou entendimento de parâmetros e argumentos, tipo de retorno, declaração com `:=`, atribuição com `=` e valor zero.
- Fundamentos passaram para `em prática`; ainda falta demonstrar transferência em outro exercício.
- Próxima ação: introduzir expressões booleanas e condicionais em uma atividade curta.

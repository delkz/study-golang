# Respostas do diagnóstico

Preencha este arquivo sem pesquisar as respostas. `Não sei` é uma resposta válida.

## Parte 1 — modelo mental

### 1. Linguagem compilada e estaticamente tipada
Uma linguagem compilada é uma linguagem que passa por um compilador e é traduzida para a maquina, como C#, em vez de rodar direto no ambiente como javascript
Estaticamente tipada é como typescript, onde você especifica o tipo da variavel

### 2. Módulos e pacotes
Acredito que seja como os pacotes do javascript, no npm, sao codigos feitos por outras pessoas que podem ser reutiliados

### 3. Tratamento de falhas
não sei

### 4. Arrays e slices
não sei

### 5. Experiência anterior com Go
Apenas dei uma lida na doc inicial

## Parte 2 — leitura de código

### Previsão antes de executar
numbers = [2,4]
alias = [10,4,8]

Alias é uma copia de numbers, as alterações não deveriam refletir em numbers, somente em alias

### Resultado observado depois de executar

Na verdade numbers foi alterado tambem, então acredito que só atribuir a variavel faz uma copia que altera as duas
porem o append foi só no alias como imaginei

### O que explica eventuais diferenças

talvez criar uma variavel atribuindo um array linke as duas, talvez o modelo correto seja desestruturando? 
# Study Go

Repositório de estudo prático e progressivo de Go, com exercícios que evoluem de fundamentos da linguagem até construção e operação de serviços de back-end.

O objetivo não é acumular exemplos prontos. Cada etapa deve produzir evidências de que os conceitos foram compreendidos, aplicados, testados e conseguem ser explicados. A implementação principal dos exercícios é escrita pelo estudante; a IA atua como mentora, revisora e parceira de diagnóstico.

## Estado atual

O repositório está na etapa de diagnóstico inicial. Ainda não existe um exercício ativo nem competências classificadas como demonstradas. O primeiro passo será identificar o ponto de partida e escolher uma atividade curta, sem presumir conhecimento específico de Go nem repetir fundamentos já dominados.

O estado atualizado da jornada está em [LEARNING_PROGRESS.md](LEARNING_PROGRESS.md).

## Como funciona

Cada ciclo de estudo segue, em linhas gerais, este fluxo:

1. Definir uma conquista pequena e verificável para a sessão.
2. Apresentar apenas os conceitos necessários para alcançá-la.
3. Implementar uma solução própria em Go.
4. Validar com testes e ferramentas adequadas.
5. Revisar comportamento, clareza, tratamento de erros e uso idiomático da linguagem.
6. Registrar evidências, dificuldades e o próximo passo.

Quando houver uma dúvida, a ajuda será gradual: explicação e perguntas primeiro; depois dicas, pseudocódigo ou pequenos exemplos. Uma solução completa não deve ser entregue antes de existir oportunidade real de prática.

## Princípios

- Aprender construindo situações próximas do trabalho real.
- Priorizar a biblioteca padrão antes de adicionar abstrações externas.
- Comparar Go com TypeScript e Node.js quando isso melhorar o modelo mental, deixando claros os limites da comparação.
- Usar testes como especificação de comportamento, não apenas como confirmação final.
- Tratar código que compila como ponto de partida, não como garantia de qualidade.
- Avançar com base em evidências observáveis, não em uma trilha rígida ou sensação de familiaridade.
- Consultar fontes oficiais e atuais ao introduzir conceitos ou comportamentos dependentes de versão.

## Trilha de competências

A ordem será adaptada ao desempenho observado, mas a jornada deverá passar gradualmente por:

- fundamentos e modelo da linguagem;
- módulos, pacotes e organização de código;
- tratamento de erros e testes;
- interfaces, composição e design de APIs;
- contexto e concorrência;
- HTTP e serviços de back-end;
- persistência de dados;
- segurança e observabilidade;
- desempenho e operação;
- arquitetura e manutenção de aplicações.

Os estados usados no acompanhamento são `não avaliado`, `apresentado`, `em prática`, `demonstrado` e `revisar`.

## Estrutura do repositório

```text
.
├── .agents/skills/go-learning-mentor/  # Skill local da mentoria
├── AGENTS.md                           # Regras pedagógicas obrigatórias
├── LEARNING_PROGRESS.md                # Estado e evidências entre sessões
├── RESOURCES.md                        # Fontes realmente utilizadas
└── README.md                           # Visão geral do projeto
```

Diretórios e módulos de exercícios serão adicionados conforme a progressão. A estrutura não será criada antecipadamente sem uma necessidade concreta.

## Preparação do ambiente

1. Instale Go seguindo a [documentação oficial](https://go.dev/doc/install).
2. Confirme a instalação no PowerShell:

   ```powershell
   go version
   ```

3. Abra este repositório no ambiente de desenvolvimento.
4. Inicie a primeira sessão dizendo: `Quero começar meu diagnóstico de Go.`

Não é necessário inicializar um módulo antes de definir o primeiro exercício. Quando existir código Go, os comandos de validação serão documentados junto ao respectivo exercício.

## Validação esperada

Conforme o conteúdo permitir, as soluções serão verificadas com ferramentas padrão do ecossistema:

```powershell
go fmt ./...
go test ./...
go vet ./...
go test -race ./...
```

Nem todos esses comandos serão aplicáveis desde o primeiro exercício. Cada atividade informará seus próprios critérios e comandos de validação.

## Uso de IA

A IA é parte explícita do processo de mentoria, mas não substitui a prática. Ela pode explicar conceitos, elaborar testes, revisar soluções, identificar lacunas e propor o próximo desafio. Decisões, implementação principal e demonstração de entendimento permanecem sob responsabilidade do estudante.

As regras completas dessa colaboração estão em [AGENTS.md](AGENTS.md).

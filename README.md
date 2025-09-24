# Calculadora Científica em Go

Uma calculadora científica web desenvolvida em **Go** com interface interativa no navegador. Permite realizar operações matemáticas básicas e avançadas, incluindo funções científicas como seno, cosseno, tangente, logaritmos, exponenciais, fatorial, e muito mais.

---

## Funcionalidades

- Operações básicas: `+`, `-`, `*`, `/`  
- Funções científicas: `sin`, `cos`, `tan`, `asin`, `acos`, `atan`, `ln`, `log`, `exp`, `sqrt`  
- Operações avançadas: `^` (potência), `%` (porcentagem), `abs`, `floor`, `ceil`, `n!` (fatorial)  
- Visor interativo que atualiza apenas com cliques nos botões  
- Suporte a **graus e radianos** para funções trigonométricas  
- Design inspirado em calculadoras científicas físicas  
- Interface colorida:  
  - Números em azul clarinho  
  - Funções em roxo clarinho  
  - Operadores em cinza  

---

## Pré-requisitos

- Go >= 1.25  
- Navegador moderno (Chrome, Firefox, Edge, etc.)

---

## Como rodar

1. Clone o repositório:

```bash
git clone <URL_DO_REPOSITORIO>
cd <PASTA_DO_PROJETO>
```

2. Inicialize o módulo Go:

```bash
go mod init calculadora
go mod tidy
```

3. Execute o servidor:

```bash
go run main.go
```

4. Abra o navegador em:

```
http://localhost:8080
```

5. Clique nos botões para usar a calculadora!

---

## Estrutura do projeto

```
calculadora/
│
├── main.go         # Servidor Go que processa as operações
├── go.mod          # Módulo Go
├── go.sum          # Dependências
└── static/
    └── index.html  # Interface da calculadora
```

---

## Tecnologias

- [Go](https://golang.org/) para o backend  
- HTML, CSS e JavaScript para a interface web  
- API REST para comunicação entre frontend e backend

---

## Observações

- A função `sin`, `cos` e `tan` pode ser usada em **graus ou radianos**, de acordo com o checkbox “Graus” na interface.  
- O visor é **apenas de leitura**: todas as operações devem ser feitas através dos botões.  
- Novas operações apagam automaticamente o visor após o cálculo.

---

## Autor

- Samuel De Andrade (KernelKn1ght)

---



# MyAPI
## Introdução
MyAPI é uma API desenvolvida em Golang utilizando o framework Gin. Esta API oferece endpoints para autenticação, consulta de dados protegidos e upload de arquivos XLSX.

## Requisitos
- Golang 1.16 ou superior
- Módulos Golang:
   - github.com/gin-gonic/gin
   - github.com/tealeg/xlsx

## Instalação

1. Clone o repositório:

```bash
Copiar código
git clone https://github.com/seu-usuario/myapi.git
cd myapi
```

2. Instale as dependências:

```bash
Copiar código
go mod tidy

```

3. Rodando a aplicação
Para iniciar o servidor, execute:

```bash
Copiar código
go run main.go

```
O servidor estará rodando na porta 8080.

## Endpoints
- Autenticação
  - Endpoint para autenticar usuário e obter token de acesso.
```
URL: POST /login
```
- Corpo da Requisição:
   - json
- Copiar código
```
{
  "username": "user",
  "password": "pass"
}
```
- Resposta de Sucesso (200 OK):
   - json
- Copiar código
```
{
  "token": "mock-token"
}
```
- Resposta de Falha (401 Unauthorized):
  - json
- Copiar código
```
{
  "error": "Unauthorized"
}
```
## Consulta de Dados Protegidos
 - Endpoint para consultar dados protegidos usando token de autorização.

```
URL: GET /data
```
- Cabeçalho Requerido:
   - makefile
- Copiar código

```
Authorization: Bearer mock-token
```
- Resposta de Sucesso (200 OK):
   - json
- Copiar código

```
[
  ["coluna1_linha1", "coluna2_linha1"],
  ["coluna1_linha2", "coluna2_linha2"]
  // ...
]
```

- Resposta de Falha (401 Unauthorized):
   - json
- Copiar código

```
{
  "error": "Unauthorized"
}
```

## Upload de Arquivo XLSX
- Endpoint para realizar upload de arquivo XLSX contendo dados para processamento.

```
URL: POST /upload
Corpo da Requisição:
bash
Copiar código
file=@path/to/your/file.xlsx
```
- Resposta de Sucesso (200 OK):
   - json
- Copiar código

```
{
  "message": "File processed successfully"
}
```
- Resposta de Falha (400 Bad Request):
   - json
- Copiar código

```
{
  "error": "No file is received"
}
```

## Teste
- Você pode testar os endpoints usando curl ou Postman.

## Autenticação com curl:

```bash
Copiar código
curl -X POST http://localhost:8080/login \
-H "Content-Type: application/json" \
-d '{"username": "user", "password": "pass"}'
```
- Consulta de Dados Protegidos com curl:

```bash
Copiar código
curl -X GET http://localhost:8080/data \
-H "Authorization: Bearer mock-token"
```
- Upload de Arquivo XLSX com curl:

```bash
Copiar código
curl -X POST http://localhost:8080/upload \
-F "file=@path/to/your/file.xlsx"
```
- Autenticação com Postman:

- Configure uma requisição POST para http://localhost:8080/login.
- Adicione o cabeçalho Content-Type: application/json.
- No corpo da requisição, adicione:
   - json
- Copiar código

```
{
  "username": "user",
  "password": "pass"
}
```
## Consulta de Dados Protegidos com Postman:

- Configure uma requisição GET para http://localhost:8080/data.
- Adicione o cabeçalho Authorization: Bearer mock-token.
- Upload de Arquivo XLSX com Postman:

- Configure uma requisição POST para http://localhost:8080/upload.
- Adicione um campo do tipo arquivo com chave file e selecione o arquivo XLSX.

### Conclusão
- A API MyAPI oferece funcionalidades robustas para autenticação, consulta de dados protegidos e upload de arquivos XLSX. Certifique-se de ter as dependências instaladas e siga as instruções acima para rodar e testar a aplicação.
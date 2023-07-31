[![NPM](https://img.shields.io/npm/l/react)](https://github.com/marciosenaf/api-rest/blob/main/LICENSE) 

# Documentação da API REST em Go
Aqui está a documentação da API REST em Go e do cliente para consumir essa API.

# API REST em Go
## Estrutura do Projeto

```bash
api/
  └── main.go
client/
  └── main.go
```

## Descrição
A API REST em Go é um servidor que expõe um endpoint /users para recuperar uma lista de usuários em formato JSON. A lista de usuários é fixa e não persistente.

## Endpoints
```GET /users```
Retorna uma lista de usuários em formato JSON.

Exemplo de resposta:

```bash
[
  {
    "ID": 1,
    "Name": "Marcio Sena"
  },
  {
    "ID": 5,
    "Name": "Tonico"
  }
]

```

## Código-fonte

```bash 
// api/main.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	ID   int
	Name string
}

func main() {
	http.HandleFunc("/users", getUsers)
	fmt.Println("API está em execução em: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]User{
		{
			ID:   1,
			Name: "Marcio Sena",
		},
		{
			ID:   5,
			Name: "Tonico",
		},
	})
}
```
## Cliente da API em Go

### Descrição
O cliente da API é um programa em Go que consome o endpoint /users da API REST e exibe a lista de usuários no console.

## Código-fonte

```bash
// client/main.go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	ID   int
	Name string
}

func main() {
	resp, err := http.Get("http://localhost:8080/users")
	if err != nil {
		fmt.Println("Erro ao fazer a requisição:", err.Error())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("Não foi possível obter os dados. Código de status:", resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler o corpo da resposta:", err.Error())
		return
	}

	var users []User
	err = json.Unmarshal(body, &users)
	if err != nil {
		fmt.Println("Erro ao decodificar a resposta:", err.Error())
		return
	}

	fmt.Println("Lista de Usuários:")
	for _, user := range users {
		fmt.Printf("ID: %d, Nome: %s\n", user.ID, user.Name)
	}
}
```
## Executando a API e o Cliente
- Abra um terminal e navegue até a pasta ```api```.
- Execute o seguinte comando para iniciar a API:

```bash
go run main.go
```
- Abra outro terminal e navegue até a pasta client.
- Execute o seguinte comando para iniciar o cliente:

```bash
go run main.go
```

O cliente irá consumir a API e exibir a lista de usuários retornada pelo servidor.

Nota: Certifique-se de que a API esteja em execução antes de executar o cliente. Caso contrário, o cliente mostrará um erro ao tentar se conectar à API.

# Autor

Márcio sena

https://www.linkedin.com/in/marciosenaf/









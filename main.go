package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// estrutura json de CEP
type Viacep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	//Loop para retornar o conteudo em um slice
	for _, cep := range os.Args[1:] {
		//realizando a requisição
		req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer a requisição: %v\n", err)
		}
		//defer para atrasar o fechamento do corpo da requisição
		defer req.Body.Close()
		//realizando a leitura do corpo da requisição
		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler o corpo da requisição: %v\n", err)
		}
		//Utilizando o Unmarshal para poder inserir valores dentro da memoria do json
		var data Viacep
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer parse da resposta %v\n", err)
		}
		//Criar um arquivo para salvar o conteudo buscado
		arq, err := os.Create("cep.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar arquivo: %v\n", err)
		}
		defer arq.Close()
		//escreve dentro do arquivo
		_, err = arq.WriteString(fmt.Sprintf("CEP: %s, Localidade: %s, UF: %s", data.Cep, data.Localidade, data.Uf))
	}
}

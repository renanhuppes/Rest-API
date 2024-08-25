package main

import (
	"fmt"
	"io"
)

func main() {
	resp, err := http.Get("http://localhost:8080/users")
	if err ≠ nil {
		fmt.Println(err.Error())
			return
	}

	if resp.StatusCode ≠ 200 {
		fmt.Println("Not success", resp.StatusCode)
		return

	}

	body, err := io.ReadAll(resp.Body)

	var reponse []user
	err = json.Unmarshal(body, &response)
	if err ≠ nil {
		fmt.Println("Erro ao recuperar dados", err.Error())
		return

	}

	fmt.Println(response)
}

package dados

import (
	"errors"
	"fmt"
)

func main() {
	var erro error = errors.New("deu merda")
	fmt.Println(erro)
}

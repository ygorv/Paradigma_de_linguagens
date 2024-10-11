package main
// questões 10, 11, 12 e 13
import (
	"errors"
	"fmt"
)

type Calculadora struct{}

func (c Calculadora) SomarDois(a, b int) int {
	return a + b
}

func (c Calculadora) SomarTres(a, b, d int) int {
	return a + b + d
}

type Funcionario interface {
	CalcularSalario() float64
	GetNome() string
}

type FuncionarioHorista struct {
	nome             string
	horasTrabalhadas int
	valorHora        float64
}

func NovoFuncionarioHorista(nome string, horasTrabalhadas int, valorHora float64) FuncionarioHorista {
	return FuncionarioHorista{nome, horasTrabalhadas, valorHora}
}

func (f FuncionarioHorista) CalcularSalario() float64 {
	return float64(f.horasTrabalhadas) * f.valorHora
}

func (f FuncionarioHorista) GetNome() string {
	return f.nome
}

type FuncionarioAssalariado struct {
	nome          string
	salarioMensal float64
}

func NovoFuncionarioAssalariado(nome string, salarioMensal float64) FuncionarioAssalariado {
	return FuncionarioAssalariado{nome, salarioMensal}
}

func (f FuncionarioAssalariado) CalcularSalario() float64 {
	return f.salarioMensal
}

func (f FuncionarioAssalariado) GetNome() string {
	return f.nome
}

type Produto struct {
	nome  string
	preco float64
}

func NovoProduto(nome string, preco float64) Produto {
	return Produto{nome, preco}
}

func (p Produto) Somar(outro Produto) (Produto, error) {
	if outro == (Produto{}) {
		return Produto{}, errors.New("produto não pode ser nulo")
	}
	novoNome := p.nome + " + " + outro.nome
	novoPreco := p.preco + outro.preco
	return Produto{novoNome, novoPreco}, nil
}

func (p Produto) String() string {
	return fmt.Sprintf("%s: R$%.2f", p.nome, p.preco)
}

type Matematica struct{}

func Fatorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("o fatorial não está definido para números negativos")
	}
	if n == 0 || n == 1 {
		return 1, nil
	}
	fatorialAnterior, _ := Fatorial(n - 1)
	return n * fatorialAnterior, nil
}

func main() {
	calc := Calculadora{}
	fmt.Println("Soma de 5 + 3:", calc.SomarDois(5, 3))
	fmt.Println("Soma de 5 + 3 + 2:", calc.SomarTres(5, 3, 2))

	horista := NovoFuncionarioHorista("João", 160, 20)
	assalariado := NovoFuncionarioAssalariado("Maria", 3000)

	fmt.Printf("Salário de %s: R$%.2f\n", horista.GetNome(), horista.CalcularSalario())
	fmt.Printf("Salário de %s: R$%.2f\n", assalariado.GetNome(), assalariado.CalcularSalario())

	produto1 := NovoProduto("Camiseta", 50.0)
	produto2 := NovoProduto("Calça", 100.0)
	produtoSoma, err := produto1.Somar(produto2)
	if err != nil {
		fmt.Println("Erro ao somar produtos:", err)
	} else {
		fmt.Println("Produto 1:", produto1)
		fmt.Println("Produto 2:", produto2)
		fmt.Println("Soma dos produtos:", produtoSoma)
	}

	fat5, _ := Fatorial(5)
	fmt.Println("Fatorial de 5:", fat5)

	fat0, _ := Fatorial(0)
	fmt.Println("Fatorial de 0:", fat0)
}

package main
// questões 8 e 9
import (
	"fmt"
)

type Empregado struct {
	nome    string
	cargo   string
	salario float64
}

func NewEmpregado(nome, cargo string, salario float64) *Empregado {
	return &Empregado{nome: nome, cargo: cargo, salario: salario}
}

func (e *Empregado) String() string {
	return fmt.Sprintf("%s - %s (Salário: R$%.2f)", e.nome, e.cargo, e.salario)
}

type Imprimivel interface {
	imprimir()
}

type Relatorio struct {
	titulo string
}

func NewRelatorio(titulo string) *Relatorio {
	return &Relatorio{titulo: titulo}
}

func (r *Relatorio) imprimir() {
	fmt.Println("Imprimindo relatório:", r.titulo)
	fmt.Println("Conteúdo do relatório...")
}

type Contrato struct {
	numero string
}

func NewContrato(numero string) *Contrato {
	return &Contrato{numero: numero}
}

func (c *Contrato) imprimir() {
	fmt.Println("Imprimindo contrato número:", c.numero)
	fmt.Println("Termos e condições do contrato...")
}

type Empresa struct {
	nome      string
	empregados []*Empregado
}

func NewEmpresa(nome string) *Empresa {
	return &Empresa{nome: nome, empregados: []*Empregado{}}
}

func (e *Empresa) adicionarEmpregado(empregado *Empregado) {
	e.empregados = append(e.empregados, empregado)
}

func (e *Empresa) listarEmpregados() {
	fmt.Println("Empregados da", e.nome+":")
	for _, empregado := range e.empregados {
		fmt.Println(empregado)
	}
}

func imprimirDocumento(doc Imprimivel) {
	doc.imprimir()
}

func main() {
	empresa := NewEmpresa("Minha Empresa Ltda.")

	emp1 := NewEmpregado("João Silva", "Desenvolvedor", 5000)
	emp2 := NewEmpregado("Maria Santos", "Gerente de Projetos", 7000)
	emp3 := NewEmpregado("Carlos Oliveira", "Analista de Dados", 4500)

	empresa.adicionarEmpregado(emp1)
	empresa.adicionarEmpregado(emp2)
	empresa.adicionarEmpregado(emp3)

	empresa.listarEmpregados()

	fmt.Println("\nTestando a interface Imprimivel:")
	relatorio := NewRelatorio("Relatório Anual")
	contrato := NewContrato("C-12345")

	imprimirDocumento(relatorio)
	fmt.Println()
	imprimirDocumento(contrato)
}

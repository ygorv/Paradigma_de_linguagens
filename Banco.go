cdpackage main
// questões 3 e 15
import (
	"errors"
	"fmt"
)

// ContaBancaria representa uma conta bancária simples
type ContaBancaria struct {
	titular string
	saldo   float64
}

// Novo construtor para criar uma nova conta bancária
func NovaContaBancaria(titular string, saldoInicial float64) ContaBancaria {
	return ContaBancaria{titular: titular, saldo: saldoInicial}
}

// GetTitular retorna o titular da conta
func (c ContaBancaria) GetTitular() string {
	return c.titular
}

// GetSaldo retorna o saldo da conta
func (c ContaBancaria) GetSaldo() float64 {
	return c.saldo
}

// SetSaldo lança um erro porque não é permitido modificar o saldo diretamente
func (c *ContaBancaria) SetSaldo(valor float64) error {
	return errors.New("não é possível modificar o saldo diretamente. Use os métodos depositar() ou sacar()")
}

// Depositar adiciona um valor ao saldo, se o valor for positivo
func (c *ContaBancaria) Depositar(valor float64) bool {
	if valor > 0 {
		c.saldo += valor
		return true
	}
	return false
}

// Sacar retira um valor do saldo, se possível
func (c *ContaBancaria) Sacar(valor float64) error {
	if valor > 0 {
		if c.saldo >= valor {
			c.saldo -= valor
			return nil
		} else {
			return NovaSaldoInsuficienteException(c.saldo, valor)
		}
	}
	return nil
}

// SaldoInsuficienteException representa um erro de saldo insuficiente
type SaldoInsuficienteException struct {
	saldoAtual  float64
	valorSaque  float64
	mensagem    string
}

// Novo construtor para criar um erro de saldo insuficiente
func NovaSaldoInsuficienteException(saldoAtual, valorSaque float64) *SaldoInsuficienteException {
	return &SaldoInsuficienteException{
		saldoAtual: saldoAtual,
		valorSaque: valorSaque,
		mensagem:   fmt.Sprintf("Saldo insuficiente. Saldo atual: R$%.2f, Valor do saque: R$%.2f", saldoAtual, valorSaque),
	}
}

// Método Error para o erro de saldo insuficiente
func (e *SaldoInsuficienteException) Error() string {
	return e.mensagem
}

// Função principal
func main() {
	contaMaria := NovaContaBancaria("Maria", 1000)

	fmt.Printf("Titular: %s\n", contaMaria.GetTitular())
	fmt.Printf("Saldo inicial: R$%.2f\n", contaMaria.GetSaldo())

	if contaMaria.Depositar(500) {
		fmt.Printf("Depósito realizado. Novo saldo: R$%.2f\n", contaMaria.GetSaldo())
	} else {
		fmt.Println("Falha no depósito.")
	}

	err := contaMaria.Sacar(200)
	if err == nil {
		fmt.Printf("Saque realizado. Novo saldo: R$%.2f\n", contaMaria.GetSaldo())
	} else {
		fmt.Printf("Falha no saque: %s\n", err.Error())
	}

	err = contaMaria.Sacar(2000)
	if err != nil {
		fmt.Printf("Falha no saque: %s\n", err.Error())
	}

	err = contaMaria.SetSaldo(5000)
	if err != nil {
		fmt.Printf("Erro ao tentar modificar o saldo: %s\n", err.Error())
	}
}

class SaldoInsuficienteException(Exception):
    def __init__(self, saldo, valor):
        self.saldo = saldo
        self.valor = valor
        super().__init__(f"Saldo insuficiente. Saldo atual: {self.saldo}, Tentativa de saque: {self.valor}")

class ContaBancaria:
    def __init__(self, saldo_inicial: float):
        self.saldo = saldo_inicial

    def sacar(self, valor: float):
        if valor > self.saldo:
            raise SaldoInsuficienteException(self.saldo, valor)
        self.saldo -= valor
        print(f"Saque realizado: {valor}")
        print(f"Saldo restante: {self.saldo}")

conta = ContaBancaria(100.0)
try:
    conta.sacar(150.0) 
except SaldoInsuficienteException as e:
    print(e)

class ContaBancaria:
    def __init__(self, titular, saldo_inicial=0):
        self.titular = titular
        self.__saldo = saldo_inicial   
    
   
    def depositar(self, valor):
        if valor > 0:
            self.__saldo += valor
            print(f'Depósito de R${valor:.2f} realizado com sucesso!')
        else:
            print('Valor de depósito inválido.')
    
   
    def sacar(self, valor):
        if 0 < valor <= self.__saldo:
            self.__saldo -= valor
            print(f'Saque de R${valor:.2f} realizado com sucesso!')
        else:
            print('Saque inválido ou saldo insuficiente.')
    
   
    def consultar_saldo(self):
        return self.__saldo


conta = ContaBancaria('João', 1000)
print(f'Titular: {conta.titular}')
print(f'Saldo inicial: R${conta.consultar_saldo():.2f}')

conta.depositar(500)
print(f'Saldo após depósito: R${conta.consultar_saldo():.2f}')

conta.sacar(300)
print(f'Saldo após saque: R${conta.consultar_saldo():.2f}')

conta.sacar(1500)  

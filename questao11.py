from abc import ABC, abstractmethod

class Funcionario(ABC):
    def __init__(self, nome):
        self.nome = nome

    @abstractmethod
    def calcularSalario(self):
        pass

class FuncionarioHorista(Funcionario):
    def __init__(self, nome, horas_trabalhadas, valor_hora):
        super().__init__(nome)
        self.horas_trabalhadas = horas_trabalhadas
        self.valor_hora = valor_hora

    def calcularSalario(self):
        return self.horas_trabalhadas * self.valor_hora

class FuncionarioAssalariado(Funcionario):
    def __init__(self, nome, salario_fixo):
        super().__init__(nome)
        self.salario_fixo = salario_fixo

    def calcularSalario(self):
        return self.salario_fixo


if __name__ == "__main__":

    horista = FuncionarioHorista("João", 160, 25)  
    print(f"Salário de {horista.nome} (horista): R${horista.calcularSalario():.2f}")

    assalariado = FuncionarioAssalariado("Maria", 3000)  
    print(f"Salário de {assalariado.nome} (assalariado): R${assalariado.calcularSalario():.2f}")

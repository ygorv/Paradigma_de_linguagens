class Empregado:
    def __init__(self, nome, cargo, salario):
        self.nome = nome
        self.cargo = cargo
        self.salario = salario

    def __str__(self):
        return f"{self.nome}, Cargo: {self.cargo}, Salário: R${self.salario:.2f}"


class Empresa:
    def __init__(self, nome):
        self.nome = nome
        self.empregados = []

    def adicionar_empregado(self, empregado):
        self.empregados.append(empregado)

    def exibir_empregados(self):
        print(f"Empregados da empresa {self.nome}:")
        for empregado in self.empregados:
            print(empregado)

emp1 = Empregado("João", "Desenvolvedor", 5000)
emp2 = Empregado("Maria", "Gerente", 8000)
emp3 = Empregado("Pedro", "Designer", 4000)

empresa = Empresa("Tech Solutions")

empresa.adicionar_empregado(emp1)
empresa.adicionar_empregado(emp2)
empresa.adicionar_empregado(emp3)

empresa.exibir_empregados()

print("\nEmpregados fora da empresa:")
print(emp1)
print(emp2)

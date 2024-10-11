class Carro:
    def __init__(self, marca, modelo, ano):
        self.marca = marca
        self.modelo = modelo
        self.ano = ano
        self.velocidade = 0
    
    def acelerar(self, incremento):
        self.velocidade += incremento
    
    def frear(self, decremento):
        self.velocidade = max(0, self.velocidade - decremento) 
    
    def exibir_velocidade(self):
        return f"Velocidade atual: {self.velocidade} km/h"


carro1 = Carro("Toyota", "Corolla", 2020)


carro1.acelerar(30)
print(carro1.exibir_velocidade()) 

carro1.frear(10)
print(carro1.exibir_velocidade())  

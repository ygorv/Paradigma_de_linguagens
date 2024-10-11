class Motor:
    def __init__(self, potencia):
        self.potencia = potencia
    
    def ligar(self):
        return f'Motor de {self.potencia} cavalos ligado.'
    
    def desligar(self):
        return 'Motor desligado.'


class Carro:
    def __init__(self, marca, motor):
        self.marca = marca
        self.motor = motor 
    
    def ligar_carro(self):
        print(f'Ligando o carro da marca {self.marca}.')
        print(self.motor.ligar())
    
    def desligar_carro(self):
        print(f'Desligando o carro da marca {self.marca}.')
        print(self.motor.desligar())

motor_v8 = Motor(500)  
carro = Carro('Ferrari', motor_v8) 

carro.ligar_carro()
carro.desligar_carro()

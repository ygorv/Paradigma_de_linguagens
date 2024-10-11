class Animal:
    def __init__(self, nome):
        self.nome = nome
    
    def som(self):
        pass

class Cachorro(Animal):
    def som(self):
        return f'{self.nome} faz: Au au!'

class Gato(Animal):
    def som(self):
        return f'{self.nome} faz: Miau!'

class Pato(Animal):
    def som(self):
        return f'{self.nome} faz: Quack quack!'

def emitir_sons(animais):
    for animal in animais:
        print(animal.som())


cachorro = Cachorro('Bob')
gato = Gato('Mimi')
pato = Pato('Donald')

lista_de_animais = [cachorro, gato, pato]


emitir_sons(lista_de_animais)

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


cachorro = Cachorro('Bob')
gato = Gato('Mimi')

print(cachorro.som())  
print(gato.som())      

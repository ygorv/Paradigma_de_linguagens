class Produto:
    def __init__(self, nome, preco):
        self.nome = nome
        self.preco = preco

    def __add__(self, outro_produto):
        if isinstance(outro_produto, Produto):
            return self.preco + outro_produto.preco
        raise TypeError("O objeto deve ser uma instância da classe Produto")

    def __str__(self):
        return f"Produto: {self.nome}, Preço: R${self.preco:.2f}"

if __name__ == "__main__":
    produto1 = Produto("Notebook", 2500.00)
    produto2 = Produto("Smartphone", 1500.00)

    print(produto1)
    print(produto2)

    soma_precos = produto1 + produto2
    print(f"Soma dos preços: R${soma_precos:.2f}")

class Matematica:

    @staticmethod
    def fatorial(n: int) -> int:
        if n == 0:
            return 1
        resultado = 1
        for i in range(1, n + 1):
            resultado *= i
        return resultado

numero = 5
resultado = Matematica.fatorial(numero)
print(f"Fatorial de {numero} é: {resultado}")

class Calculadora:
    def somar(self, *args):
        if len(args) == 2:
            return self._somar_dois(args[0], args[1])
        elif len(args) == 3:
            return self._somar_tres(args[0], args[1], args[2])
        else:
            raise ValueError("O método somar aceita apenas 2 ou 3 números.")

    def _somar_dois(self, a, b):
        return a + b

    def _somar_tres(self, a, b, c):
        return a + b + c

calculadora = Calculadora()

resultado_2_numeros = calculadora.somar(10, 20)
print(f"Resultado da soma de dois números: {resultado_2_numeros}")

resultado_3_numeros = calculadora.somar(5, 10, 15)
print(f"Resultado da soma de três números: {resultado_3_numeros}")

try:
    calculadora.somar(1) 
except ValueError as e:
    print(e)

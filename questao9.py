from abc import ABC, abstractmethod

class Imprimivel(ABC):
    @abstractmethod
    def imprimir(self) -> None:
        pass

class Relatorio(Imprimivel):
    def imprimir(self) -> None:
        print("Imprimindo Relatório...")

class Contrato(Imprimivel):
    def imprimir(self) -> None:
        print("Imprimindo Contrato...")

relatorio = Relatorio()
contrato = Contrato()

relatorio.imprimir()
contrato.imprimir()

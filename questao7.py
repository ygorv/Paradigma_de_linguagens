class Professor:
    def __init__(self, nome):
        self.nome = nome
        self.escolas = []

    def adicionar_escola(self, escola):
        if escola not in self.escolas:
            self.escolas.append(escola)
            escola.adicionar_professor(self)

    def __str__(self):
        return self.nome


class Escola:
    def __init__(self, nome):
        self.nome = nome
        self.professores = []

    def adicionar_professor(self, professor):
        if professor not in self.professores:
            self.professores.append(professor)
            professor.adicionar_escola(self)

    def __str__(self):
        return self.nome

prof1 = Professor("Professor A")
prof2 = Professor("Professor B")


escola1 = Escola("Escola X")
escola2 = Escola("Escola Y")

escola1.adicionar_professor(prof1)
escola1.adicionar_professor(prof2)

escola2.adicionar_professor(prof1)

print(f"Escolas de {prof1.nome}: {[escola.nome for escola in prof1.escolas]}")
print(f"Professores da {escola1.nome}: {[prof.nome for prof in escola1.professores]}")

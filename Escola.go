package main
// questão 7
import (
	"fmt"
)

type Professor struct {
	nome    string
	escolas []*Escola
}

func NewProfessor(nome string) *Professor {
	return &Professor{nome: nome, escolas: []*Escola{}}
}

func (p *Professor) adicionarEscola(escola *Escola) {
	for _, e := range p.escolas {
		if e == escola {
			return
		}
	}
	p.escolas = append(p.escolas, escola)
}

func (p *Professor) removerEscola(escola *Escola) {
	for i, e := range p.escolas {
		if e == escola {
			p.escolas = append(p.escolas[:i], p.escolas[i+1:]...)
			return
		}
	}
}

func (p *Professor) listarEscolas() []string {
	nomes := []string{}
	for _, escola := range p.escolas {
		nomes = append(nomes, escola.getNome())
	}
	return nomes
}

func (p *Professor) getNome() string {
	return p.nome
}

type Escola struct {
	nome      string
	professores []*Professor
}

func NewEscola(nome string) *Escola {
	return &Escola{nome: nome, professores: []*Professor{}}
}

func (e *Escola) adicionarProfessor(professor *Professor) {
	for _, p := range e.professores {
		if p == professor {
			return
		}
	}
	e.professores = append(e.professores, professor)
	professor.adicionarEscola(e)
}

func (e *Escola) removerProfessor(professor *Professor) {
	for i, p := range e.professores {
		if p == professor {
			e.professores = append(e.professores[:i], e.professores[i+1:]...)
			professor.removerEscola(e)
			return
		}
	}
}

func (e *Escola) listarProfessores() []string {
	nomes := []string{}
	for _, professor := range e.professores {
		nomes = append(nomes, professor.getNome())
	}
	return nomes
}

func (e *Escola) getNome() string {
	return e.nome
}

func main() {
	escola1 := NewEscola("Escola A")
	escola2 := NewEscola("Escola B")

	prof1 := NewProfessor("João")
	prof2 := NewProfessor("Maria")

	escola1.adicionarProfessor(prof1)
	escola1.adicionarProfessor(prof2)
	escola2.adicionarProfessor(prof1)

	fmt.Println("Professores da", escola1.getNome()+":", escola1.listarProfessores())
	fmt.Println("Professores da", escola2.getNome()+":", escola2.listarProfessores())
	fmt.Println("Escolas onde", prof1.getNome(), "leciona:", prof1.listarEscolas())
	fmt.Println("Escolas onde", prof2.getNome(), "leciona:", prof2.listarEscolas())
}

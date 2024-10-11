package main
// questão 4 e 5
import "fmt"

// Estrutura Animal
type Animal struct {
    nome  string
    idade int
}

// Construtor padrão para Animal
func NewAnimal() Animal {
    return Animal{}
}

func NewAnimalComNomeIdade(nome string, idade int) Animal {
    return Animal{nome: nome, idade: idade}
}

// Função emitirSom (default para Animal)
func (a Animal) EmitirSom() {
    fmt.Println("O animal emite um som.")
}

// Getters e setters para nome e idade
func (a *Animal) GetNome() string {
    return a.nome
}

func (a *Animal) SetNome(nome string) {
    a.nome = nome
}

func (a *Animal) GetIdade() int {
    return a.idade
}

func (a *Animal) SetIdade(idade int) {
    a.idade = idade
}

// Estrutura Cachorro que herda de Animal
type Cachorro struct {
    Animal
}

func NewCachorro() Cachorro {
    return Cachorro{NewAnimal()}
}

// Sobrescreve a função emitirSom para Cachorro
func (c Cachorro) EmitirSom() {
    fmt.Println("O cachorro late: Au au!")
}

// Estrutura Gato que herda de Animal
type Gato struct {
    Animal
}

func NewGato() Gato {
    return Gato{NewAnimal()}
}

// Sobrescreve a função emitirSom para Gato
func (g Gato) EmitirSom() {
    fmt.Println("O gato mia: Miau!")
}

// Estrutura Vaca que herda de Animal
type Vaca struct {
    Animal
}

func NewVaca() Vaca {
    return Vaca{NewAnimal()}
}

// Sobrescreve a função emitirSom para Vaca
func (v Vaca) EmitirSom() {
    fmt.Println("A vaca muge: Muuu!")
}

// Função fazerBarulho que percorre um array de animais e faz com que emitam som
func FazerBarulho(animais []AnimalInterface) {
    for _, animal := range animais {
        animal.EmitirSom()
    }
}

// Interface para garantir que todos os animais têm EmitirSom
type AnimalInterface interface {
    EmitirSom()
}

func main() {
    cachorro := NewCachorro()
    gato := NewGato()
    vaca := NewVaca()

    // Array de animais que implementam AnimalInterface
    animais := []AnimalInterface{cachorro, gato, vaca}

    FazerBarulho(animais)
}

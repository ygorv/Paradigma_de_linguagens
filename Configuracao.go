package main

import (
	"fmt"
	"sync"
)

type Configuracao struct {
	tema   string
	idioma string
}

var instancia *Configuracao
var once sync.Once

func getInstance() *Configuracao {
	once.Do(func() {
		instancia = &Configuracao{}
		instancia.inicializar()
	})
	return instancia
}

func (c *Configuracao) inicializar() {
	c.tema = "claro"
	c.idioma = "português"
}

func (c *Configuracao) SetTema(tema string) {
	c.tema = tema
}

func (c *Configuracao) SetIdioma(idioma string) {
	c.idioma = idioma
}

func (c *Configuracao) GetConfiguracoes() string {
	return fmt.Sprintf("Tema: %s, Idioma: %s", c.tema, c.idioma)
}

func main() {
	config1 := getInstance()
	config1.SetTema("escuro")

	config2 := getInstance()
	config2.SetIdioma("inglês")

	fmt.Println(config1.GetConfiguracoes())
	fmt.Println(config2.GetConfiguracoes())
	fmt.Println(config1 == config2)
}

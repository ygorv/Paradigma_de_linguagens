package main
// questões 1,2 e 6
import (
    "fmt"
)

type Motor struct {
    tipo     string
    potencia int
}

func NewMotor(tipo string, potencia int) Motor {
    return Motor{tipo: tipo, potencia: potencia}
}

func (m Motor) String() string {
    return fmt.Sprintf("%s - %d cv", m.tipo, m.potencia)
}

type Carro struct {
    marca      string
    modelo     string
    ano        int
    velocidade int
    motor      Motor
}

func NewCarro(marca, modelo string, ano int, tipoMotor string, potenciaMotor int) Carro {
    return Carro{
        marca:      marca,
        modelo:     modelo,
        ano:        ano,
        velocidade: 0,
        motor:      NewMotor(tipoMotor, potenciaMotor),
    }
}

func (c Carro) Motor(tipoMotor string) {
    fmt.Printf("🔧 Motor do %s: %s\n", c.modelo, tipoMotor)
}

func (c Carro) Bateria(capacidadeBateria int) {
    fmt.Printf("🔋 Bateria do %s: %d kWh\n", c.modelo, capacidadeBateria)
}

func (c Carro) Pneus(tipoPneu string) {
    fmt.Printf("🛞 Pneus do %s: %s\n", c.modelo, tipoPneu)
}

func (c Carro) Sobre() {
    fmt.Printf("🚗 %s %s (%d)\n", c.marca, c.modelo, c.ano)
    fmt.Printf("🔧 Motor: %s\n", c.motor)
}

func (c *Carro) Acelerar(incremento int) {
    c.velocidade += incremento
    fmt.Printf("🚀 %s acelerou para %d km/h\n", c.modelo, c.velocidade)
}

func (c *Carro) Frear(decremento int) {
    if c.velocidade-decremento < 0 {
        c.velocidade = 0
    } else {
        c.velocidade -= decremento
    }
    fmt.Printf("🛑 %s freou para %d km/h\n", c.modelo, c.velocidade)
}

func (c Carro) ExibirVelocidade() {
    fmt.Printf("🏁 Velocidade do %s: %d km/h\n", c.modelo, c.velocidade)
}

func ImprimirSeparador() {
    fmt.Println("========================================")
}

func main() {
    carro1 := NewCarro("Toyota", "Corolla", 2020, "V6", 200)
    carro2 := NewCarro("Honda", "Civic", 2018, "V4", 180)
    carro3 := NewCarro("Ford", "Mustang", 2022, "V8", 450)

    carros := []Carro{carro1, carro2, carro3}

    for _, carro := range carros {
        ImprimirSeparador()
        carro.Sobre()
        if carro.marca == "Toyota" {
            carro.Bateria(12)
            carro.Pneus("Radiais")
        } else if carro.marca == "Honda" {
            carro.Bateria(14)
            carro.Pneus("Esportivos")
        } else {
            carro.Bateria(15)
            carro.Pneus("Off-road")
        }
    }

    ImprimirSeparador()
    fmt.Println("\n📊 Teste de aceleração e frenagem (Toyota Corolla):")
    ImprimirSeparador()

    carro1.Acelerar(50)
    carro1.ExibirVelocidade()
    carro1.Acelerar(30)
    carro1.ExibirVelocidade()
    carro1.Frear(20)
    carro1.ExibirVelocidade()
    carro1.Frear(100)
    carro1.ExibirVelocidade()

    ImprimirSeparador()
}

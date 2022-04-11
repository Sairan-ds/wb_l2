package pattern

import "fmt"

//_______________________________________________________________________
// Интерфейс продукта


type IGun interface {
    setName(name string)
    setPower(power int)
    GetName() string
    GetPower() int
}


//_______________________________________________________________________
// Конкретный продукт

type gun struct {
    name  string
    power int
}

func (g *gun) setName(name string) {
    g.name = name
}

func (g *gun) GetName() string {
    return g.name
}

func (g *gun) setPower(power int) {
    g.power = power
}

func (g *gun) GetPower() int {
    return g.power
}

//_______________________________________________________________________
// Конкретный продукт

type ak47 struct {
    gun
}

func newAk47() IGun {
    return &ak47{
        gun: gun{
            name:  "AK47 gun",
            power: 4,
        },
    }
}
//_______________________________________________________________________
// Конкретный продукт
type musket struct {
    gun
}

func newMusket() IGun {
    return &musket{
        gun: gun{
            name:  "Musket gun",
            power: 1,
        },
    }
}
//_______________________________________________________________________
// Фабрика

func GetGun(gunType string) (IGun, error) {
    if gunType == "ak47" {
        return newAk47(), nil
    }
    if gunType == "musket" {
        return newMusket(), nil
    }
    return nil, fmt.Errorf("wrong gun type passed")
}




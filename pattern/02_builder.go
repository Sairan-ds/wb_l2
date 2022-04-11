package pattern


/* Строитель — это порождающий паттерн проектирования,
который позволяет создавать сложные объекты пошагово.
Строитель даёт возможность использовать один и тот же
код строительства для получения разных
представлений объектов.
Паттерн Строитель предлагает вынести конструирование
объекта за пределы его собственного класса, поручив это
дело отдельным объектам, называемым строителями.
Паттерн предлагает разбить процесс конструирования объ-
екта на отдельные шаги (например, построитьСтены ,
вставитьДвери и другие). Чтобы создать объект, вам нужно
поочерёдно вызывать методы строителя. Причём не нужно
запускать все шаги, а только те, что нужны для производства
объекта определённой конфигурации.*/


//_______________________________________________________________________
// Director
/* Вы можете выделить вызовы методов стро-
ителя в отдельный класс, называемый директором. В этом
случае директор будет задавать порядок шагов строитель-
ства, а строитель — выполнять их. */
type director struct {
	builder iBuilder
}

func NewDirector(b iBuilder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) SetBuilder(b iBuilder) {
	d.builder = b
}

func (d *director) BuildPC() pc {
	d.builder.SetProcessorType()
	d.builder.SetGraphicCardType()
	d.builder.SetNumRum()
	return d.builder.GetPC()
}

//_______________________________________________________________________
// PC структура нашего компьютера

type pc struct {
	ProccessorType string
	GraphicCard   string
	RAM      int
}
//_______________________________________________________________________
// Builder interface объявляет шаги конструирования продуктов, общие для всех видов строителей.

type iBuilder interface {
	SetProcessorType()
	SetGraphicCardType()
	SetNumRum()
	GetPC() pc
}

func GetBuilder(builderType string) iBuilder {
	if builderType == "office" {
		return &officeBuilder{}
	}

	if builderType == "gaming" {
		return &gamingBuilder{}
	}
	return nil
}

//_______________________________________________________________________
// Строитель игрового компьютера

type gamingBuilder struct {
	ProccessorType string
	GraphicCard   string
	RAM      int
}

func NewGamingPCBuilder() *gamingBuilder {
	return &gamingBuilder{}
}

func (b *gamingBuilder) SetProcessorType() {
	b.ProccessorType = "I7 7700 "
}

func (b *gamingBuilder) SetGraphicCardType() {
	b.GraphicCard = "Nvidia Geforce 3080"
}

func (b *gamingBuilder) SetNumRum() {
	b.RAM = 32
}

func (b *gamingBuilder) GetPC() pc {
	return pc{
		ProccessorType: b.ProccessorType,
		GraphicCard: b.GraphicCard,
		RAM: b.RAM,
	}
}

//_______________________________________________________________________
// Строитель оффисного компьютера

type officeBuilder struct {
	ProccessorType string
	GraphicCard   string
	RAM      int
}

func NewOfficePCBuilder() *officeBuilder {
	return &officeBuilder{}
}

func (b *officeBuilder) SetProcessorType() {
	b.ProccessorType = "I5 9600 "
}

func (b *officeBuilder) SetGraphicCardType() {
	b.GraphicCard = "Nvidia Geforce 1060"
}

func (b *officeBuilder) SetNumRum() {
	b.RAM = 16
}

func (b *officeBuilder) GetPC() pc {
	return pc{
		ProccessorType: b.ProccessorType,
		GraphicCard: b.GraphicCard,
		RAM: b.RAM,
	}
}

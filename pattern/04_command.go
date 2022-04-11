package pattern

import "fmt"

//_______________________________________________________________________
// Кнопка, которая отправляет команду 

type Button struct {
	Command Command
}

func (b *Button) Press() {
	b.Command.execute()
}

//_______________________________________________________________________
// command 

type Command interface {
	execute()
}
//_______________________________________________________________________
// device

type Device interface {
	On()
	Off()
}
//_______________________________________________________________________
// device

type OffCommand struct {
	Device Device
}

func (c *OffCommand) execute() {
	c.Device.Off()
}
//_______________________________________________________________________
// device

type OnCommand struct {
	Device Device
}

func (c *OnCommand) execute() {
	c.Device.On()
}

//_______________________________________________________________________
// device

type Tv struct {
	isRunning bool
}

func (t *Tv) On() {
	t.isRunning = true
	fmt.Println("Turning Tv on")
}

func (t *Tv) Off() {
	t.isRunning = false
	fmt.Println("Turning Tv off")
}
package main

import "fmt"

/*
参考资料：GPT
桥接模式是一种设计模式，用于将抽象与其实现分离，使得两者可以独立地变化。
这个例子中，我们将定义一个Device接口作为实现的抽象，并有一个Remote结构体作为“抽象”的一部分。然后我们会有两个不同的设备实现：Tv和Radio。

Device接口定义了所有设备都必须实现的方法。Tv和Radio是Device的具体实现，它们都可以被打开、关闭、调节频道和音量。
Remote则是一个抽象类型，它包含一个Device类型的成员，这样Remote就可以控制任何实现了Device接口的对象，这也就实现了桥接模式。
在main函数中，我们可以看到如何使用这些类型来控制电视和收音机。
*/

// Device 接口定义了设备的一组行为
type Device interface {
	On()
	Off()
	TuneChannel(channel int)
	GetVolume() int
	SetVolume(volume int)
}

// Remote 是一个抽象，包含一个对Device的引用
type Remote struct {
	device Device
}

func (r *Remote) TurnOn() {
	r.device.On()
}

func (r *Remote) TurnOff() {
	r.device.Off()
}

func (r *Remote) SetChannel(channel int) {
	r.device.TuneChannel(channel)
}

func (r *Remote) IncreaseVolume() {
	r.device.SetVolume(r.device.GetVolume() + 1)
}

func (r *Remote) DecreaseVolume() {
	r.device.SetVolume(r.device.GetVolume() - 1)
}

// Tv 是Device接口的一个实现
type Tv struct {
	isRunning bool
	volume    int
	channel   int
}

func (t *Tv) On() {
	t.isRunning = true
	fmt.Println("TV is turned on")
}

func (t *Tv) Off() {
	t.isRunning = false
	fmt.Println("TV is turned off")
}

func (t *Tv) TuneChannel(channel int) {
	t.channel = channel
	fmt.Printf("TV channel set to %d\n", channel)
}

func (t *Tv) GetVolume() int {
	return t.volume
}

func (t *Tv) SetVolume(volume int) {
	t.volume = volume
	fmt.Printf("TV volume set to %d\n", volume)
}

// Radio 是Device接口的另一个实现
type Radio struct {
	isRunning bool
	volume    int
	channel   int
}

func (r *Radio) On() {
	r.isRunning = true
	fmt.Println("Radio is turned on")
}

func (r *Radio) Off() {
	r.isRunning = false
	fmt.Println("Radio is turned off")
}

func (r *Radio) TuneChannel(channel int) {
	r.channel = channel
	fmt.Printf("Radio channel set to %d\n", channel)
}

func (r *Radio) GetVolume() int {
	return r.volume
}

func (r *Radio) SetVolume(volume int) {
	r.volume = volume
	fmt.Printf("Radio volume set to %d\n", volume)
}

func main() {
	tv := &Tv{}
	radio := &Radio{}

	tvRemote := Remote{device: tv}
	radioRemote := Remote{device: radio}

	tvRemote.TurnOn()
	tvRemote.SetChannel(3)
	tvRemote.IncreaseVolume()
	tvRemote.TurnOff()

	radioRemote.TurnOn()
	radioRemote.SetChannel(5)
	radioRemote.IncreaseVolume()
	radioRemote.TurnOff()
}

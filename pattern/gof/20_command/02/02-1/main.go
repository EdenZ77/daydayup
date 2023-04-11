package main

import "fmt"

type Doctor struct{}

func (d *Doctor) treatEye() {
	fmt.Println("医生治疗眼睛")
}

func (d *Doctor) treatNose() {
	fmt.Println("医生治疗鼻子")
}

// CommandTreatEye 治疗眼睛的病单
type CommandTreatEye struct {
	doctor *Doctor
}

func (cmd *CommandTreatEye) Treat() {
	cmd.doctor.treatEye()
}

// CommandTreatNose 治疗鼻子的病单
type CommandTreatNose struct {
	doctor *Doctor
}

func (cmd *CommandTreatNose) Treat() {
	cmd.doctor.treatNose()
}

// 那么可以通过下述方式，新增一个订单模块，将业务方和核心医生模块进行解耦和隔离。
// 病人可以先填写病单，并不会直接和医生进行交互和耦合，医生只对接订单的接口，实现的代码方式如下：
func main() {
	//依赖病单，通过填写病单，让医生看病
	//治疗眼睛的病单
	doctor := new(Doctor)
	cmdEye := CommandTreatEye{doctor}
	cmdEye.Treat() //通过病单来让医生看病

	cmdNose := CommandTreatNose{doctor}
	cmdNose.Treat() //通过病单来让医生看病
}

// 这样Doctor作为核心的消息接受者和计算模块，将和业务模块高耦合，每个业务方都需要直接面向Doctor依赖和编程。
//func main() {
//	doctor := new(Doctor)
//	doctor.treatEye()
//	doctor.treatNose()
//}

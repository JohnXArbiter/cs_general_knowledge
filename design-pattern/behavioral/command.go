package main

import "fmt"

// 病人
func main() {
	//依赖病单，通过填写病单，让医生看病
	doctor := new(Doctor)
	//治疗眼睛的病单
	cmdEye := CommandTreatEye{doctor}
	//治疗鼻子的病单
	cmdNose := CommandTreatNose{doctor}

	//护士
	nurse := new(Nurse)
	//收集管理病单
	nurse.CmdList = append(nurse.CmdList, &cmdEye, &cmdNose)

	//执行病单指令
	nurse.Notify()
}

// Doctor 医生-命令接收者
type Doctor struct{}

func (d *Doctor) treatEye() {
	fmt.Println("医生治疗眼睛")
}

func (d *Doctor) treatNose() {
	fmt.Println("医生治疗鼻子")
}

// Command 抽象的命令
type Command interface {
	Treat()
}

type CommandTreatEye struct {
	doctor *Doctor
}

func (cmd *CommandTreatEye) Treat() {
	cmd.doctor.treatEye()
}

type CommandTreatNose struct {
	doctor *Doctor
}

func (cmd *CommandTreatNose) Treat() {
	cmd.doctor.treatNose()
}

// Nurse 护士-调用命令者
type Nurse struct {
	CmdList []Command
}

func (n *Nurse) Notify() {
	if n.CmdList == nil {
		return
	}
	for _, cmd := range n.CmdList {
		cmd.Treat()
	}
}

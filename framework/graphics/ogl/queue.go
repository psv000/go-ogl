package ogl

type (
	// Queue ...
	Queue struct {
		oglDevice *OpenGL
		commands  []Command

		cmdLen  int
		cmdShft int
	}
)

// NewQueue ...
func NewQueue(oglDevice *OpenGL) *Queue {
	return &Queue{
		oglDevice: oglDevice,
		commands:  make([]Command, CmdLenDefault),
		cmdLen:    CmdLenDefault,
	}
}

// AddCmd ...
func (q *Queue) AddCmd(cmd Command) {
	if q.cmdShft >= q.cmdLen {
		q.commands = append(q.commands, cmd)
		q.cmdLen++
	}
	q.commands[q.cmdShft] = cmd
	q.cmdShft++

}

// Flush ...
func (q *Queue) Flush() {
	q.commands = make([]Command, q.cmdLen)
	q.cmdShft = 0

}

func (q *Queue) Process() {
	for _, cmd := range q.commands {
		switch cmd.Ct {
		case NoCmd:
			return
		case ApplyProgramCmd:
			program, ok := cmd.Args.(uint32)
			checkConversion(ok)
			q.oglDevice.ApplyProgram(program)
		case ApplyUniformsCmd:
			uni, ok := cmd.Args.([]Uniform)
			checkConversion(ok)
			for _, u := range uni {
				q.oglDevice.ApplyUniform(u)
			}
		case DrawMeshCmd:
			drawArgs, ok := cmd.Args.(DrawArgs)
			checkConversion(ok)
			q.oglDevice.Draw(drawArgs.Vao, drawArgs.IndLen)
		case ClearContextCmd:
			q.oglDevice.ClrScr()
		}
	}
}

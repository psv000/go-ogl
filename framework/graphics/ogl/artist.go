package ogl

import (
	"framework/graphics/primitives"
	"framework/graphics/program"
	"framework/graphics/utils"
	"framework/mth"

	"github.com/sirupsen/logrus"
)

type (
	// Artist ...
	Artist struct {
		queue *Queue
	}
)

// NewArtist ...
func NewArtist(queue *Queue) *Artist {
	return &Artist{queue: queue}
}

// ClearScreen ...
func (a *Artist) ClearScreen() {
	a.queue.AddCmd(Command{
		Ct:   ClearContextCmd,
		Args: nil,
	})
}

func (a *Artist) DrawMesh(glMesh interface{}, model, view, projection mth.Mat4f, color mth.Vec4f32) {
	oglMesh, ok := glMesh.(*Mesh)
	utils.CheckConversion(ok)

	for i, u := range oglMesh.uniforms {
		switch u.Dst {
		case program.ModelDst:
			u.Arg = model
		case program.ViewDst:
			u.Arg = view
		case program.ProjectionDst:
			u.Arg = projection
		case program.ColorDst:
			u.Arg = color
		default:
			logrus.Panic("unknown uniform dst")
		}
		oglMesh.uniforms[i] = u
	}

	a.queue.AddCmd(Command{
		Ct:   ApplyProgramCmd,
		Args: oglMesh.program,
	})
	a.queue.AddCmd(Command{
		Ct:   ApplyUniformsCmd,
		Args: oglMesh.uniforms,
	})
	a.queue.AddCmd(Command{
		Ct: DrawMeshCmd,
		Args: DrawArgs{
			Vao:    oglMesh.vao,
			IndLen: oglMesh.indicesLen,
		},
	})
}

// DrawMeshGroup ...
func (a *Artist) DrawMeshGroup(group *primitives.MeshGroup, view, projection mth.Mat4f){
	lu := group.GPUPack.Uniforms[program.UCLight]
	lm := group.GPUPack.Uniforms[program.UCMesh]

	var uniforms []program.Uniform
	for _, l := range group.LightSources {
		for i, u := range lu {
			switch u.Dst {
			case program.LightPosDst:
				u.Arg = l.Pos
			case program.LightColorDst:
				u.Arg = l.Col
			}
			lu[i] = u

		}
	}
	uniforms = append(uniforms, lu...)

	for _, m := range group.Meshes {
		oglMesh, ok := m.Gl().(*Mesh)
		utils.CheckConversion(ok)
		for i, u := range lm {
			switch u.Dst {
			case program.ModelDst:
				u.Arg = m.Node().Update()
			case program.ViewDst:
				u.Arg = view
			case program.ProjectionDst:
				u.Arg = projection
			case program.ColorDst:
				u.Arg = m.Color()
			default:
				logrus.Panic("unknown uniform dst")
			}
			lm[i] = u
		}

		a.queue.AddCmd(Command{
			Ct:   ApplyProgramCmd,
			Args: group.GPUPack.ID,
		})
		a.queue.AddCmd(Command{
			Ct:   ApplyUniformsCmd,
			Args: append(uniforms, lm...),
		})
		a.queue.AddCmd(Command{
			Ct: DrawMeshCmd,
			Args: DrawArgs{
				Vao:    oglMesh.vao,
				IndLen: oglMesh.indicesLen,
			},
		})
	}
}
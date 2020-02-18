package ogl

import (
	"framework/graphics"
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
	checkConversion(ok)

	for i, u := range oglMesh.uniforms {
		switch u.dst {
		case ModelDst:
			u.arg = model
		case ViewDst:
			u.arg = view
		case ProjectionDst:
			u.arg = projection
		case ColorDst:
			u.arg = color
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
func (a *Artist) DrawMeshGroup(group graphics.MeshGroup, view, projection mth.Mat4f){

	lu := group.GPUPack.Uniforms[UCLight]
	lm := group.GPUPack.Uniforms[UCMesh]

	var uniforms []Uniform
	for _, l := range group.LightSources {
		for i, u := range lu {
			switch u.dst {
			case LightPosDst:
				u.arg = l.Pos
			case LightColorDst:
				u.arg = l.Col
			}
			lu[i] = u

		}
	}
	uniforms = append(uniforms, lu...)

	for _, m := range group.Meshes {
		oglMesh, ok := m.Gl().(*Mesh)
		checkConversion(ok)
		for _, u := range lm {
			switch u.dst {
			case ModelDst:
				u.arg = m.Node().Update()
			case ViewDst:
				u.arg = view
			case ProjectionDst:
				u.arg = projection
			case ColorDst:
				u.arg = m.Color()
			default:
				logrus.Panic("unknown uniform dst")
			}
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
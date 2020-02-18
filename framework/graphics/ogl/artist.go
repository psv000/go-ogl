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
		case LightPosDst:
			u.arg = mth.Vec3f32{0., 0., 0.}
		case LightColorDst:
			u.arg = mth.Vec3f32{1., 1., 1.}
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

func (a *Artist) DrawMeshGroup(
	group graphics.MeshGroup,
	model, view, projection mth.Mat4f,
	color mth.Vec4f32,
	ls []graphics.LightSource){

	for _, m := range group.Meshes {

		m.Node()
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

}
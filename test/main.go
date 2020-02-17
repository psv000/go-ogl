package main

import (
	"fmt"
	"framework/mth"
	"log"
	"math"
	"runtime"
	"strings"
	"unsafe"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	width  = 800
	height = 600

	vertexShaderSource = `
		#version 410
		layout (location = 0) in vec3 vp;
		layout (location = 1) in vec4 color;

		out vec4 vColor;

		uniform mat4 mvp;

		void main() {
			gl_Position = mvp * vec4(vp, 1.0);
			vColor = color;
		}
	` + "\x00"

	fragmentShaderSource = `
		#version 410

		out vec4 frag_color;
		in vec4 vColor;
		void main() {
			frag_color = vColor;
		}
	` + "\x00"
)

type pack struct {
	p [3]float32
	c [4]uint8
}

var loc int32

var fovY = 45. * math.Pi / 180.
var aspect = float64(width) / float64(height)
var zFar = 1000.
var zNear = 0.1
var view = mth.NewUnitMat4f()
var p = mth.PerspectiveRH(fovY, aspect, zNear, zFar)

//var mvp = mth.OrthoRH(0, width, 0, height, zNear, zFar)

func main() {
	runtime.LockOSThread()

	window := initGlfw()
	defer glfw.Terminate()
	program := initOpenGL()

	vao := makeVao()
	for !window.ShouldClose() {
		draw(vao, window, program)
	}
}

var dt float64

func draw(vao uint32, window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	loc = gl.GetUniformLocation(program, gl.Str("mvp"+"\x00"))

	//view.Set(3, 0, -400)
	//view.Set(3, 1, -300)
	view.Set(3, 2, -(math.Sin(dt*4)+1.)*10-2)
	dt += 0.016

	m := mth.NewUnitMat4f()
	res := m.Mult(view).Mult(p)
	vals := res.Values()
	gl.UniformMatrix4fv(loc, 1, false, &vals[0])

	gl.BindVertexArray(vao)
	gl.DrawElements(gl.TRIANGLES, 3, gl.UNSIGNED_INT, gl.PtrOffset(0))
	gl.BindVertexArray(0)

	glfw.PollEvents()

	window.SwapBuffers()
}

// initGlfw initializes glfw and returns a Window to use.
func initGlfw() *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, "test gl app", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	return window
}

// initOpenGL initializes OpenGL and returns an intiialized program.
func initOpenGL() uint32 {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	gl.Enable(gl.DEPTH_TEST)

	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	prog := gl.CreateProgram()
	gl.AttachShader(prog, vertexShader)
	gl.AttachShader(prog, fragmentShader)
	gl.LinkProgram(prog)
	return prog
}

// makeVao initializes and returns a vertex array from the points provided.
func makeVao() uint32 {
	var vbo uint32

	var packed = []pack{
		{[3]float32{0, 0.5, 0}, [4]uint8{255, 0., 0., 255}},
		{[3]float32{-0.5, -0.5, 0}, [4]uint8{0., 255, 0., 255}},
		{[3]float32{0.5, -0.5, 0}, [4]uint8{0., 0., 255, 255}},
	}
	//var triangle = []pack{
	//	{[3]float32{400, 450, 0.}, [4]uint8{255, 0., 0., 255}},
	//	{[3]float32{200, 150, 0.}, [4]uint8{0., 255, 0., 255}},
	//	{[3]float32{600, 150, 0.}, [4]uint8{0., 0., 255, 255}},
	//}

	var triangle = []pack{
		{[3]float32{-0.5, -0.5, 0.}, [4]uint8{255, 0., 0., 255}},
		{[3]float32{0.5, -0.5, 0.}, [4]uint8{0., 255, 0., 255}},
		{[3]float32{0, 0.5, 0.}, [4]uint8{0., 0., 255, 255}},
	}

	packed = triangle

	var indices = []uint32{
		0, 1, 2,
	}

	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)

	var ptr interface{} = packed
	var siz = int(unsafe.Sizeof(packed[0])) * len(packed)

	gl.BufferData(gl.ARRAY_BUFFER, siz, gl.Ptr(ptr), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	var ibo uint32
	gl.GenBuffers(1, &ibo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ibo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, 4*len(indices), gl.Ptr(indices), gl.STATIC_DRAW)

	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, (4*3)+(1*4), nil)
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(1, 4, gl.UNSIGNED_BYTE, true, (4*3)+(1*4), gl.PtrOffset((4 * 3)))
	gl.EnableVertexAttribArray(1)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ibo)

	return vao
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}

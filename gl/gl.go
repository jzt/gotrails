package gl

import (
	"log"

	GL "github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"

	"github.com/jzt/gotrails/gl/shader"
)

const (
	// TODO(jzt): put shaders in a file
	vertexShaderSource = `
		#version 410
		in vec3 vp;
		void main() {
			gl_Position = vec4(vp, 1.0);
		}
	` + "\x00"

	fragmentShaderSource = `
		#version 410
		out vec4 frag_color;
		void main() {
			frag_color = vec4(1,1,1,1);
		}
		` + "\x00"
)

// Initglfw initializes and returns a window
func InitGlfw(width, height int) *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, "RENAME ME", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	return window
}

// GLFWTerminate
func GLFWTerminate() {
	glfw.Terminate()
}

func InitOpenGL() uint32 {
	if err := GL.Init(); err != nil {
		panic(err)
	}

	version := GL.GoStr(GL.GetString(GL.VERSION))
	log.Println("OpenGL version", version)

	vertexShader, err := shader.CompileShader(vertexShaderSource, GL.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}

	fragmentShader, err := shader.CompileShader(fragmentShaderSource, GL.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	prog := GL.CreateProgram()
	GL.AttachShader(prog, vertexShader)
	GL.AttachShader(prog, fragmentShader)
	GL.LinkProgram(prog)
	return prog
}

func MakeVao(points []float32) uint32 {
	var vbo uint32
	GL.GenBuffers(1, &vbo)
	GL.BindBuffer(GL.ARRAY_BUFFER, vbo)
	GL.BufferData(GL.ARRAY_BUFFER, 4*len(points), GL.Ptr(points), GL.STATIC_DRAW)

	var vao uint32
	GL.GenVertexArrays(1, &vao)
	GL.BindVertexArray(vao)
	GL.EnableVertexAttribArray(0)
	GL.BindBuffer(GL.ARRAY_BUFFER, vbo)
	GL.VertexAttribPointer(0, 3, GL.FLOAT, false, 0, nil)

	return vao
}

// PollEvents
func PollEvents() {
	glfw.PollEvents()
}

package mesh

import (
	"log"

	GL "github.com/go-gl/gl/v4.1-core/gl"

	"github.com/jzt/gotrails/gl"
)

var vertices = []float32{
	0, 0.5, 0, // top
	-0.5, -0.5, 0, // left
	0.5, -0.5, 0, // right
}

type TriangleMesh struct {
	vao, program uint32
}

// Render implements the mesh.Render() interface
func (t *TriangleMesh) Render() {
	GL.UseProgram(t.program)
	GL.BindVertexArray(t.vao)
	GL.DrawArrays(GL.TRIANGLES, 0, int32(len(vertices)/3))
}

func Triangle(program uint32) *TriangleMesh {
	log.Println("Creating triangle")
	return &TriangleMesh{
		vao:     gl.MakeVao(vertices),
		program: program,
	}
}

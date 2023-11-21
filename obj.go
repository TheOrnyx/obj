package obj

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

//basic vector with x, y, z coords
type Vertex struct {
	X, Y, Z float32
}

//basic face containing vertex indexes
type Face struct {
	Vertices []int
}

//the object parsed from the obj
type Object struct {
	Vertices  []Vertex
	Normals   []Vertex
	TexCoords []Vertex
	Faces []Face
}

// OpenFile open and parse a .obj file
func OpenFile(path string) (*Object, error) {
	newModel := new(Object)

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		err := processLine(line, newModel)
		if err != nil {
			return nil, err
		}
	}

	return newModel, nil
}

// process a line
func processLine(line string, m *Object) error {
	words := strings.Fields(line)
	if len(words) < 1 {
		return nil
	}

	switch words[0] {
	case "v":
		x, err := strconv.ParseFloat(words[1], 32)
		if err != nil {
			return err
		}
		y, err := strconv.ParseFloat(words[2], 32)
		if err != nil {
			return err
		}
		z, err := strconv.ParseFloat(words[3], 32)
		if err != nil {
			return err
		}
		m.Vertices = append(m.Vertices, Vertex{float32(x), float32(y), float32(z)})

	case "vt": //maybe handle like missing ones or smth
		u, err := strconv.ParseFloat(words[1], 32)
		if err != nil {
			return err
		}
		v, err := strconv.ParseFloat(words[2], 32)
		if err != nil {
			return err
		}
		m.TexCoords = append(m.TexCoords, Vertex{float32(u), float32(v), 0.0})

	case "vn":
		x, err := strconv.ParseFloat(words[1], 32)
		if err != nil {
			return err
		}
		y, err := strconv.ParseFloat(words[2], 32)
		if err != nil {
			return err
		}
		z, err := strconv.ParseFloat(words[3], 32)
		if err != nil {
			return err
		}
		m.Normals = append(m.Normals, Vertex{float32(x), float32(y), float32(z)})

	case "f":
		var face Face
		for _, vertexStr := range words[1:] {
			vertexIndices := strings.Split(vertexStr, "/")
			vertexIndex, _ := strconv.Atoi(vertexIndices[0])
			face.Vertices = append(face.Vertices, vertexIndex-1) // Obj indices are 1-based
		}
		m.Faces = append(m.Faces, face)
	}

	return nil
}

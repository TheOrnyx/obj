package obj

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Model struct {
	Vertices  []float32
	TexCoords []float32
	Normals   []float32
}

// OpenFile open and parse a .obj file
func OpenFile(path string) (*Model, error) {
	newModel := new(Model)

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
func processLine(line string, m *Model) error {
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
		m.Vertices = append(m.Vertices, float32(x), float32(y), float32(z))

	case "vt": //maybe handle like missing ones or smth
		u, err := strconv.ParseFloat(words[1], 32)
		if err != nil {
			return err
		}
		v, err := strconv.ParseFloat(words[2], 32)
		if err != nil {
			return err
		}
		m.TexCoords = append(m.TexCoords, float32(u), float32(v))

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
		m.Normals = append(m.Normals, float32(x), float32(y), float32(z))
	}

	return nil
}

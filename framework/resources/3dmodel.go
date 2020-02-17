package resources

import (
	"bufio"
	"fmt"
	"framework/mth"
	"os"
	"strings"

	"framework/graphics/ogl"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	threeDModelInfoTag = loaderInfoTag + ": 3model"

	vertHeader = "v "
	uvsHeader  = "vt "
	normHeader = "vn "
	indHeader  = "f "
)

// Load3ModelObj ...
func Load3ModelObj(filename string) (Obj3DModel, error) {
	var result Obj3DModel

	f, err := os.Open(filename)
	if err != nil {
		return Obj3DModel{}, errors.Wrap(err, threeDModelInfoTag)
	}
	defer func() {
		if cErr := f.Close(); cErr != nil {
			logrus.Panic(errors.Wrap(cErr, threeDModelInfoTag))
		}
	}()

	var vert [3]float64
	var uvs [2]float64
	var norm [3]float64
	var vertInd [3]uint32
	var uvsInd [3]uint32
	var normInd [3]uint32
	var header string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case strings.HasPrefix(line, vertHeader):
			_, err := fmt.Sscanf(line, "%s %f %f %f", &header, &vert[0], &vert[1], &vert[2])
			if err != nil {
				return Obj3DModel{}, errors.Wrap(err, threeDModelInfoTag)
			}
			result.Vert = append(result.Vert, vert)
		case strings.HasPrefix(line, uvsHeader):
			_, err := fmt.Sscanf(line, "%s %f %f", &header, &uvs[0], &uvs[1])
			if err != nil {
				return Obj3DModel{}, errors.Wrap(err, threeDModelInfoTag)
			}
			result.UVS = append(result.UVS, uvs)
		case strings.HasPrefix(line, normHeader):
			_, err := fmt.Sscanf(line, "%s %f %f %f", &header, &norm[0], &norm[1], &norm[2])
			if err != nil {
				return Obj3DModel{}, errors.Wrap(err, threeDModelInfoTag)
			}
			result.Norm = append(result.Norm, norm)
		case strings.HasPrefix(line, indHeader):
			matches, err := fmt.Sscanf(line, "%s %d/%d/%d %d/%d/%d %d/%d/%d",
				&header,
				&vertInd[0], &uvsInd[0], &normInd[0],
				&vertInd[1], &uvsInd[1], &normInd[1],
				&vertInd[2], &uvsInd[2], &normInd[2])
			if err != nil {
				return Obj3DModel{}, errors.Wrap(err, threeDModelInfoTag)
			}
			if matches != 10 {
				return Obj3DModel{}, errors.New(threeDModelInfoTag + ": file can't be read, try exporting with other options")
			}
			for i := 0; i < 3; i++ {
				vertInd[i]--
				uvsInd[i]--
				normInd[i]--

				result.Ind = append(result.Ind, [IndLen]uint32{vertInd[i], uvsInd[i], normInd[i]})
			}
		}
	}
	return result, nil
}

// ConvertObjToOGL ...
func ConvertObjToOGL(model Obj3DModel) ([]ogl.V3fUV2fN3f, []uint32) {
	var result []ogl.V3fUV2fN3f
	var indices []uint32

	result = make([]ogl.V3fUV2fN3f, len(model.Ind))
	for i, v := range model.Ind {
		result[i] = ogl.V3fUV2fN3f{
			Point: mth.Vec3f32{
				float32(model.Vert[v[VertInd]][0]),
				float32(model.Vert[v[VertInd]][1]),
				float32(model.Vert[v[VertInd]][2]),
			},
			UVS: mth.Vec2f32{
				float32(model.UVS[v[UVSInd]][0]),
				float32(model.UVS[v[UVSInd]][1]),
			},
			Norm: mth.Vec3f32{
				float32(model.Norm[v[NormInd]][0]),
				float32(model.Norm[v[NormInd]][1]),
			},
		}
		indices = append(indices, uint32(i))
	}
	return result, indices
}

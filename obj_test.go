package obj

import (
	"testing"
	"reflect"
)

func TestOpenFile(t *testing.T) {
	testCases := []struct {
		name     string
		filePath string
		expected *Object
		expectErr bool
	}{
		{
			name:     "ValidFile",
			filePath: "test.obj", // replace with the path to a valid .obj file for testing
			expected: &Object{
				Vertices:  []float32{1.0, 2.0, 3.0, 4.0, 5.0, 6.0},
				TexCoords: []float32{0.1, 0.2, 0.3, 0.4},
				Normals:   []float32{0.5, 0.6, 0.7, 0.8, 0.9, 1.0},
			},
			expectErr: false,
		},
		// Add more test cases for different scenarios
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := OpenFile(tc.filePath)

			if tc.expectErr && err == nil {
				t.Errorf("Expected an error but got nil")
			}

			if !tc.expectErr && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Result does not match the expected output.\nExpected: %v\nGot: %v", tc.expected, result)
			}
			t.Logf("Final Output: \n%v", result)
		})
	}
}

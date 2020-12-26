package got

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRepository(t *testing.T) {
	type field struct {
		Name   string
		DBName string
		GoType string
		Type   string
	}
	data := struct {
		ImportPath    string
		ModuleName    string
		ModelName     string
		ImportModules []string
		Fields        []field
	}{
		ImportPath:    "github.com/test/repository",
		ModuleName:    "repository",
		ModelName:     "User",
		ImportModules: []string{},
		Fields: []field{
			{Name: "Id", DBName: "id", GoType: "int", Type: "number"},
			{Name: "Name", DBName: "name", GoType: "string", Type: "string"},
		},
	}
	err := GenerateRepository(data)
	assert.Nil(t, err)
	if err != nil {
		t.Errorf("%+v", err)
	}
}

package template_demo

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"html/template"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	type User struct {
		Name string
	}

	tpl := template.New("hello-world")
	tpl, err := tpl.Parse(`Hello, {{ .Name}}`)
	require.NoError(t, err)

	buffer := &bytes.Buffer{}
	err = tpl.Execute(buffer, User{Name: "Tom"})
	require.NoError(t, err)
	assert.Equal(t, `Hello, Tom`, buffer.String())
}

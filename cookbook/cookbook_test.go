package cookbook_test

import (
	"testing"

	"github.com/mlafeldt/chef-runner/cookbook"
	"github.com/stretchr/testify/assert"
)

func TestNewCookbook(t *testing.T) {
	cb, err := cookbook.NewCookbook("../testdata")
	assert.NoError(t, err)
	if assert.NotNil(t, cb) {
		assert.Equal(t, "../testdata", cb.Path)
		assert.Equal(t, "practicingruby", cb.Name)
		assert.Equal(t, "1.3.1", cb.Version)
	}
}

func TestNewCookbook_WithoutMetadata(t *testing.T) {
	cb, err := cookbook.NewCookbook(".")
	assert.NoError(t, err)
	if assert.NotNil(t, cb) {
		assert.Equal(t, ".", cb.Path)
		assert.Equal(t, "", cb.Name)
		assert.Equal(t, "", cb.Version)
	}
}

func TestString(t *testing.T) {
	cb := cookbook.Cookbook{Name: "cats", Version: "1.2.3"}
	assert.Equal(t, "cats 1.2.3", cb.String())
}

func TestFiles(t *testing.T) {
	cb, _ := cookbook.NewCookbook("../testdata")
	expect := []string{
		"../testdata/README.md",
		"../testdata/metadata.rb",
		"../testdata/attributes",
		"../testdata/recipes",
	}
	assert.Equal(t, expect, cb.Files())
}

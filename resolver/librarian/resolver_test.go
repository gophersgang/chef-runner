package librarian_test

import (
	"testing"

	"github.com/mlafeldt/chef-runner/resolver"
	. "github.com/mlafeldt/chef-runner/resolver/librarian"
	"github.com/stretchr/testify/assert"
)

func TestResolverInterface(t *testing.T) {
	assert.Implements(t, (*resolver.Resolver)(nil), new(Resolver))
}

func TestCommand(t *testing.T) {
	expect := []string{"librarian-chef", "install", "--path", "a/b/c"}
	actual := Command("a/b/c")
	assert.Equal(t, expect, actual)
}

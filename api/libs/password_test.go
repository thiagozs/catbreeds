package libs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPasswordGen(t *testing.T) {

	passwd := NewPasswordGen()
	gen, err := passwd.Generate("abcd")
	if err != nil {
		t.Error(err)
	}
	require.NoError(t, err)

	match, err := passwd.Compare("abcd", gen)
	if !match || err != nil {
		t.Error(err)
	}
	require.NoError(t, err)
	require.True(t, match)

	match, err = passwd.Compare("123", gen)
	if match || err != nil {
		t.Error(err)
	}
	require.NoError(t, err)
	require.False(t, match)

}

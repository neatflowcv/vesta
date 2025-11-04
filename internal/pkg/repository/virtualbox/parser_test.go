package virtualbox_test

import (
	_ "embed"
	"testing"

	"github.com/neatflowcv/vesta/internal/pkg/domain"
	"github.com/neatflowcv/vesta/internal/pkg/repository/virtualbox"
	"github.com/stretchr/testify/require"
)

//go:embed list.long
var content []byte

func TestParser_Parse(t *testing.T) {
	t.Parallel()

	parser := virtualbox.NewParser()

	ret := parser.Parse(content)

	require.Len(t, ret, 2)
	require.Equal(t, "archlinux", ret[0].Name())
	require.Equal(t, "archlinux Clone", ret[1].Name())
	require.Equal(t, domain.InstanceStatusStopped, ret[0].Status())
	require.Equal(t, domain.InstanceStatusStopped, ret[1].Status())
	require.Equal(t, "d3854f3b-225f-431e-9a9b-83ef36718ca3", ret[0].ID())
	require.Equal(t, "544649c9-6c73-4f47-8664-4ebbfe96d333", ret[1].ID())
}

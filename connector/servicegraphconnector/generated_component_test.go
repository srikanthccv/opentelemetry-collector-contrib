// Code generated by mdatagen. DO NOT EDIT.

package servicegraphconnector

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componenttest"

	"go.opentelemetry.io/collector/connector"
	"go.opentelemetry.io/collector/connector/connectortest"
	"go.opentelemetry.io/collector/consumer/consumertest"

	"go.opentelemetry.io/collector/confmap/confmaptest"
)

// assertNoErrorHost implements a component.Host that asserts that there were no errors.
type assertNoErrorHost struct {
	component.Host
	*testing.T
}

var _ component.Host = (*assertNoErrorHost)(nil)

// newAssertNoErrorHost returns a new instance of assertNoErrorHost.
func newAssertNoErrorHost(t *testing.T) component.Host {
	return &assertNoErrorHost{
		componenttest.NewNopHost(),
		t,
	}
}

func (aneh *assertNoErrorHost) ReportFatalError(err error) {
	assert.NoError(aneh, err)
}

func TestComponentLifecycle(t *testing.T) {
	factory := NewFactory()

	tests := []struct {
		name     string
		createFn func(ctx context.Context, set connector.CreateSettings, cfg component.Config) (component.Component, error)
	}{

		{
			name: "traces_to_metrics",
			createFn: func(ctx context.Context, set connector.CreateSettings, cfg component.Config) (component.Component, error) {
				return factory.CreateTracesToMetrics(ctx, set, cfg, consumertest.NewNop())
			},
		},
	}

	cm, err := confmaptest.LoadConf("metadata.yaml")
	require.NoError(t, err)
	cfg := factory.CreateDefaultConfig()
	sub, err := cm.Sub("tests::config")
	require.NoError(t, err)
	require.NoError(t, component.UnmarshalConfig(sub, cfg))

	for _, test := range tests {
		t.Run(test.name+"-shutdown", func(t *testing.T) {
			c, err := test.createFn(context.Background(), connectortest.NewNopCreateSettings(), cfg)
			require.NoError(t, err)
			err = c.Shutdown(context.Background())
			require.NoError(t, err)
		})

		t.Run(test.name+"-lifecycle", func(t *testing.T) {

			firstConnector, err := test.createFn(context.Background(), connectortest.NewNopCreateSettings(), cfg)
			require.NoError(t, err)
			host := newAssertNoErrorHost(t)
			require.NoError(t, err)
			require.NoError(t, firstConnector.Start(context.Background(), host))
			require.NoError(t, firstConnector.Shutdown(context.Background()))
			secondConnector, err := test.createFn(context.Background(), connectortest.NewNopCreateSettings(), cfg)
			require.NoError(t, err)
			require.NoError(t, secondConnector.Start(context.Background(), host))
			require.NoError(t, secondConnector.Shutdown(context.Background()))
		})
	}
}

package router_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trevatk/block-broker/internal/adapter/logging"
	"github.com/trevatk/block-broker/internal/adapter/port/http/router"
	"github.com/trevatk/block-broker/internal/adapter/setup"
	"github.com/trevatk/block-broker/internal/core/domain"
)

func init() {
	_ = os.Setenv("SERVER_HTTP_PORT", "8080")
}

func Test_NewRouter(t *testing.T) {
	t.Run("provider", func(t *testing.T) {

		assert := assert.New(t)

		ctx := context.TODO()

		logger, err := logging.NewLogger()
		assert.NoError(err)

		cfg := setup.NewConfig()
		assert.NoError(setup.ProcessConfigWithEnv(ctx, cfg))

		mockMessenger := domain.NewMockMessenger(t)

		s := router.NewRouter(logger, cfg, mockMessenger)
		assert.NotNil(s)
	})
}

package source_test

import (
	"fmt"
	"testing"

	"github.com/rommms07/idream-erp/config"
	"github.com/rommms07/idream-erp/helpers/source"
	"github.com/stretchr/testify/assert"
)

// Test_mustLoadMockAppConfigFromMcoksFolder asserts whether or not the source.AppConfig does its work properly,
// the first thing it does is it loads the app_config.json from the tests/mocks folder and then asserts all its
// defined fields.
func Test_mustLoadMockAppConfigFromMocksFolder(t *testing.T) {
	bak := config.DEFAULT
	config.DEFAULT = fmt.Sprintf("%s/tests/mocks/app_config.json", config.ROOTDIR)

	conf := source.AppConfig()

	assert.Equal(t, conf.Version.Release, "testing", "Did not match the expected version.")
	assert.Equal(t, conf.Message, "This message is coming from the mocks/app_config.json", "Did not match the expected message.")
	config.DEFAULT = bak
}

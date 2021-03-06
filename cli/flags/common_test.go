package flags

import (
	"path/filepath"
	"testing"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

func TestCommonOptionsInstallFlags(t *testing.T) {
	flags := pflag.NewFlagSet("testing", pflag.ContinueOnError)
	opts := NewCommonOptions()
	opts.InstallFlags(flags)

	err := flags.Parse([]string{
		"--tlscacert=\"/foo/cafile\"",
		"--tlscert=\"/foo/cert\"",
		"--tlskey=\"/foo/key\"",
	})
	assert.NoError(t, err)
	assert.Equal(t, "/foo/cafile", opts.TLSOptions.CAFile)
	assert.Equal(t, "/foo/cert", opts.TLSOptions.CertFile)
	assert.Equal(t, opts.TLSOptions.KeyFile, "/foo/key")
}

func defaultPath(filename string) string {
	return filepath.Join(ConfigurationDir(), filename)
}

func TestCommonOptionsInstallFlagsWithDefaults(t *testing.T) {
	flags := pflag.NewFlagSet("testing", pflag.ContinueOnError)
	opts := NewCommonOptions()
	opts.InstallFlags(flags)

	err := flags.Parse([]string{})
	assert.NoError(t, err)
	assert.Equal(t, defaultPath("ca.pem"), opts.TLSOptions.CAFile)
	assert.Equal(t, defaultPath("cert.pem"), opts.TLSOptions.CertFile)
	assert.Equal(t, defaultPath("key.pem"), opts.TLSOptions.KeyFile)
}

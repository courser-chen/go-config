package cfg

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

/*
* @Author: jensen.chen
* @Date:   2025/4/28 15:23
 */
func TestCfgLoad(t *testing.T) {
	pwd, _ := os.Getwd()
	Load(filepath.Join(pwd, "test_cfg.yml"))
	val, err := GetConfig("cli.name")
	assert.Nil(t, err)
	assert.Equal(t, val, "test")
}

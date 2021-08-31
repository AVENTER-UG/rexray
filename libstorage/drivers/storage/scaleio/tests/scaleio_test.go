package scaleio

import (
	"testing"

	apitests "github.com/AVENTER-UG/rexray/libstorage/api/tests"

	// load the driver packages
	"github.com/AVENTER-UG/rexray/libstorage/drivers/storage/scaleio"
	_ "github.com/AVENTER-UG/rexray/libstorage/drivers/storage/scaleio/storage"
)

func TestSuite(t *testing.T) {
	apitests.RunSuite(t, scaleio.Name)
}

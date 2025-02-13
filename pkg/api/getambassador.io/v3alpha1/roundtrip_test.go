package v3alpha1

import (
	"encoding/json"
	"io/ioutil"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"sigs.k8s.io/yaml"
)

func TestAuthSvcRoundTrip(t *testing.T) {
	var a []AuthService
	checkRoundtrip(t, "authsvc.yaml", &a)
}

func TestDevPortalRoundTrip(t *testing.T) {
	var d []DevPortal
	checkRoundtrip(t, "devportals.yaml", &d)
}

func TestHostRoundTrip(t *testing.T) {
	var h []Host
	checkRoundtrip(t, "hosts.yaml", &h)
}

func TestLogSvcRoundTrip(t *testing.T) {
	var l []LogService
	checkRoundtrip(t, "logsvc.yaml", &l)
}

func TestMappingRoundTrip(t *testing.T) {
	var m []Mapping
	checkRoundtrip(t, "mappings.yaml", &m)
}

func TestModuleRoundTrip(t *testing.T) {
	var m []Module
	checkRoundtrip(t, "modules.yaml", &m)
}

func TestRateLimitSvcRoundTrip(t *testing.T) {
	var r []RateLimitService
	checkRoundtrip(t, "ratelimitsvc.yaml", &r)
}

func TestTCPMappingRoundTrip(t *testing.T) {
	var tm []TCPMapping
	checkRoundtrip(t, "tcpmappings.yaml", &tm)
}

func TestTLSContextRoundTrip(t *testing.T) {
	var tc []TLSContext
	checkRoundtrip(t, "tlscontexts.yaml", &tc)
}

func TestTracingSvcRoundTrip(t *testing.T) {
	var tr []TracingService
	checkRoundtrip(t, "tracingsvc.yaml", &tr)
}

func checkRoundtrip(t *testing.T, filename string, ptr interface{}) {
	bytes, err := ioutil.ReadFile(path.Join("testdata", filename))
	require.NoError(t, err)

	canonical := func() string {
		var untyped interface{}
		err = yaml.Unmarshal(bytes, &untyped)
		require.NoError(t, err)
		canonical, err := json.MarshalIndent(untyped, "", "\t")
		require.NoError(t, err)
		return string(canonical)
	}()

	actual := func() string {
		// Round-trip twice, to get map field ordering, instead of struct field ordering.

		// first
		err = yaml.UnmarshalStrict(bytes, ptr)
		require.NoError(t, err)
		first, err := json.Marshal(ptr)
		require.NoError(t, err)

		// second
		var untyped interface{}
		err = json.Unmarshal(first, &untyped)
		require.NoError(t, err)
		second, err := json.MarshalIndent(untyped, "", "\t")
		require.NoError(t, err)

		return string(second)
	}()

	assert.Equal(t, canonical, actual)
}

package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/aiven/aiven-operator/api/v1alpha1"
	cassandrauserconfig "github.com/aiven/aiven-operator/api/v1alpha1/userconfig/service/cassandra"
)

func getCassandraYaml(project, name, cloudName string) string {
	return fmt.Sprintf(`
apiVersion: aiven.io/v1alpha1
kind: Cassandra
metadata:
  name: %[2]s
spec:
  authSecretRef:
    name: aiven-token
    key: token

  project: %[1]s
  cloudName: %[3]s
  plan: startup-4
  disk_space: 450Gib

  tags:
    env: test
    instance: foo

  userConfig:
    migrate_sstableloader: true
    public_access:
      prometheus: true
    ip_filter:
      - network: 0.0.0.0/32
        description: bar
      - network: 10.20.0.0/16

`, project, name, cloudName)
}

func TestCassandra(t *testing.T) {
	t.Parallel()
	defer recoverPanic(t)

	// GIVEN
	ctx := context.Background()
	name := randName("cassandra")
	yml := getCassandraYaml(testProject, name, testPrimaryCloudName)
	s := NewSession(k8sClient, avnClient, testProject)

	// Cleans test afterwards
	defer s.Destroy()

	// WHEN
	// Applies given manifest
	require.NoError(t, s.Apply(yml))

	// Waits kube objects
	cs := new(v1alpha1.Cassandra)
	require.NoError(t, s.GetRunning(cs, name))

	// THEN
	csAvn, err := avnClient.Services.Get(ctx, testProject, name)
	require.NoError(t, err)
	assert.Equal(t, csAvn.Name, cs.GetName())
	assert.Equal(t, "RUNNING", cs.Status.State)
	assert.Equal(t, csAvn.State, cs.Status.State)
	assert.Equal(t, csAvn.Plan, cs.Spec.Plan)
	assert.Equal(t, csAvn.CloudName, cs.Spec.CloudName)
	assert.Equal(t, "450Gib", cs.Spec.DiskSpace)
	assert.Equal(t, 460800, csAvn.DiskSpaceMB)
	assert.Equal(t, map[string]string{"env": "test", "instance": "foo"}, cs.Spec.Tags)
	csResp, err := avnClient.ServiceTags.Get(ctx, testProject, name)
	require.NoError(t, err)
	assert.Equal(t, csResp.Tags, cs.Spec.Tags)

	// UserConfig test
	require.NotNil(t, cs.Spec.UserConfig)
	require.NotNil(t, cs.Spec.UserConfig.PublicAccess)
	assert.Equal(t, anyPointer(true), cs.Spec.UserConfig.PublicAccess.Prometheus)
	assert.Equal(t, anyPointer(true), cs.Spec.UserConfig.MigrateSstableloader)

	// Validates ip filters
	require.Len(t, cs.Spec.UserConfig.IpFilter, 2)

	// First entry
	assert.Equal(t, "0.0.0.0/32", cs.Spec.UserConfig.IpFilter[0].Network)
	assert.Equal(t, "bar", *cs.Spec.UserConfig.IpFilter[0].Description)

	// Second entry
	assert.Equal(t, "10.20.0.0/16", cs.Spec.UserConfig.IpFilter[1].Network)
	assert.Nil(t, cs.Spec.UserConfig.IpFilter[1].Description)

	// Compares with Aiven ip_filter
	var ipFilterAvn []*cassandrauserconfig.IpFilter
	require.NoError(t, castInterface(csAvn.UserConfig["ip_filter"], &ipFilterAvn))
	assert.Equal(t, ipFilterAvn, cs.Spec.UserConfig.IpFilter)

	// Secrets test
	secret, err := s.GetSecret(cs.GetName())
	require.NoError(t, err)
	assert.NotEmpty(t, secret.Data["CASSANDRA_HOST"])
	assert.NotEmpty(t, secret.Data["CASSANDRA_PORT"])
	assert.NotEmpty(t, secret.Data["CASSANDRA_USER"])
	assert.NotEmpty(t, secret.Data["CASSANDRA_PASSWORD"])
	assert.NotEmpty(t, secret.Data["CASSANDRA_URI"])
	assert.NotEmpty(t, secret.Data["CASSANDRA_HOSTS"])
}

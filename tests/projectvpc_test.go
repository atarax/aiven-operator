package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/aiven/aiven-go-client/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/types"

	"github.com/aiven/aiven-operator/api/v1alpha1"
)

func getProjectVPCsYaml(project, vpcName1, cloudName1, vpcName2, cloudName2 string) string {
	return fmt.Sprintf(`
apiVersion: aiven.io/v1alpha1
kind: ProjectVPC
metadata:
  name: %[2]s
spec:
  authSecretRef:
    name: aiven-token
    key: token

  project: %[1]s
  cloudName: %[3]s
  networkCidr: 10.0.0.0/24

---

apiVersion: aiven.io/v1alpha1
kind: ProjectVPC
metadata:
  name: %[4]s
spec:
  authSecretRef:
    name: aiven-token
    key: token

  project: %[1]s
  cloudName: %[5]s
  networkCidr: 10.0.0.0/24
`, project, vpcName1, cloudName1, vpcName2, cloudName2)
}

func getKafkaForProjectVPCYaml(project, vpcID, kafkaName, cloudName string) string {
	return fmt.Sprintf(`
apiVersion: aiven.io/v1alpha1
kind: Kafka
metadata:
  name: %[3]s
spec:
  authSecretRef:
    name: aiven-token
    key: token

  project: %[1]s
  cloudName: %[4]s
  plan: startup-2
  projectVpcId: %[2]s
`, project, vpcID, kafkaName, cloudName)
}

// TestProjectVPCID Kafka.Spec.ProjectVPCID
// creates vpc and then creates kafka with given vpcID
func TestProjectVPCID(t *testing.T) {
	t.Parallel()
	defer recoverPanic(t)

	// GIVEN
	ctx := context.Background()
	vpcName1 := randName("project-vpc-id")
	vpcName2 := randName("project-vpc-id")
	vpcYaml := getProjectVPCsYaml(testProject, vpcName1, testSecondaryCloudName, vpcName2, testTertiaryCloudName)
	vpcSession := NewSession(k8sClient, avnClient, testProject)

	// Cleans test afterwards
	defer vpcSession.Destroy()

	// WHEN
	// Applies given manifest
	require.NoError(t, vpcSession.Apply(vpcYaml))

	// Waits kube object
	vpc1 := new(v1alpha1.ProjectVPC)
	require.NoError(t, vpcSession.GetRunning(vpc1, vpcName1))

	vpc2 := new(v1alpha1.ProjectVPC)
	require.NoError(t, vpcSession.GetRunning(vpc2, vpcName2))

	// THEN
	// Validates VPC
	vpc1Avn, err := avnClient.VPCs.Get(ctx, testProject, vpc1.Status.ID)
	require.NoError(t, err)
	assert.Equal(t, "ACTIVE", vpc1Avn.State)
	assert.Equal(t, vpc1Avn.State, vpc1.Status.State)
	assert.Equal(t, vpc1Avn.CloudName, vpc1.Spec.CloudName)
	assert.Equal(t, "10.0.0.0/24", vpc1.Spec.NetworkCidr)
	assert.Equal(t, vpc1Avn.NetworkCIDR, vpc1.Spec.NetworkCidr)

	vpc2Avn, err := avnClient.VPCs.Get(ctx, testProject, vpc2.Status.ID)
	require.NoError(t, err)
	assert.Equal(t, "ACTIVE", vpc2Avn.State)
	assert.Equal(t, vpc2Avn.State, vpc2.Status.State)
	assert.Equal(t, vpc2Avn.CloudName, vpc2.Spec.CloudName)
	assert.Equal(t, "10.0.0.0/24", vpc2.Spec.NetworkCidr)
	assert.Equal(t, vpc2Avn.NetworkCIDR, vpc2.Spec.NetworkCidr)

	// Creates Kafka with given vpcID
	kafkaName := randName("project-vpc-id")
	kafkaYaml := getKafkaForProjectVPCYaml(testProject, vpc1.Status.ID, kafkaName, testSecondaryCloudName)
	kafkaSession := NewSession(k8sClient, avnClient, testProject)

	// Cleans test afterwards
	defer kafkaSession.Destroy()

	// WHEN
	// Applies given manifest
	require.NoError(t, kafkaSession.Apply(kafkaYaml))

	// Waits kube objects
	kafka := new(v1alpha1.Kafka)
	require.NoError(t, kafkaSession.GetRunning(kafka, kafkaName))

	// THEN
	kafkaAvn, err := avnClient.Services.Get(ctx, testProject, kafkaName)
	require.NoError(t, err)
	assert.Equal(t, kafkaAvn.Name, kafka.GetName())
	assert.Equal(t, "RUNNING", kafka.Status.State)
	assert.Equal(t, kafkaAvn.State, kafka.Status.State)
	assert.Equal(t, kafkaAvn.Plan, kafka.Spec.Plan)
	assert.Equal(t, kafkaAvn.CloudName, kafka.Spec.CloudName)

	// Validates VPC
	assert.Equal(t, vpc1.Status.ID, kafka.Spec.ProjectVPCID)
	assert.Equal(t, anyPointer(vpc1.Status.ID), kafkaAvn.ProjectVPCID)

	// Migrates the service to vpc2
	kafkaYamlUpd := getKafkaForProjectVPCYaml(testProject, vpc2.Status.ID, kafkaName, testTertiaryCloudName)
	require.NoError(t, kafkaSession.Apply(kafkaYamlUpd))

	// This migration takes too long, so we don't wait it's being in the RUNNING state in kube
	ctx, cancel := context.WithTimeout(ctx, waitRunningTimeout)
	defer cancel()

	// Gets Aiven object
	var kafkaAvnUpd *aiven.Service
	require.NoError(t, retryForever(ctx, func() (bool, error) {
		kafkaAvnUpd, err = avnClient.Services.Get(ctx, testProject, kafkaName)
		if err != nil {
			return false, err
		}

		// Just waits object being updated in Aiven
		return kafkaAvnUpd.CloudName != testTertiaryCloudName, nil
	}))

	// Gets kube object
	kafkaUpd := new(v1alpha1.Kafka)
	require.NoError(t, k8sClient.Get(ctx, types.NamespacedName{Name: kafkaName, Namespace: "default"}, kafkaUpd))
	assert.Equal(t, vpc2.Status.ID, kafkaUpd.Spec.ProjectVPCID)
	assert.Equal(t, vpc2.Status.ID, *kafkaAvnUpd.ProjectVPCID)
	assert.Equal(t, vpc2.Spec.CloudName, kafkaUpd.Spec.CloudName)
	assert.Equal(t, vpc2.Spec.CloudName, kafkaAvnUpd.CloudName)
}

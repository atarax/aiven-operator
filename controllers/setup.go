package controllers

import (
	"fmt"
	"strings"

	ctrl "sigs.k8s.io/controller-runtime"
)

func SetupControllers(mgr ctrl.Manager, defaultToken string) error {
	if err := (&SecretFinalizerGCController{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("controllers").WithName("SecretFinalizerGCController"),
	}).SetupWithManager(mgr, defaultToken != ""); err != nil {
		return fmt.Errorf("controller SecretFinalizerGCController: %w", err)
	}

	if err := (&ProjectReconciler{
		Controller: newController(mgr, "Project", defaultToken),
	}).SetupWithManager(mgr); err != nil {
		return fmt.Errorf("controller Project: %w", err)
	}

	if err := (&PostgreSQLReconciler{
		Controller: newController(mgr, "PostgreSQL", defaultToken),
	}).SetupWithManager(mgr); err != nil {
		return fmt.Errorf("controller PostgreSQL: %w", err)
	}

	if err := (&ConnectionPoolReconciler{
		Controller: newController(mgr, "ConnectionPool", defaultToken),
	}).SetupWithManager(mgr); err != nil {
		return fmt.Errorf("controller ConnectionPool: %w", err)
	}

	if err := (&DatabaseReconciler{
		Controller: newController(mgr, "Database", defaultToken),
	}).SetupWithManager(mgr); err != nil {
		return fmt.Errorf("controller Database: %w", err)
	}

	if err := (&KafkaReconciler{
		Controller: newController(mgr, "Kafka", defaultToken),
	}).SetupWithManager(mgr); err != nil {
		return fmt.Errorf("controller Kafka: %w", err)
	}

	if err := (&ProjectVPCReconciler{
		Controller: newController(mgr, "ProjectVPC", defaultToken),
	}).SetupWithManager(mgr); err != nil {
		return fmt.Errorf("controller ProjectVPC: %w", err)
	}

	if err := (&KafkaTopicReconciler{
		Controller: newController(mgr, "KafkaTopic", defaultToken),
	}).SetupWithManager(mgr); err != nil {
		return fmt.Errorf("controller KafkaTopic: %w", err)
	}

	if err := (&KafkaACLReconciler{
		Controller: newController(mgr, "KafkaACL", defaultToken),
	}).SetupWithManager(mgr); err != nil {
		return fmt.Errorf("controller KafkaACL: %w", err)
	}

	if err := (&KafkaConnectReconciler{
		Controller: newController(mgr, "KafkaConnect", defaultToken),
	}).SetupWithManager(mgr); err != nil {
		return fmt.Errorf("controller KafkaConnect: %w", err)
	}

	if err := (&ServiceUserReconciler{
		Controller: newController(mgr, "ServiceUser", defaultToken),
	}).SetupWithManager(mgr); err != nil {
		return fmt.Errorf("controller ServiceUser: %w", err)
	}

	if err := (&KafkaSchemaReconciler{
		Controller: newController(mgr, "KafkaSchema", defaultToken),
	}).SetupWithManager(mgr); err != nil {
		return fmt.Errorf("controller KafkaSchema: %w", err)
	}

	if err := (&ServiceIntegrationReconciler{
		Controller: newController(mgr, "ServiceIntegration", defaultToken),
	}).SetupWithManager(mgr); err != nil {
		return fmt.Errorf("controller ServiceIntegration: %w", err)
	}
	if err := (&KafkaConnectorReconciler{
		Controller: newController(mgr, "KafkaConnector", defaultToken),
	}).SetupWithManager(mgr); err != nil {
		return fmt.Errorf("controller KafkaConnector: %w", err)
	}

	if err := (&RedisReconciler{
		Controller: newController(mgr, "Redis", defaultToken),
	}).SetupWithManager(mgr); err != nil {
		return fmt.Errorf("controller Redis: %w", err)
	}

	if err := (&OpenSearchReconciler{
		Controller: newController(mgr, "OpenSearch", defaultToken),
	}).SetupWithManager(mgr); err != nil {
		return fmt.Errorf("controller OpenSearch: %w", err)
	}

	if err := (&ClickhouseReconciler{
		Controller: newController(mgr, "Clickhouse", defaultToken),
	}).SetupWithManager(mgr); err != nil {
		return fmt.Errorf("controller Clickhouse: %w", err)
	}

	if err := (&ClickhouseUserReconciler{
		Controller: newController(mgr, "ClickhouseUser", defaultToken),
	}).SetupWithManager(mgr); err != nil {
		return fmt.Errorf("controller ClickhouseUser: %w", err)
	}

	if err := (&MySQLReconciler{
		Controller: newController(mgr, "MySQL", defaultToken),
	}).SetupWithManager(mgr); err != nil {
		return fmt.Errorf("controller MySQL: %w", err)
	}

	if err := (&CassandraReconciler{
		Controller: newController(mgr, "Cassandra", defaultToken),
	}).SetupWithManager(mgr); err != nil {
		return fmt.Errorf("controller Cassandra: %w", err)
	}

	if err := (&GrafanaReconciler{
		Controller: newController(mgr, "Grafana", defaultToken),
	}).SetupWithManager(mgr); err != nil {
		return fmt.Errorf("controller Grafana: %w", err)
	}

	//+kubebuilder:scaffold:builder
	return nil
}

func newController(mgr ctrl.Manager, name, defaultToken string) Controller {
	return Controller{
		Client:       mgr.GetClient(),
		Log:          ctrl.Log.WithName("controllers").WithName(name),
		Scheme:       mgr.GetScheme(),
		Recorder:     mgr.GetEventRecorderFor(strings.ToLower(name) + "-reconciler"),
		DefaultToken: defaultToken,
	}
}

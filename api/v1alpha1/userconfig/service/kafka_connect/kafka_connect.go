// Code generated by user config generator. DO NOT EDIT.
// +kubebuilder:object:generate=true

package kafkaconnectuserconfig

// CIDR address block, either as a string, or in a dict with an optional description field
type IpFilter struct {
	// +kubebuilder:validation:MaxLength=1024
	// Description for IP filter list entry
	Description *string `groups:"create,update" json:"description,omitempty"`

	// +kubebuilder:validation:MaxLength=43
	// CIDR address block
	Network string `groups:"create,update" json:"network"`
}

// Kafka Connect configuration values
type KafkaConnect struct {
	// +kubebuilder:validation:Enum="None";"All"
	// Defines what client configurations can be overridden by the connector. Default is None
	ConnectorClientConfigOverridePolicy *string `groups:"create,update" json:"connector_client_config_override_policy,omitempty"`

	// +kubebuilder:validation:Enum="earliest";"latest"
	// What to do when there is no initial offset in Kafka or if the current offset does not exist any more on the server. Default is earliest
	ConsumerAutoOffsetReset *string `groups:"create,update" json:"consumer_auto_offset_reset,omitempty"`

	// +kubebuilder:validation:Minimum=1048576
	// +kubebuilder:validation:Maximum=104857600
	// Records are fetched in batches by the consumer, and if the first record batch in the first non-empty partition of the fetch is larger than this value, the record batch will still be returned to ensure that the consumer can make progress. As such, this is not a absolute maximum.
	ConsumerFetchMaxBytes *int `groups:"create,update" json:"consumer_fetch_max_bytes,omitempty"`

	// +kubebuilder:validation:Enum="read_uncommitted";"read_committed"
	// Transaction read isolation level. read_uncommitted is the default, but read_committed can be used if consume-exactly-once behavior is desired.
	ConsumerIsolationLevel *string `groups:"create,update" json:"consumer_isolation_level,omitempty"`

	// +kubebuilder:validation:Minimum=1048576
	// +kubebuilder:validation:Maximum=104857600
	// Records are fetched in batches by the consumer.If the first record batch in the first non-empty partition of the fetch is larger than this limit, the batch will still be returned to ensure that the consumer can make progress.
	ConsumerMaxPartitionFetchBytes *int `groups:"create,update" json:"consumer_max_partition_fetch_bytes,omitempty"`

	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=2147483647
	// The maximum delay in milliseconds between invocations of poll() when using consumer group management (defaults to 300000).
	ConsumerMaxPollIntervalMs *int `groups:"create,update" json:"consumer_max_poll_interval_ms,omitempty"`

	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=10000
	// The maximum number of records returned in a single call to poll() (defaults to 500).
	ConsumerMaxPollRecords *int `groups:"create,update" json:"consumer_max_poll_records,omitempty"`

	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=100000000
	// The interval at which to try committing offsets for tasks (defaults to 60000).
	OffsetFlushIntervalMs *int `groups:"create,update" json:"offset_flush_interval_ms,omitempty"`

	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=2147483647
	// Maximum number of milliseconds to wait for records to flush and partition offset data to be committed to offset storage before cancelling the process and restoring the offset data to be committed in a future attempt (defaults to 5000).
	OffsetFlushTimeoutMs *int `groups:"create,update" json:"offset_flush_timeout_ms,omitempty"`

	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=5242880
	// This setting gives the upper bound of the batch size to be sent. If there are fewer than this many bytes accumulated for this partition, the producer will 'linger' for the linger.ms time waiting for more records to show up. A batch size of zero will disable batching entirely (defaults to 16384).
	ProducerBatchSize *int `groups:"create,update" json:"producer_batch_size,omitempty"`

	// +kubebuilder:validation:Minimum=5242880
	// +kubebuilder:validation:Maximum=134217728
	// The total bytes of memory the producer can use to buffer records waiting to be sent to the broker (defaults to 33554432).
	ProducerBufferMemory *int `groups:"create,update" json:"producer_buffer_memory,omitempty"`

	// +kubebuilder:validation:Enum="gzip";"snappy";"lz4";"zstd";"none"
	// Specify the default compression type for producers. This configuration accepts the standard compression codecs ('gzip', 'snappy', 'lz4', 'zstd'). It additionally accepts 'none' which is the default and equivalent to no compression.
	ProducerCompressionType *string `groups:"create,update" json:"producer_compression_type,omitempty"`

	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=5000
	// This setting gives the upper bound on the delay for batching: once there is batch.size worth of records for a partition it will be sent immediately regardless of this setting, however if there are fewer than this many bytes accumulated for this partition the producer will 'linger' for the specified time waiting for more records to show up. Defaults to 0.
	ProducerLingerMs *int `groups:"create,update" json:"producer_linger_ms,omitempty"`

	// +kubebuilder:validation:Minimum=131072
	// +kubebuilder:validation:Maximum=67108864
	// This setting will limit the number of record batches the producer will send in a single request to avoid sending huge requests.
	ProducerMaxRequestSize *int `groups:"create,update" json:"producer_max_request_size,omitempty"`

	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=600000
	// The maximum delay that is scheduled in order to wait for the return of one or more departed workers before rebalancing and reassigning their connectors and tasks to the group. During this period the connectors and tasks of the departed workers remain unassigned.  Defaults to 5 minutes.
	ScheduledRebalanceMaxDelayMs *int `groups:"create,update" json:"scheduled_rebalance_max_delay_ms,omitempty"`

	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=2147483647
	// The timeout in milliseconds used to detect failures when using Kafka’s group management facilities (defaults to 10000).
	SessionTimeoutMs *int `groups:"create,update" json:"session_timeout_ms,omitempty"`
}

// Allow access to selected service ports from private networks
type PrivateAccess struct {
	// Allow clients to connect to kafka_connect with a DNS name that always resolves to the service's private IP addresses. Only available in certain network locations
	KafkaConnect *bool `groups:"create,update" json:"kafka_connect,omitempty"`

	// Allow clients to connect to prometheus with a DNS name that always resolves to the service's private IP addresses. Only available in certain network locations
	Prometheus *bool `groups:"create,update" json:"prometheus,omitempty"`
}

// Allow access to selected service components through Privatelink
type PrivatelinkAccess struct {
	// Enable jolokia
	Jolokia *bool `groups:"create,update" json:"jolokia,omitempty"`

	// Enable kafka_connect
	KafkaConnect *bool `groups:"create,update" json:"kafka_connect,omitempty"`

	// Enable prometheus
	Prometheus *bool `groups:"create,update" json:"prometheus,omitempty"`
}

// Allow access to selected service ports from the public Internet
type PublicAccess struct {
	// Allow clients to connect to kafka_connect from the public internet for service nodes that are in a project VPC or another type of private network
	KafkaConnect *bool `groups:"create,update" json:"kafka_connect,omitempty"`

	// Allow clients to connect to prometheus from the public internet for service nodes that are in a project VPC or another type of private network
	Prometheus *bool `groups:"create,update" json:"prometheus,omitempty"`
}
type KafkaConnectUserConfig struct {
	// +kubebuilder:validation:MaxItems=1
	// Additional Cloud Regions for Backup Replication
	AdditionalBackupRegions []string `groups:"create,update" json:"additional_backup_regions,omitempty"`

	// +kubebuilder:validation:MaxItems=1024
	// Allow incoming connections from CIDR address block, e.g. '10.20.0.0/16'
	IpFilter []*IpFilter `groups:"create,update" json:"ip_filter,omitempty"`

	// Kafka Connect configuration values
	KafkaConnect *KafkaConnect `groups:"create,update" json:"kafka_connect,omitempty"`

	// Allow access to selected service ports from private networks
	PrivateAccess *PrivateAccess `groups:"create,update" json:"private_access,omitempty"`

	// Allow access to selected service components through Privatelink
	PrivatelinkAccess *PrivatelinkAccess `groups:"create,update" json:"privatelink_access,omitempty"`

	// Allow access to selected service ports from the public Internet
	PublicAccess *PublicAccess `groups:"create,update" json:"public_access,omitempty"`

	// Use static public IP addresses
	StaticIps *bool `groups:"create,update" json:"static_ips,omitempty"`
}

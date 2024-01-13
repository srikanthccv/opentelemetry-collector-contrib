// Code generated by mdatagen. DO NOT EDIT.

package metadata

import "go.opentelemetry.io/collector/confmap"

// MetricConfig provides common config for a particular metric.
type MetricConfig struct {
	Enabled bool `mapstructure:"enabled"`

	enabledSetByUser bool
}

func (ms *MetricConfig) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(ms)
	if err != nil {
		return err
	}
	ms.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// MetricsConfig provides config for saphana metrics.
type MetricsConfig struct {
	SaphanaAlertCount                       MetricConfig `mapstructure:"saphana.alert.count"`
	SaphanaBackupLatest                     MetricConfig `mapstructure:"saphana.backup.latest"`
	SaphanaColumnMemoryUsed                 MetricConfig `mapstructure:"saphana.column.memory.used"`
	SaphanaComponentMemoryUsed              MetricConfig `mapstructure:"saphana.component.memory.used"`
	SaphanaConnectionCount                  MetricConfig `mapstructure:"saphana.connection.count"`
	SaphanaCPUUsed                          MetricConfig `mapstructure:"saphana.cpu.used"`
	SaphanaDiskSizeCurrent                  MetricConfig `mapstructure:"saphana.disk.size.current"`
	SaphanaHostMemoryCurrent                MetricConfig `mapstructure:"saphana.host.memory.current"`
	SaphanaHostSwapCurrent                  MetricConfig `mapstructure:"saphana.host.swap.current"`
	SaphanaInstanceCodeSize                 MetricConfig `mapstructure:"saphana.instance.code_size"`
	SaphanaInstanceMemoryCurrent            MetricConfig `mapstructure:"saphana.instance.memory.current"`
	SaphanaInstanceMemorySharedAllocated    MetricConfig `mapstructure:"saphana.instance.memory.shared.allocated"`
	SaphanaInstanceMemoryUsedPeak           MetricConfig `mapstructure:"saphana.instance.memory.used.peak"`
	SaphanaLicenseExpirationTime            MetricConfig `mapstructure:"saphana.license.expiration.time"`
	SaphanaLicenseLimit                     MetricConfig `mapstructure:"saphana.license.limit"`
	SaphanaLicensePeak                      MetricConfig `mapstructure:"saphana.license.peak"`
	SaphanaNetworkRequestAverageTime        MetricConfig `mapstructure:"saphana.network.request.average_time"`
	SaphanaNetworkRequestCount              MetricConfig `mapstructure:"saphana.network.request.count"`
	SaphanaNetworkRequestFinishedCount      MetricConfig `mapstructure:"saphana.network.request.finished.count"`
	SaphanaReplicationAverageTime           MetricConfig `mapstructure:"saphana.replication.average_time"`
	SaphanaReplicationBacklogSize           MetricConfig `mapstructure:"saphana.replication.backlog.size"`
	SaphanaReplicationBacklogTime           MetricConfig `mapstructure:"saphana.replication.backlog.time"`
	SaphanaRowStoreMemoryUsed               MetricConfig `mapstructure:"saphana.row_store.memory.used"`
	SaphanaSchemaMemoryUsedCurrent          MetricConfig `mapstructure:"saphana.schema.memory.used.current"`
	SaphanaSchemaMemoryUsedMax              MetricConfig `mapstructure:"saphana.schema.memory.used.max"`
	SaphanaSchemaOperationCount             MetricConfig `mapstructure:"saphana.schema.operation.count"`
	SaphanaSchemaRecordCompressedCount      MetricConfig `mapstructure:"saphana.schema.record.compressed.count"`
	SaphanaSchemaRecordCount                MetricConfig `mapstructure:"saphana.schema.record.count"`
	SaphanaServiceCodeSize                  MetricConfig `mapstructure:"saphana.service.code_size"`
	SaphanaServiceCount                     MetricConfig `mapstructure:"saphana.service.count"`
	SaphanaServiceMemoryCompactorsAllocated MetricConfig `mapstructure:"saphana.service.memory.compactors.allocated"`
	SaphanaServiceMemoryCompactorsFreeable  MetricConfig `mapstructure:"saphana.service.memory.compactors.freeable"`
	SaphanaServiceMemoryEffectiveLimit      MetricConfig `mapstructure:"saphana.service.memory.effective_limit"`
	SaphanaServiceMemoryHeapCurrent         MetricConfig `mapstructure:"saphana.service.memory.heap.current"`
	SaphanaServiceMemoryLimit               MetricConfig `mapstructure:"saphana.service.memory.limit"`
	SaphanaServiceMemorySharedCurrent       MetricConfig `mapstructure:"saphana.service.memory.shared.current"`
	SaphanaServiceMemoryUsed                MetricConfig `mapstructure:"saphana.service.memory.used"`
	SaphanaServiceStackSize                 MetricConfig `mapstructure:"saphana.service.stack_size"`
	SaphanaServiceThreadCount               MetricConfig `mapstructure:"saphana.service.thread.count"`
	SaphanaTransactionBlocked               MetricConfig `mapstructure:"saphana.transaction.blocked"`
	SaphanaTransactionCount                 MetricConfig `mapstructure:"saphana.transaction.count"`
	SaphanaUptime                           MetricConfig `mapstructure:"saphana.uptime"`
	SaphanaVolumeOperationCount             MetricConfig `mapstructure:"saphana.volume.operation.count"`
	SaphanaVolumeOperationSize              MetricConfig `mapstructure:"saphana.volume.operation.size"`
	SaphanaVolumeOperationTime              MetricConfig `mapstructure:"saphana.volume.operation.time"`
}

func DefaultMetricsConfig() MetricsConfig {
	return MetricsConfig{
		SaphanaAlertCount: MetricConfig{
			Enabled: true,
		},
		SaphanaBackupLatest: MetricConfig{
			Enabled: true,
		},
		SaphanaColumnMemoryUsed: MetricConfig{
			Enabled: true,
		},
		SaphanaComponentMemoryUsed: MetricConfig{
			Enabled: true,
		},
		SaphanaConnectionCount: MetricConfig{
			Enabled: true,
		},
		SaphanaCPUUsed: MetricConfig{
			Enabled: true,
		},
		SaphanaDiskSizeCurrent: MetricConfig{
			Enabled: true,
		},
		SaphanaHostMemoryCurrent: MetricConfig{
			Enabled: true,
		},
		SaphanaHostSwapCurrent: MetricConfig{
			Enabled: true,
		},
		SaphanaInstanceCodeSize: MetricConfig{
			Enabled: true,
		},
		SaphanaInstanceMemoryCurrent: MetricConfig{
			Enabled: true,
		},
		SaphanaInstanceMemorySharedAllocated: MetricConfig{
			Enabled: true,
		},
		SaphanaInstanceMemoryUsedPeak: MetricConfig{
			Enabled: true,
		},
		SaphanaLicenseExpirationTime: MetricConfig{
			Enabled: true,
		},
		SaphanaLicenseLimit: MetricConfig{
			Enabled: true,
		},
		SaphanaLicensePeak: MetricConfig{
			Enabled: true,
		},
		SaphanaNetworkRequestAverageTime: MetricConfig{
			Enabled: true,
		},
		SaphanaNetworkRequestCount: MetricConfig{
			Enabled: true,
		},
		SaphanaNetworkRequestFinishedCount: MetricConfig{
			Enabled: true,
		},
		SaphanaReplicationAverageTime: MetricConfig{
			Enabled: true,
		},
		SaphanaReplicationBacklogSize: MetricConfig{
			Enabled: true,
		},
		SaphanaReplicationBacklogTime: MetricConfig{
			Enabled: true,
		},
		SaphanaRowStoreMemoryUsed: MetricConfig{
			Enabled: true,
		},
		SaphanaSchemaMemoryUsedCurrent: MetricConfig{
			Enabled: true,
		},
		SaphanaSchemaMemoryUsedMax: MetricConfig{
			Enabled: true,
		},
		SaphanaSchemaOperationCount: MetricConfig{
			Enabled: true,
		},
		SaphanaSchemaRecordCompressedCount: MetricConfig{
			Enabled: true,
		},
		SaphanaSchemaRecordCount: MetricConfig{
			Enabled: true,
		},
		SaphanaServiceCodeSize: MetricConfig{
			Enabled: true,
		},
		SaphanaServiceCount: MetricConfig{
			Enabled: true,
		},
		SaphanaServiceMemoryCompactorsAllocated: MetricConfig{
			Enabled: true,
		},
		SaphanaServiceMemoryCompactorsFreeable: MetricConfig{
			Enabled: true,
		},
		SaphanaServiceMemoryEffectiveLimit: MetricConfig{
			Enabled: true,
		},
		SaphanaServiceMemoryHeapCurrent: MetricConfig{
			Enabled: true,
		},
		SaphanaServiceMemoryLimit: MetricConfig{
			Enabled: true,
		},
		SaphanaServiceMemorySharedCurrent: MetricConfig{
			Enabled: true,
		},
		SaphanaServiceMemoryUsed: MetricConfig{
			Enabled: true,
		},
		SaphanaServiceStackSize: MetricConfig{
			Enabled: true,
		},
		SaphanaServiceThreadCount: MetricConfig{
			Enabled: true,
		},
		SaphanaTransactionBlocked: MetricConfig{
			Enabled: true,
		},
		SaphanaTransactionCount: MetricConfig{
			Enabled: true,
		},
		SaphanaUptime: MetricConfig{
			Enabled: true,
		},
		SaphanaVolumeOperationCount: MetricConfig{
			Enabled: true,
		},
		SaphanaVolumeOperationSize: MetricConfig{
			Enabled: true,
		},
		SaphanaVolumeOperationTime: MetricConfig{
			Enabled: true,
		},
	}
}

// ResourceAttributeConfig provides common config for a particular resource attribute.
type ResourceAttributeConfig struct {
	Enabled bool `mapstructure:"enabled"`

	enabledSetByUser bool
}

func (rac *ResourceAttributeConfig) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(rac)
	if err != nil {
		return err
	}
	rac.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// ResourceAttributesConfig provides config for saphana resource attributes.
type ResourceAttributesConfig struct {
	DbSystem    ResourceAttributeConfig `mapstructure:"db.system"`
	SaphanaHost ResourceAttributeConfig `mapstructure:"saphana.host"`
}

func DefaultResourceAttributesConfig() ResourceAttributesConfig {
	return ResourceAttributesConfig{
		DbSystem: ResourceAttributeConfig{
			Enabled: true,
		},
		SaphanaHost: ResourceAttributeConfig{
			Enabled: true,
		},
	}
}

// MetricsBuilderConfig is a configuration for saphana metrics builder.
type MetricsBuilderConfig struct {
	Metrics            MetricsConfig            `mapstructure:"metrics"`
	ResourceAttributes ResourceAttributesConfig `mapstructure:"resource_attributes"`
}

func DefaultMetricsBuilderConfig() MetricsBuilderConfig {
	return MetricsBuilderConfig{
		Metrics:            DefaultMetricsConfig(),
		ResourceAttributes: DefaultResourceAttributesConfig(),
	}
}

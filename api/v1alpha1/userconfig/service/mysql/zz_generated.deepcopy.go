//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright (c) 2022 Aiven, Helsinki, Finland. https://aiven.io/

// Code generated by controller-gen. DO NOT EDIT.

package mysqluserconfig

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IpFilter) DeepCopyInto(out *IpFilter) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IpFilter.
func (in *IpFilter) DeepCopy() *IpFilter {
	if in == nil {
		return nil
	}
	out := new(IpFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Migration) DeepCopyInto(out *Migration) {
	*out = *in
	if in.Dbname != nil {
		in, out := &in.Dbname, &out.Dbname
		*out = new(string)
		**out = **in
	}
	if in.IgnoreDbs != nil {
		in, out := &in.IgnoreDbs, &out.IgnoreDbs
		*out = new(string)
		**out = **in
	}
	if in.Method != nil {
		in, out := &in.Method, &out.Method
		*out = new(string)
		**out = **in
	}
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(string)
		**out = **in
	}
	if in.Ssl != nil {
		in, out := &in.Ssl, &out.Ssl
		*out = new(bool)
		**out = **in
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Migration.
func (in *Migration) DeepCopy() *Migration {
	if in == nil {
		return nil
	}
	out := new(Migration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Mysql) DeepCopyInto(out *Mysql) {
	*out = *in
	if in.ConnectTimeout != nil {
		in, out := &in.ConnectTimeout, &out.ConnectTimeout
		*out = new(int)
		**out = **in
	}
	if in.DefaultTimeZone != nil {
		in, out := &in.DefaultTimeZone, &out.DefaultTimeZone
		*out = new(string)
		**out = **in
	}
	if in.GroupConcatMaxLen != nil {
		in, out := &in.GroupConcatMaxLen, &out.GroupConcatMaxLen
		*out = new(int)
		**out = **in
	}
	if in.InformationSchemaStatsExpiry != nil {
		in, out := &in.InformationSchemaStatsExpiry, &out.InformationSchemaStatsExpiry
		*out = new(int)
		**out = **in
	}
	if in.InnodbChangeBufferMaxSize != nil {
		in, out := &in.InnodbChangeBufferMaxSize, &out.InnodbChangeBufferMaxSize
		*out = new(int)
		**out = **in
	}
	if in.InnodbFlushNeighbors != nil {
		in, out := &in.InnodbFlushNeighbors, &out.InnodbFlushNeighbors
		*out = new(int)
		**out = **in
	}
	if in.InnodbFtMinTokenSize != nil {
		in, out := &in.InnodbFtMinTokenSize, &out.InnodbFtMinTokenSize
		*out = new(int)
		**out = **in
	}
	if in.InnodbFtServerStopwordTable != nil {
		in, out := &in.InnodbFtServerStopwordTable, &out.InnodbFtServerStopwordTable
		*out = new(string)
		**out = **in
	}
	if in.InnodbLockWaitTimeout != nil {
		in, out := &in.InnodbLockWaitTimeout, &out.InnodbLockWaitTimeout
		*out = new(int)
		**out = **in
	}
	if in.InnodbLogBufferSize != nil {
		in, out := &in.InnodbLogBufferSize, &out.InnodbLogBufferSize
		*out = new(int)
		**out = **in
	}
	if in.InnodbOnlineAlterLogMaxSize != nil {
		in, out := &in.InnodbOnlineAlterLogMaxSize, &out.InnodbOnlineAlterLogMaxSize
		*out = new(int)
		**out = **in
	}
	if in.InnodbPrintAllDeadlocks != nil {
		in, out := &in.InnodbPrintAllDeadlocks, &out.InnodbPrintAllDeadlocks
		*out = new(bool)
		**out = **in
	}
	if in.InnodbReadIoThreads != nil {
		in, out := &in.InnodbReadIoThreads, &out.InnodbReadIoThreads
		*out = new(int)
		**out = **in
	}
	if in.InnodbRollbackOnTimeout != nil {
		in, out := &in.InnodbRollbackOnTimeout, &out.InnodbRollbackOnTimeout
		*out = new(bool)
		**out = **in
	}
	if in.InnodbThreadConcurrency != nil {
		in, out := &in.InnodbThreadConcurrency, &out.InnodbThreadConcurrency
		*out = new(int)
		**out = **in
	}
	if in.InnodbWriteIoThreads != nil {
		in, out := &in.InnodbWriteIoThreads, &out.InnodbWriteIoThreads
		*out = new(int)
		**out = **in
	}
	if in.InteractiveTimeout != nil {
		in, out := &in.InteractiveTimeout, &out.InteractiveTimeout
		*out = new(int)
		**out = **in
	}
	if in.InternalTmpMemStorageEngine != nil {
		in, out := &in.InternalTmpMemStorageEngine, &out.InternalTmpMemStorageEngine
		*out = new(string)
		**out = **in
	}
	if in.LongQueryTime != nil {
		in, out := &in.LongQueryTime, &out.LongQueryTime
		*out = new(float64)
		**out = **in
	}
	if in.MaxAllowedPacket != nil {
		in, out := &in.MaxAllowedPacket, &out.MaxAllowedPacket
		*out = new(int)
		**out = **in
	}
	if in.MaxHeapTableSize != nil {
		in, out := &in.MaxHeapTableSize, &out.MaxHeapTableSize
		*out = new(int)
		**out = **in
	}
	if in.NetBufferLength != nil {
		in, out := &in.NetBufferLength, &out.NetBufferLength
		*out = new(int)
		**out = **in
	}
	if in.NetReadTimeout != nil {
		in, out := &in.NetReadTimeout, &out.NetReadTimeout
		*out = new(int)
		**out = **in
	}
	if in.NetWriteTimeout != nil {
		in, out := &in.NetWriteTimeout, &out.NetWriteTimeout
		*out = new(int)
		**out = **in
	}
	if in.SlowQueryLog != nil {
		in, out := &in.SlowQueryLog, &out.SlowQueryLog
		*out = new(bool)
		**out = **in
	}
	if in.SortBufferSize != nil {
		in, out := &in.SortBufferSize, &out.SortBufferSize
		*out = new(int)
		**out = **in
	}
	if in.SqlMode != nil {
		in, out := &in.SqlMode, &out.SqlMode
		*out = new(string)
		**out = **in
	}
	if in.SqlRequirePrimaryKey != nil {
		in, out := &in.SqlRequirePrimaryKey, &out.SqlRequirePrimaryKey
		*out = new(bool)
		**out = **in
	}
	if in.TmpTableSize != nil {
		in, out := &in.TmpTableSize, &out.TmpTableSize
		*out = new(int)
		**out = **in
	}
	if in.WaitTimeout != nil {
		in, out := &in.WaitTimeout, &out.WaitTimeout
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Mysql.
func (in *Mysql) DeepCopy() *Mysql {
	if in == nil {
		return nil
	}
	out := new(Mysql)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MysqlUserConfig) DeepCopyInto(out *MysqlUserConfig) {
	*out = *in
	if in.AdditionalBackupRegions != nil {
		in, out := &in.AdditionalBackupRegions, &out.AdditionalBackupRegions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AdminPassword != nil {
		in, out := &in.AdminPassword, &out.AdminPassword
		*out = new(string)
		**out = **in
	}
	if in.AdminUsername != nil {
		in, out := &in.AdminUsername, &out.AdminUsername
		*out = new(string)
		**out = **in
	}
	if in.BackupHour != nil {
		in, out := &in.BackupHour, &out.BackupHour
		*out = new(int)
		**out = **in
	}
	if in.BackupMinute != nil {
		in, out := &in.BackupMinute, &out.BackupMinute
		*out = new(int)
		**out = **in
	}
	if in.BinlogRetentionPeriod != nil {
		in, out := &in.BinlogRetentionPeriod, &out.BinlogRetentionPeriod
		*out = new(int)
		**out = **in
	}
	if in.IpFilter != nil {
		in, out := &in.IpFilter, &out.IpFilter
		*out = make([]*IpFilter, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(IpFilter)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Migration != nil {
		in, out := &in.Migration, &out.Migration
		*out = new(Migration)
		(*in).DeepCopyInto(*out)
	}
	if in.Mysql != nil {
		in, out := &in.Mysql, &out.Mysql
		*out = new(Mysql)
		(*in).DeepCopyInto(*out)
	}
	if in.MysqlVersion != nil {
		in, out := &in.MysqlVersion, &out.MysqlVersion
		*out = new(string)
		**out = **in
	}
	if in.PrivateAccess != nil {
		in, out := &in.PrivateAccess, &out.PrivateAccess
		*out = new(PrivateAccess)
		(*in).DeepCopyInto(*out)
	}
	if in.PrivatelinkAccess != nil {
		in, out := &in.PrivatelinkAccess, &out.PrivatelinkAccess
		*out = new(PrivatelinkAccess)
		(*in).DeepCopyInto(*out)
	}
	if in.ProjectToForkFrom != nil {
		in, out := &in.ProjectToForkFrom, &out.ProjectToForkFrom
		*out = new(string)
		**out = **in
	}
	if in.PublicAccess != nil {
		in, out := &in.PublicAccess, &out.PublicAccess
		*out = new(PublicAccess)
		(*in).DeepCopyInto(*out)
	}
	if in.RecoveryTargetTime != nil {
		in, out := &in.RecoveryTargetTime, &out.RecoveryTargetTime
		*out = new(string)
		**out = **in
	}
	if in.ServiceToForkFrom != nil {
		in, out := &in.ServiceToForkFrom, &out.ServiceToForkFrom
		*out = new(string)
		**out = **in
	}
	if in.StaticIps != nil {
		in, out := &in.StaticIps, &out.StaticIps
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MysqlUserConfig.
func (in *MysqlUserConfig) DeepCopy() *MysqlUserConfig {
	if in == nil {
		return nil
	}
	out := new(MysqlUserConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateAccess) DeepCopyInto(out *PrivateAccess) {
	*out = *in
	if in.Mysql != nil {
		in, out := &in.Mysql, &out.Mysql
		*out = new(bool)
		**out = **in
	}
	if in.Mysqlx != nil {
		in, out := &in.Mysqlx, &out.Mysqlx
		*out = new(bool)
		**out = **in
	}
	if in.Prometheus != nil {
		in, out := &in.Prometheus, &out.Prometheus
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateAccess.
func (in *PrivateAccess) DeepCopy() *PrivateAccess {
	if in == nil {
		return nil
	}
	out := new(PrivateAccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivatelinkAccess) DeepCopyInto(out *PrivatelinkAccess) {
	*out = *in
	if in.Mysql != nil {
		in, out := &in.Mysql, &out.Mysql
		*out = new(bool)
		**out = **in
	}
	if in.Mysqlx != nil {
		in, out := &in.Mysqlx, &out.Mysqlx
		*out = new(bool)
		**out = **in
	}
	if in.Prometheus != nil {
		in, out := &in.Prometheus, &out.Prometheus
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivatelinkAccess.
func (in *PrivatelinkAccess) DeepCopy() *PrivatelinkAccess {
	if in == nil {
		return nil
	}
	out := new(PrivatelinkAccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicAccess) DeepCopyInto(out *PublicAccess) {
	*out = *in
	if in.Mysql != nil {
		in, out := &in.Mysql, &out.Mysql
		*out = new(bool)
		**out = **in
	}
	if in.Mysqlx != nil {
		in, out := &in.Mysqlx, &out.Mysqlx
		*out = new(bool)
		**out = **in
	}
	if in.Prometheus != nil {
		in, out := &in.Prometheus, &out.Prometheus
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicAccess.
func (in *PublicAccess) DeepCopy() *PublicAccess {
	if in == nil {
		return nil
	}
	out := new(PublicAccess)
	in.DeepCopyInto(out)
	return out
}

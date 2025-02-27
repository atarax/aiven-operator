//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright (c) 2022 Aiven, Helsinki, Finland. https://aiven.io/

// Code generated by controller-gen. DO NOT EDIT.

package clickhousepostgresqluserconfig

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClickhousePostgresqlUserConfig) DeepCopyInto(out *ClickhousePostgresqlUserConfig) {
	*out = *in
	if in.Databases != nil {
		in, out := &in.Databases, &out.Databases
		*out = make([]*Databases, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Databases)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClickhousePostgresqlUserConfig.
func (in *ClickhousePostgresqlUserConfig) DeepCopy() *ClickhousePostgresqlUserConfig {
	if in == nil {
		return nil
	}
	out := new(ClickhousePostgresqlUserConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Databases) DeepCopyInto(out *Databases) {
	*out = *in
	if in.Database != nil {
		in, out := &in.Database, &out.Database
		*out = new(string)
		**out = **in
	}
	if in.Schema != nil {
		in, out := &in.Schema, &out.Schema
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Databases.
func (in *Databases) DeepCopy() *Databases {
	if in == nil {
		return nil
	}
	out := new(Databases)
	in.DeepCopyInto(out)
	return out
}

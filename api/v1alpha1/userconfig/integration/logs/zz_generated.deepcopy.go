//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright (c) 2022 Aiven, Helsinki, Finland. https://aiven.io/

// Code generated by controller-gen. DO NOT EDIT.

package logsuserconfig

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogsUserConfig) DeepCopyInto(out *LogsUserConfig) {
	*out = *in
	if in.ElasticsearchIndexDaysMax != nil {
		in, out := &in.ElasticsearchIndexDaysMax, &out.ElasticsearchIndexDaysMax
		*out = new(int)
		**out = **in
	}
	if in.ElasticsearchIndexPrefix != nil {
		in, out := &in.ElasticsearchIndexPrefix, &out.ElasticsearchIndexPrefix
		*out = new(string)
		**out = **in
	}
	if in.SelectedLogFields != nil {
		in, out := &in.SelectedLogFields, &out.SelectedLogFields
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogsUserConfig.
func (in *LogsUserConfig) DeepCopy() *LogsUserConfig {
	if in == nil {
		return nil
	}
	out := new(LogsUserConfig)
	in.DeepCopyInto(out)
	return out
}

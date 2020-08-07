// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package useragent

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson5c493fc2DecodeGithubComKsensehqEventnativeUseragent(in *jlexer.Lexer, out *Resolver) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson5c493fc2EncodeGithubComKsensehqEventnativeUseragent(out *jwriter.Writer, in Resolver) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Resolver) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5c493fc2EncodeGithubComKsensehqEventnativeUseragent(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Resolver) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5c493fc2EncodeGithubComKsensehqEventnativeUseragent(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Resolver) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5c493fc2DecodeGithubComKsensehqEventnativeUseragent(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Resolver) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5c493fc2DecodeGithubComKsensehqEventnativeUseragent(l, v)
}
func easyjson5c493fc2DecodeGithubComKsensehqEventnativeUseragent1(in *jlexer.Lexer, out *ResolvedUa) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "ua_family":
			out.UaFamily = string(in.String())
		case "ua_version":
			out.UaVersion = string(in.String())
		case "os_family":
			out.OsFamily = string(in.String())
		case "os_version":
			out.OsVersion = string(in.String())
		case "device_family":
			out.DeviceFamily = string(in.String())
		case "device_brand":
			out.DeviceBrand = string(in.String())
		case "device_model":
			out.DeviceModel = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson5c493fc2EncodeGithubComKsensehqEventnativeUseragent1(out *jwriter.Writer, in ResolvedUa) {
	out.RawByte('{')
	first := true
	_ = first
	if in.UaFamily != "" {
		const prefix string = ",\"ua_family\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.UaFamily))
	}
	if in.UaVersion != "" {
		const prefix string = ",\"ua_version\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.UaVersion))
	}
	if in.OsFamily != "" {
		const prefix string = ",\"os_family\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.OsFamily))
	}
	if in.OsVersion != "" {
		const prefix string = ",\"os_version\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.OsVersion))
	}
	if in.DeviceFamily != "" {
		const prefix string = ",\"device_family\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.DeviceFamily))
	}
	if in.DeviceBrand != "" {
		const prefix string = ",\"device_brand\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.DeviceBrand))
	}
	if in.DeviceModel != "" {
		const prefix string = ",\"device_model\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.DeviceModel))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ResolvedUa) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5c493fc2EncodeGithubComKsensehqEventnativeUseragent1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ResolvedUa) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5c493fc2EncodeGithubComKsensehqEventnativeUseragent1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ResolvedUa) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5c493fc2DecodeGithubComKsensehqEventnativeUseragent1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ResolvedUa) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5c493fc2DecodeGithubComKsensehqEventnativeUseragent1(l, v)
}

// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package crypto

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

func easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoPkgCapabilitiesCrypto(in *jlexer.Lexer, out *CertificateVerificationResponse) {
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
		case "trusted":
			out.Trusted = bool(in.Bool())
		case "reason":
			out.Reason = string(in.String())
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
func easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoPkgCapabilitiesCrypto(out *jwriter.Writer, in CertificateVerificationResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"trusted\":"
		out.RawString(prefix[1:])
		out.Bool(bool(in.Trusted))
	}
	{
		const prefix string = ",\"reason\":"
		out.RawString(prefix)
		out.String(string(in.Reason))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CertificateVerificationResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoPkgCapabilitiesCrypto(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CertificateVerificationResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoPkgCapabilitiesCrypto(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CertificateVerificationResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoPkgCapabilitiesCrypto(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CertificateVerificationResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoPkgCapabilitiesCrypto(l, v)
}
func easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoPkgCapabilitiesCrypto1(in *jlexer.Lexer, out *CertificateVerificationRequest) {
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
		case "cert":
			(out.Cert).UnmarshalEasyJSON(in)
		case "cert_chain":
			if in.IsNull() {
				in.Skip()
				out.CertChain = nil
			} else {
				in.Delim('[')
				if out.CertChain == nil {
					if !in.IsDelim(']') {
						out.CertChain = make([]Certificate, 0, 2)
					} else {
						out.CertChain = []Certificate{}
					}
				} else {
					out.CertChain = (out.CertChain)[:0]
				}
				for !in.IsDelim(']') {
					var v1 Certificate
					(v1).UnmarshalEasyJSON(in)
					out.CertChain = append(out.CertChain, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "not_after":
			out.NotAfter = string(in.String())
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
func easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoPkgCapabilitiesCrypto1(out *jwriter.Writer, in CertificateVerificationRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"cert\":"
		out.RawString(prefix[1:])
		(in.Cert).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"cert_chain\":"
		out.RawString(prefix)
		if in.CertChain == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.CertChain {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"not_after\":"
		out.RawString(prefix)
		out.String(string(in.NotAfter))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CertificateVerificationRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoPkgCapabilitiesCrypto1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CertificateVerificationRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoPkgCapabilitiesCrypto1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CertificateVerificationRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoPkgCapabilitiesCrypto1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CertificateVerificationRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoPkgCapabilitiesCrypto1(l, v)
}
func easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoPkgCapabilitiesCrypto2(in *jlexer.Lexer, out *Certificate) {
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
		case "encoding":
			out.Encoding = CertificateEncoding(in.Int())
		case "data":
			if in.IsNull() {
				in.Skip()
				out.Data = nil
			} else {
				in.Delim('[')
				if out.Data == nil {
					if !in.IsDelim(']') {
						out.Data = make([]int32, 0, 16)
					} else {
						out.Data = []int32{}
					}
				} else {
					out.Data = (out.Data)[:0]
				}
				for !in.IsDelim(']') {
					var v4 int32
					v4 = int32(in.Int32())
					out.Data = append(out.Data, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoPkgCapabilitiesCrypto2(out *jwriter.Writer, in Certificate) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"encoding\":"
		out.RawString(prefix[1:])
		(in.Encoding).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"data\":"
		out.RawString(prefix)
		if in.Data == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Data {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.Int32(int32(v6))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Certificate) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoPkgCapabilitiesCrypto2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Certificate) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoPkgCapabilitiesCrypto2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Certificate) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoPkgCapabilitiesCrypto2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Certificate) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoPkgCapabilitiesCrypto2(l, v)
}
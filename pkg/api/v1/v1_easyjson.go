// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1

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

func easyjson102f8a2fDecodeAnagramsvcPkgApiV1(in *jlexer.Lexer, out *LoadWordsResponse) {
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
		case "data":
			out.Data = string(in.String())
		case "error":
			out.Error = bool(in.Bool())
		case "errorText":
			out.ErrorText = string(in.String())
		case "additionalErrors":
			if in.IsNull() {
				in.Skip()
				out.AdditionalErrors = nil
			} else {
				if out.AdditionalErrors == nil {
					out.AdditionalErrors = new(AdditionalErrors)
				}
				(*out.AdditionalErrors).UnmarshalEasyJSON(in)
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
func easyjson102f8a2fEncodeAnagramsvcPkgApiV1(out *jwriter.Writer, in LoadWordsResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"data\":"
		out.RawString(prefix[1:])
		out.String(string(in.Data))
	}
	{
		const prefix string = ",\"error\":"
		out.RawString(prefix)
		out.Bool(bool(in.Error))
	}
	{
		const prefix string = ",\"errorText\":"
		out.RawString(prefix)
		out.String(string(in.ErrorText))
	}
	{
		const prefix string = ",\"additionalErrors\":"
		out.RawString(prefix)
		if in.AdditionalErrors == nil {
			out.RawString("null")
		} else {
			(*in.AdditionalErrors).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v LoadWordsResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson102f8a2fEncodeAnagramsvcPkgApiV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v LoadWordsResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson102f8a2fEncodeAnagramsvcPkgApiV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *LoadWordsResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson102f8a2fDecodeAnagramsvcPkgApiV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *LoadWordsResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson102f8a2fDecodeAnagramsvcPkgApiV1(l, v)
}
func easyjson102f8a2fDecodeAnagramsvcPkgApiV11(in *jlexer.Lexer, out *GetAnagramsResponse) {
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
		case "data":
			if in.IsNull() {
				in.Skip()
				out.Data = nil
			} else {
				if out.Data == nil {
					out.Data = new([]string)
				}
				if in.IsNull() {
					in.Skip()
					*out.Data = nil
				} else {
					in.Delim('[')
					if *out.Data == nil {
						if !in.IsDelim(']') {
							*out.Data = make([]string, 0, 4)
						} else {
							*out.Data = []string{}
						}
					} else {
						*out.Data = (*out.Data)[:0]
					}
					for !in.IsDelim(']') {
						var v1 string
						v1 = string(in.String())
						*out.Data = append(*out.Data, v1)
						in.WantComma()
					}
					in.Delim(']')
				}
			}
		case "error":
			out.Error = bool(in.Bool())
		case "errorText":
			out.ErrorText = string(in.String())
		case "additionalErrors":
			if in.IsNull() {
				in.Skip()
				out.AdditionalErrors = nil
			} else {
				if out.AdditionalErrors == nil {
					out.AdditionalErrors = new(AdditionalErrors)
				}
				(*out.AdditionalErrors).UnmarshalEasyJSON(in)
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
func easyjson102f8a2fEncodeAnagramsvcPkgApiV11(out *jwriter.Writer, in GetAnagramsResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"data\":"
		out.RawString(prefix[1:])
		if in.Data == nil {
			out.RawString("null")
		} else {
			if *in.Data == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
				out.RawString("null")
			} else {
				out.RawByte('[')
				for v2, v3 := range *in.Data {
					if v2 > 0 {
						out.RawByte(',')
					}
					out.String(string(v3))
				}
				out.RawByte(']')
			}
		}
	}
	{
		const prefix string = ",\"error\":"
		out.RawString(prefix)
		out.Bool(bool(in.Error))
	}
	{
		const prefix string = ",\"errorText\":"
		out.RawString(prefix)
		out.String(string(in.ErrorText))
	}
	{
		const prefix string = ",\"additionalErrors\":"
		out.RawString(prefix)
		if in.AdditionalErrors == nil {
			out.RawString("null")
		} else {
			(*in.AdditionalErrors).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetAnagramsResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson102f8a2fEncodeAnagramsvcPkgApiV11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetAnagramsResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson102f8a2fEncodeAnagramsvcPkgApiV11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetAnagramsResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson102f8a2fDecodeAnagramsvcPkgApiV11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetAnagramsResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson102f8a2fDecodeAnagramsvcPkgApiV11(l, v)
}
func easyjson102f8a2fDecodeAnagramsvcPkgApiV12(in *jlexer.Lexer, out *AdditionalErrors) {
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
		case "name":
			out.Name = string(in.String())
		case "lastName":
			out.LastName = string(in.String())
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
func easyjson102f8a2fEncodeAnagramsvcPkgApiV12(out *jwriter.Writer, in AdditionalErrors) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"lastName\":"
		out.RawString(prefix)
		out.String(string(in.LastName))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AdditionalErrors) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson102f8a2fEncodeAnagramsvcPkgApiV12(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AdditionalErrors) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson102f8a2fEncodeAnagramsvcPkgApiV12(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AdditionalErrors) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson102f8a2fDecodeAnagramsvcPkgApiV12(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AdditionalErrors) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson102f8a2fDecodeAnagramsvcPkgApiV12(l, v)
}

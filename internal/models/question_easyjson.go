// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

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

func easyjson78ba5d84Decode20241TeaStealersInternalModels(in *jlexer.Lexer, out *QuestionResp) {
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
		case "id":
			out.ID = int64(in.Int64())
		case "question_text":
			out.QuestionText = string(in.String())
		case "max_mark":
			out.MaxMark = int64(in.Int64())
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
func easyjson78ba5d84Encode20241TeaStealersInternalModels(out *jwriter.Writer, in QuestionResp) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.ID))
	}
	{
		const prefix string = ",\"question_text\":"
		out.RawString(prefix)
		out.String(string(in.QuestionText))
	}
	{
		const prefix string = ",\"max_mark\":"
		out.RawString(prefix)
		out.Int64(int64(in.MaxMark))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v QuestionResp) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson78ba5d84Encode20241TeaStealersInternalModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v QuestionResp) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson78ba5d84Encode20241TeaStealersInternalModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *QuestionResp) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson78ba5d84Decode20241TeaStealersInternalModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *QuestionResp) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson78ba5d84Decode20241TeaStealersInternalModels(l, v)
}
func easyjson78ba5d84Decode20241TeaStealersInternalModels1(in *jlexer.Lexer, out *Question) {
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
		case "id":
			out.ID = int64(in.Int64())
		case "question_text":
			out.QuestionText = string(in.String())
		case "question_theme":
			out.Theme = QuestionTheme(in.String())
		case "max_mark":
			out.MaxMark = int64(in.Int64())
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
func easyjson78ba5d84Encode20241TeaStealersInternalModels1(out *jwriter.Writer, in Question) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.ID))
	}
	{
		const prefix string = ",\"question_text\":"
		out.RawString(prefix)
		out.String(string(in.QuestionText))
	}
	{
		const prefix string = ",\"question_theme\":"
		out.RawString(prefix)
		out.String(string(in.Theme))
	}
	{
		const prefix string = ",\"max_mark\":"
		out.RawString(prefix)
		out.Int64(int64(in.MaxMark))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Question) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson78ba5d84Encode20241TeaStealersInternalModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Question) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson78ba5d84Encode20241TeaStealersInternalModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Question) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson78ba5d84Decode20241TeaStealersInternalModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Question) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson78ba5d84Decode20241TeaStealersInternalModels1(l, v)
}

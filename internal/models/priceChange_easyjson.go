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

func easyjson4d8c9291Decode20241TeaStealersInternalModels(in *jlexer.Lexer, out *PriceChangeData) {
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
		case "price":
			out.Price = int64(in.Int64())
		case "data":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.DateCreation).UnmarshalJSON(data))
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
func easyjson4d8c9291Encode20241TeaStealersInternalModels(out *jwriter.Writer, in PriceChangeData) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"price\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.Price))
	}
	{
		const prefix string = ",\"data\":"
		out.RawString(prefix)
		out.Raw((in.DateCreation).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PriceChangeData) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4d8c9291Encode20241TeaStealersInternalModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PriceChangeData) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4d8c9291Encode20241TeaStealersInternalModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PriceChangeData) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4d8c9291Decode20241TeaStealersInternalModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PriceChangeData) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4d8c9291Decode20241TeaStealersInternalModels(l, v)
}
func easyjson4d8c9291Decode20241TeaStealersInternalModels1(in *jlexer.Lexer, out *PriceChange) {
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
		case "advertId":
			out.AdvertID = int64(in.Int64())
		case "price":
			out.Price = int64(in.Int64())
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
func easyjson4d8c9291Encode20241TeaStealersInternalModels1(out *jwriter.Writer, in PriceChange) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.ID))
	}
	{
		const prefix string = ",\"advertId\":"
		out.RawString(prefix)
		out.Int64(int64(in.AdvertID))
	}
	{
		const prefix string = ",\"price\":"
		out.RawString(prefix)
		out.Int64(int64(in.Price))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PriceChange) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4d8c9291Encode20241TeaStealersInternalModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PriceChange) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4d8c9291Encode20241TeaStealersInternalModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PriceChange) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4d8c9291Decode20241TeaStealersInternalModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PriceChange) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4d8c9291Decode20241TeaStealersInternalModels1(l, v)
}
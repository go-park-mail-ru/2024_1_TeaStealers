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

func easyjsonB2c7a8edDecode20241TeaStealersInternalModels(in *jlexer.Lexer, out *StatisticViewAdvert) {
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
			if data := in.UnsafeBytes(); in.Ok() {
				in.AddError((out.ID).UnmarshalText(data))
			}
		case "userId":
			if data := in.UnsafeBytes(); in.Ok() {
				in.AddError((out.UserID).UnmarshalText(data))
			}
		case "advertId":
			if data := in.UnsafeBytes(); in.Ok() {
				in.AddError((out.AdvertID).UnmarshalText(data))
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
func easyjsonB2c7a8edEncode20241TeaStealersInternalModels(out *jwriter.Writer, in StatisticViewAdvert) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.RawText((in.ID).MarshalText())
	}
	{
		const prefix string = ",\"userId\":"
		out.RawString(prefix)
		out.RawText((in.UserID).MarshalText())
	}
	{
		const prefix string = ",\"advertId\":"
		out.RawString(prefix)
		out.RawText((in.AdvertID).MarshalText())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v StatisticViewAdvert) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB2c7a8edEncode20241TeaStealersInternalModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v StatisticViewAdvert) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB2c7a8edEncode20241TeaStealersInternalModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *StatisticViewAdvert) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB2c7a8edDecode20241TeaStealersInternalModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *StatisticViewAdvert) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB2c7a8edDecode20241TeaStealersInternalModels(l, v)
}

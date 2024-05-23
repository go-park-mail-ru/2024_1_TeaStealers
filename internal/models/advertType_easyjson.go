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

func easyjson41dc4deeDecode20241TeaStealersInternalModels(in *jlexer.Lexer, out *HouseTypeAdvert) {
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
		case "houseId":
			out.HouseID = int64(in.Int64())
		case "advertId":
			out.AdvertID = int64(in.Int64())
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
func easyjson41dc4deeEncode20241TeaStealersInternalModels(out *jwriter.Writer, in HouseTypeAdvert) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"houseId\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.HouseID))
	}
	{
		const prefix string = ",\"advertId\":"
		out.RawString(prefix)
		out.Int64(int64(in.AdvertID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v HouseTypeAdvert) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson41dc4deeEncode20241TeaStealersInternalModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v HouseTypeAdvert) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson41dc4deeEncode20241TeaStealersInternalModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *HouseTypeAdvert) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson41dc4deeDecode20241TeaStealersInternalModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *HouseTypeAdvert) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson41dc4deeDecode20241TeaStealersInternalModels(l, v)
}
func easyjson41dc4deeDecode20241TeaStealersInternalModels1(in *jlexer.Lexer, out *FlatTypeAdvert) {
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
		case "flatId":
			out.FlatID = int64(in.Int64())
		case "advertId":
			out.AdvertID = int64(in.Int64())
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
func easyjson41dc4deeEncode20241TeaStealersInternalModels1(out *jwriter.Writer, in FlatTypeAdvert) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"flatId\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.FlatID))
	}
	{
		const prefix string = ",\"advertId\":"
		out.RawString(prefix)
		out.Int64(int64(in.AdvertID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FlatTypeAdvert) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson41dc4deeEncode20241TeaStealersInternalModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FlatTypeAdvert) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson41dc4deeEncode20241TeaStealersInternalModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FlatTypeAdvert) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson41dc4deeDecode20241TeaStealersInternalModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FlatTypeAdvert) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson41dc4deeDecode20241TeaStealersInternalModels1(l, v)
}
func easyjson41dc4deeDecode20241TeaStealersInternalModels2(in *jlexer.Lexer, out *AdvertType) {
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
		case "advertType":
			out.AdvertType = AdvertTypeAdvert(in.String())
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
func easyjson41dc4deeEncode20241TeaStealersInternalModels2(out *jwriter.Writer, in AdvertType) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.RawText((in.ID).MarshalText())
	}
	{
		const prefix string = ",\"advertType\":"
		out.RawString(prefix)
		out.String(string(in.AdvertType))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AdvertType) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson41dc4deeEncode20241TeaStealersInternalModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AdvertType) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson41dc4deeEncode20241TeaStealersInternalModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AdvertType) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson41dc4deeDecode20241TeaStealersInternalModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AdvertType) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson41dc4deeDecode20241TeaStealersInternalModels2(l, v)
}
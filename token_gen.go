package coby

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Token) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "t":
			z.Token, err = dc.ReadString()
			if err != nil {
				return
			}
		case "d":
			var xsz uint32
			xsz, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Data) >= int(xsz) {
				z.Data = z.Data[:xsz]
			} else {
				z.Data = make([]string, xsz)
			}
			for xvk := range z.Data {
				z.Data[xvk], err = dc.ReadString()
				if err != nil {
					return
				}
			}
		case "e":
			z.Expire, err = dc.ReadInt64()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Token) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteMapHeader(3)
	if err != nil {
		return
	}
	err = en.WriteString("t")
	if err != nil {
		return
	}
	err = en.WriteString(z.Token)
	if err != nil {
		return
	}
	err = en.WriteString("d")
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Data)))
	if err != nil {
		return
	}
	for xvk := range z.Data {
		err = en.WriteString(z.Data[xvk])
		if err != nil {
			return
		}
	}
	err = en.WriteString("e")
	if err != nil {
		return
	}
	err = en.WriteInt64(z.Expire)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Token) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendMapHeader(o, 3)
	o = msgp.AppendString(o, "t")
	o = msgp.AppendString(o, z.Token)
	o = msgp.AppendString(o, "d")
	o = msgp.AppendArrayHeader(o, uint32(len(z.Data)))
	for xvk := range z.Data {
		o = msgp.AppendString(o, z.Data[xvk])
	}
	o = msgp.AppendString(o, "e")
	o = msgp.AppendInt64(o, z.Expire)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Token) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "t":
			z.Token, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "d":
			var xsz uint32
			xsz, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Data) >= int(xsz) {
				z.Data = z.Data[:xsz]
			} else {
				z.Data = make([]string, xsz)
			}
			for xvk := range z.Data {
				z.Data[xvk], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			}
		case "e":
			z.Expire, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

func (z *Token) Msgsize() (s int) {
	s = msgp.MapHeaderSize + msgp.StringPrefixSize + 1 + msgp.StringPrefixSize + len(z.Token) + msgp.StringPrefixSize + 1 + msgp.ArrayHeaderSize
	for xvk := range z.Data {
		s += msgp.StringPrefixSize + len(z.Data[xvk])
	}
	s += msgp.StringPrefixSize + 1 + msgp.Int64Size
	return
}

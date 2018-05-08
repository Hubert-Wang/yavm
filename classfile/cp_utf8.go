package classfile

import (
	"fmt"
	"unicode/utf16"
)

/*
CONSTANT_Utf8_info {
    u1 tag;
    u2 length;
    u1 bytes[length];
}
*/

type ConstantUtf8Info struct {
	str string
}

func (self *ConstantUtf8Info) readInfo(reader *ClassReader)  {
	length := uint32(reader.readUint32())
	bytes := reader.readBytes(length)
	self.str = decodeMUTF8(bytes)
}

func (self *ConstantUtf8Info) Value() string {
	return self.str
}

/*
func decodeMUTF8(bytes []byte) string {
	return string(bytes) // not correct!
}
*/

// mutf8 -> utf16 -> utf32 -> string
// see java.io.DataInputStream.readUTF(DataInput)

func decodeMUTF8(bytes []byte) string {
	utflen := len(bytes)
	chars := make([]uint16, utflen)

	var c, char2, char3 uint16
	count :=0
	char_count := 0

	for count < utflen {
		c = uint16(bytes[count])
		if c > 127 {
			break
		}
		count++
		chars[char_count] = c
		char_count++
	}

	for count < utflen {
		c = uint16(bytes[count])
		switch c >> 4 {
		case 0, 1, 2, 3, 4, 5, 6, 7:
			/* 0xxxxxxx*/
			count++
			chars[char_count] = c
			char_count++
		case 12, 13:
			/* 110x xxxx   10xx xxxx*/
			count += 2
			if count > utflen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(chars[count-1])
			if char2&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", count))
			}
			chars[char_count] = c&0x1F<<6 | char2&0x3F
			char_count++
		case 14:
			/* 1110 xxxx  10xx xxxx  10xx xxxx*/
			count += 3
			if count > utflen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(bytes[count-2])
			char3 = uint16(bytes[count-1])
			if char2&0xC0 != 0x80 || char3&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", (count - 1)))
			}
			chars[char_count] = c&0x0F<<12 | char2&0x3F<<6 | char3&0x3F<<0
			char_count++
		default:
			/* 10xx xxxx,  1111 xxxx */
			panic(fmt.Errorf("malformed input around byte %v", count))
		}
		}
	chars = chars[0:char_count]
	runes := utf16.Decode(chars)
	return string(runes)
}


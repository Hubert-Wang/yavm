package classfile

import "fmt"

/*
ClassFile {
    u4             magic;
    u2             minor_version;
    u2             major_version;
    u2             constant_pool_count;
    cp_info        constant_pool[constant_pool_count-1];
    u2             access_flags;
    u2             this_class;
    u2             super_class;
    u2             interfaces_count;
    u2             interfaces[interfaces_count];
    u2             fields_count;
    field_info     fields[fields_count];
    u2             methods_count;
    method_info    methods[methods_count];
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/

type ClassFile struct {
	magic uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags uint16
	thisClass uint16
	superClass uint16
	interfaces [] uint16
	fields []*MemberInfo
	methods []*MemberInfo
	attributes []AttributeInfo

}

func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
	self.constantPool = readConstantPool(reader)
}

func(self *ClassFile) readAndCheckMagic(reader *ClassReader)  {
	self.magic = reader.readUint32()
	if self.magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError:invalid magic")
	}
}

func (self *ClassFile) readAndCheckVersion(reader *ClassReader)  {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()

	switch self.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 {
			return
		}

	panic("java.lang.UnsupportedClassVersionError!")
	}
}
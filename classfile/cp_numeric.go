package classfile


/*
CONSTANT_Integer_info {
    u1 tag;
    u4 bytes;
}
*/
type ConstantIntegerInfo struct {
	val int32
}

func (self *ConstantIntegerInfo) readInfo(reader *ClassReader)  {
	bytes := reader.readUint32()
	self.val = int32(bytes)
}

func (self *ConstantIntegerInfo) Value()  int32{
	return self.val
}


/*
CONSTANT_Float_info {
    u1 tag;
    u4 bytes;
}
*/
type ConstantFloatInfo struct {
	val float32
}

func (self *ConstantFloatInfo) readInfo(reader *ClassReader)  {
	bytes := reader.readUint32()
	self.val = float32(bytes)
}

func (self *ConstantFloatInfo) Value() float32  {
	return self.val
}


/*
CONSTANT_Long_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
*/
type ConstantLongInfo struct {
	val uint64
}

func (self *ConstantLongInfo) readInfo(reader *ClassReader)  {
	bytes := reader.readUint64()
	self.val = uint64(bytes)
}

func Value(self *ConstantLongInfo)  uint64{
	return self.val
}

/*
CONSTANT_Double_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
*/
type ConstantDoubleInfo struct {
	val float64
}

func (self *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = float64(bytes)
}

func (self *ConstantDoubleInfo) Value() float64 {
	return self.val
}

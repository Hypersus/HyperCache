

//===================================================================#
//	Copyright (C) 2022 Zeke. All rights reserved
// 
//	Filename:		byteview.go
//	Author:			Zeke
//	Date:			2022.08.27
//	E-mail:			hypersus@outlook.com
//	Discription:	test script
//	
//===================================================================#

package hypercache

type ByteView struct {
	b []byte
}

func (v ByteView) Len() int {
	return len(v.b)
}

func (v ByteView) String() string {
	return string(v.b)
}

func (v ByteView) cloneBytes(b []byte) []byte{
	c := make([]byte, len(b))
	copy(c, b)
	return c

}

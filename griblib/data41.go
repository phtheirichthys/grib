package griblib

import (
	"io"
)

type Data41 struct {
	Reference    float32 `json:"reference"`
	BinaryScale  uint16  `json:"binaryScale"`
	DecimalScale uint16  `json:"decimalScale"`
	Bits         uint8   `json:"bits"`
	Type         uint8   `json:"type"`
}

func ParseData41(dataReader io.Reader, dataLength int, template *Data41) []float64 {

	data0 := Data0{
		Reference:    template.Reference,
		BinaryScale:  template.BinaryScale,
		DecimalScale: template.DecimalScale,
		Type:         template.Type,
		Bits:         template.Bits,
	}
	return ParseData0(dataReader, dataLength, &data0)
}

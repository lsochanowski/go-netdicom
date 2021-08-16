package fuzzpdu

import (
	"bytes"
	"encoding/binary"
	"flag"

	"github.com/lsochanowski/go-dicom/dicomio"
	"github.com/lsochanowski/go-netdicom/dimse"
	"github.com/lsochanowski/go-netdicom/pdu"
)

func init() {
	flag.Parse()
}

func Fuzz(data []byte) int {
	in := bytes.NewBuffer(data)
	if len(data) == 0 || data[0] <= 0xc0 {
		pdu.ReadPDU(in, 4<<20) // nolint: errcheck
	} else {
		d := dicomio.NewDecoder(in, binary.LittleEndian, dicomio.ExplicitVR)
		dimse.ReadMessage(d)
	}
	return 0
}

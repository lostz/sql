package sql

import "errors"

const NamesSepChar byte = '\377'
const MysqlVersionID = 50109

const (
	LongLen             = 10
	LonglongLen         = 19
	SignedLonglongLen   = 19
	UnsignedLonglongLen = 20
)

var (
	Long             []byte = []byte{'2', '1', '4', '7', '4', '8', '3', '6', '4', '7'}
	SignedLong       []byte = []byte{'-', '2', '1', '4', '7', '4', '8', '3', '6', '4', '8'}
	Longlong         []byte = []byte{'9', '2', '2', '3', '3', '7', '2', '0', '3', '6', '8', '5', '4', '7', '7', '5', '8', '0', '7'}
	SignedLonglong   []byte = []byte{'-', '9', '2', '2', '3', '3', '7', '2', '0', '3', '6', '8', '5', '4', '7', '7', '5', '8', '0', '8'}
	UnsignedLonglong []byte = []byte{'1', '8', '4', '4', '6', '7', '4', '4', '0', '7', '3', '7', '0', '9', '5', '5', '1', '6', '1', '5'}
)
var ErrStringFormat = errors.New("text string format error")

package sql

import "errors"

const (
	LONG_LEN              = 10
	LONGLONG_LEN          = 19
	SIGNED_LONGLONG_LEN   = 19
	UNSIGNED_LONGLONG_LEN = 20
)

var (
	LONG              []byte = []byte{'2', '1', '4', '7', '4', '8', '3', '6', '4', '7'}
	SIGNED_LONG       []byte = []byte{'-', '2', '1', '4', '7', '4', '8', '3', '6', '4', '8'}
	LONGLONG          []byte = []byte{'9', '2', '2', '3', '3', '7', '2', '0', '3', '6', '8', '5', '4', '7', '7', '5', '8', '0', '7'}
	SIGNED_LONGLONG   []byte = []byte{'-', '9', '2', '2', '3', '3', '7', '2', '0', '3', '6', '8', '5', '4', '7', '7', '5', '8', '0', '8'}
	UNSIGNED_LONGLONG []byte = []byte{'1', '8', '4', '4', '6', '7', '4', '4', '0', '7', '3', '7', '0', '9', '5', '5', '1', '6', '1', '5'}
)
var ErrStringFormat = errors.New("text string format error")

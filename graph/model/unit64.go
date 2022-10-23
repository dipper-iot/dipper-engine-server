package model

import (
	"encoding/json"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"io"
	"strconv"
	"strings"
)

const bigIntUnsigned = uint64(1<<64/2 - 1)

func MarshalUint64(i uint64) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, fmt.Sprintf(`"%s"`, strconv.FormatUint(i, 10)))
	})
}

func UnmarshalUint64(v interface{}) (uint64, error) {
	switch v := v.(type) {
	case string:
		return stringToUint64(v)
	case int64:
		return uint64(v), nil
	case uint64:
		return v, nil
	case json.Number:
		stringNumber := string(v)
		return stringToUint64(stringNumber)
	default:
		return 0, fmt.Errorf("%T is not an uint64", v)
	}
}

func stringToUint64(v string) (uint64, error) {
	if v == "" {
		return 0, nil
	}
	if strings.Contains(v, "e") {
		return 0, fmt.Errorf("parse number failed because the number contains scientific notation (%s), which is not supported", v)
	}

	num, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		return 0, err
	}

	if num > bigIntUnsigned {
		return 0, fmt.Errorf("the number must be lower than %d", bigIntUnsigned)
	}

	if num < 0 {
		return 0, fmt.Errorf("cannot convert to uint64 because the value is negative")
	}

	return num, nil
}

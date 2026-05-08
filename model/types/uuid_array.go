package types

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type UUIDArray []uuid.UUID

func (a *UUIDArray) Scan(value interface{}) error {
	var str string

	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	default:
		return errors.New("Failed to parse UUIDArray: unsupported data type")
	}

	str = strings.TrimPrefix(str, "{")
	str = strings.TrimSuffix(str, "}")

	if str == "" {
		*a = nil
		return nil
	}

	parts := strings.Split(str, ",")

	*a = make(UUIDArray, 0, len(parts))

	for _, p := range parts {
		p = strings.TrimSpace(strings.Trim(p, `"`))
		if p == "" {
			continue
		}

		uuid, err := uuid.Parse(p)
		if err != nil {
			return fmt.Errorf("Invalid UUID in Array : %v", err)
		}

		*a = append(*a, uuid)
	}

	return nil
}

func (a UUIDArray) Value() (driver.Value, error) {
	if len(a) == 0 {
		return "{}", nil
	}

	postgreFormat := make([]string, 0, len(a))

	for _, value := range a {
		postgreFormat = append(postgreFormat, fmt.Sprintf(`"%s"`, value.String()))
	}

	return "{" + strings.Join(postgreFormat, ",") + "}", nil
}

func (UUIDArray) GormDataType() string {
	return "uuid[]"
}

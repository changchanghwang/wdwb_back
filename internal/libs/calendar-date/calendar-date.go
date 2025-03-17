package calendarDate

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"database/sql/driver"

	applicationError "github.com/changchanghwang/wdwb_back/pkg/application-error"
)

const calendarDateFormat = "2006-01-02"

// YYYY-MM-DD
// Can be empty string
type CalendarDate string

func (d CalendarDate) Parse() (time.Time, error) {
	return time.Parse(calendarDateFormat, string(d))
}

func (d CalendarDate) Year() (int, error) {
	splittedReportedOn := strings.Split(string(d), "-")
	year, err := strconv.Atoi(splittedReportedOn[0])
	if err != nil {
		return 0, applicationError.New(http.StatusBadRequest, fmt.Sprintf("invalid year: %s", splittedReportedOn[0]), "")
	}
	return year, nil
}

func (d CalendarDate) Quarter() (int, error) {
	quarterMap := map[int]int{
		1:  1,
		2:  1,
		3:  1,
		4:  2,
		5:  2,
		6:  2,
		7:  3,
		8:  3,
		9:  3,
		10: 4,
		11: 4,
		12: 4,
	}

	splittedReportedOn := strings.Split(string(d), "-")
	month, err := strconv.Atoi(splittedReportedOn[1])
	if err != nil {
		return 0, applicationError.New(http.StatusBadRequest, fmt.Sprintf("invalid month: %s", splittedReportedOn[1]), "")
	}

	return quarterMap[month], nil
}

func (d *CalendarDate) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return applicationError.New(http.StatusBadRequest, fmt.Sprintf("invalid date format: %v", err), "")
	}

	if s == "" {
		*d = ""
		return nil
	}

	if _, err := time.Parse(calendarDateFormat, s); err != nil {
		return applicationError.New(http.StatusBadRequest, fmt.Sprintf("invalid date format: %v", err), "")
	}

	*d = CalendarDate(s)
	return nil
}

func (d CalendarDate) MarshalJSON() ([]byte, error) {
	if d == "" {
		return json.Marshal("")
	}

	if _, err := time.Parse(calendarDateFormat, string(d)); err != nil {
		return nil, applicationError.New(http.StatusBadRequest, fmt.Sprintf("invalid date format: %v", err), "")
	}

	return json.Marshal(string(d))
}

// Scan implements the sql.Scanner interface for CalendarDate.
// It converts a database value (typically a time.Time) into a CalendarDate.
func (d *CalendarDate) Scan(value interface{}) error {
	if value == nil {
		*d = ""
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		*d = CalendarDate(v.Format(calendarDateFormat))
		return nil
	case []byte:
		str := string(v)
		if _, err := time.Parse(calendarDateFormat, str); err != nil {
			return fmt.Errorf("invalid date format: %v", err)
		}
		*d = CalendarDate(str)
		return nil
	case string:
		if _, err := time.Parse(calendarDateFormat, v); err != nil {
			return fmt.Errorf("invalid date format: %v", err)
		}
		*d = CalendarDate(v)
		return nil
	default:
		return fmt.Errorf("unsupported type %T for CalendarDate", value)
	}
}

// Value implements the driver.Valuer interface for CalendarDate.
// It converts a CalendarDate into a format that can be stored in the database.
func (d CalendarDate) Value() (driver.Value, error) {
	if d == "" {
		return nil, nil
	}

	t, err := time.Parse(calendarDateFormat, string(d))
	if err != nil {
		return nil, fmt.Errorf("invalid date format: %v", err)
	}
	return t, nil
}

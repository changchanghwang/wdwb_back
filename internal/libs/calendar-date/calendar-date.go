package calendarDate

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

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

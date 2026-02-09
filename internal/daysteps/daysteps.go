package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

var ErrNotPositive = errors.New("must be positive")

// Parse parses training data from string.
func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return fmt.Errorf("invalid data format: expected 'steps,duration', got %q", datastring)
	}
	ds.Steps, err = strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("invalid steps count, expected numbers, got %q: %w", parts[0], err)
	}
	if ds.Steps <= 0 {
		return fmt.Errorf("steps %d: %w", ds.Steps, ErrNotPositive)
	}

	ds.Duration, err = time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("invalid duration: expected format 30m, 1h, 1h30m, 10s, got %q: %w", parts[1], err)
	}
	if ds.Duration <= 0 {
		return fmt.Errorf("duration %v: %w", ds.Duration, ErrNotPositive)
	}
	return nil
}

// ActionInfo returns formatted information about daily walking activity.
func (ds DaySteps) ActionInfo() (string, error) {
	dist := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, dist, calories), nil
}

package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

var ErrInvalidFormat = errors.New("invalid format")
var ErrNotPositive = errors.New("must be positive")

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// Parse parses training data from string.
func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return fmt.Errorf("invalid data format, expected 'steps,activityType,duration', got: %q: %w", datastring, ErrInvalidFormat)
	}

	t.Steps, err = strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("invalid steps count, expected numbers, got %q: %w", parts[0], ErrInvalidFormat)
	}
	if t.Steps <= 0 {
		return fmt.Errorf("steps %d: %w", t.Steps, ErrNotPositive)
	}

	t.TrainingType = parts[1]

	t.Duration, err = time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("invalid time format, expected 30m, 1h, 1h40m, 45s, got %q: %w", parts[2], ErrInvalidFormat)
	}
	if t.Duration <= 0 {
		return fmt.Errorf("duration %v: %w", t.Duration, ErrNotPositive)
	}
	return nil
}

// ActionInfo returns formatted training data based on activity type.
func (t Training) ActionInfo() (string, error) {
	dist := spentenergy.Distance(t.Steps, t.Height)
	avgSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	switch t.TrainingType {
	case "Бег":
		calories, err := spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			t.TrainingType, t.Duration.Hours(), dist, avgSpeed, calories), nil

	case "Ходьба":
		calories, err := spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			t.TrainingType, t.Duration.Hours(), dist, avgSpeed, calories), nil

	default:
		return "", errors.New("неизвестный тип тренировки")
	}
}

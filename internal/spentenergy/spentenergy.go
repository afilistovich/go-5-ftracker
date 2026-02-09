package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("steps count must be positive")
	}
	if weight <= 0 {
		return 0, errors.New("weight must be positive")
	}
	if height <= 0 {
		return 0, errors.New("height must be positive")
	}
	if duration <= 0 {
		return 0, errors.New("time duration must be positive")
	}

	averageSpeed := MeanSpeed(steps, height, duration)
	calories := (duration.Minutes() * averageSpeed * weight) / minInH
	return calories * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("steps count must be positive")
	}
	if weight <= 0 {
		return 0, errors.New("weight must be positive")
	}
	if height <= 0 {
		return 0, errors.New("height must be positive")
	}
	if duration <= 0 {
		return 0, errors.New("time duration must be positive")
	}

	averageSpeed := MeanSpeed(steps, height, duration)
	calories := (duration.Minutes() * averageSpeed * weight) / minInH

	return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 || steps <= 0 {
		return 0
	}
	return (Distance(steps, height)) / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	dist := (stepLength * float64(steps)) / mInKm
	return dist
}

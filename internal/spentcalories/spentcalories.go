package spentcalories

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	slice := strings.Split(data, ",")
	if len(slice) != 3 {
		return 0, "", 0, nil
	}
	steps, err := strconv.Atoi(slice[0])
	if err != nil {
		return 0, "", 0, err
	}
	duration, err := time.ParseDuration(slice[2])
	if err != nil {
		return 0, "", 0, err
	}
	return steps, slice[1], duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := height * stepLengthCoefficient
	return (stepLength * float64(steps)) / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	dis := distance(steps, height)
	averageSpeed := dis / duration.Hours()
	return averageSpeed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, typeActivity, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
	}
	switch typeActivity {
	case "Ходьба":
		walkDistance := distance(steps, height)
		walkAverageSpeed := meanSpeed(steps, height, duration)
		walkCalories, _ := WalkingSpentCalories(steps, weight, height, duration)
		message := fmt.Sprintf("Тип тренировки: %s \nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f/n", typeActivity, duration.Hours(), walkDistance, walkAverageSpeed, walkCalories)
		return message, nil
	case "Бег":
		runDistance := distance(steps, height)
		runAverageSpeed := meanSpeed(steps, height, duration)
		runCalories, _ := RunningSpentCalories(steps, weight, height, duration)
		message := fmt.Sprintf("Тип тренировки: %s \nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", typeActivity, duration.Hours(), runDistance, runAverageSpeed, runCalories)
		return message, nil
	}
	return "", errors.New("неизвестный тип тренировки")
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	switch {
	case steps <= 0:
		return 0, errors.New("шагов должно быть больше нуля")
	case weight <= 0:
		return 0, errors.New("вес должен быть больше нуля")
	case height <= 0:
		return 0, errors.New("рост должен быть больше нуля")
	case duration <= 0:
		return 0, errors.New("время должно быть больше нуля")
	}
	averageSpeed := meanSpeed(steps, height, duration)
	return (weight * averageSpeed * duration.Minutes()) / minInH, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	switch {
	case steps <= 0, weight <= 0, height <= 0, duration <= 0:
		return 0, errors.New("все параметры должны быть больше нуля")
	}
	averageSpeed := meanSpeed(steps, height, duration)
	calories := (weight * averageSpeed * duration.Minutes()) / minInH
	return calories * walkingCaloriesCoefficient, nil
}

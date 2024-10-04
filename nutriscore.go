package main

type ScoreType int

// Enum
const (
	Food ScoreType = iota
	Beverage
	Water
	Cheese
)

// Output
type NutritionalScore struct {
	Value     int
	Positive  int
	Negative  int
	ScoreType ScoreType
}

// Units
type EnergyKJ float64
type SugarGram float64
type SaturatedFattyAcidsGram float64
type SodiumMilligram float64
type FruitsPercent float64
type FiberGram float64
type ProteinGram float64

// Input
type NutritionalData struct {
	Energy              EnergyKJ
	Sugars              SugarGram
	SaturatedFattyAcids SaturatedFattyAcidsGram
	Sodium              SodiumMilligram
	Fruits              FruitsPercent
	Fiber               FiberGram
	Protein             ProteinGram
	IsWater             bool
}

// Food steps
var energyLevels = []float64{3350, 3015, 2680, 2345, 2010, 1675, 1340, 1005, 670, 335}
var sugarsLevels = []float64{45, 40, 36, 31, 27, 22.5, 18, 13.5, 9, 4.5}
var saturatedFattyAcidsLevels = []float64{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
var sodiumLevels = []float64{900, 810, 720, 630, 540, 450, 360, 270, 180, 90}
var fiberLevels = []float64{4.7, 3.7, 2.8, 1.9, 0.9}
var proteinLevels = []float64{8, 6.4, 4.8, 3.2, 1.6}

// Beverage steps
var energyLevelsBeverage = []float64{270, 240, 210, 180, 150, 120, 90, 60, 30, 0}
var sugarsLevelsBeverage = []float64{13.5, 12, 10.5, 9, 7.5, 6, 4.5, 3, 1.5, 0}

// GetPoints
func (e EnergyKJ) GetPoints(st ScoreType) int {

	if st == Beverage {
		return getPointsFromRange(float64(e), energyLevelsBeverage)
	}
	return getPointsFromRange(float64(e), energyLevels)

}
func (s SugarGram) GetPoints(st ScoreType) int {

	if st == Beverage {
		return getPointsFromRange(float64(s), sugarsLevelsBeverage)
	}
	return getPointsFromRange(float64(s), sugarsLevels)

}

func (sfa SaturatedFattyAcidsGram) GetPoints(st ScoreType) int {
	return getPointsFromRange(float64(sfa), saturatedFattyAcidsLevels)
}
func (s SodiumMilligram) GetPoints(st ScoreType) int {
	return getPointsFromRange(float64(s), sodiumLevels)
}

func (f FruitsPercent) GetPoints(st ScoreType) int {
	if st == Beverage {
		if f > 80 {
			return 10
		} else if f > 60 {
			return 4
		} else if f > 40 {
			return 2
		}
		return 0
	}
	if f > 80 {
		return 5
	} else if f > 60 {
		return 2
	} else if f > 40 {
		return 1
	}
	return 0

}
func (f FiberGram) GetPoints(st ScoreType) int {
	return getPointsFromRange(float64(f), fiberLevels)

}
func (p ProteinGram) GetPoints(st ScoreType) int {
	return getPointsFromRange(float64(p), proteinLevels)
}

// Cal for Energy
func EnergyFromKcal(kcal float64) EnergyKJ {
	return EnergyKJ(kcal * 4.184)

}

// Cal for Sodium
func SodiumFromSalt(saltMg float64) SodiumMilligram {
	return SodiumMilligram(saltMg / 2.5)
}

// Output Func
func GetNutritionalScore(n NutritionalData, st ScoreType) NutritionalScore {

	value := 0
	positive := 0
	negative := 0

	if st != Water {
		fruitPoints := n.Fruits.GetPoints(st)
		fiberPoints := n.Fiber.GetPoints(st)

		positive = fruitPoints + fiberPoints + n.Protein.GetPoints(st)
		negative = n.Energy.GetPoints(st) + n.Sugars.GetPoints(st) +
			n.SaturatedFattyAcids.GetPoints(st) + n.Sodium.GetPoints(st)

		if st == Cheese {
			value = negative - positive
		} else {
			if negative >= 11 && fruitPoints < 5 {
				value = negative - positive - fruitPoints
			} else {
				value = negative - positive
			}
		}
	}

	return NutritionalScore{
		Value:     value,
		Positive:  positive,
		Negative:  negative,
		ScoreType: st,
	}
}

// Grade
var scoreToLetter = []string{"A", "B", "C", "D", "E"}

func (ns NutritionalScore) GetNutriScore() string {

	if ns.ScoreType == Food {
		return scoreToLetter[getPointsFromRange(float64(ns.Value), []float64{18, 10, 2, -1})]
	}
	if ns.ScoreType == Water {
		return scoreToLetter[0]
	}
	return scoreToLetter[getPointsFromRange(float64(ns.Value), []float64{9, 5, 1, -2})]

}

func getPointsFromRange(v float64, steps []float64) int {
	lenSteps := len(steps)
	for i, l := range steps {
		if v > l {
			return lenSteps - i
		}
	}
	return 0
}

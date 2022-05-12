package square

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type CustomInt int

const SidesCircle = 0
const SidesSquare = 4
const SidesTriangle = 3

func CalcSquare(sideLen float64, sidesNum CustomInt) float64 {
	result := 0.0
	if sidesNum == SidesSquare {
		result = sideLen * sideLen
	} else if sidesNum == SidesCircle {
		result = math.Pi * math.Pow(sideLen, 2)
	} else if sidesNum == SidesTriangle {
		result = 0.5 * sideLen * sideLen * math.Sqrt(3) / 2
	}
	return result
}

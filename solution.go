package square
import "math"
// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

const Pi float64 = 3.141592
type intCustomType int
const SidesTriangle intCustomType = 3
const SidesSquare intCustomType = 4
const SidesCircle intCustomType = 0

func CalcSquare(sideLen float64, sidesNum intCustomType) float64 {
	if sidesNum == SidesCircle {
		return sideLen * sideLen * Pi
	} else if sidesNum == SidesSquare {
		return sideLen * sideLen
	} else if sidesNum == SidesTriangle{
		return math.Sqrt(3)/4*sideLen*sideLen
	} else {
		return 0
	}
}

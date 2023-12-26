package conversion
import(
	"strconv"
	"errors"
)
func StringToFoat(strings []string) ([]float64, error) {
	var floats []float64
	for _, string:= range strings {
		floatVal, err := strconv.ParseFloat(string, 64)
		if err!= nil {
			return nil, errors.New("string to failed failed")
		}
		floats=append(floats, floatVal)
	}
	return floats, nil
}


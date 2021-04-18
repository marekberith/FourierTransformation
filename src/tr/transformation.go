package tr

import (
	"log"
	"math"
	"math/cmplx"
	"strconv"
	"strings"
)

type Vector struct {
	number complex128
}

func parseVectorIFFT(text string) []Vector {
	w := &text
	text = strings.Replace(text, " ", "", -1)
	tmp := strings.Split(*w, ",")
	values := make([]Vector, 0, len(tmp))
	for _, raw := range tmp {
		v, err := strconv.ParseComplex(raw, 128)
		if err != nil {
			log.Print(err)
			continue
		}
		var vct Vector
		vct.number = v
		values = append(values, vct)
	}
	return values
}

func parseVectorFFT(text string) []Vector {
	w := &text
	text = strings.Replace(text, " ", "", -1)
	tmp := strings.Split(*w, ",")
	values := make([]Vector, 0, len(tmp))
	for _, raw := range tmp {
		v, err := strconv.Atoi(raw)
		if err != nil {
			log.Print(err)
			continue
		}
		var vct Vector
		vct.number = complex(float64(v), 0)
		values = append(values, vct)
	}
	return values
}

func FFT(vec []Vector) []Vector {
	n := len(vec)
	if n == 1 {
		return vec
	}
	cmplPow := complex(0, 1) * complex(2.0 * math.Pi / float64(n), 0)
	omegaN := cmplx.Pow(complex(math.E, 0), cmplPow)
	omega := 1 + 0i

	limit := n / 2
	vecL := make([]Vector, limit)
	vecR := make([]Vector, limit)
	for i := 0; i < n; i++ {
		if i % 2 == 0 {
			vecL[i / 2].number = vec[i].number
		} else {
			vecR[i / 2].number = vec[i].number
		}
	}
	vecL = FFT(vecL)
	vecR = FFT(vecR)

	result := make([]Vector, n)
	for i := 0; i < limit; i++ {
		result[i].number = vecL[i].number + omega * vecR[i].number
		result[i + limit].number = vecL[i].number - omega * vecR[i].number
		omega *= omegaN
	}
	return result
}

func IFFT(vec []Vector) []Vector {
	n := len(vec)
	if n == 1 {
		return vec
	}
	cmplPow := complex(0, -1) * complex(2.0 * math.Pi / float64(n), 0)
	omegaN := cmplx.Pow(complex(math.E, 0), cmplPow)
	omega := 1 + 0i

	limit := n / 2
	vecL := make([]Vector, limit)
	vecR := make([]Vector, limit)
	for i := 0; i < n; i++ {
		if i % 2 == 0 {
			vecL[i / 2].number = vec[i].number
		} else {
			vecR[i / 2].number = vec[i].number
		}
	}
	vecL = IFFT(vecL)
	vecR = IFFT(vecR)

	result := make([]Vector, n)
	for i := 0; i < limit; i++ {
		result[i].number = vecL[i].number + omega * vecR[i].number
		result[i + limit].number = vecL[i].number - omega * vecR[i].number
		omega *= omegaN
	}
	return result
}


func TransformVectorFFT(text string) string {
	vec := parseVectorFFT(text)
	result := FFT(vec)
	limit := len(result)
	resultString := "("
	for i := 0; i < limit; i++ {
		resultString += strconv.FormatComplex(result[i].number, 'f', -1, 64)
		if i != limit - 1 {
			resultString += ","
		}
	}
	resultString += ")"
	return resultString
}

func TransformVectorIFFT(text string) string {
	vec := parseVectorIFFT(text)
	result := IFFT(vec)
	limit := len(result)
	resultString := "("
	for i := 0; i < limit; i++ {
		resultString += strconv.FormatFloat(real(result[i].number), 'f', -1, 64)
		if i != limit - 1 {
			resultString += ","
		}
	}
	resultString += ")"
	return resultString
}

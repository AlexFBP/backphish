package common

import (
	"strings"

	"github.com/brianvoe/gofakeit"
)

// From https://www.infobae.com/colombia/2023/06/27/estos-son-los-10-apellidos-mas-comunes-en-colombia/

var colombiaLastNames = []string{
	"Rodríguez",
	"Martínez",
	"García",
	"Gómez",
	"López",
	"Gonzáles",
	"Hernández",
	"Sánchez",
	"Pérez",
	"Ramírez",
}

// From https://pasaportecolombiano.wordpress.com/2014/08/17/los-nombres-y-apellidos-mas-comunes-en-colombia/

var colombiaFirstNamesMale = []string{
	"José",
	"Luís",
	"Carlos",
	"Juan",
}
var colombiaMiddleNamesMale = []string{
	"Antonio",
	"De Jesús",
	"Alberto",
	"Enrique",
	"María",
}
var colombiaFirstMidNamesMale = []string{
	"Juan Carlos",
	"Luis Alberto",
	"Carlos Alberto",
	"Luis Eduardo",
}
var colombiaFullNamesMale = []string{
	"Luis Alberto Rodríguez",
	"Luis Eduardo Rodríguez",
	"José Antonio Rodríguez",
	"Luis Alberto González",
	"Luis Eduardo Ramírez",
	"Luis Alberto Sánchez",
	"José Antonio Martínez",
	"Luis Carlos Rodríguez",
	"Luis Eduardo Gómez",
	"Jaime Rodríguez",
	"Alfonso Rodríguez",
	"José Rodríguez",
	"Rafael Rodríguez",
	"Álvaro Hernández",
	"Jaime Ramírez",
	"Pedro Rodríguez",
	"Carlos Rodríguez",
	"Jaime García",
	"Alfonso López",
}

var colombiaFirstNamesFemale = []string{
	"María",
	"Luz",
	"Ana",
}
var colombiaFirstMidNamesFemale = []string{
	"Luz Marina",
	"María del Carmen",
	"Sandra Milena",
	"Ana María",
}
var colombiaFullNamesFemale = []string{
	"María del Carmen Rodríguez",
	"Luz Marina Rodríguez",
	"Luz Marina Sánchez",
	"María del Carmen Sánchez",
	"Luz Marina González",
	"María del Carmen González",
	"Luz Marina García",
	"Luz Marina Gómez",
	"María del Carmen Gómez",
	"María del Carmen Hernández",
	"Rosalba Rodríguez",
	"María Rodríguez",
	"Mercedes Rodríguez",
	"María González",
	"Rosalba Hernández",
	"Isabel García",
	"Carmen Rodríguez",
	"Graciela Rodríguez",
	"Cecilia Rodríguez",
	"Margarita Rodríguez",
}

// Random Name Generation for Colombia

func GeneraNombresHombreCol() (s string) {
	if gofakeit.Bool() {
		s = PickRand(colombiaFirstNamesMale)
		if gofakeit.Bool() {
			s += " " + PickRand(colombiaMiddleNamesMale)
		}
	} else {
		s = PickRand(colombiaFirstMidNamesMale)
	}
	return
}

func GeneraNombresMujerCol() (s string) {
	if gofakeit.Bool() {
		s = PickRand(colombiaFirstNamesFemale)
	} else {
		s = PickRand(colombiaFirstMidNamesFemale)
	}
	return
}

func GeneraApellidosCol() string {
	return PickRand(colombiaLastNames)
}

func GeneraNombresApellidosPersonaCol(noAccents bool) (nombres, apellidos string) {
	if gofakeit.Bool() {
		// Male
		nombres = GeneraNombresHombreCol()
	} else {
		// Female
		nombres = GeneraNombresMujerCol()
	}
	apellidos = GeneraApellidosCol()

	if noAccents || gofakeit.Bool() {
		RemoveAccents(&nombres)
		RemoveAccents(&apellidos)
	}
	return
}

func GeneraNombresApellidosPersonaCombinadosCol(noAccents bool) (s string) {
	if gofakeit.Bool() {
		// Male
		if gofakeit.Bool() {
			s = GeneraNombresHombreCol() + " " + GeneraApellidosCol()
		} else {
			s = PickRand(colombiaFullNamesMale)
		}
	} else {
		// Female
		if gofakeit.Bool() {
			s = GeneraNombresMujerCol() + " " + GeneraApellidosCol()
		} else {
			s = PickRand(colombiaFullNamesFemale)
		}
	}

	if noAccents || gofakeit.Bool() {
		RemoveAccents(&s)
	}
	return
}

func RemoveAccents(s *string) {
	replacements := map[string]string{
		"á": "a",
		"é": "e",
		"í": "i",
		"ó": "o",
		"ú": "u",
		"ñ": PickRand([]string{"n", "nh"}),
	}
	for replace, with := range replacements {
		*s = strings.ReplaceAll(*s, replace, with)
	}
}

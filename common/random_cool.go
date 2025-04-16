/**
 * Random Colombian data generation functions
 */

package common

import (
	"fmt"
	"math/rand" // "crypto/rand"
	"strconv"
	"strings"

	"github.com/brianvoe/gofakeit"
)

func init() {
	iniciaIndices()
}

func GeneraNIPcolombia() string {
	minOldRange := int(1e3)
	oldRange := int(1e8) - minOldRange
	newRange := int(1e9)
	universe := oldRange + newRange
	id := rand.Intn(universe)
	if id >= oldRange {
		id += newRange - oldRange
	} else {
		id += minOldRange
	}
	return strconv.Itoa(id)
}

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

func GeneraCiudadDeptoColombia() (string, string) {
	cities := []struct {
		size float64
		cap  string
		dep  string
	}{ // FROM https://es.wikipedia.org/wiki/Anexo:Departamentos_de_Colombia_por_poblaci%C3%B3n
		{size: 8.02 + 3.50, cap: "Bogotá", dep: "Cundinamarca"},
		{size: 6.95, cap: "Medellin", dep: "Antioquia"},
		{size: 4.71, cap: "Cali", dep: "Valle del Cauca"},
		{size: 2.85, cap: "Barranquilla", dep: "Atlántico"},
		{size: 2.39, cap: "Bucaramanga", dep: "Santander"},
		{size: 2.29, cap: "Cartagena de Indias", dep: "Bolivar"},
		{size: 1.93, cap: "Montería", dep: "Córdoba"},
		{size: 1.72, cap: "San Juan de Pasto", dep: "Nariño"},
		{size: 1.72, cap: "Cúcuta", dep: "Norte de Santander"},
		{size: 1.58, cap: "Popayán", dep: "Cauca"},
		{size: 1.52, cap: "Santa Marta", dep: "Magdalena"},
		{size: 1.39, cap: "Ibague", dep: "Tolima"},
		{size: 1.37, cap: "Valledupar", dep: "Cesar"},
		{size: 1.32, cap: "Santiago de Tunja", dep: "Boyacá"},
		{size: 1.20, cap: "Neiva", dep: "Huila"},
		{size: 1.14, cap: "Villavicencio", dep: "Meta"},
		{size: 1.06, cap: "Manizales", dep: "Caldas"},
		{size: 1.06, cap: "Riohacha", dep: "La Guajira"},
		{size: 1.01, cap: "Sincelejo", dep: "Sucre"},
		{size: 0.986, cap: "Pereira", dep: "Risaralda"},
		{size: 0.604, cap: "Quibdó", dep: "Chocó"},
	}

	sum := 0.0
	for _, v := range cities {
		sum += v.size
	}
	selected := rand.Float64() * sum

	sum = 0.0
	selected_pos := 0
	for k, v := range cities {
		sum += v.size
		if selected <= sum {
			selected_pos = k
			break
		}
	}

	return cities[selected_pos].cap, cities[selected_pos].dep
}

// From https://wiki.openstreetmap.org/wiki/ES:Colombia/Gu%C3%ADa_para_mapear/n%C3%BAmeros_telef%C3%B3nicos
var colombiaRangosCelulares = [][]int{
	{300, 2000000, 9399999},
	{301, 2000000, 7999999},
	{302, 2000000, 4699999},
	{303, 2000000, 6849999},
	{304, 2000000, 3899999},
	{324, 2000000, 6999999},
	{305, 7000000, 9599999},
	{310, 2000000, 9999999},
	{311, 2000000, 9999999},
	{312, 2000000, 9999999},
	{313, 2000000, 9999999},
	{314, 2000000, 9999999},
	{320, 2000000, 9999999},
	{321, 2000000, 9999999},
	{322, 2000000, 9999999},
	{323, 2000000, 9999999},
	{315, 2000000, 9999999},
	{316, 2000000, 9999999},
	{317, 2000000, 9999999},
	{318, 2000000, 9999999},
	{319, 2000000, 9999999},
	{350, 2000000, 9399999},
	{351, 2000000, 9399999},
	{302, 4700000, 8699999},
	// {323,6000000,9999999}, // Range already covered
	{324, 1000000, 1999999},
	{324, 7000000, 9999999},
	{333, 0, 499999},
	{333, 6000000, 6999999},
}

var indicesRangosCelulares []int

func iniciaIndices() {
	if len(indicesRangosCelulares) == 0 {
		indicesRangosCelulares = make([]int, len(colombiaRangosCelulares))
		universo := 0 // Numbers that can be generated
		for k, v := range colombiaRangosCelulares {
			numeros := v[2] - v[1] + 1 // max - min + 1
			universo += numeros
			indicesRangosCelulares[k] = universo
		}
	}
}

func generaCelValido(pos int) (cel int, e error) {
	max := indicesRangosCelulares[len(indicesRangosCelulares)-1]
	if pos > max {
		return -1, fmt.Errorf("%d over %d", pos, max)
	} else if pos < 0 {
		return -1, fmt.Errorf("must be >= 0 - got: %d", pos)
	}
	for k, v := range indicesRangosCelulares {
		if pos > v {
			continue
		}
		rule := colombiaRangosCelulares[k]
		cel = rule[0] * int(1e7)
		cel += rule[1]
		if k > 0 {
			cel += pos - indicesRangosCelulares[k-1]
		} else {
			cel += pos
		}
		return cel, nil
	}
	return -1, fmt.Errorf("unhandled value: %d", pos)
}

func GeneraCelColombia() string {
	u := rand.Intn(indicesRangosCelulares[len(indicesRangosCelulares)-1])
	n, _ := generaCelValido(u)
	return strconv.Itoa(n)
}


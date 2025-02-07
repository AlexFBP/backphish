package nequi8

import (
	"github.com/AlexFBP/backphish/common"
)

type userData struct {
	Nip      string `json:"documentNumber"`
	FullName string `json:"fullName"`
	Phone    string `json:"username,omitempty"`
	Pass     string `json:"password,omitempty"`
	IP       string `json:"userIP"`
	City     string `json:"city"`
	Country  string `json:"country"`
	OTP      string `json:"otpCode,omitempty"`
	DynCode  string `json:"dynamicCode,omitempty"`
	Step     string `json:"step,omitempty"`
}

func newUser() (d userData) {
	d.Nip = common.GeneraNIPcolombia()
	d.FullName = common.GeneraNombresApellidosPersonaCombinadosCol(false)
	d.Phone = common.AddSeparator(common.GeneraCelColombia(), 0, " ")
	d.Pass = common.GeneraPin(4)
	d.IP = common.GeneraIP()
	d.City, _ = common.GeneraCiudadDeptoColombia()
	d.Country = "Colombia"
	return
}

func (d *userData) DataForStep(step uint8, includeStep bool) (nd userData) {
	nd.Nip, nd.FullName, nd.IP, nd.City, nd.Country = d.Nip, d.FullName, d.IP, d.City, d.Country
	if step >= 2 {
		nd.Phone = d.Phone
		nd.Pass = d.Pass
	}
	if step == 3 {
		nd.OTP = d.OTP
	}
	if step >= 4 {
		nd.DynCode = common.GeneraPin(6)
	}
	if includeStep {
		switch step {
		case 1:
			nd.Step = "ini"
		case 4:
			nd.Step = "dinamica2"
		case 5:
			nd.Step = "dinamica3"
		default:
		}
	}
	return
}

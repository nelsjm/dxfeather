package dxfeather

import (
	"fmt"
	"strconv"
)

const (
	Unitless    = 0
	UnitInches  = 1
	UnitFeet        = 2
	UnitMiles       = 3
	UnitMillimeters = 4
	UnitCentimeters = 5
	UnitMeters      = 6
	UnitKilometers  = 7
	UnitMicroinches = 8
	UnitMils        = 9
	UnitYards       = 10
	UnitAngstroms   = 11
	UnitNanometers = 12
	UnitMicrons = 13
	UnitDecimeters = 14
	UnitDecameters = 15
	UnitHectometers = 16
	UnitGigameters = 17
	UnitAstronomicalUnits = 18
	UnitLightYears = 19
	UnitParsecs = 20
)

type Header struct {
	Units int
}

func (h Header) toDXFString() string {
	var res string

	res = formatString("0","SECTION")
	res += formatString("2", "HEADER")
	res += formatVariable("$INSUNITS", "70", strconv.Itoa(h.Units))
	res += formatString("0","ENDSEC")

	return res
}

func formatVariable(variable string, groupCode string, value string) string {
	return fmt.Sprintf("%s\n%s\n%s\n", variable, groupCode, value)
}
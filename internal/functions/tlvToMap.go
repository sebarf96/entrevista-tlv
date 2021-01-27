package fun

import (
	"errors"
	"regexp"
	"strconv"
)

var (
	regexAlphaNumeric = `^[A-Z0-9]*$`
	regexNumeric      = `^[0-9]*$`
)

//TlvToMAP campos tlv en formato []Byte --> Mapa con los campos encontrados en el TLV
func TlvToMAP(tlv []byte) (map[string]string, error) {

	mapFinal := make(map[string]string)
	tlvString := string(tlv)

	if len(tlvString) == 0 {
		return nil, errors.New(ErrorEmptyTLV)
	}
	for len(tlvString) > 1 {
		large, fieldNumber, value, err := getLargeAndFieldNumberAndValue(tlvString)
		if err != nil {
			return nil, err
		}
		mapFinal[fieldNumber] = value
		tlvString = tlvString[large+5:]
	}

	return mapFinal, nil
}

func getLargeAndFieldNumberAndValue(tlv string) (int, string, string, error) {

	if len(tlv) < 6 {
		return 0, "", "", errors.New(ErrorInvalidTLV)
	}

	large, _ := strconv.Atoi(tlv[0:2])
	typeField := tlv[2:3]
	fieldNumber := tlv[3:5]
	value := tlv[5 : large+5]
	err := validateAlphaNumeric(typeField, value)
	if err != nil {
		return 0, "", "", err
	}
	return large, fieldNumber, value, nil
}

func validateAlphaNumeric(typeField string, value string) error {

	if typeField == "A" {
		re := regexp.MustCompile(regexAlphaNumeric)
		if re.MatchString(value) {
			return nil
		}
		return errors.New(ErrorInvalidAlphaNumericValue)
	}
	if typeField == "N" {
		re := regexp.MustCompile(regexNumeric)
		if re.MatchString(value) {
			return nil
		}
		return errors.New(ErrorInvalidNumericValue)

	}

	return errors.New(ErrorInvalidType)
}

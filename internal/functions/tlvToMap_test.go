package fun

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	tlvOK  = `11A05AB398765UJ102N2300`
	tlvNOK = `11Z05AB398765UJ102N2300`
)

func TestFunctionTlvToMapString(t *testing.T) {

	t.Run("Debe traer los campos tlv en un array de string", func(t *testing.T) {

		res, err := TlvToMAP([]byte(tlvOK))

		assert.Equal(t, "AB398765UJ1", res["05"])
		assert.Equal(t, "00", res["23"])
		assert.Nil(t, err)
	})

	t.Run("Debe fallar por un mal typo de campo", func(t *testing.T) {

		res, err := TlvToMAP([]byte(tlvNOK))

		assert.Error(t, err, "Error por falla de tlv")
		assert.Nil(t, res)
	})
	t.Run("Debe fallar por un tlv invalido de entrada", func(t *testing.T) {

		res, err := TlvToMAP([]byte("12345"))
		assert.Error(t, err, "Error por falla de tlv")
		assert.Nil(t, res)
	})

	t.Run("Debe fallar por un tlv vacio de entrada", func(t *testing.T) {

		res, err := TlvToMAP([]byte(""))
		assert.Error(t, err, "Error por falla de tlv")
		assert.Nil(t, res)
	})

	t.Run("Debe aprobar tipo alfa numerico", func(t *testing.T) {

		err := validateAlphaNumeric("A", "A345AS34456")

		assert.Nil(t, err)
	})

	t.Run("Debe fallar tipo alfa numerico", func(t *testing.T) {

		err := validateAlphaNumeric("A", "A345AS3445@@6")

		assert.Error(t, err)
	})

	t.Run("Debe aprobar tipo numerico", func(t *testing.T) {

		err := validateAlphaNumeric("N", "123456")

		assert.Nil(t, err)
	})

	t.Run("Debe fallar tipo numerico", func(t *testing.T) {

		err := validateAlphaNumeric("N", "A345AS3445@@6")

		assert.Error(t, err)
	})

	t.Run("Debe fallar tipo invalido", func(t *testing.T) {

		err := validateAlphaNumeric("Z", "A345AS3445@@6")

		assert.Error(t, err)
	})
}

package usecase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetData(t *testing.T) {

	var params = []string{"1,2,3", "4", "xx"}
	var expectedCode = []int{200, 404, 400}
	var expectedLenRes = []int{3, 0, 0}
	var expectedErr = []error{nil, errors.New("Resource with ID 4 not exist"), errors.New("Invalid or Empty ID xx")}
	for i, x := range params {
		res, code, err := GetData(x)

		assert.Equal(t, code, expectedCode[i])
		assert.Equal(t, len(res), expectedLenRes[i])
		assert.Equal(t, err, expectedErr[i])
	}
}

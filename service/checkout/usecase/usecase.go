package usecase

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/lonysutrisno/tokenomy/pkg"
)

func GetData(args string) (res []pkg.Data, code int, err error) {
	var data []pkg.Data
	err = json.Unmarshal([]byte(pkg.DummyData), &data)
	if err != nil {
		return res, 500, err
	}

	cleaned := strings.Replace(args, ",", "", -1)
	for _, x := range cleaned {
		id, err := strconv.Atoi(string(x))
		if err != nil {
			return res, 400, errors.New("Invalid or Empty ID " + args)
		}
		for _, el := range data {
			if id == el.ID {
				res = append(res, el)
			}
		}
	}

	if len(res) <= 0 {
		return res, 404, errors.New("Resource with ID " + args + " not exist")
	}
	return res, 200, nil
}

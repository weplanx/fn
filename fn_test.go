package fn_test

import (
	"context"
	"github.com/caarlos0/env/v9"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"github.com/weplanx/fn"
	"github.com/weplanx/fn/common"
	"os"
	"testing"
)

var (
	values common.Values
	x      *fn.Fn
)

func TestMain(m *testing.M) {
	var err error
	if err = env.Parse(&values); err != nil {
		panic(err)
	}
	if x, err = fn.New(
		fn.SetCos(values.Cos.Url, values.Cos.SecretId, values.Cos.SecretKey),
	); err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func TestFn_TencentCosExcel(t *testing.T) {
	data := [][]interface{}{
		{"Name", "CCType", "CCNumber", "Century", "Currency", "Date", "Email", "URL"},
	}
	for n := 0; n < 10; n++ {
		data = append(data, []interface{}{
			faker.Name(), faker.CCType(), faker.CCNumber(), faker.Century(), faker.Currency(), faker.Date(), faker.Email(), faker.URL(),
		})
	}
	ctx := context.TODO()
	err := x.TencentCosExcel(ctx, "test", map[string][][]interface{}{
		"Sheet1": data,
	})
	assert.NoError(t, err)
}

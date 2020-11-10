package controller

import (
	"context"
	pb "funcext/router"
	"testing"
)

func TestController_SimpleExcel(t *testing.T) {
	ctx := context.Background()
	result, err := client.SimpleExcel(ctx, &pb.Excel{
		Sheets: []*pb.Sheet{
			{
				Name: "Sheet1",
				Rows: []*pb.Row{
					{Axis: "A1", Value: "ID"},
					{Axis: "B1", Value: "Staff"},
					{Axis: "C1", Value: "Username"},
					{Axis: "D1", Value: "CreateTime"},
					{Axis: "E1", Value: "UpdateTime"},
				},
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(result.Url)
	}
}

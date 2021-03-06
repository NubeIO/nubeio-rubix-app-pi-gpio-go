package outputs

import (
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/numbers"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/types"
	"github.com/gin-gonic/gin"
	"time"
)

type bulkWrite struct {
	UO1 Outputs `json:"UO1"`
	UO2 Outputs `json:"UO2"`
	UO3 Outputs `json:"UO3"`
	UO4 Outputs `json:"UO4"`
	UO5 Outputs `json:"UO5"`
	UO6 Outputs `json:"UO6"`
	DO1 Outputs `json:"DO1"`
	DO2 Outputs `json:"DO2"`
}

type BulkWrite struct {
	IONum string  `json:"IONum"`
	Value float64 `json:"value"`
}

func (inst *Outputs) BulkWrite(ctx *gin.Context) {
	body, err := getBodyBulk(ctx)
	if err != nil {
		reposeHandler(false, err, ctx)
		return
	}
	for _, io := range body {
		writeValue := types.ToFloat64(io.Value)
		inst.Value = numbers.Scale(writeValue, 0, 100, 0, 1)
		inst.valueOriginal = writeValue
		inst.IONum = io.IONum
		time.Sleep(300 * time.Millisecond)
		write, err := inst.write()
		if err != nil {
			reposeHandler(write, err, ctx)
			return
		}
	}
	reposeHandler(true, nil, ctx)
	return
}

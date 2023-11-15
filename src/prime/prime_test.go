package prime

import (
	"context"
	"fmt"
	"github.com/bnb-chain/tss-lib/common"
	"testing"
	"time"
)

func TestGeneratePreParams2(t *testing.T) {
	for i := 0; i < 10; i++ {
		now := time.Now()

		sgps, err := common.GetRandomSafePrimesConcurrent(context.Background(), 1024, 1, 8)
		if err != nil {
			t.Fail()
		}

		fmt.Println(len(sgps))
		fmt.Println(sgps[0].SafePrime().String())

		fmt.Println(time.Now().Sub(now))
	}

}

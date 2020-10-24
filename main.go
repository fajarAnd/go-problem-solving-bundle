package main

import (
	"errors"
	"fmt"
	"math"
)

const (
	CAKE  int = 20
	APPLE int = 25
)

func pack(item int, count int) int8 {
	var mod int = item % count
	return int8(mod)
}

func bundling(apple int, cake int, count int) int8 {
	var addedItem int8 = pack(apple, count) + pack(cake, count)
	return addedItem
}

func itemInBox(item int, count int) int8 {
	divideItem := math.Floor(float64(item / count))
	return int8(divideItem)
}

type packDetail struct {
	box        int8
	itemLeft   int8
	appleInBox int8
	cakeInBox  int8
}

func packArrange(apple int, cake int, count int) packDetail {
	pd := packDetail{
		box:        int8(count),
		itemLeft:   bundling(apple, cake, count),
		appleInBox: itemInBox(apple, count),
		cakeInBox:  itemInBox(cake, count),
	}
	return pd
}

func getBestPack(packDetails []packDetail) (packDetail, error) {
	if len(packDetails) == 0 {
		return packDetail{}, errors.New("Slice Cannot empty")
	}

	min := packDetails[0]
	for _, pd := range packDetails {
		if pd.itemLeft < min.itemLeft {
			min = pd
		}
	}

	return min, nil
}

func main() {
	maxItem := math.Max(float64(CAKE), float64(APPLE))
	packLines := []packDetail{}

	for i := 2; i <= int(maxItem); i++ {
		packLines = append(packLines, packArrange(APPLE, CAKE, i))
	}

	p, e := getBestPack(packLines)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(p)
}

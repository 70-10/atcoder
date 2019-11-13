package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {
	tests := []struct {
		cities []City
		expect string
	}{
		{
			cities: []City{
				City{Name: "unagi", Population: 20},
				City{Name: "usagi", Population: 13},
				City{Name: "snuke", Population: 42},
				City{Name: "smeke", Population: 7},
			},
			expect: "snuke",
		},
		{
			cities: []City{
				City{Name: "a", Population: 10},
				City{Name: "b", Population: 20},
				City{Name: "c", Population: 30},
				City{Name: "d", Population: 40},
				City{Name: "e", Population: 100},
			},
			expect: "atcoder",
		},
		{
			cities: []City{
				City{Name: "a", Population: 10},
				City{Name: "b", Population: 20},
				City{Name: "c", Population: 30},
				City{Name: "d", Population: 40},
				City{Name: "e", Population: 100},
				City{Name: "yasuzuka", Population: 3340},
				City{Name: "uragawara", Population: 4032},
				City{Name: "oshima", Population: 2249},
				City{Name: "maki", Population: 2614},
				City{Name: "kakizaki", Population: 11484},
				City{Name: "ogata", Population: 10401},
				City{Name: "kubiki", Population: 9746},
				City{Name: "yoshikawa", Population: 5142},
				City{Name: "joetsu", Population: 100000},
				City{Name: "nakago", Population: 4733},
				City{Name: "itakura", Population: 7517},
				City{Name: "kiyosato", Population: 3152},
				City{Name: "sanwa", Population: 6190},
				City{Name: "nadachi", Population: 3169},
			},
			expect: "joetsu",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, logic(test.cities))
	}
}

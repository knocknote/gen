// Generated by: setup
// TypeWriter: gen
// Directive: +test on main.Other

package main

import (
	"errors"
	"sort"
)

// Others is a slice of type Other, for use with gen methods below. Use this type where you would use []Other. (This is required because slices cannot be method receivers.)
type Others []Other

// Max returns the maximum value of Others. In the case of multiple items being equally maximal, the first such element is returned. Returns error if no elements. See: http://clipperhouse.github.io/gen/#Max
func (rcv Others) Max() (result Other, err error) {
	l := len(rcv)
	if l == 0 {
		err = errors.New("cannot determine the Max of an empty slice")
		return
	}
	result = rcv[0]
	for _, v := range rcv {
		if v > result {
			result = v
		}
	}
	return
}

// Min returns the minimum value of Others. In the case of multiple items being equally minimal, the first such element is returned. Returns error if no elements. See: http://clipperhouse.github.io/gen/#Min
func (rcv Others) Min() (result Other, err error) {
	l := len(rcv)
	if l == 0 {
		err = errors.New("cannot determine the Min of an empty slice")
		return
	}
	result = rcv[0]
	for _, v := range rcv {
		if v < result {
			result = v
		}
	}
	return
}

// Sort returns a new ordered Others slice. See: http://clipperhouse.github.io/gen/#Sort
func (rcv Others) Sort() Others {
	result := make(Others, len(rcv))
	copy(result, rcv)
	sort.Sort(result)
	return result
}

// IsSorted reports whether Others is sorted. See: http://clipperhouse.github.io/gen/#Sort
func (rcv Others) IsSorted() bool {
	return sort.IsSorted(rcv)
}

// SortDesc returns a new reverse-ordered Others slice. See: http://clipperhouse.github.io/gen/#Sort
func (rcv Others) SortDesc() Others {
	result := make(Others, len(rcv))
	copy(result, rcv)
	sort.Sort(sort.Reverse(result))
	return result
}

// IsSortedDesc reports whether Others is reverse-sorted. See: http://clipperhouse.github.io/gen/#Sort
func (rcv Others) IsSortedDesc() bool {
	return sort.IsSorted(sort.Reverse(rcv))
}

func (rcv Others) Len() int {
	return len(rcv)
}
func (rcv Others) Less(i, j int) bool {
	return rcv[i] < rcv[j]
}
func (rcv Others) Swap(i, j int) {
	rcv[i], rcv[j] = rcv[j], rcv[i]
}

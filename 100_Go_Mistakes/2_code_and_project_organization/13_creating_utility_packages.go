// Problem here is util is meaningless. It could be called common, shared, or base
package util

func NewStringSet (...string) map[string]struct{} {
	// ...
}

func SortStringSet(map[string]struct{}) []string {
	// ...
}

package client

import "util"

set: = util.NewStringSet("c", "a", "b")
fmt.Println(uitl.SortStringSet(set))
// better to call the package stringset and rename the functions to New, Sort
// nano packages shouldn't be the norm, but they do have their places

set := stringSet.New("c", "a", "b")
fmt.Println(stringSet.Sort(set))

// or make methods
package stringset

type Set map[string]struct{}
func New(...string) Set { ... }
func (s Set) Sort() []string { ... }

package clinet

import "stringset"

set := stringset.New("c", "a", "b")
fmt.Println(set.Sort())

// also big junk drawers will become a depdency nightmare.

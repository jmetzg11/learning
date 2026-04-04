// every export should be document with a comment above that is a complete sentence starting with the name of the export
// this import for hover support, discoverability, or publishing your code

// Customer is a customer representation.
type Customer struct{}

// ID returns the customer identified.
func (c Customer) ID() string { return "" }


// functions - write what it intends to do, not how it does it. How is for comments in the function

// Deprecated elements

// ComputePath returns the fastest path between two points.
// Deprecated: This function uses a deprecated way to compute the fastest path.
// Use ComputeFastestPath instead.
func ComputePath() {}

// most IDEs hand deprecated comments


// Also comment the package name

// Package math provides basic constants and mathematical functions.
//
// blah blah
// blah blah
package math



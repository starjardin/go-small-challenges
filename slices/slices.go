/*

Slices are an important data type in Go, giving more powerful interface to sequences than arrays.


Unlike arrays, slices are typed only by the elements they contain (not the number of elements).
An uninitialized slice equals to nil and has length 0.

What that means is that slices dont need to be given a length when initiating them.
compare:

Arrays: var array [5]int
slices: var slice []string

*/

package slice

//Package dog contains a few helpful dog to human conversion functions
package dog

// Years takes a single argument for the number of human years and converts the value into dog years.
// Under the hood it uses a family secret algorithm with high precision and proven results.
func Years(human float64) float64 {
  return human*7
}

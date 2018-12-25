/*
Package space has a single function that converts the given seconds in the age of the given planet.
*/
package space

// Converts the given seconds in the age of the given planet
func Age(seconds float64, planet string) float64 {
	// A map to store the proportions between the planets age and the planet Earth
	proportions := make(map[string]float64)
	proportions["Earth"] = 1
	proportions["Mercury"] = 0.2408467
	proportions["Venus"] = 0.61519726
	proportions["Mars"] = 1.8808158
	proportions["Jupiter"] = 11.862615
	proportions["Saturn"] = 29.447498
	proportions["Uranus"] = 84.016846
	proportions["Neptune"] = 164.79132

	// Calculates the age based on the proportions
	// 31557600 is equivalent as 365.25 days on earth (one year)
	return seconds / 31557600 / proportions[planet]
}

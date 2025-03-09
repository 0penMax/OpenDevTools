package lorem

import (
	"strings"
)

const lorem = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin facilisis mi sapien, vitae accumsan libero malesuada in. Suspendisse sodales finibus sagittis. Proin et augue vitae dui scelerisque imperdiet. Suspendisse et pulvinar libero. Vestibulum id porttitor augue. Vivamus lobortis lacus et libero ultricies accumsan. Donec non feugiat enim, nec tempus nunc. Mauris rutrum, diam euismod elementum ultricies, purus tellus faucibus augue, sit amet tristique diam purus eu arcu. Integer elementum urna non justo fringilla fermentum. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Quisque sollicitudin elit in metus imperdiet, et gravida tortor hendrerit. In volutpat tellus quis sapien rutrum, sit amet cursus augue ultricies. Morbi tincidunt arcu id commodo mollis. Aliquam laoreet purus sed justo pulvinar, quis porta risus lobortis. In commodo leo id porta mattis."

// Generate generates a lorem ipsum string with the given number of words.
func Generate(wordCount int) string {
	// A base set of lorem ipsum words.
	loremWords := strings.Fields(lorem)

	words := make([]string, 0, wordCount)
	// Cycle through loremWords until we reach the desired count.
	for i := 0; i < wordCount; i++ {
		words = append(words, loremWords[i%len(loremWords)])
	}
	return strings.Join(words, " ")
}

package word

import (
	"math/rand"
	"testing"
	"time"
)

// randomPalindrome returns a palindrome whose length and contents are derived
// from the psuedo-random number generator rng.
func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		runes[i] = r
		runes[n-1-i] = r
	}
	nl := []rune("’()[]{}<>:,‒–—―…!" +
		".«»-‐?‘’“”;/⁄␠·&@*\\•^¤¢$€£¥₩₪†‡°¡¿¬#№%‰‱¶′§~¨_|¦⁂☞∴‽※ ")
	runes = append(runes, nl[rand.Intn(len(nl))])
	return string(runes)
}

func TestRandomPalindrome(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}

	}
}

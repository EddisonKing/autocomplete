package tests

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/EddisonKing/autocomplete"
	"github.com/stretchr/testify/require"
)

const charset string = "abcdefghijklmnopqrstuvwxyz"

func randomString(l uint) string {
	s := &strings.Builder{}
	for range l {
		s.WriteRune(rune(charset[rand.Intn(len(charset))]))
	}
	return s.String()
}

func randomStringsRandomLength(count, min, max uint) []string {
	result := make([]string, 0)

	for range count {
		size := uint(rand.Intn(int(max-min))) + min
		result = append(result, randomString(size))
	}

	return result
}

func TestNewAutoCompleteNotNil(t *testing.T) {
	ac := autocomplete.New()
	require.NotNil(t, ac)
}

func TestAutoCompleteLoadAndCompleteExactMatch(t *testing.T) {
	ac := autocomplete.New()

	test := "hello"
	ac.Load(test)

	result := ac.Complete("hello")
	require.Equal(t, 1, len(result))
	require.Equal(t, test, result[0])
}

func TestAutoCompleteLoadAndCompletePartialMatch(t *testing.T) {
	ac := autocomplete.New()

	ac.Load([]string{"hello", "healthy", "heap"}...)

	result := ac.Complete("hel")
	require.Equal(t, 1, len(result))

	result = ac.Complete("he")
	require.Equal(t, 3, len(result))

	result = ac.Complete("hea")
	require.Equal(t, 2, len(result))
}

var (
	smallSampleWords []string = randomStringsRandomLength(1000, 3, 12)
	hugeSampleWords  []string = randomStringsRandomLength(1_000_000, 3, 12)
)

func TestAutoCompleteCompleteAll(t *testing.T) {
	ac := autocomplete.New()

	ac.Load(smallSampleWords...)

	result := ac.Complete("")
	require.Equal(t, ac.Count(), len(result))
}

func TestAutoCompleteCompleteAtLeastOne(t *testing.T) {
	ac := autocomplete.New()

	ac.Load(smallSampleWords...)

	result := ac.Complete(smallSampleWords[0])
	require.GreaterOrEqual(t, len(result), 1)
}

// These are somewhat benchmarking tests but not really.
// The intention here is to guage the difference between a Load and Complete once the trie gets arbitrarily "large"
func TestAutoCompleteCompleteAllUnderHugeLoad(t *testing.T) {
	ac := autocomplete.New()

	start := time.Now()
	ac.Load(hugeSampleWords...)
	fmt.Printf("Loaded 1,000,000 entries in %dms\n", time.Since(start).Milliseconds())

	start = time.Now()
	result := ac.Complete("")
	require.Equal(t, ac.Count(), len(result))

	fmt.Printf("Counted all 1,000,000 entries in %dms\n", time.Since(start).Milliseconds())
}

func TestAutoCompleteCompleteAtLeastOneUnderHugeLoad(t *testing.T) {
	ac := autocomplete.New()

	start := time.Now()
	ac.Load(hugeSampleWords...)
	fmt.Printf("Loaded 1,000,000 entries in %dms\n", time.Since(start).Milliseconds())

	var firstRune rune
	for _, r := range hugeSampleWords[0] {
		firstRune = r
		break
	}

	start = time.Now()
	result := ac.Complete(string(firstRune))
	require.GreaterOrEqual(t, len(result), 1)

	fmt.Printf("Found %d results in 1,000,000 entries starting with '%s' in %dms\n", len(result), string(firstRune), time.Since(start).Milliseconds())
}

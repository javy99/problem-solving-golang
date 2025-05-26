package groupanagrams

import (
	"reflect"
	"sort"
	"strings"
	"testing"
)

func sortGroups(groups [][]string) [][]string {
	for _, group := range groups {
		sort.Strings(group)
	}
	sort.Slice(groups, func(i, j int) bool {
		return strings.Join(groups[i], "") < strings.Join(groups[j], "")
	})
	return groups
}

func TestGroupAnagrams(t *testing.T) {
	cases := []struct {
		name     string
		fn       func([]string) [][]string
		strs     []string
		expected [][]string
	}{
		// GroupAnagramsSort
		{"Sort - Basic", GroupAnagramsSort, []string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}},
		{"Sort - Empty", GroupAnagramsSort, []string{}, [][]string{}},
		{"Sort - Single Word", GroupAnagramsSort, []string{"abc"}, [][]string{{"abc"}}},
		{"Sort - All Same", GroupAnagramsSort, []string{"a", "a", "a"}, [][]string{{"a", "a", "a"}}},
		{"Sort - No Anagrams", GroupAnagramsSort, []string{"a", "b", "c"}, [][]string{{"a"}, {"b"}, {"c"}}},

		// GroupAnagramsCount
		{"Count - Basic", GroupAnagramsCount, []string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}},
		{"Count - Empty", GroupAnagramsCount, []string{}, [][]string{}},
		{"Count - Single Word", GroupAnagramsCount, []string{"abc"}, [][]string{{"abc"}}},
		{"Count - All Same", GroupAnagramsCount, []string{"a", "a", "a"}, [][]string{{"a", "a", "a"}}},
		{"Count - No Anagrams", GroupAnagramsCount, []string{"a", "b", "c"}, [][]string{{"a"}, {"b"}, {"c"}}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := sortGroups(c.fn(c.strs))
			want := sortGroups(c.expected)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("%s failed: expected %v, got %v", c.name, want, got)
			}
		})
	}
}

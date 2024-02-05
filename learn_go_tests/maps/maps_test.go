package maps

import "testing"

func TestSearch(t *testing.T) {

	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test") // here dictionary is a structure
		want := "this is just a test"

		assertString(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown") // here dictionary is a structure
		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}

	// creating a new method for the type Dictionary to add a new entry to the map
	dictionary.Add("test", "this is just a test")

	want := "this is just a test"

	got, err := dictionary.Search("test")

	if err != nil {
		t.Errorf("Adding was not successful")
	}

	assertString(t, got, want)
}

func assertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

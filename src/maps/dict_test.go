package maps

import "testing"

const word = "test"
const definition = "this is just a test"
const newDefinition = "new definition"

func TestDict(t *testing.T) {
	dict := Dict{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, err := dict.Search("test")
		want := "this is just a test"

		assertNoError(t, err)
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dict.Search("unknown")

		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	dict := Dict{}

	t.Run("New word", func(t *testing.T) {
		err := dict.Add(word, definition)
		assertDefinition(t, dict, word, definition)
		assertError(t, err, nil)
	})

	t.Run("Existing word", func(t *testing.T) {
		_ = dict.Add(word, definition)
		err := dict.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dict, word, definition)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestUpdate(t *testing.T) {

	t.Run("Update existing", func(t *testing.T) {
		dict := Dict{}
		_ = dict.Add(word, definition)
		err := dict.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dict, word, newDefinition)
	})

	t.Run("Update not existing", func(t *testing.T) {
		dict := Dict{}
		err := dict.Update(word, newDefinition)

		assertError(t, err, ErrWordDoesNotExists)
	})
}

func TestDelete(t *testing.T) {

	t.Run("Test deletion", func(t *testing.T) {
		dict := Dict{word: definition}
		dict.Delete(word)
		_, err := dict.Search(word)

		if err != ErrNotFound {
			t.Errorf("Expected %q to be deleted", word)
		}
	})
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Fatal("Did not expect any error but got one")
	}
}

func assertDefinition(t testing.TB, dict Dict, word, definition string) {
	t.Helper()

	got, err := dict.Search(word)
	assertNoError(t, err)
	assertStrings(t, got, definition)
}

package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("辞書の検索", func(t *testing.T) {
		got, err := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
		assertNoError(t, err)
	})

	t.Run("存在しないワードの辞書検索", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("ワードの新規登録", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("登録済みのワードの登録は不可能", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{
			word: definition,
		}

		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("正常な更新", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		newDefinition := "new definition"
		dictionary := Dictionary{
			word: definition,
		}

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})
	t.Run("ワードが存在しない場合はエラーとなる", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExists)
	})
}

func TestDelete(t *testing.T) {
	t.Run("正常な削除", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{
			word: definition,
		}

		err := dictionary.Delete(word)

		assertError(t, err, nil)
		assertNotFound(t, dictionary, word)
	})
}

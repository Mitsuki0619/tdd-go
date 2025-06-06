package dictionary

type DictionaryErr string

const (
	ErrNotFound          = DictionaryErr("could not find the word you were looking for")
	ErrWordExists        = DictionaryErr("the word already exists")
	ErrWordDoesNotExists = DictionaryErr("the word not exists")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

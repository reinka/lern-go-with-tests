package maps

const (
	ErrNotFound          = DictErr("Could not find the word you were looking for")
	ErrWordExists        = DictErr("Word already exists")
	ErrWordDoesNotExists = DictErr("Word does not exist")
)

type Dict map[string]string
type DictErr string

func (d DictErr) Error() string {
	return string(d)
}

func (d Dict) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dict) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dict) Update(word, newDefinition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExists
	case nil:
		d[word] = newDefinition
	default:
		return err
	}

	return nil
}

func (d Dict) Delete(word string) {
	delete(d, word)
}

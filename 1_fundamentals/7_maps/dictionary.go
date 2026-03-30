package maps

type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("word doesn't exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (dic Dictionary) Search(word string) (string, error) {
	definition, ok := dic[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (dic Dictionary) Add(word, definition string) error {
	_, err := dic.Search(word)

	switch err {
	case ErrNotFound:
		dic[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (dic Dictionary) Update(word, definition string) error {
	_, err := dic.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		dic[word] = definition
	default:
		return err
	}

	return nil
}

func (dic Dictionary) Delete(word string) error {
	_, err := dic.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		delete(dic, word)
	default:
		return err
	}
	
	return nil
}

package dictionary

const (
	//ErrNotFound ...
	ErrNotFound = Err("could not find the word you were looking for")

	//ErrWordExists ...
	ErrWordExists = Err("Word already exisists")

	//ErrWordDoesNotExist ..
	ErrWordDoesNotExist = Err("Word does not exist")
)

//Err errors for dictionary
type Err string

func (e Err) Error() string {
	return string(e)
}

//Search ...
func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}

//Dictionary ...
type Dictionary map[string]string

//Search ...
func (d Dictionary) Search(word string) (string, error) {
	defitinion, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return defitinion, nil
}

//Add lets you populate array with word and its definition
func (d Dictionary) Add(word, definition string) error {
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

//Update updates
func (d Dictionary) Update(word, definition string) error {

	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

//Delete removes stuff
func (d Dictionary) Delete(word string) {
	delete(d, word)
}

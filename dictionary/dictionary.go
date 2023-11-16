package dictionary

type Word struct {
	Definition string
}

func (w Word) String() string {
	return w.Definition
}

type Dictionary struct {
	words map[string]Word
}

// Dans la definition, je retourne un pointeur
func New() *Dictionary {

	// Je retourne l'adresse du nouveau dictionnaire avec &
	return &Dictionary{
		words: make(map[string]Word),
	}
}

func (d *Dictionary) Add(word string, definition string) {

	newWord := Word{
		Definition: definition,
	}

	d.words[word] = newWord
}

func (d *Dictionary) Get(word string) (*Word, bool) {
	entry, exists := d.words[word]
	if exists {
		return &entry, true
	}
	return nil, false
}

func (d *Dictionary) Remove(word string) {
	delete(d.words, word)
}

func (d *Dictionary) List() ([]string, map[string]Word) {
	wordsList := make([]string, 0)
	for word := range d.words {
		wordsList = append(wordsList, word)
	}
	return wordsList, d.words
}

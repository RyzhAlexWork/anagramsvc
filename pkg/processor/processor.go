package processor

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/go-kit/kit/log"

	"anagramsvc/pkg/api/rest"
	"anagramsvc/pkg/api/v1"
)

type generateCache interface {
	Get() (data *[]string)
	Set(data []string) (message string)
}

// Processor ...
type Processor interface {
	GetAnagrams(ctx context.Context, req rest.Request) (response v1.GetAnagramsResponse, err error)
	LoadWords(ctx context.Context, req rest.Request) (response v1.LoadWordsResponse, err error)
}

type processor struct {
	logger        log.Logger
	generateCache generateCache
}

// Метод позволяет получить все анаграммы для заданного слова, если таковые имеются в словаре
func (p *processor) GetAnagrams(ctx context.Context, req rest.Request) (response v1.GetAnagramsResponse, err error) {
	var anagrams []string

	// Получаем слова из словаря
	loadWords := p.generateCache.Get()
	// Проверяем что слова загружены
	if len(*loadWords) == 0 {
		err = fmt.Errorf(v1.MissingWordListError)
		return
	}
	// Переводим слово для поиска в нижний регистр
	*req.SearchWord() = strings.ToLower(*req.SearchWord())
	// Сравниваем слово для поиска с каждым словом из словаря
	for _, wordFromList := range *loadWords {
		if isAnagram(*req.SearchWord(), wordFromList) {
			anagrams = append(anagrams, wordFromList)
		}
	}
	response.Data = &anagrams
	return
}

// Метод позволяет загрузить слова в словарь
func (p *processor) LoadWords(ctx context.Context, req rest.Request) (response v1.LoadWordsResponse, err error) {
	response.Data = p.generateCache.Set(*req.LoadWords())
	return
}

// Функция для сортировки строки в алфавитном порядке
func sortString(str string) (result string) {
	strArray := strings.Split(str, "")
	sort.Strings(strArray)
	return strings.Join(strArray, "")
}

// Функция проверки являются ли два слова анограммами
func isAnagram(searchWord, wordFromList string) (result bool) {
	if sortString(searchWord) == sortString(strings.ToLower(wordFromList)) {
		result = true
	}
	return
}

// NewProcessor ...
func NewProcessor(
	logger log.Logger,
	generateCache generateCache,
) Processor {
	return &processor{
		logger:        logger,
		generateCache: generateCache,
	}
}

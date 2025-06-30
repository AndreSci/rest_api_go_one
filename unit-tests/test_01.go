package unittest_test

import (
	"fmt"
	"time"

	"github.com/AndreSci/rest_api_go_one/internal/models"
	"github.com/AndreSci/rest_api_go_one/internal/service"
)

type ResultTest struct {
	test1 string
	test2 string
	test3 string
}

func RunTests() {
	time.Sleep(time.Second * 2)
	result := runAll()
	fmt.Println(result)
}

func runAll() ResultTest {

	result := ResultTest{"none", "none", "none"}

	client, err := service.NewClient(time.Second * 10)
	if err != nil {
		result.test1 = fmt.Sprintf("Error creat client: %v", err)
		return result
	}
	result.test1 = "Done"

	errDelAll := client.DeleteAll()
	if errDelAll != nil {
		result.test2 = fmt.Sprintf("Error Delete All books: %v", errDelAll)
		return result
	}
	result.test2 = "Done"

	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("book_%d", i+1)
		author := "test_author"
		book := models.NewBook{Name: name, Author: author}

		errAddBook := client.AddBook(&book)
		if errAddBook != nil {
			result.test3 = fmt.Sprintf("Error Add books (index: %d): %v", i, errAddBook)
			return result
		}
	}
	result.test3 = "Done"

	return result
}

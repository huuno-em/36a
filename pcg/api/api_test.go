package api_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"GoNews/pcg/api"
	"GoNews/pcg/database"
	"GoNews/pcg/typeStruct"

	"github.com/stretchr/testify/assert"
)

func TestAPI_Posts(t *testing.T) {
	// Инициализация реальной базы данных
	database.InitDB()
	// Завершение работы с базой данных после теста
	defer database.DB.Close()

	// Создание API
	api := api.NewAPI(database.DB)

	// Добавляем 5 тестовых постов в базу данных
	for i := 1; i <= 5; i++ {
		testPost := typeStruct.Post{
			Title:   fmt.Sprintf("Test Title %d", i),
			Content: "Test Content",
			PubTime: 1692644239,
			Link:    "http://example.com/test",
		}
		err := database.SaveToDB(testPost)
		assert.NoError(t, err, "Failed to save post to DB")
	}

	// Создание HTTP запроса к обработчику /news/{n}
	req, err := http.NewRequest("GET", "/news/5", nil)
	assert.NoError(t, err, "Unexpected error")

	// Создание HTTP ResponseRecorder (регистратора ответов)
	rr := httptest.NewRecorder()

	// Обработка запроса
	api.ServeHTTP(rr, req)

	// Проверка кода состояния HTTP
	assert.Equal(t, http.StatusOK, rr.Code, "Expected status code 200")

	// Парсинг ответа JSON
	var responsePosts []typeStruct.Post
	err = json.Unmarshal(rr.Body.Bytes(), &responsePosts)
	assert.NoError(t, err, "Failed to parse response JSON")

	// Получаем последние 5 постов из базы данных
	expectedPosts, err := database.GetLatestPosts(5)
	assert.NoError(t, err, "Failed to get latest posts from DB")

	// Проверяем, что количество постов и их содержимое совпадают
	assert.Equal(t, len(expectedPosts), len(responsePosts), "Number of posts doesn't match")

	for i := 0; i < len(expectedPosts); i++ {
		assert.Equal(t, expectedPosts[i].Title, responsePosts[i].Title, "Title doesn't match")
		assert.Equal(t, expectedPosts[i].Content, responsePosts[i].Content, "Content doesn't match")
		assert.Equal(t, expectedPosts[i].PubTime, responsePosts[i].PubTime, "PubTime doesn't match")
		assert.Equal(t, expectedPosts[i].Link, responsePosts[i].Link, "Link doesn't match")
	}
}

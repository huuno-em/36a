package database

import (
	"testing"

	"GoNews/pcg/typeStruct"

	"github.com/stretchr/testify/assert"
)

func TestSaveAndReadFromDB(t *testing.T) {
	InitDB()

	defer DB.Close()

	// Создаем тестовый пост
	testPost := typeStruct.Post{
		Title:   "Test Title 2",
		Content: "Test Content",
		PubTime: 1692645688,
		Link:    "http://example.com/test",
	}

	// Сохраняем тестовый пост в базу данных
	err := SaveToDB(testPost)
	if err != nil {
		t.Fatalf("Failed to save post to DB: %v", err)
	}

	// Читаем пост из базы данных по названию
	readPost, err := ReadFromDB("Test Title 2")
	if err != nil {
		t.Fatalf("Failed to read post from DB: %v", err)
	}

	// Сравниваем значения
	if readPost.Title != testPost.Title ||
		readPost.Content != testPost.Content ||
		readPost.PubTime != testPost.PubTime ||
		readPost.Link != testPost.Link {
		t.Errorf("Saved data doesn't match expected data.")
	}
}

func TestDeleteByTitle(t *testing.T) {
	InitDB()

	defer DB.Close()

	// Создаем тестовый пост
	testPost := typeStruct.Post{
		Title:   "Test Title 1",
		Content: "Test Content",
		PubTime: 1692645688,
		Link:    "http://example.com/test",
	}

	// Сохраняем тестовый пост в базу данных
	err := SaveToDB(testPost)
	assert.NoError(t, err, "Failed to save post to DB")

	// Удаляем пост по названию
	err = DeleteByTitle("Test Title 3")
	assert.NoError(t, err, "Failed to delete post by title")

	// Пытаемся прочитать пост с удаленным названием
	_, err = ReadFromDB("Test Title 3")
	assert.Error(t, err, "Expected an error when trying to read deleted post")
}

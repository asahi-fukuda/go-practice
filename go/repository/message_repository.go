package repository

import "github.com/GIT_USER_ID/GIT_REPO_ID/go/model"

type MessageRepository interface {
	Save(message *model.Message) error
	List() ([]*model.Message, error)
}

package rdb

import (
	"context"
	"encoding/json"
	"fmt"
	"frame/models"
	"github.com/go-redis/redis/v8"
)

// 获取用户缓存
func (s *Store) GetUser(userId uint64) (*models.User, error) {
	b, err := s.client.Get(context.Background(), fmt.Sprintf("user:%v", userId)).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, err
	}

	var user models.User
	return &user, json.Unmarshal(b, &user)
}

// 保存用户的缓存
func (s *Store) SaveUser(user *models.User) error {
	b, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return s.client.Set(context.Background(), fmt.Sprintf("user:%v", user.ID), b, 0).Err()
}

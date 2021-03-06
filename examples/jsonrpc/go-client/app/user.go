// Copyright 2016-2019 Yincheng Fang, Alex Stocks
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"fmt"
	"time"
)

import (
	"github.com/dubbo/go-for-apache-dubbo/config"
)

func init() {
	config.SetConsumerService(new(UserProvider))
}

type JsonRPCUser struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int64  `json:"age"`
	Time int64  `json:"time"`
	Sex  string `json:"sex"`
}

func (u JsonRPCUser) String() string {
	return fmt.Sprintf(
		"User{ID:%s, Name:%s, Age:%d, Time:%s, Sex:%s}",
		u.ID, u.Name, u.Age, time.Unix(u.Time, 0).Format("2006-01-02 15:04:05.99999"), u.Sex,
	)
}

type UserProvider struct {
	GetUsers func(req []interface{}) ([]JsonRPCUser, error)
	GetUser  func(ctx context.Context, req []interface{}, rsp *JsonRPCUser) error
	GetUser0 func(id string, name string) (JsonRPCUser, error)
	GetUser1 func(ctx context.Context, req []interface{}, rsp *JsonRPCUser) error
	GetUser2 func(ctx context.Context, req []interface{}, rsp *JsonRPCUser) error `dubbo:"getUser"`
	Echo     func(ctx context.Context, req interface{}) (interface{}, error)      // Echo represent EchoFilter will be used
}

func (u *UserProvider) Service() string {
	return "com.ikurento.user.UserProvider"
}

func (u *UserProvider) Version() string {
	return ""
}

package sqlxx

import (
	"fmt"
	"io/ioutil"

	_ "github.com/lib/pq"
	"gopkg.in/yaml.v2"
)

type DBDescriptor struct {
	// 数据库地址
	Host string `yaml:"host"`
	// 数据库名称
	Name string `yaml:"name" `
	// 用户名称
	User string `yaml:"user" `
	// 用户密码
	Password string `yaml:"password" `
	// 端口号
	Port string `yaml:"port" `
}

func (descriptor *DBDescriptor) ConfigByYaml(fileName string) (*DBDescriptor, error) {
	if fileName == "" {
		panic("yaml read failure")
	}
	// parser
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic("yaml read failure")
	}
	err = yaml.Unmarshal(data, descriptor)
	if err != nil {
		panic(err)
	}
	return descriptor, nil
}

func (descriptor *DBDescriptor) String() string {
	return fmt.Sprintf("host=%s dbname=%s user=%s password=%s port=%s sslmode=disable", descriptor.Host, descriptor.Name,
		descriptor.User, descriptor.Password, descriptor.Port)
}

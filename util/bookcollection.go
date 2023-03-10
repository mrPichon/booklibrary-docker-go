package util

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"
)

type Book struct {
	ID       int    `mapstructure:"id"`
	Name     string `mapstructure:"name"`
	Author   string `mapstructure:"author"`
	Labels   string `mapstructure:"labels"`
	Quantity int    `mapstructure:"quantity"`
}

type Collection struct {
	BookList []Book `mapstructure:"booklist"`
}

var BookCollection *Collection

func InitMockCollection() {
	BookCollection = &Collection{
		BookList: []Book{
			{
				ID:       1,
				Name:     "name",
				Author:   "isr",
				Labels:   "none",
				Quantity: 2,
			},
			{
				ID:       2,
				Name:     "hola",
				Author:   "isr",
				Labels:   "none",
				Quantity: 5,
			},
		},
	}
}

func (i Collection) MarshalBinary() (data []byte, err error) {
	bytes, err := json.Marshal(i)
	return bytes, err
}

func LoadCollection(path string) (err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("collection")
	viper.SetConfigType("json")

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error in reading config ", err)
		return
	}

	err = viper.Unmarshal(&BookCollection)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v\n", err)
	}
	return
}

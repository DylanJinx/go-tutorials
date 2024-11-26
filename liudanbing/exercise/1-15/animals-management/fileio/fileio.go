package fileio

import (
	"encoding/json"
	"errors"
	"animal-zoo/models"
	"animal-zoo/zoo"
	"os"
)

// AnimalWrapper 用于在JSON中存储动物类型和数据
type AnimalWrapper struct {
	Type string				`json:"type"`
	Data json.RawMessage	`json:"data"`
}

// SaveToFile 将Zoo的数据保存到指定文件
func SaveToFile(z *zoo.Zoo, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	data := struct {
		Animals []AnimalWrapper `json:"animals"`
		NextID  int             `json:"next_id"`
	}{
		Animals: make([]AnimalWrapper, 0, len(z.Animals)),
		NextID: z.NextID,
	}

	for _, animal := range z.Animals {
		var wrapper AnimalWrapper

		switch animal.(type) {
		case *models.Lion:
			wrapper.Type = "Lion"
		case *models.Eagle:
			wrapper.Type = "Eagle"
		case *models.Snake:
			wrapper.Type = "Snake"
		case *models.Shark:
			wrapper.Type = "Shark"
		case *models.Frog:
			wrapper.Type = "Frog"
		default:
			return errors.New("unknown animal type during saving")
		}

		b, err := json.Marshal(animal)
		if err != nil {
			return err
		}
		wrapper.Data = b
		data.Animals = append(data.Animals, wrapper)
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")	// 设置缩进
	return encoder.Encode(data)
}

// LoadFromFile 从指定文件加载Zoo的数据
func LoadFromFile(z *zoo.Zoo, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var data struct {
		Animals []AnimalWrapper `json:"animals"`
		NextID  int             `json:"next_id"`
	}

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return err
	}
}
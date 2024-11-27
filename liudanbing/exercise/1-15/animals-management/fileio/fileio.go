package fileio

import (
	"encoding/json"
	"errors"
	"animal-zoo/models"
	"animal-zoo/zoo"
	"os"
)

// AnimalWrapper 用于在JSON文件中储存动物类型和数据
type AnimalWrapper struct {
	Type string 		 `json:"type"`
	Data json.RawMessage `json:"data"`
}

// SaveToFile 将Zoo的数据保存到指定文件
func SaveToFile(z *zoo.Zoo, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// data 用于保存所有动物的数据
	data := struct {
		Animals []AnimalWrapper `json:"animals"`
		NextID int			    `json:"next_id"`
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

	// 创建一个JSON编码器，并设置缩进格式以美化输出
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
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
		NextID int			    `json:"next_id"`
	}

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return err
	}

	animals := make(map[int]models.Animal)
	for _, wrapper := range data.Animals {
		var animal models.Animal
		switch wrapper.Type {
		case "Lion":
			var lion models.Lion
			if err := json.Unmarshal(wrapper.Data, &lion); err != nil {
				return err
			}
			animal = &lion

		case "Eagle":
			var eagle models.Eagle
			if err := json.Unmarshal(wrapper.Data, &eagle); err != nil {
				return err
			}
			animal = &eagle

		case "Snake":
			var snake models.Snake
			if err := json.Unmarshal(wrapper.Data, &snake); err != nil {
				return err
			}
			animal = &snake

		case "Shark":
			var shark models.Shark
			if err := json.Unmarshal(wrapper.Data, &shark); err != nil {
				return err
			}
			animal = &shark

		case "Frog":
			var frog models.Frog
			if err := json.Unmarshal(wrapper.Data, &frog); err != nil {
				return err
			}
			animal = &frog
		
		default:
			return errors.New("unknown animal type during loading")
		}
		
		animals[animal.GetID()] = animal //使用animal.GetID()方法获取动物的ID，将动物对象存储到animals映射中，键为ID，值为动物对象
	}

	z.SetAnimals(animals, data.NextID)
	return nil
}
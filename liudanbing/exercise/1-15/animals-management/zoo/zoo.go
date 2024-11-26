package zoo

import (
	"animal-zoo/models"
	"sync"
)

// Zoo 结构体表示动物园，包含动物的map和下一个ID
type Zoo struct {
	Animals map[int]models.Animal
	NextID  int
	mu	    sync.Mutex
}

// NewZoo 创建一个新的动物园, 返回一个指向Zoo实例的指针
func NewZoo() *Zoo {
	return &Zoo{
		Animals: make(map[int]models.Animal),
		NextID:  1,
	}
}

// AddAnimal 添加一个新动物，返回分配的ID
func (z *Zoo) AddAnimal(animal models.Animal) int {
	z.mu.Lock()
	defer z.mu.Unlock()

	id := z.NextID
	animal.SetID(id)
	z.Animals[id] = animal
	z.NextID++
	return id
}

// DeleteAnimal 删除指定ID的动物，返回是否成功
func (z *Zoo) DeleteAnimal(id int) bool {
	z.mu.Lock()
	defer z.mu.Unlock()

	if _, exists := z.Animals[id]; exists {
		delete(z.Animals, id)
		return true
	}
	return false
}

func (z *Zoo) UpdateAnimal(id int, updatedAnimal models.Animal) bool {
	z.mu.Lock()
	defer z.mu.Unlock()

	if _, exists := z.Animals[id]; exists {
		updatedAnimal.SetID(id)
		z.Animals[id] = updatedAnimal
		return true
	}

	return false
}

// GetAnimal 查询指定ID的动物，返回动物和是否存在
func (z *Zoo) GetAnimal(id int) (models.Animal, bool) {
	z.mu.Lock()
	defer z.mu.Unlock()

	animal, exists := z.Animals[id]
	return animal, exists
}

// ListAnimals 列出所有动物，返回动物切片
func (z *Zoo) ListAnimals() []models.Animal {
	z.mu.Lock()
	defer z.mu.Unlock()

	animals := make([]models.Animal, 0, len(z.Animals))
	for _, animal := range z.Animals {
		animals = append(animals, animal)
	}

	return animals
}

// SetAnimals 设置Zoo的Animals和NextID，用于加载数据
func (z *Zoo) SetAnimals(animals map[int]models.Animal, nextID int) {
	z.mu.Lock()
	defer z.mu.Unlock()

	z.Animals = animals
	z.NextID = nextID
}

// GetAnimals 获取Zoo的Animals和NextID，用于保存数据
func (z *Zoo) GetAnimals() (map[int]models.Animal, int) {
	z.mu.Lock()
	defer z.mu.Unlock()

	// 创建一个副本，防止外部修改
	animalsCopy := make(map[int]models.Animal)
	for id, animals := range z.Animals {
		animalsCopy[id] = animals
	}

	return animalsCopy, z.NextID
}
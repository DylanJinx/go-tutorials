package main

import (
	"bufio"
	"fmt"
	"animal-zoo/fileio"
	"animal-zoo/models"
	"animal-zoo/zoo"
	"os"
	"strconv"
)

func main() {
	z := zoo.NewZoo()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n==== 动物园管理系统 ====")
		fmt.Println("请选择操作：")
		fmt.Println("1. 添加动物")
		fmt.Println("2. 删除动物")
		fmt.Println("3. 修改动物")
		fmt.Println("4. 查询动物")
		fmt.Println("5. 列出所有动物")
		fmt.Println("6. 保存数据")
		fmt.Println("7. 加载数据")
		fmt.Println("8. 退出系统")
		fmt.Print("输入选项（1-8）：")

		if !scanner.Scan() {
			fmt.Println("读取输入时发生错误。")
			continue
		}
		choice := scanner.Text()

		switch choice {
		case "1":
			addAnimal(z, scanner)
		case "2":
			deleteAnimal(z, scanner)
		case "3":
			updateAnimal(z, scanner)
		case "4":
			queryAnimal(z, scanner)
		case "5":
			listAnimals(z)
		case "6":
			saveData(z, scanner)
		case "7":
			loadData(z, scanner)
		case "8":
			fmt.Println("感谢使用动物园管理系统！")
			return
		default:
			fmt.Println("无效的选项，请重新输入。")
		}
	}
}

func addAnimal(z *zoo.Zoo, scanner *bufio.Scanner) {
	fmt.Println("选择动物类型：")
	fmt.Println("1. Lion")
	fmt.Println("2. Eagle")
	fmt.Println("3. Snake")
	fmt.Println("4. Shark")
	fmt.Println("5. Frog")
	fmt.Print("输入选项（1-5）：")
	if !scanner.Scan() {
		fmt.Println("读取输入时发生错误。")
		return
	}
	typeChoice := scanner.Text()

	var animal models.Animal
	var err error

	switch typeChoice {
	case "1":
		animal, err = createLion(scanner)
	case "2":
		animal, err = createEagle(scanner)
	case "3":
		animal, err = createSnake(scanner)
	case "4":
		animal, err = createShark(scanner)
	case "5":
		animal, err = createFrog(scanner)
	default:
		fmt.Println("无效的选项，请重新输入。")
		return
	}

	if err != nil {
		fmt.Printf("创建动物时发生错误：%v\n", err)
		return
	}

	id := z.AddAnimal(animal)
	fmt.Printf("动物添加成功，ID为：%d\n", id)
}

func createLion(scanner *bufio.Scanner) (models.Animal, error) {
	fmt.Print("输入狮子的名字：")
	if !scanner.Scan() {
		return nil, fmt.Errorf("读取输入时发生错误")
	}
	name := scanner.Text()

	fmt.Print("输入狮子的年龄：")
	if !scanner.Scan() {
		return nil, fmt.Errorf("读取输入时发生错误")
	}
	ageStr := scanner.Text()
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		return nil, fmt.Errorf("无效的年龄：%v", ageStr)
	}

	lion := &models.Lion{
		Name: name,
		Age:  age,
		Category: models.Mammal,
	}
	fmt.Println("lion", lion)

	return lion, nil
}

func createEagle(scanner *bufio.Scanner) (models.Animal, error) {
	fmt.Print("输入老鹰的名字：")
	if !scanner.Scan() {
		return nil, fmt.Errorf("读取输入时发生错误")
	}
	name := scanner.Text()

	fmt.Print("输入老鹰的年龄：")
	if !scanner.Scan() {
		return nil, fmt.Errorf("读取输入时发生错误")
	}
	ageStr := scanner.Text()
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		return nil, fmt.Errorf("无效的年龄：%v", ageStr)
	}

	eagle := &models.Eagle{
		Name: name,
		Age:  age,
		Category: models.Bird,
	}

	return eagle, nil
}

func createSnake(scanner *bufio.Scanner) (models.Animal, error) {
	fmt.Print("输入蛇的名字：")
	if !scanner.Scan() {
		return nil, fmt.Errorf("读取输入时发生错误")
	}
	name := scanner.Text()

	fmt.Print("输入蛇的年龄：")
	if !scanner.Scan() {
		return nil, fmt.Errorf("读取输入时发生错误")
	}
	ageStr := scanner.Text()
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		return nil, fmt.Errorf("无效的年龄：%v", ageStr)
	}

	snake := &models.Snake{
		Name: name,
		Age:  age,
		Category: models.Reptile,
	}

	return snake, nil
}

func createShark(scanner *bufio.Scanner) (models.Animal, error) {
	fmt.Print("输入鲨鱼的名字：")
	if !scanner.Scan() {
		return nil, fmt.Errorf("读取输入时发生错误")
	}
	name := scanner.Text()

	fmt.Print("输入鲨鱼的年龄：")
	if !scanner.Scan() {
		return nil, fmt.Errorf("读取输入时发生错误")
	}
	ageStr := scanner.Text()
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		return nil, fmt.Errorf("无效的年龄：%v", ageStr)
	}

	shark := &models.Shark{
		Name: name,
		Age:  age,
		Category: models.Fish,
	}

	return shark, nil
}

func createFrog(scanner *bufio.Scanner) (models.Animal, error) {
	fmt.Print("输入青蛙的名字：")
	if !scanner.Scan() {
		return nil, fmt.Errorf("读取输入时发生错误")
	}
	name := scanner.Text()

	fmt.Print("输入青蛙的年龄：")
	if !scanner.Scan() {
		return nil, fmt.Errorf("读取输入时发生错误")
	}
	ageStr := scanner.Text()
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		return nil, fmt.Errorf("无效的年龄：%v", ageStr)
	}

	frog := &models.Frog{
		Name: name,
		Age:  age,
		Category: models.Amphibian,
	}

	return frog, nil
}

func deleteAnimal(z *zoo.Zoo, scanner *bufio.Scanner) {
	fmt.Println("输入要删除的动物ID：")
	if !scanner.Scan() {
		fmt.Println("读取输入时发生错误。")
		return
	}

	idStr := scanner.Text()
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("无效的ID。")
		return
	}

	success := z.DeleteAnimal(id)
	if success {
		fmt.Println("动物删除成功。")
	} else {
		fmt.Println("未找到指定ID的动物。")
	}
}

func updateAnimal(z *zoo.Zoo, scanner *bufio.Scanner) {
	fmt.Println("输入要修改的动物ID：")
	if !scanner.Scan() {
		fmt.Println("读取输入时发生错误。")
		return
	}
	idStr := scanner.Text()
	id, err := strconv.Atoi(idStr) // convert string to int
	if err != nil {
		fmt.Println("无效的ID。")
		return
	}

	animal, exists := z.GetAnimal(id)
	if !exists {
		fmt.Println("未找到指定ID的动物。")
		return
	}

	fmt.Printf("当前动物详情：%s\n", animal.GetDetails())

	fmt.Print("输入新的名字（留空表示不修改）：")
	if !scanner.Scan() {
		fmt.Println("读取输入时发生错误。")
		return
	}
	name := scanner.Text()
	if name == "" {
		name = animal.GetName()
	}

	fmt.Print("输入新的年龄（留空表示不修改）：")
	if !scanner.Scan() {
		fmt.Println("读取输入时发生错误。")
		return
	}
	ageStr := scanner.Text()
	var age int
	if ageStr == "" {
		// 保持原有年龄
		switch a := animal.(type) {
		case *models.Lion:
			age = a.Age
		case *models.Eagle:
			age = a.Age
		case *models.Snake:
			age = a.Age
		case *models.Shark:
			age = a.Age
		case *models.Frog:
			age = a.Age
		default:
			fmt.Println("不支持的动物类型。")
			return
		}
	} else {
		age, err = strconv.Atoi(ageStr)
		if err != nil {
			fmt.Println("无效的年龄。")
			return
		}
	}

	var updatedAnimal models.Animal
	switch a := animal.(type) {
	case *models.Lion:
		updatedLion := &models.Lion{
			Name: name,
			Age:  age,
			Category: a.Category,
		}
		updatedAnimal = updatedLion
	case *models.Eagle:
		updatedEagle := &models.Eagle{
			Name: name,
			Age:  age,
			Category: a.Category,
		}
		updatedAnimal = updatedEagle
	case *models.Snake:
		updatedSnake := &models.Snake{
			Name: name,
			Age:  age,
			Category: a.Category,
		}
		updatedAnimal = updatedSnake
	case *models.Shark:
		updatedShark := &models.Shark{
			Name: name,
			Age:  age,
			Category: a.Category,
		}
		updatedAnimal = updatedShark
	case *models.Frog:
		updatedFrog := &models.Frog{
			Name: name,
			Age:  age,
			Category: a.Category,
		}
		updatedAnimal = updatedFrog
	default:
		fmt.Println("不支持的动物类型。")
		return
	}

	success := z.UpdateAnimal(id, updatedAnimal)
	if success {
		fmt.Println("动物更新成功。")
	} else {
		fmt.Println("更新失败。")
	}

}

func queryAnimal(z *zoo.Zoo, scanner *bufio.Scanner) {
	fmt.Println("输入要查询的动物ID：")
	if !scanner.Scan() {
		fmt.Println("读取输入时发生错误。")
		return
	}

	idStr := scanner.Text()
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("无效的ID。")
		return
	}

	animal, exists := z.GetAnimal(id)
	if exists {
		fmt.Println("动物信息：", animal.GetDetails())
		fmt.Println("动物叫声：", animal.Speak())
		fmt.Println("动物移动方式：", animal.Move())
	} else {
		fmt.Println("未找到指定ID的动物。")
	}
}

func listAnimals(z *zoo.Zoo) {
	animals := z.ListAnimals()
	if len(animals) == 0 {
		fmt.Println("动物园中没有动物。")
		return
	}

	fmt.Println("动物园中的动物：")
	for _, animal := range animals {
		fmt.Println(animal.GetDetails())
	}
}

func saveData(z *zoo.Zoo, scanner *bufio.Scanner) {
	fmt.Println("输入要保存的文件名(e.g. zoo.json):")
	if !scanner.Scan() {
		fmt.Println("读取输入时发生错误。")
		return
	}
	filename := scanner.Text()
	err := fileio.SaveToFile(z, filename)
	if err != nil {
		fmt.Printf("保存数据时发生错误：%v\n", err)
	} else {
		fmt.Println("数据保存成功。")
	}
}

func loadData(z *zoo.Zoo, scanner *bufio.Scanner) {
	fmt.Println("输入要加载的文件名(e.g. zoo.json):")
	if !scanner.Scan() {
		fmt.Println("读取输入时发生错误。")
		return
	}

	filename := scanner.Text()
	err := fileio.LoadFromFile(z, filename)
	if err != nil {
		fmt.Printf("加载数据时发生错误：%v\n", err)
	} else {
		fmt.Println("数据加载成功。")
	}
}
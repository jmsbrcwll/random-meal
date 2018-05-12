package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"strconv"
	"bufio"
	"os"
	"encoding/csv"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("api/v1")
	{
		v1.GET("/instructions", GetInstructions)
	}

	return router
}

func GetInstructions(c *gin.Context) {
	c.JSON(200, gin.H{"ok": "Welcome to Chicago!"})
	// curl -i http://localhost:8080/api/v1/Instructions
}

func GetFoodStuffs(filename string) []FoodStuff {
	lines, err := readFoodStuffsFromFile(filename)
	if err != nil {
		fmt.Println(err)
		return []FoodStuff{}
	}
	return lines
}


func readFoodStuffsFromFile(filename string) ([]FoodStuff, error) {
	var lines []FoodStuff
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file ", err)
		os.Exit(1)
	}
	defer f.Close()
	r := csv.NewReader(bufio.NewReader(f))
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	for _, line := range records {
		lines = append(lines, getFoodStuffFromFile(line))
	}
	return lines, nil

}

func getFoodStuffFromFile(line []string) FoodStuff {
	name := line[1]
	calories, err := strconv.ParseFloat(line[2], 64)
	if err != nil {
		fmt.Println(err)
	}
	protein, err := strconv.ParseFloat(line[3], 64)
	if err != nil {
		fmt.Println(err)
	}
	sugars, err := strconv.ParseFloat(line[4], 64)
	if err != nil {
		fmt.Println(err)
	}
	return FoodStuff{name:name, calories:calories, protein:protein, sugars:sugars}



}

type FoodStuff struct {
	name string
	calories float64
	protein float64
	sugars float64
}

func main() {
	router := SetupRouter()
	beef := GetFoodStuffs("beef.csv")
	veg := GetFoodStuffs("veg.csv")
	carbs := GetFoodStuffs("carbs.csv")
	fmt.Println(beef)
	fmt.Println(veg)
	fmt.Println(carbs)
	router.Run(":8080")
}
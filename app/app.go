package app

import (
	"strconv"
	"fmt"
	"os"
	"encoding/csv"
	"bufio"
	"github.com/gin-gonic/gin"
	"math/rand"
)

func GetInstructions(c *gin.Context) {
	caloriesInput := c.Request.URL.Query().Get("calories")
	calories, err := strconv.ParseInt(caloriesInput, 10, 32)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, "supply a proper calorie")
		return
	}

	//otherwise, we can proceed to get some food-stuffs

	beef := GetFoodStuffs("beef.csv")
	veg := GetFoodStuffs("veg.csv")
	carbs := GetFoodStuffs("carbs.csv")

	var meal = make(map[string]*FoodStuff)
	beefSelection := getRandomFoodStuff(beef)
	meal["beef"] = beefSelection
	meal["veg"] = getRandomFoodStuff(veg)
	meal["carbs"] = getRandomFoodStuff(carbs)

	totalCalories := meal["beef"].Calories + meal["veg"].Calories + meal["carbs"].Calories
	c.JSON(200, gin.H{"meal": meal, "calories": calories, "totalCalories": totalCalories})
}

func getRandomFoodStuff(foodGroup []FoodStuff) *FoodStuff {
	randomIndex := rand.Int() % len(foodGroup)
	return &foodGroup[randomIndex]
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
	return FoodStuff{Name:name, Calories:calories, Protein:protein, Sugars:sugars}



}

type FoodStuff struct {
	Name string `json:"name"`
	Calories float64 `json:"calories"`
	Protein float64 `json:"protein"`
	Sugars float64 `json:"sugars"`
}

package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
)

type pepper struct {
	Name     string
	HeatUnit int
	Peppers  int
}

func main() {
	PromptSelect()
	PromptCustomSelect()
	PromptConfirm()
	PromptValidate()
	PromptCustomValidate()

}

func PromptValidate() {
	// 入力が不正な場合errorを返す関数を作成
	validate := func(input string) error {
		_, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("Invalid number")
		}
		return nil
	}

	// インタラクションの表示やバリデーションを設定
	prompt := promptui.Prompt{
		Label:    "Number", // 表示する文言
		Validate: validate, // validate
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", result)
}

func PromptSelect() {
	prompt := promptui.Select{
		Label: "Select Day",
		Items: []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday",
			"Saturday", "Sunday"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", result)
}

func PromptConfirm() {
	prompt := promptui.Prompt{
		Label:     "Delete Resource",
		IsConfirm: true,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", result)
}

func PromptCustomValidate() {
	validate := func(input string) error {
		_, err := strconv.ParseFloat(input, 64)
		return err
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     "Spicy Level",
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You answered %s\n", result)
}

func PromptCustomSelect() {
	peppers := []pepper{
		{Name: "Bell Pepper", HeatUnit: 0, Peppers: 0},
		{Name: "Banana Pepper", HeatUnit: 100, Peppers: 1},
		{Name: "Poblano", HeatUnit: 1000, Peppers: 2},
		{Name: "Jalapeño", HeatUnit: 3500, Peppers: 3},
		{Name: "Aleppo", HeatUnit: 10000, Peppers: 4},
		{Name: "Tabasco", HeatUnit: 30000, Peppers: 5},
		{Name: "Malagueta", HeatUnit: 50000, Peppers: 6},
		{Name: "Habanero", HeatUnit: 100000, Peppers: 7},
		{Name: "Red Savina Habanero", HeatUnit: 350000, Peppers: 8},
		{Name: "Dragon’s Breath", HeatUnit: 855000, Peppers: 9},
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U0001F336 {{ .Name | cyan }} ({{ .HeatUnit | red }})",
		Inactive: "  {{ .Name | cyan }} ({{ .HeatUnit | red }})",
		Selected: "\U0001F336 {{ .Name | red | cyan }}",
		Details: `
--------- Pepper ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Heat Unit:" | faint }}	{{ .HeatUnit }}
{{ "Peppers:" | faint }}	{{ .Peppers }}`,
	}

	searcher := func(input string, index int) bool {
		pepper := peppers[index]
		name := strings.Replace(strings.ToLower(pepper.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     "Spicy Level",
		Items:     peppers,
		Templates: templates,
		Size:      4,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose number %d: %s\n", i+1, peppers[i].Name)
}

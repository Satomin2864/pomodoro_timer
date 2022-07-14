package console

import (
	"fmt"
	"strconv"

	"github.com/manifoldco/promptui"
)

type pepper struct {
	Name     string
	HeatUnit int
	Peppers  int
}

func Console() (string, error) {
	validate := func(input string) error {
		num, err := strconv.ParseFloat(input, 64)
		switch {
		case num > 30:
			return fmt.Errorf("Over 30 minutes")
		case num < 1:
			return fmt.Errorf("Less than 1 minute")
		}
		return err
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     "wark time",
		Templates: templates,
		Validate:  validate,
	}
	return prompt.Run()
	// result, err := prompt.Run()

	// if err != nil {
	// 	fmt.Printf("Prompt failed %v\n", err)
	// 	return "", err
	// }

	// fmt.Printf("You answered %s\n", result)
	// return result, nil
}

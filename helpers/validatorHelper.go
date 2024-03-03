package helpers

import (
	"fmt"
	"gokanban/structs/projectitem"
	strc "strconv"
	s "strings"
)

func ValidateProjectItem(p projectitem.ProjectItemViewModel) (map[string][]string, bool) {
	containsErrors := false
	validations := map[string][]string{
		"title":         ValidateTitle(p.Title),
		"description":   ValidateDescription(p.Description),
		"estimatedtime": ValidateTime(p.EstimatedTime),
	}
	for _, v := range validations {
		if len(v) > 0 {
			containsErrors = true
			break
		}
	}

	return validations, containsErrors
}

func ValidateTitle(title string) []string {
	errors := []string{}
	if len(title) == 0 {
		errors = append(errors, "Tittel kan ikke være tom")
		return errors
	}
	errors = append(errors, validateIllegalChars(title)...)
	if len(title) > 100 {
		errors = append(errors, "Tittel er for lang")
	}

	return errors
}

func ValidateDescription(description string) []string {
	errors := []string{}
	errors = append(errors, validateIllegalChars(description)...)
	if len(description) > 200 {
		errors = append(errors, "Beskrivelsen er for lang")
	}

	return errors
}

func ValidateTime(estimatedtime string) []string {
	errors := []string{}
	parsedInt, err := strc.ParseInt(estimatedtime, 10, 64)
	if err != nil {
		errors = append(errors, "Verdien må være et heltall")
		return errors
	}

	err = validateNumberRange(int(parsedInt), 0, 1000)
	if err != nil {
		errors = append(errors, err.Error())
	}

	return errors
}

func validateNumberRange(value int, min int, max int) error {
	if value < min {
		return fmt.Errorf("Verdien må være større enn %d", min)
	}
	if value > max {
		return fmt.Errorf("Verdien må være mindre enn %d", max)
	}

	return nil
}

func validateIllegalChars(value string) []string {
	illegalChars := []string{"<", ">", "~", "$", "{", "}", "%"}
	errors := []string{}
	for _, char := range illegalChars {
		if s.ContainsAny(value, char) {
			errors = append(errors, fmt.Sprintf("Verdien inneholder et ulovlig tegn: %s", char))
		}
	}

	return errors
}

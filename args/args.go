package args

import (
	"errors"
	"regexp"
)

func IsValidRoleMention(arg string) (bool, error) {
	pattern := "^<@&([0-9]{17,19})>$" // Matches role arguments

	r, err := regexp.Compile(pattern)
	if err != nil {
		return false, err
	}

	match := r.MatchString(arg)

	return match, nil
}

func ParseRoleIDFromMention(mention string) (string, error) {
	pattern := "[0-9]\\d+"

	r, err := regexp.Compile(pattern)
	if err != nil {
		return "", err
	}
	isValidMention, err := IsValidRoleMention(mention)
	if err != nil {
		return "", err
	}
	if isValidMention {
		return r.FindString(mention), nil
	} else {
		return "", errors.New("invalid mention")
	}
}

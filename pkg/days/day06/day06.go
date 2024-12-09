package days

import "errors"

func Process(content *[]byte) (string, string, error) {
	result1, err := part1(content)
	if err != nil {
		return "", "", err
	}
	result2, err := part2(content)
	if err != nil {
		return result1, "", err
	}
	return result1, result2, nil
}

func part1(content *[]byte) (string, error) {
	// todo: implement
	return "", errors.New("not implemented")
}

func part2(content *[]byte) (string, error) {
	// todo: implement
	return "", errors.New("not implemented")
}

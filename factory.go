package main

import "fmt"

func CreateObjectFactory(objType string, props *Properties) (obj Object, err error) {
	switch objType {
	case "Player":
		obj = NewPlayer(props)
		return obj, nil
	case "Enemy":
		obj, err = NewEnemy(props)
		if err != nil {
			return nil, err
		}
		return obj, nil
	}

	return nil, fmt.Errorf("invalid object type, %q", objType)
}

package main

func CreateObjectFactory(objType string, props *Properties) (obj Object, err error) {
	switch objType {
	case "Player":
		obj = NewPlayer(props)
	case "Enemy":
		obj, err = NewEnemy(props)
		if err != nil {
			return nil, err
		}
	}

	return obj, nil
}

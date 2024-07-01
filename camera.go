package main

type Camera struct {
	instance *Camera
}

func (c *Camera) GetInstance() *Camera {
	if c.instance == nil {
		return &Camera{}
	}

	return c.instance
}

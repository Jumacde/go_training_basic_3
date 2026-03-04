package main

type Calcu interface {
	GetCalc() (uint, uint, uint)
	calc_add() (uint, uint, uint)
}

type Calculate struct {
	x, y, result uint
}

func (c *Calculate) SetCalc(x, y, result uint) {
	c.x = x
	c.y = y
	c.result = result
	c.calc_add()
}

func (c *Calculate) GetCalc() (uint, uint, uint) {
	return c.x, c.y, c.result
}

func (c *Calculate) calc_add() (uint, uint, uint) {
	c.result = c.x + c.y
	return c.x, c.y, c.result

}

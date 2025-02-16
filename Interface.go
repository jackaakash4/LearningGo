package main

import (
    "fmt"
    "math"
)

type geometry interface{
    area() float64
    peri() float64
}

type rect struct{
    width, height float64
}

type circle struct{
    radius float64
}

func (r rect) area float64{
    return r.width * r.height
}

funct (r rect) peri float64{
    return 2*(r.width + r.height)
}

func (c circle) area float64{
    return math.Pi * c.radius^2
}

func (c circle) peri float64{
    return 2 * math.Pi * c.radius\
}


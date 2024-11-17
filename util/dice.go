package util

import (
    "math/rand"
    "time"
)


type DiceRoller interface {
    Roll() int
}


type StandardDice struct {
    sides int
}


func NewStandardDice() *StandardDice {
    return &StandardDice{
        sides: 6,
    }
}


func (d *StandardDice) Roll() int {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(d.sides) + 1
}
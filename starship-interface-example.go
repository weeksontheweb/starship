package main

import (
	"fmt"
)

// Multiple types of starships will have common functionality.
// Therefore put the common functionality in the interface.
type starship interface {
	move(direction string, distance int)
	calculateFuelUsed(distance int) int
}

// Lets have a Battle Cruiser for the heavy artillery.
type battlecruiser struct {
	fuelAmountLeft int
}

//Lets have a Star Fighter (something light and fast).
type starfighter struct {
	fuelAmountLeft int
}

// The Battle Cruiser will have different movement capabilities
// to the Star Fighter, so we calculate their movements differently.
// Also fuel consumption is different.
// Hence the reason to give each Star Ship it's own Methods.
// All methods defined in the interface, must be implemented,
// otherwise Go won't know which interface to implement.

func (s battlecruiser) move(direction string, amount int) {
	//Move code for the Battle Cruiser goes here.
}

func (s battlecruiser) calculateFuelUsed(distance int) int {
	s.fuelAmountLeft = s.fuelAmountLeft - (distance * 2)
	return distance * 2
}

func (s starfighter) move(direction string, amount int) {
	//Move code for the Star Fighter goes here.
}

func (s starfighter) calculateFuelUsed(distance int) int {
	s.fuelAmountLeft = s.fuelAmountLeft - distance
	return distance
}

// This is where the movement and subsequent fuel consumption is calculated.
// Although the move capabilities of each ship is different, there will
// be common code used in the calculation (such as the notification).
func moveShip(s starship, direction string, distance int) {

	// The actual movement is different for each star ship.
	// So call the specific movement code for that ship.
	s.move("north", distance)

	//Same for fuel consumption
	fuelUsed := s.calculateFuelUsed(distance)

	// The notification is common to all ships.
	fmt.Printf("Moving the %T.\nDirection: %s\nDistance: %d km(s)\nFuel used: %d litre(s)\n\n", s, direction, distance, fuelUsed)
}

func main() {

	// Lets define our starship types.
	// Behind the scenes these are linked to the interface.
	// We do not have to explicitly state which interface they
	// link to.
	var bc battlecruiser
	var sf starfighter

	// Let's move the ships.
	// This can now be done by calling the moveShip function.
	// This will accept any Star Ship via the interface.
	// Also note, that the methods defined in the interface will be passed.
	// Any ship specific info (variables or ship specific methods, not defined
	// in the interface) will not be available.
	moveShip(bc, "north", 2)
	moveShip(sf, "east", 1)

}

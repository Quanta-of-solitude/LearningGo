package main

import "fmt"

// Create the function fuelGauge() here
func fuelGauge(fuel int){
  fmt.Println("Remaining fuel: ",fuel)
}

// Create the function calculateFuel() here
func calculateFuel(planet string) int{
  var fuel int
  switch planet{
    case "Venus":
      fuel = 300000
    case "Mercury":
      fuel = 500000
    case "Mars":
      fuel = 700000
    default:
      fuel = 0
  }

  return fuel
}

// Create the function greetPlanet() here
func greetPlanet(planet string){
  fmt.Printf("\nTrying to go to %v \n",planet)
}

// Create the function cantFly() here
func cantFly(){
   fmt.Println("We do not have the available fuel to fly there\n")
}

// Create the function flyToPlanet() here
func flyToPlanet(planet string, fuel int) int{
  var fuelRemaining, fuelCost int
  fuelRemaining = fuel
  fuelCost = calculateFuel(planet)

  if (fuelRemaining >= fuelCost){
    greetPlanet(planet)
    fuelRemaining -= fuelCost
  }

  if fuelCost >=fuelRemaining{
    cantFly()
  }
  return fuelRemaining
}

func main() {
  // Test your functions!
  fuel,planetChoice:=1000000, "Venus"
  fuel = flyToPlanet(planetChoice, fuel)
  fuelGauge(fuel)
  
  planetChoice= "Mars"
  fuel = flyToPlanet(planetChoice, fuel)
  fuelGauge(fuel)


}

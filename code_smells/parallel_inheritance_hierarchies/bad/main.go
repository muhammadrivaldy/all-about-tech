package main

import "fmt"

/*
	# Structure codes:

	- Brand
		- Toyota
			- Crossover
			- Sedan
		- Honda
			- Crossover
			- Sedan
	- Electric Brand
		- Tesla
			- Crossover
			- Sedan
*/

type Brand struct {
	Name string
}

type Toyota struct {
	Brand
}

type CrossoverToyota struct {
	Toyota
}

func (c *CrossoverToyota) BuildCar() {
	fmt.Printf("We build %s crossover\n", c.Name)
}

type SedanToyota struct {
	Toyota
}

func (s *SedanToyota) BuildCar() {
	fmt.Printf("We build %s sedan\n", s.Name)
}

type Honda struct {
	Brand
}

type CrossoverHonda struct {
	Honda
}

func (c *CrossoverHonda) BuildCar() {
	fmt.Printf("We build %s crossover\n", c.Name)
}

type SedanHonda struct {
	Honda
}

func (s *SedanHonda) BuildCar() {
	fmt.Printf("We build %s sedan\n", s.Name)
}

type ElectricBrand struct {
	Name string
}

type Tesla struct {
	ElectricBrand
}

type CrossoverTesla struct {
	Tesla
}

func (c *CrossoverTesla) BuildCar() {
	fmt.Printf("We build %s crossover\n", c.Name)
}

type SedanTesla struct {
	Tesla
}

func (s *SedanTesla) BuildCar() {
	fmt.Printf("We build %s sedan\n", s.Name)
}

func main() {
	brand := Brand{Name: "Toyota"}
	toyota := Toyota{Brand: brand}
	crossoverToyota := CrossoverToyota{Toyota: toyota}

	crossoverToyota.BuildCar()
}

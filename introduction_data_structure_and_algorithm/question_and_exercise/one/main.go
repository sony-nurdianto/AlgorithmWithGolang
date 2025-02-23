// Give  an example where you can use composite pattern
package main

import "fmt"

type OrganizationComponent interface {
	showDetails()
}

type Division struct {
	Name string
}

func (dv *Division) getName() string {
	return "Division: " + dv.Name
}

type Department struct {
	Name     string
	Division []Division
}

func (d *Department) getName() string {
	return "Department: " + d.Name
}

func (dp *Department) addDivision(dv Division) {
	dp.Division = append(dp.Division, dv)
}

func (dp *Department) getDivision() []Division {
	return dp.Division
}

func (dp *Department) showDetails() {
	fmt.Println(dp.getName())
	defer fmt.Println()
	for _, dv := range dp.getDivision() {
		fmt.Println(dv.getName())
	}
}

type Companys struct {
	Name       string
	Department []Department
}

func (c *Companys) addDepartment(d Department) {
	c.Department = append(c.Department, d)
}

func (c *Companys) getDepartment() []Department {
	return c.Department
}

func (c *Companys) showDetails() {
	fmt.Println(c.Name)
	for _, dp := range c.getDepartment() {
		dp.showDetails()
	}
}

func main() {
	company := &Companys{
		Name: "Futures Company",
	}

	researchAndDevelopment := Department{
		Name: "Research and Development",
	}

	RNDElektronik := Division{
		Name: "Elektronik",
	}

	RNDSoftwareEngineering := Division{
		Name: "Software",
	}

	researchAndDevelopment.addDivision(RNDElektronik)
	researchAndDevelopment.addDivision(RNDSoftwareEngineering)

	company.addDepartment(researchAndDevelopment)

	softwareEngineering := Department{
		Name: "software Engineering",
	}

	systemEngineering := Division{
		Name: "System Engineering",
	}

	iot := Division{
		Name: "Internet Of Things",
	}

	softwareEngineering.addDivision(systemEngineering)
	softwareEngineering.addDivision(iot)

	company.addDepartment(softwareEngineering)

	company.showDetails()
}

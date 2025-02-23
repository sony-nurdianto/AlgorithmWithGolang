package main

import "fmt"

const (
	CustomerType = "customer"
	EmployeeType = "employee"
	ManagerType  = "manager"
	AddressType  = "address"
)

type DataTransferObjectFactory struct {
	pool map[string]DataTransferObject
}

func (factory DataTransferObjectFactory) getDataTransferObject(dtoType string) DataTransferObject {
	dto := factory.pool[dtoType]

	if dto == nil {
		fmt.Printf("Membuat Dto baru dengan tipe: %v\n", dto)
		switch dtoType {
		case CustomerType:
			factory.pool[dtoType] = Customer{id: "1"}
		case EmployeeType:
			factory.pool[dtoType] = Employee{id: "1"}
		case ManagerType:
			factory.pool[dtoType] = Manager{id: "1"}
		case AddressType:
			factory.pool[dtoType] = Address{id: "1"}
		}
		dto = factory.pool[dtoType]
	}

	return dto
}

type DataTransferObject interface {
	getId() string
}

type Customer struct {
	id   string
	name string
	ssn  string
}

func (customer Customer) getId() string {
	return customer.id
}

type Employee struct {
	id   string
	name string
}

func (employee Employee) getId() string {
	return employee.id
}

type Manager struct {
	id   string
	name string
	dept string
}

func (manager Manager) getId() string {
	return manager.id
}

type Address struct {
	id          string
	streetLine1 string
	streetLine2 string
	state       string
	city        string
}

func (address Address) getId() string {
	return address.id
}

func main() {
	factory := DataTransferObjectFactory{
		pool: make(map[string]DataTransferObject),
	}

	customer := factory.getDataTransferObject(CustomerType)
	fmt.Printf("Customer ID: %s\n", customer.getId())

	employee := factory.getDataTransferObject(EmployeeType)
	fmt.Printf("Employee ID: %s\n", employee.getId())

	manager := factory.getDataTransferObject(ManagerType)
	fmt.Printf("Manager ID: %s\n", manager.getId())

	address := factory.getDataTransferObject(AddressType)
	fmt.Printf("Adress ID: %s\n", address.getId())
}

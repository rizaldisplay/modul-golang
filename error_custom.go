package main

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{Message: "ID cannot be empty"}
	}

	if id != "rizal" {
		return &notFoundError{Message: "Data not found"}
	}

	return nil
}

func main() {
	err := SaveData("", nil)
	if err != nil {
		if validationError, ok := err.(*validationError); ok {
			println("Validation Error:", validationError.Error())
		} else if notFoundError, ok := err.(*notFoundError); ok {
			println("Not Found Error:", notFoundError.Error())
		} else {
			println("Unknown Error:", err.Error())
		}
	} else {
		println("Data saved successfully")
	}
}

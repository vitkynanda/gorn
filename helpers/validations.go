package helpers

import (
	"gorm_crud/dtos"
)

func GenerateValidationResponse(err error) (response dtos.ValidationResponse) {
	response.Success = false
	response.Message = "Bad Request" 


// var validations []dtos.Validation

	// get validation errors
	// validationErrors := err.(validator.ValidationErrors)

	//  for _, value := range validationErrors {
	// // 	// get field & rule (tag)
	// 	// field, rule := value.Field, value.Tag


	// 	// fmt.Println(value.Field())

	// 	// // create validation object
	// 	// validation := dtos.Validation{Field: field, Message: langs.GenerateValidationMessage(field, rule)}

	// 	// // add validation object to validations
	// 	// validations = append(validations, validation)
	// }

	// set Validations response
	

	return response
}

package handler

import "fmt"

func erroParamIsRequired(name, typ string)error{
	return fmt.Errorf("param: role %s (type: %s) is required", name, typ)
}

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"Company"`
	Location string `json:"Location"`
	Remote   *bool  `json:"Remote"`
	Link     string `json:"Link"`
	Salary   int64  `json:"Salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" && r.Company== "" && r.Location =="" && r.Remote==nil && r.Salary <= 0{
		return fmt.Errorf("request body is empty or bad format")
	}
	if r.Role == "" {
		return erroParamIsRequired("role", "string")
	}
	if r.Company == "" {
		return erroParamIsRequired("company", "string")
	}
	if r.Location == "" {
		return erroParamIsRequired("location", "string")
	}
	if r.Link == "" {
		return erroParamIsRequired("link", "string")
	}
	if r.Remote == nil {
		return erroParamIsRequired("remote", "bool")
	}
	if r.Salary <= 0 {
		return erroParamIsRequired("salary", "int64")
	}
	return nil
}
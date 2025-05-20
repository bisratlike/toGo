package validator

import (
    "encoding/json"
    "net/http"
    "github.com/go-playground/validator/v10"
    "github.com/bisratlike/toGo/pkg/response"
)

var validate = validator.New()

func ParseAndValidate(w http.ResponseWriter, r *http.Request, dst interface{}) bool {
    if err := json.NewDecoder(r.Body).Decode(dst); err != nil {
        response.Error(w, http.StatusBadRequest, "Invalid JSON body", err)
        return false
    }

    if err := validate.Struct(dst); err != nil {
        var errs []string
        for _, e := range err.(validator.ValidationErrors) {
            errs = append(errs, e.Field()+": "+e.Tag())
        }
        response.ValidationError(w, errs)
        return false
    }
    return true
}
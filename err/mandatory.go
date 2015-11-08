package err

type MandatoryArgumentError struct {
    Key string
}

func (e *MandatoryArgumentError) Error() string {
    return "Mandatory argument " + e.Key + " is missing"
}
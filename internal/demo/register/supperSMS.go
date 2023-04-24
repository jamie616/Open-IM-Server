package register

type SupperSMS struct {
}

func (s SupperSMS) SendSms(code int, phoneNumber string) (resp interface{}, err error) {
	return nil, nil
}

func NewSupperSMS() (*SupperSMS, error) {
	return &SupperSMS{}, nil
}

package phone

type Phone struct {
	subscribeNumber         string //SN
	nationalDestinationCode string //NDC
	contryCode              string //CC
	isValid                 bool
	err                     error
}

func NewPhone(subscribeNumber, nationalDestinationCode, contryCode string) *Phone {
	return &Phone{
		subscribeNumber:         subscribeNumber,
		nationalDestinationCode: nationalDestinationCode,
		contryCode:              contryCode,
	}
}

func NewPhone2(number string) *Phone {

	return &Phone{
		subscribeNumber:         PhoneSN(number),
		nationalDestinationCode: PhoneNDC(number),
		contryCode:              PhoneCC(number),
	}
}

func PhoneCC(phone string) string {
	//TODO
	return phone
}
func PhoneNDC(phone string) string {
	//TODO
	return phone
}

func PhoneSN(phone string) string {
	//TODO
	return phone
}

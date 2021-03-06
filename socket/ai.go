package socket

type AddressInfo struct {
	HOST string
	PORT string
	TYPE string
}

func NewAddressInfo(service int) AddressInfo {
	// for client
	if service == 1 {
		return AddressInfo{"192.168.1.6", "6969", "tcp"}
	}

	return AddressInfo{"192.168.1.6", "6969", "tcp"}
}

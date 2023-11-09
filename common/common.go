package common

type Server_config struct {
	ServerName, ServerIpPort, ServerProtocol string
	ServerPort                               int
}

func stringToBytes(s ...string) [][]byte {
	var result [][]byte
	for _, str := range s {
		bytes := []byte(str)
		result = append(result, bytes)
	}
	return result
}

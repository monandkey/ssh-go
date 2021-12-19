package ssh

// singleNodeStruct is a function that returns the SshMethod
func singleNodeStruct() SshMethod {
	return &singleNode{}
}

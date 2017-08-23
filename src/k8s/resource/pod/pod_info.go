package pod

//acquire pod information
type PodInfo struct{
	Name string
	Ip string
}

func (pi *PodInfo) GetName() string {
	return pi.Name
}

func (pi *PodInfo) GetIp() string {
	return pi.Ip
}


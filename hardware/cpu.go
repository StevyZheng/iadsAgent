package hardware

import (
	"fmt"
	"iadsAgent/common"
	"io/ioutil"
	"strings"
)

type CpuHwInfo struct {
	Model     string
	Count     int
	CoreCount int
	Stepping  string
}

func (e *CpuHwInfo) GetCpuHwInfo() (err error) {
	tmp, err := ioutil.ReadFile("/proc/cpuinfo")
	tmpStr := strings.Replace(string(tmp), "\n", "", 1)
	if err != nil {
		fmt.Println(err)
	}
	e.Model = common.SearchSplitStringColumnFirst(tmpStr, ".*model name.*", ":", 2)
	e.Stepping = common.SearchSplitStringColumnFirst(tmpStr, ".*stepping.*", ":", 2)
	countTmp1 := common.SearchString(tmpStr, ".*physical id.*")
	countTmp := common.UniqStringList(countTmp1)
	e.Count = len(countTmp)
	coreCountTmp1 := common.SearchString(tmpStr, ".*processor.*")
	coreCountTmp := common.UniqStringList(coreCountTmp1)
	e.CoreCount = len(coreCountTmp)
	return err
}

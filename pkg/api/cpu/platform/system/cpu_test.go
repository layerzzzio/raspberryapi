// package system_test

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/raspibuddy/rpi/utl/mock/mocksys"
// 	cext "github.com/shirou/gopsutil/cpu"
// )

// func TestListError(t *testing.T) {
// 	csys := &mocksys.CPU{
// 		// ListFn: func() ([]cext.InfoStat, []float64, []cext.TimesStat, error) {
// 		// 	return nil, nil, nil, errors.New("test error")
// 		// },
// 		CPUInfoFn: func() ([]cext.InfoStat, error) {
// 			return []cext.InfoStat{}, nil
// 		},
// 	}
// 	i, err := csys.CPUInfo()
// 	fmt.Println(i)
// 	fmt.Println(err)

// 	// info, percent, vCore, err := csys.List()
// 	// fmt.Println(info)
// 	// fmt.Println(percent)
// 	// fmt.Println(vCore)
// 	// assert.Nil(t, info)
// 	// assert.Nil(t, percent)
// 	// assert.Nil(t, vCore)
// 	// assert.NotNil(t, err)
// 	// assert.EqualValues(t, "Could not retrieve the CPU metrics", err.(*echo.HTTPError).Message)
// 	// assert.EqualValues(t, http.StatusInternalServerError, err.(*echo.HTTPError).Code)
// }

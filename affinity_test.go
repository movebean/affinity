package affinity_test

//import "affinity"
import "github.com/movebean/affinity"
import "runtime"
import "testing"
import "math"

var cpuNum = float64(runtime.NumCPU())

var getTestCases = []struct {
	args     int
	expected int64
}{
	{0, int64(math.Pow(2, cpuNum)) - 1},
}

var setTestCases = []struct {
	args     int
	expected int64
}{
	{0, int64(math.Pow(2, cpuNum)) - 2},
}

func TestGet(t *testing.T) {
	for _, ca := range getTestCases {
		var cpuSet int64
		affinity.GetAffinity(ca.args, &cpuSet)
		if cpuSet != ca.expected {
			t.Errorf("expected %d, but get %d", ca.expected, cpuSet)
		}
	}
}

func TestSet(t *testing.T) {
	for _, ca := range setTestCases {
		var cpuSet int64
		affinity.GetAffinity(ca.args, &cpuSet)
		cpuSet = cpuSet & 0XFFFFFFE
		affinity.SetAffinity(ca.args, &cpuSet)
		affinity.GetAffinity(ca.args, &cpuSet)
		if cpuSet != ca.expected {
			t.Errorf("expected %d, but get %d", ca.expected, cpuSet)
		}
	}
}

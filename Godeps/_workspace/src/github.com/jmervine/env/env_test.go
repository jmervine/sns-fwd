package env

import (
	. "github.com/jmervine/sns-fwd/Godeps/_workspace/src/github.com/jmervine/env/_fixtures"

	"fmt"
	"log"
	"os"
	"testing"
	"time"

	. "github.com/jmervine/GoT"
)

var env = "_fixtures/fixtures.env"

func init() {
	UnSetFixtures()
}

func TestLoad(T *testing.T) {
	T.Skip("not testing godotenv internals")
}

func TestOverload(T *testing.T) {
	defer UnSetFixtures()
	os.Setenv("F_INT", "999")

	err := Overload(env)
	Go(T).AssertNil(err)

	// ensure clobber
	Go(T).AssertEqual(os.Getenv("F_INT"), "9")
}

func TestSet(T *testing.T) {
	defer UnSetFixtures()

	for key, val := range Fixtures {
		Set(key, val)

		Go(T).AssertEqual(os.Getenv(key), fmt.Sprintf("%v", val))
	}
}

func TestSetMap(T *testing.T) {
	defer UnSetFixtures()
	SetMap(Fixtures)

	for key, val := range Fixtures {
		Go(T).AssertEqual(os.Getenv(key), fmt.Sprintf("%v", val))
	}
}

func TestGet(T *testing.T) {
	defer UnSetFixtures()

	s := Get("F_STRING")
	Go(T).AssertEqual(s, "")

	SetFixtures()

	s = Get("F_STRING")
	Go(T).AssertEqual(s, "string")

	s = GetString("F_STRING")
	Go(T).AssertEqual(s, "string")
}

func TestRequire(T *testing.T) {
	defer UnSetFixtures()

	s, e := Require("F_STRING")
	Go(T).AssertEqual(e.Error(), "missing required string from F_STRING")
	Go(T).AssertEqual(s, "")

	SetFixtures()

	s, e = Require("F_STRING")
	Go(T).AssertEqual(s, "string")
	Go(T).AssertNil(e)

	s, e = RequireString("F_STRING")
	Go(T).AssertEqual(s, "string")
	Go(T).AssertNil(e)
}

func TestGetOrSet(T *testing.T) {
	defer UnSetFixtures()

	s := GetOrSet("F_STRING", "default")
	Go(T).AssertEqual(s, "default")
	UnSetFixtures()

	SetFixtures()
	s = GetOrSet("F_STRING", "default")
	Go(T).AssertEqual(s, "string")
}

func TestRequireDuration(T *testing.T) {
	defer UnSetFixtures()

	emptyDur := new(time.Duration)

	d, e := RequireDuration("F_DURATION")
	Go(T).RefuteNil(e)
	Go(T).AssertEqual(d, *emptyDur)

	SetFixtures()
	d, e = RequireDuration("F_DURATION")
	Go(T).AssertEqual(d, Fixtures["F_DURATION"])
	Go(T).AssertNil(e)
}

func TestGetDuration(T *testing.T) {
	defer UnSetFixtures()
	SetFixtures()

	d := GetDuration("F_DURATION")
	Go(T).AssertEqual(d, Fixtures["F_DURATION"])
}

func TestGetOrSetDuration(T *testing.T) {
	defer UnSetFixtures()

	def, _ := time.ParseDuration("1d")

	d := GetOrSetDuration("F_DURATION", def)
	Go(T).AssertEqual(d, def)

	SetFixtures()

	d = GetOrSetDuration("F_DURATION", def)
	Go(T).AssertEqual(d, Fixtures["F_DURATION"])
}

func TestGetInt(T *testing.T) {
	defer UnSetFixtures()

	i := GetInt("F_INT")
	Go(T).AssertEqual(i, int(0))

	SetFixtures()

	i = GetInt("F_INT")
	Go(T).AssertEqual(i, Fixtures["F_INT"])
}

func TestGetOrSetInt(T *testing.T) {
	defer UnSetFixtures()

	def := int(2)

	i := GetOrSetInt("F_INT", def)
	Go(T).AssertEqual(i, int(2))

	SetFixtures()

	i = GetInt("F_INT")
	Go(T).AssertEqual(i, Fixtures["F_INT"])
}

func TestRequireInt(T *testing.T) {
	defer UnSetFixtures()

	i, e := RequireInt("F_INT")
	Go(T).RefuteNil(e)
	Go(T).AssertEqual(i, int(0))

	SetFixtures()
	i, e = RequireInt("F_INT")
	Go(T).AssertEqual(i, Fixtures["F_INT"])
	Go(T).AssertNil(e)
}

func TestGetInt32(T *testing.T) {
	defer UnSetFixtures()

	i := GetInt32("F_INT32")
	Go(T).AssertEqual(i, int32(0))

	SetFixtures()

	i = GetInt32("F_INT32")
	Go(T).AssertEqual(i, Fixtures["F_INT32"])
}

func TestGetOrSetInt32(T *testing.T) {
	defer UnSetFixtures()

	def := int32(2)

	i := GetOrSetInt32("F_INT32", def)
	Go(T).AssertEqual(i, int32(2))

	SetFixtures()

	i = GetInt32("F_INT32")
	Go(T).AssertEqual(i, Fixtures["F_INT32"])
}

func TestRequireInt32(T *testing.T) {
	defer UnSetFixtures()

	i, e := RequireInt32("F_INT32")
	Go(T).RefuteNil(e)
	Go(T).AssertEqual(i, int32(0))

	SetFixtures()
	i, e = RequireInt32("F_INT32")
	Go(T).AssertEqual(i, Fixtures["F_INT32"])
	Go(T).AssertNil(e)
}

func TestGetInt64(T *testing.T) {
	defer UnSetFixtures()

	i := GetInt64("F_INT64")
	Go(T).AssertEqual(i, int64(0))

	SetFixtures()

	i = GetInt64("F_INT64")
	Go(T).AssertEqual(i, Fixtures["F_INT64"])
}

func TestGetOrSetInt64(T *testing.T) {
	defer UnSetFixtures()

	def := int64(2)

	i := GetOrSetInt64("F_INT64", def)
	Go(T).AssertEqual(i, int64(2))

	SetFixtures()

	i = GetInt64("F_INT64")
	Go(T).AssertEqual(i, Fixtures["F_INT64"])
}

func TestRequireInt64(T *testing.T) {
	defer UnSetFixtures()

	i, e := RequireInt64("F_INT64")
	Go(T).RefuteNil(e)
	Go(T).AssertEqual(i, float64(0))

	SetFixtures()
	i, e = RequireInt64("F_INT64")
	Go(T).AssertEqual(i, Fixtures["F_INT64"])
	Go(T).AssertNil(e)
}

func TestGetFloat32(T *testing.T) {
	defer UnSetFixtures()

	i := GetFloat32("F_FLOAT32")
	Go(T).AssertEqual(i, float32(0))

	SetFixtures()

	i = GetFloat32("F_FLOAT32")
	Go(T).AssertEqual(i, Fixtures["F_FLOAT32"])
}

func TestGetOrSetFloat32(T *testing.T) {
	defer UnSetFixtures()

	def := float32(2)

	i := GetOrSetFloat32("F_FLOAT32", def)
	Go(T).AssertEqual(i, float32(2))

	SetFixtures()

	i = GetFloat32("F_FLOAT32")
	Go(T).AssertEqual(i, Fixtures["F_FLOAT32"])
}

func TestRequireFloat32(T *testing.T) {
	defer UnSetFixtures()

	i, e := RequireFloat32("F_FLOAT32")
	Go(T).RefuteNil(e)
	Go(T).AssertEqual(i, float32(0))

	SetFixtures()
	i, e = RequireFloat32("F_FLOAT32")
	Go(T).AssertEqual(i, Fixtures["F_FLOAT32"])
	Go(T).AssertNil(e)
}

func TestGetFloat64(T *testing.T) {
	defer UnSetFixtures()

	i := GetFloat64("F_FLOAT64")
	Go(T).AssertEqual(i, float64(0))

	SetFixtures()

	i = GetFloat64("F_FLOAT64")
	Go(T).AssertEqual(i, Fixtures["F_FLOAT64"])
}

func TestGetOrSetFloat64(T *testing.T) {
	defer UnSetFixtures()

	def := float64(2)

	i := GetOrSetFloat64("F_FLOAT64", def)
	Go(T).AssertEqual(i, float64(2))

	SetFixtures()

	i = GetFloat64("F_FLOAT64")
	Go(T).AssertEqual(i, Fixtures["F_FLOAT64"])
}

func TestRequireFloat64(T *testing.T) {
	defer UnSetFixtures()

	i, e := RequireFloat64("F_FLOAT64")
	Go(T).RefuteNil(e)
	Go(T).AssertEqual(i, float64(0))

	SetFixtures()
	i, e = RequireFloat64("F_FLOAT64")
	Go(T).AssertEqual(i, Fixtures["F_FLOAT64"])
	Go(T).AssertNil(e)
}

func TestGetBool(T *testing.T) {
	defer UnSetFixtures()

	i := GetBool("F_BOOL")
	Go(T).AssertEqual(i, false)

	SetFixtures()

	i = GetBool("F_BOOL")
	Go(T).AssertEqual(i, Fixtures["F_BOOL"])
}

func TestGetOrSetBool(T *testing.T) {
	defer UnSetFixtures()

	def := bool(false)

	i := GetOrSetBool("F_BOOL", def)
	Go(T).Refute(i)

	SetFixtures()

	i = GetBool("F_BOOL")
	Go(T).AssertEqual(i, Fixtures["F_BOOL"])
}

func TestRequireBool(T *testing.T) {
	defer UnSetFixtures()

	i, e := RequireBool("F_BOOL")
	Go(T).RefuteNil(e)
	Go(T).Refute(i)

	SetFixtures()
	i, e = RequireBool("F_BOOL")
	Go(T).AssertEqual(i, Fixtures["F_BOOL"])
	Go(T).AssertNil(e)
}

func Test_toString(T *testing.T) {
	Go(T).AssertEqual(toString(9), "9")
}

func Test_onError(T *testing.T) {
	Go(T).AssertNil(onError(nil))
	Go(T).RefuteNil(onError(fmt.Errorf("error")))
}

func Example() {
	PanicOnRequire = true
	Set("EX_PORT", 3000)

	var (
		addr  string
		port  int
		debug bool
		err   error
	)

	port, err = RequireInt("EX_PORT")
	if err != nil {
		log.Fatal(err)
	}

	addr = GetString("EX_ADDR")

	debug = GetOrSetBool("EX_DEBUG", false)

	fmt.Printf("addr=%v port=%v debug=%v", addr, port, debug)
	// Output: addr= port=3000 debug=false
}

func ExampleSet() {
	Set("SOME_INT", 1)
}

func ExampleGet() {
	// where: SOME_INT=1
	i := Get("SOME_INT")

	if i == "1" {
		fmt.Printf("%v", i)
	}
}

func ExampleGetInt() {
	// where: SOME_INT=1
	i := GetInt("SOME_INT")

	if i == 1 {
		fmt.Printf("%v", i)
	}
}

func ExampleLoad() {
	defer UnSetFixtures()

	Set("F_STRING", "old_string")

	// using `_fixtures/fixtures.env`
	Load(env)

	var f = GetFloat32("F_FLOAT32")
	var s = GetString("F_STRING")
	var b = GetBool("F_BOOL")
	var i = GetInt("F_INT")

	fmt.Printf("F_FLOAT32 ::: %v\n", f)
	fmt.Printf("F_STRING  ::: %v\n", s)
	fmt.Printf("F_BOOL    ::: %v\n", b)
	fmt.Printf("F_INT     ::: %v\n", i)

	// Output:
	// F_FLOAT32 ::: 9.1
	// F_STRING  ::: old_string
	// F_BOOL    ::: false
	// F_INT     ::: 9
}

func ExampleOverload() {
	defer UnSetFixtures()

	Set("F_STRING", "old_string")

	// using `_fixtures/fixtures.env`
	Overload(env)

	var f = GetFloat32("F_FLOAT32")
	var s = GetString("F_STRING")
	var b = GetBool("F_BOOL")
	var i = GetInt("F_INT")

	fmt.Printf("F_FLOAT32 ::: %v\n", f)
	fmt.Printf("F_STRING  ::: %v\n", s)
	fmt.Printf("F_BOOL    ::: %v\n", b)
	fmt.Printf("F_INT     ::: %v\n", i)

	// Output:
	// F_FLOAT32 ::: 9.1
	// F_STRING  ::: sample file
	// F_BOOL    ::: false
	// F_INT     ::: 9
}

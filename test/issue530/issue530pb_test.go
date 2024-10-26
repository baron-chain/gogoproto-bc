// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: issue530/issue530.proto

package issue530

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	github_com_cosmos_gogoproto_jsonpb "github.com/cosmos/gogoproto/jsonpb"
	github_com_cosmos_gogoproto_proto "github.com/cosmos/gogoproto/proto"
	proto "github.com/cosmos/gogoproto/proto"
	math "math"
	math_rand "math/rand"
	testing "testing"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func TestFoo5Proto(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedFoo5(popr, false)
	dAtA, err := github_com_cosmos_gogoproto_proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &Foo5{}
	if err := github_com_cosmos_gogoproto_proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	littlefuzz := make([]byte, len(dAtA))
	copy(littlefuzz, dAtA)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
	if len(littlefuzz) > 0 {
		fuzzamount := 100
		for i := 0; i < fuzzamount; i++ {
			littlefuzz[popr.Intn(len(littlefuzz))] = byte(popr.Intn(256))
			littlefuzz = append(littlefuzz, byte(popr.Intn(256)))
		}
		// shouldn't panic
		_ = github_com_cosmos_gogoproto_proto.Unmarshal(littlefuzz, msg)
	}
}

func TestFoo5MarshalTo(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedFoo5(popr, false)
	size := p.Size()
	dAtA := make([]byte, size)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(dAtA)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &Foo5{}
	if err := github_com_cosmos_gogoproto_proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestBar1Proto(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar1(popr, false)
	dAtA, err := github_com_cosmos_gogoproto_proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &Bar1{}
	if err := github_com_cosmos_gogoproto_proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	littlefuzz := make([]byte, len(dAtA))
	copy(littlefuzz, dAtA)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
	if len(littlefuzz) > 0 {
		fuzzamount := 100
		for i := 0; i < fuzzamount; i++ {
			littlefuzz[popr.Intn(len(littlefuzz))] = byte(popr.Intn(256))
			littlefuzz = append(littlefuzz, byte(popr.Intn(256)))
		}
		// shouldn't panic
		_ = github_com_cosmos_gogoproto_proto.Unmarshal(littlefuzz, msg)
	}
}

func TestBar1MarshalTo(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar1(popr, false)
	size := p.Size()
	dAtA := make([]byte, size)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(dAtA)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &Bar1{}
	if err := github_com_cosmos_gogoproto_proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestBar2Proto(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar2(popr, false)
	dAtA, err := github_com_cosmos_gogoproto_proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &Bar2{}
	if err := github_com_cosmos_gogoproto_proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	littlefuzz := make([]byte, len(dAtA))
	copy(littlefuzz, dAtA)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
	if len(littlefuzz) > 0 {
		fuzzamount := 100
		for i := 0; i < fuzzamount; i++ {
			littlefuzz[popr.Intn(len(littlefuzz))] = byte(popr.Intn(256))
			littlefuzz = append(littlefuzz, byte(popr.Intn(256)))
		}
		// shouldn't panic
		_ = github_com_cosmos_gogoproto_proto.Unmarshal(littlefuzz, msg)
	}
}

func TestBar2MarshalTo(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar2(popr, false)
	size := p.Size()
	dAtA := make([]byte, size)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(dAtA)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &Bar2{}
	if err := github_com_cosmos_gogoproto_proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestBar3Proto(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar3(popr, false)
	dAtA, err := github_com_cosmos_gogoproto_proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &Bar3{}
	if err := github_com_cosmos_gogoproto_proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	littlefuzz := make([]byte, len(dAtA))
	copy(littlefuzz, dAtA)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
	if len(littlefuzz) > 0 {
		fuzzamount := 100
		for i := 0; i < fuzzamount; i++ {
			littlefuzz[popr.Intn(len(littlefuzz))] = byte(popr.Intn(256))
			littlefuzz = append(littlefuzz, byte(popr.Intn(256)))
		}
		// shouldn't panic
		_ = github_com_cosmos_gogoproto_proto.Unmarshal(littlefuzz, msg)
	}
}

func TestBar3MarshalTo(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar3(popr, false)
	size := p.Size()
	dAtA := make([]byte, size)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(dAtA)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &Bar3{}
	if err := github_com_cosmos_gogoproto_proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestBar4Proto(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar4(popr, false)
	dAtA, err := github_com_cosmos_gogoproto_proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &Bar4{}
	if err := github_com_cosmos_gogoproto_proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	littlefuzz := make([]byte, len(dAtA))
	copy(littlefuzz, dAtA)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
	if len(littlefuzz) > 0 {
		fuzzamount := 100
		for i := 0; i < fuzzamount; i++ {
			littlefuzz[popr.Intn(len(littlefuzz))] = byte(popr.Intn(256))
			littlefuzz = append(littlefuzz, byte(popr.Intn(256)))
		}
		// shouldn't panic
		_ = github_com_cosmos_gogoproto_proto.Unmarshal(littlefuzz, msg)
	}
}

func TestBar4MarshalTo(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar4(popr, false)
	size := p.Size()
	dAtA := make([]byte, size)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(dAtA)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &Bar4{}
	if err := github_com_cosmos_gogoproto_proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestBar5Proto(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar5(popr, false)
	dAtA, err := github_com_cosmos_gogoproto_proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &Bar5{}
	if err := github_com_cosmos_gogoproto_proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	littlefuzz := make([]byte, len(dAtA))
	copy(littlefuzz, dAtA)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
	if len(littlefuzz) > 0 {
		fuzzamount := 100
		for i := 0; i < fuzzamount; i++ {
			littlefuzz[popr.Intn(len(littlefuzz))] = byte(popr.Intn(256))
			littlefuzz = append(littlefuzz, byte(popr.Intn(256)))
		}
		// shouldn't panic
		_ = github_com_cosmos_gogoproto_proto.Unmarshal(littlefuzz, msg)
	}
}

func TestBar5MarshalTo(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar5(popr, false)
	size := p.Size()
	dAtA := make([]byte, size)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(dAtA)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &Bar5{}
	if err := github_com_cosmos_gogoproto_proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestBar7Proto(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar7(popr, false)
	dAtA, err := github_com_cosmos_gogoproto_proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &Bar7{}
	if err := github_com_cosmos_gogoproto_proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	littlefuzz := make([]byte, len(dAtA))
	copy(littlefuzz, dAtA)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
	if len(littlefuzz) > 0 {
		fuzzamount := 100
		for i := 0; i < fuzzamount; i++ {
			littlefuzz[popr.Intn(len(littlefuzz))] = byte(popr.Intn(256))
			littlefuzz = append(littlefuzz, byte(popr.Intn(256)))
		}
		// shouldn't panic
		_ = github_com_cosmos_gogoproto_proto.Unmarshal(littlefuzz, msg)
	}
}

func TestBar7MarshalTo(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar7(popr, false)
	size := p.Size()
	dAtA := make([]byte, size)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(dAtA)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &Bar7{}
	if err := github_com_cosmos_gogoproto_proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestBar8Proto(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar8(popr, false)
	dAtA, err := github_com_cosmos_gogoproto_proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &Bar8{}
	if err := github_com_cosmos_gogoproto_proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	littlefuzz := make([]byte, len(dAtA))
	copy(littlefuzz, dAtA)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
	if len(littlefuzz) > 0 {
		fuzzamount := 100
		for i := 0; i < fuzzamount; i++ {
			littlefuzz[popr.Intn(len(littlefuzz))] = byte(popr.Intn(256))
			littlefuzz = append(littlefuzz, byte(popr.Intn(256)))
		}
		// shouldn't panic
		_ = github_com_cosmos_gogoproto_proto.Unmarshal(littlefuzz, msg)
	}
}

func TestBar8MarshalTo(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar8(popr, false)
	size := p.Size()
	dAtA := make([]byte, size)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(dAtA)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &Bar8{}
	if err := github_com_cosmos_gogoproto_proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestBar9Proto(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar9(popr, false)
	dAtA, err := github_com_cosmos_gogoproto_proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &Bar9{}
	if err := github_com_cosmos_gogoproto_proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	littlefuzz := make([]byte, len(dAtA))
	copy(littlefuzz, dAtA)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
	if len(littlefuzz) > 0 {
		fuzzamount := 100
		for i := 0; i < fuzzamount; i++ {
			littlefuzz[popr.Intn(len(littlefuzz))] = byte(popr.Intn(256))
			littlefuzz = append(littlefuzz, byte(popr.Intn(256)))
		}
		// shouldn't panic
		_ = github_com_cosmos_gogoproto_proto.Unmarshal(littlefuzz, msg)
	}
}

func TestBar9MarshalTo(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar9(popr, false)
	size := p.Size()
	dAtA := make([]byte, size)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(dAtA)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &Bar9{}
	if err := github_com_cosmos_gogoproto_proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestFoo5JSON(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedFoo5(popr, true)
	marshaler := github_com_cosmos_gogoproto_jsonpb.Marshaler{}
	jsondata, err := marshaler.MarshalToString(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &Foo5{}
	err = github_com_cosmos_gogoproto_jsonpb.UnmarshalString(jsondata, msg)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Json Equal %#v", seed, msg, p)
	}
}
func TestBar1JSON(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar1(popr, true)
	marshaler := github_com_cosmos_gogoproto_jsonpb.Marshaler{}
	jsondata, err := marshaler.MarshalToString(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &Bar1{}
	err = github_com_cosmos_gogoproto_jsonpb.UnmarshalString(jsondata, msg)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Json Equal %#v", seed, msg, p)
	}
}
func TestBar2JSON(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar2(popr, true)
	marshaler := github_com_cosmos_gogoproto_jsonpb.Marshaler{}
	jsondata, err := marshaler.MarshalToString(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &Bar2{}
	err = github_com_cosmos_gogoproto_jsonpb.UnmarshalString(jsondata, msg)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Json Equal %#v", seed, msg, p)
	}
}
func TestBar3JSON(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar3(popr, true)
	marshaler := github_com_cosmos_gogoproto_jsonpb.Marshaler{}
	jsondata, err := marshaler.MarshalToString(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &Bar3{}
	err = github_com_cosmos_gogoproto_jsonpb.UnmarshalString(jsondata, msg)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Json Equal %#v", seed, msg, p)
	}
}
func TestBar4JSON(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar4(popr, true)
	marshaler := github_com_cosmos_gogoproto_jsonpb.Marshaler{}
	jsondata, err := marshaler.MarshalToString(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &Bar4{}
	err = github_com_cosmos_gogoproto_jsonpb.UnmarshalString(jsondata, msg)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Json Equal %#v", seed, msg, p)
	}
}
func TestBar5JSON(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar5(popr, true)
	marshaler := github_com_cosmos_gogoproto_jsonpb.Marshaler{}
	jsondata, err := marshaler.MarshalToString(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &Bar5{}
	err = github_com_cosmos_gogoproto_jsonpb.UnmarshalString(jsondata, msg)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Json Equal %#v", seed, msg, p)
	}
}
func TestBar7JSON(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar7(popr, true)
	marshaler := github_com_cosmos_gogoproto_jsonpb.Marshaler{}
	jsondata, err := marshaler.MarshalToString(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &Bar7{}
	err = github_com_cosmos_gogoproto_jsonpb.UnmarshalString(jsondata, msg)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Json Equal %#v", seed, msg, p)
	}
}
func TestBar8JSON(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar8(popr, true)
	marshaler := github_com_cosmos_gogoproto_jsonpb.Marshaler{}
	jsondata, err := marshaler.MarshalToString(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &Bar8{}
	err = github_com_cosmos_gogoproto_jsonpb.UnmarshalString(jsondata, msg)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Json Equal %#v", seed, msg, p)
	}
}
func TestBar9JSON(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar9(popr, true)
	marshaler := github_com_cosmos_gogoproto_jsonpb.Marshaler{}
	jsondata, err := marshaler.MarshalToString(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &Bar9{}
	err = github_com_cosmos_gogoproto_jsonpb.UnmarshalString(jsondata, msg)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Json Equal %#v", seed, msg, p)
	}
}
func TestFoo5ProtoText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedFoo5(popr, true)
	dAtA := github_com_cosmos_gogoproto_proto.MarshalTextString(p)
	msg := &Foo5{}
	if err := github_com_cosmos_gogoproto_proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestFoo5ProtoCompactText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedFoo5(popr, true)
	dAtA := github_com_cosmos_gogoproto_proto.CompactTextString(p)
	msg := &Foo5{}
	if err := github_com_cosmos_gogoproto_proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestBar1ProtoText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar1(popr, true)
	dAtA := github_com_cosmos_gogoproto_proto.MarshalTextString(p)
	msg := &Bar1{}
	if err := github_com_cosmos_gogoproto_proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestBar1ProtoCompactText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar1(popr, true)
	dAtA := github_com_cosmos_gogoproto_proto.CompactTextString(p)
	msg := &Bar1{}
	if err := github_com_cosmos_gogoproto_proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestBar2ProtoText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar2(popr, true)
	dAtA := github_com_cosmos_gogoproto_proto.MarshalTextString(p)
	msg := &Bar2{}
	if err := github_com_cosmos_gogoproto_proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestBar2ProtoCompactText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar2(popr, true)
	dAtA := github_com_cosmos_gogoproto_proto.CompactTextString(p)
	msg := &Bar2{}
	if err := github_com_cosmos_gogoproto_proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestBar3ProtoText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar3(popr, true)
	dAtA := github_com_cosmos_gogoproto_proto.MarshalTextString(p)
	msg := &Bar3{}
	if err := github_com_cosmos_gogoproto_proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestBar3ProtoCompactText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar3(popr, true)
	dAtA := github_com_cosmos_gogoproto_proto.CompactTextString(p)
	msg := &Bar3{}
	if err := github_com_cosmos_gogoproto_proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestBar4ProtoText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar4(popr, true)
	dAtA := github_com_cosmos_gogoproto_proto.MarshalTextString(p)
	msg := &Bar4{}
	if err := github_com_cosmos_gogoproto_proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestBar4ProtoCompactText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar4(popr, true)
	dAtA := github_com_cosmos_gogoproto_proto.CompactTextString(p)
	msg := &Bar4{}
	if err := github_com_cosmos_gogoproto_proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestBar5ProtoText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar5(popr, true)
	dAtA := github_com_cosmos_gogoproto_proto.MarshalTextString(p)
	msg := &Bar5{}
	if err := github_com_cosmos_gogoproto_proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestBar5ProtoCompactText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar5(popr, true)
	dAtA := github_com_cosmos_gogoproto_proto.CompactTextString(p)
	msg := &Bar5{}
	if err := github_com_cosmos_gogoproto_proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestBar7ProtoText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar7(popr, true)
	dAtA := github_com_cosmos_gogoproto_proto.MarshalTextString(p)
	msg := &Bar7{}
	if err := github_com_cosmos_gogoproto_proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestBar7ProtoCompactText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar7(popr, true)
	dAtA := github_com_cosmos_gogoproto_proto.CompactTextString(p)
	msg := &Bar7{}
	if err := github_com_cosmos_gogoproto_proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestBar8ProtoText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar8(popr, true)
	dAtA := github_com_cosmos_gogoproto_proto.MarshalTextString(p)
	msg := &Bar8{}
	if err := github_com_cosmos_gogoproto_proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestBar8ProtoCompactText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar8(popr, true)
	dAtA := github_com_cosmos_gogoproto_proto.CompactTextString(p)
	msg := &Bar8{}
	if err := github_com_cosmos_gogoproto_proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestBar9ProtoText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar9(popr, true)
	dAtA := github_com_cosmos_gogoproto_proto.MarshalTextString(p)
	msg := &Bar9{}
	if err := github_com_cosmos_gogoproto_proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestBar9ProtoCompactText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar9(popr, true)
	dAtA := github_com_cosmos_gogoproto_proto.CompactTextString(p)
	msg := &Bar9{}
	if err := github_com_cosmos_gogoproto_proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestFoo5Size(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedFoo5(popr, true)
	size2 := github_com_cosmos_gogoproto_proto.Size(p)
	dAtA, err := github_com_cosmos_gogoproto_proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	size := p.Size()
	if len(dAtA) != size {
		t.Errorf("seed = %d, size %v != marshalled size %v", seed, size, len(dAtA))
	}
	if size2 != size {
		t.Errorf("seed = %d, size %v != before marshal proto.Size %v", seed, size, size2)
	}
	size3 := github_com_cosmos_gogoproto_proto.Size(p)
	if size3 != size {
		t.Errorf("seed = %d, size %v != after marshal proto.Size %v", seed, size, size3)
	}
}

func TestBar1Size(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar1(popr, true)
	size2 := github_com_cosmos_gogoproto_proto.Size(p)
	dAtA, err := github_com_cosmos_gogoproto_proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	size := p.Size()
	if len(dAtA) != size {
		t.Errorf("seed = %d, size %v != marshalled size %v", seed, size, len(dAtA))
	}
	if size2 != size {
		t.Errorf("seed = %d, size %v != before marshal proto.Size %v", seed, size, size2)
	}
	size3 := github_com_cosmos_gogoproto_proto.Size(p)
	if size3 != size {
		t.Errorf("seed = %d, size %v != after marshal proto.Size %v", seed, size, size3)
	}
}

func TestBar2Size(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar2(popr, true)
	size2 := github_com_cosmos_gogoproto_proto.Size(p)
	dAtA, err := github_com_cosmos_gogoproto_proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	size := p.Size()
	if len(dAtA) != size {
		t.Errorf("seed = %d, size %v != marshalled size %v", seed, size, len(dAtA))
	}
	if size2 != size {
		t.Errorf("seed = %d, size %v != before marshal proto.Size %v", seed, size, size2)
	}
	size3 := github_com_cosmos_gogoproto_proto.Size(p)
	if size3 != size {
		t.Errorf("seed = %d, size %v != after marshal proto.Size %v", seed, size, size3)
	}
}

func TestBar3Size(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar3(popr, true)
	size2 := github_com_cosmos_gogoproto_proto.Size(p)
	dAtA, err := github_com_cosmos_gogoproto_proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	size := p.Size()
	if len(dAtA) != size {
		t.Errorf("seed = %d, size %v != marshalled size %v", seed, size, len(dAtA))
	}
	if size2 != size {
		t.Errorf("seed = %d, size %v != before marshal proto.Size %v", seed, size, size2)
	}
	size3 := github_com_cosmos_gogoproto_proto.Size(p)
	if size3 != size {
		t.Errorf("seed = %d, size %v != after marshal proto.Size %v", seed, size, size3)
	}
}

func TestBar4Size(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar4(popr, true)
	size2 := github_com_cosmos_gogoproto_proto.Size(p)
	dAtA, err := github_com_cosmos_gogoproto_proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	size := p.Size()
	if len(dAtA) != size {
		t.Errorf("seed = %d, size %v != marshalled size %v", seed, size, len(dAtA))
	}
	if size2 != size {
		t.Errorf("seed = %d, size %v != before marshal proto.Size %v", seed, size, size2)
	}
	size3 := github_com_cosmos_gogoproto_proto.Size(p)
	if size3 != size {
		t.Errorf("seed = %d, size %v != after marshal proto.Size %v", seed, size, size3)
	}
}

func TestBar5Size(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar5(popr, true)
	size2 := github_com_cosmos_gogoproto_proto.Size(p)
	dAtA, err := github_com_cosmos_gogoproto_proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	size := p.Size()
	if len(dAtA) != size {
		t.Errorf("seed = %d, size %v != marshalled size %v", seed, size, len(dAtA))
	}
	if size2 != size {
		t.Errorf("seed = %d, size %v != before marshal proto.Size %v", seed, size, size2)
	}
	size3 := github_com_cosmos_gogoproto_proto.Size(p)
	if size3 != size {
		t.Errorf("seed = %d, size %v != after marshal proto.Size %v", seed, size, size3)
	}
}

func TestBar7Size(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar7(popr, true)
	size2 := github_com_cosmos_gogoproto_proto.Size(p)
	dAtA, err := github_com_cosmos_gogoproto_proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	size := p.Size()
	if len(dAtA) != size {
		t.Errorf("seed = %d, size %v != marshalled size %v", seed, size, len(dAtA))
	}
	if size2 != size {
		t.Errorf("seed = %d, size %v != before marshal proto.Size %v", seed, size, size2)
	}
	size3 := github_com_cosmos_gogoproto_proto.Size(p)
	if size3 != size {
		t.Errorf("seed = %d, size %v != after marshal proto.Size %v", seed, size, size3)
	}
}

func TestBar8Size(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar8(popr, true)
	size2 := github_com_cosmos_gogoproto_proto.Size(p)
	dAtA, err := github_com_cosmos_gogoproto_proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	size := p.Size()
	if len(dAtA) != size {
		t.Errorf("seed = %d, size %v != marshalled size %v", seed, size, len(dAtA))
	}
	if size2 != size {
		t.Errorf("seed = %d, size %v != before marshal proto.Size %v", seed, size, size2)
	}
	size3 := github_com_cosmos_gogoproto_proto.Size(p)
	if size3 != size {
		t.Errorf("seed = %d, size %v != after marshal proto.Size %v", seed, size, size3)
	}
}

func TestBar9Size(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar9(popr, true)
	size2 := github_com_cosmos_gogoproto_proto.Size(p)
	dAtA, err := github_com_cosmos_gogoproto_proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	size := p.Size()
	if len(dAtA) != size {
		t.Errorf("seed = %d, size %v != marshalled size %v", seed, size, len(dAtA))
	}
	if size2 != size {
		t.Errorf("seed = %d, size %v != before marshal proto.Size %v", seed, size, size2)
	}
	size3 := github_com_cosmos_gogoproto_proto.Size(p)
	if size3 != size {
		t.Errorf("seed = %d, size %v != after marshal proto.Size %v", seed, size, size3)
	}
}

func TestFoo5Stringer(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedFoo5(popr, false)
	s1 := p.String()
	s2 := fmt.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func TestBar1Stringer(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedBar1(popr, false)
	s1 := p.String()
	s2 := fmt.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func TestBar2Stringer(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedBar2(popr, false)
	s1 := p.String()
	s2 := fmt.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func TestBar3Stringer(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedBar3(popr, false)
	s1 := p.String()
	s2 := fmt.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func TestBar4Stringer(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedBar4(popr, false)
	s1 := p.String()
	s2 := fmt.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func TestBar5Stringer(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedBar5(popr, false)
	s1 := p.String()
	s2 := fmt.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func TestBar7Stringer(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedBar7(popr, false)
	s1 := p.String()
	s2 := fmt.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func TestBar8Stringer(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedBar8(popr, false)
	s1 := p.String()
	s2 := fmt.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}

//These tests are generated by github.com/cosmos/gogoproto/plugin/testgen
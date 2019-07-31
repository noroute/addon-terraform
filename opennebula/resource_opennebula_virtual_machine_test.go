package opennebula

import (
	"addon-terraform/opennebula/test"
	"encoding/xml"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/stretchr/testify/assert"
	"testing"
)

func UnsetVmResourceData() *test.StubResourceData {

	// represents "unset" answers to Get/GetOk

	testObj := new(test.StubResourceData)
	testObj.
		On("Get","cpu").
			Return(0.0).
		On("Get","vcpu").
			Return(0).
		On("Get", "memory").
			Return(0).
		On("Get", "context").
			Return(make(map[string]interface{})).
		On("Get","nic").
			Return(&schema.Set{}).
		On("Get", "disk").
			Return(&schema.Set{}).
		On("GetOk","graphics").
			Return(&schema.Set{}, false).
		On("GetOk","os").
			Return(&schema.Set{}, false)

	return testObj
}

func VmResourceDataWithCpuVcpuAndMem() *test.StubResourceData {

	// represents answers to Get/GetOk with Cpu, VCpu and Memory set

	testObj := new(test.StubResourceData)
	testObj.
		On("Get","cpu").
			Return(0.8).
		On("Get","vcpu").
			Return(1).
		On("Get", "memory").
			Return(2048).
		On("Get", "context").
			Return(make(map[string]interface{})).
		On("Get","nic").
			Return(&schema.Set{}).
		On("Get", "disk").
			Return(&schema.Set{}).
		On("GetOk","graphics").
			Return(&schema.Set{}, false).
		On("GetOk","os").
			Return(&schema.Set{}, false)

	return testObj
}

func TestGenerateXmlReturnsNilValuesForUnsetResourceData(t *testing.T) {

	rd := UnsetVmResourceData()
	vmTemplate := &vmTemplate{}

	generatedXml, _ := generateVmXML(rd)
	_ = xml.Unmarshal([]byte(generatedXml), vmTemplate)

	rd.AssertExpectations(t)
	assert.Equal(t, 0.0, vmTemplate.CPU)
	assert.Equal(t, 0, vmTemplate.VCPU)
	assert.Equal(t, 0, vmTemplate.Memory)
}

func TestGenerateXmlReturnsValuesForCpuVcpuAndMemSetInResourceData(t *testing.T) {

	rd := VmResourceDataWithCpuVcpuAndMem()
	vmTemplate := &vmTemplate{}

	generatedXml, _ := generateVmXML(rd)
	_ = xml.Unmarshal([]byte(generatedXml), vmTemplate)

	rd.AssertExpectations(t)
	assert.Equal(t, 0.8, vmTemplate.CPU)
	assert.Equal(t, 1, vmTemplate.VCPU)
	assert.Equal(t, 2048, vmTemplate.Memory)

}

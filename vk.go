// Copyright (c) 2015-2018 Khronos Group
// Copyright (c) 2018 Dominik Honnef

package vk

// #cgo pkg-config: vulkan
//
// #include <stdlib.h>
// #include "vk.h"
import "C"
import (
	"bytes"
	"fmt"
	"math"
	"os"
	"time"
	"unsafe"
)

// OPT(dh): we could replace large chunks of C info struct
// initializers with memcpys of our Go info structs. A lot of the
// time, they're mostly identical, aside from pNext and slices. This
// would replace a whole lot of MOV instructions with single
// runtime.memmove calls.
//
// We can create special structs with copy-range markers as [0]byte
// fields, cast our Go structs to these, and use them for
// straightforward copying. We should be able to code-generate these.

const debug = true

type (
	DeviceSize = uint64
)

const (
	// Vulkan 1.0 version number
	APIVersion10 = uint32(C.VK_API_VERSION_1_0)
	// Vulkan 1.1 version number
	APIVersion11 = uint32(C.VK_API_VERSION_1_1)
)

var vkEnumerateInstanceVersion C.PFN_vkEnumerateInstanceVersion
var vkEnumerateInstanceExtensionProperties C.PFN_vkEnumerateInstanceExtensionProperties
var vkEnumerateInstanceLayerProperties C.PFN_vkEnumerateInstanceLayerProperties
var vkCreateInstance C.PFN_vkCreateInstance

func init() {
	assertSameSize(unsafe.Sizeof(AttachmentDescription{}), C.sizeof_VkAttachmentDescription)
	assertSameSize(unsafe.Sizeof(AttachmentReference{}), C.sizeof_VkAttachmentReference)
	assertSameSize(unsafe.Sizeof(DescriptorSetLayout{}), C.sizeof_VkDescriptorSetLayout)
	assertSameSize(unsafe.Sizeof(Fence{}), C.sizeof_VkFence)
	assertSameSize(unsafe.Sizeof(ImageView{}), C.sizeof_VkImageView)
	assertSameSize(unsafe.Sizeof(MemoryHeap{}), C.sizeof_VkMemoryHeap)
	assertSameSize(unsafe.Sizeof(MemoryRequirements{}), C.sizeof_VkMemoryRequirements)
	assertSameSize(unsafe.Sizeof(MemoryType{}), C.sizeof_VkMemoryType)
	assertSameSize(unsafe.Sizeof(PushConstantRange{}), C.sizeof_VkPushConstantRange)
	assertSameSize(unsafe.Sizeof(Rect2D{}), C.sizeof_VkRect2D)
	assertSameSize(unsafe.Sizeof(Semaphore{}), C.sizeof_VkSemaphore)
	assertSameSize(unsafe.Sizeof(SubpassDependency{}), C.sizeof_VkSubpassDependency)
	assertSameSize(unsafe.Sizeof(VertexInputAttributeDescription{}), C.sizeof_VkVertexInputAttributeDescription)
	assertSameSize(unsafe.Sizeof(VertexInputBindingDescription{}), C.sizeof_VkVertexInputBindingDescription)
	assertSameSize(unsafe.Sizeof(Viewport{}), C.sizeof_VkViewport)
	assertSameSize(unsafe.Sizeof(ComponentMapping{}), C.sizeof_VkComponentMapping)
	assertSameSize(unsafe.Sizeof(ImageSubresourceRange{}), C.sizeof_VkImageSubresourceRange)
	assertSameSize(unsafe.Sizeof(ClearDepthStencilValue{}), C.sizeof_VkClearDepthStencilValue)
	assertSameSize(unsafe.Sizeof(BufferCopy{}), C.sizeof_VkBufferCopy)
	assertSameSize(unsafe.Sizeof(BufferImageCopy{}), C.sizeof_VkBufferImageCopy)
	assertSameSize(unsafe.Sizeof(ImageSubresourceLayers{}), C.sizeof_VkImageSubresourceLayers)
	assertSameSize(unsafe.Sizeof(ImageCopy{}), C.sizeof_VkImageCopy)

	vkEnumerateInstanceVersion =
		C.PFN_vkEnumerateInstanceVersion(mustVkGetInstanceProcAddr(nil, "vkEnumerateInstanceVersion"))
	vkEnumerateInstanceExtensionProperties =
		C.PFN_vkEnumerateInstanceExtensionProperties(mustVkGetInstanceProcAddr(nil, "vkEnumerateInstanceExtensionProperties"))
	vkEnumerateInstanceLayerProperties =
		C.PFN_vkEnumerateInstanceLayerProperties(mustVkGetInstanceProcAddr(nil, "vkEnumerateInstanceLayerProperties"))
	vkCreateInstance =
		C.PFN_vkCreateInstance(mustVkGetInstanceProcAddr(nil, "vkCreateInstance"))
}

// MakeVersion constructs an API version number.
func MakeVersion(major, minor, patch uint32) uint32 {
	return major<<22 | minor<<12 | patch
}

type InstanceCreateInfo struct {
	Extensions []Extension
	// If not nil, this information helps implementations recognize behavior inherent to classes of applications
	ApplicationInfo *ApplicationInfo
	// Names of layers to enable for the created instance
	EnabledLayerNames []string
	// Names of extensions to enable
	EnabledExtensionNames []string
}

type ApplicationInfo struct {
	Extensions []Extension
	// The name of the application
	ApplicationName string
	// The developer-supplied version number of the application
	ApplicationVersion uint32
	// The name of the engine (if any) used to create the application
	EngineName string
	// The developer-supplied version number of the engine used to create the application
	EngineVersion uint32
	// The highest version of Vulkan that the application is designed to use
	APIVersion uint32
}

func CreateInstance(info *InstanceCreateInfo) (*Instance, error) {
	// TODO(dh): support a custom allocator
	ptr := (*C.VkInstanceCreateInfo)(alloc(C.sizeof_VkInstanceCreateInfo))
	ptr.sType = C.VkStructureType(StructureTypeInstanceCreateInfo)
	ptr.pNext = buildChain(info.Extensions)
	defer internalizeChain(info.Extensions, ptr.pNext)
	ptr.enabledLayerCount = C.uint32_t(len(info.EnabledLayerNames))
	ptr.enabledExtensionCount = C.uint32_t(len(info.EnabledExtensionNames))
	ptr.ppEnabledLayerNames = externStrings(info.EnabledLayerNames)
	ptr.ppEnabledExtensionNames = externStrings(info.EnabledExtensionNames)
	defer free(uptr(ptr))
	defer free(uptr(ptr.ppEnabledLayerNames))
	defer free(uptr(ptr.ppEnabledExtensionNames))
	if info.ApplicationInfo != nil {
		ptr.pApplicationInfo = (*C.VkApplicationInfo)(alloc(C.sizeof_VkApplicationInfo))
		ptr.pApplicationInfo.sType = C.VkStructureType(StructureTypeApplicationInfo)
		ptr.pApplicationInfo.pNext = buildChain(info.ApplicationInfo.Extensions)
		ptr.pApplicationInfo.pApplicationName = C.CString(info.ApplicationInfo.ApplicationName)
		ptr.pApplicationInfo.applicationVersion = C.uint32_t(info.ApplicationInfo.ApplicationVersion)
		ptr.pApplicationInfo.pEngineName = C.CString(info.ApplicationInfo.EngineName)
		ptr.pApplicationInfo.engineVersion = C.uint32_t(info.ApplicationInfo.EngineVersion)
		ptr.pApplicationInfo.apiVersion = C.uint32_t(info.ApplicationInfo.APIVersion)
		defer free(uptr(ptr.pApplicationInfo))
		defer free(uptr(ptr.pApplicationInfo.pApplicationName))
		defer free(uptr(ptr.pApplicationInfo.pEngineName))
		defer internalizeChain(info.ApplicationInfo.Extensions, ptr.pApplicationInfo.pNext)
	}

	var instance C.VkInstance
	res := Result(C.domVkCreateInstance(vkCreateInstance, ptr, nil, &instance))
	if res != Success {
		return nil, res
	}

	out := &Instance{hnd: instance}
	out.init()

	return out, nil
}

type Instance struct {
	// VK_DEFINE_HANDLE(VkInstance)
	hnd C.VkInstance
	fps [instanceMaxPFN]C.PFN_vkVoidFunction
}

func (ins *Instance) init() {
	for i, name := range instanceFpNames {
		ins.fps[i] = vkGetInstanceProcAddr(ins.hnd, name)
	}
}

func (ins *Instance) EnumeratePhysicalDevices() ([]*PhysicalDevice, error) {
	count := C.uint32_t(1)
	var devs *C.VkPhysicalDevice
	for {
		devs = (*C.VkPhysicalDevice)(allocn(int(count), C.sizeof_VkPhysicalDevice))
		defer free(uptr(devs))
		res := Result(C.domVkEnumeratePhysicalDevices(ins.fps[vkEnumeratePhysicalDevices], ins.hnd, &count, devs))
		if res != Success && res != Incomplete {
			return nil, res
		}
		if res == Success {
			break
		}
		if res == Incomplete {
			continue
		}
		panic(fmt.Sprintf("unexpected result %s", res))
	}
	var out []*PhysicalDevice
	for _, dev := range (*[math.MaxInt32]C.VkPhysicalDevice)(uptr(devs))[:count] {
		out = append(out, &PhysicalDevice{dev, ins})
	}
	return out, nil
}

type PhysicalDevice struct {
	// VK_DEFINE_HANDLE(VkPhysicalDevice)
	hnd      C.VkPhysicalDevice
	instance *Instance
}

type PhysicalDeviceLimits struct {
	MaxImageDimension1D                             uint32
	MaxImageDimension2D                             uint32
	MaxImageDimension3D                             uint32
	MaxImageDimensionCube                           uint32
	MaxImageArrayLayers                             uint32
	MaxTexelBufferElements                          uint32
	MaxUniformBufferRange                           uint32
	MaxStorageBufferRange                           uint32
	MaxPushConstantsSize                            uint32
	MaxMemoryAllocationCount                        uint32
	MaxSamplerAllocationCount                       uint32
	BufferImageGranularity                          DeviceSize
	SparseAddressSpaceSize                          DeviceSize
	MaxBoundDescriptorSets                          uint32
	MaxPerStageDescriptorSamplers                   uint32
	MaxPerStageDescriptorUniformBuffers             uint32
	MaxPerStageDescriptorStorageBuffers             uint32
	MaxPerStageDescriptorSampledImages              uint32
	MaxPerStageDescriptorStorageImages              uint32
	MaxPerStageDescriptorInputAttachments           uint32
	MaxPerStageResources                            uint32
	MaxDescriptorSetSamplers                        uint32
	MaxDescriptorSetUniformBuffers                  uint32
	MaxDescriptorSetUniformBuffersDynamic           uint32
	MaxDescriptorSetStorageBuffers                  uint32
	MaxDescriptorSetStorageBuffersDynamic           uint32
	MaxDescriptorSetSampledImages                   uint32
	MaxDescriptorSetStorageImages                   uint32
	MaxDescriptorSetInputAttachments                uint32
	MaxVertexInputAttributes                        uint32
	MaxVertexInputBindings                          uint32
	MaxVertexInputAttributeOffset                   uint32
	MaxVertexInputBindingStride                     uint32
	MaxVertexOutputComponents                       uint32
	MaxTessellationGenerationLevel                  uint32
	MaxTessellationPatchSize                        uint32
	MaxTessellationControlPerVertexInputComponents  uint32
	MaxTessellationControlPerVertexOutputComponents uint32
	MaxTessellationControlPerPatchOutputComponents  uint32
	MaxTessellationControlTotalOutputComponents     uint32
	MaxTessellationEvaluationInputComponents        uint32
	MaxTessellationEvaluationOutputComponents       uint32
	MaxGeometryShaderInvocations                    uint32
	MaxGeometryInputComponents                      uint32
	MaxGeometryOutputComponents                     uint32
	MaxGeometryOutputVertices                       uint32
	MaxGeometryTotalOutputComponents                uint32
	MaxFragmentInputComponents                      uint32
	MaxFragmentOutputAttachments                    uint32
	MaxFragmentDualSrcAttachments                   uint32
	MaxFragmentCombinedOutputResources              uint32
	MaxComputeSharedMemorySize                      uint32
	MaxComputeWorkGroupCount                        [3]uint32
	MaxComputeWorkGroupInvocations                  uint32
	MaxComputeWorkGroupSize                         [3]uint32
	SubPixelPrecisionBits                           uint32
	SubTexelPrecisionBits                           uint32
	MipmapPrecisionBits                             uint32
	MaxDrawIndexedIndexValue                        uint32
	MaxDrawIndirectCount                            uint32
	MaxSamplerLodBias                               float32
	MaxSamplerAnisotropy                            float32
	MaxViewports                                    uint32
	MaxViewportDimensions                           [2]uint32
	ViewportBoundsRange                             [2]float32
	ViewportSubPixelBits                            uint32
	MinMemoryMapAlignment                           uintptr
	MinTexelBufferOffsetAlignment                   DeviceSize
	MinUniformBufferOffsetAlignment                 DeviceSize
	MinStorageBufferOffsetAlignment                 DeviceSize
	MinTexelOffset                                  int32
	MaxTexelOffset                                  uint32
	MinTexelGatherOffset                            int32
	MaxTexelGatherOffset                            uint32
	MinInterpolationOffset                          float32
	MaxInterpolationOffset                          float32
	SubPixelInterpolationOffsetBits                 uint32
	MaxFramebufferWidth                             uint32
	MaxFramebufferHeight                            uint32
	MaxFramebufferLayers                            uint32
	FramebufferColorSampleCounts                    SampleCountFlags
	FramebufferDepthSampleCounts                    SampleCountFlags
	FramebufferStencilSampleCounts                  SampleCountFlags
	FramebufferNoAttachmentsSampleCounts            SampleCountFlags
	MaxColorAttachments                             uint32
	SampledImageColorSampleCounts                   SampleCountFlags
	SampledImageIntegerSampleCounts                 SampleCountFlags
	SampledImageDepthSampleCounts                   SampleCountFlags
	SampledImageStencilSampleCounts                 SampleCountFlags
	StorageImageSampleCounts                        SampleCountFlags
	MaxSampleMaskWords                              uint32
	TimestampComputeAndGraphics                     bool
	TimestampPeriod                                 float32
	MaxClipDistances                                uint32
	MaxCullDistances                                uint32
	MaxCombinedClipAndCullDistances                 uint32
	DiscreteQueuePriorities                         uint32
	PointSizeRange                                  [2]float32
	LineWidthRange                                  [2]float32
	PointSizeGranularity                            float32
	LineWidthGranularity                            float32
	StrictLines                                     bool
	StandardSampleLocations                         bool
	OptimalBufferCopyOffsetAlignment                DeviceSize
	OptimalBufferCopyRowPitchAlignment              DeviceSize
	NonCoherentAtomSize                             DeviceSize
}

type PhysicalDeviceSparseProperties struct {
	ResidencyStandard2DBlockShape            bool
	ResidencyStandard2DMultisampleBlockShape bool
	ResidencyStandard3DBlockShape            bool
	ResidencyAlignedMipSize                  bool
	ResidencyNonResidentStrict               bool
}

type PhysicalDeviceProperties struct {
	// The version of Vulkan supported by the device.
	APIVersion uint32
	// The vendor-specified version of the driver.
	DriverVersion uint32
	// A unique identifier for the vendor of the physical device.
	VendorID uint32
	// a unique identifier for the physical device among devices available from the vendor.
	DeviceID   uint32
	DeviceType PhysicalDeviceType
	DeviceName string
	// A universally unique identifier for the device.
	PipelineCacheUUID []byte
	// Device-specific limits of the physical device.
	Limits           PhysicalDeviceLimits
	SparseProperties PhysicalDeviceSparseProperties
}

func (props *PhysicalDeviceProperties) internalize(cprops *C.VkPhysicalDeviceProperties) {
	*props = PhysicalDeviceProperties{
		APIVersion:        uint32(cprops.apiVersion),
		DriverVersion:     uint32(cprops.driverVersion),
		VendorID:          uint32(cprops.vendorID),
		DeviceID:          uint32(cprops.deviceID),
		DeviceType:        PhysicalDeviceType(cprops.deviceType),
		DeviceName:        C.GoString(&cprops.deviceName[0]),
		PipelineCacheUUID: (*[C.VK_UUID_SIZE]byte)(uptr(&cprops.pipelineCacheUUID))[:],
		Limits: PhysicalDeviceLimits{
			MaxImageDimension1D:                             uint32(cprops.limits.maxImageDimension1D),
			MaxImageDimension2D:                             uint32(cprops.limits.maxImageDimension2D),
			MaxImageDimension3D:                             uint32(cprops.limits.maxImageDimension3D),
			MaxImageDimensionCube:                           uint32(cprops.limits.maxImageDimensionCube),
			MaxImageArrayLayers:                             uint32(cprops.limits.maxImageArrayLayers),
			MaxTexelBufferElements:                          uint32(cprops.limits.maxTexelBufferElements),
			MaxUniformBufferRange:                           uint32(cprops.limits.maxUniformBufferRange),
			MaxStorageBufferRange:                           uint32(cprops.limits.maxStorageBufferRange),
			MaxPushConstantsSize:                            uint32(cprops.limits.maxPushConstantsSize),
			MaxMemoryAllocationCount:                        uint32(cprops.limits.maxMemoryAllocationCount),
			MaxSamplerAllocationCount:                       uint32(cprops.limits.maxSamplerAllocationCount),
			BufferImageGranularity:                          DeviceSize(cprops.limits.bufferImageGranularity),
			SparseAddressSpaceSize:                          DeviceSize(cprops.limits.sparseAddressSpaceSize),
			MaxBoundDescriptorSets:                          uint32(cprops.limits.maxBoundDescriptorSets),
			MaxPerStageDescriptorSamplers:                   uint32(cprops.limits.maxPerStageDescriptorSamplers),
			MaxPerStageDescriptorUniformBuffers:             uint32(cprops.limits.maxPerStageDescriptorUniformBuffers),
			MaxPerStageDescriptorStorageBuffers:             uint32(cprops.limits.maxPerStageDescriptorStorageBuffers),
			MaxPerStageDescriptorSampledImages:              uint32(cprops.limits.maxPerStageDescriptorSampledImages),
			MaxPerStageDescriptorStorageImages:              uint32(cprops.limits.maxPerStageDescriptorStorageImages),
			MaxPerStageDescriptorInputAttachments:           uint32(cprops.limits.maxPerStageDescriptorInputAttachments),
			MaxPerStageResources:                            uint32(cprops.limits.maxPerStageResources),
			MaxDescriptorSetSamplers:                        uint32(cprops.limits.maxDescriptorSetSamplers),
			MaxDescriptorSetUniformBuffers:                  uint32(cprops.limits.maxDescriptorSetUniformBuffers),
			MaxDescriptorSetUniformBuffersDynamic:           uint32(cprops.limits.maxDescriptorSetUniformBuffersDynamic),
			MaxDescriptorSetStorageBuffers:                  uint32(cprops.limits.maxDescriptorSetStorageBuffers),
			MaxDescriptorSetStorageBuffersDynamic:           uint32(cprops.limits.maxDescriptorSetStorageBuffersDynamic),
			MaxDescriptorSetSampledImages:                   uint32(cprops.limits.maxDescriptorSetSampledImages),
			MaxDescriptorSetStorageImages:                   uint32(cprops.limits.maxDescriptorSetStorageImages),
			MaxDescriptorSetInputAttachments:                uint32(cprops.limits.maxDescriptorSetInputAttachments),
			MaxVertexInputAttributes:                        uint32(cprops.limits.maxVertexInputAttributes),
			MaxVertexInputBindings:                          uint32(cprops.limits.maxVertexInputBindings),
			MaxVertexInputAttributeOffset:                   uint32(cprops.limits.maxVertexInputAttributeOffset),
			MaxVertexInputBindingStride:                     uint32(cprops.limits.maxVertexInputBindingStride),
			MaxVertexOutputComponents:                       uint32(cprops.limits.maxVertexOutputComponents),
			MaxTessellationGenerationLevel:                  uint32(cprops.limits.maxTessellationGenerationLevel),
			MaxTessellationPatchSize:                        uint32(cprops.limits.maxTessellationPatchSize),
			MaxTessellationControlPerVertexInputComponents:  uint32(cprops.limits.maxTessellationControlPerVertexInputComponents),
			MaxTessellationControlPerVertexOutputComponents: uint32(cprops.limits.maxTessellationControlPerVertexOutputComponents),
			MaxTessellationControlPerPatchOutputComponents:  uint32(cprops.limits.maxTessellationControlPerPatchOutputComponents),
			MaxTessellationControlTotalOutputComponents:     uint32(cprops.limits.maxTessellationControlTotalOutputComponents),
			MaxTessellationEvaluationInputComponents:        uint32(cprops.limits.maxTessellationEvaluationInputComponents),
			MaxTessellationEvaluationOutputComponents:       uint32(cprops.limits.maxTessellationEvaluationOutputComponents),
			MaxGeometryShaderInvocations:                    uint32(cprops.limits.maxGeometryShaderInvocations),
			MaxGeometryInputComponents:                      uint32(cprops.limits.maxGeometryInputComponents),
			MaxGeometryOutputComponents:                     uint32(cprops.limits.maxGeometryOutputComponents),
			MaxGeometryOutputVertices:                       uint32(cprops.limits.maxGeometryOutputVertices),
			MaxGeometryTotalOutputComponents:                uint32(cprops.limits.maxGeometryTotalOutputComponents),
			MaxFragmentInputComponents:                      uint32(cprops.limits.maxFragmentInputComponents),
			MaxFragmentOutputAttachments:                    uint32(cprops.limits.maxFragmentOutputAttachments),
			MaxFragmentDualSrcAttachments:                   uint32(cprops.limits.maxFragmentDualSrcAttachments),
			MaxFragmentCombinedOutputResources:              uint32(cprops.limits.maxFragmentCombinedOutputResources),
			MaxComputeSharedMemorySize:                      uint32(cprops.limits.maxComputeSharedMemorySize),
			MaxComputeWorkGroupCount: [3]uint32{
				uint32(cprops.limits.maxComputeWorkGroupCount[0]),
				uint32(cprops.limits.maxComputeWorkGroupCount[1]),
				uint32(cprops.limits.maxComputeWorkGroupCount[2]),
			},
			MaxComputeWorkGroupInvocations: uint32(cprops.limits.maxComputeWorkGroupInvocations),
			MaxComputeWorkGroupSize: [3]uint32{
				uint32(cprops.limits.maxComputeWorkGroupSize[0]),
				uint32(cprops.limits.maxComputeWorkGroupSize[1]),
				uint32(cprops.limits.maxComputeWorkGroupSize[2]),
			},
			SubPixelPrecisionBits:    uint32(cprops.limits.subPixelPrecisionBits),
			SubTexelPrecisionBits:    uint32(cprops.limits.subTexelPrecisionBits),
			MipmapPrecisionBits:      uint32(cprops.limits.mipmapPrecisionBits),
			MaxDrawIndexedIndexValue: uint32(cprops.limits.maxDrawIndexedIndexValue),
			MaxDrawIndirectCount:     uint32(cprops.limits.maxDrawIndirectCount),
			MaxSamplerLodBias:        float32(cprops.limits.maxSamplerLodBias),
			MaxSamplerAnisotropy:     float32(cprops.limits.maxSamplerAnisotropy),
			MaxViewports:             uint32(cprops.limits.maxViewports),
			MaxViewportDimensions: [2]uint32{
				uint32(cprops.limits.maxViewportDimensions[0]),
				uint32(cprops.limits.maxViewportDimensions[1]),
			},
			ViewportBoundsRange: [2]float32{
				float32(cprops.limits.viewportBoundsRange[0]),
				float32(cprops.limits.viewportBoundsRange[1]),
			},
			ViewportSubPixelBits:                 uint32(cprops.limits.viewportSubPixelBits),
			MinMemoryMapAlignment:                uintptr(cprops.limits.minMemoryMapAlignment),
			MinTexelBufferOffsetAlignment:        DeviceSize(cprops.limits.minTexelBufferOffsetAlignment),
			MinUniformBufferOffsetAlignment:      DeviceSize(cprops.limits.minUniformBufferOffsetAlignment),
			MinStorageBufferOffsetAlignment:      DeviceSize(cprops.limits.minStorageBufferOffsetAlignment),
			MinTexelOffset:                       int32(cprops.limits.minTexelOffset),
			MaxTexelOffset:                       uint32(cprops.limits.maxTexelOffset),
			MinTexelGatherOffset:                 int32(cprops.limits.minTexelGatherOffset),
			MaxTexelGatherOffset:                 uint32(cprops.limits.maxTexelGatherOffset),
			MinInterpolationOffset:               float32(cprops.limits.minInterpolationOffset),
			MaxInterpolationOffset:               float32(cprops.limits.maxInterpolationOffset),
			SubPixelInterpolationOffsetBits:      uint32(cprops.limits.subPixelInterpolationOffsetBits),
			MaxFramebufferWidth:                  uint32(cprops.limits.maxFramebufferWidth),
			MaxFramebufferHeight:                 uint32(cprops.limits.maxFramebufferHeight),
			MaxFramebufferLayers:                 uint32(cprops.limits.maxFramebufferLayers),
			FramebufferColorSampleCounts:         SampleCountFlags(cprops.limits.framebufferColorSampleCounts),
			FramebufferDepthSampleCounts:         SampleCountFlags(cprops.limits.framebufferDepthSampleCounts),
			FramebufferStencilSampleCounts:       SampleCountFlags(cprops.limits.framebufferStencilSampleCounts),
			FramebufferNoAttachmentsSampleCounts: SampleCountFlags(cprops.limits.framebufferNoAttachmentsSampleCounts),
			MaxColorAttachments:                  uint32(cprops.limits.maxColorAttachments),
			SampledImageColorSampleCounts:        SampleCountFlags(cprops.limits.sampledImageColorSampleCounts),
			SampledImageIntegerSampleCounts:      SampleCountFlags(cprops.limits.sampledImageIntegerSampleCounts),
			SampledImageDepthSampleCounts:        SampleCountFlags(cprops.limits.sampledImageDepthSampleCounts),
			SampledImageStencilSampleCounts:      SampleCountFlags(cprops.limits.sampledImageStencilSampleCounts),
			StorageImageSampleCounts:             SampleCountFlags(cprops.limits.storageImageSampleCounts),
			MaxSampleMaskWords:                   uint32(cprops.limits.maxSampleMaskWords),
			TimestampComputeAndGraphics:          cprops.limits.timestampComputeAndGraphics == C.VK_TRUE,
			TimestampPeriod:                      float32(cprops.limits.timestampPeriod),
			MaxClipDistances:                     uint32(cprops.limits.maxClipDistances),
			MaxCullDistances:                     uint32(cprops.limits.maxCullDistances),
			MaxCombinedClipAndCullDistances:      uint32(cprops.limits.maxCombinedClipAndCullDistances),
			DiscreteQueuePriorities:              uint32(cprops.limits.discreteQueuePriorities),
			PointSizeRange: [2]float32{
				float32(cprops.limits.pointSizeRange[0]),
				float32(cprops.limits.pointSizeRange[1]),
			},
			LineWidthRange: [2]float32{
				float32(cprops.limits.lineWidthRange[0]),
				float32(cprops.limits.lineWidthRange[1]),
			},
			PointSizeGranularity:               float32(cprops.limits.pointSizeGranularity),
			LineWidthGranularity:               float32(cprops.limits.lineWidthGranularity),
			StrictLines:                        cprops.limits.strictLines == C.VK_TRUE,
			StandardSampleLocations:            cprops.limits.standardSampleLocations == C.VK_TRUE,
			OptimalBufferCopyOffsetAlignment:   DeviceSize(cprops.limits.optimalBufferCopyOffsetAlignment),
			OptimalBufferCopyRowPitchAlignment: DeviceSize(cprops.limits.optimalBufferCopyRowPitchAlignment),
			NonCoherentAtomSize:                DeviceSize(cprops.limits.nonCoherentAtomSize),
		},
		SparseProperties: PhysicalDeviceSparseProperties{
			ResidencyStandard2DBlockShape:            cprops.sparseProperties.residencyStandard2DBlockShape == C.VK_TRUE,
			ResidencyStandard2DMultisampleBlockShape: cprops.sparseProperties.residencyStandard2DMultisampleBlockShape == C.VK_TRUE,
			ResidencyStandard3DBlockShape:            cprops.sparseProperties.residencyStandard3DBlockShape == C.VK_TRUE,
			ResidencyAlignedMipSize:                  cprops.sparseProperties.residencyAlignedMipSize == C.VK_TRUE,
			ResidencyNonResidentStrict:               cprops.sparseProperties.residencyNonResidentStrict == C.VK_TRUE,
		},
	}
}

func (dev *PhysicalDevice) Properties2(extensions []Extension) *PhysicalDeviceProperties {
	cprops := (*C.VkPhysicalDeviceProperties2)(alloc(C.sizeof_VkPhysicalDeviceProperties2))
	cprops.sType = C.VkStructureType(StructureTypePhysicalDeviceProperties2)
	cprops.pNext = buildChain(extensions)
	C.domVkGetPhysicalDeviceProperties2(dev.instance.fps[vkGetPhysicalDeviceProperties2], dev.hnd, cprops)
	internalizeChain(extensions, cprops.pNext)
	var out PhysicalDeviceProperties
	out.internalize(&cprops.properties)
	return &out
}

// Properties returns general properties of the physical device.
func (dev *PhysicalDevice) Properties() *PhysicalDeviceProperties {
	var props C.VkPhysicalDeviceProperties
	C.domVkGetPhysicalDeviceProperties(dev.instance.fps[vkGetPhysicalDeviceProperties], dev.hnd, &props)
	var out PhysicalDeviceProperties
	out.internalize(&props)
	return &out
}

type MemoryType struct {
	PropertyFlags MemoryPropertyFlags
	HeapIndex     uint32

	// must be kept identical to C struct
}

type MemoryHeap struct {
	Size  DeviceSize
	Flags MemoryHeapFlags

	// must be kept identical to C struct
}

type PhysicalDeviceMemoryProperties struct {
	Types []MemoryType
	Heaps []MemoryHeap
}

func (dev *PhysicalDevice) MemoryProperties() PhysicalDeviceMemoryProperties {
	var props C.VkPhysicalDeviceMemoryProperties
	C.domVkGetPhysicalDeviceMemoryProperties(dev.instance.fps[vkGetPhysicalDeviceMemoryProperties], dev.hnd, &props)

	return PhysicalDeviceMemoryProperties{
		Types: (*[C.VK_MAX_MEMORY_TYPES]MemoryType)(uptr(&props.memoryTypes))[:props.memoryTypeCount],
		Heaps: (*[C.VK_MAX_MEMORY_TYPES]MemoryHeap)(uptr(&props.memoryHeaps))[:props.memoryHeapCount],
	}
}

type ExtensionProperties struct {
	Name        string
	SpecVersion uint32
}

func (dev *PhysicalDevice) ExtensionProperties(layer string) ([]ExtensionProperties, error) {
	var count C.uint32_t
	var cLayer *C.char
	if layer != "" {
		cLayer := C.CString(layer)
		defer free(uptr(cLayer))
	}
	res := Result(C.domVkEnumerateDeviceExtensionProperties(dev.instance.fps[vkEnumerateDeviceExtensionProperties], dev.hnd, cLayer, &count, nil))
	if res != Success {
		return nil, res
	}
	properties := make([]C.VkExtensionProperties, count)
	res = Result(C.domVkEnumerateDeviceExtensionProperties(dev.instance.fps[vkEnumerateDeviceExtensionProperties], dev.hnd, cLayer, &count, (*C.VkExtensionProperties)(slice2ptr(uptr(&properties)))))
	if res != Success {
		return nil, res
	}
	out := make([]ExtensionProperties, count)

	for i, s := range properties {
		name := (*[256]byte)(uptr(&s.extensionName))[:]
		idx := bytes.IndexByte(name, 0)
		out[i] = ExtensionProperties{
			Name:        string(name[:idx]),
			SpecVersion: uint32(s.specVersion),
		}
	}
	return out, nil
}

type PhysicalDeviceFeatures struct {
	RobustBufferAccess                      bool
	FullDrawIndexUint32                     bool
	ImageCubeArray                          bool
	IndependentBlend                        bool
	GeometryShader                          bool
	TessellationShader                      bool
	SampleRateShading                       bool
	DualSrcBlend                            bool
	LogicOp                                 bool
	MultiDrawIndirect                       bool
	DrawIndirectFirstInstance               bool
	DepthClamp                              bool
	DepthBiasClamp                          bool
	FillModeNonSolid                        bool
	DepthBounds                             bool
	WideLines                               bool
	LargePoints                             bool
	AlphaToOne                              bool
	MultiViewport                           bool
	SamplerAnisotropy                       bool
	TextureCompressionETC2                  bool
	TextureCompressionASTC_LDR              bool
	TextureCompressionBC                    bool
	OcclusionQueryPrecise                   bool
	PipelineStatisticsQuery                 bool
	VertexPipelineStoresAndAtomics          bool
	FragmentStoresAndAtomics                bool
	ShaderTessellationAndGeometryPointSize  bool
	ShaderImageGatherExtended               bool
	ShaderStorageImageExtendedFormats       bool
	ShaderStorageImageMultisample           bool
	ShaderStorageImageReadWithoutFormat     bool
	ShaderStorageImageWriteWithoutFormat    bool
	ShaderUniformBufferArrayDynamicIndexing bool
	ShaderSampledImageArrayDynamicIndexing  bool
	ShaderStorageBufferArrayDynamicIndexing bool
	ShaderStorageImageArrayDynamicIndexing  bool
	ShaderClipDistance                      bool
	ShaderCullDistance                      bool
	ShaderFloat64                           bool
	ShaderInt64                             bool
	ShaderInt16                             bool
	ShaderResourceResidency                 bool
	ShaderResourceMinLod                    bool
	SparseBinding                           bool
	SparseResidencyBuffer                   bool
	SparseResidencyImage2D                  bool
	SparseResidencyImage3D                  bool
	SparseResidency2Samples                 bool
	SparseResidency4Samples                 bool
	SparseResidency8Samples                 bool
	SparseResidency16Samples                bool
	SparseResidencyAliased                  bool
	VariableMultisampleRate                 bool
	InheritedQueries                        bool
}

func (dev *PhysicalDevice) Features() *PhysicalDeviceFeatures {
	var features C.VkPhysicalDeviceFeatures
	C.domVkGetPhysicalDeviceFeatures(dev.instance.fps[vkGetPhysicalDeviceFeatures], dev.hnd, &features)

	return &PhysicalDeviceFeatures{
		RobustBufferAccess:                      features.robustBufferAccess == C.VK_TRUE,
		FullDrawIndexUint32:                     features.fullDrawIndexUint32 == C.VK_TRUE,
		ImageCubeArray:                          features.imageCubeArray == C.VK_TRUE,
		IndependentBlend:                        features.independentBlend == C.VK_TRUE,
		GeometryShader:                          features.geometryShader == C.VK_TRUE,
		TessellationShader:                      features.tessellationShader == C.VK_TRUE,
		SampleRateShading:                       features.sampleRateShading == C.VK_TRUE,
		DualSrcBlend:                            features.dualSrcBlend == C.VK_TRUE,
		LogicOp:                                 features.logicOp == C.VK_TRUE,
		MultiDrawIndirect:                       features.multiDrawIndirect == C.VK_TRUE,
		DrawIndirectFirstInstance:               features.drawIndirectFirstInstance == C.VK_TRUE,
		DepthClamp:                              features.depthClamp == C.VK_TRUE,
		DepthBiasClamp:                          features.depthBiasClamp == C.VK_TRUE,
		FillModeNonSolid:                        features.fillModeNonSolid == C.VK_TRUE,
		DepthBounds:                             features.depthBounds == C.VK_TRUE,
		WideLines:                               features.wideLines == C.VK_TRUE,
		LargePoints:                             features.largePoints == C.VK_TRUE,
		AlphaToOne:                              features.alphaToOne == C.VK_TRUE,
		MultiViewport:                           features.multiViewport == C.VK_TRUE,
		SamplerAnisotropy:                       features.samplerAnisotropy == C.VK_TRUE,
		TextureCompressionETC2:                  features.textureCompressionETC2 == C.VK_TRUE,
		TextureCompressionASTC_LDR:              features.textureCompressionASTC_LDR == C.VK_TRUE,
		TextureCompressionBC:                    features.textureCompressionBC == C.VK_TRUE,
		OcclusionQueryPrecise:                   features.occlusionQueryPrecise == C.VK_TRUE,
		PipelineStatisticsQuery:                 features.pipelineStatisticsQuery == C.VK_TRUE,
		VertexPipelineStoresAndAtomics:          features.vertexPipelineStoresAndAtomics == C.VK_TRUE,
		FragmentStoresAndAtomics:                features.fragmentStoresAndAtomics == C.VK_TRUE,
		ShaderTessellationAndGeometryPointSize:  features.shaderTessellationAndGeometryPointSize == C.VK_TRUE,
		ShaderImageGatherExtended:               features.shaderImageGatherExtended == C.VK_TRUE,
		ShaderStorageImageExtendedFormats:       features.shaderStorageImageExtendedFormats == C.VK_TRUE,
		ShaderStorageImageMultisample:           features.shaderStorageImageMultisample == C.VK_TRUE,
		ShaderStorageImageReadWithoutFormat:     features.shaderStorageImageReadWithoutFormat == C.VK_TRUE,
		ShaderStorageImageWriteWithoutFormat:    features.shaderStorageImageWriteWithoutFormat == C.VK_TRUE,
		ShaderUniformBufferArrayDynamicIndexing: features.shaderUniformBufferArrayDynamicIndexing == C.VK_TRUE,
		ShaderSampledImageArrayDynamicIndexing:  features.shaderSampledImageArrayDynamicIndexing == C.VK_TRUE,
		ShaderStorageBufferArrayDynamicIndexing: features.shaderStorageBufferArrayDynamicIndexing == C.VK_TRUE,
		ShaderStorageImageArrayDynamicIndexing:  features.shaderStorageImageArrayDynamicIndexing == C.VK_TRUE,
		ShaderClipDistance:                      features.shaderClipDistance == C.VK_TRUE,
		ShaderCullDistance:                      features.shaderCullDistance == C.VK_TRUE,
		ShaderFloat64:                           features.shaderFloat64 == C.VK_TRUE,
		ShaderInt64:                             features.shaderInt64 == C.VK_TRUE,
		ShaderInt16:                             features.shaderInt16 == C.VK_TRUE,
		ShaderResourceResidency:                 features.shaderResourceResidency == C.VK_TRUE,
		ShaderResourceMinLod:                    features.shaderResourceMinLod == C.VK_TRUE,
		SparseBinding:                           features.sparseBinding == C.VK_TRUE,
		SparseResidencyBuffer:                   features.sparseResidencyBuffer == C.VK_TRUE,
		SparseResidencyImage2D:                  features.sparseResidencyImage2D == C.VK_TRUE,
		SparseResidencyImage3D:                  features.sparseResidencyImage3D == C.VK_TRUE,
		SparseResidency2Samples:                 features.sparseResidency2Samples == C.VK_TRUE,
		SparseResidency4Samples:                 features.sparseResidency4Samples == C.VK_TRUE,
		SparseResidency8Samples:                 features.sparseResidency8Samples == C.VK_TRUE,
		SparseResidency16Samples:                features.sparseResidency16Samples == C.VK_TRUE,
		SparseResidencyAliased:                  features.sparseResidencyAliased == C.VK_TRUE,
		VariableMultisampleRate:                 features.variableMultisampleRate == C.VK_TRUE,
		InheritedQueries:                        features.inheritedQueries == C.VK_TRUE,
	}
}

type QueueFamilyProperties struct {
	QueueFlags                  QueueFlags
	QueueCount                  uint32
	TimestampValidBits          uint32
	MinImageTransferGranularity Extent3D

	// must be kept identical to C struct
}

type Extent2D struct {
	Width  uint32
	Height uint32

	// must be kept identical to C struct
}

type Extent3D struct {
	Width  uint32
	Height uint32
	Depth  uint32

	// must be kept identical to C struct
}

func (dev *PhysicalDevice) QueueFamilyProperties() []QueueFamilyProperties {
	var count C.uint32_t
	C.domVkGetPhysicalDeviceQueueFamilyProperties(dev.instance.fps[vkGetPhysicalDeviceQueueFamilyProperties], dev.hnd, &count, nil)
	props := make([]QueueFamilyProperties, count)
	C.domVkGetPhysicalDeviceQueueFamilyProperties(dev.instance.fps[vkGetPhysicalDeviceQueueFamilyProperties], dev.hnd, &count, (*C.VkQueueFamilyProperties)(slice2ptr(uptr(&props))))
	return props
}

type DeviceQueueCreateInfo struct {
	Extensions       []Extension
	Flags            DeviceQueueCreateFlags
	QueueFamilyIndex uint32
	QueuePriorities  []float32
}

type DeviceCreateInfo struct {
	Extensions            []Extension
	QueueCreateInfos      []DeviceQueueCreateInfo
	EnabledExtensionNames []string
	EnabledFeatures       *PhysicalDeviceFeatures
}

type Device struct {
	// VK_DEFINE_HANDLE(VkDevice)
	hnd C.VkDevice

	fps                 [deviceMaxPFN]C.PFN_vkVoidFunction
	vkGetDeviceProcAddr C.PFN_vkGetDeviceProcAddr
}

func (dev *PhysicalDevice) CreateDevice(info *DeviceCreateInfo) (*Device, error) {
	// TODO(dh): support custom allocators
	ptr := (*C.VkDeviceCreateInfo)(alloc(C.sizeof_VkDeviceCreateInfo))
	ptr.sType = C.VkStructureType(StructureTypeDeviceCreateInfo)
	ptr.pNext = buildChain(info.Extensions)
	defer internalizeChain(info.Extensions, ptr.pNext)
	ptr.queueCreateInfoCount = C.uint32_t(len(info.QueueCreateInfos))
	ptr.pQueueCreateInfos = (*C.VkDeviceQueueCreateInfo)(allocn(len(info.QueueCreateInfos), C.sizeof_VkDeviceQueueCreateInfo))
	defer free(uptr(ptr.pQueueCreateInfos))
	arr := (*[math.MaxInt32]C.VkDeviceQueueCreateInfo)(uptr(ptr.pQueueCreateInfos))[:len(info.QueueCreateInfos)]
	for i, obj := range info.QueueCreateInfos {
		arr[i] = C.VkDeviceQueueCreateInfo{
			sType:            C.VkStructureType(StructureTypeDeviceQueueCreateInfo),
			pNext:            buildChain(obj.Extensions),
			flags:            C.VkDeviceQueueCreateFlags(obj.Flags),
			queueFamilyIndex: C.uint32_t(obj.QueueFamilyIndex),
			queueCount:       C.uint32_t(len(obj.QueuePriorities)),
			pQueuePriorities: externFloat32(obj.QueuePriorities),
		}
		defer free(uptr(arr[i].pQueuePriorities))
		defer internalizeChain(obj.Extensions, arr[i].pNext)
	}
	ptr.enabledExtensionCount = C.uint32_t(len(info.EnabledExtensionNames))
	ptr.ppEnabledExtensionNames = externStrings(info.EnabledExtensionNames)
	defer free(uptr(ptr.ppEnabledExtensionNames))
	if info.EnabledFeatures != nil {
		ptr.pEnabledFeatures = (*C.VkPhysicalDeviceFeatures)(alloc(C.sizeof_VkPhysicalDeviceFeatures))
		ptr.pEnabledFeatures.robustBufferAccess = vkBool(info.EnabledFeatures.RobustBufferAccess)
		ptr.pEnabledFeatures.fullDrawIndexUint32 = vkBool(info.EnabledFeatures.FullDrawIndexUint32)
		ptr.pEnabledFeatures.imageCubeArray = vkBool(info.EnabledFeatures.ImageCubeArray)
		ptr.pEnabledFeatures.independentBlend = vkBool(info.EnabledFeatures.IndependentBlend)
		ptr.pEnabledFeatures.geometryShader = vkBool(info.EnabledFeatures.GeometryShader)
		ptr.pEnabledFeatures.tessellationShader = vkBool(info.EnabledFeatures.TessellationShader)
		ptr.pEnabledFeatures.sampleRateShading = vkBool(info.EnabledFeatures.SampleRateShading)
		ptr.pEnabledFeatures.dualSrcBlend = vkBool(info.EnabledFeatures.DualSrcBlend)
		ptr.pEnabledFeatures.logicOp = vkBool(info.EnabledFeatures.LogicOp)
		ptr.pEnabledFeatures.multiDrawIndirect = vkBool(info.EnabledFeatures.MultiDrawIndirect)
		ptr.pEnabledFeatures.drawIndirectFirstInstance = vkBool(info.EnabledFeatures.DrawIndirectFirstInstance)
		ptr.pEnabledFeatures.depthClamp = vkBool(info.EnabledFeatures.DepthClamp)
		ptr.pEnabledFeatures.depthBiasClamp = vkBool(info.EnabledFeatures.DepthBiasClamp)
		ptr.pEnabledFeatures.fillModeNonSolid = vkBool(info.EnabledFeatures.FillModeNonSolid)
		ptr.pEnabledFeatures.depthBounds = vkBool(info.EnabledFeatures.DepthBounds)
		ptr.pEnabledFeatures.wideLines = vkBool(info.EnabledFeatures.WideLines)
		ptr.pEnabledFeatures.largePoints = vkBool(info.EnabledFeatures.LargePoints)
		ptr.pEnabledFeatures.alphaToOne = vkBool(info.EnabledFeatures.AlphaToOne)
		ptr.pEnabledFeatures.multiViewport = vkBool(info.EnabledFeatures.MultiViewport)
		ptr.pEnabledFeatures.samplerAnisotropy = vkBool(info.EnabledFeatures.SamplerAnisotropy)
		ptr.pEnabledFeatures.textureCompressionETC2 = vkBool(info.EnabledFeatures.TextureCompressionETC2)
		ptr.pEnabledFeatures.textureCompressionASTC_LDR = vkBool(info.EnabledFeatures.TextureCompressionASTC_LDR)
		ptr.pEnabledFeatures.textureCompressionBC = vkBool(info.EnabledFeatures.TextureCompressionBC)
		ptr.pEnabledFeatures.occlusionQueryPrecise = vkBool(info.EnabledFeatures.OcclusionQueryPrecise)
		ptr.pEnabledFeatures.pipelineStatisticsQuery = vkBool(info.EnabledFeatures.PipelineStatisticsQuery)
		ptr.pEnabledFeatures.vertexPipelineStoresAndAtomics = vkBool(info.EnabledFeatures.VertexPipelineStoresAndAtomics)
		ptr.pEnabledFeatures.fragmentStoresAndAtomics = vkBool(info.EnabledFeatures.FragmentStoresAndAtomics)
		ptr.pEnabledFeatures.shaderTessellationAndGeometryPointSize = vkBool(info.EnabledFeatures.ShaderTessellationAndGeometryPointSize)
		ptr.pEnabledFeatures.shaderImageGatherExtended = vkBool(info.EnabledFeatures.ShaderImageGatherExtended)
		ptr.pEnabledFeatures.shaderStorageImageExtendedFormats = vkBool(info.EnabledFeatures.ShaderStorageImageExtendedFormats)
		ptr.pEnabledFeatures.shaderStorageImageMultisample = vkBool(info.EnabledFeatures.ShaderStorageImageMultisample)
		ptr.pEnabledFeatures.shaderStorageImageReadWithoutFormat = vkBool(info.EnabledFeatures.ShaderStorageImageReadWithoutFormat)
		ptr.pEnabledFeatures.shaderStorageImageWriteWithoutFormat = vkBool(info.EnabledFeatures.ShaderStorageImageWriteWithoutFormat)
		ptr.pEnabledFeatures.shaderUniformBufferArrayDynamicIndexing = vkBool(info.EnabledFeatures.ShaderUniformBufferArrayDynamicIndexing)
		ptr.pEnabledFeatures.shaderSampledImageArrayDynamicIndexing = vkBool(info.EnabledFeatures.ShaderSampledImageArrayDynamicIndexing)
		ptr.pEnabledFeatures.shaderStorageBufferArrayDynamicIndexing = vkBool(info.EnabledFeatures.ShaderStorageBufferArrayDynamicIndexing)
		ptr.pEnabledFeatures.shaderStorageImageArrayDynamicIndexing = vkBool(info.EnabledFeatures.ShaderStorageImageArrayDynamicIndexing)
		ptr.pEnabledFeatures.shaderClipDistance = vkBool(info.EnabledFeatures.ShaderClipDistance)
		ptr.pEnabledFeatures.shaderCullDistance = vkBool(info.EnabledFeatures.ShaderCullDistance)
		ptr.pEnabledFeatures.shaderFloat64 = vkBool(info.EnabledFeatures.ShaderFloat64)
		ptr.pEnabledFeatures.shaderInt64 = vkBool(info.EnabledFeatures.ShaderInt64)
		ptr.pEnabledFeatures.shaderInt16 = vkBool(info.EnabledFeatures.ShaderInt16)
		ptr.pEnabledFeatures.shaderResourceResidency = vkBool(info.EnabledFeatures.ShaderResourceResidency)
		ptr.pEnabledFeatures.shaderResourceMinLod = vkBool(info.EnabledFeatures.ShaderResourceMinLod)
		ptr.pEnabledFeatures.sparseBinding = vkBool(info.EnabledFeatures.SparseBinding)
		ptr.pEnabledFeatures.sparseResidencyBuffer = vkBool(info.EnabledFeatures.SparseResidencyBuffer)
		ptr.pEnabledFeatures.sparseResidencyImage2D = vkBool(info.EnabledFeatures.SparseResidencyImage2D)
		ptr.pEnabledFeatures.sparseResidencyImage3D = vkBool(info.EnabledFeatures.SparseResidencyImage3D)
		ptr.pEnabledFeatures.sparseResidency2Samples = vkBool(info.EnabledFeatures.SparseResidency2Samples)
		ptr.pEnabledFeatures.sparseResidency4Samples = vkBool(info.EnabledFeatures.SparseResidency4Samples)
		ptr.pEnabledFeatures.sparseResidency8Samples = vkBool(info.EnabledFeatures.SparseResidency8Samples)
		ptr.pEnabledFeatures.sparseResidency16Samples = vkBool(info.EnabledFeatures.SparseResidency16Samples)
		ptr.pEnabledFeatures.sparseResidencyAliased = vkBool(info.EnabledFeatures.SparseResidencyAliased)
		ptr.pEnabledFeatures.variableMultisampleRate = vkBool(info.EnabledFeatures.VariableMultisampleRate)
		ptr.pEnabledFeatures.inheritedQueries = vkBool(info.EnabledFeatures.InheritedQueries)
		defer free(uptr(ptr.pEnabledFeatures))
	}
	var out C.VkDevice
	res := Result(C.domVkCreateDevice(dev.instance.fps[vkCreateDevice], dev.hnd, ptr, nil, &out))
	if res != Success {
		return nil, res
	}
	ldev := &Device{
		hnd:                 out,
		vkGetDeviceProcAddr: C.PFN_vkGetDeviceProcAddr(dev.instance.fps[vkGetDeviceProcAddr]),
	}
	ldev.init()

	return ldev, nil
}

func (dev *Device) Destroy() {
	// TODO(dh): support custom allocators
	C.domVkDestroyDevice(dev.fps[vkDestroyDevice], dev.hnd, nil)
}

func (dev *Device) init() {
	for i, name := range deviceFpNames {
		dev.fps[i] = dev.getDeviceProcAddr(name)
	}
}

func (dev *Device) getDeviceProcAddr(name string) C.PFN_vkVoidFunction {
	cName := C.CString(name)
	defer free(uptr(cName))
	fp := C.domVkGetDeviceProcAddr(dev.vkGetDeviceProcAddr, dev.hnd, cName)
	if debug {
		fmt.Fprintf(os.Stderr, "%s = %p\n", name, fp)
	}
	return fp
}

type Queue struct {
	// VK_DEFINE_HANDLE(VkQueue)
	hnd C.VkQueue
	fps *[deviceMaxPFN]C.PFN_vkVoidFunction
}

func (q *Queue) WaitIdle() error {
	res := Result(C.domVkQueueWaitIdle(q.fps[vkQueueWaitIdle], q.hnd))
	return result2error(res)
}

func (dev *Device) Queue(family, index uint32) *Queue {
	var out C.VkQueue
	C.domVkGetDeviceQueue(dev.fps[vkGetDeviceQueue], dev.hnd, C.uint32_t(family), C.uint32_t(index), &out)
	return &Queue{hnd: out, fps: &dev.fps}
}

type CommandBuffer struct {
	// VK_DEFINE_HANDLE(VkCommandBuffer)
	hnd C.VkCommandBuffer
	fps *[deviceMaxPFN]C.PFN_vkVoidFunction
}

func (buf *CommandBuffer) Reset(flags CommandBufferResetFlags) error {
	res := Result(C.domVkResetCommandBuffer(buf.fps[vkResetCommandBuffer], buf.hnd, C.VkCommandBufferResetFlags(flags)))
	return result2error(res)
}

type CommandBufferBeginInfo struct {
	Extensions      []Extension
	Flags           CommandBufferUsageFlags
	InheritanceInfo *CommandBufferInheritanceInfo
}

type RenderPass struct {
	// VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkRenderPass)
	hnd C.VkRenderPass
}

type Framebuffer struct {
	// VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkFramebuffer)
	hnd C.VkFramebuffer
}

type CommandBufferInheritanceInfo struct {
	Extensions           []Extension
	RenderPass           RenderPass
	Subpass              uint32
	Framebuffer          Framebuffer
	OcclusionQueryEnable bool
	QueryFlags           QueryControlFlags
	PipelineStatistics   QueryPipelineStatisticFlags
}

func (buf *CommandBuffer) Begin(info *CommandBufferBeginInfo) error {
	ptr := (*C.VkCommandBufferBeginInfo)(alloc(C.sizeof_VkCommandBufferBeginInfo))
	defer free(uptr(ptr))
	ptr.sType = C.VkStructureType(StructureTypeCommandBufferBeginInfo)
	ptr.pNext = buildChain(info.Extensions)
	defer internalizeChain(info.Extensions, ptr.pNext)
	ptr.flags = C.VkCommandBufferUsageFlags(info.Flags)
	if info.InheritanceInfo != nil {
		ptr.pInheritanceInfo = (*C.VkCommandBufferInheritanceInfo)(alloc(C.sizeof_VkCommandBufferInheritanceInfo))
		defer free(uptr(ptr.pInheritanceInfo))
		ptr.pInheritanceInfo.sType = C.VkStructureType(StructureTypeCommandBufferInheritanceInfo)
		ptr.pInheritanceInfo.pNext = buildChain(info.InheritanceInfo.Extensions)
		defer internalizeChain(info.InheritanceInfo.Extensions, ptr.pInheritanceInfo.pNext)
		ptr.pInheritanceInfo.renderPass = C.VkRenderPass(info.InheritanceInfo.RenderPass.hnd)
		ptr.pInheritanceInfo.subpass = C.uint32_t(info.InheritanceInfo.Subpass)
		ptr.pInheritanceInfo.framebuffer = C.VkFramebuffer(info.InheritanceInfo.Framebuffer.hnd)
		ptr.pInheritanceInfo.occlusionQueryEnable = vkBool(info.InheritanceInfo.OcclusionQueryEnable)
		ptr.pInheritanceInfo.queryFlags = C.VkQueryControlFlags(info.InheritanceInfo.QueryFlags)
		ptr.pInheritanceInfo.pipelineStatistics = C.VkQueryPipelineStatisticFlags(info.InheritanceInfo.PipelineStatistics)
	}
	res := Result(C.domVkBeginCommandBuffer(buf.fps[vkBeginCommandBuffer], buf.hnd, ptr))
	return result2error(res)
}

func (buf *CommandBuffer) End() error {
	res := Result(C.domVkEndCommandBuffer(buf.fps[vkEndCommandBuffer], buf.hnd))
	return result2error(res)
}

func (buf *CommandBuffer) SetLineWidth(lineWidth float32) {
	C.domVkCmdSetLineWidth(buf.fps[vkCmdSetLineWidth], buf.hnd, C.float(lineWidth))
}

func (buf *CommandBuffer) SetDepthBias(constantFactor, clamp, slopeFactor float32) {
	C.domVkCmdSetDepthBias(buf.fps[vkCmdSetDepthBias], buf.hnd, C.float(constantFactor), C.float(clamp), C.float(slopeFactor))
}

func (buf *CommandBuffer) SetBlendConstants(blendConstants [4]float32) {
	C.domVkCmdSetBlendConstants(buf.fps[vkCmdSetBlendConstants], buf.hnd, (*C.float)(slice2ptr(uptr(&blendConstants))))
}

func (buf *CommandBuffer) Draw(vertexCount, instanceCount, firstVertex, firstInstance uint32) {
	C.domVkCmdDraw(buf.fps[vkCmdDraw], buf.hnd, C.uint32_t(vertexCount), C.uint32_t(instanceCount), C.uint32_t(firstVertex), C.uint32_t(firstInstance))
}

func (buf *CommandBuffer) SetViewport(firstViewport uint32, viewports []Viewport) {
	C.domVkCmdSetViewport(buf.fps[vkCmdSetViewport], buf.hnd, C.uint32_t(firstViewport), C.uint32_t(len(viewports)), (*C.VkViewport)(slice2ptr(uptr(&viewports))))
}

func (buf *CommandBuffer) SetScissor(firstScissor uint32, scissors []Rect2D) {
	C.domVkCmdSetScissor(buf.fps[vkCmdSetScissor], buf.hnd, C.uint32_t(firstScissor), C.uint32_t(len(scissors)), (*C.VkRect2D)(slice2ptr(uptr(&scissors))))
}

func (buf *CommandBuffer) SetDeviceMask(deviceMask uint32) {
	C.domVkCmdSetDeviceMask(buf.fps[vkCmdSetDeviceMask], buf.hnd, C.uint32_t(deviceMask))
}

func (buf *CommandBuffer) SetDepthBounds(min, max float32) {
	C.domVkCmdSetDepthBounds(buf.fps[vkCmdSetDepthBounds], buf.hnd, C.float(min), C.float(max))
}

func (buf *CommandBuffer) PushConstants(layout PipelineLayout, stageFlags ShaderStageFlags, offset uint32, size uint32, data []byte) {
	C.domVkCmdPushConstants(buf.fps[vkCmdPushConstants], buf.hnd, layout.hnd, C.VkShaderStageFlags(stageFlags), C.uint32_t(offset), C.uint32_t(len(data)), slice2ptr(uptr(&data)))
}

func (buf *CommandBuffer) FillBuffer(dstBuffer Buffer, dstOffset DeviceSize, size DeviceSize, data uint32) {
	C.domVkCmdFillBuffer(buf.fps[vkCmdFillBuffer], buf.hnd, dstBuffer.hnd, C.VkDeviceSize(dstOffset), C.VkDeviceSize(size), C.uint32_t(data))
}

func (buf *CommandBuffer) Dispatch(x, y, z uint32) {
	C.domVkCmdDispatch(buf.fps[vkCmdDispatch], buf.hnd, C.uint32_t(x), C.uint32_t(y), C.uint32_t(z))
}

func (buf *CommandBuffer) SetEvent(event Event, stageMask PipelineStageFlags) {
	C.domVkCmdSetEvent(buf.fps[vkCmdSetEvent], buf.hnd, event.hnd, C.VkPipelineStageFlags(stageMask))
}

type ClearAttachment struct {
	AspectMask      ImageAspectFlags
	ColorAttachment uint32
	ClearValue      ClearValue
}

type ClearRect struct {
	Rect           Rect2D
	BaseArrayLayer uint32
	LayerCount     uint32

	// must be kept identical to C struct
}

func (buf *CommandBuffer) ClearAttachments(attachments []ClearAttachment, rects []ClearRect) {
	mem := allocn(len(attachments), C.sizeof_VkClearAttachment)
	arr := (*[math.MaxInt32]C.VkClearAttachment)(mem)[:len(attachments)]
	for i := range arr {
		arr[i] = C.VkClearAttachment{
			aspectMask:      C.VkImageAspectFlags(attachments[i].AspectMask),
			colorAttachment: C.uint32_t(attachments[i].ColorAttachment),
		}
		switch v := attachments[i].ClearValue.(type) {
		case ClearColorValueFloat32s:
			copy(arr[i].clearValue[:], (*[16]byte)(uptr(&v))[:])
		case ClearColorValueInt32s:
			copy(arr[i].clearValue[:], (*[16]byte)(uptr(&v))[:])
		case ClearColorValueUint32s:
			copy(arr[i].clearValue[:], (*[16]byte)(uptr(&v))[:])
		case ClearDepthStencilValue:
			ucopy1(uptr(&arr[i].clearValue), uptr(&v), C.sizeof_VkClearDepthStencilValue)
		default:
			panic(fmt.Sprintf("unreachable: %T", v))
		}
	}
	C.domVkCmdClearAttachments(buf.fps[vkCmdClearAttachments], buf.hnd, C.uint32_t(len(attachments)), (*C.VkClearAttachment)(mem), C.uint32_t(len(rects)), (*C.VkClearRect)(slice2ptr(uptr(&rects))))
	free(uptr(mem))
}

func (buf *CommandBuffer) ClearColorImage(image Image, imageLayout ImageLayout, color ClearColorValue, ranges []ImageSubresourceRange) {
	cColor := (*C.VkClearColorValue)(alloc(C.sizeof_VkClearColorValue))
	switch v := color.(type) {
	case ClearColorValueFloat32s:
		copy(cColor[:], (*[16]byte)(uptr(&v))[:])
	case ClearColorValueInt32s:
		copy(cColor[:], (*[16]byte)(uptr(&v))[:])
	case ClearColorValueUint32s:
		copy(cColor[:], (*[16]byte)(uptr(&v))[:])
	default:
		panic(fmt.Sprintf("unreachable: %T", v))
	}
	C.domVkCmdClearColorImage(buf.fps[vkCmdClearColorImage], buf.hnd, image.hnd, C.VkImageLayout(imageLayout), cColor, C.uint32_t(len(ranges)), (*C.VkImageSubresourceRange)(slice2ptr(uptr(&ranges))))
	free(uptr(cColor))
}

func (buf *CommandBuffer) ClearDepthStencilImage(image Image, imageLayout ImageLayout, depthStencil ClearDepthStencilValue, ranges []ImageSubresourceRange) {
	C.domVkCmdClearDepthStencilImage(buf.fps[vkCmdClearDepthStencilImage], buf.hnd, image.hnd, C.VkImageLayout(imageLayout), (*C.VkClearDepthStencilValue)(uptr(&depthStencil)), C.uint32_t(len(ranges)), (*C.VkImageSubresourceRange)(slice2ptr(uptr(&ranges))))
}

func (info *RenderPassBeginInfo) c() *C.VkRenderPassBeginInfo {
	size0 := align(C.sizeof_VkRenderPassBeginInfo)
	size1 := align(C.sizeof_VkClearValue * uintptr(len(info.ClearValues)))
	size := size0 + size1
	mem := alloc(C.size_t(size))
	cinfo := (*C.VkRenderPassBeginInfo)(mem)
	*cinfo = C.VkRenderPassBeginInfo{
		sType:           C.VkStructureType(StructureTypeRenderPassBeginInfo),
		pNext:           buildChain(info.Extensions),
		renderPass:      info.RenderPass.hnd,
		framebuffer:     info.Framebuffer.hnd,
		clearValueCount: C.uint32_t(len(info.ClearValues)),
		pClearValues:    (*C.VkClearValue)(uptr(uintptr(mem) + size0)),
	}
	ucopy1(uptr(&cinfo.renderArea), uptr(&info.RenderArea), C.sizeof_VkRect2D)
	arr := (*[math.MaxInt32]C.VkClearValue)(uptr(cinfo.pClearValues))[:len(info.ClearValues)]
	for i := range arr {
		switch v := info.ClearValues[i].(type) {
		case ClearColorValueFloat32s:
			copy(arr[i][:], (*[16]byte)(uptr(&v))[:])
		case ClearColorValueInt32s:
			copy(arr[i][:], (*[16]byte)(uptr(&v))[:])
		case ClearColorValueUint32s:
			copy(arr[i][:], (*[16]byte)(uptr(&v))[:])
		case ClearDepthStencilValue:
			ucopy1(uptr(&arr[i]), uptr(&v), C.sizeof_VkClearDepthStencilValue)
		default:
			panic(fmt.Sprintf("unreachable: %T", v))
		}
	}
	return cinfo
}

func (buf *CommandBuffer) BeginRenderPass(info *RenderPassBeginInfo, contents SubpassContents) {
	cinfo := info.c()
	C.domVkCmdBeginRenderPass(buf.fps[vkCmdBeginRenderPass], buf.hnd, cinfo, C.VkSubpassContents(contents))
	internalizeChain(info.Extensions, cinfo.pNext)
	free(uptr(cinfo))
}

func (buf *CommandBuffer) EndRenderPass() {
	C.domVkCmdEndRenderPass(buf.fps[vkCmdEndRenderPass], buf.hnd)
}

func (buf *CommandBuffer) BindPipeline(pipelineBindPoint PipelineBindPoint, pipeline Pipeline) {
	C.domVkCmdBindPipeline(buf.fps[vkCmdBindPipeline], buf.hnd, C.VkPipelineBindPoint(pipelineBindPoint), pipeline.hnd)
}

func (buf *CommandBuffer) BindIndexBuffer(buffer Buffer, offset DeviceSize, indexType IndexType) {
	C.domVkCmdBindIndexBuffer(buf.fps[vkCmdBindIndexBuffer], buf.hnd, buffer.hnd, C.VkDeviceSize(offset), C.VkIndexType(indexType))
}

type BufferCopy struct {
	SrcOffset DeviceSize
	DstOffset DeviceSize
	Size      DeviceSize

	// must be kept identical to C struct
}

func (buf *CommandBuffer) CopyBuffer(srcBuffer, dstBuffer Buffer, regions []BufferCopy) {
	C.domVkCmdCopyBuffer(buf.fps[vkCmdCopyBuffer], buf.hnd, srcBuffer.hnd, dstBuffer.hnd, C.uint32_t(len(regions)), (*C.VkBufferCopy)(slice2ptr(uptr(&regions))))
}

type BufferImageCopy struct {
	BufferOfset       DeviceSize
	BufferRowLength   uint32
	BufferImageHeight uint32
	ImageSubresource  ImageSubresourceLayers
	ImageOffset       Offset3D
	ImageExtent       Extent3D

	// must be kept identical to C struct
}

type ImageSubresourceLayers struct {
	AspectMask     ImageAspectFlags
	MipLevel       uint32
	BaseArrayLayer uint32
	LayerCount     uint32

	// must be kept identical to C struct
}

func (buf *CommandBuffer) CopyBufferToImage(srcBuffer Buffer, dstImage Image, dstImageLayout ImageLayout, regions []BufferImageCopy) {
	C.domVkCmdCopyBufferToImage(
		buf.fps[vkCmdCopyBufferToImage],
		buf.hnd,
		srcBuffer.hnd,
		dstImage.hnd,
		C.VkImageLayout(dstImageLayout),
		C.uint32_t(len(regions)),
		(*C.VkBufferImageCopy)(slice2ptr(uptr(&regions))))
}

type ImageCopy struct {
	SrcSubresource ImageSubresourceLayers
	SrcOffset      Offset3D
	DstSubresource ImageSubresourceLayers
	DstOffset      Offset3D
	Extent         Extent3D

	// must be kept identical to C struct
}

func (buf *CommandBuffer) CopyImage(srcImage Image, srcImageLayout ImageLayout, dstImage Image, dstImageLayout ImageLayout, regions []ImageCopy) {
	C.domVkCmdCopyImage(
		buf.fps[vkCmdCopyImage],
		buf.hnd,
		srcImage.hnd,
		C.VkImageLayout(srcImageLayout),
		dstImage.hnd,
		C.VkImageLayout(dstImageLayout),
		C.uint32_t(len(regions)),
		(*C.VkImageCopy)(slice2ptr(uptr(&regions))))
}

func (buf *CommandBuffer) CopyImageToBuffer(srcImage Image, srcImageLayout ImageLayout, dstBuffer Buffer, regions []BufferImageCopy) {
	C.domVkCmdCopyImageToBuffer(
		buf.fps[vkCmdCopyImageToBuffer],
		buf.hnd,
		srcImage.hnd,
		C.VkImageLayout(srcImageLayout),
		dstBuffer.hnd,
		C.uint32_t(len(regions)),
		(*C.VkBufferImageCopy)(slice2ptr(uptr(&regions))))
}

func (buf *CommandBuffer) ResetEvent(event Event, stageMask PipelineStageFlags) {
	C.domVkCmdResetEvent(buf.fps[vkCmdResetEvent], buf.hnd, event.hnd, C.VkPipelineStageFlags(stageMask))
}

type CommandPoolCreateInfo struct {
	Extensions       []Extension
	Flags            CommandPoolCreateFlags
	QueueFamilyIndex uint32
}

type CommandPool struct {
	// VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkCommandPool)
	hnd C.VkCommandPool
}

func (dev *Device) CreateCommandPool(info *CommandPoolCreateInfo) (CommandPool, error) {
	// TODO(dh): support callbacks
	ptr := (*C.VkCommandPoolCreateInfo)(alloc(C.sizeof_VkCommandPoolCreateInfo))
	ptr.sType = C.VkStructureType(StructureTypeCommandPoolCreateInfo)
	ptr.pNext = buildChain(info.Extensions)
	ptr.flags = C.VkCommandPoolCreateFlags(info.Flags)
	ptr.queueFamilyIndex = C.uint32_t(info.QueueFamilyIndex)
	var pool CommandPool
	res := Result(C.domVkCreateCommandPool(dev.fps[vkCreateCommandPool], dev.hnd, ptr, nil, &pool.hnd))
	internalizeChain(info.Extensions, ptr.pNext)
	free(uptr(ptr))
	return pool, result2error(res)
}

func (dev *Device) DestroyCommandPool(pool CommandPool) {
	// TODO(dh): support callbacks
	C.domVkDestroyCommandPool(dev.fps[vkDestroyCommandPool], dev.hnd, pool.hnd, nil)
}

func (dev *Device) TrimCommandPool(pool CommandPool, flags CommandPoolTrimFlags) {
	C.domVkTrimCommandPool(dev.fps[vkTrimCommandPool], dev.hnd, pool.hnd, C.VkCommandPoolTrimFlags(flags))
}

func (dev *Device) ResetCommandPool(pool CommandPool, flags CommandPoolResetFlags) error {
	res := Result(C.domVkResetCommandPool(dev.fps[vkResetCommandPool], dev.hnd, pool.hnd, C.VkCommandPoolResetFlags(flags)))
	return result2error(res)
}

type CommandBufferAllocateInfo struct {
	Extensions         []Extension
	Level              CommandBufferLevel
	CommandBufferCount uint32
}

func (dev *Device) AllocateCommandBuffers(pool CommandPool, info *CommandBufferAllocateInfo) ([]*CommandBuffer, error) {
	ptr := (*C.VkCommandBufferAllocateInfo)(alloc(C.sizeof_VkCommandBufferAllocateInfo))
	ptr.sType = C.VkStructureType(StructureTypeCommandBufferAllocateInfo)
	ptr.pNext = buildChain(info.Extensions)
	ptr.commandPool = pool.hnd
	ptr.level = C.VkCommandBufferLevel(info.Level)
	ptr.commandBufferCount = C.uint32_t(info.CommandBufferCount)
	bufs := make([]C.VkCommandBuffer, info.CommandBufferCount)
	res := Result(C.domVkAllocateCommandBuffers(dev.fps[vkAllocateCommandBuffers], dev.hnd, ptr, (*C.VkCommandBuffer)(slice2ptr(uptr(&bufs)))))
	internalizeChain(info.Extensions, ptr.pNext)
	free(uptr(ptr))
	if res != Success {
		return nil, res
	}
	out := make([]*CommandBuffer, info.CommandBufferCount)
	for i, buf := range bufs {
		out[i] = &CommandBuffer{hnd: buf, fps: &dev.fps}
	}
	return out, nil
}

func (dev *Device) FreeCommandBuffers(pool CommandPool, bufs []*CommandBuffer) {
	if len(bufs) == 1 {
		C.domVkFreeCommandBuffers(dev.fps[vkFreeCommandBuffers], dev.hnd, pool.hnd, 1, (*C.VkCommandBuffer)(uptr(bufs[0])))
		return
	}

	// OPT(dh): cache this slice and reuse it for multiple
	// FreeCommandBuffers calls. Since the function has to be
	// reentrant, and we don't want to store the slice in the
	// CommandPool itself, we'd probably best use a sync.Pool.
	ptrs := make([]C.VkCommandBuffer, len(bufs))
	for i, buf := range bufs {
		ptrs[i] = buf.hnd
	}
	C.domVkFreeCommandBuffers(dev.fps[vkFreeCommandBuffers], dev.hnd, pool.hnd, C.uint32_t(len(bufs)), (*C.VkCommandBuffer)(slice2ptr(uptr(&ptrs))))
}

func (dev *Device) WaitIdle() error {
	res := Result(C.domVkDeviceWaitIdle(dev.fps[vkDeviceWaitIdle], dev.hnd))
	return result2error(res)
}

type Image struct {
	// VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkImage)
	hnd C.VkImage
}

type ImageView struct {
	// VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkImageView)
	hnd C.VkImageView

	// must be kept identical to C struct
}

type ImageViewCreateInfo struct {
	Extensions       []Extension
	Image            Image
	ViewType         ImageViewType
	Format           Format
	Components       ComponentMapping
	SubresourceRange ImageSubresourceRange
}

type ComponentMapping struct {
	R ComponentSwizzle
	G ComponentSwizzle
	B ComponentSwizzle
	A ComponentSwizzle

	// must be kept identical to C struct
}

type ImageSubresourceRange struct {
	AspectMask     ImageAspectFlags
	BaseMipLevel   uint32
	LevelCount     uint32
	BaseArrayLayer uint32
	LayerCount     uint32

	// must be kept identical to C struct
}

func (dev *Device) CreateImageView(info *ImageViewCreateInfo) (ImageView, error) {
	// TODO(dh): support custom allocator
	ptr := (*C.VkImageViewCreateInfo)(alloc(C.sizeof_VkImageViewCreateInfo))
	ptr.sType = C.VkStructureType(StructureTypeImageViewCreateInfo)
	ptr.pNext = buildChain(info.Extensions)
	ptr.image = info.Image.hnd
	ptr.viewType = C.VkImageViewType(info.ViewType)
	ptr.format = C.VkFormat(info.Format)
	ucopy1(uptr(&ptr.components), uptr(&info.Components), C.sizeof_VkComponentMapping)
	ucopy1(uptr(&ptr.subresourceRange), uptr(&info.SubresourceRange), C.sizeof_VkImageSubresourceRange)

	var out ImageView
	res := Result(C.domVkCreateImageView(dev.fps[vkCreateImageView], dev.hnd, ptr, nil, &out.hnd))
	internalizeChain(info.Extensions, ptr.pNext)
	free(uptr(ptr))
	return out, result2error(res)
}

func (dev *Device) DestroyImageView(view ImageView) {
	// TODO(dh): support custom allocator
	C.domVkDestroyImageView(dev.fps[vkDestroyImageView], dev.hnd, view.hnd, nil)
}

type ShaderModule struct {
	// VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkShaderModule)
	hnd C.VkShaderModule
}

type ShaderModuleCreateInfo struct {
	Extensions []Extension
	Code       []byte
}

func (dev *Device) CreateShaderModule(info *ShaderModuleCreateInfo) (ShaderModule, error) {
	// TODO(dh): support custom allocator
	ptr := (*C.VkShaderModuleCreateInfo)(alloc(C.sizeof_VkShaderModuleCreateInfo))
	ptr.sType = C.VkStructureType(StructureTypeShaderModuleCreateInfo)
	ptr.pNext = buildChain(info.Extensions)
	ptr.codeSize = C.size_t(len(info.Code))
	ptr.pCode = (*C.uint32_t)(C.CBytes(info.Code))
	defer free(uptr(ptr.pCode))
	var out ShaderModule
	res := Result(C.domVkCreateShaderModule(dev.fps[vkCreateShaderModule], dev.hnd, ptr, nil, &out.hnd))
	internalizeChain(info.Extensions, ptr.pNext)
	free(uptr(ptr))
	return out, result2error(res)
}

type PipelineShaderStageCreateInfo struct {
	Extensions []Extension
	Stage      ShaderStageFlags
	Module     ShaderModule
	Name       string
	// TODO(dh): support specialization info
}

type PipelineVertexInputStateCreateInfo struct {
	Extensions                  []Extension
	VertexBindingDescriptions   []VertexInputBindingDescription
	VertexAttributeDescriptions []VertexInputAttributeDescription
}

func (info *PipelineVertexInputStateCreateInfo) c() *C.VkPipelineVertexInputStateCreateInfo {
	size0 := align(C.sizeof_VkPipelineVertexInputStateCreateInfo)
	size1 := align(uintptr(len(info.VertexBindingDescriptions)) * C.sizeof_VkVertexInputBindingDescription)
	size2 := align(uintptr(len(info.VertexAttributeDescriptions)) * C.sizeof_VkVertexInputAttributeDescription)
	size := size0 + size1 + size2

	mem := alloc(C.size_t(size))
	cinfo := (*C.VkPipelineVertexInputStateCreateInfo)(mem)
	bindings := (*C.VkVertexInputBindingDescription)(uptr(uintptr(mem) + size0))
	attribs := (*C.VkVertexInputAttributeDescription)(uptr(uintptr(mem) + size0 + size1))
	*cinfo = C.VkPipelineVertexInputStateCreateInfo{
		sType: C.VkStructureType(StructureType(StructureTypePipelineVertexInputStateCreateInfo)),
		pNext: buildChain(info.Extensions),
		flags: 0,
		vertexBindingDescriptionCount:   C.uint32_t(len(info.VertexBindingDescriptions)),
		pVertexBindingDescriptions:      bindings,
		vertexAttributeDescriptionCount: C.uint32_t(len(info.VertexAttributeDescriptions)),
		pVertexAttributeDescriptions:    attribs,
	}
	ucopy(uptr(bindings), uptr(&info.VertexBindingDescriptions), C.sizeof_VkVertexInputBindingDescription)
	ucopy(uptr(attribs), uptr(&info.VertexAttributeDescriptions), C.sizeof_VkVertexInputAttributeDescription)
	return cinfo
}

type VertexInputBindingDescription struct {
	Binding   uint32
	Stride    uint32
	InputRate VertexInputRate

	// must be kept identical to C struct
}

type VertexInputAttributeDescription struct {
	Location uint32
	Binding  uint32
	Format   Format
	Offset   uint32

	// must be kept identical to C struct
}

type PipelineInputAssemblyStateCreateInfo struct {
	Extensions             []Extension
	Topology               PrimitiveTopology
	PrimitiveRestartEnable bool
}

func (info *PipelineInputAssemblyStateCreateInfo) c() *C.VkPipelineInputAssemblyStateCreateInfo {
	cinfo := (*C.VkPipelineInputAssemblyStateCreateInfo)(alloc(C.sizeof_VkPipelineInputAssemblyStateCreateInfo))
	*cinfo = C.VkPipelineInputAssemblyStateCreateInfo{
		sType:                  C.VkStructureType(StructureTypePipelineInputAssemblyStateCreateInfo),
		pNext:                  buildChain(info.Extensions),
		flags:                  0,
		topology:               C.VkPrimitiveTopology(info.Topology),
		primitiveRestartEnable: vkBool(info.PrimitiveRestartEnable),
	}
	return cinfo
}

type Viewport struct {
	X        float32
	Y        float32
	Width    float32
	Height   float32
	MinDepth float32
	MaxDepth float32

	// must be kept identical to C struct
}

type Rect2D struct {
	Offset Offset2D
	Extent Extent2D

	// must be kept identical to C struct
}

type Offset2D struct {
	X int32
	Y int32

	// must be kept identical to C struct
}

type Offset3D struct {
	X int32
	Y int32
	Z int32

	// must be kept identical to C struct
}

type PipelineViewportStateCreateInfo struct {
	Extensions []Extension
	Viewports  []Viewport
	Scissors   []Rect2D
}

func (info *PipelineViewportStateCreateInfo) c() *C.VkPipelineViewportStateCreateInfo {
	size0 := align(C.sizeof_VkPipelineViewportStateCreateInfo)
	size1 := align(uintptr(len(info.Viewports)) * C.sizeof_VkViewport)
	size2 := align(uintptr(len(info.Scissors)) * C.sizeof_VkRect2D)
	size := size0 + size1 + size2
	mem := alloc(C.size_t(size))
	cinfo := (*C.VkPipelineViewportStateCreateInfo)(mem)
	viewports := (*C.VkViewport)(uptr(uintptr(mem) + size0))
	scissors := (*C.VkRect2D)(uptr(uintptr(mem) + size0 + size1))
	*cinfo = C.VkPipelineViewportStateCreateInfo{
		sType:         C.VkStructureType(StructureTypePipelineViewportStateCreateInfo),
		pNext:         buildChain(info.Extensions),
		flags:         0,
		viewportCount: C.uint32_t(len(info.Viewports)),
		pViewports:    viewports,
		scissorCount:  C.uint32_t(len(info.Scissors)),
		pScissors:     scissors,
	}
	ucopy(uptr(viewports), uptr(&info.Viewports), C.sizeof_VkViewport)
	ucopy(uptr(scissors), uptr(&info.Scissors), C.sizeof_VkRect2D)
	return cinfo
}

type PipelineRasterizationStateCreateInfo struct {
	Extensions              []Extension
	DepthClampEnable        bool
	RasterizerDiscardEnable bool
	PolygonMode             PolygonMode
	CullMode                CullModeFlags
	FrontFace               FrontFace
	DepthBiasEnable         bool
	DepthBiasConstantFactor float32
	DepthBiasClamp          float32
	DepthBiasSlopeFactor    float32
	LineWidth               float32
}

func (info *PipelineRasterizationStateCreateInfo) c() *C.VkPipelineRasterizationStateCreateInfo {
	cinfo := (*C.VkPipelineRasterizationStateCreateInfo)(alloc(C.sizeof_VkPipelineRasterizationStateCreateInfo))
	*cinfo = C.VkPipelineRasterizationStateCreateInfo{
		sType:                   C.VkStructureType(StructureTypePipelineRasterizationStateCreateInfo),
		pNext:                   buildChain(info.Extensions),
		flags:                   0,
		depthClampEnable:        vkBool(info.DepthClampEnable),
		rasterizerDiscardEnable: vkBool(info.RasterizerDiscardEnable),
		polygonMode:             C.VkPolygonMode(info.PolygonMode),
		cullMode:                C.VkCullModeFlags(info.CullMode),
		frontFace:               C.VkFrontFace(info.FrontFace),
		depthBiasEnable:         vkBool(info.DepthBiasEnable),
		depthBiasConstantFactor: C.float(info.DepthBiasConstantFactor),
		depthBiasClamp:          C.float(info.DepthBiasClamp),
		depthBiasSlopeFactor:    C.float(info.DepthBiasSlopeFactor),
		lineWidth:               C.float(info.LineWidth),
	}
	return cinfo
}

type PipelineMultisampleStateCreateInfo struct {
	Extensions            []Extension
	RasterizationSamples  SampleCountFlags
	SampleShadingEnable   bool
	MinSampleShading      float32
	SampleMask            []SampleMask
	AlphaToCoverageEnable bool
	AlphaToOneEnable      bool
}

func (info *PipelineMultisampleStateCreateInfo) c() *C.VkPipelineMultisampleStateCreateInfo {
	size0 := align(C.sizeof_VkPipelineMultisampleStateCreateInfo)
	size1 := align(uintptr(len(info.SampleMask)) * C.sizeof_VkSampleMask)
	size := size0 + size1
	mem := alloc(C.size_t(size))
	cinfo := (*C.VkPipelineMultisampleStateCreateInfo)(mem)
	var sampleMask *C.VkSampleMask
	if info.SampleMask != nil {
		sampleMask = (*C.VkSampleMask)(uptr(uintptr(mem) + size0))
		ucopy(uptr(sampleMask), uptr(&info.SampleMask), C.sizeof_VkSampleMask)
	}

	*cinfo = C.VkPipelineMultisampleStateCreateInfo{
		sType:                 C.VkStructureType(StructureTypePipelineMultisampleStateCreateInfo),
		pNext:                 buildChain(info.Extensions),
		flags:                 0,
		rasterizationSamples:  C.VkSampleCountFlagBits(info.RasterizationSamples),
		sampleShadingEnable:   vkBool(info.SampleShadingEnable),
		minSampleShading:      C.float(info.MinSampleShading),
		pSampleMask:           sampleMask,
		alphaToCoverageEnable: vkBool(info.AlphaToCoverageEnable),
		alphaToOneEnable:      vkBool(info.AlphaToOneEnable),
	}
	return cinfo
}

type SampleMask uint32

type PipelineColorBlendAttachmentState struct {
	BlendEnable         bool
	SrcColorBlendFactor BlendFactor
	DstColorBlendFactor BlendFactor
	ColorBlendOp        BlendOp
	SrcAlphaBlendFactor BlendFactor
	DstAlphaBlendFactor BlendFactor
	AlphaBlendOp        BlendOp
	ColorWriteMask      ColorComponentFlags
}

type PipelineColorBlendStateCreateInfo struct {
	Extensions     []Extension
	LogicOpEnable  bool
	LogicOp        LogicOp
	Attachments    []PipelineColorBlendAttachmentState
	BlendConstants [4]float32
}

func (info *PipelineColorBlendStateCreateInfo) c() *C.VkPipelineColorBlendStateCreateInfo {
	size0 := align(C.sizeof_VkPipelineColorBlendStateCreateInfo)
	size1 := align(C.sizeof_VkPipelineColorBlendAttachmentState * uintptr(len(info.Attachments)))
	size := size0 + size1
	mem := alloc(C.size_t(size))
	cinfo := (*C.VkPipelineColorBlendStateCreateInfo)(mem)
	attachments := (*C.VkPipelineColorBlendAttachmentState)(uptr(uintptr(mem) + size0))
	*cinfo = C.VkPipelineColorBlendStateCreateInfo{
		sType:           C.VkStructureType(StructureTypePipelineColorBlendStateCreateInfo),
		pNext:           buildChain(info.Extensions),
		flags:           0,
		logicOpEnable:   vkBool(info.LogicOpEnable),
		logicOp:         C.VkLogicOp(info.LogicOp),
		attachmentCount: C.uint32_t(len(info.Attachments)),
		pAttachments:    attachments,
		blendConstants: [4]C.float{
			C.float(info.BlendConstants[0]),
			C.float(info.BlendConstants[1]),
			C.float(info.BlendConstants[2]),
			C.float(info.BlendConstants[3]),
		},
	}
	attachmentsArr := (*[math.MaxInt32]C.VkPipelineColorBlendAttachmentState)(uptr(attachments))[:len(info.Attachments)]
	for i := range attachmentsArr {
		attachmentsArr[i] = C.VkPipelineColorBlendAttachmentState{
			blendEnable:         vkBool(info.Attachments[i].BlendEnable),
			srcColorBlendFactor: C.VkBlendFactor(info.Attachments[i].SrcColorBlendFactor),
			dstColorBlendFactor: C.VkBlendFactor(info.Attachments[i].DstColorBlendFactor),
			colorBlendOp:        C.VkBlendOp(info.Attachments[i].ColorBlendOp),
			srcAlphaBlendFactor: C.VkBlendFactor(info.Attachments[i].SrcAlphaBlendFactor),
			dstAlphaBlendFactor: C.VkBlendFactor(info.Attachments[i].DstAlphaBlendFactor),
			alphaBlendOp:        C.VkBlendOp(info.Attachments[i].AlphaBlendOp),
			colorWriteMask:      C.VkColorComponentFlags(info.Attachments[i].ColorWriteMask),
		}
	}
	return cinfo
}

type PipelineDynamicStateCreateInfo struct {
	Extensions    []Extension
	DynamicStates []DynamicState
}

func (info *PipelineDynamicStateCreateInfo) c() *C.VkPipelineDynamicStateCreateInfo {
	size0 := align(C.sizeof_VkPipelineDynamicStateCreateInfo)
	size1 := align(C.sizeof_VkDynamicState * uintptr(len(info.DynamicStates)))
	size := size0 + size1
	mem := alloc(C.size_t(size))
	cinfo := (*C.VkPipelineDynamicStateCreateInfo)(mem)
	dynamicStates := (*C.VkDynamicState)(uptr(uintptr(mem) + size0))
	*cinfo = C.VkPipelineDynamicStateCreateInfo{
		sType:             C.VkStructureType(StructureTypePipelineDynamicStateCreateInfo),
		pNext:             buildChain(info.Extensions),
		flags:             0,
		dynamicStateCount: C.uint32_t(len(info.DynamicStates)),
		pDynamicStates:    dynamicStates,
	}
	ucopy(uptr(dynamicStates), uptr(&info.DynamicStates), C.sizeof_VkDynamicState)
	return cinfo
}

type PipelineLayout struct {
	// VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkPipelineLayout)
	hnd C.VkPipelineLayout
}

type PipelineLayoutCreateInfo struct {
	Extensions         []Extension
	SetLayouts         []DescriptorSetLayout
	PushConstantRanges []PushConstantRange
}

func (info *PipelineLayoutCreateInfo) c() *C.VkPipelineLayoutCreateInfo {
	size0 := align(C.sizeof_VkPipelineLayoutCreateInfo)
	size1 := align(C.sizeof_VkDescriptorSetLayout * uintptr(len(info.SetLayouts)))
	size2 := align(C.sizeof_VkPushConstantRange * uintptr(len(info.PushConstantRanges)))
	size := size0 + size1 + size2
	mem := alloc(C.size_t(size))
	cinfo := (*C.VkPipelineLayoutCreateInfo)(mem)
	setLayouts := (*C.VkDescriptorSetLayout)(uptr(uintptr(mem) + size0))
	push := (*C.VkPushConstantRange)(uptr(uintptr(mem) + size0 + size1))
	*cinfo = C.VkPipelineLayoutCreateInfo{
		sType:                  C.VkStructureType(StructureTypePipelineLayoutCreateInfo),
		pNext:                  buildChain(info.Extensions),
		flags:                  0,
		setLayoutCount:         C.uint32_t(len(info.SetLayouts)),
		pSetLayouts:            setLayouts,
		pushConstantRangeCount: C.uint32_t(len(info.PushConstantRanges)),
		pPushConstantRanges:    push,
	}
	ucopy(uptr(setLayouts), uptr(&info.SetLayouts), C.sizeof_VkDescriptorSetLayout)
	ucopy(uptr(push), uptr(&info.PushConstantRanges), C.sizeof_VkPushConstantRange)
	return cinfo
}

func (dev *Device) CreatePipelineLayout(info *PipelineLayoutCreateInfo) (PipelineLayout, error) {
	// TODO(dh): support custom allocators
	cinfo := info.c()
	var out PipelineLayout
	res := Result(C.domVkCreatePipelineLayout(dev.fps[vkCreatePipelineLayout], dev.hnd, cinfo, nil, &out.hnd))
	internalizeChain(info.Extensions, cinfo.pNext)
	free(uptr(cinfo))
	return out, result2error(res)
}

func (dev *Device) DestroyPipelineLayout(layout PipelineLayout) {
	// TODO(dh): support custom allocators
	C.domVkDestroyPipelineLayout(dev.fps[vkDestroyPipelineLayout], dev.hnd, layout.hnd, nil)
}

type DescriptorSetLayout struct {
	// VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkDescriptorSetLayout)
	hnd C.VkDescriptorSetLayout

	// must be kept identical to C struct
}

type PushConstantRange struct {
	StageFlags ShaderStageFlags
	Offset     uint32
	Size       uint32

	// must be kept identical to C struct
}

type PipelineTessellationStateCreateInfo struct {
	Extensions         []Extension
	PatchControlPoints uint32
}

func (info *PipelineTessellationStateCreateInfo) c() *C.VkPipelineTessellationStateCreateInfo {
	cinfo := (*C.VkPipelineTessellationStateCreateInfo)(alloc(C.sizeof_VkPipelineTessellationStateCreateInfo))
	*cinfo = C.VkPipelineTessellationStateCreateInfo{
		sType:              C.VkStructureType(StructureTypePipelineTessellationStateCreateInfo),
		pNext:              buildChain(info.Extensions),
		flags:              0,
		patchControlPoints: C.uint32_t(info.PatchControlPoints),
	}
	return cinfo
}

type PipelineDepthStencilStateCreateInfo struct {
	Extensions            []Extension
	DepthTestEnable       bool
	DepthWriteEnable      bool
	DepthCompareOp        CompareOp
	DepthBoundsTestEnable bool
	StencilTestEnable     bool
	Front                 StencilOpState
	Back                  StencilOpState
	MinDepthBounds        float32
	MaxDepthBounds        float32
}

func (info *PipelineDepthStencilStateCreateInfo) c() *C.VkPipelineDepthStencilStateCreateInfo {
	cinfo := (*C.VkPipelineDepthStencilStateCreateInfo)(alloc(C.sizeof_VkPipelineDepthStencilStateCreateInfo))
	*cinfo = C.VkPipelineDepthStencilStateCreateInfo{
		sType:                 C.VkStructureType(StructureTypePipelineDepthStencilStateCreateInfo),
		pNext:                 buildChain(info.Extensions),
		flags:                 0,
		depthTestEnable:       vkBool(info.DepthTestEnable),
		depthWriteEnable:      vkBool(info.DepthWriteEnable),
		depthCompareOp:        C.VkCompareOp(info.DepthCompareOp),
		depthBoundsTestEnable: vkBool(info.DepthBoundsTestEnable),
		stencilTestEnable:     vkBool(info.StencilTestEnable),
		front:                 *(*C.VkStencilOpState)(uptr(&info.Front)),
		back:                  *(*C.VkStencilOpState)(uptr(&info.Back)),
		minDepthBounds:        C.float(info.MinDepthBounds),
		maxDepthBounds:        C.float(info.MaxDepthBounds),
	}
	return cinfo
}

type StencilOpState struct {
	FailOp      StencilOp
	PassOp      StencilOp
	DepthFailOp StencilOp
	CompareOp   CompareOp
	CompareMask uint32
	WriteMask   uint32
	Reference   uint32
}

type Pipeline struct {
	// VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkPipeline)
	hnd C.VkPipeline
}

type GraphicsPipelineCreateInfo struct {
	Extensions         []Extension
	Flags              PipelineCreateFlags
	Stages             []PipelineShaderStageCreateInfo
	VertexInputState   *PipelineVertexInputStateCreateInfo
	InputAssemblyState *PipelineInputAssemblyStateCreateInfo
	TessellationState  *PipelineTessellationStateCreateInfo
	ViewportState      *PipelineViewportStateCreateInfo
	RasterizationState *PipelineRasterizationStateCreateInfo
	MultisampleState   *PipelineMultisampleStateCreateInfo
	DepthStencilState  *PipelineDepthStencilStateCreateInfo
	ColorBlendState    *PipelineColorBlendStateCreateInfo
	DynamicState       *PipelineDynamicStateCreateInfo
	Layout             PipelineLayout
	RenderPass         RenderPass
	Subpass            uint32
	BasePipelineHandle *Pipeline
	BasePipelineIndex  int32
}

func (dev *Device) CreateGraphicsPipelines(infos []GraphicsPipelineCreateInfo) ([]Pipeline, error) {
	// TODO(dh): support pipeline cache
	// TODO(dh): support custom allocators
	ptrs := (*C.VkGraphicsPipelineCreateInfo)(allocn(len(infos), C.sizeof_VkGraphicsPipelineCreateInfo))
	defer free(uptr(ptrs))

	ptrsArr := (*[math.MaxInt32]C.VkGraphicsPipelineCreateInfo)(uptr(ptrs))[:len(infos)]
	for i := range ptrsArr {
		ptr := &ptrsArr[i]
		info := &infos[i]

		ptr.sType = C.VkStructureType(StructureTypeGraphicsPipelineCreateInfo)
		ptr.pNext = buildChain(info.Extensions)
		defer internalizeChain(info.Extensions, ptr.pNext)
		ptr.flags = C.VkPipelineCreateFlags(info.Flags)
		ptr.stageCount = C.uint32_t(len(info.Stages))

		ptr.pStages = (*C.VkPipelineShaderStageCreateInfo)(allocn(len(info.Stages), C.sizeof_VkPipelineShaderStageCreateInfo))
		defer free(uptr(ptr.pStages))
		arr := (*[math.MaxInt32]C.VkPipelineShaderStageCreateInfo)(uptr(ptr.pStages))[:len(info.Stages)]
		for i := range arr {
			arr[i] = C.VkPipelineShaderStageCreateInfo{
				sType:  C.VkStructureType(StructureTypePipelineShaderStageCreateInfo),
				pNext:  buildChain(info.Stages[i].Extensions),
				stage:  C.VkShaderStageFlagBits(info.Stages[i].Stage),
				module: info.Stages[i].Module.hnd,
				pName:  C.CString(info.Stages[i].Name),
			}
			defer free(uptr(arr[i].pName))
			defer internalizeChain(info.Stages[i].Extensions, arr[i].pNext)
		}

		if info.VertexInputState != nil {
			ptr.pVertexInputState = info.VertexInputState.c()
			defer free(uptr(ptr.pVertexInputState))
			defer internalizeChain(info.VertexInputState.Extensions, ptr.pVertexInputState.pNext)
		}
		if info.InputAssemblyState != nil {
			ptr.pInputAssemblyState = info.InputAssemblyState.c()
			defer free(uptr(ptr.pInputAssemblyState))
			defer internalizeChain(info.InputAssemblyState.Extensions, ptr.pInputAssemblyState.pNext)
		}
		if info.TessellationState != nil {
			ptr.pTessellationState = info.TessellationState.c()
			defer free(uptr(ptr.pTessellationState))
			defer internalizeChain(info.TessellationState.Extensions, ptr.pTessellationState.pNext)
		}
		if info.ViewportState != nil {
			ptr.pViewportState = info.ViewportState.c()
			defer free(uptr(ptr.pViewportState))
			defer internalizeChain(info.ViewportState.Extensions, ptr.pViewportState.pNext)
		}
		if info.RasterizationState != nil {
			ptr.pRasterizationState = info.RasterizationState.c()
			defer free(uptr(ptr.pRasterizationState))
			defer internalizeChain(info.RasterizationState.Extensions, ptr.pRasterizationState.pNext)
		}
		if info.MultisampleState != nil {
			ptr.pMultisampleState = info.MultisampleState.c()
			defer free(uptr(ptr.pMultisampleState))
			defer internalizeChain(info.MultisampleState.Extensions, ptr.pMultisampleState.pNext)
		}
		if info.DepthStencilState != nil {
			ptr.pDepthStencilState = info.DepthStencilState.c()
			defer free(uptr(ptr.pDepthStencilState))
			defer internalizeChain(info.DepthStencilState.Extensions, ptr.pDepthStencilState.pNext)
		}
		if info.ColorBlendState != nil {
			ptr.pColorBlendState = info.ColorBlendState.c()
			defer free(uptr(ptr.pColorBlendState))
			defer internalizeChain(info.ColorBlendState.Extensions, ptr.pColorBlendState.pNext)
		}
		if info.DynamicState != nil {
			ptr.pDynamicState = info.DynamicState.c()
			defer free(uptr(ptr.pDynamicState))
			defer internalizeChain(info.DynamicState.Extensions, ptr.pDynamicState.pNext)
		}
		ptr.layout = info.Layout.hnd
		ptr.renderPass = info.RenderPass.hnd
		ptr.subpass = C.uint32_t(info.Subpass)
		if info.BasePipelineHandle != nil {
			ptr.basePipelineHandle = info.BasePipelineHandle.hnd
		}
		ptr.basePipelineIndex = C.int32_t(info.BasePipelineIndex)
	}

	hnds := make([]C.VkPipeline, len(infos))
	res := Result(C.domVkCreateGraphicsPipelines(dev.fps[vkCreateGraphicsPipelines], dev.hnd, 0, C.uint32_t(len(infos)), ptrs, nil, (*C.VkPipeline)(slice2ptr(uptr(&hnds)))))
	if res != Success {
		return nil, res
	}
	out := make([]Pipeline, len(infos))
	for i, hnd := range hnds {
		out[i] = Pipeline{hnd}
	}
	return out, nil
}

type AttachmentDescription struct {
	Flags          AttachmentDescriptionFlags
	Format         Format
	Samples        SampleCountFlags
	LoadOp         AttachmentLoadOp
	StoreOp        AttachmentStoreOp
	StencilLoadOp  AttachmentLoadOp
	StencilStoreOp AttachmentStoreOp
	InitialLayout  ImageLayout
	FinalLayout    ImageLayout

	// must be kept identical to C struct
}

type AttachmentReference struct {
	Attachment uint32
	Layout     ImageLayout

	// must be kept identical to C struct
}

type SubpassDescription struct {
	Flags                  SubpassDescriptionFlags
	PipelineBindPoint      PipelineBindPoint
	InputAttachments       []AttachmentReference
	ColorAttachments       []AttachmentReference
	ResolveAttachments     []AttachmentReference
	DepthStencilAttachment *AttachmentReference
	PreserveAttachments    []uint32
}

type RenderPassCreateInfo struct {
	Extensions   []Extension
	Attachments  []AttachmentDescription
	Subpasses    []SubpassDescription
	Dependencies []SubpassDependency
}

type SubpassDependency struct {
	SrcSubpass      uint32
	DstSubpass      uint32
	SrcStageMask    PipelineStageFlags
	DstStageMask    PipelineStageFlags
	SrcAccessMask   AccessFlags
	DstAccessMask   AccessFlags
	DependencyFlags DependencyFlags

	// must be kept identical to C struct
}

func (dev *Device) CreateRenderPass(info *RenderPassCreateInfo) (RenderPass, error) {
	// TODO(dh): support custom allocators
	size0 := align(C.sizeof_VkRenderPassCreateInfo)
	size1 := align(C.sizeof_VkAttachmentDescription * uintptr(len(info.Attachments)))
	size2 := align(C.sizeof_VkSubpassDescription * uintptr(len(info.Subpasses)))
	size3 := align(C.sizeof_VkSubpassDependency * uintptr(len(info.Dependencies)))
	size := size0 + size1 + size2 + size3
	mem := alloc(C.size_t(size))
	defer free(mem)
	cinfo := (*C.VkRenderPassCreateInfo)(mem)
	attachments := (*C.VkAttachmentDescription)(uptr(uintptr(mem) + size0))
	subpasses := (*C.VkSubpassDescription)(uptr(uintptr(mem) + size0 + size1))
	dependencies := (*C.VkSubpassDependency)(uptr(uintptr(mem) + size0 + size1 + size2))
	*cinfo = C.VkRenderPassCreateInfo{
		sType:           C.VkStructureType(StructureTypeRenderPassCreateInfo),
		pNext:           buildChain(info.Extensions),
		flags:           0,
		attachmentCount: C.uint32_t(len(info.Attachments)),
		pAttachments:    attachments,
		subpassCount:    C.uint32_t(len(info.Subpasses)),
		pSubpasses:      subpasses,
		dependencyCount: C.uint32_t(len(info.Dependencies)),
		pDependencies:   dependencies,
	}
	defer internalizeChain(info.Extensions, cinfo.pNext)
	ucopy(uptr(attachments), uptr(&info.Attachments), C.sizeof_VkAttachmentDescription)
	subpassesArr := (*[math.MaxInt32]C.VkSubpassDescription)(uptr(subpasses))[:len(info.Subpasses)]
	for i := range subpassesArr {
		subpass := &info.Subpasses[i]
		csubpass := &subpassesArr[i]
		*csubpass = C.VkSubpassDescription{
			flags:                   C.VkSubpassDescriptionFlags(subpass.Flags),
			pipelineBindPoint:       C.VkPipelineBindPoint(subpass.PipelineBindPoint),
			inputAttachmentCount:    C.uint32_t(len(subpass.InputAttachments)),
			colorAttachmentCount:    C.uint32_t(len(subpass.ColorAttachments)),
			preserveAttachmentCount: C.uint32_t(len(subpass.PreserveAttachments)),
			pInputAttachments:       (*C.VkAttachmentReference)(allocn(len(subpass.InputAttachments), C.sizeof_VkAttachmentReference)),
			pColorAttachments:       (*C.VkAttachmentReference)(allocn(len(subpass.ColorAttachments), C.sizeof_VkAttachmentReference)),
			pPreserveAttachments:    (*C.uint32_t)(allocn(len(subpass.PreserveAttachments), C.sizeof_uint32_t)),
		}
		ucopy(uptr(csubpass.pInputAttachments), uptr(&subpass.InputAttachments), C.sizeof_VkAttachmentReference)
		ucopy(uptr(csubpass.pColorAttachments), uptr(&subpass.ColorAttachments), C.sizeof_VkAttachmentReference)
		if len(subpass.ResolveAttachments) > 0 {
			csubpass.pResolveAttachments = (*C.VkAttachmentReference)(allocn(len(subpass.ResolveAttachments), C.sizeof_VkAttachmentReference))
			defer free(uptr(csubpass.pResolveAttachments))
			ucopy(uptr(csubpass.pResolveAttachments), uptr(&subpass.ResolveAttachments), C.sizeof_VkAttachmentReference)
		}
		if subpass.DepthStencilAttachment != nil {
			csubpass.pDepthStencilAttachment = (*C.VkAttachmentReference)(alloc(C.sizeof_VkAttachmentReference))
			ucopy1(uptr(csubpass.pDepthStencilAttachment), uptr(subpass.DepthStencilAttachment), C.sizeof_VkAttachmentReference)
		}
		ucopy(uptr(csubpass.pPreserveAttachments), uptr(&subpass.PreserveAttachments), C.sizeof_uint32_t)
	}
	ucopy(uptr(dependencies), uptr(&info.Dependencies), C.sizeof_VkSubpassDependency)
	var out RenderPass
	res := Result(C.domVkCreateRenderPass(dev.fps[vkCreateRenderPass], dev.hnd, cinfo, nil, &out.hnd))
	return out, result2error(res)
}

type FramebufferCreateInfo struct {
	Extensions  []Extension
	RenderPass  RenderPass
	Attachments []ImageView
	Width       uint32
	Height      uint32
	Layers      uint32
}

func (info *FramebufferCreateInfo) c() *C.VkFramebufferCreateInfo {
	size0 := align(C.sizeof_VkFramebufferCreateInfo)
	size1 := align(uintptr(C.sizeof_VkImageView) * uintptr(len(info.Attachments)))
	size := size0 + size1
	mem := alloc(C.size_t(size))
	cinfo := (*C.VkFramebufferCreateInfo)(mem)
	*cinfo = C.VkFramebufferCreateInfo{
		sType:           C.VkStructureType(StructureTypeFramebufferCreateInfo),
		pNext:           buildChain(info.Extensions),
		flags:           0,
		renderPass:      info.RenderPass.hnd,
		attachmentCount: C.uint32_t(len(info.Attachments)),
		pAttachments:    (*C.VkImageView)(uptr(uintptr(mem) + size0)),
		width:           C.uint32_t(info.Width),
		height:          C.uint32_t(info.Height),
		layers:          C.uint32_t(info.Layers),
	}
	ucopy(uptr(cinfo.pAttachments), uptr(&info.Attachments), C.sizeof_VkImageView)
	return cinfo
}

func (dev *Device) CreateFramebuffer(info *FramebufferCreateInfo) (Framebuffer, error) {
	// TODO(dh): support custom allocators
	cinfo := info.c()
	var fb Framebuffer
	res := Result(C.domVkCreateFramebuffer(dev.fps[vkCreateFramebuffer], dev.hnd, cinfo, nil, &fb.hnd))
	internalizeChain(info.Extensions, cinfo.pNext)
	free(uptr(cinfo))
	return fb, result2error(res)
}

func (dev *Device) DestroyFramebuffer(fb Framebuffer) {
	// TODO(dh): support custom allocators
	C.domVkDestroyFramebuffer(dev.fps[vkDestroyFramebuffer], dev.hnd, fb.hnd, nil)
}

type RenderPassBeginInfo struct {
	Extensions  []Extension
	RenderPass  RenderPass
	Framebuffer Framebuffer
	RenderArea  Rect2D
	ClearValues []ClearValue
}

type ClearValue interface {
	isClearValue()
}

type ClearColorValue interface {
	isClearColorValue()
}

type ClearColorValueFloat32s [4]float32
type ClearColorValueInt32s [4]int32
type ClearColorValueUint32s [4]uint32

type ClearDepthStencilValue struct {
	Depth   float32
	Stencil uint32

	// must be kept identical to C struct
}

func (ClearColorValueFloat32s) isClearColorValue() {}
func (ClearColorValueInt32s) isClearColorValue()   {}
func (ClearColorValueUint32s) isClearColorValue()  {}

func (ClearColorValueFloat32s) isClearValue() {}
func (ClearColorValueInt32s) isClearValue()   {}
func (ClearColorValueUint32s) isClearValue()  {}
func (ClearDepthStencilValue) isClearValue()  {}

type Semaphore struct {
	// VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkSemaphore)
	hnd C.VkSemaphore

	// must be kept identical to C struct
}

type SemaphoreCreateInfo struct {
	Extensions []Extension
}

func (info *SemaphoreCreateInfo) c() *C.VkSemaphoreCreateInfo {
	cinfo := (*C.VkSemaphoreCreateInfo)(alloc(C.sizeof_VkSemaphoreCreateInfo))
	*cinfo = C.VkSemaphoreCreateInfo{
		sType: C.VkStructureType(StructureTypeSemaphoreCreateInfo),
		pNext: buildChain(info.Extensions),
	}
	return cinfo
}

func (dev *Device) CreateSemaphore(info *SemaphoreCreateInfo) (Semaphore, error) {
	// TODO(dh): support custom allocators
	cinfo := info.c()
	var sem Semaphore
	res := Result(C.domVkCreateSemaphore(dev.fps[vkCreateSemaphore], dev.hnd, cinfo, nil, &sem.hnd))
	internalizeChain(info.Extensions, cinfo.pNext)
	free(uptr(cinfo))
	return sem, result2error(res)
}

func (dev *Device) DestroySemaphore(sem Semaphore) {
	// TODO(dh): support custom allocators
	C.domVkDestroySemaphore(dev.fps[vkDestroySemaphore], dev.hnd, sem.hnd, nil)
}

type SubmitInfo struct {
	Extensions       []Extension
	WaitSemaphores   []Semaphore
	WaitDstStageMask []PipelineStageFlags
	CommandBuffers   []*CommandBuffer
	SignalSemaphores []Semaphore
}

func (queue *Queue) Submit(infos []SubmitInfo, fence *Fence) error {
	var (
		waitSemaphoreCount   uintptr
		commandBufferCount   uintptr
		signalSemaphoreCount uintptr
	)
	for _, info := range infos {
		waitSemaphoreCount += uintptr(len(info.WaitSemaphores))
		commandBufferCount += uintptr(len(info.CommandBuffers))
		signalSemaphoreCount += uintptr(len(info.SignalSemaphores))
	}
	size0 := align(C.sizeof_VkSubmitInfo * uintptr(len(infos)))
	size1 := align(C.sizeof_VkSemaphore * waitSemaphoreCount)
	size2 := align(C.sizeof_VkPipelineStageFlags * waitSemaphoreCount)
	size3 := align(C.sizeof_VkCommandBuffer * commandBufferCount)
	size4 := align(C.sizeof_VkSemaphore * signalSemaphoreCount)
	size := size0 + size1 + size2 + size3 + size4
	mem := uintptr(alloc(C.size_t(size)))
	defer free(uptr(mem))

	cinfos := mem
	waitSemaphores := mem + size0
	waitDstStageMask := mem + size0 + size1
	commandBuffers := mem + size0 + size1 + size2
	signalSemaphores := mem + size0 + size1 + size2 + size3

	for _, info := range infos {
		if len(info.WaitSemaphores) != len(info.WaitDstStageMask) {
			panic("WaitSemaphores and WaitDstStageMask must have same length")
		}
		*(*C.VkSubmitInfo)(uptr(cinfos)) = C.VkSubmitInfo{
			sType:                C.VkStructureType(StructureTypeSubmitInfo),
			pNext:                buildChain(info.Extensions),
			waitSemaphoreCount:   C.uint32_t(len(info.WaitSemaphores)),
			pWaitSemaphores:      (*C.VkSemaphore)(uptr(waitSemaphores)),
			pWaitDstStageMask:    (*C.VkPipelineStageFlags)(uptr(waitDstStageMask)),
			commandBufferCount:   C.uint32_t(len(info.CommandBuffers)),
			pCommandBuffers:      (*C.VkCommandBuffer)(uptr(commandBuffers)),
			signalSemaphoreCount: C.uint32_t(len(info.SignalSemaphores)),
			pSignalSemaphores:    (*C.VkSemaphore)(uptr(signalSemaphores)),
		}
		defer internalizeChain(info.Extensions, (*C.VkSubmitInfo)(uptr(cinfos)).pNext)
		ucopy(uptr(waitSemaphores), uptr(&info.WaitSemaphores), C.sizeof_VkSemaphore)
		ucopy(uptr(waitDstStageMask), uptr(&info.WaitDstStageMask), C.sizeof_VkPipelineStageFlags)
		ucopy(uptr(signalSemaphores), uptr(&info.SignalSemaphores), C.sizeof_VkSemaphore)
		arr := (*[math.MaxInt32]C.VkCommandBuffer)(uptr(commandBuffers))[:len(info.CommandBuffers)]
		for i := range arr {
			arr[i] = info.CommandBuffers[i].hnd
		}

		cinfos += C.sizeof_VkSubmitInfo
		waitSemaphores += C.sizeof_VkSemaphore * uintptr(len(info.WaitSemaphores))
		waitDstStageMask += C.sizeof_VkPipelineStageFlags * uintptr(len(info.WaitSemaphores))
		commandBuffers += C.sizeof_VkCommandBuffer * uintptr(len(info.CommandBuffers))
		signalSemaphores += C.sizeof_VkSemaphore * uintptr(len(info.SignalSemaphores))
	}

	var fenceHnd C.VkFence
	if fence != nil {
		fenceHnd = fence.hnd
	}
	res := Result(C.domVkQueueSubmit(queue.fps[vkQueueSubmit], queue.hnd, C.uint32_t(len(infos)), (*C.VkSubmitInfo)(uptr(mem)), fenceHnd))
	return result2error(res)
}

type Fence struct {
	// VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkFence)
	hnd C.VkFence

	// must be kept identical to C struct
}

type FenceCreateInfo struct {
	Extensions []Extension
	Flags      FenceCreateFlags
}

func (dev *Device) CreateFence(info *FenceCreateInfo) (Fence, error) {
	// TODO(dh): support custom allocators
	cinfo := (*C.VkFenceCreateInfo)(alloc(C.sizeof_VkFenceCreateInfo))
	*cinfo = C.VkFenceCreateInfo{
		sType: C.VkStructureType(StructureTypeFenceCreateInfo),
		pNext: buildChain(info.Extensions),
		flags: C.VkFenceCreateFlags(info.Flags),
	}
	var fence Fence
	res := Result(C.domVkCreateFence(dev.fps[vkCreateFence], dev.hnd, cinfo, nil, &fence.hnd))
	internalizeChain(info.Extensions, cinfo.pNext)
	free(uptr(cinfo))
	return fence, result2error(res)
}

func (dev *Device) DestroyFence(fence Fence) {
	// TODO(dh): support custom allocators
	C.domVkDestroyFence(dev.fps[vkDestroyFence], dev.hnd, fence.hnd, nil)
}

func (dev *Device) FenceStatus(fence Fence) (Result, error) {
	res := Result(C.domVkGetFenceStatus(dev.fps[vkGetFenceStatus], dev.hnd, fence.hnd))
	switch res {
	case Success, NotReady:
		return res, nil
	default:
		return res, res
	}
}

func (dev *Device) WaitForFences(fences []Fence, waitAll bool, timeout time.Duration) error {
	res := Result(C.domVkWaitForFences(dev.fps[vkWaitForFences], dev.hnd, C.uint32_t(len(fences)), (*C.VkFence)(slice2ptr(uptr(&fences))), vkBool(waitAll), C.uint64_t(timeout)))
	return result2error(res)
}

func (dev *Device) ResetFences(fences []Fence) error {
	res := Result(C.domVkResetFences(dev.fps[vkResetFences], dev.hnd, C.uint32_t(len(fences)), (*C.VkFence)(slice2ptr(uptr(&fences)))))
	return result2error(res)
}

type BufferCreateInfo struct {
	Extensions         []Extension
	Flags              BufferCreateFlags
	Size               DeviceSize
	Usage              BufferUsageFlags
	SharingMode        SharingMode
	QueueFamilyIndices []uint32
}

type Buffer struct {
	// VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkBuffer)
	hnd C.VkBuffer
}

func (info *BufferCreateInfo) c() *C.VkBufferCreateInfo {
	size0 := align(C.sizeof_VkBufferCreateInfo)
	size1 := align(C.sizeof_uint32_t * uintptr(len(info.QueueFamilyIndices)))
	size := size0 + size1
	mem := alloc(C.size_t(size))
	cinfo := (*C.VkBufferCreateInfo)(mem)
	*cinfo = C.VkBufferCreateInfo{
		sType:                 C.VkStructureType(StructureTypeBufferCreateInfo),
		pNext:                 buildChain(info.Extensions),
		flags:                 C.VkBufferCreateFlags(info.Flags),
		size:                  C.VkDeviceSize(info.Size),
		usage:                 C.VkBufferUsageFlags(info.Usage),
		sharingMode:           C.VkSharingMode(info.SharingMode),
		queueFamilyIndexCount: C.uint32_t(len(info.QueueFamilyIndices)),
		pQueueFamilyIndices:   (*C.uint32_t)(uptr(uintptr(mem) + size0)),
	}
	ucopy(uptr(cinfo.pQueueFamilyIndices), uptr(&info.QueueFamilyIndices), C.sizeof_uint32_t)
	return cinfo
}

func (dev *Device) CreateBuffer(info *BufferCreateInfo) (Buffer, error) {
	// TODO(dh): support custom allocators
	cinfo := info.c()
	var buf Buffer
	res := Result(C.domVkCreateBuffer(dev.fps[vkCreateBuffer], dev.hnd, cinfo, nil, &buf.hnd))
	internalizeChain(info.Extensions, cinfo.pNext)
	free(uptr(cinfo))
	return buf, result2error(res)
}

func (dev *Device) DestroyBuffer(buf Buffer) {
	// TODO(dh): support custom allocators
	C.domVkDestroyBuffer(dev.fps[vkDestroyBuffer], dev.hnd, buf.hnd, nil)
}

type MemoryRequirements struct {
	Size           DeviceSize
	Alignment      DeviceSize
	MemoryTypeBits uint32

	// must be kept identical to C struct
}

func (dev *Device) BufferMemoryRequirements(buf Buffer) MemoryRequirements {
	var reqs MemoryRequirements
	C.domVkGetBufferMemoryRequirements(dev.fps[vkGetBufferMemoryRequirements], dev.hnd, buf.hnd, (*C.VkMemoryRequirements)(uptr(&reqs)))
	return reqs
}

type MemoryAllocateInfo struct {
	Extensions      []Extension
	AllocationSize  DeviceSize
	MemoryTypeIndex uint32
}

func (info *MemoryAllocateInfo) c() *C.VkMemoryAllocateInfo {
	cinfo := (*C.VkMemoryAllocateInfo)(alloc(C.sizeof_VkMemoryAllocateInfo))
	*cinfo = C.VkMemoryAllocateInfo{
		sType:           C.VkStructureType(StructureTypeMemoryAllocateInfo),
		pNext:           buildChain(info.Extensions),
		allocationSize:  C.VkDeviceSize(info.AllocationSize),
		memoryTypeIndex: C.uint32_t(info.MemoryTypeIndex),
	}
	return cinfo
}

// DeviceMemory is an opaque handle to a device memory object.
type DeviceMemory struct {
	// VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkDeviceMemory)
	hnd C.VkDeviceMemory
}

func (dev *Device) AllocateMemory(info *MemoryAllocateInfo) (DeviceMemory, error) {
	// TODO(dh): support custom allocators
	cinfo := info.c()
	var mem DeviceMemory
	res := Result(C.domVkAllocateMemory(dev.fps[vkAllocateMemory], dev.hnd, cinfo, nil, &mem.hnd))
	internalizeChain(info.Extensions, cinfo.pNext)
	free(uptr(cinfo))
	return mem, result2error(res)
}

func (dev *Device) FreeMemory(mem DeviceMemory) {
	// TODO(dh): support custom allocators
	C.domVkFreeMemory(dev.fps[vkFreeMemory], dev.hnd, mem.hnd, nil)
}

func (dev *Device) BindBufferMemory(buf Buffer, mem DeviceMemory, offset DeviceSize) error {
	res := Result(C.domVkBindBufferMemory(dev.fps[vkBindBufferMemory], dev.hnd, buf.hnd, mem.hnd, C.VkDeviceSize(offset)))
	return result2error(res)
}

type BindBufferMemoryInfo struct {
	Extensions   []Extension
	Buffer       Buffer
	Memory       DeviceMemory
	MemoryOffset DeviceSize
}

func (dev *Device) BindBufferMemory2(infos []BindBufferMemoryInfo) error {
	mem := allocn(len(infos), C.sizeof_VkBindBufferMemoryInfo)
	defer free(mem)
	cinfos := (*[1 << 31]C.VkBindBufferMemoryInfo)(mem)[:len(infos)]
	for i, info := range infos {
		cinfos[i] = C.VkBindBufferMemoryInfo{
			sType:        C.VkStructureType(StructureTypeBindBufferMemoryInfo),
			pNext:        buildChain(info.Extensions),
			buffer:       info.Buffer.hnd,
			memory:       info.Memory.hnd,
			memoryOffset: C.VkDeviceSize(info.MemoryOffset),
		}
		defer internalizeChain(info.Extensions, cinfos[i].pNext)
	}
	res := Result(C.domVkBindBufferMemory2(dev.fps[vkBindBufferMemory2], dev.hnd, C.uint32_t(len(infos)), (*C.VkBindBufferMemoryInfo)(mem)))
	return result2error(res)
}

func (dev *Device) MapMemory(mem DeviceMemory, offset, size DeviceSize, flags MemoryMapFlags) (uintptr, error) {
	var ptr uptr
	res := Result(C.domVkMapMemory(dev.fps[vkMapMemory], dev.hnd, mem.hnd, C.VkDeviceSize(offset), C.VkDeviceSize(size), C.VkMemoryMapFlags(flags), &ptr))
	return uintptr(ptr), result2error(res)
}

func (dev *Device) UnmapMemory(mem DeviceMemory) {
	C.domVkUnmapMemory(dev.fps[vkUnmapMemory], dev.hnd, mem.hnd)
}

type ImageCreateInfo struct {
	Extensions         []Extension
	Flags              ImageCreateFlags
	ImageType          ImageType
	Format             Format
	Extent             Extent3D
	MipLevels          uint32
	ArrayLayers        uint32
	Samples            SampleCountFlags
	Tiling             ImageTiling
	Usage              ImageUsageFlags
	SharingMode        SharingMode
	QueueFamilyIndices []uint32
	InitialLayout      ImageLayout
}

func (info *ImageCreateInfo) c() *C.VkImageCreateInfo {
	size0 := align(C.sizeof_VkImageCreateInfo)
	size1 := align(C.sizeof_uint32_t * uintptr(len(info.QueueFamilyIndices)))
	size := size0 + size1
	mem := alloc(C.size_t(size))
	cinfo := (*C.VkImageCreateInfo)(mem)
	*cinfo = C.VkImageCreateInfo{
		sType:                 C.VkStructureType(StructureTypeImageCreateInfo),
		pNext:                 buildChain(info.Extensions),
		flags:                 C.VkImageCreateFlags(info.Flags),
		imageType:             C.VkImageType(info.ImageType),
		format:                C.VkFormat(info.Format),
		mipLevels:             C.uint32_t(info.MipLevels),
		arrayLayers:           C.uint32_t(info.ArrayLayers),
		samples:               C.VkSampleCountFlagBits(info.Samples),
		tiling:                C.VkImageTiling(info.Tiling),
		usage:                 C.VkImageUsageFlags(info.Usage),
		sharingMode:           C.VkSharingMode(info.SharingMode),
		queueFamilyIndexCount: C.uint32_t(len(info.QueueFamilyIndices)),
		pQueueFamilyIndices:   (*C.uint32_t)(uptr(uintptr(mem) + size0)),
		initialLayout:         C.VkImageLayout(info.InitialLayout),
	}
	ucopy(uptr(cinfo.pQueueFamilyIndices), uptr(&info.QueueFamilyIndices), C.sizeof_uint32_t)
	ucopy1(uptr(&cinfo.extent), uptr(&info.Extent), C.sizeof_VkExtent3D)
	return cinfo
}

func (dev *Device) CreateImage(info *ImageCreateInfo) (Image, error) {
	// TODO(dh): support custom allocators
	cinfo := info.c()
	var img Image
	res := Result(C.domVkCreateImage(dev.fps[vkCreateImage], dev.hnd, cinfo, nil, &img.hnd))
	internalizeChain(info.Extensions, cinfo.pNext)
	free(uptr(cinfo))
	return img, result2error(res)
}

func (dev *Device) DestroyImage(img Image) {
	// TODO(dh): support custom allocators
	C.domVkDestroyImage(dev.fps[vkDestroyImage], dev.hnd, img.hnd, nil)
}

type Event struct {
	// VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkEvent)
	hnd C.VkEvent
}

type EventCreateInfo struct {
	Extensions []Extension
}

func (info *EventCreateInfo) c() *C.VkEventCreateInfo {
	cinfo := (*C.VkEventCreateInfo)(alloc(C.sizeof_VkEventCreateInfo))
	*cinfo = C.VkEventCreateInfo{
		sType: C.VkStructureType(StructureTypeEventCreateInfo),
		pNext: buildChain(info.Extensions),
		flags: 0,
	}
	return cinfo
}

func (dev *Device) CreateEvent(info *EventCreateInfo) (Event, error) {
	// TODO(dh): support custom allocators
	cinfo := info.c()
	var ev Event
	res := Result(C.domVkCreateEvent(dev.fps[vkCreateEvent], dev.hnd, cinfo, nil, &ev.hnd))
	internalizeChain(info.Extensions, cinfo.pNext)
	free(uptr(cinfo))
	return ev, result2error(res)
}

func (dev *Device) DestroyEvent(ev Event) {
	// TODO(dh): support custom allocators
	C.domVkDestroyEvent(dev.fps[vkDestroyEvent], dev.hnd, ev.hnd, nil)
}

func (dev *Device) SetEvent(ev Event) error {
	res := Result(C.domVkSetEvent(dev.fps[vkSetEvent], dev.hnd, ev.hnd))
	return result2error(res)
}

func (dev *Device) ResetEvent(ev Event) error {
	res := Result(C.domVkResetEvent(dev.fps[vkResetEvent], dev.hnd, ev.hnd))
	return result2error(res)
}

func (dev *Device) EventStatus(ev Event) (Result, error) {
	res := Result(C.domVkGetEventStatus(dev.fps[vkGetEventStatus], dev.hnd, ev.hnd))
	switch res {
	case EventSet, EventReset:
		return res, nil
	default:
		return res, res
	}
}

func vkGetInstanceProcAddr(instance C.VkInstance, name string) C.PFN_vkVoidFunction {
	// TODO(dh): return a mock function pointer that panics with a nice message

	cName := C.CString(name)
	fp := C.vkGetInstanceProcAddr(instance, cName)
	if debug {
		fmt.Fprintf(os.Stderr, "%s = %p\n", name, fp)
	}
	free(uptr(cName))
	return fp
}

func mustVkGetInstanceProcAddr(instance C.VkInstance, name string) C.PFN_vkVoidFunction {
	fp := vkGetInstanceProcAddr(instance, name)
	if fp == nil {
		panic(fmt.Sprintf("couldn't load function %s", name))
	}
	return fp
}

func (hnd *CommandBuffer) String() string  { return fmt.Sprintf("VkCommandBuffer(%#x)", hnd.hnd) }
func (hnd *Device) String() string         { return fmt.Sprintf("VkDevice(%#x)", hnd.hnd) }
func (hnd *Instance) String() string       { return fmt.Sprintf("VkInstance(%#x)", hnd.hnd) }
func (hnd *PhysicalDevice) String() string { return fmt.Sprintf("VkPhysicalDevice(%#x)", hnd.hnd) }
func (hnd *Queue) String() string          { return fmt.Sprintf("VkQueue(%#x)", hnd.hnd) }
func (hnd Buffer) String() string          { return fmt.Sprintf("VkBuffer(%#x)", hnd.hnd) }
func (hnd CommandPool) String() string     { return fmt.Sprintf("VkCommandPool(%#x)", hnd.hnd) }
func (hnd DeviceMemory) String() string    { return fmt.Sprintf("VkDeviceMemory(%#x)", hnd.hnd) }
func (hnd Fence) String() string           { return fmt.Sprintf("VkFence(%#x)", hnd.hnd) }
func (hnd Framebuffer) String() string     { return fmt.Sprintf("VkFramebuffer(%#x)", hnd.hnd) }
func (hnd Image) String() string           { return fmt.Sprintf("VkImage(%#x)", hnd.hnd) }
func (hnd ImageView) String() string       { return fmt.Sprintf("VkImageView(%#x)", hnd.hnd) }
func (hnd Pipeline) String() string        { return fmt.Sprintf("VkPipeline(%#x)", hnd.hnd) }
func (hnd PipelineLayout) String() string  { return fmt.Sprintf("VkPipelineLayout(%#x)", hnd.hnd) }
func (hnd RenderPass) String() string      { return fmt.Sprintf("VkRenderPass(%#x)", hnd.hnd) }
func (hnd Semaphore) String() string       { return fmt.Sprintf("VkSemaphore(%#x)", hnd.hnd) }
func (hnd ShaderModule) String() string    { return fmt.Sprintf("VkShaderModule(%#x)", hnd.hnd) }
func (hnd DescriptorSetLayout) String() string {
	return fmt.Sprintf("VkDescriptorSetLayout(%#x)", hnd.hnd)
}

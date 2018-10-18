// Copyright (c) 2015-2018 Khronos Group
// Copyright (c) 2018 Dominik Honnef

package vk

// #cgo pkg-config: vulkan
// #cgo CFLAGS: -DVK_NO_PROTOTYPES
//
// #include <stdlib.h>
//
// #define VK_DEFINE_NON_DISPATCHABLE_HANDLE(object) typedef uintptr_t object;
//
// #include <vulkan/vulkan_core.h>
//
// VKAPI_ATTR PFN_vkVoidFunction VKAPI_CALL vkGetInstanceProcAddr(VkInstance instance, const char *pName);
//
// VkResult domVkCreateInstance(PFN_vkCreateInstance fp, const VkInstanceCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkInstance* pInstance);
// VkResult domVkEnumeratePhysicalDevices(PFN_vkEnumeratePhysicalDevices fp, VkInstance instance, uint32_t* pPhysicalDeviceCount, VkPhysicalDevice* pPhysicalDevices);
// void     domVkGetPhysicalDeviceProperties(PFN_vkGetPhysicalDeviceProperties fp, VkPhysicalDevice physicalDevice, VkPhysicalDeviceProperties* pProperties);
// void     domVkGetPhysicalDeviceFeatures(PFN_vkGetPhysicalDeviceFeatures fp, VkPhysicalDevice physicalDevice, VkPhysicalDeviceFeatures* pFeatures);
// void     domVkGetPhysicalDeviceQueueFamilyProperties(PFN_vkGetPhysicalDeviceQueueFamilyProperties fp, VkPhysicalDevice physicalDevice, uint32_t* pQueueFamilyPropertyCount, VkQueueFamilyProperties* pQueueFamilyProperties);
// VkResult domVkCreateDevice(PFN_vkCreateDevice fp, VkPhysicalDevice physicalDevice, const VkDeviceCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkDevice* pDevice);
// void     domVkGetDeviceQueue(PFN_vkGetDeviceQueue fp, VkDevice device, uint32_t queueFamilyIndex, uint32_t queueIndex, VkQueue* pQueue);
// PFN_vkVoidFunction domVkGetDeviceProcAddr(PFN_vkGetDeviceProcAddr fp, VkDevice device, const char* pName);
// VkResult domVkCreateCommandPool(PFN_vkCreateCommandPool fp, VkDevice device, const VkCommandPoolCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkCommandPool* pCommandPool);
// void     domVkTrimCommandPool(PFN_vkTrimCommandPool fp, VkDevice device, VkCommandPool commandPool, VkCommandPoolTrimFlags flags);
// VkResult domVkResetCommandPool(PFN_vkResetCommandPool fp, VkDevice device, VkCommandPool commandPool, VkCommandPoolResetFlags flags);
// VkResult domVkAllocateCommandBuffers(PFN_vkAllocateCommandBuffers fp, VkDevice device, const VkCommandBufferAllocateInfo* pAllocateInfo, VkCommandBuffer* pCommandBuffers);
// VkResult domVkResetCommandBuffer(PFN_vkResetCommandBuffer fp, VkCommandBuffer commandBuffer, VkCommandBufferResetFlags flags);
// void     domVkFreeCommandBuffers(PFN_vkFreeCommandBuffers fp, VkDevice device, VkCommandPool commandPool, uint32_t commandBufferCount, const VkCommandBuffer* pCommandBuffers);
// VkResult domVkEndCommandBuffer(PFN_vkEndCommandBuffer fp, VkCommandBuffer commandBuffer);
// VkResult domVkBeginCommandBuffer(PFN_vkBeginCommandBuffer fp, VkCommandBuffer commandBuffer, const VkCommandBufferBeginInfo* pBeginInfo);
import "C"
import (
	"fmt"
	"os"
	"unsafe"
)

const debug = true

type (
	DeviceSize       = uint64
	SampleCountFlags = uint32
)

var vkEnumerateInstanceVersion C.PFN_vkEnumerateInstanceVersion
var vkEnumerateInstanceExtensionProperties C.PFN_vkEnumerateInstanceExtensionProperties
var vkEnumerateInstanceLayerProperties C.PFN_vkEnumerateInstanceLayerProperties
var vkCreateInstance C.PFN_vkCreateInstance

func init() {
	vkEnumerateInstanceVersion =
		C.PFN_vkEnumerateInstanceVersion(mustVkGetInstanceProcAddr(nil, "vkEnumerateInstanceVersion"))
	vkEnumerateInstanceExtensionProperties =
		C.PFN_vkEnumerateInstanceExtensionProperties(mustVkGetInstanceProcAddr(nil, "vkEnumerateInstanceExtensionProperties"))
	vkEnumerateInstanceLayerProperties =
		C.PFN_vkEnumerateInstanceLayerProperties(mustVkGetInstanceProcAddr(nil, "vkEnumerateInstanceLayerProperties"))
	vkCreateInstance =
		C.PFN_vkCreateInstance(mustVkGetInstanceProcAddr(nil, "vkCreateInstance"))
}

type InstanceCreateInfo struct {
	Next                  unsafe.Pointer
	ApplicationInfo       *ApplicationInfo
	EnabledLayerNames     []string
	EnabledExtensionNames []string
}

type ApplicationInfo struct {
	Next               unsafe.Pointer
	ApplicationName    string
	ApplicationVersion uint32
	EngineName         string
	EngineVersion      uint32
	APIVersion         uint32
}

func externStrings(ss []string) (**C.char, func()) {
	var ptrs []unsafe.Pointer

	ptr := C.calloc(C.size_t(len(ss)), C.size_t(unsafe.Sizeof(uintptr(0))))
	ptrs = append(ptrs, ptr)
	slice := (*[1 << 31]*C.char)(ptr)[:len(ss)]
	for i, s := range ss {
		slice[i] = C.CString(s)
		ptrs = append(ptrs, unsafe.Pointer(slice[i]))
	}
	return (**C.char)(ptr), func() {
		for _, ptr := range ptrs {
			C.free(ptr)
		}
	}
}

func CreateInstance(info *InstanceCreateInfo) (*Instance, error) {
	// TODO(dh): support a custom allocator
	var free1, free2 func()

	ptr := (*C.VkInstanceCreateInfo)(C.calloc(1, C.size_t(unsafe.Sizeof(C.VkInstanceCreateInfo{}))))
	ptr.sType = C.VkStructureType(StructureTypeInstanceCreateInfo)
	ptr.pNext = info.Next
	ptr.enabledLayerCount = C.uint32_t(len(info.EnabledLayerNames))
	ptr.enabledExtensionCount = C.uint32_t(len(info.EnabledExtensionNames))
	ptr.ppEnabledLayerNames, free1 = externStrings(info.EnabledLayerNames)
	ptr.ppEnabledExtensionNames, free2 = externStrings(info.EnabledExtensionNames)
	defer C.free(unsafe.Pointer(ptr))
	defer free1()
	defer free2()
	if info.ApplicationInfo != nil {
		ptr.pApplicationInfo = (*C.VkApplicationInfo)(C.calloc(1, C.size_t(unsafe.Sizeof(C.VkApplicationInfo{}))))
		ptr.pApplicationInfo.sType = C.VkStructureType(StructureTypeApplicationInfo)
		ptr.pApplicationInfo.pNext = info.ApplicationInfo.Next
		ptr.pApplicationInfo.pApplicationName = C.CString(info.ApplicationInfo.ApplicationName)
		ptr.pApplicationInfo.applicationVersion = C.uint32_t(info.ApplicationInfo.ApplicationVersion)
		ptr.pApplicationInfo.pEngineName = C.CString(info.ApplicationInfo.EngineName)
		ptr.pApplicationInfo.engineVersion = C.uint32_t(info.ApplicationInfo.EngineVersion)
		ptr.pApplicationInfo.apiVersion = C.uint32_t(info.ApplicationInfo.APIVersion)
		defer C.free(unsafe.Pointer(ptr.pApplicationInfo))
		defer C.free(unsafe.Pointer(ptr.pApplicationInfo.pApplicationName))
		defer C.free(unsafe.Pointer(ptr.pApplicationInfo.pEngineName))
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
	hnd C.VkInstance
	fps [instanceMaxPFN]C.PFN_vkVoidFunction
}

func (ins *Instance) init() {
	for i, name := range instanceFpNames {
		ins.fps[i] = vkGetInstanceProcAddr(ins.hnd, name)
	}
}

func (ins *Instance) String() string {
	return fmt.Sprintf("VkInstance(%p)", ins.hnd)
}

func (ins *Instance) EnumeratePhysicalDevices() ([]*PhysicalDevice, error) {
	count := C.uint32_t(1)
	var devs *C.VkPhysicalDevice
	for {
		devs = (*C.VkPhysicalDevice)(C.calloc(C.size_t(count), C.size_t(unsafe.Sizeof(C.VkPhysicalDevice(nil)))))
		defer C.free(unsafe.Pointer(devs))
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
	for _, dev := range (*[1 << 31]C.VkPhysicalDevice)(unsafe.Pointer(devs))[:count] {
		out = append(out, &PhysicalDevice{dev, ins})
	}
	return out, nil
}

type PhysicalDevice struct {
	hnd      C.VkPhysicalDevice
	instance *Instance
}

func (dev *PhysicalDevice) String() string {
	return fmt.Sprintf("VkPhysicalDevice(%p)", dev.hnd)
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

// Properties returns general properties of the physical device.
func (dev *PhysicalDevice) Properties() *PhysicalDeviceProperties {
	var props C.VkPhysicalDeviceProperties
	C.domVkGetPhysicalDeviceProperties(dev.instance.fps[vkGetPhysicalDeviceProperties], dev.hnd, &props)
	return &PhysicalDeviceProperties{
		APIVersion:        uint32(props.apiVersion),
		DriverVersion:     uint32(props.driverVersion),
		VendorID:          uint32(props.vendorID),
		DeviceID:          uint32(props.deviceID),
		DeviceType:        PhysicalDeviceType(props.deviceType),
		DeviceName:        C.GoString(&props.deviceName[0]),
		PipelineCacheUUID: (*[C.VK_UUID_SIZE]byte)(unsafe.Pointer(&props.pipelineCacheUUID))[:],
		Limits: PhysicalDeviceLimits{
			MaxImageDimension1D:                             uint32(props.limits.maxImageDimension1D),
			MaxImageDimension2D:                             uint32(props.limits.maxImageDimension2D),
			MaxImageDimension3D:                             uint32(props.limits.maxImageDimension3D),
			MaxImageDimensionCube:                           uint32(props.limits.maxImageDimensionCube),
			MaxImageArrayLayers:                             uint32(props.limits.maxImageArrayLayers),
			MaxTexelBufferElements:                          uint32(props.limits.maxTexelBufferElements),
			MaxUniformBufferRange:                           uint32(props.limits.maxUniformBufferRange),
			MaxStorageBufferRange:                           uint32(props.limits.maxStorageBufferRange),
			MaxPushConstantsSize:                            uint32(props.limits.maxPushConstantsSize),
			MaxMemoryAllocationCount:                        uint32(props.limits.maxMemoryAllocationCount),
			MaxSamplerAllocationCount:                       uint32(props.limits.maxSamplerAllocationCount),
			BufferImageGranularity:                          DeviceSize(props.limits.bufferImageGranularity),
			SparseAddressSpaceSize:                          DeviceSize(props.limits.sparseAddressSpaceSize),
			MaxBoundDescriptorSets:                          uint32(props.limits.maxBoundDescriptorSets),
			MaxPerStageDescriptorSamplers:                   uint32(props.limits.maxPerStageDescriptorSamplers),
			MaxPerStageDescriptorUniformBuffers:             uint32(props.limits.maxPerStageDescriptorUniformBuffers),
			MaxPerStageDescriptorStorageBuffers:             uint32(props.limits.maxPerStageDescriptorStorageBuffers),
			MaxPerStageDescriptorSampledImages:              uint32(props.limits.maxPerStageDescriptorSampledImages),
			MaxPerStageDescriptorStorageImages:              uint32(props.limits.maxPerStageDescriptorStorageImages),
			MaxPerStageDescriptorInputAttachments:           uint32(props.limits.maxPerStageDescriptorInputAttachments),
			MaxPerStageResources:                            uint32(props.limits.maxPerStageResources),
			MaxDescriptorSetSamplers:                        uint32(props.limits.maxDescriptorSetSamplers),
			MaxDescriptorSetUniformBuffers:                  uint32(props.limits.maxDescriptorSetUniformBuffers),
			MaxDescriptorSetUniformBuffersDynamic:           uint32(props.limits.maxDescriptorSetUniformBuffersDynamic),
			MaxDescriptorSetStorageBuffers:                  uint32(props.limits.maxDescriptorSetStorageBuffers),
			MaxDescriptorSetStorageBuffersDynamic:           uint32(props.limits.maxDescriptorSetStorageBuffersDynamic),
			MaxDescriptorSetSampledImages:                   uint32(props.limits.maxDescriptorSetSampledImages),
			MaxDescriptorSetStorageImages:                   uint32(props.limits.maxDescriptorSetStorageImages),
			MaxDescriptorSetInputAttachments:                uint32(props.limits.maxDescriptorSetInputAttachments),
			MaxVertexInputAttributes:                        uint32(props.limits.maxVertexInputAttributes),
			MaxVertexInputBindings:                          uint32(props.limits.maxVertexInputBindings),
			MaxVertexInputAttributeOffset:                   uint32(props.limits.maxVertexInputAttributeOffset),
			MaxVertexInputBindingStride:                     uint32(props.limits.maxVertexInputBindingStride),
			MaxVertexOutputComponents:                       uint32(props.limits.maxVertexOutputComponents),
			MaxTessellationGenerationLevel:                  uint32(props.limits.maxTessellationGenerationLevel),
			MaxTessellationPatchSize:                        uint32(props.limits.maxTessellationPatchSize),
			MaxTessellationControlPerVertexInputComponents:  uint32(props.limits.maxTessellationControlPerVertexInputComponents),
			MaxTessellationControlPerVertexOutputComponents: uint32(props.limits.maxTessellationControlPerVertexOutputComponents),
			MaxTessellationControlPerPatchOutputComponents:  uint32(props.limits.maxTessellationControlPerPatchOutputComponents),
			MaxTessellationControlTotalOutputComponents:     uint32(props.limits.maxTessellationControlTotalOutputComponents),
			MaxTessellationEvaluationInputComponents:        uint32(props.limits.maxTessellationEvaluationInputComponents),
			MaxTessellationEvaluationOutputComponents:       uint32(props.limits.maxTessellationEvaluationOutputComponents),
			MaxGeometryShaderInvocations:                    uint32(props.limits.maxGeometryShaderInvocations),
			MaxGeometryInputComponents:                      uint32(props.limits.maxGeometryInputComponents),
			MaxGeometryOutputComponents:                     uint32(props.limits.maxGeometryOutputComponents),
			MaxGeometryOutputVertices:                       uint32(props.limits.maxGeometryOutputVertices),
			MaxGeometryTotalOutputComponents:                uint32(props.limits.maxGeometryTotalOutputComponents),
			MaxFragmentInputComponents:                      uint32(props.limits.maxFragmentInputComponents),
			MaxFragmentOutputAttachments:                    uint32(props.limits.maxFragmentOutputAttachments),
			MaxFragmentDualSrcAttachments:                   uint32(props.limits.maxFragmentDualSrcAttachments),
			MaxFragmentCombinedOutputResources:              uint32(props.limits.maxFragmentCombinedOutputResources),
			MaxComputeSharedMemorySize:                      uint32(props.limits.maxComputeSharedMemorySize),
			MaxComputeWorkGroupCount: [3]uint32{
				uint32(props.limits.maxComputeWorkGroupCount[0]),
				uint32(props.limits.maxComputeWorkGroupCount[1]),
				uint32(props.limits.maxComputeWorkGroupCount[2]),
			},
			MaxComputeWorkGroupInvocations: uint32(props.limits.maxComputeWorkGroupInvocations),
			MaxComputeWorkGroupSize: [3]uint32{
				uint32(props.limits.maxComputeWorkGroupSize[0]),
				uint32(props.limits.maxComputeWorkGroupSize[1]),
				uint32(props.limits.maxComputeWorkGroupSize[2]),
			},
			SubPixelPrecisionBits:    uint32(props.limits.subPixelPrecisionBits),
			SubTexelPrecisionBits:    uint32(props.limits.subTexelPrecisionBits),
			MipmapPrecisionBits:      uint32(props.limits.mipmapPrecisionBits),
			MaxDrawIndexedIndexValue: uint32(props.limits.maxDrawIndexedIndexValue),
			MaxDrawIndirectCount:     uint32(props.limits.maxDrawIndirectCount),
			MaxSamplerLodBias:        float32(props.limits.maxSamplerLodBias),
			MaxSamplerAnisotropy:     float32(props.limits.maxSamplerAnisotropy),
			MaxViewports:             uint32(props.limits.maxViewports),
			MaxViewportDimensions: [2]uint32{
				uint32(props.limits.maxViewportDimensions[0]),
				uint32(props.limits.maxViewportDimensions[1]),
			},
			ViewportBoundsRange: [2]float32{
				float32(props.limits.viewportBoundsRange[0]),
				float32(props.limits.viewportBoundsRange[1]),
			},
			ViewportSubPixelBits:                 uint32(props.limits.viewportSubPixelBits),
			MinMemoryMapAlignment:                uintptr(props.limits.minMemoryMapAlignment),
			MinTexelBufferOffsetAlignment:        DeviceSize(props.limits.minTexelBufferOffsetAlignment),
			MinUniformBufferOffsetAlignment:      DeviceSize(props.limits.minUniformBufferOffsetAlignment),
			MinStorageBufferOffsetAlignment:      DeviceSize(props.limits.minStorageBufferOffsetAlignment),
			MinTexelOffset:                       int32(props.limits.minTexelOffset),
			MaxTexelOffset:                       uint32(props.limits.maxTexelOffset),
			MinTexelGatherOffset:                 int32(props.limits.minTexelGatherOffset),
			MaxTexelGatherOffset:                 uint32(props.limits.maxTexelGatherOffset),
			MinInterpolationOffset:               float32(props.limits.minInterpolationOffset),
			MaxInterpolationOffset:               float32(props.limits.maxInterpolationOffset),
			SubPixelInterpolationOffsetBits:      uint32(props.limits.subPixelInterpolationOffsetBits),
			MaxFramebufferWidth:                  uint32(props.limits.maxFramebufferWidth),
			MaxFramebufferHeight:                 uint32(props.limits.maxFramebufferHeight),
			MaxFramebufferLayers:                 uint32(props.limits.maxFramebufferLayers),
			FramebufferColorSampleCounts:         SampleCountFlags(props.limits.framebufferColorSampleCounts),
			FramebufferDepthSampleCounts:         SampleCountFlags(props.limits.framebufferDepthSampleCounts),
			FramebufferStencilSampleCounts:       SampleCountFlags(props.limits.framebufferStencilSampleCounts),
			FramebufferNoAttachmentsSampleCounts: SampleCountFlags(props.limits.framebufferNoAttachmentsSampleCounts),
			MaxColorAttachments:                  uint32(props.limits.maxColorAttachments),
			SampledImageColorSampleCounts:        SampleCountFlags(props.limits.sampledImageColorSampleCounts),
			SampledImageIntegerSampleCounts:      SampleCountFlags(props.limits.sampledImageIntegerSampleCounts),
			SampledImageDepthSampleCounts:        SampleCountFlags(props.limits.sampledImageDepthSampleCounts),
			SampledImageStencilSampleCounts:      SampleCountFlags(props.limits.sampledImageStencilSampleCounts),
			StorageImageSampleCounts:             SampleCountFlags(props.limits.storageImageSampleCounts),
			MaxSampleMaskWords:                   uint32(props.limits.maxSampleMaskWords),
			TimestampComputeAndGraphics:          props.limits.timestampComputeAndGraphics == C.VK_TRUE,
			TimestampPeriod:                      float32(props.limits.timestampPeriod),
			MaxClipDistances:                     uint32(props.limits.maxClipDistances),
			MaxCullDistances:                     uint32(props.limits.maxCullDistances),
			MaxCombinedClipAndCullDistances:      uint32(props.limits.maxCombinedClipAndCullDistances),
			DiscreteQueuePriorities:              uint32(props.limits.discreteQueuePriorities),
			PointSizeRange: [2]float32{
				float32(props.limits.pointSizeRange[0]),
				float32(props.limits.pointSizeRange[1]),
			},
			LineWidthRange: [2]float32{
				float32(props.limits.lineWidthRange[0]),
				float32(props.limits.lineWidthRange[1]),
			},
			PointSizeGranularity:               float32(props.limits.pointSizeGranularity),
			LineWidthGranularity:               float32(props.limits.lineWidthGranularity),
			StrictLines:                        props.limits.strictLines == C.VK_TRUE,
			StandardSampleLocations:            props.limits.standardSampleLocations == C.VK_TRUE,
			OptimalBufferCopyOffsetAlignment:   DeviceSize(props.limits.optimalBufferCopyOffsetAlignment),
			OptimalBufferCopyRowPitchAlignment: DeviceSize(props.limits.optimalBufferCopyRowPitchAlignment),
			NonCoherentAtomSize:                DeviceSize(props.limits.nonCoherentAtomSize),
		},
		SparseProperties: PhysicalDeviceSparseProperties{
			ResidencyStandard2DBlockShape:            props.sparseProperties.residencyStandard2DBlockShape == C.VK_TRUE,
			ResidencyStandard2DMultisampleBlockShape: props.sparseProperties.residencyStandard2DMultisampleBlockShape == C.VK_TRUE,
			ResidencyStandard3DBlockShape:            props.sparseProperties.residencyStandard3DBlockShape == C.VK_TRUE,
			ResidencyAlignedMipSize:                  props.sparseProperties.residencyAlignedMipSize == C.VK_TRUE,
			ResidencyNonResidentStrict:               props.sparseProperties.residencyNonResidentStrict == C.VK_TRUE,
		},
	}
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
}

type Extent2D struct {
	Width  uint32
	Height uint32
}

type Extent3D struct {
	Width  uint32
	Height uint32
	Depth  uint32
}

func (dev *PhysicalDevice) QueueFamilyProperties() []*QueueFamilyProperties {
	var count C.uint32_t
	C.domVkGetPhysicalDeviceQueueFamilyProperties(dev.instance.fps[vkGetPhysicalDeviceQueueFamilyProperties], dev.hnd, &count, nil)
	props := (*C.VkQueueFamilyProperties)(C.calloc(C.size_t(count), C.size_t(unsafe.Sizeof(C.VkQueueFamilyProperties{}))))
	C.domVkGetPhysicalDeviceQueueFamilyProperties(dev.instance.fps[vkGetPhysicalDeviceQueueFamilyProperties], dev.hnd, &count, props)
	var out []*QueueFamilyProperties
	for _, prop := range (*[1 << 31]C.VkQueueFamilyProperties)(unsafe.Pointer(props))[:count] {
		out = append(out, &QueueFamilyProperties{
			QueueFlags:         QueueFlags(prop.queueFlags),
			QueueCount:         uint32(prop.queueCount),
			TimestampValidBits: uint32(prop.timestampValidBits),
			MinImageTransferGranularity: Extent3D{
				Width:  uint32(prop.minImageTransferGranularity.width),
				Height: uint32(prop.minImageTransferGranularity.height),
				Depth:  uint32(prop.minImageTransferGranularity.depth),
			},
		})
	}
	return out
}

type DeviceQueueCreateInfo struct {
	Next             unsafe.Pointer
	Flags            DeviceQueueCreateFlags
	QueueFamilyIndex uint32
	QueuePriorities  []float32
}

type DeviceCreateInfo struct {
	Next                  unsafe.Pointer
	QueueCreateInfos      []*DeviceQueueCreateInfo
	EnabledExtensionNames []string
	EnabledFeatures       *PhysicalDeviceFeatures
}

func (dev *PhysicalDevice) CreateDevice(info *DeviceCreateInfo) (*Device, Result) {
	// TODO(dh): support custom allocators
	var free1 func()
	ptr := (*C.VkDeviceCreateInfo)(C.calloc(1, C.size_t(unsafe.Sizeof(C.VkDeviceCreateInfo{}))))
	ptr.sType = C.VkStructureType(StructureTypeDeviceCreateInfo)
	ptr.pNext = info.Next
	ptr.queueCreateInfoCount = C.uint32_t(len(info.QueueCreateInfos))
	ptr.pQueueCreateInfos = (*C.VkDeviceQueueCreateInfo)(C.calloc(C.size_t(len(info.QueueCreateInfos)), C.size_t(unsafe.Sizeof(C.VkDeviceQueueCreateInfo{}))))
	defer C.free(unsafe.Pointer(ptr.pQueueCreateInfos))
	arr := (*[1 << 31]C.VkDeviceQueueCreateInfo)(unsafe.Pointer(ptr.pQueueCreateInfos))[:len(info.QueueCreateInfos)]
	for i, obj := range info.QueueCreateInfos {
		arr[i] = C.VkDeviceQueueCreateInfo{
			sType:            C.VkStructureType(StructureTypeDeviceQueueCreateInfo),
			pNext:            obj.Next,
			flags:            C.VkDeviceQueueCreateFlags(obj.Flags),
			queueFamilyIndex: C.uint32_t(obj.QueueFamilyIndex),
			queueCount:       C.uint32_t(len(obj.QueuePriorities)),
			pQueuePriorities: (*C.float)(&obj.QueuePriorities[0]),
		}
	}
	ptr.enabledExtensionCount = C.uint32_t(len(info.EnabledExtensionNames))
	ptr.ppEnabledExtensionNames, free1 = externStrings(info.EnabledExtensionNames)
	defer free1()
	if info.EnabledFeatures != nil {
		ptr.pEnabledFeatures = (*C.VkPhysicalDeviceFeatures)(C.calloc(1, C.size_t(unsafe.Sizeof(C.VkPhysicalDeviceFeatures{}))))
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
		defer C.free(unsafe.Pointer(ptr.pEnabledFeatures))
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

	return ldev, Success
}

type Device struct {
	hnd C.VkDevice

	fps                 [deviceMaxPFN]C.PFN_vkVoidFunction
	vkGetDeviceProcAddr C.PFN_vkGetDeviceProcAddr
}

func (dev *Device) init() {
	for i, name := range deviceFpNames {
		dev.fps[i] = dev.mustGetDeviceProcAddr(name)
	}
}

func (dev *Device) String() string {
	return fmt.Sprintf("VkDevice(%p)", dev.hnd)
}

func (dev *Device) mustGetDeviceProcAddr(name string) C.PFN_vkVoidFunction {
	fp := dev.getDeviceProcAddr(name)
	if fp == nil {
		panic(fmt.Sprintf("couldn't load function %s", name))
	}
	return fp
}

func (dev *Device) getDeviceProcAddr(name string) C.PFN_vkVoidFunction {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return C.domVkGetDeviceProcAddr(dev.vkGetDeviceProcAddr, dev.hnd, cName)
}

type Queue struct {
	hnd C.VkQueue
}

func (q *Queue) String() string {
	return fmt.Sprintf("VkQueue(%p)", q.hnd)
}

func (dev *Device) Queue(family, index uint32) *Queue {
	var out C.VkQueue
	C.domVkGetDeviceQueue(dev.fps[vkGetDeviceQueue], dev.hnd, C.uint32_t(family), C.uint32_t(index), &out)
	return &Queue{hnd: out}
}

type CommandPool struct {
	hnd C.VkCommandPool
	dev *Device

	freePtrs []C.VkCommandBuffer
}

func (pool *CommandPool) String() string {
	return fmt.Sprintf("VkCommandPool(%p)", pool.hnd)
}

type CommandBuffer struct {
	hnd  C.VkCommandBuffer
	pool *CommandPool
}

func (buf *CommandBuffer) String() string {
	return fmt.Sprintf("VkCommandBuffer(%p)", buf)
}

func (buf *CommandBuffer) Free() {
	C.domVkFreeCommandBuffers(buf.pool.dev.fps[vkFreeCommandBuffers], buf.pool.dev.hnd, buf.pool.hnd, 1, &buf.hnd)
}

func (buf *CommandBuffer) Reset(flags CommandBufferResetFlags) error {
	res := Result(C.domVkResetCommandBuffer(buf.pool.dev.fps[vkResetCommandBuffer], buf.hnd, C.VkCommandBufferResetFlags(flags)))
	if res != Success {
		return res
	}
	return nil
}

type CommandBufferBeginInfo struct {
	Next            unsafe.Pointer
	Flags           CommandBufferUsageFlags
	InheritanceInfo *CommandBufferInheritanceInfo
}

type RenderPass C.VkRenderPass
type Framebuffer C.VkFramebuffer

type CommandBufferInheritanceInfo struct {
	Next                 unsafe.Pointer
	RenderPass           RenderPass
	Subpass              uint32
	Framebuffer          Framebuffer
	OcclusionQueryEnable bool
	QueryFlags           QueryControlFlags
	PipelineStatistics   QueryPipelineStatisticFlags
}

func (buf *CommandBuffer) Begin(info *CommandBufferBeginInfo) error {
	ptr := (*C.VkCommandBufferBeginInfo)(C.calloc(1, C.size_t(unsafe.Sizeof(C.VkCommandBufferBeginInfo{}))))
	defer C.free(unsafe.Pointer(ptr))
	ptr.sType = C.VkStructureType(StructureTypeCommandBufferBeginInfo)
	ptr.pNext = info.Next
	ptr.flags = C.VkCommandBufferUsageFlags(info.Flags)
	if info.InheritanceInfo != nil {
		ptr.pInheritanceInfo = (*C.VkCommandBufferInheritanceInfo)(C.calloc(1, C.size_t(unsafe.Sizeof(C.VkCommandBufferInheritanceInfo{}))))
		defer C.free(unsafe.Pointer(ptr.pInheritanceInfo))
		ptr.pInheritanceInfo.sType = C.VkStructureType(StructureTypeCommandBufferInheritanceInfo)
		ptr.pInheritanceInfo.pNext = info.InheritanceInfo.Next
		ptr.pInheritanceInfo.renderPass = C.VkRenderPass(info.InheritanceInfo.RenderPass)
		ptr.pInheritanceInfo.subpass = C.uint32_t(info.InheritanceInfo.Subpass)
		ptr.pInheritanceInfo.framebuffer = C.VkFramebuffer(info.InheritanceInfo.Framebuffer)
		ptr.pInheritanceInfo.occlusionQueryEnable = vkBool(info.InheritanceInfo.OcclusionQueryEnable)
		ptr.pInheritanceInfo.queryFlags = C.VkQueryControlFlags(info.InheritanceInfo.QueryFlags)
		ptr.pInheritanceInfo.pipelineStatistics = C.VkQueryPipelineStatisticFlags(info.InheritanceInfo.PipelineStatistics)
	}
	res := Result(C.domVkBeginCommandBuffer(buf.pool.dev.fps[vkBeginCommandBuffer], buf.hnd, ptr))
	if res != Success {
		return res
	}
	return nil
}

func (buf *CommandBuffer) End() error {
	res := Result(C.domVkEndCommandBuffer(buf.pool.dev.fps[vkEndCommandBuffer], buf.hnd))
	if res != Success {
		return res
	}
	return nil
}

type CommandPoolCreateInfo struct {
	Next             unsafe.Pointer
	Flags            CommandPoolCreateFlags
	QueueFamilyIndex uint32
}

func (dev *Device) CreateCommandPool(info *CommandPoolCreateInfo) (*CommandPool, error) {
	// TODO(dh): support callbacks
	ptr := (*C.VkCommandPoolCreateInfo)(C.calloc(1, C.size_t(unsafe.Sizeof(C.VkCommandPoolCreateInfo{}))))
	defer C.free(unsafe.Pointer(ptr))
	ptr.sType = C.VkStructureType(StructureTypeCommandPoolCreateInfo)
	ptr.pNext = info.Next
	ptr.flags = C.VkCommandPoolCreateFlags(info.Flags)
	ptr.queueFamilyIndex = C.uint32_t(info.QueueFamilyIndex)
	var pool C.VkCommandPool
	res := Result(C.domVkCreateCommandPool(dev.fps[vkCreateCommandPool], dev.hnd, ptr, nil, &pool))
	if res != Success {
		return nil, res
	}
	return &CommandPool{hnd: pool, dev: dev}, nil
}

func (pool *CommandPool) Trim(flags CommandPoolTrimFlags) {
	C.domVkTrimCommandPool(pool.dev.fps[vkTrimCommandPool], pool.dev.hnd, pool.hnd, C.VkCommandPoolTrimFlags(flags))
}

func vkBool(b bool) C.VkBool32 {
	if b {
		return C.VK_TRUE
	}
	return C.VK_FALSE
}

func (pool *CommandPool) Reset(flags CommandPoolResetFlags) error {
	res := Result(C.domVkResetCommandPool(pool.dev.fps[vkResetCommandPool], pool.dev.hnd, pool.hnd, C.VkCommandPoolResetFlags(flags)))
	if res != Success {
		return res
	}
	return nil
}

type CommandBufferAllocateInfo struct {
	Next               unsafe.Pointer
	Level              CommandBufferLevel
	CommandBufferCount uint32
}

func (pool *CommandPool) AllocateCommandBuffers(info *CommandBufferAllocateInfo) ([]*CommandBuffer, error) {
	ptr := (*C.VkCommandBufferAllocateInfo)(C.calloc(1, C.size_t(unsafe.Sizeof(C.VkCommandBufferAllocateInfo{}))))
	defer C.free(unsafe.Pointer(ptr))
	ptr.sType = C.VkStructureType(StructureTypeCommandBufferAllocateInfo)
	ptr.pNext = info.Next
	ptr.commandPool = pool.hnd
	ptr.level = C.VkCommandBufferLevel(info.Level)
	ptr.commandBufferCount = C.uint32_t(info.CommandBufferCount)
	bufs := make([]C.VkCommandBuffer, info.CommandBufferCount)
	res := Result(C.domVkAllocateCommandBuffers(pool.dev.fps[vkAllocateCommandBuffers], pool.dev.hnd, ptr, &bufs[0]))
	if res != Success {
		return nil, res
	}
	out := make([]*CommandBuffer, info.CommandBufferCount)
	for i, buf := range bufs {
		out[i] = &CommandBuffer{hnd: buf}
	}
	return out, nil
}

func (pool *CommandPool) FreeBuffers(bufs []*CommandBuffer) {
	ptrs := pool.freePtrs[:0]
	if cap(ptrs) >= len(bufs) {
		ptrs = ptrs[:len(bufs)]
	} else {
		ptrs = make([]C.VkCommandBuffer, len(bufs))
	}
	for i, buf := range bufs {
		ptrs[i] = buf.hnd
	}
	C.domVkFreeCommandBuffers(pool.dev.fps[vkFreeCommandBuffers], pool.dev.hnd, pool.hnd, C.uint32_t(len(bufs)), &ptrs[0])
	pool.freePtrs = ptrs[:0]
}

func vkGetInstanceProcAddr(instance C.VkInstance, name string) C.PFN_vkVoidFunction {
	// TODO(dh): return a mock function pointer that panics with a nice message

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	fp := C.vkGetInstanceProcAddr(instance, cName)
	if debug && fp == nil {
		fmt.Fprintln(os.Stderr, "no function pointer for", name)
	}
	return fp
}

func mustVkGetInstanceProcAddr(instance C.VkInstance, name string) C.PFN_vkVoidFunction {
	fp := vkGetInstanceProcAddr(instance, name)
	if fp == nil {
		panic(fmt.Sprintf("couldn't load function %s", name))
	}
	return fp
}

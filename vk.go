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
// void     domVkCmdSetLineWidth(PFN_vkCmdSetLineWidth fp, VkCommandBuffer commandBuffer, float lineWidth);
// void     domVkCmdSetDepthBias(PFN_vkCmdSetDepthBias fp, VkCommandBuffer commandBuffer, float depthBiasConstantFactor, float depthBiasClamp, float depthBiasSlopeFactor);
// void     domVkCmdSetBlendConstants(PFN_vkCmdSetBlendConstants fp, VkCommandBuffer commandBuffer, const float blendConstants[4]);
// void     domVkCmdDraw(PFN_vkCmdDraw fp, VkCommandBuffer commandBuffer, uint32_t vertexCount, uint32_t instanceCount, uint32_t firstVertex, uint32_t firstInstance);
// VkResult domVkQueueWaitIdle(PFN_vkQueueWaitIdle fp, VkQueue queue);
// VkResult domVkDeviceWaitIdle(PFN_vkDeviceWaitIdle fp, VkDevice device);
// VkResult domVkCreateImageView(PFN_vkCreateImageView fp, VkDevice device, const VkImageViewCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkImageView* pView);
// VkResult domVkCreateShaderModule(PFN_vkCreateShaderModule fp, VkDevice device, const VkShaderModuleCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkShaderModule* pShaderModule);
// VkResult domVkCreateGraphicsPipelines(PFN_vkCreateGraphicsPipelines fp, VkDevice device, VkPipelineCache pipelineCache, uint32_t createInfoCount, const VkGraphicsPipelineCreateInfo* pCreateInfos, const VkAllocationCallbacks* pAllocator, VkPipeline* pPipelines);
// VkResult domVkCreatePipelineLayout(PFN_vkCreatePipelineLayout fp, VkDevice device, const VkPipelineLayoutCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkPipelineLayout* pPipelineLayout);
// VkResult domVkCreateRenderPass(PFN_vkCreateRenderPass fp, VkDevice device, const VkRenderPassCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkRenderPass* pRenderPass);
// VkResult domVkCreateFramebuffer(PFN_vkCreateFramebuffer fp, VkDevice device, const VkFramebufferCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkFramebuffer* pFramebuffer);
// void     domVkCmdBeginRenderPass(PFN_vkCmdBeginRenderPass fp, VkCommandBuffer commandBuffer, const VkRenderPassBeginInfo* pRenderPassBegin, VkSubpassContents contents);
// void     domVkCmdBindPipeline(PFN_vkCmdBindPipeline fp, VkCommandBuffer commandBuffer, VkPipelineBindPoint pipelineBindPoint, VkPipeline pipeline);
// void     domVkCmdEndRenderPass(PFN_vkCmdEndRenderPass fp, VkCommandBuffer commandBuffer);
// VkResult domVkCreateSemaphore(PFN_vkCreateSemaphore fp, VkDevice device, const VkSemaphoreCreateInfo* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkSemaphore* pSemaphore);
// VkResult domVkQueueSubmit(PFN_vkQueueSubmit fp, VkQueue queue, uint32_t submitCount, const VkSubmitInfo* pSubmits, VkFence fence);
import "C"
import (
	"fmt"
	"math"
	"os"
	"unsafe"
)

// XXX add assertions that some of our structs are the same size as the C structs.
// any struct used in ucopy must have an assertion.

const debug = true

type (
	DeviceSize = uint64
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
	slice := (*[math.MaxInt32]*C.char)(ptr)[:len(ss)]
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

func externFloat32(vs []float32) *C.float {
	if len(vs) == 0 {
		return nil
	}
	return (*C.float)(C.CBytes((*[math.MaxInt32]byte)(unsafe.Pointer(&vs[0]))[:uintptr(len(vs))*unsafe.Sizeof(float32(0))]))
}

func CreateInstance(info *InstanceCreateInfo) (*Instance, error) {
	// TODO(dh): support a custom allocator
	var free1, free2 func()

	ptr := (*C.VkInstanceCreateInfo)(C.calloc(1, C.sizeof_VkInstanceCreateInfo))
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
		ptr.pApplicationInfo = (*C.VkApplicationInfo)(C.calloc(1, C.sizeof_VkApplicationInfo))
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
	// VK_DEFINE_HANDLE(VkInstance)
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
		devs = (*C.VkPhysicalDevice)(C.calloc(C.size_t(count), C.sizeof_VkPhysicalDevice))
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
	for _, dev := range (*[math.MaxInt32]C.VkPhysicalDevice)(unsafe.Pointer(devs))[:count] {
		out = append(out, &PhysicalDevice{dev, ins})
	}
	return out, nil
}

type PhysicalDevice struct {
	// VK_DEFINE_HANDLE(VkPhysicalDevice)
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
	props := (*C.VkQueueFamilyProperties)(C.calloc(C.size_t(count), C.sizeof_VkQueueFamilyProperties))
	C.domVkGetPhysicalDeviceQueueFamilyProperties(dev.instance.fps[vkGetPhysicalDeviceQueueFamilyProperties], dev.hnd, &count, props)
	var out []*QueueFamilyProperties
	for _, prop := range (*[math.MaxInt32]C.VkQueueFamilyProperties)(unsafe.Pointer(props))[:count] {
		// XXX can we use ucopy here?
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

func (dev *PhysicalDevice) CreateDevice(info *DeviceCreateInfo) (*Device, error) {
	// TODO(dh): support custom allocators
	var free1 func()
	ptr := (*C.VkDeviceCreateInfo)(C.calloc(1, C.sizeof_VkDeviceCreateInfo))
	ptr.sType = C.VkStructureType(StructureTypeDeviceCreateInfo)
	ptr.pNext = info.Next
	ptr.queueCreateInfoCount = C.uint32_t(len(info.QueueCreateInfos))
	ptr.pQueueCreateInfos = (*C.VkDeviceQueueCreateInfo)(C.calloc(C.size_t(len(info.QueueCreateInfos)), C.sizeof_VkDeviceQueueCreateInfo))
	defer C.free(unsafe.Pointer(ptr.pQueueCreateInfos))
	arr := (*[math.MaxInt32]C.VkDeviceQueueCreateInfo)(unsafe.Pointer(ptr.pQueueCreateInfos))[:len(info.QueueCreateInfos)]
	for i, obj := range info.QueueCreateInfos {
		arr[i] = C.VkDeviceQueueCreateInfo{
			sType:            C.VkStructureType(StructureTypeDeviceQueueCreateInfo),
			pNext:            obj.Next,
			flags:            C.VkDeviceQueueCreateFlags(obj.Flags),
			queueFamilyIndex: C.uint32_t(obj.QueueFamilyIndex),
			queueCount:       C.uint32_t(len(obj.QueuePriorities)),
			pQueuePriorities: externFloat32(obj.QueuePriorities),
		}
		defer C.free(unsafe.Pointer(arr[i].pQueuePriorities))
	}
	ptr.enabledExtensionCount = C.uint32_t(len(info.EnabledExtensionNames))
	ptr.ppEnabledExtensionNames, free1 = externStrings(info.EnabledExtensionNames)
	defer free1()
	if info.EnabledFeatures != nil {
		ptr.pEnabledFeatures = (*C.VkPhysicalDeviceFeatures)(C.calloc(1, C.sizeof_VkPhysicalDeviceFeatures))
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

	return ldev, nil
}

type Device struct {
	// VK_DEFINE_HANDLE(VkDevice)
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
	// VK_DEFINE_HANDLE(VkQueue)
	hnd C.VkQueue
	fps *[deviceMaxPFN]C.PFN_vkVoidFunction
}

func (q *Queue) String() string {
	return fmt.Sprintf("VkQueue(%p)", q.hnd)
}

func (q *Queue) WaitIdle() error {
	res := Result(C.domVkQueueWaitIdle(q.fps[vkQueueWaitIdle], q.hnd))
	if res != Success {
		return res
	}
	return nil
}

func (dev *Device) Queue(family, index uint32) *Queue {
	var out C.VkQueue
	C.domVkGetDeviceQueue(dev.fps[vkGetDeviceQueue], dev.hnd, C.uint32_t(family), C.uint32_t(index), &out)
	return &Queue{hnd: out, fps: &dev.fps}
}

type CommandPool struct {
	// VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkCommandPool)
	hnd C.VkCommandPool
	dev *Device

	freePtrs []C.VkCommandBuffer
}

func (pool *CommandPool) String() string {
	return fmt.Sprintf("VkCommandPool(%p)", pool.hnd)
}

type CommandBuffer struct {
	// VK_DEFINE_HANDLE(VkCommandBuffer)
	hnd C.VkCommandBuffer
	fps *[deviceMaxPFN]C.PFN_vkVoidFunction
}

func (buf *CommandBuffer) String() string {
	return fmt.Sprintf("VkCommandBuffer(%p)", buf)
}

func (buf *CommandBuffer) Reset(flags CommandBufferResetFlags) error {
	res := Result(C.domVkResetCommandBuffer(buf.fps[vkResetCommandBuffer], buf.hnd, C.VkCommandBufferResetFlags(flags)))
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

type RenderPass struct {
	// VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkRenderPass)
	hnd C.VkRenderPass
}

type Framebuffer struct {
	// VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkFramebuffer)
	hnd C.VkFramebuffer
}

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
	ptr := (*C.VkCommandBufferBeginInfo)(C.calloc(1, C.sizeof_VkCommandBufferBeginInfo))
	defer C.free(unsafe.Pointer(ptr))
	ptr.sType = C.VkStructureType(StructureTypeCommandBufferBeginInfo)
	ptr.pNext = info.Next
	ptr.flags = C.VkCommandBufferUsageFlags(info.Flags)
	if info.InheritanceInfo != nil {
		ptr.pInheritanceInfo = (*C.VkCommandBufferInheritanceInfo)(C.calloc(1, C.sizeof_VkCommandBufferInheritanceInfo))
		defer C.free(unsafe.Pointer(ptr.pInheritanceInfo))
		ptr.pInheritanceInfo.sType = C.VkStructureType(StructureTypeCommandBufferInheritanceInfo)
		ptr.pInheritanceInfo.pNext = info.InheritanceInfo.Next
		ptr.pInheritanceInfo.renderPass = C.VkRenderPass(info.InheritanceInfo.RenderPass.hnd)
		ptr.pInheritanceInfo.subpass = C.uint32_t(info.InheritanceInfo.Subpass)
		ptr.pInheritanceInfo.framebuffer = C.VkFramebuffer(info.InheritanceInfo.Framebuffer.hnd)
		ptr.pInheritanceInfo.occlusionQueryEnable = vkBool(info.InheritanceInfo.OcclusionQueryEnable)
		ptr.pInheritanceInfo.queryFlags = C.VkQueryControlFlags(info.InheritanceInfo.QueryFlags)
		ptr.pInheritanceInfo.pipelineStatistics = C.VkQueryPipelineStatisticFlags(info.InheritanceInfo.PipelineStatistics)
	}
	res := Result(C.domVkBeginCommandBuffer(buf.fps[vkBeginCommandBuffer], buf.hnd, ptr))
	if res != Success {
		return res
	}
	return nil
}

func (buf *CommandBuffer) End() error {
	res := Result(C.domVkEndCommandBuffer(buf.fps[vkEndCommandBuffer], buf.hnd))
	if res != Success {
		return res
	}
	return nil
}

func (buf *CommandBuffer) SetLineWidth(lineWidth float32) {
	C.domVkCmdSetLineWidth(buf.fps[vkCmdSetLineWidth], buf.hnd, C.float(lineWidth))
}

func (buf *CommandBuffer) SetDepthBias(constantFactor, clamp, slopeFactor float32) {
	C.domVkCmdSetDepthBias(buf.fps[vkCmdSetDepthBias], buf.hnd, C.float(constantFactor), C.float(clamp), C.float(slopeFactor))
}

func (buf *CommandBuffer) SetBlendConstants(blendConstants [4]float32) {
	C.domVkCmdSetBlendConstants(buf.fps[vkCmdSetBlendConstants], buf.hnd, (*C.float)(unsafe.Pointer(&blendConstants[0])))
}

func (buf *CommandBuffer) Draw(vertexCount, instanceCount, firstVertex, firstInstance uint32) {
	C.domVkCmdDraw(buf.fps[vkCmdDraw], buf.hnd, C.uint32_t(vertexCount), C.uint32_t(instanceCount), C.uint32_t(firstVertex), C.uint32_t(firstInstance))
}

func (info RenderPassBeginInfo) c() *C.VkRenderPassBeginInfo {
	size0 := uintptr(C.sizeof_VkRenderPassBeginInfo)
	size1 := C.sizeof_VkClearValue * uintptr(len(info.ClearValues))
	size := size0 + size1
	alloc := C.calloc(1, C.size_t(size))
	cinfo := (*C.VkRenderPassBeginInfo)(alloc)
	*cinfo = C.VkRenderPassBeginInfo{
		sType:           C.VkStructureType(StructureTypeRenderPassBeginInfo),
		pNext:           info.Next,
		renderPass:      info.RenderPass.hnd,
		framebuffer:     info.Framebuffer.hnd,
		clearValueCount: C.uint32_t(len(info.ClearValues)),
		pClearValues:    (*C.VkClearValue)(unsafe.Pointer(uintptr(alloc) + size0)),
	}
	ucopy1(unsafe.Pointer(&cinfo.renderArea), unsafe.Pointer(&info.RenderArea), C.sizeof_VkRect2D)
	arr := (*[math.MaxInt32]C.VkClearValue)(unsafe.Pointer(cinfo.pClearValues))[:len(info.ClearValues)]
	for i := range arr {
		switch v := info.ClearValues[i].(type) {
		case ClearColorValueFloat32s:
			copy(arr[i][:], (*[16]byte)(unsafe.Pointer(&v))[:])
		case ClearColorValueInt32s:
			copy(arr[i][:], (*[16]byte)(unsafe.Pointer(&v))[:])
		case ClearColorValueUint32s:
			copy(arr[i][:], (*[16]byte)(unsafe.Pointer(&v))[:])
		case ClearDepthStencilValue:
			ucopy1(unsafe.Pointer(&arr[i]), unsafe.Pointer(&v), C.sizeof_VkClearDepthStencilValue)
		default:
			panic(fmt.Sprintf("unreachable: %T", v))
		}
	}
	return cinfo
}

func (buf *CommandBuffer) BeginRenderPass(info *RenderPassBeginInfo, contents SubpassContents) {
	cinfo := info.c()
	defer C.free(unsafe.Pointer(cinfo))
	C.domVkCmdBeginRenderPass(buf.fps[vkCmdBeginRenderPass], buf.hnd, cinfo, C.VkSubpassContents(contents))
}

func (buf *CommandBuffer) EndRenderPass() {
	C.domVkCmdEndRenderPass(buf.fps[vkCmdEndRenderPass], buf.hnd)
}

func (buf *CommandBuffer) BindPipeline(pipelineBindPoint PipelineBindPoint, pipeline Pipeline) {
	C.domVkCmdBindPipeline(buf.fps[vkCmdBindPipeline], buf.hnd, C.VkPipelineBindPoint(pipelineBindPoint), pipeline.hnd)
}

type CommandPoolCreateInfo struct {
	Next             unsafe.Pointer
	Flags            CommandPoolCreateFlags
	QueueFamilyIndex uint32
}

func (dev *Device) CreateCommandPool(info *CommandPoolCreateInfo) (*CommandPool, error) {
	// TODO(dh): support callbacks
	ptr := (*C.VkCommandPoolCreateInfo)(C.calloc(1, C.sizeof_VkCommandPoolCreateInfo))
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
	ptr := (*C.VkCommandBufferAllocateInfo)(C.calloc(1, C.sizeof_VkCommandBufferAllocateInfo))
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
		out[i] = &CommandBuffer{hnd: buf, fps: &pool.dev.fps}
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

func (dev *Device) WaitIdle() error {
	res := Result(C.domVkDeviceWaitIdle(dev.fps[vkDeviceWaitIdle], dev.hnd))
	if res != Success {
		return res
	}
	return nil
}

type Image struct {
	// VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkImage)
	hnd C.VkImage
	dev *Device
}

func (img Image) String() string {
	return fmt.Sprintf("VkImage(%p)", img.hnd)
}

type ImageView struct {
	// VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkImageView)
	hnd C.VkImageView
}

type ImageViewCreateInfo struct {
	Next             unsafe.Pointer
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
}

type ImageSubresourceRange struct {
	AspectMask     ImageAspectFlags
	BaseMipLevel   uint32
	LevelCount     uint32
	BaseArrayLayer uint32
	LayerCount     uint32
}

func (img Image) CreateView(info *ImageViewCreateInfo) (ImageView, error) {
	// TODO(dh): support custom allocator
	ptr := (*C.VkImageViewCreateInfo)(C.calloc(1, C.sizeof_VkImageViewCreateInfo))
	defer C.free(unsafe.Pointer(ptr))
	ptr.sType = C.VkStructureType(StructureTypeImageViewCreateInfo)
	ptr.pNext = info.Next
	ptr.image = img.hnd
	ptr.viewType = C.VkImageViewType(info.ViewType)
	ptr.format = C.VkFormat(info.Format)
	ptr.components = C.VkComponentMapping{
		r: C.VkComponentSwizzle(info.Components.R),
		g: C.VkComponentSwizzle(info.Components.G),
		b: C.VkComponentSwizzle(info.Components.B),
		a: C.VkComponentSwizzle(info.Components.A),
	}
	ptr.subresourceRange = C.VkImageSubresourceRange{
		aspectMask:     C.VkImageAspectFlags(info.SubresourceRange.AspectMask),
		baseMipLevel:   C.uint32_t(info.SubresourceRange.BaseMipLevel),
		levelCount:     C.uint32_t(info.SubresourceRange.LevelCount),
		baseArrayLayer: C.uint32_t(info.SubresourceRange.BaseArrayLayer),
		layerCount:     C.uint32_t(info.SubresourceRange.LayerCount),
	}

	var hnd C.VkImageView
	res := Result(C.domVkCreateImageView(img.dev.fps[vkCreateImageView], img.dev.hnd, ptr, nil, &hnd))
	if res != Success {
		return ImageView{}, res
	}
	return ImageView{hnd: hnd}, nil
}

type ShaderModule struct {
	// VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkShaderModule)
	hnd C.VkShaderModule
}

type ShaderModuleCreateInfo struct {
	Next unsafe.Pointer
	Code []byte
}

func (dev *Device) CreateShaderModule(info *ShaderModuleCreateInfo) (ShaderModule, error) {
	// TODO(dh): support custom allocator
	ptr := (*C.VkShaderModuleCreateInfo)(C.calloc(1, C.sizeof_VkShaderModuleCreateInfo))
	defer C.free(unsafe.Pointer(ptr))
	ptr.sType = C.VkStructureType(StructureTypeShaderModuleCreateInfo)
	ptr.pNext = info.Next
	ptr.codeSize = C.size_t(len(info.Code))
	ptr.pCode = (*C.uint32_t)(C.CBytes(info.Code))
	defer C.free(unsafe.Pointer(ptr.pCode))
	var hnd C.VkShaderModule
	res := Result(C.domVkCreateShaderModule(dev.fps[vkCreateShaderModule], dev.hnd, ptr, nil, &hnd))
	if res != Success {
		return ShaderModule{}, res
	}
	return ShaderModule{hnd: hnd}, nil
}

type PipelineShaderStageCreateInfo struct {
	Next   unsafe.Pointer
	Stage  ShaderStageFlags
	Module ShaderModule
	Name   string
	// TODO(dh): support specialization info
}

type PipelineVertexInputStateCreateInfo struct {
	Next                        unsafe.Pointer
	VertexBindingDescriptions   []VertexInputBindingDescription
	VertexAttributeDescriptions []VertexInputAttributeDescription
}

func (info PipelineVertexInputStateCreateInfo) c() *C.VkPipelineVertexInputStateCreateInfo {
	size0 := uintptr(C.sizeof_VkPipelineVertexInputStateCreateInfo)
	size1 := uintptr(len(info.VertexBindingDescriptions)) * C.sizeof_VkVertexInputBindingDescription
	size2 := uintptr(len(info.VertexAttributeDescriptions)) * C.sizeof_VkVertexInputAttributeDescription
	size := size0 + size1 + size2

	alloc := C.calloc(1, C.size_t(size))
	cinfo := (*C.VkPipelineVertexInputStateCreateInfo)(alloc)
	bindings := (*C.VkVertexInputBindingDescription)(unsafe.Pointer(uintptr(alloc) + size0))
	attribs := (*C.VkVertexInputAttributeDescription)(unsafe.Pointer(uintptr(alloc) + size0 + size1))
	*cinfo = C.VkPipelineVertexInputStateCreateInfo{
		sType: C.VkStructureType(StructureType(StructureTypePipelineVertexInputStateCreateInfo)),
		pNext: info.Next,
		flags: 0,
		vertexBindingDescriptionCount:   C.uint32_t(len(info.VertexBindingDescriptions)),
		pVertexBindingDescriptions:      bindings,
		vertexAttributeDescriptionCount: C.uint32_t(len(info.VertexAttributeDescriptions)),
		pVertexAttributeDescriptions:    attribs,
	}
	ucopy(unsafe.Pointer(bindings), unsafe.Pointer(&info.VertexBindingDescriptions), C.sizeof_VkVertexInputBindingDescription)
	ucopy(unsafe.Pointer(attribs), unsafe.Pointer(&info.VertexAttributeDescriptions), C.sizeof_VkVertexInputAttributeDescription)
	return cinfo
}

type VertexInputBindingDescription struct {
	Binding   uint32
	Stride    uint32
	InputRate VertexInputRate
}

type VertexInputAttributeDescription struct {
	Location uint32
	Binding  uint32
	Format   Format
	Offset   uint32
}

type PipelineInputAssemblyStateCreateInfo struct {
	Next                   unsafe.Pointer
	Topology               PrimitiveTopology
	PrimitiveRestartEnable bool
}

func (info PipelineInputAssemblyStateCreateInfo) c() *C.VkPipelineInputAssemblyStateCreateInfo {
	cinfo := (*C.VkPipelineInputAssemblyStateCreateInfo)(C.calloc(1, C.sizeof_VkPipelineInputAssemblyStateCreateInfo))
	*cinfo = C.VkPipelineInputAssemblyStateCreateInfo{
		sType:                  C.VkStructureType(StructureTypePipelineInputAssemblyStateCreateInfo),
		pNext:                  info.Next,
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
}

type Rect2D struct {
	Offset Offset2D
	Extent Extent2D
}

type Offset2D struct {
	X int32
	Y int32
}

type PipelineViewportStateCreateInfo struct {
	Next      unsafe.Pointer
	Viewports []Viewport
	Scissors  []Rect2D
}

func (info PipelineViewportStateCreateInfo) c() *C.VkPipelineViewportStateCreateInfo {
	size0 := uintptr(C.sizeof_VkPipelineViewportStateCreateInfo)
	size1 := uintptr(len(info.Viewports)) * C.sizeof_VkViewport
	size2 := uintptr(len(info.Scissors)) * C.sizeof_VkRect2D
	size := size0 + size1 + size2
	alloc := C.calloc(1, C.size_t(size))
	cinfo := (*C.VkPipelineViewportStateCreateInfo)(alloc)
	viewports := (*C.VkViewport)(unsafe.Pointer(uintptr(alloc) + size0))
	scissors := (*C.VkRect2D)(unsafe.Pointer(uintptr(alloc) + size0 + size1))
	*cinfo = C.VkPipelineViewportStateCreateInfo{
		sType:         C.VkStructureType(StructureTypePipelineViewportStateCreateInfo),
		pNext:         info.Next,
		flags:         0,
		viewportCount: C.uint32_t(len(info.Viewports)),
		pViewports:    viewports,
		scissorCount:  C.uint32_t(len(info.Scissors)),
		pScissors:     scissors,
	}
	ucopy(unsafe.Pointer(viewports), unsafe.Pointer(&info.Viewports), C.sizeof_VkViewport)
	ucopy(unsafe.Pointer(scissors), unsafe.Pointer(&info.Scissors), C.sizeof_VkRect2D)
	return cinfo
}

type PipelineRasterizationStateCreateInfo struct {
	Next                    unsafe.Pointer
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

func (info PipelineRasterizationStateCreateInfo) c() *C.VkPipelineRasterizationStateCreateInfo {
	cinfo := (*C.VkPipelineRasterizationStateCreateInfo)(C.calloc(1, C.sizeof_VkPipelineRasterizationStateCreateInfo))
	*cinfo = C.VkPipelineRasterizationStateCreateInfo{
		sType:                   C.VkStructureType(StructureTypePipelineRasterizationStateCreateInfo),
		pNext:                   info.Next,
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
	Next                  unsafe.Pointer
	RasterizationSamples  SampleCountFlags
	SampleShadingEnable   bool
	MinSampleShading      float32
	SampleMask            []SampleMask
	AlphaToCoverageEnable bool
	AlphaToOneEnable      bool
}

func (info PipelineMultisampleStateCreateInfo) c() *C.VkPipelineMultisampleStateCreateInfo {
	size0 := uintptr(C.sizeof_VkPipelineMultisampleStateCreateInfo)
	size1 := uintptr(len(info.SampleMask)) * C.sizeof_VkSampleMask
	size := size0 + size1
	alloc := C.calloc(1, C.size_t(size))
	cinfo := (*C.VkPipelineMultisampleStateCreateInfo)(alloc)
	var sampleMask *C.VkSampleMask
	if info.SampleMask != nil {
		sampleMask = (*C.VkSampleMask)(unsafe.Pointer(uintptr(alloc) + size0))
		ucopy(unsafe.Pointer(sampleMask), unsafe.Pointer(&info.SampleMask), C.sizeof_VkSampleMask)
	}

	*cinfo = C.VkPipelineMultisampleStateCreateInfo{
		sType:                 C.VkStructureType(StructureTypePipelineMultisampleStateCreateInfo),
		pNext:                 info.Next,
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
	Next           unsafe.Pointer
	LogicOpEnable  bool
	LogicOp        LogicOp
	Attachments    []PipelineColorBlendAttachmentState
	BlendConstants [4]float32
}

func (info PipelineColorBlendStateCreateInfo) c() *C.VkPipelineColorBlendStateCreateInfo {
	size0 := uintptr(C.sizeof_VkPipelineColorBlendStateCreateInfo)
	size1 := C.sizeof_VkPipelineColorBlendAttachmentState * uintptr(len(info.Attachments))
	size := size0 + size1
	alloc := C.calloc(1, C.size_t(size))
	cinfo := (*C.VkPipelineColorBlendStateCreateInfo)(alloc)
	attachments := (*C.VkPipelineColorBlendAttachmentState)(unsafe.Pointer(uintptr(alloc) + size0))
	*cinfo = C.VkPipelineColorBlendStateCreateInfo{
		sType:           C.VkStructureType(StructureTypePipelineColorBlendStateCreateInfo),
		pNext:           info.Next,
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
	attachmentsArr := (*[math.MaxInt32]C.VkPipelineColorBlendAttachmentState)(unsafe.Pointer(attachments))[:len(info.Attachments)]
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
	Next          unsafe.Pointer
	DynamicStates []DynamicState
}

func (info PipelineDynamicStateCreateInfo) c() *C.VkPipelineDynamicStateCreateInfo {
	size0 := uintptr(C.sizeof_VkPipelineDynamicStateCreateInfo)
	size1 := C.sizeof_VkDynamicState * uintptr(len(info.DynamicStates))
	size := size0 + size1
	alloc := C.calloc(1, C.size_t(size))
	cinfo := (*C.VkPipelineDynamicStateCreateInfo)(alloc)
	dynamicStates := (*C.VkDynamicState)(unsafe.Pointer(uintptr(alloc) + size0))
	*cinfo = C.VkPipelineDynamicStateCreateInfo{
		sType:             C.VkStructureType(StructureTypePipelineDynamicStateCreateInfo),
		pNext:             info.Next,
		flags:             0,
		dynamicStateCount: C.uint32_t(len(info.DynamicStates)),
		pDynamicStates:    dynamicStates,
	}
	ucopy(unsafe.Pointer(dynamicStates), unsafe.Pointer(&info.DynamicStates), C.sizeof_VkDynamicState)
	return cinfo
}

type PipelineLayout struct {
	// VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkPipelineLayout)
	hnd C.VkPipelineLayout
}

type PipelineLayoutCreateInfo struct {
	Next               unsafe.Pointer
	SetLayouts         []DescriptorSetLayout
	PushConstantRanges []PushConstantRange
}

func (info PipelineLayoutCreateInfo) c() *C.VkPipelineLayoutCreateInfo {
	size0 := uintptr(C.sizeof_VkPipelineLayoutCreateInfo)
	size1 := C.sizeof_VkDescriptorSetLayout * uintptr(len(info.SetLayouts))
	size2 := C.sizeof_VkPushConstantRange * uintptr(len(info.PushConstantRanges))
	size := size0 + size1 + size2
	alloc := C.calloc(1, C.size_t(size))
	cinfo := (*C.VkPipelineLayoutCreateInfo)(alloc)
	setLayouts := (*C.VkDescriptorSetLayout)(unsafe.Pointer(uintptr(alloc) + size0))
	push := (*C.VkPushConstantRange)(unsafe.Pointer(uintptr(alloc) + size0 + size1))
	*cinfo = C.VkPipelineLayoutCreateInfo{
		sType:                  C.VkStructureType(StructureTypePipelineLayoutCreateInfo),
		pNext:                  info.Next,
		flags:                  0,
		setLayoutCount:         C.uint32_t(len(info.SetLayouts)),
		pSetLayouts:            setLayouts,
		pushConstantRangeCount: C.uint32_t(len(info.PushConstantRanges)),
		pPushConstantRanges:    push,
	}
	ucopy(unsafe.Pointer(setLayouts), unsafe.Pointer(&info.SetLayouts), C.sizeof_VkDescriptorSetLayout)
	ucopy(unsafe.Pointer(push), unsafe.Pointer(&info.PushConstantRanges), C.sizeof_VkPushConstantRange)
	return cinfo
}

func (dev *Device) CreatePipelineLayout(info *PipelineLayoutCreateInfo) (PipelineLayout, error) {
	// TODO(dh): support custom allocators
	cinfo := info.c()
	defer C.free(unsafe.Pointer(cinfo))
	var hnd C.VkPipelineLayout
	res := Result(C.domVkCreatePipelineLayout(dev.fps[vkCreatePipelineLayout], dev.hnd, cinfo, nil, &hnd))
	if res != Success {
		return PipelineLayout{}, res
	}
	return PipelineLayout{hnd}, nil
}

type DescriptorSetLayout struct {
	// VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkDescriptorSetLayout)
	hnd C.VkDescriptorSetLayout
}

type PushConstantRange struct {
	StageFlags ShaderStageFlags
	Offset     uint32
	Size       uint32
}

type PipelineTessellationStateCreateInfo struct {
	Next               unsafe.Pointer
	PatchControlPoints uint32
}

func (info PipelineTessellationStateCreateInfo) c() *C.VkPipelineTessellationStateCreateInfo {
	cinfo := (*C.VkPipelineTessellationStateCreateInfo)(C.calloc(1, C.sizeof_VkPipelineTessellationStateCreateInfo))
	*cinfo = C.VkPipelineTessellationStateCreateInfo{
		sType:              C.VkStructureType(StructureTypePipelineTessellationStateCreateInfo),
		pNext:              info.Next,
		flags:              0,
		patchControlPoints: C.uint32_t(info.PatchControlPoints),
	}
	return cinfo
}

type PipelineDepthStencilStateCreateInfo struct {
	Next                  unsafe.Pointer
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

func (info PipelineDepthStencilStateCreateInfo) c() *C.VkPipelineDepthStencilStateCreateInfo {
	cinfo := (*C.VkPipelineDepthStencilStateCreateInfo)(C.calloc(1, C.sizeof_VkPipelineDepthStencilStateCreateInfo))
	*cinfo = C.VkPipelineDepthStencilStateCreateInfo{
		sType:                 C.VkStructureType(StructureTypePipelineDepthStencilStateCreateInfo),
		pNext:                 info.Next,
		flags:                 0,
		depthTestEnable:       vkBool(info.DepthTestEnable),
		depthWriteEnable:      vkBool(info.DepthWriteEnable),
		depthCompareOp:        C.VkCompareOp(info.DepthCompareOp),
		depthBoundsTestEnable: vkBool(info.DepthBoundsTestEnable),
		stencilTestEnable:     vkBool(info.StencilTestEnable),
		front:                 *(*C.VkStencilOpState)(unsafe.Pointer(&info.Front)),
		back:                  *(*C.VkStencilOpState)(unsafe.Pointer(&info.Back)),
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
	Next               unsafe.Pointer
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
	ptrs := (*C.VkGraphicsPipelineCreateInfo)(C.calloc(C.size_t(len(infos)), C.sizeof_VkGraphicsPipelineCreateInfo))
	defer C.free(unsafe.Pointer(ptrs))

	ptrsArr := (*[math.MaxInt32]C.VkGraphicsPipelineCreateInfo)(unsafe.Pointer(ptrs))[:len(infos)]
	for i := range ptrsArr {
		ptr := &ptrsArr[i]
		info := &infos[i]

		ptr.sType = C.VkStructureType(StructureTypeGraphicsPipelineCreateInfo)
		ptr.pNext = info.Next
		ptr.flags = C.VkPipelineCreateFlags(info.Flags)
		ptr.stageCount = C.uint32_t(len(info.Stages))

		ptr.pStages = (*C.VkPipelineShaderStageCreateInfo)(C.calloc(C.size_t(len(info.Stages)), C.sizeof_VkPipelineShaderStageCreateInfo))
		defer C.free(unsafe.Pointer(ptr.pStages))
		arr := (*[math.MaxInt32]C.VkPipelineShaderStageCreateInfo)(unsafe.Pointer(ptr.pStages))[:len(info.Stages)]
		for i := range arr {
			arr[i] = C.VkPipelineShaderStageCreateInfo{
				sType:  C.VkStructureType(StructureTypePipelineShaderStageCreateInfo),
				pNext:  info.Stages[i].Next,
				stage:  C.VkShaderStageFlagBits(info.Stages[i].Stage),
				module: info.Stages[i].Module.hnd,
				pName:  C.CString(info.Stages[i].Name),
			}
			defer C.free(unsafe.Pointer(arr[i].pName))
		}

		if info.VertexInputState != nil {
			ptr.pVertexInputState = info.VertexInputState.c()
			defer C.free(unsafe.Pointer(ptr.pVertexInputState))
		}
		if info.InputAssemblyState != nil {
			ptr.pInputAssemblyState = info.InputAssemblyState.c()
			defer C.free(unsafe.Pointer(ptr.pInputAssemblyState))
		}
		if info.TessellationState != nil {
			ptr.pTessellationState = info.TessellationState.c()
			defer C.free(unsafe.Pointer(ptr.pTessellationState))
		}
		if info.ViewportState != nil {
			ptr.pViewportState = info.ViewportState.c()
			defer C.free(unsafe.Pointer(ptr.pViewportState))
		}
		if info.RasterizationState != nil {
			ptr.pRasterizationState = info.RasterizationState.c()
			defer C.free(unsafe.Pointer(ptr.pRasterizationState))
		}
		if info.MultisampleState != nil {
			ptr.pMultisampleState = info.MultisampleState.c()
			defer C.free(unsafe.Pointer(ptr.pMultisampleState))
		}
		if info.DepthStencilState != nil {
			ptr.pDepthStencilState = info.DepthStencilState.c()
			defer C.free(unsafe.Pointer(ptr.pDepthStencilState))
		}
		if info.ColorBlendState != nil {
			ptr.pColorBlendState = info.ColorBlendState.c()
			defer C.free(unsafe.Pointer(ptr.pColorBlendState))
		}
		if info.DynamicState != nil {
			ptr.pDynamicState = info.DynamicState.c()
			defer C.free(unsafe.Pointer(ptr.pDynamicState))
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
	res := Result(C.domVkCreateGraphicsPipelines(dev.fps[vkCreateGraphicsPipelines], dev.hnd, 0, C.uint32_t(len(infos)), ptrs, nil, &hnds[0]))
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
}

type AttachmentReference struct {
	Attachment uint32
	Layout     ImageLayout
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
	Next         unsafe.Pointer
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
}

func (dev *Device) CreateRenderPass(info *RenderPassCreateInfo) (RenderPass, error) {
	// TODO(dh): support custom allocators
	size0 := uintptr(C.sizeof_VkRenderPassCreateInfo)
	size1 := C.sizeof_VkAttachmentDescription * uintptr(len(info.Attachments))
	size2 := C.sizeof_VkSubpassDescription * uintptr(len(info.Subpasses))
	size3 := C.sizeof_VkSubpassDependency * uintptr(len(info.Dependencies))
	size := size0 + size1 + size2 + size3
	alloc := C.calloc(1, C.size_t(size))
	defer C.free(alloc)
	cinfo := (*C.VkRenderPassCreateInfo)(alloc)
	attachments := (*C.VkAttachmentDescription)(unsafe.Pointer(uintptr(alloc) + size0))
	subpasses := (*C.VkSubpassDescription)(unsafe.Pointer(uintptr(alloc) + size0 + size1))
	dependencies := (*C.VkSubpassDependency)(unsafe.Pointer(uintptr(alloc) + size0 + size1 + size2))
	*cinfo = C.VkRenderPassCreateInfo{
		sType:           C.VkStructureType(StructureTypeRenderPassCreateInfo),
		pNext:           info.Next,
		flags:           0,
		attachmentCount: C.uint32_t(len(info.Attachments)),
		pAttachments:    attachments,
		subpassCount:    C.uint32_t(len(info.Subpasses)),
		pSubpasses:      subpasses,
		dependencyCount: C.uint32_t(len(info.Dependencies)),
		pDependencies:   dependencies,
	}
	ucopy(unsafe.Pointer(attachments), unsafe.Pointer(&info.Attachments), C.sizeof_VkAttachmentDescription)
	subpassesArr := (*[math.MaxInt32]C.VkSubpassDescription)(unsafe.Pointer(subpasses))[:len(info.Subpasses)]
	for i := range subpassesArr {
		subpass := &info.Subpasses[i]
		csubpass := &subpassesArr[i]
		*csubpass = C.VkSubpassDescription{
			flags:                   C.VkSubpassDescriptionFlags(subpass.Flags),
			pipelineBindPoint:       C.VkPipelineBindPoint(subpass.PipelineBindPoint),
			inputAttachmentCount:    C.uint32_t(len(subpass.InputAttachments)),
			colorAttachmentCount:    C.uint32_t(len(subpass.ColorAttachments)),
			preserveAttachmentCount: C.uint32_t(len(subpass.PreserveAttachments)),
			pInputAttachments:       (*C.VkAttachmentReference)(calloc(C.size_t(len(subpass.InputAttachments)), C.sizeof_VkAttachmentReference)),
			pColorAttachments:       (*C.VkAttachmentReference)(calloc(C.size_t(len(subpass.ColorAttachments)), C.sizeof_VkAttachmentReference)),
			pPreserveAttachments:    (*C.uint32_t)(calloc(C.size_t(len(subpass.PreserveAttachments)), C.sizeof_uint32_t)),
		}
		ucopy(unsafe.Pointer(csubpass.pInputAttachments), unsafe.Pointer(&subpass.InputAttachments), C.sizeof_VkAttachmentReference)
		ucopy(unsafe.Pointer(csubpass.pColorAttachments), unsafe.Pointer(&subpass.ColorAttachments), C.sizeof_VkAttachmentReference)
		if len(subpass.ResolveAttachments) > 0 {
			csubpass.pResolveAttachments = (*C.VkAttachmentReference)(calloc(C.size_t(len(subpass.ResolveAttachments)), C.sizeof_VkAttachmentReference))
			defer C.free(unsafe.Pointer(csubpass.pResolveAttachments))
			ucopy(unsafe.Pointer(csubpass.pResolveAttachments), unsafe.Pointer(&subpass.ResolveAttachments), C.sizeof_VkAttachmentReference)
		}
		if subpass.DepthStencilAttachment != nil {
			csubpass.pDepthStencilAttachment = (*C.VkAttachmentReference)(calloc(1, C.sizeof_VkAttachmentReference))
			ucopy1(unsafe.Pointer(csubpass.pDepthStencilAttachment), unsafe.Pointer(&subpass.DepthStencilAttachment), C.sizeof_VkAttachmentReference)
		}
		ucopy(unsafe.Pointer(csubpass.pPreserveAttachments), unsafe.Pointer(&subpass.PreserveAttachments), C.sizeof_uint32_t)
	}
	ucopy(unsafe.Pointer(dependencies), unsafe.Pointer(&info.Dependencies), C.sizeof_VkSubpassDependency)
	var hnd C.VkRenderPass
	res := Result(C.domVkCreateRenderPass(dev.fps[vkCreateRenderPass], dev.hnd, cinfo, nil, &hnd))
	if res != Success {
		return RenderPass{}, res
	}
	return RenderPass{hnd: hnd}, nil
}

type FramebufferCreateInfo struct {
	Next        unsafe.Pointer
	RenderPass  RenderPass
	Attachments []ImageView
	Width       uint32
	Height      uint32
	Layers      uint32
}

func (info FramebufferCreateInfo) c() *C.VkFramebufferCreateInfo {
	size0 := uintptr(C.sizeof_VkFramebufferCreateInfo)
	size1 := uintptr(C.sizeof_VkImageView) * uintptr(len(info.Attachments))
	size := size0 + size1
	alloc := C.calloc(1, C.size_t(size))
	cinfo := (*C.VkFramebufferCreateInfo)(alloc)
	*cinfo = C.VkFramebufferCreateInfo{
		sType:           C.VkStructureType(StructureTypeFramebufferCreateInfo),
		pNext:           info.Next,
		flags:           0,
		renderPass:      info.RenderPass.hnd,
		attachmentCount: C.uint32_t(len(info.Attachments)),
		pAttachments:    (*C.VkImageView)(unsafe.Pointer(uintptr(alloc) + size0)),
		width:           C.uint32_t(info.Width),
		height:          C.uint32_t(info.Height),
		layers:          C.uint32_t(info.Layers),
	}
	ucopy(unsafe.Pointer(cinfo.pAttachments), unsafe.Pointer(&info.Attachments), C.sizeof_VkImageView)
	return cinfo
}

func (dev *Device) CreateFramebuffer(info *FramebufferCreateInfo) (Framebuffer, error) {
	// TODO(dh): support custom allocators
	cinfo := info.c()
	defer C.free(unsafe.Pointer(cinfo))
	var hnd C.VkFramebuffer
	res := Result(C.domVkCreateFramebuffer(dev.fps[vkCreateFramebuffer], dev.hnd, cinfo, nil, &hnd))
	if res != Success {
		return Framebuffer{}, res
	}
	return Framebuffer{hnd}, nil
}

type RenderPassBeginInfo struct {
	Next        unsafe.Pointer
	RenderPass  RenderPass
	Framebuffer Framebuffer
	RenderArea  Rect2D
	ClearValues []ClearValue
}

type ClearValue interface {
	isClearValue()
}

type ClearColorValueFloat32s [4]float32
type ClearColorValueInt32s [4]int32
type ClearColorValueUint32s [4]uint32

type ClearDepthStencilValue struct {
	Depth   float32
	Stencil uint32
}

func (ClearColorValueFloat32s) isClearValue() {}
func (ClearColorValueInt32s) isClearValue()   {}
func (ClearColorValueUint32s) isClearValue()  {}
func (ClearDepthStencilValue) isClearValue()  {}

type Semaphore struct {
	// VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkSemaphore)
	hnd C.VkSemaphore
}

type SemaphoreCreateInfo struct {
	Next unsafe.Pointer
}

func (info SemaphoreCreateInfo) c() *C.VkSemaphoreCreateInfo {
	cinfo := (*C.VkSemaphoreCreateInfo)(C.calloc(1, C.sizeof_VkSemaphoreCreateInfo))
	*cinfo = C.VkSemaphoreCreateInfo{
		sType: C.VkStructureType(StructureTypeSemaphoreCreateInfo),
		pNext: info.Next,
	}
	return cinfo
}

func (dev *Device) CreateSemaphore(info *SemaphoreCreateInfo) (Semaphore, error) {
	// TODO(dh): support custom allocators
	cinfo := info.c()
	defer C.free(unsafe.Pointer(cinfo))
	var hnd C.VkSemaphore
	res := Result(C.domVkCreateSemaphore(dev.fps[vkCreateSemaphore], dev.hnd, cinfo, nil, &hnd))
	if res != Success {
		return Semaphore{}, res
	}
	return Semaphore{hnd}, nil
}

type SubmitInfo struct {
	Next             unsafe.Pointer
	WaitSemaphores   []Semaphore
	WaitDstStageMask []PipelineStageFlags
	CommandBuffers   []*CommandBuffer
	SignalSemaphores []Semaphore
}

func (queue *Queue) Submit(infos []SubmitInfo) error {
	// TODO(dh): support fence
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
	size0 := C.sizeof_VkSubmitInfo * uintptr(len(infos))
	size1 := C.sizeof_VkSemaphore * waitSemaphoreCount
	size2 := C.sizeof_VkPipelineStageFlags * waitSemaphoreCount
	size3 := C.sizeof_VkCommandBuffer * commandBufferCount
	size4 := C.sizeof_VkSemaphore * signalSemaphoreCount
	size := size0 + size1 + size2 + size3 + size4
	alloc := uintptr(C.calloc(1, C.size_t(size)))
	defer C.free(unsafe.Pointer(alloc))

	cinfos := alloc
	waitSemaphores := alloc + size0
	waitDstStageMask := alloc + size0 + size1
	commandBuffers := alloc + size0 + size1 + size2
	signalSemaphores := alloc + size0 + size1 + size2 + size3

	for _, info := range infos {
		if len(info.WaitSemaphores) != len(info.WaitDstStageMask) {
			panic("WaitSemaphores and WaitDstStageMask must have same length")
		}
		*(*C.VkSubmitInfo)(unsafe.Pointer(cinfos)) = C.VkSubmitInfo{
			sType:                C.VkStructureType(StructureTypeSubmitInfo),
			pNext:                info.Next,
			waitSemaphoreCount:   C.uint32_t(len(info.WaitSemaphores)),
			pWaitSemaphores:      (*C.VkSemaphore)(unsafe.Pointer(waitSemaphores)),
			pWaitDstStageMask:    (*C.VkPipelineStageFlags)(unsafe.Pointer(waitDstStageMask)),
			commandBufferCount:   C.uint32_t(len(info.CommandBuffers)),
			pCommandBuffers:      (*C.VkCommandBuffer)(unsafe.Pointer(commandBuffers)),
			signalSemaphoreCount: C.uint32_t(len(info.SignalSemaphores)),
			pSignalSemaphores:    (*C.VkSemaphore)(unsafe.Pointer(signalSemaphores)),
		}
		ucopy(unsafe.Pointer(waitSemaphores), unsafe.Pointer(&info.WaitSemaphores), C.sizeof_VkSemaphore)
		ucopy(unsafe.Pointer(waitDstStageMask), unsafe.Pointer(&info.WaitDstStageMask), C.sizeof_VkPipelineStageFlags)
		ucopy(unsafe.Pointer(signalSemaphores), unsafe.Pointer(&info.SignalSemaphores), C.sizeof_VkSemaphore)
		arr := (*[math.MaxInt32]C.VkCommandBuffer)(unsafe.Pointer(commandBuffers))[:len(info.CommandBuffers)]
		for i := range arr {
			arr[i] = info.CommandBuffers[i].hnd
		}

		cinfos += C.sizeof_VkSubmitInfo
		waitSemaphores += C.sizeof_VkSemaphore * uintptr(len(info.WaitSemaphores))
		waitDstStageMask += C.sizeof_VkPipelineStageFlags * uintptr(len(info.WaitSemaphores))
		commandBuffers += C.sizeof_VkCommandBuffer * uintptr(len(info.CommandBuffers))
		signalSemaphores += C.sizeof_VkSemaphore * uintptr(len(info.SignalSemaphores))
	}

	res := Result(C.domVkQueueSubmit(queue.fps[vkQueueSubmit], queue.hnd, C.uint32_t(len(infos)), (*C.VkSubmitInfo)(unsafe.Pointer(alloc)), nil))
	if res != Success {
		return res
	}
	return nil
}

func calloc(nmemb C.size_t, size C.size_t) unsafe.Pointer {
	if nmemb == 0 {
		return nil
	}
	return C.calloc(nmemb, size)
}

// ucopy copies data from src to dst,
// where dst must be a C pointer and src must be a pointer to a Go slice.
func ucopy(dst, src unsafe.Pointer, size uintptr) {
	elems := *(*int)(unsafe.Pointer(uintptr(src) + unsafe.Sizeof(uintptr(0))))
	if elems == 0 {
		return
	}
	// Access the slice's underlying data
	src = (*(*unsafe.Pointer)(src))
	copy(
		(*[math.MaxInt32]byte)(dst)[:uintptr(elems)*size],
		(*[math.MaxInt32]byte)(src)[:uintptr(elems)*size],
	)
}

func ucopy1(dst, src unsafe.Pointer, size uintptr) {
	copy(
		(*[math.MaxInt32]byte)(dst)[:size],
		(*[math.MaxInt32]byte)(src)[:size],
	)
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

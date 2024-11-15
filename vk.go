// Copyright (c) 2015-2018 Khronos Group
// Copyright (c) 2018 Dominik Honnef

package vk

// #cgo pkg-config: vulkan
//
// #cgo noescape domVkCreateInstance
// #cgo noescape domVkEnumeratePhysicalDevices
// #cgo noescape domVkGetPhysicalDeviceProperties2
// #cgo noescape domVkGetPhysicalDeviceMemoryProperties2
// #cgo noescape domVkGetPhysicalDeviceFeatures
// #cgo noescape domVkCreateDevice
// #cgo noescape domVkGetDeviceQueue2
// #cgo noescape domVkBeginCommandBuffer
// #cgo noescape domVkCmdClearColorImage
// #cgo noescape domVkCreateCommandPool
// #cgo noescape domVkAllocateCommandBuffers
// #cgo noescape domVkCreateImageView
// #cgo noescape domVkCreateShaderModule
// #cgo noescape domVkCreateFence
// #cgo noescape domVkGetBufferMemoryRequirements2
// #cgo noescape domVkGetImageMemoryRequirements2
// #cgo noescape domVkGetPhysicalDeviceFormatProperties2
// #cgo noescape domVkGetPhysicalDeviceImageFormatProperties2
// #cgo noescape domVkCreateSwapchainKHR
//
// #cgo nocallback domVkCreateInstance
// #cgo nocallback domVkEnumeratePhysicalDevices
// #cgo nocallback domVkGetPhysicalDeviceProperties2
// #cgo nocallback domVkGetPhysicalDeviceMemoryProperties2
// #cgo nocallback domVkGetPhysicalDeviceFeatures
// #cgo nocallback domVkCreateDevice
// #cgo nocallback domVkGetDeviceQueue2
// #cgo nocallback domVkBeginCommandBuffer
// #cgo nocallback domVkCmdClearColorImage
// #cgo nocallback domVkCreateCommandPool
// #cgo nocallback domVkAllocateCommandBuffers
// #cgo nocallback domVkCreateImageView
// #cgo nocallback domVkCreateShaderModule
// #cgo nocallback domVkCreateFence
// #cgo nocallback domVkGetBufferMemoryRequirements2
// #cgo nocallback domVkGetImageMemoryRequirements2
// #cgo nocallback domVkGetPhysicalDeviceFormatProperties2
// #cgo nocallback domVkGetPhysicalDeviceImageFormatProperties2
// #cgo nocallback domVkCreateSwapchainKHR
//
// #include <stdlib.h>
// #include "vk.h"
import "C"
import (
	"bytes"
	"fmt"
	"math"
	"os"
	"structs"
	"time"
	"unsafe"

	"honnef.co/go/safeish"
)

// OPT(dh): if we wrote our own memory allocator, we could avoid the
// significant overhead of calling malloc and free.

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
	APIVersion10 = Version(C.VK_API_VERSION_1_0)
	// Vulkan 1.1 version number
	APIVersion11 = Version(C.VK_API_VERSION_1_1)
)

var vkEnumerateInstanceVersion C.PFN_vkEnumerateInstanceVersion
var vkEnumerateInstanceExtensionProperties C.PFN_vkEnumerateInstanceExtensionProperties
var vkEnumerateInstanceLayerProperties C.PFN_vkEnumerateInstanceLayerProperties
var vkCreateInstance C.PFN_vkCreateInstance

var _ = "_"[unsafe.Sizeof(AttachmentDescription{})-C.sizeof_VkAttachmentDescription]
var _ = "_"[unsafe.Sizeof(AttachmentReference{})-C.sizeof_VkAttachmentReference]
var _ = "_"[unsafe.Sizeof(DescriptorSetLayout{})-C.sizeof_VkDescriptorSetLayout]
var _ = "_"[unsafe.Sizeof(Fence{})-C.sizeof_VkFence]
var _ = "_"[unsafe.Sizeof(ImageView{})-C.sizeof_VkImageView]
var _ = "_"[unsafe.Sizeof(MemoryHeap{})-C.sizeof_VkMemoryHeap]
var _ = "_"[unsafe.Sizeof(MemoryRequirements{})-C.sizeof_VkMemoryRequirements]
var _ = "_"[unsafe.Sizeof(MemoryType{})-C.sizeof_VkMemoryType]
var _ = "_"[unsafe.Sizeof(PushConstantRange{})-C.sizeof_VkPushConstantRange]
var _ = "_"[unsafe.Sizeof(Rect2D{})-C.sizeof_VkRect2D]
var _ = "_"[unsafe.Sizeof(Semaphore{})-C.sizeof_VkSemaphore]
var _ = "_"[unsafe.Sizeof(SubpassDependency{})-C.sizeof_VkSubpassDependency]
var _ = "_"[unsafe.Sizeof(VertexInputAttributeDescription{})-C.sizeof_VkVertexInputAttributeDescription]
var _ = "_"[unsafe.Sizeof(VertexInputBindingDescription{})-C.sizeof_VkVertexInputBindingDescription]
var _ = "_"[unsafe.Sizeof(Viewport{})-C.sizeof_VkViewport]
var _ = "_"[unsafe.Sizeof(ComponentMapping{})-C.sizeof_VkComponentMapping]
var _ = "_"[unsafe.Sizeof(ImageSubresourceRange{})-C.sizeof_VkImageSubresourceRange]
var _ = "_"[unsafe.Sizeof(ClearDepthStencilValue{})-C.sizeof_VkClearDepthStencilValue]
var _ = "_"[unsafe.Sizeof(BufferCopy{})-C.sizeof_VkBufferCopy]
var _ = "_"[unsafe.Sizeof(BufferImageCopy{})-C.sizeof_VkBufferImageCopy]
var _ = "_"[unsafe.Sizeof(ImageSubresourceLayers{})-C.sizeof_VkImageSubresourceLayers]
var _ = "_"[unsafe.Sizeof(ImageCopy{})-C.sizeof_VkImageCopy]
var _ = "_"[unsafe.Sizeof(ImageBlit{})-C.sizeof_VkImageBlit]
var _ = "_"[unsafe.Sizeof(Event{})-C.sizeof_VkEvent]
var _ = "_"[unsafe.Sizeof(ImageResolve{})-C.sizeof_VkImageResolve]
var _ = "_"[unsafe.Sizeof(DescriptorPoolSize{})-C.sizeof_VkDescriptorPoolSize]
var _ = "_"[unsafe.Sizeof(DescriptorSet{})-C.sizeof_VkDescriptorSet]
var _ = "_"[unsafe.Sizeof(DescriptorBufferInfo{})-C.sizeof_VkDescriptorBufferInfo]
var _ = "_"[unsafe.Sizeof(DescriptorImageInfo{})-C.sizeof_VkDescriptorImageInfo]
var _ = "_"[unsafe.Sizeof(FormatProperties{})-C.sizeof_VkFormatProperties]

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

// MakeVersion constructs an API version number.
func MakeVersion(major, minor, patch uint32) Version {
	return Version(major<<22 | minor<<12 | patch)
}

type Version uint32

func (v Version) String() string {
	// return major<<22 | minor<<12 | patch
	major := (v >> 22)
	minor := (v >> 12) & 0x3FF
	patch := v & 0xFFF
	return fmt.Sprintf("%d.%d.%d", major, minor, patch)
}

func InstanceVersion() Version {
	var v Version
	C.domVkEnumerateInstanceVersion(vkEnumerateInstanceVersion, (*C.uint32_t)(unsafe.Pointer(&v)))
	return v
}

type LayerProperties struct {
	LayerName             string
	SpecVersion           Version
	ImplementationVersion uint32
	Description           string
}

func InstanceLayerProperties() ([]LayerProperties, error) {
	for {
		var count C.uint32_t
		res := Result(C.domVkEnumerateInstanceLayerProperties(vkEnumerateInstanceLayerProperties, &count, nil))
		if res != Success {
			return nil, res
		}
		cprops := make([]C.VkLayerProperties, count)
		res = Result(C.domVkEnumerateInstanceLayerProperties(vkEnumerateInstanceLayerProperties, &count, unsafe.SliceData(cprops)))
		if res == Success {
			out := make([]LayerProperties, count)
			cprops = cprops[:count]
			for i := range cprops {
				out[i] = LayerProperties{
					LayerName:             str(cprops[i].layerName[:]),
					SpecVersion:           Version(cprops[i].specVersion),
					ImplementationVersion: uint32(cprops[i].implementationVersion),
					Description:           str(cprops[i].description[:]),
				}
			}
			return out, nil
		}
		if res == Incomplete {
			continue
		}
		return nil, res
	}
}

func InstanceExtensionProperties(layerName string) ([]ExtensionProperties, error) {
	var cname *C.char
	if layerName != "" {
		cname = C.CString(layerName)
		defer C.free(unsafe.Pointer(cname))
	}

	for {
		var count C.uint32_t
		res := Result(C.domVkEnumerateInstanceExtensionProperties(vkEnumerateInstanceExtensionProperties, cname, &count, nil))
		if res != Success {
			return nil, res
		}
		cprops := make([]C.VkExtensionProperties, count)
		res = Result(C.domVkEnumerateInstanceExtensionProperties(vkEnumerateInstanceExtensionProperties, cname, &count, unsafe.SliceData(cprops)))
		if res == Success {
			out := make([]ExtensionProperties, count)
			cprops = cprops[:count]
			for i := range cprops {
				out[i] = ExtensionProperties{
					Name:        str(cprops[i].extensionName[:]),
					SpecVersion: uint32(cprops[i].specVersion),
				}
			}
			return out, nil
		}
		if res == Incomplete {
			continue
		}
		return nil, res
	}
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
	ApplicationVersion Version
	// The name of the engine (if any) used to create the application
	EngineName string
	// The developer-supplied version number of the engine used to create the application
	EngineVersion Version
	// The highest version of Vulkan that the application is designed to use
	APIVersion Version
}

// There is no global state in Vulkan and all per-application state is stored in an Instance object.
// Creating an Instance object initializes the Vulkan library
// and allows the application to pass information about itself to the implementation.
type Instance struct {
	// VK_DEFINE_HANDLE(VkInstance)
	hnd C.VkInstance
	fps [instanceMaxPFN]C.PFN_vkVoidFunction
}

func CreateInstance(info *InstanceCreateInfo) (*Instance, error) {
	a := new(allocator)
	defer a.free()

	var ptr C.VkInstanceCreateInfo
	ptr.sType = C.VkStructureType(StructureTypeInstanceCreateInfo)
	ptr.pNext = buildChain(a, info.Extensions)
	defer internalizeChain(info.Extensions, ptr.pNext)
	ptr.enabledLayerCount = C.uint32_t(len(info.EnabledLayerNames))
	ptr.enabledExtensionCount = C.uint32_t(len(info.EnabledExtensionNames))
	ptr.ppEnabledLayerNames = externStrings(a, info.EnabledLayerNames)
	ptr.ppEnabledExtensionNames = externStrings(a, info.EnabledExtensionNames)
	if info.ApplicationInfo != nil {
		ptr.pApplicationInfo = alloc[C.VkApplicationInfo](a)
		ptr.pApplicationInfo.sType = C.VkStructureType(StructureTypeApplicationInfo)
		ptr.pApplicationInfo.pNext = buildChain(a, info.ApplicationInfo.Extensions)
		ptr.pApplicationInfo.pApplicationName = externString(a, info.ApplicationInfo.ApplicationName)
		ptr.pApplicationInfo.applicationVersion = C.uint32_t(info.ApplicationInfo.ApplicationVersion)
		ptr.pApplicationInfo.pEngineName = externString(a, info.ApplicationInfo.EngineName)
		ptr.pApplicationInfo.engineVersion = C.uint32_t(info.ApplicationInfo.EngineVersion)
		ptr.pApplicationInfo.apiVersion = C.uint32_t(info.ApplicationInfo.APIVersion)
		defer internalizeChain(info.ApplicationInfo.Extensions, ptr.pApplicationInfo.pNext)
	}

	var instance C.VkInstance
	res := Result(C.domVkCreateInstance(vkCreateInstance, &ptr, nil, &instance))
	if res != Success {
		return nil, res
	}

	out := &Instance{hnd: instance}
	out.init()

	return out, nil
}

func (ins *Instance) Destroy() {
	// TODO(dh): support a custom allocator
	C.domVkDestroyInstance(ins.fps[vkDestroyInstance], ins.hnd, nil)
}

func (ins *Instance) init() {
	for i, name := range instanceFpNames {
		ins.fps[i] = vkGetInstanceProcAddr(ins.hnd, name)
	}
}

func (ins *Instance) PhysicalDevices() ([]*PhysicalDevice, error) {
	count := C.uint32_t(1)
	var devs []C.VkPhysicalDevice
	for {
		devs = make([]C.VkPhysicalDevice, count)
		res := Result(C.domVkEnumeratePhysicalDevices(ins.fps[vkEnumeratePhysicalDevices], ins.hnd, &count, unsafe.SliceData(devs)))
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
	for _, dev := range devs {
		out = append(out, &PhysicalDevice{dev, ins})
	}
	return out, nil
}

// Vulkan separates the concept of physical and logical devices.
// A physical device usually represents a single complete implementation of Vulkan
// (excluding instance-level functionality)
// available to the host, of which there are a finite number.
// A logical device represents an instance of that implementation
// with its own state and resources independent of other logical devices.
type PhysicalDevice struct {
	// VK_DEFINE_HANDLE(VkPhysicalDevice)
	hnd      C.VkPhysicalDevice
	instance *Instance
}

type PhysicalDeviceLimits struct {
	// MaxImageDimension1D is the maximum dimension (width) supported
	// for all images created with an imageType of ImageType1D.
	MaxImageDimension1D uint32
	// MaxImageDimension2D is the maximum dimension (width or height)
	// supported for all images created with an imageType of
	// ImageType2D and without ImageCreateCubeCompatibleBit set in
	// flags.
	MaxImageDimension2D uint32
	// MaxImageDimension3D is the maximum dimension (width, height, or
	// depth) supported for all images created with an imageType of
	// ImageType3D.
	MaxImageDimension3D uint32
	// MaxImageDimensionCube is the maximum dimension (width or
	// height) supported for all images created with an imageType of
	// ImageType2D and with ImageCreateCubeCompatibleBit set in flags.
	MaxImageDimensionCube uint32
	// MaxImageArrayLayers is the maximum number of layers
	// (arrayLayers) for an image.
	MaxImageArrayLayers uint32
	// MaxTexelBufferElements is the maximum number of addressable
	// texels for a buffer view created on a buffer which was created
	// with the BufferUsageUniformTexelBufferBit or
	// BufferUsageStorageTexelBufferBit set in the Usage field of the
	// BufferCreateInfo structure.
	MaxTexelBufferElements uint32
	// MaxUniformBufferRange is the maximum value that can be
	// specified in the Range field of any DescriptorBufferInfo
	// structures passed to a call to UpdateDescriptorSets for
	// descriptors of type DescriptorTypeUniformBuffer or
	// DescriptorTypeUniformBufferDynamic.
	MaxUniformBufferRange uint32
	// MaxStorageBufferRange is the maximum value that can be
	// specified in the Range field of any DescriptorBufferInfo
	// structures passed to a call to UpdateDescriptorSets for
	// descriptors of type DescriptorTypeStorageBuffer or
	// DescriptorTypeStorageBufferDynamic.
	MaxStorageBufferRange uint32
	// MaxPushConstantsSize is the maximum size, in bytes, of the pool
	// of push constant memory. For each of the push constant ranges
	// indicated by the PushConstantRanges field of the
	// PipelineLayoutCreateInfo structure, (offset + size) must be
	// less than or equal to this limit.
	MaxPushConstantsSize uint32
	// MaxMemoryAllocationCount is the maximum number of device memory
	// allocations, as created by AllocateMemory, which can
	// simultaneously exist.
	MaxMemoryAllocationCount uint32
	// MaxSamplerAllocationCount is the maximum number of sampler
	// objects, as created by CreateSampler, which can
	// simultaneously exist on a device.
	MaxSamplerAllocationCount uint32
	// BufferImageGranularity is the granularity, in bytes, at which
	// buffer or linear image resources, and optimal image resources
	// can be bound to adjacent offsets in the same DeviceMemory
	// object without aliasing. See Buffer-Image Granularity for more
	// details.
	BufferImageGranularity DeviceSize
	// SparseAddressSpaceSize is the total amount of address space
	// available, in bytes, for sparse memory resources. This is an
	// upper bound on the sum of the size of all sparse resources,
	// regardless of whether any memory is bound to them.
	SparseAddressSpaceSize DeviceSize
	// MaxBoundDescriptorSets is the maximum number of descriptor sets
	// that can be simultaneously used by a pipeline. All
	// DescriptorSet decorations in shader modules must have a value
	// less than maxBoundDescriptorSets.
	MaxBoundDescriptorSets uint32
	// MaxPerStageDescriptorSamplers is the maximum number of samplers
	// that can be accessible to a single shader stage in a pipeline
	// layout. Descriptors with a type of DescriptorTypeSampler or
	// DescriptorTypeCombinedImageSampler count against this limit.
	// Only descriptors in descriptor set layouts created without the
	// DescriptorSetLayoutCreateUpdateAfterBindPoolBitEXT bit set
	// count against this limit. A descriptor is accessible to a
	// shader stage when the StageFlags field of the
	// DescriptorSetLayoutBinding structure has the bit for that
	// shader stage set.
	MaxPerStageDescriptorSamplers uint32
	// MaxPerStageDescriptorUniformBuffers is the maximum number of
	// uniform buffers that can be accessible to a single shader stage
	// in a pipeline layout. Descriptors with a type of
	// DescriptorTypeUniformBuffer or
	// DescriptorTypeUniformBufferDynamic count against this limit.
	// Only descriptors in descriptor set layouts created without the
	// DescriptorSetLayoutCreateUpdateAfterBindPoolBitEXT bit set
	// count against this limit. A descriptor is accessible to a
	// shader stage when the StageFlags field of the
	// DescriptorSetLayoutBinding structure has the bit for that
	// shader stage set.
	MaxPerStageDescriptorUniformBuffers uint32
	// MaxPerStageDescriptorStorageBuffers is the maximum number of
	// storage buffers that can be accessible to a single shader stage
	// in a pipeline layout. Descriptors with a type of
	// DescriptorTypeStorageBuffer or
	// DescriptorTypeStorageBufferDynamic count against this limit.
	// Only descriptors in descriptor set layouts created without the
	// DescriptorSetLayoutCreateUpdateAfterBindPoolBitEXT bit set
	// count against this limit. A descriptor is accessible to a
	// pipeline shader stage when the StageFlags field of the
	// DescriptorSetLayoutBinding structure has the bit for that
	// shader stage set.
	MaxPerStageDescriptorStorageBuffers uint32
	// MaxPerStageDescriptorSampledImages is the maximum number of
	// sampled images that can be accessible to a single shader stage
	// in a pipeline layout. Descriptors with a type of
	// DescriptorTypeCombinedImageSampler, DescriptorTypeSampledImage,
	// or DescriptorTypeUniformTexelBuffer count against this limit.
	// Only descriptors in descriptor set layouts created without the
	// DescriptorSetLayoutCreateUpdateAfterBindPoolBitEXT bit set
	// count against this limit. A descriptor is accessible to a
	// pipeline shader stage when the StageFlags field of the
	// DescriptorSetLayoutBinding structure has the bit for that
	// shader stage set.
	MaxPerStageDescriptorSampledImages uint32
	// MaxPerStageDescriptorStorageImages is the maximum number of
	// storage images that can be accessible to a single shader stage
	// in a pipeline layout. Descriptors with a type of
	// DescriptorTypeStorageImage, or DescriptorTypeStorageTexelBuffer
	// count against this limit. Only descriptors in descriptor set
	// layouts created without the
	// DescriptorSetLayoutCreateUpdateAfterBindPoolBitEXT bit set
	// count against this limit. A descriptor is accessible to a
	// pipeline shader stage when the StageFlags field of the
	// DescriptorSetLayoutBinding structure has the bit for that
	// shader stage set.
	MaxPerStageDescriptorStorageImages uint32
	// MaxPerStageDescriptorInputAttachments is the maximum number of
	// input attachments that can be accessible to a single shader
	// stage in a pipeline layout. Descriptors with a type of
	// DescriptorTypeInputAttachment count against this limit. Only
	// descriptors in descriptor set layouts created without the
	// DescriptorSetLayoutCreateUpdateAfterBindPoolBitEXT bit set
	// count against this limit. A descriptor is accessible to a
	// pipeline shader stage when the StageFlags field of the
	// DescriptorSetLayoutBinding structure has the bit for that
	// shader stage set. These are only supported for the fragment
	// stage.
	MaxPerStageDescriptorInputAttachments uint32
	// MaxPerStageResources is the maximum number of resources that
	// can be accessible to a single shader stage in a pipeline
	// layout. Descriptors with a type of
	// DescriptorTypeCombinedImageSampler, DescriptorTypeSampledImage,
	// DescriptorTypeStorageImage, DescriptorTypeUniformTexelBuffer,
	// DescriptorTypeStorageTexelBuffer, DescriptorTypeUniformBuffer,
	// DescriptorTypeStorageBuffer,
	// DescriptorTypeUniformBufferDynamic,
	// DescriptorTypeStorageBufferDynamic, or
	// DescriptorTypeInputAttachment count against this limit. Only
	// descriptors in descriptor set layouts created without the
	// DescriptorSetLayoutCreateUpdateAfterBindPoolBitEXT bit set
	// count against this limit. For the fragment shader stage the
	// framebuffer color attachments also count against this limit.
	MaxPerStageResources uint32
	// MaxDescriptorSetSamplers is the maximum number of samplers that
	// can be included in descriptor bindings in a pipeline layout
	// across all pipeline shader stages and descriptor set numbers.
	// Descriptors with a type of DescriptorTypeSampler or
	// DescriptorTypeCombinedImageSampler count against this limit.
	// Only descriptors in descriptor set layouts created without the
	// DescriptorSetLayoutCreateUpdateAfterBindPoolBitEXT bit set
	// count against this limit.
	MaxDescriptorSetSamplers uint32
	// MaxDescriptorSetUniformBuffers is the maximum number of uniform
	// buffers that can be included in descriptor bindings in a
	// pipeline layout across all pipeline shader stages and
	// descriptor set numbers. Descriptors with a type of
	// DescriptorTypeUniformBuffer or
	// DescriptorTypeUniformBufferDynamic count against this limit.
	// Only descriptors in descriptor set layouts created without the
	// DescriptorSetLayoutCreateUpdateAfterBindPoolBitEXT bit set
	// count against this limit.
	MaxDescriptorSetUniformBuffers uint32
	// MaxDescriptorSetUniformBuffersDynamic is the maximum number of
	// dynamic uniform buffers that can be included in descriptor
	// bindings in a pipeline layout across all pipeline shader stages
	// and descriptor set numbers. Descriptors with a type of
	// DescriptorTypeUniformBufferDynamic count against this limit.
	// Only descriptors in descriptor set layouts created without the
	// DescriptorSetLayoutCreateUpdateAfterBindPoolBitEXT bit set
	// count against this limit.
	MaxDescriptorSetUniformBuffersDynamic uint32
	// MaxDescriptorSetStorageBuffers is the maximum number of storage
	// buffers that can be included in descriptor bindings in a
	// pipeline layout across all pipeline shader stages and
	// descriptor set numbers. Descriptors with a type of
	// DescriptorTypeStorageBuffer or
	// DescriptorTypeStorageBufferDynamic count against this limit.
	// Only descriptors in descriptor set layouts created without the
	// DescriptorSetLayoutCreateUpdateAfterBindPoolBitEXT bit set
	// count against this limit.
	MaxDescriptorSetStorageBuffers uint32
	// MaxDescriptorSetStorageBuffersDynamic is the maximum number of
	// dynamic storage buffers that can be included in descriptor
	// bindings in a pipeline layout across all pipeline shader stages
	// and descriptor set numbers. Descriptors with a type of
	// DescriptorTypeStorageBufferDynamic count against this limit.
	// Only descriptors in descriptor set layouts created without the
	// DescriptorSetLayoutCreateUpdateAfterBindPoolBitEXT bit set
	// count against this limit.
	MaxDescriptorSetStorageBuffersDynamic uint32
	// MaxDescriptorSetSampledImages is the maximum number of sampled
	// images that can be included in descriptor bindings in a
	// pipeline layout across all pipeline shader stages and
	// descriptor set numbers. Descriptors with a type of
	// DescriptorTypeCombinedImageSampler, DescriptorTypeSampledImage,
	// or DescriptorTypeUniformTexelBuffer count against this limit.
	// Only descriptors in descriptor set layouts created without the
	// DescriptorSetLayoutCreateUpdateAfterBindPoolBitEXT bit set
	// count against this limit.
	MaxDescriptorSetSampledImages uint32
	// MaxDescriptorSetStorageImages is the maximum number of storage
	// images that can be included in descriptor bindings in a
	// pipeline layout across all pipeline shader stages and
	// descriptor set numbers. Descriptors with a type of
	// DescriptorTypeStorageImage, or DescriptorTypeStorageTexelBuffer
	// count against this limit. Only descriptors in descriptor set
	// layouts created without the
	// DescriptorSetLayoutCreateUpdateAfterBindPoolBitEXT bit set
	// count against this limit.
	MaxDescriptorSetStorageImages uint32
	// MaxDescriptorSetInputAttachments is the maximum number of input
	// attachments that can be included in descriptor bindings in a
	// pipeline layout across all pipeline shader stages and
	// descriptor set numbers. Descriptors with a type of
	// DescriptorTypeInputAttachment count against this limit. Only
	// descriptors in descriptor set layouts created without the
	// DescriptorSetLayoutCreateUpdateAfterBindPoolBitEXT bit set
	// count against this limit.
	MaxDescriptorSetInputAttachments uint32
	// MaxVertexInputAttributes is the maximum number of vertex input
	// attributes that can be specified for a graphics pipeline. These
	// are described in the array of VertexInputAttributeDescription
	// structures that are provided at graphics pipeline creation time
	// via the VertexAttributeDescriptions field of the
	// PipelineVertexInputStateCreateInfo structure.
	MaxVertexInputAttributes uint32
	// MaxVertexInputBindings is the maximum number of vertex buffers
	// that can be specified for providing vertex attributes to a
	// graphics pipeline. These are described in the array of
	// VertexInputBindingDescription structures that are provided at
	// graphics pipeline creation time via the
	// VertexBindingDescriptions field of the
	// PipelineVertexInputStateCreateInfo structure. The Binding field
	// of VertexInputBindingDescription must be less than this limit.
	MaxVertexInputBindings uint32
	// MaxVertexInputAttributeOffset is the maximum vertex input
	// attribute offset that can be added to the vertex input binding
	// stride. The Offset field of the VertexInputAttributeDescription
	// structure must be less than or equal to this limit.
	MaxVertexInputAttributeOffset uint32
	// MaxVertexInputBindingStride is the maximum vertex input binding
	// stride that can be specified in a vertex input binding. The
	// Stride field of the VertexInputBindingDescription structure
	// must be less than or equal to this limit.
	MaxVertexInputBindingStride uint32
	// MaxVertexOutputComponents is the maximum number of components
	// of output variables which can be output by a vertex shader.
	MaxVertexOutputComponents uint32
	// MaxTessellationGenerationLevel is the maximum tessellation
	// generation level supported by the fixed-function tessellation
	// primitive generator.
	MaxTessellationGenerationLevel uint32
	// MaxTessellationPatchSize is the maximum patch size, in
	// vertices, of patches that can be processed by the tessellation
	// control shader and tessellation primitive generator. The
	// PatchControlPoints field of the
	// PipelineTessellationStateCreateInfo structure specified at
	// pipeline creation time and the value provided in the
	// OutputVertices execution mode of shader modules must be less
	// than or equal to this limit.
	MaxTessellationPatchSize uint32
	// MaxTessellationControlPerVertexInputComponents is the maximum
	// number of components of input variables which can be provided
	// as per-vertex inputs to the tessellation control shader stage.
	MaxTessellationControlPerVertexInputComponents uint32
	// MaxTessellationControlPerVertexOutputComponents is the maximum
	// number of components of per-vertex output variables which can
	// be output from the tessellation control shader stage.
	MaxTessellationControlPerVertexOutputComponents uint32
	// MaxTessellationControlPerPatchOutputComponents is the maximum
	// number of components of per-patch output variables which can be
	// output from the tessellation control shader stage.
	MaxTessellationControlPerPatchOutputComponents uint32
	// MaxTessellationControlTotalOutputComponents is the maximum
	// total number of components of per-vertex and per-patch output
	// variables which can be output from the tessellation control
	// shader stage.
	MaxTessellationControlTotalOutputComponents uint32
	// MaxTessellationEvaluationInputComponents is the maximum number
	// of components of input variables which can be provided as
	// per-vertex inputs to the tessellation evaluation shader stage.
	MaxTessellationEvaluationInputComponents uint32
	// MaxTessellationEvaluationOutputComponents is the maximum number
	// of components of per-vertex output variables which can be
	// output from the tessellation evaluation shader stage.
	MaxTessellationEvaluationOutputComponents uint32
	// MaxGeometryShaderInvocations is the maximum invocation count
	// supported for instanced geometry shaders. The value provided in
	// the Invocations execution mode of shader modules must be less
	// than or equal to this limit.
	MaxGeometryShaderInvocations uint32
	// MaxGeometryInputComponents is the maximum number of components
	// of input variables which can be provided as inputs to the
	// geometry shader stage.
	MaxGeometryInputComponents uint32
	// MaxGeometryOutputComponents is the maximum number of components
	// of output variables which can be output from the geometry
	// shader stage.
	MaxGeometryOutputComponents uint32
	// MaxGeometryOutputVertices is the maximum number of vertices
	// which can be emitted by any geometry shader.
	MaxGeometryOutputVertices uint32
	// MaxGeometryTotalOutputComponents is the maximum total number of
	// components of output, across all emitted vertices, which can be
	// output from the geometry shader stage.
	MaxGeometryTotalOutputComponents uint32
	// MaxFragmentInputComponents is the maximum number of components
	// of input variables which can be provided as inputs to the
	// fragment shader stage.
	MaxFragmentInputComponents uint32
	// MaxFragmentOutputAttachments is the maximum number of output
	// attachments which can be written to by the fragment shader
	// stage.
	MaxFragmentOutputAttachments uint32
	// MaxFragmentDualSrcAttachments is the maximum number of output
	// attachments which can be written to by the fragment shader
	// stage when blending is enabled and one of the dual source blend
	// modes is in use.
	MaxFragmentDualSrcAttachments uint32
	// MaxFragmentCombinedOutputResources is the total number of
	// storage buffers, storage images, and output buffers which can
	// be used in the fragment shader stage.
	MaxFragmentCombinedOutputResources uint32
	// MaxComputeSharedMemorySize is the maximum total storage size,
	// in bytes, of all variables declared with the WorkgroupLocal
	// storage class in shader modules (or with the shared storage
	// qualifier in GLSL) in the compute shader stage.
	MaxComputeSharedMemorySize uint32
	// MaxComputeWorkGroupCount is the maximum number of local
	// workgroups that can be dispatched by a single dispatch command.
	// These three values represent the maximum number of local
	// workgroups for the X, Y, and Z dimensions, respectively. The
	// workgroup count parameters to the dispatch commands must be
	// less than or equal to the corresponding limit.
	MaxComputeWorkGroupCount [3]uint32
	// MaxComputeWorkGroupInvocations is the maximum total number of
	// compute shader invocations in a single local workgroup. The
	// product of the X, Y, and Z sizes as specified by the LocalSize
	// execution mode in shader modules and by the object decorated by
	// the WorkgroupSize decoration must be less than or equal to this
	// limit.
	MaxComputeWorkGroupInvocations uint32
	// MaxComputeWorkGroupSize is the maximum size of a local compute
	// workgroup, per dimension. These three values represent the
	// maximum local workgroup size in the X, Y, and Z dimensions,
	// respectively. The x, y, and z sizes specified by the LocalSize
	// execution mode and by the object decorated by the WorkgroupSize
	// decoration in shader modules must be less than or equal to the
	// corresponding limit.
	MaxComputeWorkGroupSize [3]uint32
	// SubPixelPrecisionBits is the number of bits of subpixel
	// precision in framebuffer coordinates xf and yf.
	SubPixelPrecisionBits uint32
	// SubTexelPrecisionBits is the number of bits of precision in the
	// division along an axis of an image used for minification and
	// magnification filters. 2subTexelPrecisionBits is the actual
	// number of divisions along each axis of the image represented.
	// Sub-texel values calculated during image sampling will snap to
	// these locations when generating the filtered results.
	SubTexelPrecisionBits uint32
	// MipmapPrecisionBits is the number of bits of division that the
	// LOD calculation for mipmap fetching get snapped to when
	// determining the contribution from each mip level to the mip
	// filtered results. 2mipmapPrecisionBits is the actual number of
	// divisions.
	MipmapPrecisionBits uint32
	// MaxDrawIndexedIndexValue is the maximum index value that can be
	// used for indexed draw calls when using 32-bit indices. This
	// excludes the primitive restart index value of 0xFFFFFFFF.
	MaxDrawIndexedIndexValue uint32
	// MaxDrawIndirectCount is the maximum draw count that is
	// supported for indirect draw calls.
	MaxDrawIndirectCount uint32
	// MaxSamplerLodBias is the maximum absolute sampler LOD bias. The
	// sum of the MipLodBias field of the SamplerCreateInfo structure
	// and the Bias operand of image sampling operations in shader
	// modules (or 0 if no Bias operand is provided to an image
	// sampling operation) are clamped to the range
	// [-maxSamplerLodBias,+maxSamplerLodBias].
	MaxSamplerLodBias float32
	// MaxSamplerAnisotropy is the maximum degree of sampler
	// anisotropy. The maximum degree of anisotropic filtering used
	// for an image sampling operation is the minimum of the
	// MaxAnisotropy field of the SamplerCreateInfo structure and this
	// limit.
	MaxSamplerAnisotropy float32
	// MaxViewports is the maximum number of active viewports. The
	// ViewportCount field of the PipelineViewportStateCreateInfo
	// structure that is provided at pipeline creation must be less
	// than or equal to this limit.
	MaxViewports uint32
	// MaxViewportDimensions are the maximum viewport dimensions in
	// the X (width) and Y (height) dimensions, respectively. The
	// maximum viewport dimensions must be greater than or equal to
	// the largest image which can be created and used as a
	// framebuffer attachment.
	MaxViewportDimensions [2]uint32
	// ViewportBoundsRange is the [minimum, maximum] range that the
	// corners of a viewport must be contained in. This range must be
	// at least [-2 × size, 2 × size - 1], where size =
	// max(maxViewportDimensions[0], maxViewportDimensions[1]).
	ViewportBoundsRange [2]float32
	// ViewportSubPixelBits is the number of bits of subpixel
	// precision for viewport bounds. The subpixel precision that
	// floating-point viewport bounds are interpreted at is given by
	// this limit.
	ViewportSubPixelBits uint32
	// MinMemoryMapAlignment is the minimum required alignment, in
	// bytes, of host visible memory allocations within the host
	// address space. When mapping a memory allocation with
	// MapMemory, subtracting offset bytes from the returned pointer
	// will always produce an integer multiple of this limit.
	MinMemoryMapAlignment uintptr
	// MinTexelBufferOffsetAlignment is the minimum required
	// alignment, in bytes, for the Offset field of the
	// BufferViewCreateInfo structure for texel buffers. When a buffer
	// view is created for a buffer which was created with
	// BufferUsageUniformTexelBufferBit or
	// BufferUsageStorageTexelBufferBit set in the Usage field of the
	// BufferCreateInfo structure, the offset must be an integer
	// multiple of this limit.
	MinTexelBufferOffsetAlignment DeviceSize
	// MinUniformBufferOffsetAlignment is the minimum required
	// alignment, in bytes, for the Offset field of the
	// DescriptorBufferInfo structure for uniform buffers. When a
	// descriptor of type DescriptorTypeUniformBuffer or
	// DescriptorTypeUniformBufferDynamic is updated, the offset must
	// be an integer multiple of this limit. Similarly, dynamic
	// offsets for uniform buffers must be multiples of this limit.
	MinUniformBufferOffsetAlignment DeviceSize
	// MinStorageBufferOffsetAlignment is the minimum required
	// alignment, in bytes, for the Offset field of the
	// DescriptorBufferInfo structure for storage buffers. When a
	// descriptor of type DescriptorTypeStorageBuffer or
	// DescriptorTypeStorageBufferDynamic is updated, the offset must
	// be an integer multiple of this limit. Similarly, dynamic
	// offsets for storage buffers must be multiples of this limit.
	MinStorageBufferOffsetAlignment DeviceSize
	// MinTexelOffset is the minimum offset value for the ConstOffset
	// image operand of any of the OpImageSample* or OpImageFetch*
	// image instructions.
	MinTexelOffset int32
	// MaxTexelOffset is the maximum offset value for the ConstOffset
	// image operand of any of the OpImageSample* or OpImageFetch*
	// image instructions.
	MaxTexelOffset uint32
	// MinTexelGatherOffset is the minimum offset value for the Offset
	// or ConstOffsets image operands of any of the OpImage*Gather
	// image instructions.
	MinTexelGatherOffset int32
	// MaxTexelGatherOffset is the maximum offset value for the Offset
	// or ConstOffsets image operands of any of the OpImage*Gather
	// image instructions.
	MaxTexelGatherOffset uint32
	// MinInterpolationOffset is the minimum negative offset value for
	// the offset operand of the InterpolateAtOffset extended
	// instruction.
	MinInterpolationOffset float32
	// MaxInterpolationOffset is the maximum positive offset value for
	// the offset operand of the InterpolateAtOffset extended
	// instruction.
	MaxInterpolationOffset float32
	// SubPixelInterpolationOffsetBits is the number of subpixel
	// fractional bits that the x and y offsets to the
	// InterpolateAtOffset extended instruction may be rounded to as
	// fixed-point values.
	SubPixelInterpolationOffsetBits uint32
	// MaxFramebufferWidth is the maximum width for a framebuffer. The
	// Width field of the FramebufferCreateInfo structure must be less
	// than or equal to this limit.
	MaxFramebufferWidth uint32
	// MaxFramebufferHeight is the maximum height for a framebuffer.
	// The Height field of the FramebufferCreateInfo structure must be
	// less than or equal to this limit.
	MaxFramebufferHeight uint32
	// MaxFramebufferLayers is the maximum layer count for a layered
	// framebuffer. The Layers field of the FramebufferCreateInfo
	// structure must be less than or equal to this limit.
	MaxFramebufferLayers uint32
	// FramebufferColorSampleCounts is a bitmask1 of
	// SampleCountFlagBits indicating the color sample counts that are
	// supported for all framebuffer color attachments with floating-
	// or fixed-point formats. There is no limit that specifies the
	// color sample counts that are supported for all color
	// attachments with integer formats.
	FramebufferColorSampleCounts SampleCountFlags
	// FramebufferDepthSampleCounts is a bitmask1 of
	// SampleCountFlagBits indicating the supported depth sample
	// counts for all framebuffer depth/stencil attachments, when the
	// format includes a depth component.
	FramebufferDepthSampleCounts SampleCountFlags
	// FramebufferStencilSampleCounts is a bitmask1 of
	// SampleCountFlagBits indicating the supported stencil sample
	// counts for all framebuffer depth/stencil attachments, when the
	// format includes a stencil component.
	FramebufferStencilSampleCounts SampleCountFlags
	// FramebufferNoAttachmentsSampleCounts is a bitmask1 of
	// SampleCountFlagBits indicating the supported sample counts for
	// a framebuffer with no attachments.
	FramebufferNoAttachmentsSampleCounts SampleCountFlags
	// MaxColorAttachments is the maximum number of color attachments
	// that can be used by a subpass in a render pass. The
	// ColorAttachmentCount field of the SubpassDescription structure
	// must be less than or equal to this limit.
	MaxColorAttachments uint32
	// SampledImageColorSampleCounts is a bitmask1 of
	// SampleCountFlagBits indicating the sample counts supported for
	// all 2D images created with ImageTilingOptimal, usage containing
	// ImageUsageSampledBit, and a non-integer color format.
	SampledImageColorSampleCounts SampleCountFlags
	// SampledImageIntegerSampleCounts is a bitmask1 of
	// SampleCountFlagBits indicating the sample counts supported for
	// all 2D images created with ImageTilingOptimal, usage containing
	// ImageUsageSampledBit, and an integer color format.
	SampledImageIntegerSampleCounts SampleCountFlags
	// SampledImageDepthSampleCounts is a bitmask1 of
	// SampleCountFlagBits indicating the sample counts supported for
	// all 2D images created with ImageTilingOptimal, usage containing
	// ImageUsageSampledBit, and a depth format.
	SampledImageDepthSampleCounts SampleCountFlags
	// SampledImageStencilSampleCounts is a bitmask1 of
	// SampleCountFlagBits indicating the sample supported for all 2D
	// images created with ImageTilingOptimal, usage containing
	// ImageUsageSampledBit, and a stencil format.
	SampledImageStencilSampleCounts SampleCountFlags
	// StorageImageSampleCounts is a bitmask1 of SampleCountFlagBits
	// indicating the sample counts supported for all 2D images
	// created with ImageTilingOptimal, and usage containing
	// ImageUsageStorageBit.
	StorageImageSampleCounts SampleCountFlags
	// MaxSampleMaskWords is the maximum number of array elements of a
	// variable decorated with the SampleMask built-in decoration.
	MaxSampleMaskWords uint32
	// TimestampComputeAndGraphics specifies support for timestamps on
	// all graphics and compute queues. If this limit is set to true,
	// all queues that advertise the QueueGraphicsBit or
	// QueueComputeBit in the QueueFamilyProperties::queueFlags
	// support QueueFamilyProperties::timestampValidBits of at least
	// 36.
	TimestampComputeAndGraphics bool
	// TimestampPeriod is the number of nanoseconds required for a
	// timestamp query to be incremented by 1.
	TimestampPeriod float32
	// MaxClipDistances is the maximum number of clip distances that
	// can be used in a single shader stage. The size of any array
	// declared with the ClipDistance built-in decoration in a shader
	// module must be less than or equal to this limit.
	MaxClipDistances uint32
	// MaxCullDistances is the maximum number of cull distances that
	// can be used in a single shader stage. The size of any array
	// declared with the CullDistance built-in decoration in a shader
	// module must be less than or equal to this limit.
	MaxCullDistances uint32
	// MaxCombinedClipAndCullDistances is the maximum combined number
	// of clip and cull distances that can be used in a single shader
	// stage. The sum of the sizes of any pair of arrays declared with
	// the ClipDistance and CullDistance built-in decoration used by a
	// single shader stage in a shader module must be less than or
	// equal to this limit.
	MaxCombinedClipAndCullDistances uint32
	// DiscreteQueuePriorities is the number of discrete priorities
	// that can be assigned to a queue based on the value of each
	// field of DeviceQueueCreateInfo.QueuePriorities. This must be at
	// least 2, and levels must be spread evenly over the range, with
	// at least one level at 1.0, and another at 0.0.
	DiscreteQueuePriorities uint32
	// PointSizeRange is the range [minimum,maximum] of supported
	// sizes for points. Values written to variables decorated with
	// the PointSize built-in decoration are clamped to this range.
	PointSizeRange [2]float32
	// LineWidthRange is the range [minimum,maximum] of supported
	// widths for lines. Values specified by the LineWidth field of
	// the PipelineRasterizationStateCreateInfo or the lineWidth
	// parameter to SetLineWidth are clamped to this range.
	LineWidthRange [2]float32
	// PointSizeGranularity is the granularity of supported point
	// sizes. Not all point sizes in the range defined by
	// pointSizeRange are supported. This limit specifies the
	// granularity (or increment) between successive supported point
	// sizes.
	PointSizeGranularity float32
	// LineWidthGranularity is the granularity of supported line
	// widths. Not all line widths in the range defined by
	// lineWidthRange are supported. This limit specifies the
	// granularity (or increment) between successive supported line
	// widths.
	LineWidthGranularity float32
	// StrictLines specifies whether lines are rasterized according to
	// the preferred method of rasterization. If set to false, lines
	// may be rasterized under a relaxed set of rules. If set to true,
	// lines are rasterized as per the strict definition.
	StrictLines bool
	// StandardSampleLocations specifies whether rasterization uses
	// the standard sample locations as documented in Multisampling.
	// If set to true, the implementation uses the documented sample
	// locations. If set to false, the implementation may use
	// different sample locations.
	StandardSampleLocations bool
	// OptimalBufferCopyOffsetAlignment is the optimal buffer offset
	// alignment in bytes for CopyBufferToImage and
	// CopyImageToBuffer. The per texel alignment requirements
	// are enforced, but applications should use the optimal alignment
	// for optimal performance and power use.
	OptimalBufferCopyOffsetAlignment DeviceSize
	// OptimalBufferCopyRowPitchAlignment is the optimal buffer row
	// pitch alignment in bytes for CopyBufferToImage and
	// CopyImageToBuffer. Row pitch is the number of bytes
	// between texels with the same X coordinate in adjacent rows (Y
	// coordinates differ by one). The per texel alignment
	// requirements are enforced, but applications should use the
	// optimal alignment for optimal performance and power use.
	OptimalBufferCopyRowPitchAlignment DeviceSize
	// NonCoherentAtomSize is the size and alignment in bytes that
	// bounds concurrent access to host-mapped device memory.
	NonCoherentAtomSize DeviceSize
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
		PipelineCacheUUID: (*[C.VK_UUID_SIZE]byte)(unsafe.Pointer(&cprops.pipelineCacheUUID))[:],
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
	a := new(allocator)
	defer a.free()

	var cprops C.VkPhysicalDeviceProperties2
	cprops.sType = C.VkStructureType(StructureTypePhysicalDeviceProperties2)
	cprops.pNext = buildChain(a, extensions)
	defer internalizeChain(extensions, cprops.pNext)
	C.domVkGetPhysicalDeviceProperties2(dev.instance.fps[vkGetPhysicalDeviceProperties2], dev.hnd, &cprops)
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
	// must be kept identical to C struct
	_ structs.HostLayout

	PropertyFlags MemoryPropertyFlags
	HeapIndex     uint32
}

type MemoryHeap struct {
	// must be kept identical to C struct
	_ structs.HostLayout

	Size  DeviceSize
	Flags MemoryHeapFlags
}

type PhysicalDeviceMemoryProperties struct {
	Types []MemoryType
	Heaps []MemoryHeap
}

func (dev *PhysicalDevice) MemoryProperties() PhysicalDeviceMemoryProperties {
	var props C.VkPhysicalDeviceMemoryProperties
	C.domVkGetPhysicalDeviceMemoryProperties(dev.instance.fps[vkGetPhysicalDeviceMemoryProperties], dev.hnd, &props)

	return PhysicalDeviceMemoryProperties{
		Types: (*[C.VK_MAX_MEMORY_TYPES]MemoryType)(unsafe.Pointer(&props.memoryTypes))[:props.memoryTypeCount],
		Heaps: (*[C.VK_MAX_MEMORY_TYPES]MemoryHeap)(unsafe.Pointer(&props.memoryHeaps))[:props.memoryHeapCount],
	}
}

type PhysicalDeviceMemoryProperties2 struct {
	Extensions       []Extension
	MemoryProperties PhysicalDeviceMemoryProperties
}

func (dev *PhysicalDevice) MemoryProperties2(props PhysicalDeviceMemoryProperties2) {
	a := new(allocator)
	defer a.free()

	cprops := C.VkPhysicalDeviceMemoryProperties2{
		sType: C.VkStructureType(StructureTypePhysicalDeviceMemoryProperties2),
		pNext: buildChain(a, props.Extensions),
	}
	defer internalizeChain(props.Extensions, cprops.pNext)
	C.domVkGetPhysicalDeviceMemoryProperties2(dev.instance.fps[vkGetPhysicalDeviceMemoryProperties2], dev.hnd, &cprops)
	props.MemoryProperties = PhysicalDeviceMemoryProperties{
		Types: (*[C.VK_MAX_MEMORY_TYPES]MemoryType)(unsafe.Pointer(&cprops.memoryProperties.memoryTypes))[:cprops.memoryProperties.memoryTypeCount],
		Heaps: (*[C.VK_MAX_MEMORY_TYPES]MemoryHeap)(unsafe.Pointer(&cprops.memoryProperties.memoryHeaps))[:cprops.memoryProperties.memoryHeapCount],
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
		defer C.free(unsafe.Pointer(cLayer))
	}
	res := Result(C.domVkEnumerateDeviceExtensionProperties(dev.instance.fps[vkEnumerateDeviceExtensionProperties], dev.hnd, cLayer, &count, nil))
	if res != Success {
		return nil, res
	}
	properties := make([]C.VkExtensionProperties, count)
	res = Result(C.domVkEnumerateDeviceExtensionProperties(
		dev.instance.fps[vkEnumerateDeviceExtensionProperties],
		dev.hnd,
		cLayer,
		&count,
		unsafe.SliceData(properties)))
	if res != Success {
		return nil, res
	}
	out := make([]ExtensionProperties, count)

	for i, s := range properties {
		name := (*[256]byte)(unsafe.Pointer(&s.extensionName))[:]
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

type PhysicalDeviceFeatures2 struct {
	Extensions []Extension
	Features   PhysicalDeviceFeatures
}

func (dev *PhysicalDevice) Features2(features *PhysicalDeviceFeatures2) {
	var cfeatures C.VkPhysicalDeviceFeatures2
	C.domVkGetPhysicalDeviceFeatures2(dev.instance.fps[vkGetPhysicalDeviceFeatures2], dev.hnd, &cfeatures)
	features.Features = PhysicalDeviceFeatures{
		RobustBufferAccess:                      cfeatures.features.robustBufferAccess == C.VK_TRUE,
		FullDrawIndexUint32:                     cfeatures.features.fullDrawIndexUint32 == C.VK_TRUE,
		ImageCubeArray:                          cfeatures.features.imageCubeArray == C.VK_TRUE,
		IndependentBlend:                        cfeatures.features.independentBlend == C.VK_TRUE,
		GeometryShader:                          cfeatures.features.geometryShader == C.VK_TRUE,
		TessellationShader:                      cfeatures.features.tessellationShader == C.VK_TRUE,
		SampleRateShading:                       cfeatures.features.sampleRateShading == C.VK_TRUE,
		DualSrcBlend:                            cfeatures.features.dualSrcBlend == C.VK_TRUE,
		LogicOp:                                 cfeatures.features.logicOp == C.VK_TRUE,
		MultiDrawIndirect:                       cfeatures.features.multiDrawIndirect == C.VK_TRUE,
		DrawIndirectFirstInstance:               cfeatures.features.drawIndirectFirstInstance == C.VK_TRUE,
		DepthClamp:                              cfeatures.features.depthClamp == C.VK_TRUE,
		DepthBiasClamp:                          cfeatures.features.depthBiasClamp == C.VK_TRUE,
		FillModeNonSolid:                        cfeatures.features.fillModeNonSolid == C.VK_TRUE,
		DepthBounds:                             cfeatures.features.depthBounds == C.VK_TRUE,
		WideLines:                               cfeatures.features.wideLines == C.VK_TRUE,
		LargePoints:                             cfeatures.features.largePoints == C.VK_TRUE,
		AlphaToOne:                              cfeatures.features.alphaToOne == C.VK_TRUE,
		MultiViewport:                           cfeatures.features.multiViewport == C.VK_TRUE,
		SamplerAnisotropy:                       cfeatures.features.samplerAnisotropy == C.VK_TRUE,
		TextureCompressionETC2:                  cfeatures.features.textureCompressionETC2 == C.VK_TRUE,
		TextureCompressionASTC_LDR:              cfeatures.features.textureCompressionASTC_LDR == C.VK_TRUE,
		TextureCompressionBC:                    cfeatures.features.textureCompressionBC == C.VK_TRUE,
		OcclusionQueryPrecise:                   cfeatures.features.occlusionQueryPrecise == C.VK_TRUE,
		PipelineStatisticsQuery:                 cfeatures.features.pipelineStatisticsQuery == C.VK_TRUE,
		VertexPipelineStoresAndAtomics:          cfeatures.features.vertexPipelineStoresAndAtomics == C.VK_TRUE,
		FragmentStoresAndAtomics:                cfeatures.features.fragmentStoresAndAtomics == C.VK_TRUE,
		ShaderTessellationAndGeometryPointSize:  cfeatures.features.shaderTessellationAndGeometryPointSize == C.VK_TRUE,
		ShaderImageGatherExtended:               cfeatures.features.shaderImageGatherExtended == C.VK_TRUE,
		ShaderStorageImageExtendedFormats:       cfeatures.features.shaderStorageImageExtendedFormats == C.VK_TRUE,
		ShaderStorageImageMultisample:           cfeatures.features.shaderStorageImageMultisample == C.VK_TRUE,
		ShaderStorageImageReadWithoutFormat:     cfeatures.features.shaderStorageImageReadWithoutFormat == C.VK_TRUE,
		ShaderStorageImageWriteWithoutFormat:    cfeatures.features.shaderStorageImageWriteWithoutFormat == C.VK_TRUE,
		ShaderUniformBufferArrayDynamicIndexing: cfeatures.features.shaderUniformBufferArrayDynamicIndexing == C.VK_TRUE,
		ShaderSampledImageArrayDynamicIndexing:  cfeatures.features.shaderSampledImageArrayDynamicIndexing == C.VK_TRUE,
		ShaderStorageBufferArrayDynamicIndexing: cfeatures.features.shaderStorageBufferArrayDynamicIndexing == C.VK_TRUE,
		ShaderStorageImageArrayDynamicIndexing:  cfeatures.features.shaderStorageImageArrayDynamicIndexing == C.VK_TRUE,
		ShaderClipDistance:                      cfeatures.features.shaderClipDistance == C.VK_TRUE,
		ShaderCullDistance:                      cfeatures.features.shaderCullDistance == C.VK_TRUE,
		ShaderFloat64:                           cfeatures.features.shaderFloat64 == C.VK_TRUE,
		ShaderInt64:                             cfeatures.features.shaderInt64 == C.VK_TRUE,
		ShaderInt16:                             cfeatures.features.shaderInt16 == C.VK_TRUE,
		ShaderResourceResidency:                 cfeatures.features.shaderResourceResidency == C.VK_TRUE,
		ShaderResourceMinLod:                    cfeatures.features.shaderResourceMinLod == C.VK_TRUE,
		SparseBinding:                           cfeatures.features.sparseBinding == C.VK_TRUE,
		SparseResidencyBuffer:                   cfeatures.features.sparseResidencyBuffer == C.VK_TRUE,
		SparseResidencyImage2D:                  cfeatures.features.sparseResidencyImage2D == C.VK_TRUE,
		SparseResidencyImage3D:                  cfeatures.features.sparseResidencyImage3D == C.VK_TRUE,
		SparseResidency2Samples:                 cfeatures.features.sparseResidency2Samples == C.VK_TRUE,
		SparseResidency4Samples:                 cfeatures.features.sparseResidency4Samples == C.VK_TRUE,
		SparseResidency8Samples:                 cfeatures.features.sparseResidency8Samples == C.VK_TRUE,
		SparseResidency16Samples:                cfeatures.features.sparseResidency16Samples == C.VK_TRUE,
		SparseResidencyAliased:                  cfeatures.features.sparseResidencyAliased == C.VK_TRUE,
		VariableMultisampleRate:                 cfeatures.features.variableMultisampleRate == C.VK_TRUE,
		InheritedQueries:                        cfeatures.features.inheritedQueries == C.VK_TRUE,
	}
	internalizeChain(features.Extensions, cfeatures.pNext)
}

type QueueFamilyProperties struct {
	// must be kept identical to C struct
	_ structs.HostLayout

	QueueFlags                  QueueFlags
	QueueCount                  uint32
	TimestampValidBits          uint32
	MinImageTransferGranularity Extent3D
}

type Extent2D struct {
	// must be kept identical to C struct
	_ structs.HostLayout

	Width  uint32
	Height uint32
}

type Extent3D struct {
	// must be kept identical to C struct
	_ structs.HostLayout

	Width  uint32
	Height uint32
	Depth  uint32
}

func (dev *PhysicalDevice) QueueFamilyProperties() []QueueFamilyProperties {
	var count C.uint32_t
	C.domVkGetPhysicalDeviceQueueFamilyProperties(
		dev.instance.fps[vkGetPhysicalDeviceQueueFamilyProperties],
		dev.hnd,
		&count,
		nil)
	props := make([]QueueFamilyProperties, count)
	C.domVkGetPhysicalDeviceQueueFamilyProperties(
		dev.instance.fps[vkGetPhysicalDeviceQueueFamilyProperties],
		dev.hnd,
		&count,
		safeish.SliceCastPtr[*C.VkQueueFamilyProperties](props))
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

// Device is an opaque handle to a device object.
type Device struct {
	// VK_DEFINE_HANDLE(VkDevice)
	hnd C.VkDevice

	fps                 [deviceMaxPFN]C.PFN_vkVoidFunction
	vkGetDeviceProcAddr C.PFN_vkGetDeviceProcAddr
}

func (dev *PhysicalDevice) CreateDevice(info *DeviceCreateInfo) (*Device, error) {
	a := new(allocator)
	defer a.free()

	var ptr C.VkDeviceCreateInfo
	ptr.sType = C.VkStructureType(StructureTypeDeviceCreateInfo)
	ptr.pNext = buildChain(a, info.Extensions)
	defer internalizeChain(info.Extensions, ptr.pNext)
	ptr.queueCreateInfoCount = C.uint32_t(len(info.QueueCreateInfos))
	ptr.pQueueCreateInfos = allocn[C.VkDeviceQueueCreateInfo](a, len(info.QueueCreateInfos))
	arr := (*[math.MaxInt32]C.VkDeviceQueueCreateInfo)(unsafe.Pointer(ptr.pQueueCreateInfos))[:len(info.QueueCreateInfos)]
	for i, obj := range info.QueueCreateInfos {
		arr[i] = C.VkDeviceQueueCreateInfo{
			sType:            C.VkStructureType(StructureTypeDeviceQueueCreateInfo),
			pNext:            buildChain(a, obj.Extensions),
			flags:            C.VkDeviceQueueCreateFlags(obj.Flags),
			queueFamilyIndex: C.uint32_t(obj.QueueFamilyIndex),
			queueCount:       C.uint32_t(len(obj.QueuePriorities)),
			pQueuePriorities: externFloat32(a, obj.QueuePriorities),
		}
		defer internalizeChain(obj.Extensions, arr[i].pNext)
	}
	ptr.enabledExtensionCount = C.uint32_t(len(info.EnabledExtensionNames))
	ptr.ppEnabledExtensionNames = externStrings(a, info.EnabledExtensionNames)
	if info.EnabledFeatures != nil {
		ptr.pEnabledFeatures = alloc[C.VkPhysicalDeviceFeatures](a)
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
	}
	var out C.VkDevice
	res := Result(C.domVkCreateDevice(dev.instance.fps[vkCreateDevice], dev.hnd, &ptr, nil, &out))
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
	defer C.free(unsafe.Pointer(cName))
	fp := C.domVkGetDeviceProcAddr(dev.vkGetDeviceProcAddr, dev.hnd, cName)
	if debug {
		fmt.Fprintf(os.Stderr, "%s = %p\n", name, fp)
	}
	return fp
}

// Creating a logical device also creates the queues associated with that device.
// The queues to create are described by a set of DeviceQueueCreateInfo structures
// that are passed to CreateDevice.
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

type DeviceQueueInfo2 struct {
	Extensions       []Extension
	Flags            DeviceQueueCreateFlags
	QueueFamilyIndex uint32
	QueueIndex       uint32
}

func (info *DeviceQueueInfo2) c(a *allocator) *C.VkDeviceQueueInfo2 {
	cinfo := alloc[C.VkDeviceQueueInfo2](a)
	*cinfo = C.VkDeviceQueueInfo2{
		sType:            C.VkStructureType(StructureTypeDeviceQueueInfo2),
		pNext:            buildChain(a, info.Extensions),
		flags:            C.VkDeviceQueueCreateFlags(info.Flags),
		queueFamilyIndex: C.uint32_t(info.QueueFamilyIndex),
		queueIndex:       C.uint32_t(info.QueueIndex),
	}
	return cinfo
}

func (dev *Device) Queue2(info *DeviceQueueInfo2) *Queue {
	a := new(allocator)
	defer a.free()

	cinfo := info.c(a)
	defer internalizeChain(info.Extensions, cinfo.pNext)
	out := &Queue{fps: &dev.fps}
	C.domVkGetDeviceQueue2(dev.fps[vkGetDeviceQueue2], dev.hnd, cinfo, &out.hnd)
	return out
}

// Command buffers are objects used to record commands which can be subsequently submitted to a device queue for execution.
// There are two levels of command buffers - primary command buffers, which can execute secondary command buffers, and which are submitted to queues,
// and secondary command buffers, which can be executed by primary command buffers, and which are not directly submitted to queues.
type CommandBuffer struct {
	// VK_DEFINE_HANDLE(VkCommandBuffer)
	hnd C.VkCommandBuffer
	fps *[deviceMaxPFN]C.PFN_vkVoidFunction

	bufs []C.VkCommandBuffer
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

// A render pass represents a collection of attachments, subpasses, and dependencies between the subpasses,
// and describes how the attachments are used over the course of the subpasses.
// The use of a render pass in a command buffer is a render pass instance.
type RenderPass struct {
	// VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkRenderPass)
	hnd C.VkRenderPass
}

// Render passes operate in conjunction with framebuffers.
// Framebuffers represent a collection of specific memory attachments that a render pass instance uses.
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

// Begin starts recording a command buffer.
func (buf *CommandBuffer) Begin(info *CommandBufferBeginInfo) error {
	a := new(allocator)
	defer a.free()

	var ptr C.VkCommandBufferBeginInfo
	ptr.sType = C.VkStructureType(StructureTypeCommandBufferBeginInfo)
	ptr.pNext = buildChain(a, info.Extensions)
	defer internalizeChain(info.Extensions, ptr.pNext)
	ptr.flags = C.VkCommandBufferUsageFlags(info.Flags)
	if info.InheritanceInfo != nil {
		ptr.pInheritanceInfo = alloc[C.VkCommandBufferInheritanceInfo](a)
		ptr.pInheritanceInfo.sType = C.VkStructureType(StructureTypeCommandBufferInheritanceInfo)
		ptr.pInheritanceInfo.pNext = buildChain(a, info.InheritanceInfo.Extensions)
		defer internalizeChain(info.InheritanceInfo.Extensions, ptr.pInheritanceInfo.pNext)
		ptr.pInheritanceInfo.renderPass = C.VkRenderPass(info.InheritanceInfo.RenderPass.hnd)
		ptr.pInheritanceInfo.subpass = C.uint32_t(info.InheritanceInfo.Subpass)
		ptr.pInheritanceInfo.framebuffer = C.VkFramebuffer(info.InheritanceInfo.Framebuffer.hnd)
		ptr.pInheritanceInfo.occlusionQueryEnable = vkBool(info.InheritanceInfo.OcclusionQueryEnable)
		ptr.pInheritanceInfo.queryFlags = C.VkQueryControlFlags(info.InheritanceInfo.QueryFlags)
		ptr.pInheritanceInfo.pipelineStatistics = C.VkQueryPipelineStatisticFlags(info.InheritanceInfo.PipelineStatistics)
	}
	res := Result(C.domVkBeginCommandBuffer(buf.fps[vkBeginCommandBuffer], buf.hnd, &ptr))
	return result2error(res)
}

// End finishes recording a command buffer.
//
// If there was an error during recording,
// the application will be notified by an unsuccessful return code returned by End.
// If the application wishes to further use the command buffer, the command buffer must be reset.
// The command buffer must have been in the recording state, and is moved to the executable state.
func (buf *CommandBuffer) End() error {
	res := Result(C.domVkEndCommandBuffer(buf.fps[vkEndCommandBuffer], buf.hnd))
	return result2error(res)
}

// SetLineWidth sets the dynamic line width state.
func (buf *CommandBuffer) SetLineWidth(lineWidth float32) {
	C.domVkCmdSetLineWidth(buf.fps[vkCmdSetLineWidth], buf.hnd, C.float(lineWidth))
}

// SetDepthBias sets the depth bias dynamic state.
//
// The depth values of all fragments generated by the rasterization of a polygon
// can be offset by a single value that is computed for that polygon.
// This behavior is controlled by the
// DepthBiasEnable, DepthBiasConstantFactor, DepthBiasClamp, and DepthBiasSlopeFactor fields of PipelineRasterizationStateCreateInfo,
// or by the corresponding parameters to the SetDepthBias command if depth bias state is dynamic.
func (buf *CommandBuffer) SetDepthBias(constantFactor, clamp, slopeFactor float32) {
	C.domVkCmdSetDepthBias(buf.fps[vkCmdSetDepthBias], buf.hnd, C.float(constantFactor), C.float(clamp), C.float(slopeFactor))
}

// SetBlendConstants sets the values of blend constants.
func (buf *CommandBuffer) SetBlendConstants(blendConstants [4]float32) {
	C.domVkCmdSetBlendConstants(buf.fps[vkCmdSetBlendConstants], buf.hnd, (*C.float)(&blendConstants[0]))
}

// Draw draws primitives.
//
// When the command is executed, primitives are assembled
// using the current primitive topology and vertexCount consecutive vertex indices with the first vertexIndex value equal to firstVertex.
// The primitives are drawn instanceCount times with instanceIndex starting with firstInstance and increasing sequentially for each instance.
// The assembled primitives execute the bound graphics pipeline.
func (buf *CommandBuffer) Draw(vertexCount, instanceCount, firstVertex, firstInstance uint32) {
	C.domVkCmdDraw(
		buf.fps[vkCmdDraw],
		buf.hnd,
		C.uint32_t(vertexCount),
		C.uint32_t(instanceCount),
		C.uint32_t(firstVertex),
		C.uint32_t(firstInstance))
}

// SetViewport sets the viewport on a command buffer.
func (buf *CommandBuffer) SetViewport(firstViewport uint32, viewports []Viewport) {
	C.domVkCmdSetViewport(
		buf.fps[vkCmdSetViewport],
		buf.hnd,
		C.uint32_t(firstViewport),
		C.uint32_t(len(viewports)),
		safeish.SliceCastPtr[*C.VkViewport](viewports))
}

func (buf *CommandBuffer) SetScissor(firstScissor uint32, scissors []Rect2D) {
	C.domVkCmdSetScissor(
		buf.fps[vkCmdSetScissor],
		buf.hnd,
		C.uint32_t(firstScissor),
		C.uint32_t(len(scissors)),
		safeish.SliceCastPtr[*C.VkRect2D](scissors))
}

func (buf *CommandBuffer) SetDeviceMask(deviceMask uint32) {
	C.domVkCmdSetDeviceMask(buf.fps[vkCmdSetDeviceMask], buf.hnd, C.uint32_t(deviceMask))
}

func (buf *CommandBuffer) SetDepthBounds(min, max float32) {
	C.domVkCmdSetDepthBounds(buf.fps[vkCmdSetDepthBounds], buf.hnd, C.float(min), C.float(max))
}

func (buf *CommandBuffer) PushConstants(layout PipelineLayout, stageFlags ShaderStageFlags, offset uint32, size uint32, data []byte) {
	C.domVkCmdPushConstants(
		buf.fps[vkCmdPushConstants],
		buf.hnd,
		layout.hnd,
		C.VkShaderStageFlags(stageFlags),
		C.uint32_t(offset),
		C.uint32_t(len(data)),
		unsafe.Pointer(unsafe.SliceData(data)))
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
	// must be kept identical to C struct
	_ structs.HostLayout

	Rect           Rect2D
	BaseArrayLayer uint32
	LayerCount     uint32
}

func (buf *CommandBuffer) ClearAttachments(attachments []ClearAttachment, rects []ClearRect) {
	a := new(allocator)
	defer a.free()

	mem := allocn[C.VkClearAttachment](a, len(attachments))
	arr := unsafe.Slice(mem, len(attachments))
	for i := range arr {
		arr[i] = C.VkClearAttachment{
			aspectMask:      C.VkImageAspectFlags(attachments[i].AspectMask),
			colorAttachment: C.uint32_t(attachments[i].ColorAttachment),
		}
		switch v := attachments[i].ClearValue.(type) {
		case ClearColorValueFloat32s:
			copy(arr[i].clearValue[:], (*[16]byte)(unsafe.Pointer(&v))[:])
		case ClearColorValueInt32s:
			copy(arr[i].clearValue[:], (*[16]byte)(unsafe.Pointer(&v))[:])
		case ClearColorValueUint32s:
			copy(arr[i].clearValue[:], (*[16]byte)(unsafe.Pointer(&v))[:])
		case ClearDepthStencilValue:
			ucopy1(unsafe.Pointer(&arr[i].clearValue), unsafe.Pointer(&v), C.sizeof_VkClearDepthStencilValue)
		default:
			panic(fmt.Sprintf("unreachable: %T", v))
		}
	}
	C.domVkCmdClearAttachments(
		buf.fps[vkCmdClearAttachments],
		buf.hnd,
		C.uint32_t(len(attachments)),
		(*C.VkClearAttachment)(mem),
		C.uint32_t(len(rects)),
		safeish.SliceCastPtr[*C.VkClearRect](rects))
}

func (buf *CommandBuffer) ClearColorImage(image Image, imageLayout ImageLayout, color ClearColorValue, ranges []ImageSubresourceRange) {
	var cColor C.VkClearColorValue
	switch v := color.(type) {
	case ClearColorValueFloat32s:
		copy(cColor[:], (*[16]byte)(unsafe.Pointer(&v))[:])
	case ClearColorValueInt32s:
		copy(cColor[:], (*[16]byte)(unsafe.Pointer(&v))[:])
	case ClearColorValueUint32s:
		copy(cColor[:], (*[16]byte)(unsafe.Pointer(&v))[:])
	default:
		panic(fmt.Sprintf("unreachable: %T", v))
	}
	C.domVkCmdClearColorImage(
		buf.fps[vkCmdClearColorImage],
		buf.hnd,
		image.hnd,
		C.VkImageLayout(imageLayout),
		&cColor,
		C.uint32_t(len(ranges)),
		safeish.SliceCastPtr[*C.VkImageSubresourceRange](ranges))
}

func (buf *CommandBuffer) ClearDepthStencilImage(
	image Image,
	imageLayout ImageLayout,
	depthStencil ClearDepthStencilValue,
	ranges []ImageSubresourceRange,
) {
	C.domVkCmdClearDepthStencilImage(
		buf.fps[vkCmdClearDepthStencilImage],
		buf.hnd,
		image.hnd,
		C.VkImageLayout(imageLayout),
		(*C.VkClearDepthStencilValue)(unsafe.Pointer(&depthStencil)),
		C.uint32_t(len(ranges)),
		safeish.SliceCastPtr[*C.VkImageSubresourceRange](ranges))
}

func (info *RenderPassBeginInfo) c(a *allocator) *C.VkRenderPassBeginInfo {
	cinfo := alloc[C.VkRenderPassBeginInfo](a)
	*cinfo = C.VkRenderPassBeginInfo{
		sType:           C.VkStructureType(StructureTypeRenderPassBeginInfo),
		pNext:           buildChain(a, info.Extensions),
		renderPass:      info.RenderPass.hnd,
		framebuffer:     info.Framebuffer.hnd,
		clearValueCount: C.uint32_t(len(info.ClearValues)),
		pClearValues:    allocn[C.VkClearValue](a, len(info.ClearValues)),
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
	a := new(allocator)
	defer a.free()

	cinfo := info.c(a)
	defer internalizeChain(info.Extensions, cinfo.pNext)
	C.domVkCmdBeginRenderPass(buf.fps[vkCmdBeginRenderPass], buf.hnd, cinfo, C.VkSubpassContents(contents))
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
	// must be kept identical to C struct
	_ structs.HostLayout

	SrcOffset DeviceSize
	DstOffset DeviceSize
	Size      DeviceSize
}

func (buf *CommandBuffer) CopyBuffer(srcBuffer, dstBuffer Buffer, regions []BufferCopy) {
	C.domVkCmdCopyBuffer(
		buf.fps[vkCmdCopyBuffer],
		buf.hnd,
		srcBuffer.hnd,
		dstBuffer.hnd,
		C.uint32_t(len(regions)),
		safeish.SliceCastPtr[*C.VkBufferCopy](regions))
}

type BufferImageCopy struct {
	// must be kept identical to C struct
	_ structs.HostLayout

	BufferOfset       DeviceSize
	BufferRowLength   uint32
	BufferImageHeight uint32
	ImageSubresource  ImageSubresourceLayers
	ImageOffset       Offset3D
	ImageExtent       Extent3D
}

type ImageSubresourceLayers struct {
	// must be kept identical to C struct
	_ structs.HostLayout

	AspectMask     ImageAspectFlags
	MipLevel       uint32
	BaseArrayLayer uint32
	LayerCount     uint32
}

func (buf *CommandBuffer) CopyBufferToImage(srcBuffer Buffer, dstImage Image, dstImageLayout ImageLayout, regions []BufferImageCopy) {
	C.domVkCmdCopyBufferToImage(
		buf.fps[vkCmdCopyBufferToImage],
		buf.hnd,
		srcBuffer.hnd,
		dstImage.hnd,
		C.VkImageLayout(dstImageLayout),
		C.uint32_t(len(regions)),
		safeish.SliceCastPtr[*C.VkBufferImageCopy](regions))
}

type ImageCopy struct {
	// must be kept identical to C struct
	_ structs.HostLayout

	SrcSubresource ImageSubresourceLayers
	SrcOffset      Offset3D
	DstSubresource ImageSubresourceLayers
	DstOffset      Offset3D
	Extent         Extent3D
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
		safeish.SliceCastPtr[*C.VkImageCopy](regions))
}

func (buf *CommandBuffer) CopyImageToBuffer(srcImage Image, srcImageLayout ImageLayout, dstBuffer Buffer, regions []BufferImageCopy) {
	C.domVkCmdCopyImageToBuffer(
		buf.fps[vkCmdCopyImageToBuffer],
		buf.hnd,
		srcImage.hnd,
		C.VkImageLayout(srcImageLayout),
		dstBuffer.hnd,
		C.uint32_t(len(regions)),
		safeish.SliceCastPtr[*C.VkBufferImageCopy](regions))
}

func (buf *CommandBuffer) ResetEvent(event Event, stageMask PipelineStageFlags) {
	C.domVkCmdResetEvent(buf.fps[vkCmdResetEvent], buf.hnd, event.hnd, C.VkPipelineStageFlags(stageMask))
}

func (buf *CommandBuffer) ResetQueryPool(queryPool QueryPool, firstQuery, queryCount uint32) {
	C.domVkCmdResetQueryPool(buf.fps[vkCmdResetQueryPool], buf.hnd, queryPool.hnd, C.uint32_t(firstQuery), C.uint32_t(queryCount))
}

func (buf *CommandBuffer) UpdateBuffer(dstBuffer Buffer, dstOffset DeviceSize, data []byte) {
	C.domVkCmdUpdateBuffer(
		buf.fps[vkCmdUpdateBuffer],
		buf.hnd,
		dstBuffer.hnd,
		C.VkDeviceSize(dstOffset),
		C.VkDeviceSize(len(data)),
		unsafe.Pointer(unsafe.SliceData(data)))
}

func (buf *CommandBuffer) BeginQuery(queryPool QueryPool, query uint32, flags QueryControlFlags) {
	C.domVkCmdBeginQuery(buf.fps[vkCmdBeginQuery], buf.hnd, queryPool.hnd, C.uint32_t(query), C.VkQueryControlFlags(flags))
}

func (buf *CommandBuffer) EndQuery(queryPool QueryPool, query uint32) {
	C.domVkCmdEndQuery(buf.fps[vkCmdEndQuery], buf.hnd, queryPool.hnd, C.uint32_t(query))
}

func (buf *CommandBuffer) CopyQueryPoolResults(
	queryPool QueryPool,
	firstQuery, queryCount uint32,
	dstBuffer Buffer,
	dstOffset, stride DeviceSize,
	flags QueryResultFlags,
) {
	C.domVkCmdCopyQueryPoolResults(buf.fps[vkCmdCopyQueryPoolResults],
		buf.hnd,
		queryPool.hnd,
		C.uint32_t(firstQuery),
		C.uint32_t(queryCount),
		dstBuffer.hnd,
		C.VkDeviceSize(dstOffset),
		C.VkDeviceSize(stride),
		C.VkQueryResultFlags(flags))
}

type ImageBlit struct {
	// must be kept identical to C struct
	_ structs.HostLayout

	SrcSubresource ImageSubresourceLayers
	SrcOffsets     [2]Offset3D
	DstSubresource ImageSubresourceLayers
	DstOffsets     [2]Offset3D
}

func (buf *CommandBuffer) BlitImage(
	srcImage Image,
	srcImageLayout ImageLayout,
	dstImage Image,
	dstImageLayout ImageLayout,
	regions []ImageBlit,
	filter Filter,
) {
	C.domVkCmdBlitImage(
		buf.fps[vkCmdBlitImage],
		buf.hnd,
		srcImage.hnd,
		C.VkImageLayout(srcImageLayout),
		dstImage.hnd,
		C.VkImageLayout(dstImageLayout),
		C.uint32_t(len(regions)),
		safeish.SliceCastPtr[*C.VkImageBlit](regions),
		C.VkFilter(filter))
}

func (buf *CommandBuffer) DrawIndexed(indexCount, instanceCount, firstIndex uint32, vertexOffset int32, firstInstance uint32) {
	C.domVkCmdDrawIndexed(
		buf.fps[vkCmdDrawIndexed],
		buf.hnd,
		C.uint32_t(indexCount),
		C.uint32_t(instanceCount),
		C.uint32_t(firstIndex),
		C.int32_t(vertexOffset),
		C.uint32_t(firstInstance))
}

func (buf *CommandBuffer) DrawIndexedIndirect(buffer Buffer, offset DeviceSize, drawCount, stride uint32) {
	C.domVkCmdDrawIndexedIndirect(
		buf.fps[vkCmdDrawIndexedIndirect],
		buf.hnd,
		buffer.hnd,
		C.VkDeviceSize(offset),
		C.uint32_t(drawCount),
		C.uint32_t(stride))
}

func (buf *CommandBuffer) DrawIndirect(buffer Buffer, offset DeviceSize, drawCount, stride uint32) {
	C.domVkCmdDrawIndirect(buf.fps[vkCmdDrawIndirect], buf.hnd, buffer.hnd, C.VkDeviceSize(offset), C.uint32_t(drawCount), C.uint32_t(stride))
}

type MemoryBarrier struct {
	Extensions    []Extension
	SrcAccessMask AccessFlags
	DstAccessMask AccessFlags
}

type BufferMemoryBarrier struct {
	Extensions          []Extension
	SrcAccessMask       AccessFlags
	DstAccessMask       AccessFlags
	SrcQueueFamilyIndex uint32
	DstQueueFamilyIndex uint32
	Buffer              Buffer
	Offset              DeviceSize
	Size                DeviceSize
}

type ImageMemoryBarrier struct {
	Extensions          []Extension
	SrcAccessMask       AccessFlags
	DstAccessMask       AccessFlags
	OldLayout           ImageLayout
	NewLayout           ImageLayout
	SrcQueueFamilyIndex uint32
	DstQueueFamilyIndex uint32
	Image               Image
	SubresourceRange    ImageSubresourceRange
}

func barriers(
	a *allocator,
	memoryBarriers []MemoryBarrier,
	bufferMemoryBarriers []BufferMemoryBarrier,
	imageMemoryBarriers []ImageMemoryBarrier,
) (*C.VkMemoryBarrier, *C.VkBufferMemoryBarrier, *C.VkImageMemoryBarrier) {
	cmem := allocn[C.VkMemoryBarrier](a, len(memoryBarriers))
	cbuf := allocn[C.VkBufferMemoryBarrier](a, len(bufferMemoryBarriers))
	cimg := allocn[C.VkImageMemoryBarrier](a, len(imageMemoryBarriers))
	memArr := unsafe.Slice(cmem, len(memoryBarriers))
	bufArr := unsafe.Slice(cbuf, len(bufferMemoryBarriers))
	imgArr := unsafe.Slice(cimg, len(imageMemoryBarriers))

	for i := range memArr {
		memArr[i] = C.VkMemoryBarrier{
			sType:         C.VkStructureType(StructureTypeMemoryBarrier),
			pNext:         buildChain(a, memoryBarriers[i].Extensions),
			srcAccessMask: C.VkAccessFlags(memoryBarriers[i].SrcAccessMask),
			dstAccessMask: C.VkAccessFlags(memoryBarriers[i].DstAccessMask),
		}
		// XXX this defer is probably wrong and needs to be in the caller of barriers
		defer internalizeChain(memoryBarriers[i].Extensions, memArr[i].pNext)
	}

	for i := range bufArr {
		bufArr[i] = C.VkBufferMemoryBarrier{
			sType:               C.VkStructureType(StructureTypeBufferMemoryBarrier),
			pNext:               buildChain(a, bufferMemoryBarriers[i].Extensions),
			srcAccessMask:       C.VkAccessFlags(bufferMemoryBarriers[i].SrcAccessMask),
			dstAccessMask:       C.VkAccessFlags(bufferMemoryBarriers[i].DstAccessMask),
			srcQueueFamilyIndex: C.uint32_t(bufferMemoryBarriers[i].SrcQueueFamilyIndex),
			dstQueueFamilyIndex: C.uint32_t(bufferMemoryBarriers[i].DstQueueFamilyIndex),
			buffer:              bufferMemoryBarriers[i].Buffer.hnd,
			offset:              C.VkDeviceSize(bufferMemoryBarriers[i].Offset),
			size:                C.VkDeviceSize(bufferMemoryBarriers[i].Size),
		}
		// XXX this defer is probably wrong and needs to be in the caller of barriers
		defer internalizeChain(bufferMemoryBarriers[i].Extensions, bufArr[i].pNext)
	}

	for i := range imgArr {
		imgArr[i] = C.VkImageMemoryBarrier{
			sType:               C.VkStructureType(StructureTypeImageMemoryBarrier),
			pNext:               buildChain(a, imageMemoryBarriers[i].Extensions),
			srcAccessMask:       C.VkAccessFlags(imageMemoryBarriers[i].SrcAccessMask),
			dstAccessMask:       C.VkAccessFlags(imageMemoryBarriers[i].DstAccessMask),
			oldLayout:           C.VkImageLayout(imageMemoryBarriers[i].OldLayout),
			newLayout:           C.VkImageLayout(imageMemoryBarriers[i].NewLayout),
			srcQueueFamilyIndex: C.uint32_t(imageMemoryBarriers[i].SrcQueueFamilyIndex),
			dstQueueFamilyIndex: C.uint32_t(imageMemoryBarriers[i].DstQueueFamilyIndex),
			image:               imageMemoryBarriers[i].Image.hnd,
		}
		ucopy1(unsafe.Pointer(&imgArr[i].subresourceRange), unsafe.Pointer(&imageMemoryBarriers[i].SubresourceRange), C.sizeof_VkImageSubresourceRange)
		// XXX this defer is probably wrong and needs to be in the caller of barriers
		defer internalizeChain(imageMemoryBarriers[i].Extensions, imgArr[i].pNext)
	}

	return (*C.VkMemoryBarrier)(cmem),
		(*C.VkBufferMemoryBarrier)(cbuf),
		(*C.VkImageMemoryBarrier)(cimg)
}

func (buf *CommandBuffer) WaitEvents(
	events []Event,
	srcStageMask, dstStageMask PipelineStageFlags,
	memoryBarriers []MemoryBarrier,
	bufferMemoryBarriers []BufferMemoryBarrier,
	imageMemoryBarriers []ImageMemoryBarrier,
) {
	a := new(allocator)
	defer a.free()

	cmem, cbuf, cimg := barriers(a, memoryBarriers, bufferMemoryBarriers, imageMemoryBarriers)
	C.domVkCmdWaitEvents(
		buf.fps[vkCmdWaitEvents],
		buf.hnd,
		C.uint32_t(len(events)),
		safeish.SliceCastPtr[*C.VkEvent](events),
		C.VkPipelineStageFlags(srcStageMask),
		C.VkPipelineStageFlags(dstStageMask),
		C.uint32_t(len(memoryBarriers)),
		cmem,
		C.uint32_t(len(bufferMemoryBarriers)),
		cbuf,
		C.uint32_t(len(imageMemoryBarriers)),
		cimg)
}

func (buf *CommandBuffer) NextSubpass(contents SubpassContents) {
	C.domVkCmdNextSubpass(buf.fps[vkCmdNextSubpass], buf.hnd, C.VkSubpassContents(contents))
}

func (buf *CommandBuffer) SetStencilCompareMask(faceMask StencilFaceFlags, compareMask uint32) {
	C.domVkCmdSetStencilCompareMask(buf.fps[vkCmdSetStencilCompareMask], buf.hnd, C.VkStencilFaceFlags(faceMask), C.uint32_t(compareMask))
}

func (buf *CommandBuffer) SetStencilReference(faceMask StencilFaceFlags, reference uint32) {
	C.domVkCmdSetStencilReference(buf.fps[vkCmdSetStencilReference], buf.hnd, C.VkStencilFaceFlags(faceMask), C.uint32_t(reference))
}

func (buf *CommandBuffer) SetStencilWriteMask(faceMask StencilFaceFlags, writeMask uint32) {
	C.domVkCmdSetStencilWriteMask(buf.fps[vkCmdSetStencilWriteMask], buf.hnd, C.VkStencilFaceFlags(faceMask), C.uint32_t(writeMask))
}

func (buf *CommandBuffer) WriteTimestamp(pipelineStage PipelineStageFlags, queryPool QueryPool, query uint32) {
	C.domVkCmdWriteTimestamp(buf.fps[vkCmdWriteTimestamp], buf.hnd, C.VkPipelineStageFlagBits(pipelineStage), queryPool.hnd, C.uint32_t(query))
}

func (buf *CommandBuffer) BindVertexBuffers(firstBinding uint32, buffers []Buffer, offsets []DeviceSize) {
	if safe && len(buffers) != len(offsets) {
		panic("buffers and offsets must have same length")
	}
	C.domVkCmdBindVertexBuffers(buf.fps[vkCmdBindVertexBuffers],
		buf.hnd,
		C.uint32_t(firstBinding),
		C.uint32_t(len(buffers)),
		safeish.SliceCastPtr[*C.VkBuffer](buffers),
		safeish.SliceCastPtr[*C.VkDeviceSize](offsets))
}

func (buf *CommandBuffer) ExecuteCommands(buffers []CommandBuffer) {
	if len(buffers) == 1 {
		C.domVkCmdExecuteCommands(buf.fps[vkCmdExecuteCommands], buf.hnd, 1, (*C.VkCommandBuffer)(unsafe.Pointer(&buffers[0].hnd)))
		return
	}
	arr := buf.bufs
	if cap(arr) >= len(buffers) {
		arr = arr[:len(buffers)]
	} else {
		arr = make([]C.VkCommandBuffer, len(buffers))
	}
	for i, cmd := range buffers {
		arr[i] = cmd.hnd
	}
	C.domVkCmdExecuteCommands(buf.fps[vkCmdExecuteCommands], buf.hnd, C.uint32_t(len(buffers)), unsafe.SliceData(arr))
	buf.bufs = arr
}

func (buf *CommandBuffer) DispatchBase(baseGroupX, baseGroupY, baseGroupZ, groupCountX, groupCountY, groupCountZ uint32) {
	C.domVkCmdDispatchBase(
		buf.fps[vkCmdDispatchBase],
		buf.hnd,
		C.uint32_t(baseGroupX),
		C.uint32_t(baseGroupY),
		C.uint32_t(baseGroupZ),
		C.uint32_t(groupCountX),
		C.uint32_t(groupCountY),
		C.uint32_t(groupCountZ))
}

func (buf *CommandBuffer) DispatchIndirect(buffer Buffer, offset DeviceSize) {
	C.domVkCmdDispatchIndirect(buf.fps[vkCmdDispatchIndirect], buf.hnd, buffer.hnd, C.VkDeviceSize(offset))
}

// PipelineBarrier inserts a memory dependency.
//
// When PipelineBarrier is submitted to a queue, it defines a memory dependency between commands that were submitted before it, and those submitted after it.
//
// If PipelineBarrier was recorded outside a render pass instance,
// the first synchronization scope includes all commands that occur earlier in submission order.
// If PipelineBarrier was recorded inside a render pass instance,
// the first synchronization scope includes only commands that occur earlier in submission order within the same subpass.
// In either case, the first synchronization scope is limited to operations on the pipeline stages determined by the source stage mask specified by srcStageMask.
//
// If PipelineBarrier was recorded outside a render pass instance,
// the second synchronization scope includes all commands that occur later in submission order.
// If PipelineBarrier was recorded inside a render pass instance,
// the second synchronization scope includes only commands that occur later in submission order within the same subpass.
// In either case, the second synchronization scope is limited to operations on the pipeline stages determined by the destination stage mask specified by dstStageMask.
//
// The first access scope is limited to access in the pipeline stages determined by the source stage mask specified by srcStageMask.
// Within that, the first access scope only includes the first access scopes defined by elements of
// the memoryBarriers, bufferMemoryBarriers and imageMemoryBarriers arrays, which each define a set of memory barriers.
// If no memory barriers are specified, then the first access scope includes no accesses.
//
// The second access scope is limited to access in the pipeline stages determined by the destination stage mask specified by dstStageMask.
// Within that, the second access scope only includes the second access scopes defined by elements of
// the memoryBarriers, bufferMemoryBarriers and imageMemoryBarriers arrays, which each define a set of memory barriers.
// If no memory barriers are specified, then the second access scope includes no accesses.
//
// If dependencyFlags includes DependencyByRegionBit,
// then any dependency between framebuffer-space pipeline stages is framebuffer-local -
// otherwise it is framebuffer-global.
func (buf *CommandBuffer) PipelineBarrier(
	srcStageMask PipelineStageFlags,
	dstStageMask PipelineStageFlags,
	dependencyFlags DependencyFlags,
	memoryBarriers []MemoryBarrier,
	bufferMemoryBarriers []BufferMemoryBarrier,
	imageMemoryBarriers []ImageMemoryBarrier,
) {
	a := new(allocator)
	defer a.free()

	cmem, cbuf, cimg := barriers(a, memoryBarriers, bufferMemoryBarriers, imageMemoryBarriers)
	C.domVkCmdPipelineBarrier(
		buf.fps[vkCmdPipelineBarrier],
		buf.hnd,
		C.VkPipelineStageFlags(srcStageMask),
		C.VkPipelineStageFlags(dstStageMask),
		C.VkDependencyFlags(dependencyFlags),
		C.uint32_t(len(memoryBarriers)),
		cmem,
		C.uint32_t(len(bufferMemoryBarriers)),
		cbuf,
		C.uint32_t(len(imageMemoryBarriers)),
		cimg)
}

type ImageResolve struct {
	// must be kept identical to C struct
	_ structs.HostLayout

	SrcSubresource ImageSubresourceLayers
	SrcOffset      Offset3D
	DstSubresource ImageSubresourceLayers
	DstOffset      Offset3D
	Extent         Extent3D
}

// ResolveImage resolves a multisample image to a non-multisample image.
//
// During the resolve the samples corresponding to each pixel location in the source
// are converted to a single sample before being written to the destination.
// If the source formats are floating-point or normalized types,
// the sample values for each pixel are resolved in an implementation-dependent manner.
// If the source formats are integer types, a single sample’s value is selected for each pixel.
//
// srcOffset and dstOffset select the initial x, y, and z offsets
// in texels of the sub-regions of the source and destination image data.
// extent is the size in texels of the source image to resolve in width, height and depth.
//
// Resolves are done layer by layer starting with BaseArrayLayer field of srcSubresource for the source
// and dstSubresource for the destination.
// layerCount layers are resolved to the destination image.
func (buf *CommandBuffer) ResolveImage(srcImage Image, srcImageLayout ImageLayout, dstImage Image, dstImageLayout ImageLayout, regions []ImageResolve) {
	C.domVkCmdResolveImage(
		buf.fps[vkCmdResolveImage],
		buf.hnd,
		srcImage.hnd,
		C.VkImageLayout(srcImageLayout),
		dstImage.hnd,
		C.VkImageLayout(dstImageLayout),
		C.uint32_t(len(regions)),
		safeish.SliceCastPtr[*C.VkImageResolve](regions))
}

func (buf *CommandBuffer) BindDescriptorSets(
	pipelineBindPoint PipelineBindPoint,
	layout PipelineLayout,
	firstSet uint32,
	descriptorSets []DescriptorSet,
	dynamicOffsets []uint32,
) {
	C.domVkCmdBindDescriptorSets(
		buf.fps[vkCmdBindDescriptorSets],
		buf.hnd,
		C.VkPipelineBindPoint(pipelineBindPoint),
		layout.hnd,
		C.uint32_t(firstSet),
		C.uint32_t(len(descriptorSets)),
		safeish.SliceCastPtr[*C.VkDescriptorSet](descriptorSets),
		C.uint32_t(len(dynamicOffsets)),
		safeish.SliceCastPtr[*C.uint32_t](dynamicOffsets),
	)
}

type CommandPoolCreateInfo struct {
	Extensions       []Extension
	Flags            CommandPoolCreateFlags
	QueueFamilyIndex uint32
}

// Command pools are opaque objects that command buffer memory is allocated from,
// and which allow the implementation to amortize the cost of resource creation across multiple command buffers.
// Command pools are externally synchronized, meaning that a command pool must not be used concurrently in multiple threads.
// That includes use via recording commands on any command buffers allocated from the pool,
// as well as operations that allocate, free, and reset command buffers or the pool itself.
type CommandPool struct {
	// VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkCommandPool)
	hnd C.VkCommandPool
}

func (dev *Device) CreateCommandPool(info *CommandPoolCreateInfo) (CommandPool, error) {
	a := new(allocator)
	defer a.free()

	// TODO(dh): support callbacks
	var ptr C.VkCommandPoolCreateInfo
	ptr.sType = C.VkStructureType(StructureTypeCommandPoolCreateInfo)
	ptr.pNext = buildChain(a, info.Extensions)
	defer internalizeChain(info.Extensions, ptr.pNext)
	ptr.flags = C.VkCommandPoolCreateFlags(info.Flags)
	ptr.queueFamilyIndex = C.uint32_t(info.QueueFamilyIndex)
	var pool CommandPool
	res := Result(C.domVkCreateCommandPool(dev.fps[vkCreateCommandPool], dev.hnd, &ptr, nil, &pool.hnd))
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
	a := new(allocator)
	defer a.free()

	var ptr C.VkCommandBufferAllocateInfo
	ptr.sType = C.VkStructureType(StructureTypeCommandBufferAllocateInfo)
	ptr.pNext = buildChain(a, info.Extensions)
	defer internalizeChain(info.Extensions, ptr.pNext)
	ptr.commandPool = pool.hnd
	ptr.level = C.VkCommandBufferLevel(info.Level)
	ptr.commandBufferCount = C.uint32_t(info.CommandBufferCount)
	bufs := make([]C.VkCommandBuffer, info.CommandBufferCount)
	res := Result(C.domVkAllocateCommandBuffers(dev.fps[vkAllocateCommandBuffers], dev.hnd, &ptr, unsafe.SliceData(bufs)))
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
		C.domVkFreeCommandBuffers(dev.fps[vkFreeCommandBuffers], dev.hnd, pool.hnd, 1, (*C.VkCommandBuffer)(unsafe.Pointer(bufs[0].hnd)))
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
	C.domVkFreeCommandBuffers(dev.fps[vkFreeCommandBuffers], dev.hnd, pool.hnd, C.uint32_t(len(bufs)), unsafe.SliceData(ptrs))
}

func (dev *Device) WaitIdle() error {
	res := Result(C.domVkDeviceWaitIdle(dev.fps[vkDeviceWaitIdle], dev.hnd))
	return result2error(res)
}

// Images represent multidimensional - up to 3 - arrays of data,
// which can be used for various purposes (e.g. attachments, textures),
// by binding them to a graphics or compute pipeline via descriptor sets,
// or by directly specifying them as parameters to certain commands.
type Image struct {
	// VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkImage)
	hnd C.VkImage
}

// Image objects are not directly accessed by pipeline shaders for reading or writing image data.
// Instead, image views representing contiguous ranges of the image subresources
// and containing additional metadata are used for that purpose.
// Views must be created on images of compatible types, and must represent a valid subset of image subresources.
type ImageView struct {
	// must be kept identical to C struct
	_ structs.HostLayout

	// VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkImageView)
	hnd C.VkImageView
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
	// must be kept identical to C struct
	_ structs.HostLayout

	R ComponentSwizzle
	G ComponentSwizzle
	B ComponentSwizzle
	A ComponentSwizzle
}

type ImageSubresourceRange struct {
	// must be kept identical to C struct
	_ structs.HostLayout

	AspectMask     ImageAspectFlags
	BaseMipLevel   uint32
	LevelCount     uint32
	BaseArrayLayer uint32
	LayerCount     uint32
}

func (dev *Device) CreateImageView(info *ImageViewCreateInfo) (ImageView, error) {
	a := new(allocator)
	defer a.free()

	// TODO(dh): support custom allocator
	var ptr C.VkImageViewCreateInfo
	ptr.sType = C.VkStructureType(StructureTypeImageViewCreateInfo)
	ptr.pNext = buildChain(a, info.Extensions)
	defer internalizeChain(info.Extensions, ptr.pNext)
	ptr.image = info.Image.hnd
	ptr.viewType = C.VkImageViewType(info.ViewType)
	ptr.format = C.VkFormat(info.Format)
	ucopy1(unsafe.Pointer(&ptr.components), unsafe.Pointer(&info.Components), C.sizeof_VkComponentMapping)
	ucopy1(unsafe.Pointer(&ptr.subresourceRange), unsafe.Pointer(&info.SubresourceRange), C.sizeof_VkImageSubresourceRange)

	var out ImageView
	res := Result(C.domVkCreateImageView(dev.fps[vkCreateImageView], dev.hnd, &ptr, nil, &out.hnd))
	return out, result2error(res)
}

func (dev *Device) DestroyImageView(view ImageView) {
	// TODO(dh): support custom allocator
	C.domVkDestroyImageView(dev.fps[vkDestroyImageView], dev.hnd, view.hnd, nil)
}

// Shader modules contain shader code and one or more entry points.
// Shaders are selected from a shader module by specifying an entry point as part of pipeline creation.
// The stages of a pipeline can use shaders that come from different modules.
// The shader code defining a shader module must be in the SPIR-V format, as described by the Vulkan Environment for SPIR-V appendix.
type ShaderModule struct {
	// VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkShaderModule)
	hnd C.VkShaderModule
}

type ShaderModuleCreateInfo struct {
	Extensions []Extension
	Code       []byte
}

func (dev *Device) CreateShaderModule(info *ShaderModuleCreateInfo) (ShaderModule, error) {
	a := new(allocator)
	defer a.free()

	var ptr C.VkShaderModuleCreateInfo
	ptr.sType = C.VkStructureType(StructureTypeShaderModuleCreateInfo)
	ptr.pNext = buildChain(a, info.Extensions)
	defer internalizeChain(info.Extensions, ptr.pNext)
	ptr.codeSize = C.size_t(len(info.Code))
	ptr.pCode = (*C.uint32_t)(C.CBytes(info.Code))
	defer C.free(unsafe.Pointer(ptr.pCode))
	var out ShaderModule
	res := Result(C.domVkCreateShaderModule(dev.fps[vkCreateShaderModule], dev.hnd, &ptr, nil, &out.hnd))
	return out, result2error(res)
}

func (dev *Device) DestroyShaderModule(module ShaderModule) {
	// TODO(dh): support custom allocator
	C.domVkDestroyShaderModule(dev.fps[vkDestroyShaderModule], dev.hnd, module.hnd, nil)
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

func (info *PipelineVertexInputStateCreateInfo) c(a *allocator) *C.VkPipelineVertexInputStateCreateInfo {
	cinfo := alloc[C.VkPipelineVertexInputStateCreateInfo](a)
	bindings := allocn[C.VkVertexInputBindingDescription](a, len(info.VertexBindingDescriptions))
	attribs := allocn[C.VkVertexInputAttributeDescription](a, len(info.VertexAttributeDescriptions))
	*cinfo = C.VkPipelineVertexInputStateCreateInfo{
		sType:                           C.VkStructureType(StructureType(StructureTypePipelineVertexInputStateCreateInfo)),
		pNext:                           buildChain(a, info.Extensions),
		flags:                           0,
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
	// must be kept identical to C struct
	_ structs.HostLayout

	Binding   uint32
	Stride    uint32
	InputRate VertexInputRate
}

type VertexInputAttributeDescription struct {
	// must be kept identical to C struct
	_ structs.HostLayout

	Location uint32
	Binding  uint32
	Format   Format
	Offset   uint32
}

type PipelineInputAssemblyStateCreateInfo struct {
	Extensions             []Extension
	Topology               PrimitiveTopology
	PrimitiveRestartEnable bool
}

func (info *PipelineInputAssemblyStateCreateInfo) c(a *allocator) *C.VkPipelineInputAssemblyStateCreateInfo {
	cinfo := alloc[C.VkPipelineInputAssemblyStateCreateInfo](a)
	*cinfo = C.VkPipelineInputAssemblyStateCreateInfo{
		sType:                  C.VkStructureType(StructureTypePipelineInputAssemblyStateCreateInfo),
		pNext:                  buildChain(a, info.Extensions),
		flags:                  0,
		topology:               C.VkPrimitiveTopology(info.Topology),
		primitiveRestartEnable: vkBool(info.PrimitiveRestartEnable),
	}
	return cinfo
}

type Viewport struct {
	// must be kept identical to C struct
	_ structs.HostLayout

	X        float32
	Y        float32
	Width    float32
	Height   float32
	MinDepth float32
	MaxDepth float32
}

type Rect2D struct {
	// must be kept identical to C struct
	_ structs.HostLayout

	Offset Offset2D
	Extent Extent2D
}

type Offset2D struct {
	// must be kept identical to C struct
	_ structs.HostLayout

	X int32
	Y int32
}

type Offset3D struct {
	// must be kept identical to C struct
	_ structs.HostLayout

	X int32
	Y int32
	Z int32
}

type PipelineViewportStateCreateInfo struct {
	Extensions []Extension
	Viewports  []Viewport
	Scissors   []Rect2D
}

func (info *PipelineViewportStateCreateInfo) c(a *allocator) *C.VkPipelineViewportStateCreateInfo {
	cinfo := alloc[C.VkPipelineViewportStateCreateInfo](a)
	viewports := allocn[C.VkViewport](a, len(info.Viewports))
	scissors := allocn[C.VkRect2D](a, len(info.Scissors))
	*cinfo = C.VkPipelineViewportStateCreateInfo{
		sType:         C.VkStructureType(StructureTypePipelineViewportStateCreateInfo),
		pNext:         buildChain(a, info.Extensions),
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

func (info *PipelineRasterizationStateCreateInfo) c(a *allocator) *C.VkPipelineRasterizationStateCreateInfo {
	cinfo := alloc[C.VkPipelineRasterizationStateCreateInfo](a)
	*cinfo = C.VkPipelineRasterizationStateCreateInfo{
		sType:                   C.VkStructureType(StructureTypePipelineRasterizationStateCreateInfo),
		pNext:                   buildChain(a, info.Extensions),
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

func (info *PipelineMultisampleStateCreateInfo) c(a *allocator) *C.VkPipelineMultisampleStateCreateInfo {
	if safe && info.SampleMask != nil {
		if (info.RasterizationSamples == 64 && len(info.SampleMask) < 2) || len(info.SampleMask) < 1 {
			panic("SampleMask must be nil or have ceil(rasterizationSamples / 32) elements")
		}
	}
	cinfo := alloc[C.VkPipelineMultisampleStateCreateInfo](a)
	var sampleMask *C.VkSampleMask
	if info.SampleMask != nil {
		sampleMask = allocn[C.VkSampleMask](a, len(info.SampleMask))
		ucopy(unsafe.Pointer(sampleMask), unsafe.Pointer(&info.SampleMask), C.sizeof_VkSampleMask)
	}

	*cinfo = C.VkPipelineMultisampleStateCreateInfo{
		sType:                 C.VkStructureType(StructureTypePipelineMultisampleStateCreateInfo),
		pNext:                 buildChain(a, info.Extensions),
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

func (info *PipelineColorBlendStateCreateInfo) c(a *allocator) *C.VkPipelineColorBlendStateCreateInfo {
	cinfo := alloc[C.VkPipelineColorBlendStateCreateInfo](a)
	attachments := allocn[C.VkPipelineColorBlendAttachmentState](a, len(info.Attachments))
	*cinfo = C.VkPipelineColorBlendStateCreateInfo{
		sType:           C.VkStructureType(StructureTypePipelineColorBlendStateCreateInfo),
		pNext:           buildChain(a, info.Extensions),
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
	attachmentsArr := unsafe.Slice(attachments, len(info.Attachments))
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

func (info *PipelineDynamicStateCreateInfo) c(a *allocator) *C.VkPipelineDynamicStateCreateInfo {
	cinfo := alloc[C.VkPipelineDynamicStateCreateInfo](a)
	dynamicStates := allocn[C.VkDynamicState](a, len(info.DynamicStates))
	*cinfo = C.VkPipelineDynamicStateCreateInfo{
		sType:             C.VkStructureType(StructureTypePipelineDynamicStateCreateInfo),
		pNext:             buildChain(a, info.Extensions),
		flags:             0,
		dynamicStateCount: C.uint32_t(len(info.DynamicStates)),
		pDynamicStates:    dynamicStates,
	}
	ucopy(unsafe.Pointer(dynamicStates), unsafe.Pointer(&info.DynamicStates), C.sizeof_VkDynamicState)
	return cinfo
}

// Access to descriptor sets from a pipeline is accomplished through a pipeline layout.
// Zero or more descriptor set layouts and zero or more push constant ranges are combined
// to form a pipeline layout object which describes the complete set of resources that can be accessed by a pipeline.
// The pipeline layout represents a sequence of descriptor sets with each having a specific layout.
// This sequence of layouts is used to determine the interface between shader stages and shader resources.
// Each pipeline is created using a pipeline layout.
type PipelineLayout struct {
	// VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkPipelineLayout)
	hnd C.VkPipelineLayout
}

type PipelineLayoutCreateInfo struct {
	Extensions         []Extension
	SetLayouts         []DescriptorSetLayout
	PushConstantRanges []PushConstantRange
}

func (info *PipelineLayoutCreateInfo) c(a *allocator) *C.VkPipelineLayoutCreateInfo {
	cinfo := alloc[C.VkPipelineLayoutCreateInfo](a)
	setLayouts := allocn[C.VkDescriptorSetLayout](a, len(info.SetLayouts))
	push := allocn[C.VkPushConstantRange](a, len(info.PushConstantRanges))
	*cinfo = C.VkPipelineLayoutCreateInfo{
		sType:                  C.VkStructureType(StructureTypePipelineLayoutCreateInfo),
		pNext:                  buildChain(a, info.Extensions),
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
	a := new(allocator)
	defer a.free()

	cinfo := info.c(a)
	var out PipelineLayout
	defer internalizeChain(info.Extensions, cinfo.pNext)
	res := Result(C.domVkCreatePipelineLayout(dev.fps[vkCreatePipelineLayout], dev.hnd, cinfo, nil, &out.hnd))
	return out, result2error(res)
}

func (dev *Device) DestroyPipelineLayout(layout PipelineLayout) {
	// TODO(dh): support custom allocators
	C.domVkDestroyPipelineLayout(dev.fps[vkDestroyPipelineLayout], dev.hnd, layout.hnd, nil)
}

// A descriptor set layout object is defined by an array of zero or more descriptor bindings.
// Each individual descriptor binding is specified by
// a descriptor type,
// a count (array size) of the number of descriptors in the binding,
// a set of shader stages that can access the binding,
// and (if using immutable samplers) an array of sampler descriptors.
type DescriptorSetLayout struct {
	// must be kept identical to C struct
	_ structs.HostLayout

	// VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkDescriptorSetLayout)
	hnd C.VkDescriptorSetLayout
}

type PushConstantRange struct {
	// must be kept identical to C struct
	_ structs.HostLayout

	StageFlags ShaderStageFlags
	Offset     uint32
	Size       uint32
}

type PipelineTessellationStateCreateInfo struct {
	Extensions         []Extension
	PatchControlPoints uint32
}

func (info *PipelineTessellationStateCreateInfo) c(a *allocator) *C.VkPipelineTessellationStateCreateInfo {
	cinfo := alloc[C.VkPipelineTessellationStateCreateInfo](a)
	*cinfo = C.VkPipelineTessellationStateCreateInfo{
		sType:              C.VkStructureType(StructureTypePipelineTessellationStateCreateInfo),
		pNext:              buildChain(a, info.Extensions),
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

func (info *PipelineDepthStencilStateCreateInfo) c(a *allocator) *C.VkPipelineDepthStencilStateCreateInfo {
	cinfo := alloc[C.VkPipelineDepthStencilStateCreateInfo](a)
	*cinfo = C.VkPipelineDepthStencilStateCreateInfo{
		sType:                 C.VkStructureType(StructureTypePipelineDepthStencilStateCreateInfo),
		pNext:                 buildChain(a, info.Extensions),
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

// Pipeline is an opaque handle to a pipeline object.
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

func (dev *Device) CreateGraphicsPipelines(cache *PipelineCache, infos []GraphicsPipelineCreateInfo) ([]Pipeline, error) {
	a := new(allocator)
	defer a.free()

	ptrs := allocn[C.VkGraphicsPipelineCreateInfo](a, len(infos))

	ptrsArr := (*[math.MaxInt32]C.VkGraphicsPipelineCreateInfo)(unsafe.Pointer(ptrs))[:len(infos)]
	for i := range ptrsArr {
		ptr := &ptrsArr[i]
		info := &infos[i]

		ptr.sType = C.VkStructureType(StructureTypeGraphicsPipelineCreateInfo)
		ptr.pNext = buildChain(a, info.Extensions)
		defer internalizeChain(info.Extensions, ptr.pNext)
		ptr.flags = C.VkPipelineCreateFlags(info.Flags)
		ptr.stageCount = C.uint32_t(len(info.Stages))

		ptr.pStages = allocn[C.VkPipelineShaderStageCreateInfo](a, len(info.Stages))
		arr := (*[math.MaxInt32]C.VkPipelineShaderStageCreateInfo)(unsafe.Pointer(ptr.pStages))[:len(info.Stages)]
		for i := range arr {
			arr[i] = C.VkPipelineShaderStageCreateInfo{
				sType:  C.VkStructureType(StructureTypePipelineShaderStageCreateInfo),
				pNext:  buildChain(a, info.Stages[i].Extensions),
				stage:  C.VkShaderStageFlagBits(info.Stages[i].Stage),
				module: info.Stages[i].Module.hnd,
				pName:  externString(a, info.Stages[i].Name),
			}
			defer internalizeChain(info.Stages[i].Extensions, arr[i].pNext)
		}

		if info.VertexInputState != nil {
			ptr.pVertexInputState = info.VertexInputState.c(a)
			defer internalizeChain(info.VertexInputState.Extensions, ptr.pVertexInputState.pNext)
		}
		if info.InputAssemblyState != nil {
			ptr.pInputAssemblyState = info.InputAssemblyState.c(a)
			defer internalizeChain(info.InputAssemblyState.Extensions, ptr.pInputAssemblyState.pNext)
		}
		if info.TessellationState != nil {
			ptr.pTessellationState = info.TessellationState.c(a)
			defer internalizeChain(info.TessellationState.Extensions, ptr.pTessellationState.pNext)
		}
		if info.ViewportState != nil {
			ptr.pViewportState = info.ViewportState.c(a)
			defer internalizeChain(info.ViewportState.Extensions, ptr.pViewportState.pNext)
		}
		if info.RasterizationState != nil {
			ptr.pRasterizationState = info.RasterizationState.c(a)
			defer internalizeChain(info.RasterizationState.Extensions, ptr.pRasterizationState.pNext)
		}
		if info.MultisampleState != nil {
			ptr.pMultisampleState = info.MultisampleState.c(a)
			defer internalizeChain(info.MultisampleState.Extensions, ptr.pMultisampleState.pNext)
		}
		if info.DepthStencilState != nil {
			ptr.pDepthStencilState = info.DepthStencilState.c(a)
			defer internalizeChain(info.DepthStencilState.Extensions, ptr.pDepthStencilState.pNext)
		}
		if info.ColorBlendState != nil {
			ptr.pColorBlendState = info.ColorBlendState.c(a)
			defer internalizeChain(info.ColorBlendState.Extensions, ptr.pColorBlendState.pNext)
		}
		if info.DynamicState != nil {
			ptr.pDynamicState = info.DynamicState.c(a)
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

	var cacheHnd C.VkPipelineCache
	if cache != nil {
		cacheHnd = cache.hnd
	}
	hnds := make([]C.VkPipeline, len(infos))
	res := Result(C.domVkCreateGraphicsPipelines(
		dev.fps[vkCreateGraphicsPipelines],
		dev.hnd,
		cacheHnd,
		C.uint32_t(len(infos)),
		ptrs,
		nil,
		safeish.SliceCastPtr[*C.VkPipeline](hnds)))
	if res != Success {
		return nil, res
	}
	out := make([]Pipeline, len(infos))
	for i, hnd := range hnds {
		out[i] = Pipeline{hnd}
	}
	return out, nil
}

func (dev *Device) DestroyPipeline(pipeline Pipeline) {
	// TODO(dh): support a custom allocator
	C.domVkDestroyPipeline(dev.fps[vkDestroyPipeline], dev.hnd, pipeline.hnd, nil)
}

type AttachmentDescription struct {
	_ structs.HostLayout

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
	_ structs.HostLayout

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
	Extensions   []Extension
	Attachments  []AttachmentDescription
	Subpasses    []SubpassDescription
	Dependencies []SubpassDependency
}

type SubpassDependency struct {
	// must be kept identical to C struct
	_ structs.HostLayout

	SrcSubpass      uint32
	DstSubpass      uint32
	SrcStageMask    PipelineStageFlags
	DstStageMask    PipelineStageFlags
	SrcAccessMask   AccessFlags
	DstAccessMask   AccessFlags
	DependencyFlags DependencyFlags
}

func (dev *Device) CreateRenderPass(info *RenderPassCreateInfo) (RenderPass, error) {
	a := new(allocator)
	defer a.free()

	cinfo := alloc[C.VkRenderPassCreateInfo](a)
	attachments := allocn[C.VkAttachmentDescription](a, len(info.Attachments))
	subpasses := allocn[C.VkSubpassDescription](a, len(info.Subpasses))
	dependencies := allocn[C.VkSubpassDependency](a, len(info.Dependencies))
	*cinfo = C.VkRenderPassCreateInfo{
		sType:           C.VkStructureType(StructureTypeRenderPassCreateInfo),
		pNext:           buildChain(a, info.Extensions),
		flags:           0,
		attachmentCount: C.uint32_t(len(info.Attachments)),
		pAttachments:    attachments,
		subpassCount:    C.uint32_t(len(info.Subpasses)),
		pSubpasses:      subpasses,
		dependencyCount: C.uint32_t(len(info.Dependencies)),
		pDependencies:   dependencies,
	}
	defer internalizeChain(info.Extensions, cinfo.pNext)
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
			pInputAttachments:       allocn[C.VkAttachmentReference](a, len(subpass.InputAttachments)),
			pColorAttachments:       allocn[C.VkAttachmentReference](a, len(subpass.ColorAttachments)),
			pPreserveAttachments:    allocn[C.uint32_t](a, len(subpass.PreserveAttachments)),
		}
		ucopy(unsafe.Pointer(csubpass.pInputAttachments), unsafe.Pointer(&subpass.InputAttachments), C.sizeof_VkAttachmentReference)
		ucopy(unsafe.Pointer(csubpass.pColorAttachments), unsafe.Pointer(&subpass.ColorAttachments), C.sizeof_VkAttachmentReference)
		if len(subpass.ResolveAttachments) > 0 {
			csubpass.pResolveAttachments = allocn[C.VkAttachmentReference](a, len(subpass.ResolveAttachments))
			ucopy(unsafe.Pointer(csubpass.pResolveAttachments), unsafe.Pointer(&subpass.ResolveAttachments), C.sizeof_VkAttachmentReference)
		}
		if subpass.DepthStencilAttachment != nil {
			csubpass.pDepthStencilAttachment = alloc[C.VkAttachmentReference](a)
			ucopy1(unsafe.Pointer(csubpass.pDepthStencilAttachment), unsafe.Pointer(subpass.DepthStencilAttachment), C.sizeof_VkAttachmentReference)
		}
		ucopy(unsafe.Pointer(csubpass.pPreserveAttachments), unsafe.Pointer(&subpass.PreserveAttachments), C.sizeof_uint32_t)
	}
	ucopy(unsafe.Pointer(dependencies), unsafe.Pointer(&info.Dependencies), C.sizeof_VkSubpassDependency)
	var out RenderPass
	res := Result(C.domVkCreateRenderPass(dev.fps[vkCreateRenderPass], dev.hnd, cinfo, nil, &out.hnd))
	return out, result2error(res)
}

func (dev *Device) DestroyRenderPass(renderPass RenderPass) {
	// TODO(dh): support a custom allocator
	C.domVkDestroyRenderPass(dev.fps[vkDestroyRenderPass], dev.hnd, renderPass.hnd, nil)
}

type FramebufferCreateInfo struct {
	Extensions  []Extension
	RenderPass  RenderPass
	Attachments []ImageView
	Width       uint32
	Height      uint32
	Layers      uint32
}

func (info *FramebufferCreateInfo) c(a *allocator) *C.VkFramebufferCreateInfo {
	cinfo := alloc[C.VkFramebufferCreateInfo](a)
	*cinfo = C.VkFramebufferCreateInfo{
		sType:           C.VkStructureType(StructureTypeFramebufferCreateInfo),
		pNext:           buildChain(a, info.Extensions),
		flags:           0,
		renderPass:      info.RenderPass.hnd,
		attachmentCount: C.uint32_t(len(info.Attachments)),
		pAttachments:    allocn[C.VkImageView](a, len(info.Attachments)),
		width:           C.uint32_t(info.Width),
		height:          C.uint32_t(info.Height),
		layers:          C.uint32_t(info.Layers),
	}
	ucopy(unsafe.Pointer(cinfo.pAttachments), unsafe.Pointer(&info.Attachments), C.sizeof_VkImageView)
	return cinfo
}

func (dev *Device) CreateFramebuffer(info *FramebufferCreateInfo) (Framebuffer, error) {
	a := new(allocator)
	defer a.free()

	cinfo := info.c(a)
	defer internalizeChain(info.Extensions, cinfo.pNext)
	var fb Framebuffer
	res := Result(C.domVkCreateFramebuffer(dev.fps[vkCreateFramebuffer], dev.hnd, cinfo, nil, &fb.hnd))
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
	// must be kept identical to C struct
	_ structs.HostLayout

	Depth   float32
	Stencil uint32
}

func (ClearColorValueFloat32s) isClearColorValue() {}
func (ClearColorValueInt32s) isClearColorValue()   {}
func (ClearColorValueUint32s) isClearColorValue()  {}

func (ClearColorValueFloat32s) isClearValue() {}
func (ClearColorValueInt32s) isClearValue()   {}
func (ClearColorValueUint32s) isClearValue()  {}
func (ClearDepthStencilValue) isClearValue()  {}

// Semaphores are a synchronization primitive that can be used to insert a dependency between batches submitted to queues.
// Semaphores have two states - signaled and unsignaled.
// The state of a semaphore can be signaled after execution of a batch of commands is completed.
// A batch can wait for a semaphore to become signaled before it begins execution, and the semaphore is also unsignaled before the batch begins execution.
//
// As with most objects in Vulkan, semaphores are an interface to internal data which is typically opaque to applications.
// This internal data is referred to as a semaphore’s payload.
//
// However, in order to enable communication with agents outside of the current device,
// it is necessary to be able to export that payload to a commonly understood format,
// and subsequently import from that format as well.
//
// The internal data of a semaphore may include a reference to any resources
// and pending work associated with signal or unsignal operations performed on that semaphore object.
// Mechanisms to import and export that internal data to and from semaphores exist.
// These mechanisms indirectly enable applications to share semaphore state
// between two or more semaphores and other synchronization primitives across process and API boundaries.
type Semaphore struct {
	// must be kept identical to C struct
	_ structs.HostLayout

	// VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkSemaphore)
	hnd C.VkSemaphore
}

type SemaphoreCreateInfo struct {
	Extensions []Extension
}

func (info *SemaphoreCreateInfo) c(a *allocator) *C.VkSemaphoreCreateInfo {
	cinfo := alloc[C.VkSemaphoreCreateInfo](a)
	*cinfo = C.VkSemaphoreCreateInfo{
		sType: C.VkStructureType(StructureTypeSemaphoreCreateInfo),
		pNext: buildChain(a, info.Extensions),
	}
	return cinfo
}

func (dev *Device) CreateSemaphore(info *SemaphoreCreateInfo) (Semaphore, error) {
	a := new(allocator)
	defer a.free()

	cinfo := info.c(a)
	defer internalizeChain(info.Extensions, cinfo.pNext)
	var sem Semaphore
	res := Result(C.domVkCreateSemaphore(dev.fps[vkCreateSemaphore], dev.hnd, cinfo, nil, &sem.hnd))
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
	a := new(allocator)
	defer a.free()

	cinfos := allocn[C.VkSubmitInfo](a, len(infos))
	cinfosArr := unsafe.Slice(cinfos, len(infos))

	for i, info := range infos {
		if safe && len(info.WaitSemaphores) != len(info.WaitDstStageMask) {
			panic("WaitSemaphores and WaitDstStageMask must have same length")
		}
		waitSemaphores := allocn[C.VkSemaphore](a, len(info.WaitSemaphores))
		waitDstStageMask := allocn[C.VkPipelineStageFlags](a, len(info.WaitDstStageMask))
		commandBuffers := allocn[C.VkCommandBuffer](a, len(info.CommandBuffers))
		signalSemaphores := allocn[C.VkSemaphore](a, len(info.SignalSemaphores))
		cinfosArr[i] = C.VkSubmitInfo{
			sType:                C.VkStructureType(StructureTypeSubmitInfo),
			pNext:                buildChain(a, info.Extensions),
			waitSemaphoreCount:   C.uint32_t(len(info.WaitSemaphores)),
			pWaitSemaphores:      (*C.VkSemaphore)(unsafe.Pointer(waitSemaphores)),
			pWaitDstStageMask:    (*C.VkPipelineStageFlags)(unsafe.Pointer(waitDstStageMask)),
			commandBufferCount:   C.uint32_t(len(info.CommandBuffers)),
			pCommandBuffers:      (*C.VkCommandBuffer)(unsafe.Pointer(commandBuffers)),
			signalSemaphoreCount: C.uint32_t(len(info.SignalSemaphores)),
			pSignalSemaphores:    (*C.VkSemaphore)(unsafe.Pointer(signalSemaphores)),
		}
		defer internalizeChain(info.Extensions, cinfosArr[i].pNext)
		ucopy(unsafe.Pointer(waitSemaphores), unsafe.Pointer(&info.WaitSemaphores), C.sizeof_VkSemaphore)
		ucopy(unsafe.Pointer(waitDstStageMask), unsafe.Pointer(&info.WaitDstStageMask), C.sizeof_VkPipelineStageFlags)
		ucopy(unsafe.Pointer(signalSemaphores), unsafe.Pointer(&info.SignalSemaphores), C.sizeof_VkSemaphore)
		arr := unsafe.Slice(commandBuffers, len(info.CommandBuffers))
		for i := range arr {
			arr[i] = info.CommandBuffers[i].hnd
		}
	}

	var fenceHnd C.VkFence
	if fence != nil {
		fenceHnd = fence.hnd
	}
	res := Result(C.domVkQueueSubmit(queue.fps[vkQueueSubmit], queue.hnd, C.uint32_t(len(infos)), cinfos, fenceHnd))
	return result2error(res)
}

// Fences are a synchronization primitive that can be used to insert a dependency from a queue to the host.
// Fences have two states - signaled and unsignaled.
// A fence can be signaled as part of the execution of a queue submission command.
// Fences can be unsignaled on the host with ResetFences.
// Fences can be waited on by the host with the WaitForFences command,
// and the current state can be queried with GetFenceStatus.
//
// As with most objects in Vulkan, fences are an interface to internal data
// which is typically opaque to applications.
// This internal data is referred to as a fence’s payload.
//
// However, in order to enable communication with agents outside of the current device,
// it is necessary to be able to export that payload to a commonly understood format,
// and subsequently import from that format as well.
//
// The internal data of a fence may include a reference to any resources
// and pending work associated with signal or unsignal operations performed on that fence object.
// Mechanisms to import and export that internal data to and from fences exist.
// These mechanisms indirectly enable applications to share fence state between
// two or more fences and other synchronization primitives across process and API boundaries.
type Fence struct {
	// must be kept identical to C struct
	_ structs.HostLayout

	// VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkFence)
	hnd C.VkFence
}

type FenceCreateInfo struct {
	Extensions []Extension
	Flags      FenceCreateFlags
}

func (dev *Device) CreateFence(info *FenceCreateInfo) (Fence, error) {
	a := new(allocator)
	defer a.free()

	cinfo := C.VkFenceCreateInfo{
		sType: C.VkStructureType(StructureTypeFenceCreateInfo),
		pNext: buildChain(a, info.Extensions),
		flags: C.VkFenceCreateFlags(info.Flags),
	}
	defer internalizeChain(info.Extensions, cinfo.pNext)
	var fence Fence
	res := Result(C.domVkCreateFence(dev.fps[vkCreateFence], dev.hnd, &cinfo, nil, &fence.hnd))
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
	res := Result(C.domVkWaitForFences(
		dev.fps[vkWaitForFences],
		dev.hnd,
		C.uint32_t(len(fences)),
		safeish.SliceCastPtr[*C.VkFence](fences),
		vkBool(waitAll),
		C.uint64_t(timeout)))
	return result2error(res)
}

func (dev *Device) ResetFences(fences []Fence) error {
	res := Result(C.domVkResetFences(dev.fps[vkResetFences], dev.hnd, C.uint32_t(len(fences)), safeish.SliceCastPtr[*C.VkFence](fences)))
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

// Buffers represent linear arrays of data which are used for various purposes
// by binding them to a graphics or compute pipeline
// via descriptor sets or via certain commands,
// or by directly specifying them as parameters to certain commands.
type Buffer struct {
	// must be kept identical to C struct
	_ structs.HostLayout

	// VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkBuffer)
	hnd C.VkBuffer
}

func (info *BufferCreateInfo) c(a *allocator) *C.VkBufferCreateInfo {
	cinfo := alloc[C.VkBufferCreateInfo](a)
	*cinfo = C.VkBufferCreateInfo{
		sType:                 C.VkStructureType(StructureTypeBufferCreateInfo),
		pNext:                 buildChain(a, info.Extensions),
		flags:                 C.VkBufferCreateFlags(info.Flags),
		size:                  C.VkDeviceSize(info.Size),
		usage:                 C.VkBufferUsageFlags(info.Usage),
		sharingMode:           C.VkSharingMode(info.SharingMode),
		queueFamilyIndexCount: C.uint32_t(len(info.QueueFamilyIndices)),
		pQueueFamilyIndices:   allocn[C.uint32_t](a, len(info.QueueFamilyIndices)),
	}
	ucopy(unsafe.Pointer(cinfo.pQueueFamilyIndices), unsafe.Pointer(&info.QueueFamilyIndices), C.sizeof_uint32_t)
	return cinfo
}

func (dev *Device) CreateBuffer(info *BufferCreateInfo) (Buffer, error) {
	a := new(allocator)
	defer a.free()

	cinfo := info.c(a)
	defer internalizeChain(info.Extensions, cinfo.pNext)
	var buf Buffer
	res := Result(C.domVkCreateBuffer(dev.fps[vkCreateBuffer], dev.hnd, cinfo, nil, &buf.hnd))
	return buf, result2error(res)
}

func (dev *Device) DestroyBuffer(buf Buffer) {
	// TODO(dh): support custom allocators
	C.domVkDestroyBuffer(dev.fps[vkDestroyBuffer], dev.hnd, buf.hnd, nil)
}

type MemoryRequirements struct {
	// must be kept identical to C struct
	_ structs.HostLayout

	Size           DeviceSize
	Alignment      DeviceSize
	MemoryTypeBits uint32
}

func (dev *Device) BufferMemoryRequirements(buf Buffer) MemoryRequirements {
	var reqs MemoryRequirements
	C.domVkGetBufferMemoryRequirements(dev.fps[vkGetBufferMemoryRequirements], dev.hnd, buf.hnd, (*C.VkMemoryRequirements)(unsafe.Pointer(&reqs)))
	return reqs
}

type BufferMemoryRequirementsInfo2 struct {
	Extensions []Extension
	Buffer     Buffer
}

func (info *BufferMemoryRequirementsInfo2) c(a *allocator) *C.VkBufferMemoryRequirementsInfo2 {
	cinfo := alloc[C.VkBufferMemoryRequirementsInfo2](a)
	*cinfo = C.VkBufferMemoryRequirementsInfo2{
		sType:  C.VkStructureType(StructureTypeBufferMemoryRequirementsInfo2),
		pNext:  buildChain(a, info.Extensions),
		buffer: info.Buffer.hnd,
	}
	return cinfo
}

func (dev *Device) BufferMemoryRequirements2(info *BufferMemoryRequirementsInfo2, reqs *MemoryRequirements2) {
	a := new(allocator)
	defer a.free()

	cinfo := info.c(a)
	defer internalizeChain(info.Extensions, cinfo.pNext)
	creqs := C.VkMemoryRequirements2{
		sType: C.VkStructureType(StructureTypeMemoryRequirements2),
		pNext: buildChain(a, reqs.Extensions),
	}
	defer internalizeChain(reqs.Extensions, creqs.pNext)
	C.domVkGetBufferMemoryRequirements2(dev.fps[vkGetBufferMemoryRequirements2], dev.hnd, cinfo, &creqs)
	ucopy1(unsafe.Pointer(&reqs.MemoryRequirements), unsafe.Pointer(&creqs.memoryRequirements), C.sizeof_VkMemoryRequirements)
}

type MemoryAllocateInfo struct {
	Extensions      []Extension
	AllocationSize  DeviceSize
	MemoryTypeIndex uint32
}

func (info *MemoryAllocateInfo) c(a *allocator) *C.VkMemoryAllocateInfo {
	cinfo := alloc[C.VkMemoryAllocateInfo](a)
	*cinfo = C.VkMemoryAllocateInfo{
		sType:           C.VkStructureType(StructureTypeMemoryAllocateInfo),
		pNext:           buildChain(a, info.Extensions),
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
	a := new(allocator)
	defer a.free()

	cinfo := info.c(a)
	defer internalizeChain(info.Extensions, cinfo.pNext)
	var mem DeviceMemory
	res := Result(C.domVkAllocateMemory(dev.fps[vkAllocateMemory], dev.hnd, cinfo, nil, &mem.hnd))
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
	a := new(allocator)
	defer a.free()

	mem := allocn[C.VkBindBufferMemoryInfo](a, len(infos))
	cinfos := unsafe.Slice(mem, len(infos))
	for i, info := range infos {
		cinfos[i] = C.VkBindBufferMemoryInfo{
			sType:        C.VkStructureType(StructureTypeBindBufferMemoryInfo),
			pNext:        buildChain(a, info.Extensions),
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
	var ptr unsafe.Pointer
	res := Result(C.domVkMapMemory(
		dev.fps[vkMapMemory],
		dev.hnd,
		mem.hnd,
		C.VkDeviceSize(offset),
		C.VkDeviceSize(size),
		C.VkMemoryMapFlags(flags),
		&ptr))
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

func (info *ImageCreateInfo) c(a *allocator) *C.VkImageCreateInfo {
	cinfo := alloc[C.VkImageCreateInfo](a)
	*cinfo = C.VkImageCreateInfo{
		sType:                 C.VkStructureType(StructureTypeImageCreateInfo),
		pNext:                 buildChain(a, info.Extensions),
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
		pQueueFamilyIndices:   allocn[C.uint32_t](a, len(info.QueueFamilyIndices)),
		initialLayout:         C.VkImageLayout(info.InitialLayout),
	}
	ucopy(unsafe.Pointer(cinfo.pQueueFamilyIndices), unsafe.Pointer(&info.QueueFamilyIndices), C.sizeof_uint32_t)
	ucopy1(unsafe.Pointer(&cinfo.extent), unsafe.Pointer(&info.Extent), C.sizeof_VkExtent3D)
	return cinfo
}

func (dev *Device) CreateImage(info *ImageCreateInfo) (Image, error) {
	a := new(allocator)
	defer a.free()

	cinfo := info.c(a)
	defer internalizeChain(info.Extensions, cinfo.pNext)
	var img Image
	res := Result(C.domVkCreateImage(dev.fps[vkCreateImage], dev.hnd, cinfo, nil, &img.hnd))
	return img, result2error(res)
}

func (dev *Device) DestroyImage(img Image) {
	// TODO(dh): support custom allocators
	C.domVkDestroyImage(dev.fps[vkDestroyImage], dev.hnd, img.hnd, nil)
}

// Events are a synchronization primitive
// that can be used to insert a fine-grained dependency between commands submitted to the same queue,
// or between the host and a queue.
// Events must not be used to insert a dependency between commands submitted to different queues.
// Events have two states - signaled and unsignaled.
// An application can signal an event, or unsignal it, on either the host or the device.
// A device can wait for an event to become signaled before executing further operations.
// No command exists to wait for an event to become signaled on the host,
// but the current state of an event can be queried.
type Event struct {
	// must be kept identical to C struct
	_ structs.HostLayout

	// VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkEvent)
	hnd C.VkEvent
}

type EventCreateInfo struct {
	Extensions []Extension
}

func (info *EventCreateInfo) c(a *allocator) *C.VkEventCreateInfo {
	cinfo := alloc[C.VkEventCreateInfo](a)
	*cinfo = C.VkEventCreateInfo{
		sType: C.VkStructureType(StructureTypeEventCreateInfo),
		pNext: buildChain(a, info.Extensions),
		flags: 0,
	}
	return cinfo
}

func (dev *Device) CreateEvent(info *EventCreateInfo) (Event, error) {
	a := new(allocator)
	defer a.free()

	cinfo := info.c(a)
	defer internalizeChain(info.Extensions, cinfo.pNext)
	var ev Event
	res := Result(C.domVkCreateEvent(dev.fps[vkCreateEvent], dev.hnd, cinfo, nil, &ev.hnd))
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

// Queries are managed using query pool objects.
// Each query pool is a collection of a specific number of queries of a particular type.
type QueryPool struct {
	// VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkQueryPool)
	hnd C.VkQueryPool
}

type QueryPoolCreateInfo struct {
	Extensions         []Extension
	QueryType          QueryType
	QueryCount         uint32
	PipelineStatistics QueryPipelineStatisticFlags
}

func (info *QueryPoolCreateInfo) c(a *allocator) *C.VkQueryPoolCreateInfo {
	cinfo := alloc[C.VkQueryPoolCreateInfo](a)
	*cinfo = C.VkQueryPoolCreateInfo{
		sType:              C.VkStructureType(StructureTypeQueryPoolCreateInfo),
		pNext:              buildChain(a, info.Extensions),
		queryType:          C.VkQueryType(info.QueryType),
		queryCount:         C.uint32_t(info.QueryCount),
		pipelineStatistics: C.VkQueryPipelineStatisticFlags(info.PipelineStatistics),
	}
	return cinfo
}

func (dev *Device) CreateQueryPool(info *QueryPoolCreateInfo) (QueryPool, error) {
	a := new(allocator)
	defer a.free()

	cinfo := info.c(a)
	defer internalizeChain(info.Extensions, cinfo.pNext)
	var out QueryPool
	res := Result(C.domVkCreateQueryPool(dev.fps[vkCreateQueryPool], dev.hnd, cinfo, nil, &out.hnd))
	return out, result2error(res)
}

func (dev *Device) DestroyQueryPool(queryPool QueryPool) {
	// TODO(dh(): support custom allocator
	C.domVkDestroyQueryPool(dev.fps[vkDestroyQueryPool], dev.hnd, queryPool.hnd, nil)
}

// Sampler objects represent the state of an image sampler
// which is used by the implementation to read image data
// and apply filtering and other transformations for the shader.
type Sampler struct {
	// must be kept identical to C struct
	_ structs.HostLayout

	// VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkSampler)
	hnd C.VkSampler
}

type SamplerCreateInfo struct {
	Extensions              []Extension
	MagFilter               Filter
	MinFilter               Filter
	MipmapMode              SamplerMipmapMode
	AddressModeU            SamplerAddressMode
	AddressModeV            SamplerAddressMode
	AddressModeW            SamplerAddressMode
	MipLodBias              float32
	AnisotropyEnable        bool
	MaxAnisotropy           float32
	CompareEnable           bool
	CompareOp               CompareOp
	MinLod                  float32
	MaxLod                  float32
	BorderColor             BorderColor
	UnnormalizedCoordinates bool
}

func (info *SamplerCreateInfo) c(a *allocator) *C.VkSamplerCreateInfo {
	cinfo := alloc[C.VkSamplerCreateInfo](a)
	*cinfo = C.VkSamplerCreateInfo{
		sType:                   C.VkStructureType(StructureTypeSamplerCreateInfo),
		pNext:                   buildChain(a, info.Extensions),
		magFilter:               C.VkFilter(info.MagFilter),
		minFilter:               C.VkFilter(info.MinFilter),
		mipmapMode:              C.VkSamplerMipmapMode(info.MipmapMode),
		addressModeU:            C.VkSamplerAddressMode(info.AddressModeU),
		addressModeV:            C.VkSamplerAddressMode(info.AddressModeV),
		addressModeW:            C.VkSamplerAddressMode(info.AddressModeW),
		mipLodBias:              C.float(info.MipLodBias),
		anisotropyEnable:        vkBool(info.AnisotropyEnable),
		maxAnisotropy:           C.float(info.MaxAnisotropy),
		compareEnable:           vkBool(info.CompareEnable),
		compareOp:               C.VkCompareOp(info.CompareOp),
		minLod:                  C.float(info.MinLod),
		maxLod:                  C.float(info.MaxLod),
		borderColor:             C.VkBorderColor(info.BorderColor),
		unnormalizedCoordinates: vkBool(info.UnnormalizedCoordinates),
	}
	return cinfo
}

func (dev *Device) CreateSampler(info *SamplerCreateInfo) (Sampler, error) {
	a := new(allocator)
	defer a.free()

	cinfo := info.c(a)
	defer internalizeChain(info.Extensions, cinfo.pNext)
	var out Sampler
	res := Result(C.domVkCreateSampler(dev.fps[vkCreateSampler], dev.hnd, cinfo, nil, &out.hnd))
	return out, result2error(res)
}

func (dev *Device) DestroySampler(sampler Sampler) {
	// TODO(dh): support custom allocators
	C.domVkDestroySampler(dev.fps[vkDestroySampler], dev.hnd, sampler.hnd, nil)
}

// A buffer view represents a contiguous range of a buffer and a specific format to be used to interpret the data.
// Buffer views are used to enable shaders to access buffer contents interpreted as formatted data.
// In order to create a valid buffer view, the buffer must have been created with at least one of the following usage flags:
//   - BufferUsageUniformTexelBufferBit
//   - BufferUsageStorageTexelBufferBit
type BufferView struct {
	// VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkBufferView)
	hnd C.VkBufferView
}

type BufferViewCreateInfo struct {
	Extensions []Extension
	Buffer     Buffer
	Format     Format
	Offset     DeviceSize
	Range      DeviceSize
}

func (info *BufferViewCreateInfo) c(a *allocator) *C.VkBufferViewCreateInfo {
	cinfo := alloc[C.VkBufferViewCreateInfo](a)
	*cinfo = C.VkBufferViewCreateInfo{
		sType:  C.VkStructureType(StructureTypeBufferViewCreateInfo),
		pNext:  buildChain(a, info.Extensions),
		flags:  0,
		buffer: info.Buffer.hnd,
		format: C.VkFormat(info.Format),
		offset: C.VkDeviceSize(info.Offset),
		_range: C.VkDeviceSize(info.Range),
	}
	return cinfo
}

func (dev *Device) CreateBufferView(info *BufferViewCreateInfo) (BufferView, error) {
	a := new(allocator)
	defer a.free()

	cinfo := info.c(a)
	defer internalizeChain(info.Extensions, cinfo.pNext)
	var view BufferView
	res := Result(C.domVkCreateBufferView(dev.fps[vkCreateBufferView], dev.hnd, cinfo, nil, &view.hnd))
	return view, result2error(res)
}

func (dev *Device) DestroyBufferView(view BufferView) {
	// TODO(dh): support custom allocators
	C.domVkDestroyBufferView(dev.fps[vkDestroyBufferView], dev.hnd, view.hnd, nil)
}

// Pipeline cache objects allow the result of pipeline construction to be reused between pipelines and between runs of an application.
// Reuse between pipelines is achieved by passing the same pipeline cache object when creating multiple related pipelines.
// Reuse across runs of an application is achieved by retrieving pipeline cache contents in one run of an application,
// saving the contents, and using them to preinitialize a pipeline cache on a subsequent run.
// The contents of the pipeline cache objects are managed by the implementation.
// Applications can manage the host memory consumed by a pipeline cache object and control the amount of data retrieved from a pipeline cache object.
type PipelineCache struct {
	// must be kept identical to C struct
	_ structs.HostLayout

	// VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkPipelineCache)
	hnd C.VkPipelineCache
}

type PipelineCacheCreateInfo struct {
	Extensions  []Extension
	InitialData []byte
}

func (info *PipelineCacheCreateInfo) c(a *allocator) *C.VkPipelineCacheCreateInfo {
	cinfo := alloc[C.VkPipelineCacheCreateInfo](a)
	*cinfo = C.VkPipelineCacheCreateInfo{
		sType:           C.VkStructureType(StructureTypePipelineCacheCreateInfo),
		pNext:           buildChain(a, info.Extensions),
		flags:           0,
		initialDataSize: C.size_t(len(info.InitialData)),
		pInitialData:    unsafe.Pointer(unsafe.SliceData(info.InitialData)),
	}
	return cinfo
}

func (dev *Device) CreatePipelineCache(info *PipelineCacheCreateInfo) (PipelineCache, error) {
	a := new(allocator)
	defer a.free()

	cinfo := info.c(a)
	defer internalizeChain(info.Extensions, cinfo.pNext)
	var out PipelineCache
	res := Result(C.domVkCreatePipelineCache(dev.fps[vkCreatePipelineCache], dev.hnd, cinfo, nil, &out.hnd))
	return out, result2error(res)
}

func (dev *Device) DestroyPipelineCache(cache PipelineCache) {
	// TODO(dh): support custom allocators
	C.domVkDestroyPipelineCache(dev.fps[vkDestroyPipelineCache], dev.hnd, cache.hnd, nil)
}

func (dev *Device) MergePipelineCaches(dstCache PipelineCache, srcCaches []PipelineCache) error {
	res := Result(C.domVkMergePipelineCaches(
		dev.fps[vkMergePipelineCaches],
		dev.hnd,
		dstCache.hnd,
		C.uint32_t(len(srcCaches)),
		safeish.SliceCastPtr[*C.VkPipelineCache](srcCaches)))
	return result2error(res)
}

func (dev *Device) PipelineCacheData(cache PipelineCache) ([]byte, error) {
	var size C.size_t
	var data []byte
	for {
		res := Result(C.domVkGetPipelineCacheData(dev.fps[vkGetPipelineCacheData], dev.hnd, cache.hnd, &size, nil))
		if res != Success {
			return nil, res
		}
		data = make([]byte, size)
		res = Result(C.domVkGetPipelineCacheData(dev.fps[vkGetPipelineCacheData], dev.hnd, cache.hnd, &size, unsafe.Pointer(unsafe.SliceData(data))))
		if res == Success {
			return data[:size], nil
		}
		if res == Incomplete {
			continue
		}
		return nil, res
	}
}

type DescriptorPoolSize struct {
	// must be kept identical to C struct
	_ structs.HostLayout

	Type            DescriptorType
	DescriptorCount uint32
}

type DescriptorPoolCreateInfo struct {
	Extensions []Extension
	Flags      DescriptorPoolCreateFlags
	MaxSets    uint32
	PoolSizes  []DescriptorPoolSize
}

func (info *DescriptorPoolCreateInfo) c(a *allocator) *C.VkDescriptorPoolCreateInfo {
	cinfo := alloc[C.VkDescriptorPoolCreateInfo](a)
	*cinfo = C.VkDescriptorPoolCreateInfo{
		sType:         C.VkStructureType(StructureTypeDescriptorPoolCreateInfo),
		pNext:         buildChain(a, info.Extensions),
		flags:         C.VkDescriptorPoolCreateFlags(info.Flags),
		maxSets:       C.uint32_t(info.MaxSets),
		poolSizeCount: C.uint32_t(len(info.PoolSizes)),
		pPoolSizes:    allocn[C.VkDescriptorPoolSize](a, len(info.PoolSizes)),
	}
	ucopy(unsafe.Pointer(cinfo.pPoolSizes), unsafe.Pointer(&info.PoolSizes), C.sizeof_VkDescriptorPoolSize)
	return cinfo
}

// A descriptor pool maintains a pool of descriptors, from which descriptor sets are allocated.
// Descriptor pools are externally synchronized, meaning that the application must not
// allocate and/or free descriptor sets from the same pool in multiple threads simultaneously.
type DescriptorPool struct {
	// VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkDescriptorPool)
	hnd C.VkDescriptorPool
}

func (dev *Device) CreateDescriptorPool(info *DescriptorPoolCreateInfo) (DescriptorPool, error) {
	a := new(allocator)
	defer a.free()

	cinfo := info.c(a)
	defer internalizeChain(info.Extensions, cinfo.pNext)
	var out DescriptorPool
	res := Result(C.domVkCreateDescriptorPool(dev.fps[vkCreateDescriptorPool], dev.hnd, cinfo, nil, &out.hnd))
	return out, result2error(res)
}

func (dev *Device) DestroyDescriptorPool(pool DescriptorPool) {
	// TODO(dh): support custom allocators
	C.domVkDestroyDescriptorPool(dev.fps[vkDestroyDescriptorPool], dev.hnd, pool.hnd, nil)
}

func (dev *Device) ResetDescriptorPool(pool DescriptorPool, flags DescriptorPoolResetFlags) error {
	res := Result(C.domVkResetDescriptorPool(dev.fps[vkResetDescriptorPool], dev.hnd, pool.hnd, C.VkDescriptorPoolResetFlags(flags)))
	return result2error(res)
}

type DescriptorSetLayoutBinding struct {
	Binding           uint32
	DescriptorType    DescriptorType
	DescriptorCount   uint32
	StageFlags        ShaderStageFlags
	ImmutableSamplers []Sampler
}

type DescriptorSetLayoutCreateInfo struct {
	Extensions []Extension
	Flags      DescriptorSetLayoutCreateFlags
	Bindings   []DescriptorSetLayoutBinding
}

func (info *DescriptorSetLayoutCreateInfo) c(a *allocator) *C.VkDescriptorSetLayoutCreateInfo {
	cinfo := alloc[C.VkDescriptorSetLayoutCreateInfo](a)
	*cinfo = C.VkDescriptorSetLayoutCreateInfo{
		sType:        C.VkStructureType(StructureTypeDescriptorSetLayoutCreateInfo),
		pNext:        buildChain(a, info.Extensions),
		flags:        C.VkDescriptorSetLayoutCreateFlags(info.Flags),
		bindingCount: C.uint32_t(len(info.Bindings)),
		pBindings:    allocn[C.VkDescriptorSetLayoutBinding](a, len(info.Bindings)),
	}
	arr := unsafe.Slice(cinfo.pBindings, len(info.Bindings))
	for i := range arr {
		samplers := allocn[C.VkSampler](a, len(info.Bindings[i].ImmutableSamplers))
		arr[i] = C.VkDescriptorSetLayoutBinding{
			binding:            C.uint32_t(info.Bindings[i].Binding),
			descriptorType:     C.VkDescriptorType(info.Bindings[i].DescriptorType),
			descriptorCount:    C.uint32_t(info.Bindings[i].DescriptorCount),
			stageFlags:         C.VkShaderStageFlags(info.Bindings[i].StageFlags),
			pImmutableSamplers: (*C.VkSampler)(samplers),
		}
		ucopy(unsafe.Pointer(samplers), unsafe.Pointer(&info.Bindings[i].ImmutableSamplers), C.sizeof_VkSampler)
	}
	return cinfo
}

func (dev *Device) CreateDescriptorSetLayout(info *DescriptorSetLayoutCreateInfo) (DescriptorSetLayout, error) {
	a := new(allocator)
	defer a.free()

	cinfo := info.c(a)
	defer internalizeChain(info.Extensions, cinfo.pNext)
	var out DescriptorSetLayout
	res := Result(C.domVkCreateDescriptorSetLayout(dev.fps[vkCreateDescriptorSetLayout], dev.hnd, cinfo, nil, &out.hnd))
	return out, result2error(res)
}

func (dev *Device) DestroyDescriptorSetLayout(layout DescriptorSetLayout) {
	// TODO(dh): support custom allocators
	C.domVkDestroyDescriptorSetLayout(dev.fps[vkDestroyDescriptorSetLayout], dev.hnd, layout.hnd, nil)
}

type DescriptorSetLayoutSupport struct {
	Extensions []Extension
	Supported  bool
}

func (dev *Device) DescriptorSetLayoutSupport(info DescriptorSetLayoutCreateInfo, support *DescriptorSetLayoutSupport) bool {
	// XXX audit this function, it seems to be missing an actual call
	a := new(allocator)
	defer a.free()

	cinfo := info.c(a)
	defer internalizeChain(info.Extensions, cinfo.pNext)
	csupport := alloc[C.VkDescriptorSetLayoutSupport](a)
	csupport.sType = C.VkStructureType(StructureTypeDescriptorSetLayoutSupport)
	if support != nil {
		csupport.pNext = buildChain(a, support.Extensions)
		defer internalizeChain(support.Extensions, csupport.pNext)
	}
	if support != nil {
		support.Supported = csupport.supported == C.VK_TRUE
	}
	return csupport.supported == C.VK_TRUE
}

func (dev *Device) BindImageMemory(image Image, memory DeviceMemory, offset DeviceSize) error {
	res := Result(C.domVkBindImageMemory(dev.fps[vkBindImageMemory], dev.hnd, image.hnd, memory.hnd, C.VkDeviceSize(offset)))
	return result2error(res)
}

type BindImageMemoryInfo struct {
	Extensions   []Extension
	Image        Image
	Memory       DeviceMemory
	MemoryOffset DeviceSize
}

func (dev *Device) BindImageMemory2(infos []BindImageMemoryInfo) error {
	a := new(allocator)
	defer a.free()

	mem := allocn[C.VkBindImageMemoryInfo](a, len(infos))
	arr := unsafe.Slice(mem, len(infos))
	for i := range arr {
		arr[i] = C.VkBindImageMemoryInfo{
			sType:        C.VkStructureType(StructureTypeBindImageMemoryInfo),
			pNext:        buildChain(a, infos[i].Extensions),
			image:        infos[i].Image.hnd,
			memory:       infos[i].Memory.hnd,
			memoryOffset: C.VkDeviceSize(infos[i].MemoryOffset),
		}
	}
	res := Result(C.domVkBindImageMemory2(dev.fps[vkBindImageMemory2], dev.hnd, C.uint32_t(len(infos)), (*C.VkBindImageMemoryInfo)(mem)))
	for i := range arr {
		internalizeChain(infos[i].Extensions, arr[i].pNext)
	}
	return result2error(res)
}

func (dev *Device) ImageMemoryRequirements(image Image) MemoryRequirements {
	var out MemoryRequirements
	C.domVkGetImageMemoryRequirements(dev.fps[vkGetImageMemoryRequirements], dev.hnd, image.hnd, (*C.VkMemoryRequirements)(unsafe.Pointer(&out)))
	return out
}

type ImageMemoryRequirementsInfo2 struct {
	Extensions []Extension
	Image      Image
}

func (info *ImageMemoryRequirementsInfo2) c(a *allocator) *C.VkImageMemoryRequirementsInfo2 {
	cinfo := alloc[C.VkImageMemoryRequirementsInfo2](a)
	*cinfo = C.VkImageMemoryRequirementsInfo2{
		sType: C.VkStructureType(StructureTypeImageMemoryRequirementsInfo2),
		pNext: buildChain(a, info.Extensions),
		image: info.Image.hnd,
	}
	return cinfo
}

type MemoryRequirements2 struct {
	Extensions         []Extension
	MemoryRequirements MemoryRequirements
}

func (dev *Device) ImageMemoryRequirements2(info *ImageMemoryRequirementsInfo2, reqs *MemoryRequirements2) {
	a := new(allocator)
	defer a.free()

	cinfo := info.c(a)
	defer internalizeChain(info.Extensions, cinfo.pNext)
	creqs := C.VkMemoryRequirements2{
		sType: C.VkStructureType(StructureTypeMemoryRequirements2),
		pNext: buildChain(a, reqs.Extensions),
	}
	defer internalizeChain(reqs.Extensions, creqs.pNext)
	C.domVkGetImageMemoryRequirements2(dev.fps[vkGetImageMemoryRequirements2], dev.hnd, cinfo, &creqs)
	ucopy1(unsafe.Pointer(&reqs.MemoryRequirements), unsafe.Pointer(&creqs.memoryRequirements), C.sizeof_VkMemoryRequirements)
}

type DescriptorSetAllocateInfo struct {
	Extensions     []Extension
	DescriptorPool DescriptorPool
	Layouts        []DescriptorSetLayout
}

func (info *DescriptorSetAllocateInfo) c(a *allocator) *C.VkDescriptorSetAllocateInfo {
	cinfo := alloc[C.VkDescriptorSetAllocateInfo](a)
	*cinfo = C.VkDescriptorSetAllocateInfo{
		sType:              C.VkStructureType(StructureTypeDescriptorSetAllocateInfo),
		pNext:              buildChain(a, info.Extensions),
		descriptorPool:     info.DescriptorPool.hnd,
		descriptorSetCount: C.uint32_t(len(info.Layouts)),
		pSetLayouts:        allocn[C.VkDescriptorSetLayout](a, len(info.Layouts)),
	}
	ucopy(unsafe.Pointer(cinfo.pSetLayouts), unsafe.Pointer(&info.Layouts), C.sizeof_VkDescriptorSetLayout)
	return cinfo
}

type DescriptorSet struct {
	// must be kept identical to C struct
	_ structs.HostLayout

	// VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkDescriptorSet)
	hnd C.VkDescriptorSet
}

func (dev *Device) AllocateDescriptorSets(info DescriptorSetAllocateInfo) ([]DescriptorSet, error) {
	a := new(allocator)
	defer a.free()

	cinfo := info.c(a)
	defer internalizeChain(info.Extensions, cinfo.pNext)
	out := make([]DescriptorSet, len(info.Layouts))
	res := Result(C.domVkAllocateDescriptorSets(dev.fps[vkAllocateDescriptorSets], dev.hnd, cinfo, safeish.SliceCastPtr[*C.VkDescriptorSet](out)))
	return out, result2error(res)
}

func (dev *Device) FreeDescriptorSets(pool DescriptorPool, sets []DescriptorSet) error {
	res := Result(C.domVkFreeDescriptorSets(dev.fps[vkFreeDescriptorSets], dev.hnd, pool.hnd, C.uint32_t(len(sets)), safeish.SliceCastPtr[*C.VkDescriptorSet](sets)))
	return result2error(res)
}

func (dev *Device) QueryPoolResults(
	pool QueryPool,
	firstQuery uint32,
	queryCount uint32,
	data []byte,
	stride DeviceSize,
	flags QueryResultFlags,
) error {
	res := Result(C.domVkGetQueryPoolResults(
		dev.fps[vkGetQueryPoolResults],
		dev.hnd,
		pool.hnd,
		C.uint32_t(firstQuery),
		C.uint32_t(queryCount),
		C.size_t(len(data)),
		unsafe.Pointer(unsafe.SliceData(data)),
		C.VkDeviceSize(stride),
		C.VkQueryResultFlags(flags)))
	return result2error(res)
}

func (dev *Device) RenderAreaGranularity(renderPass RenderPass) Extent2D {
	var out Extent2D
	C.domVkGetRenderAreaGranularity(dev.fps[vkGetRenderAreaGranularity], dev.hnd, renderPass.hnd, (*C.VkExtent2D)(unsafe.Pointer(&out)))
	return out
}

type WriteDescriptorSet struct {
	Extensions      []Extension
	DstSet          DescriptorSet
	DstBinding      uint32
	DstArrayElement uint32
	DescriptorType  DescriptorType
	ImageInfo       []DescriptorImageInfo
	BufferInfo      []DescriptorBufferInfo
	TexelBufferView []BufferView
}

func (set *WriteDescriptorSet) c(a *allocator, cset *C.VkWriteDescriptorSet) {
	if safe {
		n := 0
		if set.ImageInfo != nil {
			n++
		}
		if set.BufferInfo != nil {
			n++
		}
		if set.TexelBufferView != nil {
			n++
		}
		if n > 1 {
			panic("only one of ImageInfo, BufferInfo, or TexelBufferView must be provided")
		}
	}

	// We trust the user that only one of ImageInfo, BufferInfo, or
	// TexelBufferView has been provided. If that invariant is broken,
	// and safety checks are disable, invalid memory may be read.

	*cset = C.VkWriteDescriptorSet{
		sType:            C.VkStructureType(StructureTypeWriteDescriptorSet),
		pNext:            buildChain(a, set.Extensions),
		dstSet:           set.DstSet.hnd,
		dstBinding:       C.uint32_t(set.DstBinding),
		dstArrayElement:  C.uint32_t(set.DstArrayElement),
		descriptorCount:  C.uint32_t(len(set.ImageInfo) + len(set.BufferInfo) + len(set.TexelBufferView)),
		descriptorType:   C.VkDescriptorType(set.DescriptorType),
		pImageInfo:       allocn[C.VkDescriptorImageInfo](a, len(set.ImageInfo)),
		pBufferInfo:      allocn[C.VkDescriptorBufferInfo](a, len(set.BufferInfo)),
		pTexelBufferView: allocn[C.VkBufferView](a, len(set.TexelBufferView)),
	}
	ucopy(unsafe.Pointer(cset.pImageInfo), unsafe.Pointer(&set.ImageInfo), C.sizeof_VkDescriptorImageInfo)
	ucopy(unsafe.Pointer(cset.pBufferInfo), unsafe.Pointer(&set.BufferInfo), C.sizeof_VkDescriptorBufferInfo)
	ucopy(unsafe.Pointer(cset.pTexelBufferView), unsafe.Pointer(&set.TexelBufferView), C.sizeof_VkBufferView)
}

type DescriptorBufferInfo struct {
	// must be kept identical to C struct
	_ structs.HostLayout

	Buffer Buffer
	Offset DeviceSize
	Range  DeviceSize
}

type DescriptorImageInfo struct {
	// must be kept identical to C struct
	_ structs.HostLayout

	Sampler     Sampler
	ImageView   ImageView
	ImageLayout ImageLayout
}

type CopyDescriptorSet struct {
	Extensions      []Extension
	SrcSet          DescriptorSet
	SrcBinding      uint32
	SrcArrayElement uint32
	DstSet          DescriptorSet
	DstBinding      uint32
	DstArrayElement uint32
	DescriptorCount uint32
}

func (set *CopyDescriptorSet) c(a *allocator, cset *C.VkCopyDescriptorSet) {
	*cset = C.VkCopyDescriptorSet{
		sType:           C.VkStructureType(StructureTypeCopyDescriptorSet),
		pNext:           buildChain(a, set.Extensions),
		srcSet:          set.SrcSet.hnd,
		srcBinding:      C.uint32_t(set.SrcBinding),
		srcArrayElement: C.uint32_t(set.SrcArrayElement),
		dstSet:          set.DstSet.hnd,
		dstBinding:      C.uint32_t(set.DstBinding),
		dstArrayElement: C.uint32_t(set.DstArrayElement),
		descriptorCount: C.uint32_t(set.DescriptorCount),
	}
}

func (dev *Device) UpdateDescriptorSets(writes []WriteDescriptorSet, copies []CopyDescriptorSet) {
	a := new(allocator)
	defer a.free()

	cwrites := make([]C.VkWriteDescriptorSet, len(writes))
	ccopies := make([]C.VkCopyDescriptorSet, len(copies))
	for i := range cwrites {
		writes[i].c(a, &cwrites[i])
	}
	for i := range ccopies {
		copies[i].c(a, &ccopies[i])
	}
	C.domVkUpdateDescriptorSets(
		dev.fps[vkUpdateDescriptorSets],
		dev.hnd,
		C.uint32_t(len(cwrites)),
		safeish.SliceCastPtr[*C.VkWriteDescriptorSet](cwrites),
		C.uint32_t(len(ccopies)),
		safeish.SliceCastPtr[*C.VkCopyDescriptorSet](ccopies))
	for i := range cwrites {
		internalizeChain(writes[i].Extensions, cwrites[i].pNext)
	}
	for i := range ccopies {
		internalizeChain(copies[i].Extensions, ccopies[i].pNext)
	}
}

type MappedMemoryRange struct {
	Extensions []Extension
	Memory     DeviceMemory
	Offset     DeviceSize
	Size       DeviceSize
}

func (rng *MappedMemoryRange) c(a *allocator, crng *C.VkMappedMemoryRange) {
	*crng = C.VkMappedMemoryRange{
		sType:  C.VkStructureType(StructureTypeMappedMemoryRange),
		pNext:  buildChain(a, rng.Extensions),
		memory: rng.Memory.hnd,
		offset: C.VkDeviceSize(rng.Offset),
		size:   C.VkDeviceSize(rng.Size),
	}
}

func (dev *Device) FlushMappedMemoryRanges(ranges []MappedMemoryRange) error {
	a := new(allocator)
	defer a.free()

	cranges := make([]C.VkMappedMemoryRange, len(ranges))
	for i := range cranges {
		ranges[i].c(a, &cranges[i])
	}
	res := Result(C.domVkFlushMappedMemoryRanges(dev.fps[vkFlushMappedMemoryRanges], dev.hnd, C.uint32_t(len(cranges)), unsafe.SliceData(cranges)))
	for i := range cranges {
		internalizeChain(ranges[i].Extensions, cranges[i].pNext)
	}
	return result2error(res)
}

type FormatProperties struct {
	// must be kept identical to C struct
	_ structs.HostLayout

	LinearTilingFeatures  FormatFeatureFlags
	OptimalTilingFeatures FormatFeatureFlags
	BufferFeatures        FormatFeatureFlags
}

func (dev *PhysicalDevice) FormatProperties(format Format) FormatProperties {
	var props FormatProperties
	C.domVkGetPhysicalDeviceFormatProperties(dev.instance.fps[vkGetPhysicalDeviceFormatProperties], dev.hnd, C.VkFormat(format), (*C.VkFormatProperties)(unsafe.Pointer(&props)))
	return props
}

type FormatProperties2 struct {
	Extensions       []Extension
	FormatProperties FormatProperties
}

func (dev *PhysicalDevice) FormatProperties2(format Format, props *FormatProperties2) {
	a := new(allocator)
	defer a.free()

	cprops := C.VkFormatProperties2{
		sType: C.VkStructureType(StructureTypeFormatProperties2),
		pNext: buildChain(a, props.Extensions),
	}
	defer internalizeChain(props.Extensions, cprops.pNext)
	C.domVkGetPhysicalDeviceFormatProperties2(dev.instance.fps[vkGetPhysicalDeviceFormatProperties2], dev.hnd, C.VkFormat(format), &cprops)
	ucopy1(unsafe.Pointer(&props.FormatProperties), unsafe.Pointer(&cprops.formatProperties), C.sizeof_VkFormatProperties)
}

type ImageFormatProperties struct {
	// must be kept identical to C struct
	_ structs.HostLayout

	MaxExtent       Extent3D
	MaxMipLevels    uint32
	MaxArrayLayers  uint32
	SampleCounts    SampleCountFlags
	MaxResourceSize DeviceSize
}

func (dev *PhysicalDevice) ImageFormatProperties(
	format Format,
	typ ImageType,
	tiling ImageTiling,
	usage ImageUsageFlags,
	flags ImageCreateFlags,
) (ImageFormatProperties, error) {
	var props ImageFormatProperties
	res := Result(C.domVkGetPhysicalDeviceImageFormatProperties(
		dev.instance.fps[vkGetPhysicalDeviceImageFormatProperties],
		dev.hnd,
		C.VkFormat(format),
		C.VkImageType(typ),
		C.VkImageTiling(tiling),
		C.VkImageUsageFlags(usage),
		C.VkImageCreateFlags(flags),
		(*C.VkImageFormatProperties)(unsafe.Pointer(&props))))
	return props, result2error(res)
}

type PhysicalDeviceImageFormatInfo2 struct {
	Extensions []Extension
	Format     Format
	Type       ImageType
	Tiling     ImageTiling
	Usage      ImageUsageFlags
	Flags      ImageCreateFlags
}

func (info *PhysicalDeviceImageFormatInfo2) c(a *allocator) *C.VkPhysicalDeviceImageFormatInfo2 {
	cinfo := alloc[C.VkPhysicalDeviceImageFormatInfo2](a)
	*cinfo = C.VkPhysicalDeviceImageFormatInfo2{
		sType:  C.VkStructureType(StructureTypePhysicalDeviceImageFormatInfo2),
		pNext:  buildChain(a, info.Extensions),
		format: C.VkFormat(info.Format),
		_type:  C.VkImageType(info.Type),
		tiling: C.VkImageTiling(info.Tiling),
		usage:  C.VkImageUsageFlags(info.Usage),
		flags:  C.VkImageCreateFlags(info.Flags),
	}
	return cinfo
}

type ImageFormatProperties2 struct {
	Extensions            []Extension
	ImageFormatProperties ImageFormatProperties
}

func (dev *PhysicalDevice) ImageFormatProperties2(info *PhysicalDeviceImageFormatInfo2, props *ImageFormatProperties2) error {
	a := new(allocator)
	defer a.free()

	cinfo := info.c(a)
	defer internalizeChain(info.Extensions, cinfo.pNext)
	cprops := C.VkImageFormatProperties2{
		sType: C.VkStructureType(StructureTypeImageFormatProperties2),
		pNext: buildChain(a, props.Extensions),
	}
	defer internalizeChain(props.Extensions, cprops.pNext)
	res := Result(C.domVkGetPhysicalDeviceImageFormatProperties2(dev.instance.fps[vkGetPhysicalDeviceImageFormatProperties2], dev.hnd, cinfo, &cprops))
	ucopy1(unsafe.Pointer(&props.ImageFormatProperties), unsafe.Pointer(&cprops.imageFormatProperties), C.sizeof_VkImageFormatProperties)
	return result2error(res)
}

func vkGetInstanceProcAddr(instance C.VkInstance, name string) C.PFN_vkVoidFunction {
	// TODO(dh): return a mock function pointer that panics with a nice message

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	fp := C.vkGetInstanceProcAddr(instance, cName)
	if debug {
		fmt.Fprintf(os.Stderr, "%s = %p\n", name, fp)
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

func (hnd *CommandBuffer) String() string  { return fmt.Sprintf("VkCommandBuffer(%#x)", hnd.hnd) }
func (hnd *Device) String() string         { return fmt.Sprintf("VkDevice(%#x)", hnd.hnd) }
func (hnd *Instance) String() string       { return fmt.Sprintf("VkInstance(%#x)", hnd.hnd) }
func (hnd *PhysicalDevice) String() string { return fmt.Sprintf("VkPhysicalDevice(%#x)", hnd.hnd) }
func (hnd *Queue) String() string          { return fmt.Sprintf("VkQueue(%#x)", hnd.hnd) }
func (hnd Buffer) String() string          { return fmt.Sprintf("VkBuffer(%#x)", hnd.hnd) }
func (hnd BufferView) String() string      { return fmt.Sprintf("VkBufferView(%#x)", hnd.hnd) }
func (hnd CommandPool) String() string     { return fmt.Sprintf("VkCommandPool(%#x)", hnd.hnd) }
func (hnd DescriptorSet) String() string   { return fmt.Sprintf("VkDescriptorSet(%#x)", hnd.hnd) }
func (hnd DeviceMemory) String() string    { return fmt.Sprintf("VkDeviceMemory(%#x)", hnd.hnd) }
func (hnd Event) String() string           { return fmt.Sprintf("VkEvent(%#x)", hnd.hnd) }
func (hnd Fence) String() string           { return fmt.Sprintf("VkFence(%#x)", hnd.hnd) }
func (hnd Framebuffer) String() string     { return fmt.Sprintf("VkFramebuffer(%#x)", hnd.hnd) }
func (hnd Image) String() string           { return fmt.Sprintf("VkImage(%#x)", hnd.hnd) }
func (hnd ImageView) String() string       { return fmt.Sprintf("VkImageView(%#x)", hnd.hnd) }
func (hnd Pipeline) String() string        { return fmt.Sprintf("VkPipeline(%#x)", hnd.hnd) }
func (hnd PipelineCache) String() string   { return fmt.Sprintf("VkPipelineCache(%#x)", hnd.hnd) }
func (hnd PipelineLayout) String() string  { return fmt.Sprintf("VkPipelineLayout(%#x)", hnd.hnd) }
func (hnd QueryPool) String() string       { return fmt.Sprintf("VkQueryPool(%#x)", hnd.hnd) }
func (hnd RenderPass) String() string      { return fmt.Sprintf("VkRenderPass(%#x)", hnd.hnd) }
func (hnd Sampler) String() string         { return fmt.Sprintf("VkSampler(%#x)", hnd.hnd) }
func (hnd Semaphore) String() string       { return fmt.Sprintf("VkSemaphore(%#x)", hnd.hnd) }
func (hnd ShaderModule) String() string    { return fmt.Sprintf("VkShaderModule(%#x)", hnd.hnd) }
func (hnd DescriptorSetLayout) String() string {
	return fmt.Sprintf("VkDescriptorSetLayout(%#x)", hnd.hnd)
}

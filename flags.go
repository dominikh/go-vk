// Copyright (c) 2015-2018 Khronos Group
// Copyright (c) 2018 Dominik Honnef

package vk

// #include "vk.h"
import "C"
import (
	"strings"
)

//go:generate stringer -output flags_string.go -type=PresentModeKHR,CommandBufferLevel,ColorSpaceKHR,Format,StructureType,Result,PhysicalDeviceType,SharingMode,ImageViewType,ComponentSwizzle,VertexInputRate,PrimitiveTopology,PolygonMode,FrontFace,BlendFactor,BlendOp,LogicOp,DynamicState,CompareOp,StencilOp,AttachmentLoadOp,AttachmentStoreOp,ImageLayout,PipelineBindPoint,SubpassContents,ImageTiling,ImageType,IndexType,QueryType,Filter,SamplerAddressMode,SamplerMipmapMode,BorderColor,DescriptorType

type DeviceQueueCreateFlags uint32
type QueueFlags uint32
type PhysicalDeviceType uint32
type Result int32
type StructureType uint32
type SurfaceTransformFlagsKHR uint32
type CompositeAlphaFlagsKHR uint32
type ImageUsageFlags uint32
type Format uint32
type ColorSpaceKHR uint32
type PresentModeKHR uint32
type CommandPoolCreateFlags uint32
type CommandPoolTrimFlags uint32
type CommandPoolResetFlags uint32
type CommandBufferLevel uint32
type CommandBufferResetFlags uint32
type CommandBufferUsageFlags uint32
type QueryPipelineStatisticFlags uint32
type QueryControlFlags uint32
type SharingMode uint32
type ImageViewType uint32
type ComponentSwizzle uint32
type ImageAspectFlags uint32
type ShaderStageFlags uint32
type VertexInputRate uint32
type PrimitiveTopology uint32
type PolygonMode uint32
type CullModeFlags uint32
type FrontFace uint32
type SampleCountFlags uint32
type BlendFactor uint32
type BlendOp uint32
type ColorComponentFlags uint32
type LogicOp uint32
type DynamicState uint32
type PipelineCreateFlags uint32
type CompareOp uint32
type StencilOp uint32
type AttachmentDescriptionFlags uint32
type AttachmentLoadOp uint32
type AttachmentStoreOp uint32
type ImageLayout uint32
type PipelineBindPoint uint32
type SubpassDescriptionFlags uint32
type PipelineStageFlags uint32

// Bitmask specifying memory access types that will participate in a memory dependency
type AccessFlags uint32
type DependencyFlags uint32
type SubpassContents uint32
type FenceCreateFlags uint32
type BufferCreateFlags uint32
type BufferUsageFlags uint32
type MemoryPropertyFlags uint32
type MemoryHeapFlags uint32
type MemoryMapFlags uint32
type ImageTiling uint32
type ImageType uint32
type ImageCreateFlags uint32
type IndexType uint32
type QueryType uint32
type Filter uint32
type SamplerAddressMode uint32
type SamplerMipmapMode uint32
type BorderColor uint32
type QueryResultFlags uint32
type StencilFaceFlags uint32
type DescriptorType uint32
type DescriptorPoolCreateFlags uint32
type DescriptorPoolResetFlags uint32
type DescriptorSetLayoutCreateFlags uint32
type FormatFeatureFlags uint32

const (
	SubpassExternal = C.VK_SUBPASS_EXTERNAL
)

const (
	DeviceQueueCreateProtectedBit DeviceQueueCreateFlags = 0x00000001
)

const (
	// QueueGraphicsBit specifies that queues in this queue family support graphics operations.
	QueueGraphicsBit QueueFlags = 0x00000001

	// QueueComputeBit specifies that queues in this queue family support compute operations.
	QueueComputeBit QueueFlags = 0x00000002

	// QueueTransferBit specifies that queues in this queue family support transfer operations.
	QueueTransferBit QueueFlags = 0x00000004

	// QueueSparseBindingBit specifies that queues in this queue family support sparse memory management operations.
	// If any of the sparse resource features are enabled, then at least one queue family must support this bit.
	QueueSparseBindingBit QueueFlags = 0x00000008

	// If QueueProtectedBit is set, then the queues in this queue family support the DeviceQueueCreateProtectedBit bit.
	// If the protected memory physical device feature is supported,
	// then at least one queue family of at least one physical device exposed by the implementation must support this bit.
	QueueProtectedBit QueueFlags = 0x00000010
)

const (
	// The device does not match any other available types.
	PhysicalDeviceTypeOther PhysicalDeviceType = 0

	// The device is typically one embedded in or tightly coupled with the host.
	PhysicalDeviceTypeIntegratedGPU PhysicalDeviceType = 1

	// The device is typically a separate processor connected to the host via an interlink.
	PhysicalDeviceTypeDiscreteGPU PhysicalDeviceType = 2

	// The device is typically a virtual node in a virtualization environment.
	PhysicalDeviceTypeVirtualGPU PhysicalDeviceType = 3

	// The device is typically running on the same processors as the host.
	PhysicalDeviceTypeCPU PhysicalDeviceType = 4
)

const (
	// Command successfully completed.
	Success Result = 0

	// A fence or query has not yet completed.
	NotReady Result = 1

	// A wait operation has not completed in the specified time.
	Timeout Result = 2

	// An event is signaled.
	EventSet Result = 3

	// An event is unsignaled.
	EventReset Result = 4

	// A return array was too small for the result.
	Incomplete Result = 5

	// A swapchain no longer matches the surface properties exactly, but can still be used to present to the surface successfully.
	SuboptimalKHR Result = 1000001003

	// A host memory allocation has failed.
	ErrOutOfHostMemory Result = -1

	// A device memory allocation has failed.
	ErrOutOfDeviceMemory Result = -2

	// Initialization of an object could not be completed for implementation-specific reasons.
	ErrInitializationFailed Result = -3

	// The logical or physical device has been lost.
	ErrDeviceLost Result = -4

	// Mapping of a memory object has failed.
	ErrMemoryMapFailed Result = -5

	// A requested layer is not present or could not be loaded.
	ErrLayerNotPresent Result = -6

	// A requested extension is not supported.
	ErrExtensionNotPresent Result = -7

	// A requested feature is not supported.
	ErrFeatureNotPresent Result = -8

	// The requested version of Vulkan is not supported by the driver or is otherwise incompatible for implementation-specific reasons.
	ErrIncompatibleDriver Result = -9

	// Too many objects of the type have already been created.
	ErrTooManyObjects Result = -10

	// A requested format is not supported on this device.
	ErrFormatNotSupported Result = -11

	// A pool allocation has failed due to fragmentation of the pool’s memory.
	ErrFragmentedPool Result = -12

	// A surface is no longer available.
	ErrSurfaceLostKHR Result = -1000000000

	// The requested window is already in use by Vulkan or another API in a manner which prevents it from being used again.
	ErrNativeWindowInUseKHR Result = -1000000001

	// A surface has changed in such a way that it is no longer compatible with the swapchain,
	// and further presentation requests using the swapchain will fail.
	// Applications must query the new surface properties and recreate their swapchain if they wish to continue presenting to the surface.
	ErrOutOfDateKHR Result = -1000001004

	// The display used by a swapchain does not use the same presentable image layout,
	// or is incompatible in a way that prevents sharing an image.
	ErrIncompatibleDisplayKHR Result = -1000003001

	// A pool memory allocation has failed.
	ErrOutOfPoolMemory Result = -1000069000

	// An external handle is not a valid handle of the specified type.
	ErrInvalidExternalHandle Result = -1000072003

	ErrValidationFailedEXT      Result = -1000011001
	ErrInvalidShaderNV          Result = -1000012000
	ErrFragmentationEXT         Result = -1000161000
	ErrNotPermittedEXT          Result = -1000174001
	ErrOutOfPoolMemoryKHR       Result = ErrOutOfPoolMemory
	ErrInvalidExternalHandleKHR Result = ErrInvalidExternalHandle
)

const (
	StructureTypeAccelerationStructureCreateInfoNVX                    StructureType = 1000165001
	StructureTypeAccelerationStructureMemoryRequirementsInfoNVX        StructureType = 1000165008
	StructureTypeAcquireNextImageInfoKHR                               StructureType = 1000060010
	StructureTypeAndroidHardwareBufferFormatPropertiesAndroid          StructureType = 1000129002
	StructureTypeAndroidHardwareBufferPropertiesAndroid                StructureType = 1000129001
	StructureTypeAndroidHardwareBufferUsageAndroid                     StructureType = 1000129000
	StructureTypeAndroidSurfaceCreateInfoKHR                           StructureType = 1000008000
	StructureTypeAttachmentDescription2KHR                             StructureType = 1000109000
	StructureTypeAttachmentReference2KHR                               StructureType = 1000109001
	StructureTypeApplicationInfo                                       StructureType = 0
	StructureTypeBindAccelerationStructureMemoryInfoNVX                StructureType = 1000165006
	StructureTypeBindBufferMemoryDeviceGroupInfoKHR                    StructureType = 1000060013
	StructureTypeBindBufferMemoryInfoKHR                               StructureType = 1000157000
	StructureTypeBindImageMemoryDeviceGroupInfoKHR                     StructureType = 1000060014
	StructureTypeBindImageMemoryInfoKHR                                StructureType = 1000157001
	StructureTypeBindImageMemorySwapchainInfoKHR                       StructureType = 1000060009
	StructureTypeBindImagePlaneMemoryINFO                              StructureType = 1000156002
	StructureTypeBindImagePlaneMemoryInfoKHR                           StructureType = 1000156002
	StructureTypeBufferMemoryRequirementsInfo2                         StructureType = 1000146000
	StructureTypeBufferMemoryRequirementsInfo2KHR                      StructureType = 1000146000
	StructureTypeBindBufferMemoryDeviceGroupInfo                       StructureType = 1000060013
	StructureTypeBindBufferMemoryInfo                                  StructureType = 1000157000
	StructureTypeBindImageMemoryDeviceGroupInfo                        StructureType = 1000060014
	StructureTypeBindImageMemoryInfo                                   StructureType = 1000157001
	StructureTypeBindSparseInfo                                        StructureType = 7
	StructureTypeBufferCreateInfo                                      StructureType = 12
	StructureTypeBufferMemoryBarrier                                   StructureType = 44
	StructureTypeBufferViewCreateInfo                                  StructureType = 13
	StructureTypeCheckpointDataNV                                      StructureType = 1000206000
	StructureTypeCmdProcessCommandsInfoNVX                             StructureType = 1000086002
	StructureTypeCmdReserveSpaceForCommandsInfoNVX                     StructureType = 1000086003
	StructureTypeCommandBufferInheritanceConditionalRenderingInfoEXT   StructureType = 1000081000
	StructureTypeConditionalRenderingBeginInfoEXT                      StructureType = 1000081002
	StructureTypeCommandBufferAllocateInfo                             StructureType = 40
	StructureTypeCommandBufferBeginInfo                                StructureType = 42
	StructureTypeCommandBufferInheritanceInfo                          StructureType = 41
	StructureTypeCommandPoolCreateInfo                                 StructureType = 39
	StructureTypeComputePipelineCreateInfo                             StructureType = 29
	StructureTypeCopyDescriptorSet                                     StructureType = 36
	StructureTypeD3D12FenceSubmitInfoKHR                               StructureType = 1000078002
	StructureTypeDebugMarkerMarkerInfoEXT                              StructureType = 1000022002
	StructureTypeDebugMarkerObjectNameInfoEXT                          StructureType = 1000022000
	StructureTypeDebugMarkerObjectTagInfoEXT                           StructureType = 1000022001
	StructureTypeDebugReportCallbackCreateInfoEXT                      StructureType = 1000011000
	StructureTypeDebugReportCreateInfoEXT                              StructureType = 1000011000
	StructureTypeDebugUtilsLabelEXT                                    StructureType = 1000128002
	StructureTypeDebugUtilsMessengerCallbackDataEXT                    StructureType = 1000128003
	StructureTypeDebugUtilsMessengerCreateInfoEXT                      StructureType = 1000128004
	StructureTypeDebugUtilsObjectNameInfoEXT                           StructureType = 1000128000
	StructureTypeDebugUtilsObjectTagInfoEXT                            StructureType = 1000128001
	StructureTypeDedicatedAllocationBufferCreateInfoNV                 StructureType = 1000026001
	StructureTypeDedicatedAllocationImageCreateInfoNV                  StructureType = 1000026000
	StructureTypeDedicatedAllocationMemoryAllocateInfoNV               StructureType = 1000026002
	StructureTypeDescriptorAccelerationStructureInfoNVX                StructureType = 1000165007
	StructureTypeDescriptorPoolInlineUniformBlockCreateInfoEXT         StructureType = 1000138003
	StructureTypeDescriptorSetLayoutBindingFlagsCreateInfoEXT          StructureType = 1000161000
	StructureTypeDescriptorSetLayoutSupport                            StructureType = 1000168001
	StructureTypeDescriptorSetLayoutSupportKHR                         StructureType = 1000168001
	StructureTypeDescriptorSetVariableDescriptorCountAllocateInfoExt   StructureType = 1000161003
	StructureTypeDescriptorSetVariableDescriptorCountLayoutSupportEXT  StructureType = 1000161004
	StructureTypeDescriptorUpdateTemplateCreateINFO                    StructureType = 1000085000
	StructureTypeDescriptorUpdateTemplateCreateInfoKHR                 StructureType = 1000085000
	StructureTypeDeviceEventInfoEXT                                    StructureType = 1000091001
	StructureTypeDeviceGeneratedCommandsFeaturesNVX                    StructureType = 1000086005
	StructureTypeDeviceGeneratedCommandsLimitsNVX                      StructureType = 1000086004
	StructureTypeDeviceGroupBindSparseInfoKHR                          StructureType = 1000060006
	StructureTypeDeviceGroupCommandBufferBeginInfoKHR                  StructureType = 1000060004
	StructureTypeDeviceGroupDeviceCreateInfoKHR                        StructureType = 1000070001
	StructureTypeDeviceGroupPresentCapabilitiesKHR                     StructureType = 1000060007
	StructureTypeDeviceGroupPresentInfoKHR                             StructureType = 1000060011
	StructureTypeDeviceGroupRenderPassBeginInfoKHR                     StructureType = 1000060003
	StructureTypeDeviceGroupSubmitInfoKHR                              StructureType = 1000060005
	StructureTypeDeviceGroupSwapchainCreateInfoKHR                     StructureType = 1000060012
	StructureTypeDeviceQueueGlobalPriorityCreateInfoEXT                StructureType = 1000174000
	StructureTypeDeviceQueueInfo2                                      StructureType = 1000145003
	StructureTypeDisplayEventInfoEXT                                   StructureType = 1000091002
	StructureTypeDisplayModeCreateInfoKHR                              StructureType = 1000002000
	StructureTypeDisplayModeProperties2KHR                             StructureType = 1000121002
	StructureTypeDisplayPlaneCapabilities2KHR                          StructureType = 1000121004
	StructureTypeDisplayPlaneInfo2KHR                                  StructureType = 1000121003
	StructureTypeDisplayPlaneProperties2KHR                            StructureType = 1000121001
	StructureTypeDisplayPowerInfoEXT                                   StructureType = 1000091000
	StructureTypeDisplayPresentInfoKHR                                 StructureType = 1000003000
	StructureTypeDisplayProperties2KHR                                 StructureType = 1000121000
	StructureTypeDisplaySurfaceCreateInfoKHR                           StructureType = 1000002001
	StructureTypeDescriptorPoolCreateInfo                              StructureType = 33
	StructureTypeDescriptorSetAllocateInfo                             StructureType = 34
	StructureTypeDescriptorSetLayoutCreateInfo                         StructureType = 32
	StructureTypeDeviceCreateInfo                                      StructureType = 3
	StructureTypeDeviceGroupBindSparseInfo                             StructureType = 1000060006
	StructureTypeDeviceGroupCommandBufferBeginInfo                     StructureType = 1000060004
	StructureTypeDeviceGroupDeviceCreateInfo                           StructureType = 1000070001
	StructureTypeDeviceGroupRenderPassBeginInfo                        StructureType = 1000060003
	StructureTypeDeviceGroupSubmitInfo                                 StructureType = 1000060005
	StructureTypeDeviceQueueCreateInfo                                 StructureType = 2
	StructureTypeExportFenceCreateInfo                                 StructureType = 1000113000
	StructureTypeExportFenceCreateInfoKHR                              StructureType = 1000113000
	StructureTypeExportFenceWin32HandleInfoKHR                         StructureType = 1000114001
	StructureTypeExportMemoryAllocateInfo                              StructureType = 1000072002
	StructureTypeExportMemoryAllocateInfoKHR                           StructureType = 1000072002
	StructureTypeExportMemoryAllocateInfoNV                            StructureType = 1000056001
	StructureTypeExportMemoryWin32HandleInfoKHR                        StructureType = 1000073001
	StructureTypeExportMemoryWin32HandleInfoNV                         StructureType = 1000057001
	StructureTypeExportSemaphoreCreateInfo                             StructureType = 1000077000
	StructureTypeExportSemaphoreCreateInfoKHR                          StructureType = 1000077000
	StructureTypeExportSemaphoreWin32HandleInfoKHR                     StructureType = 1000078001
	StructureTypeExternalBufferProperties                              StructureType = 1000071003
	StructureTypeExternalBufferPropertiesKHR                           StructureType = 1000071003
	StructureTypeExternalFenceProperties                               StructureType = 1000112001
	StructureTypeExternalFencePropertiesKHR                            StructureType = 1000112001
	StructureTypeExternalFormatAndroid                                 StructureType = 1000129005
	StructureTypeExternalImageFormatProperties                         StructureType = 1000071001
	StructureTypeExternalImageFormatPropertiesKHR                      StructureType = 1000071001
	StructureTypeExternalMemoryBufferCreateInfo                        StructureType = 1000072000
	StructureTypeExternalMemoryBufferCreateInfoKHR                     StructureType = 1000072000
	StructureTypeExternalMemoryImageCreateInfo                         StructureType = 1000072001
	StructureTypeExternalMemoryImageCreateInfoKHR                      StructureType = 1000072001
	StructureTypeExternalMemoryImageCreateInfoNV                       StructureType = 1000056000
	StructureTypeExternalSemaphoreProperties                           StructureType = 1000076001
	StructureTypeExternalSemaphorePropertiesKHR                        StructureType = 1000076001
	StructureTypeEventCreateInfo                                       StructureType = 10
	StructureTypeFenceGetFdInfoKHR                                     StructureType = 1000115001
	StructureTypeFenceGetWin32HandleInfoKHR                            StructureType = 1000114002
	StructureTypeFormatProperties2                                     StructureType = 1000059002
	StructureTypeFormatProperties2KHR                                  StructureType = 1000059002
	StructureTypeFenceCreateInfo                                       StructureType = 8
	StructureTypeFramebufferCreateInfo                                 StructureType = 37
	StructureTypeGeometryAABB_NVX                                      StructureType = 1000165005
	StructureTypeGeometryInstanceNVX                                   StructureType = 1000165002
	StructureTypeGeometryNVX                                           StructureType = 1000165003
	StructureTypeGeometryTrianglesNVX                                  StructureType = 1000165004
	StructureTypeGraphicsPipelineCreateInfo                            StructureType = 28
	StructureTypeHdrMetadataEXT                                        StructureType = 1000105000
	StructureTypeHitShaderModuleCreateInfoNVX                          StructureType = 1000165010
	StructureTypeImageFormatListCreateInfoKHR                          StructureType = 1000147000
	StructureTypeImageFormatProperties2                                StructureType = 1000059003
	StructureTypeImageFormatProperties2KHR                             StructureType = 1000059003
	StructureTypeImageMemoryRequirementsInfo2                          StructureType = 1000146001
	StructureTypeImageMemoryRequirementsInfo2KHR                       StructureType = 1000146001
	StructureTypeImagePlaneMemoryRequirementsInfo                      StructureType = 1000156003
	StructureTypeImagePlaneMemoryRequirementsInfoKHR                   StructureType = 1000156003
	StructureTypeImageSparseMemoryRequirementsInfo2                    StructureType = 1000146002
	StructureTypeImageSparseMemoryRequirementsInfo2KHR                 StructureType = 1000146002
	StructureTypeImageSwapchainCreateInfoKHR                           StructureType = 1000060008
	StructureTypeImageViewAstcDecodeModeEXT                            StructureType = 1000067000
	StructureTypeImageViewUsageCreateInfo                              StructureType = 1000117002
	StructureTypeImageViewUsageCreateInfoKHR                           StructureType = 1000117002
	StructureTypeImportAndroidHardwareBufferInfoAndroid                StructureType = 1000129003
	StructureTypeImportFenceFdInfoKHR                                  StructureType = 1000115000
	StructureTypeImportFenceWin32HandleInfoKHR                         StructureType = 1000114000
	StructureTypeImportMemoryFdInfoKHR                                 StructureType = 1000074000
	StructureTypeImportMemoryHostPointerInfoEXT                        StructureType = 1000178000
	StructureTypeImportMemoryWin32HandleInfoKHR                        StructureType = 1000073000
	StructureTypeImportMemoryWin32HandleInfoNV                         StructureType = 1000057000
	StructureTypeImportSemaphoreFdInfoKHR                              StructureType = 1000079000
	StructureTypeImportSemaphoreWin32HandleInfoKHR                     StructureType = 1000078000
	StructureTypeIndirectCommandsLayoutCreateInfoNVX                   StructureType = 1000086001
	StructureTypeIosSurfaceCreateInfoMVK                               StructureType = 1000122000
	StructureTypeImageCreateInfo                                       StructureType = 14
	StructureTypeImageMemoryBarrier                                    StructureType = 45
	StructureTypeImageViewCreateInfo                                   StructureType = 15
	StructureTypeInstanceCreateInfo                                    StructureType = 1
	StructureTypeLoaderDeviceCreateInfo                                StructureType = 48
	StructureTypeLoaderInstanceCreateInfo                              StructureType = 47
	StructureTypeMacosSurfaceCreateInfoMVK                             StructureType = 1000123000
	StructureTypeMemoryAllocateFlagsInfoKHR                            StructureType = 1000060000
	StructureTypeMemoryDedicatedAllocateInfoKHR                        StructureType = 1000127001
	StructureTypeMemoryDedicatedRequirementsKHR                        StructureType = 1000127000
	StructureTypeMemoryFdPropertiesKHR                                 StructureType = 1000074001
	StructureTypeMemoryGetAndroidHardwareBufferInfoAndroid             StructureType = 1000129004
	StructureTypeMemoryGetFdInfoKHR                                    StructureType = 1000074002
	StructureTypeMemoryGetWin32HandleInfoKHR                           StructureType = 1000073003
	StructureTypeMemoryHostPointerPropertiesEXT                        StructureType = 1000178001
	StructureTypeMemoryRequirements2                                   StructureType = 1000146003
	StructureTypeMemoryRequirements2KHR                                StructureType = 1000146003
	StructureTypeMemoryWin32HandlePropertiesKHR                        StructureType = 1000073002
	StructureTypeMirSurfaceCreateInfoKHR                               StructureType = 1000007000
	StructureTypeMultisamplePropertiesEXT                              StructureType = 1000143004
	StructureTypeMappedMemoryRange                                     StructureType = 6
	StructureTypeMemoryAllocateFlagsInfo                               StructureType = 1000060000
	StructureTypeMemoryAllocateInfo                                    StructureType = 5
	StructureTypeMemoryBarrier                                         StructureType = 46
	StructureTypeMemoryDedicatedAllocateInfo                           StructureType = 1000127001
	StructureTypeMemoryDedicatedRequirements                           StructureType = 1000127000
	StructureTypeObjectTableCreateInfoNVX                              StructureType = 1000086000
	StructureTypePhysicalDevice16BitStorageFeaturesKHR                 StructureType = 1000083000
	StructureTypePhysicalDevice8BitStorageFeaturesKHR                  StructureType = 1000177000
	StructureTypePhysicalDeviceAstcDecodeFeaturesEXT                   StructureType = 1000067001
	StructureTypePhysicalDeviceBlendOperationAdvancedFeaturesEXT       StructureType = 1000148000
	StructureTypePhysicalDeviceBlendOperationAdvancedPropertiesEXT     StructureType = 1000148001
	StructureTypePhysicalDeviceComputeShaderDerivativesFeatures_NV     StructureType = 1000201000
	StructureTypePhysicalDeviceConditionalRenderingFeaturesEXT         StructureType = 1000081001
	StructureTypePhysicalDeviceConservativeRasterizationPropertiesEXT  StructureType = 1000101000
	StructureTypePhysicalDeviceCornerSampledImageFeaturesNV            StructureType = 1000050000
	StructureTypePhysicalDeviceDescriptorIndexingFeaturesEXT           StructureType = 1000161001
	StructureTypePhysicalDeviceDescriptorIndexingPropertiesEXT         StructureType = 1000161002
	StructureTypePhysicalDeviceDiscardRectanglePropertiesEXT           StructureType = 1000099000
	StructureTypePhysicalDeviceExclusiveScissorFeaturesNV              StructureType = 1000205002
	StructureTypePhysicalDeviceExternalBufferInfo                      StructureType = 1000071002
	StructureTypePhysicalDeviceExternalBufferInfoKHR                   StructureType = 1000071002
	StructureTypePhysicalDeviceExternalFenceInfo                       StructureType = 1000112000
	StructureTypePhysicalDeviceExternalFenceInfoKHR                    StructureType = 1000112000
	StructureTypePhysicalDeviceExternalImageFormatInfo                 StructureType = 1000071000
	StructureTypePhysicalDeviceExternalImageFormatInfoKHR              StructureType = 1000071000
	StructureTypePhysicalDeviceExternalMemoryHostPropertiesEXT         StructureType = 1000178002
	StructureTypePhysicalDeviceExternalSemaphoreInfo                   StructureType = 1000076000
	StructureTypePhysicalDeviceExternalSemaphoreInfoKHR                StructureType = 1000076000
	StructureTypePhysicalDeviceFeatures2                               StructureType = 1000059000
	StructureTypePhysicalDeviceFeatures2KHR                            StructureType = 1000059000
	StructureTypePhysicalDeviceFragmentShaderBarycentricFeaturesNV     StructureType = 1000203000
	StructureTypePhysicalDeviceGroupPropertiesKHR                      StructureType = 1000070000
	StructureTypePhysicalDeviceIdProperties                            StructureType = 1000071004
	StructureTypePhysicalDeviceIdPropertiesKHR                         StructureType = 1000071004
	StructureTypePhysicalDeviceImageFormatInfo2                        StructureType = 1000059004
	StructureTypePhysicalDeviceImageFormatInfo2KHR                     StructureType = 1000059004
	StructureTypePhysicalDeviceInlineUniformBlockFeaturesEXT           StructureType = 1000138000
	StructureTypePhysicalDeviceInlineUniformBlockPropertiesEXT         StructureType = 1000138001
	StructureTypePhysicalDeviceMaintenance3Properties                  StructureType = 1000168000
	StructureTypePhysicalDeviceMaintenance3PropertiesKHR               StructureType = 1000168000
	StructureTypePhysicalDeviceMemoryProperties2                       StructureType = 1000059006
	StructureTypePhysicalDeviceMemoryProperties2KHR                    StructureType = 1000059006
	StructureTypePhysicalDeviceMeshShaderFeaturesNV                    StructureType = 1000202000
	StructureTypePhysicalDeviceMeshShaderPropertiesNV                  StructureType = 1000202001
	StructureTypePhysicalDeviceMultiviewFeatures                       StructureType = 1000053001
	StructureTypePhysicalDeviceMultiviewFeaturesKHR                    StructureType = 1000053001
	StructureTypePhysicalDeviceMultiviewPerViewAttributesPropertiesNVX StructureType = 1000097000
	StructureTypePhysicalDeviceMultiviewProperties                     StructureType = 1000053002
	StructureTypePhysicalDeviceMultiviewPropertiesKHR                  StructureType = 1000053002
	StructureTypePhysicalDevicePointClippingProperties                 StructureType = 1000117000
	StructureTypePhysicalDevicePointClippingPropertiesKHR              StructureType = 1000117000
	StructureTypePhysicalDeviceProperties2                             StructureType = 1000059001
	StructureTypePhysicalDeviceProperties2KHR                          StructureType = 1000059001
	StructureTypePhysicalDeviceProtectedMemoryFeatures                 StructureType = 1000145001
	StructureTypePhysicalDeviceProtectedMemoryProperties               StructureType = 1000145002
	StructureTypePhysicalDevicePushDescriptorPropertiesKHR             StructureType = 1000080000
	StructureTypePhysicalDeviceRaytracingPropertiesNVX                 StructureType = 1000165009
	StructureTypePhysicalDeviceRepresentativeFragmentTestFeaturesNV    StructureType = 1000166000
	StructureTypePhysicalDeviceSamplerFilterMinmaxPropertiesEXT        StructureType = 1000130000
	StructureTypePhysicalDeviceSamplerYCbCrConversionFeatures          StructureType = 1000156004
	StructureTypePhysicalDeviceSamplerYCbCrConversionFeaturesKHR       StructureType = 1000156004
	StructureTypePhysicalDeviceSampleLocationsPropertiesEXT            StructureType = 1000143003
	StructureTypePhysicalDeviceShaderCorePropertiesAMD                 StructureType = 1000185000
	StructureTypePhysicalDeviceShaderDrawParameterFeatures             StructureType = 1000063000
	StructureTypePhysicalDeviceShaderImageFootprintFeaturesNV          StructureType = 1000204000
	StructureTypePhysicalDeviceShadingRateImageFeaturesNV              StructureType = 1000164001
	StructureTypePhysicalDeviceShadingRateImagePropertiesNV            StructureType = 1000164002
	StructureTypePhysicalDeviceSparseImageFormatInfo2                  StructureType = 1000059008
	StructureTypePhysicalDeviceSparseImageFormatInfo2KHR               StructureType = 1000059008
	StructureTypePhysicalDeviceSurfaceInfo2KHR                         StructureType = 1000119000
	StructureTypePhysicalDeviceVariablePointerFeatures                 StructureType = 1000120000
	StructureTypePhysicalDeviceVariablePointerFeaturesKHR              StructureType = 1000120000
	StructureTypePhysicalDeviceVertexAttributeDivisorFeaturesEXT       StructureType = 1000190002
	StructureTypePhysicalDeviceVertexAttributeDivisorPropertiesEXT     StructureType = 1000190000
	StructureTypePhysicalDeviceVulkanMemoryModelFeaturesKHR            StructureType = 1000211000
	StructureTypePipelineColorBlendAdvancedStateCreateInfoEXT          StructureType = 1000148002
	StructureTypePipelineCoverageModulationStateCreateInfoNV           StructureType = 1000152000
	StructureTypePipelineCoverageToColorStateCreateInfoNV              StructureType = 1000149000
	StructureTypePipelineDiscardRectangleStateCreateInfoEXT            StructureType = 1000099001
	StructureTypePipelineRasterizationConservativeStateCreateInfoEXT   StructureType = 1000101001
	StructureTypePipelineRasterizationStateRasterizationOrderAMD       StructureType = 1000018000
	StructureTypePipelineRepresentativeFragmentTestStateCreateInfoNV   StructureType = 1000166001
	StructureTypePipelineSampleLocationsStateCreateInfoEXT             StructureType = 1000143002
	StructureTypePipelineTessellationDomainOriginStateCreateInfo       StructureType = 1000117003
	StructureTypePipelineTessellationDomainOriginStateCreateInfoKHR    StructureType = 1000117003
	StructureTypePipelineVertexInputDivisorStateCreateInfoEXT          StructureType = 1000190001
	StructureTypePipelineViewportCoarseSampleOrderStateCreateInfoNV    StructureType = 1000164005
	StructureTypePipelineViewportExclusiveScissorStateCreateInfoNV     StructureType = 1000205000
	StructureTypePipelineViewportShadingRateImageStateCreateInfoNV     StructureType = 1000164000
	StructureTypePipelineViewportSwizzleStateCreateInfoNV              StructureType = 1000098000
	StructureTypePipelineViewportWScalingStateCreateInfoNV             StructureType = 1000087000
	StructureTypePresentInfoKHR                                        StructureType = 1000001001
	StructureTypePresentRegionsKHR                                     StructureType = 1000084000
	StructureTypePresentTimesInfoGoogle                                StructureType = 1000092000
	StructureTypeProtectedSubmitInfo                                   StructureType = 1000145000
	StructureTypePhysicalDevice16ItStorageFeatures                     StructureType = 1000083000
	StructureTypePhysicalDeviceGroupProperties                         StructureType = 1000070000
	StructureTypePhysicalDeviceSubgroupProperties                      StructureType = 1000094000
	StructureTypePipelineCacheCreateInfo                               StructureType = 17
	StructureTypePipelineColorBlendStateCreateInfo                     StructureType = 26
	StructureTypePipelineDepthStencilStateCreateInfo                   StructureType = 25
	StructureTypePipelineDynamicStateCreateInfo                        StructureType = 27
	StructureTypePipelineInputAssemblyStateCreateInfo                  StructureType = 20
	StructureTypePipelineLayoutCreateInfo                              StructureType = 30
	StructureTypePipelineMultisampleStateCreateInfo                    StructureType = 24
	StructureTypePipelineRasterizationStateCreateInfo                  StructureType = 23
	StructureTypePipelineShaderStageCreateInfo                         StructureType = 18
	StructureTypePipelineTessellationStateCreateInfo                   StructureType = 21
	StructureTypePipelineVertexInputStateCreateInfo                    StructureType = 19
	StructureTypePipelineViewportStateCreateInfo                       StructureType = 22
	StructureTypeQueueFamilyCheckpointPropertiesNV                     StructureType = 1000206001
	StructureTypeQueueFamilyProperties2                                StructureType = 1000059005
	StructureTypeQueueFamilyProperties2KHR                             StructureType = 1000059005
	StructureTypeQueryPoolCreateInfo                                   StructureType = 11
	StructureTypeRaytracingPipelineCreateInfoNVX                       StructureType = 1000165000
	StructureTypeRenderPassCreateInfo2KHR                              StructureType = 1000109004
	StructureTypeRenderPassInputAttachmentAspectCreateInfo             StructureType = 1000117001
	StructureTypeRenderPassInputAttachmentAspectCreateInfoKHR          StructureType = 1000117001
	StructureTypeRenderPassMultiviewCreateInfo                         StructureType = 1000053000
	StructureTypeRenderPassMultiviewCreateInfoKHR                      StructureType = 1000053000
	StructureTypeRenderPassSampleLocationsBeginInfoEXT                 StructureType = 1000143001
	StructureTypeRenderPassBeginInfo                                   StructureType = 43
	StructureTypeRenderPassCreateInfo                                  StructureType = 38
	StructureTypeSamplerReductionModeCreateInfoEXT                     StructureType = 1000130001
	StructureTypeSamplerYCbCrConversionCreateInfo                      StructureType = 1000156000
	StructureTypeSamplerYCbCrConversionCreateInfoKHR                   StructureType = 1000156000
	StructureTypeSamplerYCbCrConversionImageFormatProperties           StructureType = 1000156005
	StructureTypeSamplerYCbCrConversionImageFormatPropertiesKHR        StructureType = 1000156005
	StructureTypeSamplerYCbCrConversionINFO                            StructureType = 1000156001
	StructureTypeSamplerYCbCrConversionINFO_KHR                        StructureType = 1000156001
	StructureTypeSampleLocationsInfoEXT                                StructureType = 1000143000
	StructureTypeSemaphoreGetFdInfoKHR                                 StructureType = 1000079001
	StructureTypeSemaphoreGetWin32HandleInfoKHR                        StructureType = 1000078003
	StructureTypeShaderModuleValidationCacheCreateInfoEXT              StructureType = 1000160001
	StructureTypeSharedPresentSurfaceCapabilitiesKHR                   StructureType = 1000111000
	StructureTypeSparseImageFormatProperties2                          StructureType = 1000059007
	StructureTypeSparseImageFormatProperties2KHR                       StructureType = 1000059007
	StructureTypeSparseImageMemoryRequirements2                        StructureType = 1000146004
	StructureTypeSparseImageMemoryRequirements2KHR                     StructureType = 1000146004
	StructureTypeSubpassBeginInfoKHR                                   StructureType = 1000109005
	StructureTypeSubpassDependency2KHR                                 StructureType = 1000109003
	StructureTypeSubpassDescription2KHR                                StructureType = 1000109002
	StructureTypeSubpassEndInfoKHR                                     StructureType = 1000109006
	StructureTypeSurfaceCapabilities2EXT                               StructureType = 1000090000
	StructureTypeSurfaceCapabilities2KHR                               StructureType = 1000119001
	StructureTypeSurfaceFormat2KHR                                     StructureType = 1000119002
	StructureTypeSwapchainCounterCreateInfoEXT                         StructureType = 1000091003
	StructureTypeSwapchainCreateInfoKHR                                StructureType = 1000001000
	StructureTypeSamplerCreateInfo                                     StructureType = 31
	StructureTypeSemaphoreCreateInfo                                   StructureType = 9
	StructureTypeShaderModuleCreateInfo                                StructureType = 16
	StructureTypeSubmitInfo                                            StructureType = 4
	StructureTypeTextureLodGatherFormatPropertiesAMD                   StructureType = 1000041000
	StructureTypeValidationCacheCreateInfoEXT                          StructureType = 1000160000
	StructureTypeValidationFlagsEXT                                    StructureType = 1000061000
	StructureTypeViSurfaceCreateInfoNN                                 StructureType = 1000062000
	StructureTypeWaylandSurfaceCreateInfoKHR                           StructureType = 1000006000
	StructureTypeWin32KeyedMutexAcquireReleaseInfoKHR                  StructureType = 1000075000
	StructureTypeWin32KeyedMutexAcquireReleaseInfoNV                   StructureType = 1000058000
	StructureTypeWin32SurfaceCreateInfoKHR                             StructureType = 1000009000
	StructureTypeWriteDescriptorSetInlineUniformBlockEXT               StructureType = 1000138002
	StructureTypeWriteDescriptorSet                                    StructureType = 35
	StructureTypeXcbSurfaceCreateInfoKHR                               StructureType = 1000005000
	StructureTypeXlibSurfaceCreateInfoKHR                              StructureType = 1000004000
)

const (
	// SurfaceTransformIdentityBitKHR specifies that image content is presented without being transformed.
	SurfaceTransformIdentityBitKHR SurfaceTransformFlagsKHR = 0x00000001

	// SurfaceTransformRotate90BitKHR  specifies that image content is rotated 90 degrees clockwise.
	SurfaceTransformRotate90BitKHR SurfaceTransformFlagsKHR = 0x00000002

	// SurfaceTransformRotate180BitKHR specifies that image content is rotated 180 degrees clockwise.
	SurfaceTransformRotate180BitKHR SurfaceTransformFlagsKHR = 0x00000004

	// SurfaceTransformRotate270BitKHR specifies that image content is rotated 270 degrees clockwise.
	SurfaceTransformRotate270BitKHR SurfaceTransformFlagsKHR = 0x00000008

	// SurfaceTransformHorizontalMirrorBitKhr specifies that image content is mirrored horizontally.
	SurfaceTransformHorizontalMirrorBitKHR SurfaceTransformFlagsKHR = 0x00000010

	// SurfaceTransformHorizontalMirrorRotate90BitKHR specifies that image content is mirrored horizontally, then rotated 90 degrees clockwise.
	SurfaceTransformHorizontalMirrorRotate90BitKHR SurfaceTransformFlagsKHR = 0x00000020

	// SurfaceTransformHorizontalMirrorRotate180BitKHR specifies that image content is mirrored horizontally, then rotated 180 degrees clockwise.
	SurfaceTransformHorizontalMirrorRotate180BitKHR SurfaceTransformFlagsKHR = 0x00000040

	// SurfaceTransformHorizontalMirrorRotate270BitKHR specifies that image content is mirrored horizontally, then rotated 270 degrees clockwise.
	SurfaceTransformHorizontalMirrorRotate270BitKHR SurfaceTransformFlagsKHR = 0x00000080

	// SurfaceTransformInheritBitKHR specifies that the presentation transform is not specified,
	// and is instead determined by platform-specific considerations and mechanisms outside Vulkan.
	SurfaceTransformInheritBitKHR SurfaceTransformFlagsKHR = 0x00000100
)

const (
	// The alpha channel, if it exists, of the images is ignored in the compositing process.
	// Instead, the image is treated as if it has a constant alpha of 1.0.
	CompositeAlphaOpaqueBitKHR CompositeAlphaFlagsKHR = 0x00000001

	// The alpha channel, if it exists, of the images is respected in the compositing process.
	// The non-alpha channels of the image are expected to already be multiplied by the alpha channel by the application.
	CompositeAlphaPreMultipliedBitKHR CompositeAlphaFlagsKHR = 0x00000002

	// The alpha channel, if it exists, of the images is respected in the compositing process.
	// The non-alpha channels of the image are not expected to already be multiplied by the alpha channel by the application;
	// instead, the compositor will multiply the non-alpha channels of the image by the alpha channel during compositing.
	CompositeAlphaPostMultipliedBitKHR CompositeAlphaFlagsKHR = 0x00000004

	// The way in which the presentation engine treats the alpha channel in the images is unknown to the Vulkan API.
	// Instead, the application is responsible for setting the composite alpha blending mode using native window system commands.
	// If the application does not set the blending mode using native window system commands, then a platform-specific default will be used.
	CompositeAlphaInheritBitKHR CompositeAlphaFlagsKHR = 0x00000008
)

const (
	// ImageUsageTransferSrcBit specifies that the image can be used as the source of a transfer command.
	ImageUsageTransferSrcBit ImageUsageFlags = 0x00000001

	// ImageUsageTransferDstBit specifies that the image can be used as the destination of a transfer command.
	ImageUsageTransferDstBit ImageUsageFlags = 0x00000002

	// ImageUsageSampledBit specifies that the image can be used to create an ImageView
	// suitable for occupying a DescriptorSet slot either of type DescriptorTypeSampledImage
	// or DescriptorTypeCombinedImageSampler, and be sampled by a shader.
	ImageUsageSampledBit ImageUsageFlags = 0x00000004

	// ImageUsageStorageBit specifies that the image can be used to create an ImageView
	// suitable for occupying a DescriptorSet slot of type DescriptorTypeStorageImage.
	ImageUsageStorageBit ImageUsageFlags = 0x00000008

	// ImageUsageColorAttachmentBit specifies that the image can be used to create an ImageView
	// suitable for use as a color or resolve attachment in a Framebuffer.
	ImageUsageColorAttachmentBit ImageUsageFlags = 0x00000010

	// ImageUsageDepthStencilAttachmentBit specifies that the image can be used to create an ImageView
	// suitable for use as a depth/stencil attachment in a Framebuffer.
	ImageUsageDepthStencilAttachmentBit ImageUsageFlags = 0x00000020

	// ImageUsageTransientAttachmentBit specifies that the memory bound to this image will have been allocated
	// with the MemoryPropertyLazilyAllocatedBit (see Memory Allocation for more detail).
	// This bit can be set for any image that can be used to create an ImageView
	// suitable for use as a color, resolve, depth/stencil, or input attachment.
	ImageUsageTransientAttachmentBit ImageUsageFlags = 0x00000040

	// ImageUsageInputAttachmentBit specifies that the image can be used to create an ImageView suitable
	// for occupying DescriptorSet slot of type DescriptorTypeInputAttachment;
	// be read from a shader as an input attachment; and be used as an input attachment in a framebuffer.
	ImageUsageInputAttachmentBit ImageUsageFlags = 0x00000080
)

const (
	FormatUndefined                  Format = 0
	FormatR4G4_UNORM_PACK8           Format = 1
	FormatR4G4B4A4_UNORM_PACK16      Format = 2
	FormatB4G4R4A4_UNORM_PACK16      Format = 3
	FormatR5G6B5_UNORM_PACK16        Format = 4
	FormatB5G6R5_UNORM_PACK16        Format = 5
	FormatR5G5B5A1_UNORM_PACK16      Format = 6
	FormatB5G5R5A1_UNORM_PACK16      Format = 7
	FormatA1R5G5B5_UNORM_PACK16      Format = 8
	FormatR8_UNORM                   Format = 9
	FormatR8_SNORM                   Format = 10
	FormatR8_USCALED                 Format = 11
	FormatR8_SSCALED                 Format = 12
	FormatR8_UINT                    Format = 13
	FormatR8_SINT                    Format = 14
	FormatR8_SRGB                    Format = 15
	FormatR8G8_UNORM                 Format = 16
	FormatR8G8_SNORM                 Format = 17
	FormatR8G8_USCALED               Format = 18
	FormatR8G8_SSCALED               Format = 19
	FormatR8G8_UINT                  Format = 20
	FormatR8G8_SINT                  Format = 21
	FormatR8G8_SRGB                  Format = 22
	FormatR8G8B8_UNORM               Format = 23
	FormatR8G8B8_SNORM               Format = 24
	FormatR8G8B8_USCALED             Format = 25
	FormatR8G8B8_SSCALED             Format = 26
	FormatR8G8B8_UINT                Format = 27
	FormatR8G8B8_SINT                Format = 28
	FormatR8G8B8_SRGB                Format = 29
	FormatB8G8R8_UNORM               Format = 30
	FormatB8G8R8_SNORM               Format = 31
	FormatB8G8R8_USCALED             Format = 32
	FormatB8G8R8_SSCALED             Format = 33
	FormatB8G8R8_UINT                Format = 34
	FormatB8G8R8_SINT                Format = 35
	FormatB8G8R8_SRGB                Format = 36
	FormatR8G8B8A8_UNORM             Format = 37
	FormatR8G8B8A8_SNORM             Format = 38
	FormatR8G8B8A8_USCALED           Format = 39
	FormatR8G8B8A8_SSCALED           Format = 40
	FormatR8G8B8A8_UINT              Format = 41
	FormatR8G8B8A8_SINT              Format = 42
	FormatR8G8B8A8_SRGB              Format = 43
	FormatB8G8R8A8_UNORM             Format = 44
	FormatB8G8R8A8_SNORM             Format = 45
	FormatB8G8R8A8_USCALED           Format = 46
	FormatB8G8R8A8_SSCALED           Format = 47
	FormatB8G8R8A8_UINT              Format = 48
	FormatB8G8R8A8_SINT              Format = 49
	FormatB8G8R8A8_SRGB              Format = 50
	FormatA8B8G8R8_UNORM_PACK32      Format = 51
	FormatA8B8G8R8_SNORM_PACK32      Format = 52
	FormatA8B8G8R8_USCALED_PACK32    Format = 53
	FormatA8B8G8R8_SSCALED_PACK32    Format = 54
	FormatA8B8G8R8_UINT_PACK32       Format = 55
	FormatA8B8G8R8_SINT_PACK32       Format = 56
	FormatA8B8G8R8_SRGB_PACK32       Format = 57
	FormatA2R10G10B10_UNORM_PACK32   Format = 58
	FormatA2R10G10B10_SNORM_PACK32   Format = 59
	FormatA2R10G10B10_USCALED_PACK32 Format = 60
	FormatA2R10G10B10_SSCALED_PACK32 Format = 61
	FormatA2R10G10B10_UINT_PACK32    Format = 62
	FormatA2R10G10B10_SINT_PACK32    Format = 63
	FormatA2B10G10R10_UNORM_PACK32   Format = 64
	FormatA2B10G10R10_SNORM_PACK32   Format = 65
	FormatA2B10G10R10_USCALED_PACK32 Format = 66
	FormatA2B10G10R10_SSCALED_PACK32 Format = 67
	FormatA2B10G10R10_UINT_PACK32    Format = 68
	FormatA2B10G10R10_SINT_PACK32    Format = 69
	FormatR16_UNORM                  Format = 70
	FormatR16_SNORM                  Format = 71
	FormatR16_USCALED                Format = 72
	FormatR16_SSCALED                Format = 73
	FormatR16_UINT                   Format = 74
	FormatR16_SINT                   Format = 75
	FormatR16_SFLOAT                 Format = 76
	FormatR16G16_UNORM               Format = 77
	FormatR16G16_SNORM               Format = 78
	FormatR16G16_USCALED             Format = 79
	FormatR16G16_SSCALED             Format = 80
	FormatR16G16_UINT                Format = 81
	FormatR16G16_SINT                Format = 82
	FormatR16G16_SFLOAT              Format = 83
	FormatR16G16B16_UNORM            Format = 84
	FormatR16G16B16_SNORM            Format = 85
	FormatR16G16B16_USCALED          Format = 86
	FormatR16G16B16_SSCALED          Format = 87
	FormatR16G16B16_UINT             Format = 88
	FormatR16G16B16_SINT             Format = 89
	FormatR16G16B16_SFLOAT           Format = 90
	FormatR16G16B16A16_UNORM         Format = 91
	FormatR16G16B16A16_SNORM         Format = 92
	FormatR16G16B16A16_USCALED       Format = 93
	FormatR16G16B16A16_SSCALED       Format = 94
	FormatR16G16B16A16_UINT          Format = 95
	FormatR16G16B16A16_SINT          Format = 96
	FormatR16G16B16A16_SFLOAT        Format = 97
	FormatR32_UINT                   Format = 98
	FormatR32_SINT                   Format = 99
	FormatR32_SFLOAT                 Format = 100
	FormatR32G32_UINT                Format = 101
	FormatR32G32_SINT                Format = 102
	FormatR32G32_SFLOAT              Format = 103
	FormatR32G32B32_UINT             Format = 104
	FormatR32G32B32_SINT             Format = 105
	FormatR32G32B32_SFLOAT           Format = 106
	FormatR32G32B32A32_UINT          Format = 107
	FormatR32G32B32A32_SINT          Format = 108
	FormatR32G32B32A32_SFLOAT        Format = 109
	FormatR64_UINT                   Format = 110
	FormatR64_SINT                   Format = 111
	FormatR64_SFLOAT                 Format = 112
	FormatR64G64_UINT                Format = 113
	FormatR64G64_SINT                Format = 114
	FormatR64G64_SFLOAT              Format = 115
	FormatR64G64B64_UINT             Format = 116
	FormatR64G64B64_SINT             Format = 117
	FormatR64G64B64_SFLOAT           Format = 118
	FormatR64G64B64A64_UINT          Format = 119
	FormatR64G64B64A64_SINT          Format = 120
	FormatR64G64B64A64_SFLOAT        Format = 121
	FormatB10G11R11_UFLOAT_PACK32    Format = 122
	FormatE5B9G9R9_UFLOAT_PACK32     Format = 123
	FormatD16_UNORM                  Format = 124
	FormatX8_D24_UNORM_PACK32        Format = 125
	FormatD32_SFLOAT                 Format = 126
	FormatS8_UINT                    Format = 127
	FormatD16_UNORM_S8_UINT          Format = 128
	FormatD24_UNORM_S8_UINT          Format = 129
	FormatD32_SFLOAT_S8_UINT         Format = 130
	FormatBC1_RGB_UNORM_BLOCK        Format = 131
	FormatBC1_RGB_SRGB_BLOCK         Format = 132
	FormatBC1_RGBA_UNORM_BLOCK       Format = 133
	FormatBC1_RGBA_SRGB_BLOCK        Format = 134
	FormatBC2_UNORM_BLOCK            Format = 135
	FormatBC2_SRGB_BLOCK             Format = 136
	FormatBC3_UNORM_BLOCK            Format = 137
	FormatBC3_SRGB_BLOCK             Format = 138
	FormatBC4_UNORM_BLOCK            Format = 139
	FormatBC4_SNORM_BLOCK            Format = 140
	FormatBC5_UNORM_BLOCK            Format = 141
	FormatBC5_SNORM_BLOCK            Format = 142
	FormatBC6H_UFLOAT_BLOCK          Format = 143
	FormatBC6H_SFLOAT_BLOCK          Format = 144
	FormatBC7_UNORM_BLOCK            Format = 145
	FormatBC7_SRGB_BLOCK             Format = 146
	FormatETC2_R8G8B8_UNORM_BLOCK    Format = 147
	FormatETC2_R8G8B8_SRGB_BLOCK     Format = 148
	FormatETC2_R8G8B8A1_UNORM_BLOCK  Format = 149
	FormatETC2_R8G8B8A1_SRGB_BLOCK   Format = 150
	FormatETC2_R8G8B8A8_UNORM_BLOCK  Format = 151
	FormatETC2_R8G8B8A8_SRGB_BLOCK   Format = 152
	FormatEAC_R11_UNORM_BLOCK        Format = 153
	FormatEAC_R11_SNORM_BLOCK        Format = 154
	FormatEAC_R11G11_UNORM_BLOCK     Format = 155
	FormatEAC_R11G11_SNORM_BLOCK     Format = 156
	FormatASTC_4x4_UNORM_BLOCK       Format = 157
	FormatASTC_4x4_SRGB_BLOCK        Format = 158
	FormatASTC_5x4_UNORM_BLOCK       Format = 159
	FormatASTC_5x4_SRGB_BLOCK        Format = 160
	FormatASTC_5x5_UNORM_BLOCK       Format = 161
	FormatASTC_5x5_SRGB_BLOCK        Format = 162
	FormatASTC_6x5_UNORM_BLOCK       Format = 163
	FormatASTC_6x5_SRGB_BLOCK        Format = 164
	FormatASTC_6x6_UNORM_BLOCK       Format = 165
	FormatASTC_6x6_SRGB_BLOCK        Format = 166
	FormatASTC_8x5_UNORM_BLOCK       Format = 167
	FormatASTC_8x5_SRGB_BLOCK        Format = 168
	FormatASTC_8x6_UNORM_BLOCK       Format = 169
	FormatASTC_8x6_SRGB_BLOCK        Format = 170
	FormatASTC_8x8_UNORM_BLOCK       Format = 171
	FormatASTC_8x8_SRGB_BLOCK        Format = 172
	FormatASTC_10x5_UNORM_BLOCK      Format = 173
	FormatASTC_10x5_SRGB_BLOCK       Format = 174
	FormatASTC_10x6_UNORM_BLOCK      Format = 175
	FormatASTC_10x6_SRGB_BLOCK       Format = 176
	FormatASTC_10x8_UNORM_BLOCK      Format = 177
	FormatASTC_10x8_SRGB_BLOCK       Format = 178
	FormatASTC_10x10_UNORM_BLOCK     Format = 179
	FormatASTC_10x10_SRGB_BLOCK      Format = 180
	FormatASTC_12x10_UNORM_BLOCK     Format = 181
	FormatASTC_12x10_SRGB_BLOCK      Format = 182
	FormatASTC_12x12_UNORM_BLOCK     Format = 183
	FormatASTC_12x12_SRGB_BLOCK      Format = 184
)

const (
	COLOR_SPACE_SRGB_NONLINEAR_KHR ColorSpaceKHR = 0
)

const (
	// PresentModeImmediateKHR specifies that the presentation engine does not wait for a vertical blanking period to update the current image,
	// meaning this mode may result in visible tearing.
	// No internal queuing of presentation requests is needed, as the requests are applied immediately.
	PresentModeImmediateKHR PresentModeKHR = 0

	// PresentModeMailboxKHR specifies that the presentation engine waits for the next vertical blanking period to update the current image.
	// Tearing cannot be observed. An internal single-entry queue is used to hold pending presentation requests.
	// If the queue is full when a new presentation request is received, the new request replaces the existing entry,
	// and any images associated with the prior entry become available for re-use by the application.
	// One request is removed from the queue and processed during each vertical blanking period in which the queue is non-empty.
	PresentModeMailboxKHR PresentModeKHR = 1

	// PresentModeFifoKHR specifies that the presentation engine waits for the next vertical blanking period to update the current image.
	// Tearing cannot be observed. An internal queue is used to hold pending presentation requests.
	// New requests are appended to the end of the queue,
	// and one request is removed from the beginning of the queue
	// and processed during each vertical blanking period in which the queue is non-empty.
	//
	// This is the only value of presentMode that is required to be supported.
	PresentModeFifoKHR PresentModeKHR = 2

	// PresentModeFifoRelaxedKHR specifies that the presentation engine generally waits for the next vertical blanking period to update the current image.
	// If a vertical blanking period has already passed since the last update of the current image
	// then the presentation engine does not wait for another vertical blanking period for the update,
	// meaning this mode may result in visible tearing in this case.
	//
	// This mode is useful for reducing visual stutter with an application that will mostly present a new image before the next vertical blanking period,
	// but may occasionally be late, and present a new image just after the next vertical blanking period.
	// An internal queue is used to hold pending presentation requests.
	// New requests are appended to the end of the queue,
	// and one request is removed from the beginning of the queue and processed during
	// or after each vertical blanking period in which the queue is non-empty.
	PresentModeFifoRelaxedKHR PresentModeKHR = 3

	// PresentModeSharedDemandRefreshKHR specifies that the presentation engine and application have concurrent access to a single image,
	// which is referred to as a shared presentable image.
	// The presentation engine is only required to update the current image after a new presentation request is received.
	// Therefore the application must make a presentation request whenever an update is required.
	// However, the presentation engine may update the current image at any point, meaning this mode may result in visible tearing.
	PresentModeSharedDemandRefreshKHR PresentModeKHR = 1000111000

	// PresentModeSharedContinuousRefreshKHR specifies that the presentation engine and application have concurrent access to a single image,
	// which is referred to as a shared presentable image.
	// The presentation engine periodically updates the current image on its regular refresh cycle.
	// The application is only required to make one initial presentation request,
	// after which the presentation engine must update the current image without any need for further presentation requests.
	// The application can indicate the image contents have been updated by making a presentation request,
	// but this does not guarantee the timing of when it will be updated.
	// This mode may result in visible tearing if rendering to the image is not timed correctly.
	PresentModeSharedContinuousRefreshKHR PresentModeKHR = 1000111001
)

const (
	// CommandPoolCreateTransientBit specifies that command buffers allocated from the pool will be short-lived,
	// meaning that they will be reset or freed in a relatively short timeframe.
	// This flag may be used by the implementation to control memory allocation behavior within the pool.
	CommandPoolCreateTransientBit CommandPoolCreateFlags = 0x00000001

	// CommandPoolCreateResetCommandBufferBit allows any command buffer allocated from a pool
	// to be individually reset to the initial state;
	// either by calling ResetCommandBuffer, or via the implicit reset when calling BeginCommandBuffer.
	// If this flag is not set on a pool, then ResetCommandBuffer must not be called for any command buffer allocated from that pool.
	CommandPoolCreateResetCommandBufferBit CommandPoolCreateFlags = 0x00000002

	// CommandPoolCreateProtectedBit specifies that command buffers allocated from the pool are protected command buffers.
	// If the protected memory feature is not enabled, the CommandPoolCreateProtectedBit bit of flags must not be set.
	CommandPoolCreateProtectedBit CommandPoolCreateFlags = 0x00000004
)

const (
	// CommandlPoolResetReleaseResourcesBit specifies that resetting a command pool
	// recycles all of the resources from the command pool back to the system.
	CommandlPoolResetReleaseResourcesBit CommandPoolResetFlags = 0x00000001
)

const (
	// CommandBufferLevelPrimary specifies a primary command buffer.
	CommandBufferLevelPrimary CommandBufferLevel = 0

	// CommandBufferLevelSecondary specifies a secondary command buffer.
	CommandBufferLevelSecondary CommandBufferLevel = 1
)

const (
	// CommandBufferResetReleaseResourcesBit specifies that most or all memory resources
	// currently owned by the command buffer should be returned to the parent command pool.
	// If this flag is not set, then the command buffer may hold onto memory resources
	// and reuse them when recording commands. commandBuffer is moved to the initial state.
	CommandBufferResetReleaseResourcesBit CommandBufferResetFlags = 0x00000001
)

const (
	// CommandBufferUsageOneTimeSubmitBit specifies that each recording of the command buffer will only be submitted once,
	// and the command buffer will be reset and recorded again between each submission.
	CommandBufferUsageOneTimeSubmitBit CommandBufferUsageFlags = 0x00000001

	// CommandBufferUsageRenderPassContinueBit specifies that a secondary command buffer is considered to be entirely inside a render pass.
	// If this is a primary command buffer, then this bit is ignored.
	CommandBufferUsageRenderPassContinueBit CommandBufferUsageFlags = 0x00000002

	// CommandBufferUsageSimultaneousUseBit specifies that a command buffer can be resubmitted to a queue while it is in the pending state,
	// and recorded into multiple primary command buffers.
	CommandBufferUsageSimultaneousUseBit CommandBufferUsageFlags = 0x00000004
)

const (
	// QueryPipelineStatisticInputAssemblyVerticesBit specifies that queries managed by the pool
	// will count the number of vertices processed by the input assembly stage.
	// Vertices corresponding to incomplete primitives may contribute to the count.
	QueryPipelineStatisticInputAssemblyVerticesBit QueryPipelineStatisticFlags = 0x00000001

	// QueryPipelineStatisticInputAssemblyPrimitivesBit specifies that queries managed by the pool will
	// count the number of primitives processed by the input assembly stage.
	// If primitive restart is enabled, restarting the primitive topology has no effect on the count.
	// Incomplete primitives may be counted.
	QueryPipelineStatisticInputAssemblyPrimitivesBit QueryPipelineStatisticFlags = 0x00000002

	// QueryPipelineStatisticVertexShaderInvocationsBit specifies that queries managed by the pool will
	// count the number of vertex shader invocations.
	// This counter’s value is incremented each time a vertex shader is invoked.
	QueryPipelineStatisticVertexShaderInvocationsBit QueryPipelineStatisticFlags = 0x00000004

	// QueryPipelineStatisticGeometryShaderInvocationsBit specifies that queries managed by the pool will
	// count the number of geometry shader invocations.
	// This counter’s value is incremented each time a geometry shader is invoked.
	// In the case of instanced geometry shaders,
	// the geometry shader invocations count is incremented for each separate instanced invocation.
	QueryPipelineStatisticGeometryShaderInvocationsBit QueryPipelineStatisticFlags = 0x00000008

	// QueryPipelineStatisticGeometryShaderPrimitivesBit specifies that queries managed by the pool will
	// count the number of primitives generated by geometry shader invocations.
	// The counter’s value is incremented each time the geometry shader emits a primitive.
	// Restarting primitive topology using the SPIR-V instructions OpEndPrimitive or OpEndStreamPrimitive
	// has no effect on the geometry shader output primitives count.
	QueryPipelineStatisticGeometryShaderPrimitivesBit QueryPipelineStatisticFlags = 0x00000010

	// QueryPipelineStatisticClippingInvocationsBit specifies that queries managed by the pool will
	// count the number of primitives processed by the Primitive Clipping stage of the pipeline.
	// The counter’s value is incremented each time a primitive reaches the primitive clipping stage.
	QueryPipelineStatisticClippingInvocationsBit QueryPipelineStatisticFlags = 0x00000020

	// QueryPipelineStatisticClippingPrimitivesBit specifies that queries managed by the pool will
	// count the number of primitives output by the Primitive Clipping stage of the pipeline.
	// The counter’s value is incremented each time a primitive passes the primitive clipping stage.
	// The actual number of primitives output by the primitive clipping stage for a particular input primitive
	// is implementation-dependent but must satisfy the following conditions:
	//
	// - If at least one vertex of the input primitive lies inside the clipping volume, the counter is incremented by one or more.
	// - Otherwise, the counter is incremented by zero or more.
	QueryPipelineStatisticClippingPrimitivesBit QueryPipelineStatisticFlags = 0x00000040

	// QueryPipelineStatisticFragmentShaderInvocationsBit specifies that queries managed by the pool will
	// count the number of fragment shader invocations.
	// The counter’s value is incremented each time the fragment shader is invoked.
	QueryPipelineStatisticFragmentShaderInvocationsBit QueryPipelineStatisticFlags = 0x00000080

	// QueryPipelineStatisticTessellationControlShaderPatchesBit specifies that queries managed by the pool will
	// count the number of patches processed by the tessellation control shader.
	// The counter’s value is incremented once for each patch for which a tessellation control shader is invoked.
	QueryPipelineStatisticTessellationControlShaderPatchesBit QueryPipelineStatisticFlags = 0x00000100

	// QueryPipelineStatisticTessellationEvaluationShaderInvocationsBit specifies that queries managed by the pool will
	// count the number of invocations of the tessellation evaluation shader.
	// The counter’s value is incremented each time the tessellation evaluation shader is invoked.
	QueryPipelineStatisticTessellationEvaluationShaderInvocationsBit QueryPipelineStatisticFlags = 0x00000200

	// QueryPipelineStatisticComputeShaderInvocationsBit specifies that queries managed by the pool will
	// count the number of compute shader invocations.
	// The counter’s value is incremented every time the compute shader is invoked.
	// Implementations may skip the execution of certain compute shader invocations or execute additional compute shader invocations
	// for implementation-dependent reasons as long as the results of rendering otherwise remain unchanged.
	QueryPipelineStatisticComputeShaderInvocationsBit QueryPipelineStatisticFlags = 0x00000400
)

const (
	// QueryControlPreciseBit specifies the precision of occlusion queries.
	QueryControlPreciseBit QueryControlFlags = 0x00000001
)

const (
	// SharingModeExclusive specifies that access to any range or image subresource
	// of the object will be exclusive to a single queue family at a time.
	SharingModeExclusive SharingMode = 0
	// SharingModeConcurrent specifies that concurrent access to any range or image subresource
	// of the object from multiple queue families is supported.
	SharingModeConcurrent SharingMode = 1
)

const (
	ImageViewType1D        ImageViewType = 0
	ImageViewType2D        ImageViewType = 1
	ImageViewType3D        ImageViewType = 2
	ImageViewTypeCube      ImageViewType = 3
	ImageViewType1DArray   ImageViewType = 4
	ImageViewType2DArray   ImageViewType = 5
	ImageViewTypeCubeArray ImageViewType = 6
)

const (
	ComponentSwizzleIdentity ComponentSwizzle = 0
	ComponentSwizzleZero     ComponentSwizzle = 1
	ComponentSwizzleOne      ComponentSwizzle = 2
	ComponentSwizzleR        ComponentSwizzle = 3
	ComponentSwizzleG        ComponentSwizzle = 4
	ComponentSwizzleB        ComponentSwizzle = 5
	ComponentSwizzleA        ComponentSwizzle = 6
)

const (
	ImageAspectColorBit           ImageAspectFlags = 0x00000001
	ImageAspectDepthBit           ImageAspectFlags = 0x00000002
	ImageAspectStencilBit         ImageAspectFlags = 0x00000004
	ImageAspectMetadataBit        ImageAspectFlags = 0x00000008
	ImageAspectPlane0Bit          ImageAspectFlags = 0x00000010
	ImageAspectPlane1Bit          ImageAspectFlags = 0x00000020
	ImageAspectPlane2Bit          ImageAspectFlags = 0x00000040
	ImageAspectMemoryPlane0BitEXT ImageAspectFlags = 0x00000080
	ImageAspectMemoryPlane1BitEXT ImageAspectFlags = 0x00000100
	ImageAspectMemoryPlane2BitEXT ImageAspectFlags = 0x00000200
	ImageAspectMemoryPlane3BitEXT ImageAspectFlags = 0x00000400
	ImageAspectPlane0BitKHR       ImageAspectFlags = ImageAspectPlane0Bit
	ImageAspectPlane1BitKHR       ImageAspectFlags = ImageAspectPlane1Bit
	ImageAspectPlane2BitKHR       ImageAspectFlags = ImageAspectPlane2Bit
)

const (
	ShaderStageVertexBit                 ShaderStageFlags = 0x00000001
	ShaderStageTessellationControlBit    ShaderStageFlags = 0x00000002
	ShaderStageTessellationEvaluationBit ShaderStageFlags = 0x00000004
	ShaderStageGeometryBit               ShaderStageFlags = 0x00000008
	ShaderStageFragmentBit               ShaderStageFlags = 0x00000010
	ShaderStageComputeBit                ShaderStageFlags = 0x00000020
	ShaderStageAllGraphics               ShaderStageFlags = 0x0000001F
	ShaderStageAll                       ShaderStageFlags = 0x7FFFFFFF
	ShaderStageRaygenBitNVX              ShaderStageFlags = 0x00000100
	ShaderStageAnyHitBitNVX              ShaderStageFlags = 0x00000200
	ShaderStageClosestHitBitNVX          ShaderStageFlags = 0x00000400
	ShaderStageMissBitNVX                ShaderStageFlags = 0x00000800
	ShaderStageIntersectionBitNVX        ShaderStageFlags = 0x00001000
	ShaderStageCallableBitNVX            ShaderStageFlags = 0x00002000
	ShaderStageTaskBitNV                 ShaderStageFlags = 0x00000040
	ShaderStageMeshBitNV                 ShaderStageFlags = 0x00000080
)

const (
	VertexInputRateVertex   VertexInputRate = 0
	VertexInputRateInstance VertexInputRate = 1
)

const (
	PrimitiveTopologyPointList                  PrimitiveTopology = 0
	PrimitiveTopologyLineList                   PrimitiveTopology = 1
	PrimitiveTopologyLineStrip                  PrimitiveTopology = 2
	PrimitiveTopologyTriangleList               PrimitiveTopology = 3
	PrimitiveTopologyTriangleStrip              PrimitiveTopology = 4
	PrimitiveTopologyTriangleFan                PrimitiveTopology = 5
	PrimitiveTopologyLineListWithAdjacency      PrimitiveTopology = 6
	PrimitiveTopologyLineStripWithAdjacency     PrimitiveTopology = 7
	PrimitiveTopologyTriangleListWithAdjacency  PrimitiveTopology = 8
	PrimitiveTopologyTriangleStripWithAdjacency PrimitiveTopology = 9
	PrimitiveTopologyPatchList                  PrimitiveTopology = 10
)

const (
	PolygonModeFill            PolygonMode = 0
	PolygonModeLine            PolygonMode = 1
	PolygonModePoint           PolygonMode = 2
	PolygonModeFillRectangleNV PolygonMode = 1000153000
)

const (
	CullModeNone         CullModeFlags = 0
	CullModeFrontBit     CullModeFlags = 0x00000001
	CullModeBackBit      CullModeFlags = 0x00000002
	CullModeFrontAndBack CullModeFlags = 0x00000003
)

const (
	FrontFaceCounterClockwise FrontFace = 0
	FrontFaceClockwise        FrontFace = 1
)

const (
	SampleCount1Bit  SampleCountFlags = 0x00000001
	SampleCount2Bit  SampleCountFlags = 0x00000002
	SampleCount4Bit  SampleCountFlags = 0x00000004
	SampleCount8Bit  SampleCountFlags = 0x00000008
	SampleCount16Bit SampleCountFlags = 0x00000010
	SampleCount32Bit SampleCountFlags = 0x00000020
	SampleCount64Bit SampleCountFlags = 0x00000040
)

const (
	BlendFactorZero                  BlendFactor = 0
	BlendFactorOne                   BlendFactor = 1
	BlendFactorSrcColor              BlendFactor = 2
	BlendFactorOneMinusSrcColor      BlendFactor = 3
	BlendFactorDstColor              BlendFactor = 4
	BlendFactorOneMinusDstColor      BlendFactor = 5
	BlendFactorSrcAlpha              BlendFactor = 6
	BlendFactorOneMinusSrcAlpha      BlendFactor = 7
	BlendFactorDstAlpha              BlendFactor = 8
	BlendFactorOneMinusDstAlpha      BlendFactor = 9
	BlendFactorConstantColor         BlendFactor = 10
	BlendFactorOneMinusConstantColor BlendFactor = 11
	BlendFactorConstantAlpha         BlendFactor = 12
	BlendFactorOneMinusConstantAlpha BlendFactor = 13
	BlendFactorSrcAlphaSaturate      BlendFactor = 14
	BlendFactorSrc1Color             BlendFactor = 15
	BlendFactorOneMinusSrc1Color     BlendFactor = 16
	BlendFactorSrc1Alpha             BlendFactor = 17
	BlendFactorOneMinusSrc1Alpha     BlendFactor = 18
)

const (
	BlendOpAdd                 BlendOp = 0
	BlendOpSubtract            BlendOp = 1
	BlendOpReverseSubtract     BlendOp = 2
	BlendOpMin                 BlendOp = 3
	BlendOpMax                 BlendOp = 4
	BlendOpZeroEXT             BlendOp = 1000148000
	BlendOpSrcEXT              BlendOp = 1000148001
	BlendOpDstEXT              BlendOp = 1000148002
	BlendOpSrcOverEXT          BlendOp = 1000148003
	BlendOpDstOverEXT          BlendOp = 1000148004
	BlendOpSrcInEXT            BlendOp = 1000148005
	BlendOpDstInEXT            BlendOp = 1000148006
	BlendOpSrcOutEXT           BlendOp = 1000148007
	BlendOpDstOutEXT           BlendOp = 1000148008
	BlendOpSrcAtopEXT          BlendOp = 1000148009
	BlendOpDstAtopEXT          BlendOp = 1000148010
	BlendOpXorEXT              BlendOp = 1000148011
	BlendOpMultiplyEXT         BlendOp = 1000148012
	BlendOpScreenEXT           BlendOp = 1000148013
	BlendOpOverlayEXT          BlendOp = 1000148014
	BlendOpDarkenEXT           BlendOp = 1000148015
	BlendOpLightenEXT          BlendOp = 1000148016
	BlendOpColordodgeEXT       BlendOp = 1000148017
	BlendOpColorburnEXT        BlendOp = 1000148018
	BlendOpHardlightEXT        BlendOp = 1000148019
	BlendOpSoftlightEXT        BlendOp = 1000148020
	BlendOpDifferenceEXT       BlendOp = 1000148021
	BlendOpExclusionEXT        BlendOp = 1000148022
	BlendOpInvertEXT           BlendOp = 1000148023
	BlendOpInvertRGB_EXT       BlendOp = 1000148024
	BlendOpLineardodgeEXT      BlendOp = 1000148025
	BlendOpLinearburnEXT       BlendOp = 1000148026
	BlendOpVividlightEXT       BlendOp = 1000148027
	BlendOpLinearlightEXT      BlendOp = 1000148028
	BlendOpPinlightEXT         BlendOp = 1000148029
	BlendOpHardmixEXT          BlendOp = 1000148030
	BlendOpHSLHueEXT           BlendOp = 1000148031
	BlendOpHSLSaturationEXT    BlendOp = 1000148032
	BlendOpHSLColorEXT         BlendOp = 1000148033
	BlendOpHSLLuminosityEXT    BlendOp = 1000148034
	BlendOpPlusEXT             BlendOp = 1000148035
	BlendOpPlusClampedEXT      BlendOp = 1000148036
	BlendOpPlusClampedAlphaEXT BlendOp = 1000148037
	BlendOpPlusDarkerEXT       BlendOp = 1000148038
	BlendOpMinusEXT            BlendOp = 1000148039
	BlendOpMinusClampedEXT     BlendOp = 1000148040
	BlendOpContrastEXT         BlendOp = 1000148041
	BlendOpInvertOVG_EXT       BlendOp = 1000148042
	BlendOpRedEXT              BlendOp = 1000148043
	BlendOpGreenEXT            BlendOp = 1000148044
	BlendOpBlueEXT             BlendOp = 1000148045
)

const (
	ColorComponentR ColorComponentFlags = 0x00000001
	ColorComponentG ColorComponentFlags = 0x00000002
	ColorComponentB ColorComponentFlags = 0x00000004
	ColorComponentA ColorComponentFlags = 0x00000008
)

const (
	LogicOpClear        LogicOp = 0
	LogicOpAnd          LogicOp = 1
	LogicOpAndReverse   LogicOp = 2
	LogicOpCopy         LogicOp = 3
	LogicOpAndInverted  LogicOp = 4
	LogicOpNoOp         LogicOp = 5
	LogicOpXor          LogicOp = 6
	LogicOpOr           LogicOp = 7
	LogicOpNor          LogicOp = 8
	LogicOpEquivalent   LogicOp = 9
	LogicOpInvert       LogicOp = 10
	LogicOpOrReverse    LogicOp = 11
	LogicOpCopyInverted LogicOp = 12
	LogicOpOrInverted   LogicOp = 13
	LogicOpNand         LogicOp = 14
	LogicOpSet          LogicOp = 15
)

const (
	DynamicStateViewport                     DynamicState = 0
	DynamicStateScissor                      DynamicState = 1
	DynamicStateLineWidth                    DynamicState = 2
	DynamicStateDepthBias                    DynamicState = 3
	DynamicStateBlendConstants               DynamicState = 4
	DynamicStateDepthBounds                  DynamicState = 5
	DynamicStateStencilCompareMask           DynamicState = 6
	DynamicStateStencilWriteMask             DynamicState = 7
	DynamicStateStencilReference             DynamicState = 8
	DynamicStateViewportWScalingNV           DynamicState = 1000087000
	DynamicStateDiscardRectangleEXT          DynamicState = 1000099000
	DynamicStateSampleLocationsEXT           DynamicState = 1000143000
	DynamicStateViewportShadingRatePaletteNV DynamicState = 1000164004
	DynamicStateViewportCoarseSampleOrderNV  DynamicState = 1000164006
	DynamicStateExclusiveScissorNV           DynamicState = 1000205001
)

const (
	PipelineCreateDisableOptimizationBit         PipelineCreateFlags = 0x00000001
	PipelineCreateAllowDerivativesBit            PipelineCreateFlags = 0x00000002
	PipelineCreateDerivativeBit                  PipelineCreateFlags = 0x00000004
	PipelineCreateViewIndexFromDeviceIndexBit    PipelineCreateFlags = 0x00000008
	PipelineCreateDispatchBase                   PipelineCreateFlags = 0x00000010
	PipelineCreateDeferCompileBitNVX             PipelineCreateFlags = 0x00000020
	PipelineCreateViewIndexFromDeviceIndexBitKHR PipelineCreateFlags = PipelineCreateViewIndexFromDeviceIndexBit
	PipelineCreateDispatchBaseKHR                PipelineCreateFlags = PipelineCreateDispatchBase
)

const (
	CompareOpNever          CompareOp = 0
	CompareOpLess           CompareOp = 1
	CompareOpEqual          CompareOp = 2
	CompareOpLessOrEqual    CompareOp = 3
	CompareOpGreater        CompareOp = 4
	CompareOpNotEqual       CompareOp = 5
	CompareOpGreaterOrEqual CompareOp = 6
	CompareOpAlways         CompareOp = 7
)

const (
	StencilOpKeep              StencilOp = 0
	StencilOpZero              StencilOp = 1
	StencilOpReplace           StencilOp = 2
	StencilOpIncrementAndClamp StencilOp = 3
	StencilOpDecrementAndClamp StencilOp = 4
	StencilOpInvert            StencilOp = 5
	StencilOpIncrementAndWrap  StencilOp = 6
	StencilOpDecrementAndWrap  StencilOp = 7
)

const (
	AttachmentDescriptionMayAliasBit AttachmentDescriptionFlags = 0x00000001
)

const (
	AttachmentLoadOpLoad     AttachmentLoadOp = 0
	AttachmentLoadOpClear    AttachmentLoadOp = 1
	AttachmentLoadOpDontCare AttachmentLoadOp = 2
)

const (
	AttachmentStoreOpStore    AttachmentStoreOp = 0
	AttachmentStoreOpDontCare AttachmentStoreOp = 1
)

const (
	ImageLayoutUndefined                                ImageLayout = 0
	ImageLayoutGeneral                                  ImageLayout = 1
	ImageLayoutColorAttachmentOptimal                   ImageLayout = 2
	ImageLayoutDepthStencilAttachmentOptimal            ImageLayout = 3
	ImageLayoutDepthStencilReadOnlyOptimal              ImageLayout = 4
	ImageLayoutShaderReadOnlyOptimal                    ImageLayout = 5
	ImageLayoutTransferSrcOptimal                       ImageLayout = 6
	ImageLayoutTransferDstOptimal                       ImageLayout = 7
	ImageLayoutPreinitialized                           ImageLayout = 8
	ImageLayoutDepthReadOnlyStencilAttachmentOptimal    ImageLayout = 1000117000
	ImageLayoutDepthAttachmentStencilReadOnlyOptimal    ImageLayout = 1000117001
	ImageLayoutPresentSrcKHR                            ImageLayout = 1000001002
	ImageLayoutSharedPresentKHR                         ImageLayout = 1000111000
	ImageLayoutShadingRateOptimalNV                     ImageLayout = 1000164003
	ImageLayoutDepthReadOnlyStencilAttachmentOptimalKHR ImageLayout = ImageLayoutDepthReadOnlyStencilAttachmentOptimal
	ImageLayoutDepthAttachmentStencilReadOnlyOptimalKHR ImageLayout = ImageLayoutDepthAttachmentStencilReadOnlyOptimal
)

const (
	PipelineBindPointGraphics      PipelineBindPoint = 0
	PipelineBindPointCompute       PipelineBindPoint = 1
	PipelineBindPointRaytracingNVX PipelineBindPoint = 1000165000
)

const (
	SubpassDescriptionPerViewAttributesBitNVX    SubpassDescriptionFlags = 0x00000001
	SubpassDescriptionPerViewPositionXOnlyBitNVX SubpassDescriptionFlags = 0x00000002
)

const (
	PipelineStageTopOfPipeBit                    PipelineStageFlags = 0x00000001
	PipelineStageDrawIndirectBit                 PipelineStageFlags = 0x00000002
	PipelineStageVertexInputBit                  PipelineStageFlags = 0x00000004
	PipelineStageVertexShaderBit                 PipelineStageFlags = 0x00000008
	PipelineStageTessellationControlShaderBit    PipelineStageFlags = 0x00000010
	PipelineStageTessellationEvaluationShaderBit PipelineStageFlags = 0x00000020
	PipelineStageGeometryShaderBit               PipelineStageFlags = 0x00000040
	PipelineStageFragmentShaderBit               PipelineStageFlags = 0x00000080
	PipelineStageEarlyFragmentTestsBit           PipelineStageFlags = 0x00000100
	PipelineStageLateFragmentTestsBit            PipelineStageFlags = 0x00000200
	PipelineStageColorAttachmentOutputBit        PipelineStageFlags = 0x00000400
	PipelineStageComputeShaderBit                PipelineStageFlags = 0x00000800
	PipelineStageTransferBit                     PipelineStageFlags = 0x00001000
	PipelineStageBottomOfPipeBit                 PipelineStageFlags = 0x00002000
	PipelineStageHostBit                         PipelineStageFlags = 0x00004000
	PipelineStageAllGraphicsBit                  PipelineStageFlags = 0x00008000
	PipelineStageAllCommandsBit                  PipelineStageFlags = 0x00010000
	PipelineStageTransformFeedbackBitEXT         PipelineStageFlags = 0x01000000
	PipelineStageConditionalRenderingBitEXT      PipelineStageFlags = 0x00040000
	PipelineStageCommandProcessBitNVX            PipelineStageFlags = 0x00020000
	PipelineStageShadingRateImageBitNV           PipelineStageFlags = 0x00400000
	PipelineStageRaytracingBitNVX                PipelineStageFlags = 0x00200000
	PipelineStageTaskShaderBitNV                 PipelineStageFlags = 0x00080000
	PipelineStageMeshShaderBitNV                 PipelineStageFlags = 0x00100000
)

const (
	// AccessIndirectCommandReadBit specifies read access to indirect command data read as part of an indirect drawing or dispatch command.
	AccessIndirectCommandReadBit AccessFlags = 0x00000001

	// AccessIndexReadBit specifies read access to an index buffer as part of an indexed drawing command, bound by BindIndexBuffer.
	AccessIndexReadBit AccessFlags = 0x00000002

	// AccessVertexAttributeReadBit specifies read access to a vertex buffer as part of a drawing command, bound by BindVertexBuffers.
	AccessVertexAttributeReadBit AccessFlags = 0x00000004

	// AccessUniformReadBit specifies read access to a uniform buffer.
	AccessUniformReadBit AccessFlags = 0x00000008

	// AccessInputAttachmentReadBit specifies read access to an input attachment within a render pass during fragment shading.
	AccessInputAttachmentReadBit AccessFlags = 0x00000010

	// AccessShaderReadBit specifies read access to a storage buffer, uniform texel buffer, storage texel buffer, sampled image, or storage image.
	AccessShaderReadBit AccessFlags = 0x00000020

	// AccessShaderWriteBit specifies write access to a storage buffer, storage texel buffer, or storage image.
	AccessShaderWriteBit AccessFlags = 0x00000040

	// AccessColorAttachmentReadBit specifies read access to a color attachment, such as via blending, logic operations, or via certain subpass load operations.
	// It does not include advanced blend operations.
	AccessColorAttachmentReadBit AccessFlags = 0x00000080

	// AccessColorAttachmentWriteBit specifies write access to a color or resolve attachment during a render pass or via certain subpass load and store operations.
	AccessColorAttachmentWriteBit AccessFlags = 0x00000100

	// AccessDepthStencilAttachmentReadBit specifies read access to a depth/stencil attachment, via depth or stencil operations or via certain subpass load operations.
	AccessDepthStencilAttachmentReadBit AccessFlags = 0x00000200

	// AccessDepthStencilAttachmentWriteBit specifies write access to a depth/stencil attachment, via depth or stencil operations or via certain subpass load and store operations.
	AccessDepthStencilAttachmentWriteBit AccessFlags = 0x00000400

	// AccessTransferReadBit specifies read access to an image or buffer in a copy operation.
	AccessTransferReadBit AccessFlags = 0x00000800

	// AccessTransferWriteBit specifies write access to an image or buffer in a clear or copy operation.
	AccessTransferWriteBit AccessFlags = 0x00001000

	// AccessHostReadBit specifies read access by a host operation.
	// Accesses of this type are not performed through a resource, but directly on memory.
	AccessHostReadBit AccessFlags = 0x00002000

	// AccessHostWriteBit specifies write access by a host operation.
	// Accesses of this type are not performed through a resource, but directly on memory.
	AccessHostWriteBit AccessFlags = 0x00004000

	// AccessMemoryReadBit specifies read access via non-specific entities.
	// These entities include the Vulkan device and host, but may also include entities external to the Vulkan device or otherwise not part of the core Vulkan pipeline.
	// When included in a destination access mask, makes all available writes visible to all future read accesses on entities known to the Vulkan device.
	AccessMemoryReadBit AccessFlags = 0x00008000

	// AccessMemoryWriteBit specifies write access via non-specific entities.
	// These entities include the Vulkan device and host, but may also include entities external to the Vulkan device or otherwise not part of the core Vulkan pipeline.
	// When included in a source access mask, all writes that are performed by entities known to the Vulkan device are made available.
	// When included in a destination access mask, makes all available writes visible to all future write accesses on entities known to the Vulkan device.
	AccessMemoryWriteBit AccessFlags = 0x00010000

	// AccessTransformFeedbackWriteBitEXT specifies write access to a transform feedback buffer made when transform feedback is active.
	AccessTransformFeedbackWriteBitEXT AccessFlags = 0x02000000

	// AccessTransformFeedbackCounterReadBitEXT specifies read access to a transform feedback counter buffer which is read when BeginTransformFeedbackEXT executes.
	AccessTransformFeedbackCounterReadBitEXT AccessFlags = 0x04000000

	// AccessTransformFeedbackCounterWriteBitEXT specifies write access to a transform feedback counter buffer which is written when EndTransformFeedbackEXT executes.
	AccessTransformFeedbackCounterWriteBitEXT AccessFlags = 0x08000000

	// AccessConditionalRenderingReadBitEXT specifies read access to a predicate as part of conditional rendering.
	AccessConditionalRenderingReadBitEXT AccessFlags = 0x00100000

	// AccessCommandProcessReadBitNVX specifies reads from Buffer inputs to ProcessCommandsNVX.
	AccessCommandProcessReadBitNVX AccessFlags = 0x00020000

	// AccessCommandProcessWriteBitNVX specifies writes to the target command buffer in ProcessCommandsNVX.
	AccessCommandProcessWriteBitNVX AccessFlags = 0x00040000

	// AccessColorAttachmentReadNoncoherentBitEXT is similar to AccessColorAttachmentReadBit, but also includes advanced blend operations.
	AccessColorAttachmentReadNoncoherentBitEXT AccessFlags = 0x00080000

	// AccessShadingRateImageReadBitNV specifies read access to a shading rate image as part of a drawing command, as bound by BindShadingRateImageNV.
	AccessShadingRateImageReadBitNV AccessFlags = 0x00800000

	AccessAccelerationStructureReadBitNVX  AccessFlags = 0x00200000
	AccessAccelerationStructureWriteBitNVX AccessFlags = 0x00400000
)

const (
	DependencyByRegionBit       DependencyFlags = 0x00000001
	DependencyDeviceGroupBit    DependencyFlags = 0x00000004
	DependencyViewLocalBit      DependencyFlags = 0x00000002
	DependencyViewLocalBitKHR   DependencyFlags = DependencyViewLocalBit
	DependencyDeviceGroupBitKHR DependencyFlags = DependencyDeviceGroupBit
)

const (
	SubpassContentsInline                  SubpassContents = 0
	SubpassContentsSecondaryCommandBuffers SubpassContents = 1
)

const (
	FenceCreateSignaledBit FenceCreateFlags = 0x00000001
)

const (
	BufferCreateSparseBindingBit   BufferCreateFlags = 0x00000001
	BufferCreateSparseResidencyBit BufferCreateFlags = 0x00000002
	BufferCreateSparseAliasedBit   BufferCreateFlags = 0x00000004
	BufferCreateProtectedBit       BufferCreateFlags = 0x00000008
)

const (
	// BufferUsageTransferSrcBit specifies that the buffer can be used as the source of a transfer command
	// (see the definition of PipelineStageTransferBit).
	BufferUsageTransferSrcBit BufferUsageFlags = 0x00000001

	// BufferUsageTransferDstBit specifies that the buffer can be used as the destination of a transfer command.
	BufferUsageTransferDstBit BufferUsageFlags = 0x00000002

	// BufferUsageUniformTexelBufferBit specifies that the buffer can be used to create a BufferView suitable for
	// occupying a DescriptorSet slot of type DescriptorTypeUniformTexelBuffer.
	BufferUsageUniformTexelBufferBit BufferUsageFlags = 0x00000004

	// BufferUsageStorageTexelBufferBit specifies that the buffer can be used to create a BufferView suitable for
	// occupying a DescriptorSet slot of type DescriptorTypeStorageTexelBUFFER.
	BufferUsageStorageTexelBufferBit BufferUsageFlags = 0x00000008

	// BufferUsageUniformBufferBit specifies that the buffer can be used in a DescriptorBufferInfo suitable for
	// occupying a DescriptorSet slot either of type DescriptorTypeUniformBuffer or DescriptorTypeUniformBufferDynamic.
	BufferUsageUniformBufferBit BufferUsageFlags = 0x00000010

	// BufferUsageStorageBufferBit specifies that the buffer can be used in a DescriptorBufferInfo suitable for occupying
	// a DescriptorSet slot either of type DescriptorTypeStorageBuffer or DescriptorTypeStorageBufferDynamic.
	BufferUsageStorageBufferBit BufferUsageFlags = 0x00000020

	// BufferUsageIndexBufferBit specifies that the buffer is suitable for passing as the buffer parameter to BindIndexBuffer.
	BufferUsageIndexBufferBit BufferUsageFlags = 0x00000040

	// BufferUsageVertexBufferBit specifies that the buffer is suitable for passing as an element of the pBuffers array to BindVertexBuffers.
	BufferUsageVertexBufferBit BufferUsageFlags = 0x00000080

	// BufferUsageIndirectBufferBit specifies that the buffer is suitable for passing as the buffer parameter to
	// DrawIndirect, DrawIndexedIndirect, DrawMeshTasksIndirectNV, DrawMeshTasksIndirectCount, or DispatchIndirect.
	// It is also suitable for passing as the buffer member of IndirectCommandsTokenNVX, or sequencesCountBuffer or
	// sequencesIndexBuffer member of CmdProcessCommandsInfoNVX
	BufferUsageIndirectBufferBit BufferUsageFlags = 0x00000100

	// BufferUsageTransformFeedbackBufferBitEXT specifies that the buffer is suitable for using for binding as a
	// transform feedback buffer with BindTransformFeedbackBuffersEXT.
	BufferUsageTransformFeedbackBufferBitEXT BufferUsageFlags = 0x00000800

	// BufferUsageTransformFeedbackCounterBufferBitEXT specifies that the buffer is suitable for using as a
	// counter buffer with BeginTransformFeedbackEXT and EndTransformFeedbackEXT.
	BufferUsageTransformFeedbackCounterBufferBitEXT BufferUsageFlags = 0x00001000

	// BufferUsageConditionalRenderingBitEXT specifies that the buffer is suitable for passing as
	// the buffer parameter to BeginConditionalRenderingEXT.
	BufferUsageConditionalRenderingBitEXT BufferUsageFlags = 0x00000200

	BufferUsageRaytracingBitNVX BufferUsageFlags = 0x00000400
)

const (
	MemoryPropertyDeviceLocalBit     MemoryPropertyFlags = 0x00000001
	MemoryPropertyHostVisibleBit     MemoryPropertyFlags = 0x00000002
	MemoryPropertyHostCoherentBit    MemoryPropertyFlags = 0x00000004
	MemoryPropertyHostCachedBit      MemoryPropertyFlags = 0x00000008
	MemoryPropertyLazilyAllocatedBit MemoryPropertyFlags = 0x00000010
	MemoryPropertyProtectedBit       MemoryPropertyFlags = 0x00000020
)

const (
	MemoryHeapDeviceLocalBit      MemoryHeapFlags = 0x00000001
	MemoryHeapMultiInstanceBit    MemoryHeapFlags = 0x00000002
	MemoryHeapMultiInstanceBitKHR MemoryHeapFlags = MemoryHeapMultiInstanceBit
)

const (
	// ImageTilingOptimal specifies optimal tiling (texels are laid out in an implementation-dependent arrangement, for more optimal memory access).
	ImageTilingOptimal ImageTiling = 0
	// ImageTilingLinear specifies linear tiling (texels are laid out in memory in row-major order, possibly with some padding on each row).
	ImageTilingLinear ImageTiling = 1
	// ImageTilingDrmFormatModifierEXT indicates that the image’s tiling is defined by a Linux DRM format modifier.
	ImageTilingDrmFormatModifierEXT ImageTiling = 1000158000
)

const (
	// ImageType1D specifies a one-dimensional image.
	ImageType1D ImageType = 0
	// ImageType2D specifies a two-dimensional image.
	ImageType2D ImageType = 1
	// ImageType3D specifies a three-dimensional image.
	ImageType3D ImageType = 2
)

const (
	ImageCreateSparseBindingBit                     ImageCreateFlags = 0x00000001
	ImageCreateSparseResidencyBit                   ImageCreateFlags = 0x00000002
	ImageCreateSparseAliasedBit                     ImageCreateFlags = 0x00000004
	ImageCreateMutableFormatBit                     ImageCreateFlags = 0x00000008
	ImageCreateCubeCompatibleBit                    ImageCreateFlags = 0x00000010
	ImageCreateAliasBit                             ImageCreateFlags = 0x00000400
	ImageCreateSplitInstanceBindRegionsBit          ImageCreateFlags = 0x00000040
	ImageCreate2DArrayCompatibleBit                 ImageCreateFlags = 0x00000020
	ImageCreateBlockTexelViewCompatibleBit          ImageCreateFlags = 0x00000080
	ImageCreateExtendedUsageBit                     ImageCreateFlags = 0x00000100
	ImageCreateProtectedBit                         ImageCreateFlags = 0x00000800
	ImageCreateDisjointBit                          ImageCreateFlags = 0x00000200
	ImageCreateCornerSampledBitNV                   ImageCreateFlags = 0x00002000
	ImageCreateSampleLocationsCompatibleDepthBitEXT ImageCreateFlags = 0x00001000
	ImageCreateSplitInstanceBindRegionsBitKHR       ImageCreateFlags = ImageCreateSplitInstanceBindRegionsBit
	ImageCreate2DArrayCompatibleBitKHR              ImageCreateFlags = ImageCreate2DArrayCompatibleBit
	ImageCreateBlockTexelViewCompatibleBitKHR       ImageCreateFlags = ImageCreateBlockTexelViewCompatibleBit
	ImageCreateExtendedUsageBitKHR                  ImageCreateFlags = ImageCreateExtendedUsageBit
	ImageCreateDisjointBitKHR                       ImageCreateFlags = ImageCreateDisjointBit
	ImageCreateAliasBitKHR                          ImageCreateFlags = ImageCreateAliasBit
)

const (
	IndexTypeUint16 IndexType = 0
	IndexTypeUint32 IndexType = 1
)

const (
	QueryTypeOcclusion                  QueryType = 0
	QueryTypePipelineStatistics         QueryType = 1
	QueryTypeTimestamp                  QueryType = 2
	QueryTypeTransformFeedbackStreamEXT QueryType = 1000028004
	QueryTypeCompactedSizeNVX           QueryType = 1000165000
)

const (
	FilterNearest  Filter = 0
	FilterLinear   Filter = 1
	FilterCubicIMG Filter = 1000015000
)

const (
	SamplerAddressModeRepeat            SamplerAddressMode = 0
	SamplerAddressModeMirroredRepeat    SamplerAddressMode = 1
	SamplerAddressModeClampToEdge       SamplerAddressMode = 2
	SamplerAddressModeClampToBorder     SamplerAddressMode = 3
	SamplerAddressModeMirrorClampToEdge SamplerAddressMode = 4
)

const (
	SamplerMipmapModeNearest SamplerMipmapMode = 0
	SamplerMipmapModeLinear  SamplerMipmapMode = 1
)

const (
	BorderColorFloatTransparentBlack BorderColor = 0
	BorderColorIntTransparentBlack   BorderColor = 1
	BorderColorFloatOpaqueBlack      BorderColor = 2
	BorderColorIntOpaqueBlack        BorderColor = 3
	BorderColorFloatOpaqueWhite      BorderColor = 4
	BorderColorIntOpaqueWhite        BorderColor = 5
)

const (
	QueryResult64Bit               QueryResultFlags = 0x00000001
	QueryResultWaitBit             QueryResultFlags = 0x00000002
	QueryResultWithAvailabilityBit QueryResultFlags = 0x00000004
	QueryResultPartialBit          QueryResultFlags = 0x00000008
)

const (
	StencilFaceFrontBit StencilFaceFlags = 0x00000001
	StencilFaceBackBit  StencilFaceFlags = 0x00000002
	StencilFrontAndBack StencilFaceFlags = 0x00000003
)

const (
	DescriptorTypeSampler                  DescriptorType = 0
	DescriptorTypeCombinedImageSampler     DescriptorType = 1
	DescriptorTypeSampledImage             DescriptorType = 2
	DescriptorTypeStorageImage             DescriptorType = 3
	DescriptorTypeUniformTexelBuffer       DescriptorType = 4
	DescriptorTypeStorageTexelBuffer       DescriptorType = 5
	DescriptorTypeUniformBuffer            DescriptorType = 6
	DescriptorTypeStorageBuffer            DescriptorType = 7
	DescriptorTypeUniformBufferDynamic     DescriptorType = 8
	DescriptorTypeStorageBufferDynamic     DescriptorType = 9
	DescriptorTypeInputAttachment          DescriptorType = 10
	DescriptorTypeInlineUniformBlockEXT    DescriptorType = 1000138000
	DescriptorTypeAccelerationStructureNVX DescriptorType = 1000165000
)

const (
	DescriptorPoolCreateFreeDescriptorSetBit  DescriptorPoolCreateFlags = 0x00000001
	DescriptorPoolCreateUpdateAfterBindBitEXT DescriptorPoolCreateFlags = 0x00000002
)

const (
	DescriptorSetLayoutCreatePushDescriptorBitKHR      DescriptorSetLayoutCreateFlags = 0x00000001
	DescriptorSetLayoutCreateUpdateAfterBindPoolBitEXT DescriptorSetLayoutCreateFlags = 0x00000002
)

const (
	FormatFeatureSampledImageBit                                                        = 0x00000001
	FormatFeatureStorageImageBit                                                        = 0x00000002
	FormatFeatureStorageImageAtomicBit                                                  = 0x00000004
	FormatFeatureUniformTexelBufferBit                                                  = 0x00000008
	FormatFeatureStorageTexelBufferBit                                                  = 0x00000010
	FormatFeatureStorageTexelBufferAtomicBit                                            = 0x00000020
	FormatFeatureVertexBufferBit                                                        = 0x00000040
	FormatFeatureColorAttachmentBit                                                     = 0x00000080
	FormatFeatureColorAttachmentBlendBit                                                = 0x00000100
	FormatFeatureDepthStencilAttachmentBit                                              = 0x00000200
	FormatFeatureBlitSrcBit                                                             = 0x00000400
	FormatFeatureBlitDstBit                                                             = 0x00000800
	FormatFeatureSampledImageFilterLinearBit                                            = 0x00001000
	FormatFeatureTransferSrcBit                                                         = 0x00004000
	FormatFeatureTransferDstBit                                                         = 0x00008000
	FormatFeatureMidpointChromaSamplesBit                                               = 0x00020000
	FormatFeatureSampledImageYcbcrConversionLinearFilterBit                             = 0x00040000
	FormatFeatureSampledImageYcbcrConversionSeparateReconstructionFilterBit             = 0x00080000
	FormatFeatureSampledImageYcbcrConversionChromaReconstructionExplicitBit             = 0x00100000
	FormatFeatureSampledImageYcbcrConversionChromaReconstructionExplicitForceableBit    = 0x00200000
	FormatFeatureDisjointBit                                                            = 0x00400000
	FormatFeatureCositedChromaSamplesBit                                                = 0x00800000
	FormatFeatureSampledImageFilterCubicBitIMG                                          = 0x00002000
	FormatFeatureSampledImageFilterMinmaxBitEXT                                         = 0x00010000
	FormatFeatureTransferSrcBitKHR                                                      = FormatFeatureTransferSrcBit
	FormatFeatureTransferDstBitKHR                                                      = FormatFeatureTransferDstBit
	FormatFeatureMidpointChromaSamplesBitKHR                                            = FormatFeatureMidpointChromaSamplesBit
	FormatFeatureSampledImageYcbcrConversionLinearFilterBitKHR                          = FormatFeatureSampledImageYcbcrConversionLinearFilterBit
	FormatFeatureSampledImageYcbcrConversionSeparateReconstructionFilterBitKHR          = FormatFeatureSampledImageYcbcrConversionSeparateReconstructionFilterBit
	FormatFeatureSampledImageYcbcrConversionChromaReconstructionExplicitBitKHR          = FormatFeatureSampledImageYcbcrConversionChromaReconstructionExplicitBit
	FormatFeatureSampledImageYcbcrConversionChromaReconstructionExplicitForceableBitKHR = FormatFeatureSampledImageYcbcrConversionChromaReconstructionExplicitForceableBit
	FormatFeatureDisjointBitKHR                                                         = FormatFeatureDisjointBit
	FormatFeatureCositedChromaSamplesBitKHR                                             = FormatFeatureCositedChromaSamplesBit
)

func (res Result) Error() string { return res.String() }

func (flags DeviceQueueCreateFlags) String() string {
	var props []string
	if (flags & DeviceQueueCreateProtectedBit) != 0 {
		props = append(props, "DeviceQueueCreateProtectedBit")
	}
	return strings.Join(props, " | ")
}

func (flags QueueFlags) String() string {
	var props []string
	if (flags & QueueGraphicsBit) != 0 {
		props = append(props, "QueueGraphicsBit")
	}
	if (flags & QueueComputeBit) != 0 {
		props = append(props, "QueueComputeBit")
	}
	if (flags & QueueTransferBit) != 0 {
		props = append(props, "QueueTransferBit")
	}
	if (flags & QueueSparseBindingBit) != 0 {
		props = append(props, "QueueSparseBindingBit")
	}
	if (flags & QueueProtectedBit) != 0 {
		props = append(props, "QueueProtectedBit")
	}
	return strings.Join(props, " | ")
}

func (flags SurfaceTransformFlagsKHR) String() string {
	var out []string
	if (flags & SurfaceTransformIdentityBitKHR) != 0 {
		out = append(out, "SurfaceTransformIdentityBitKHR")
	}
	if (flags & SurfaceTransformRotate90BitKHR) != 0 {
		out = append(out, "SurfaceTransformRotate90BitKHR")
	}
	if (flags & SurfaceTransformRotate180BitKHR) != 0 {
		out = append(out, "SurfaceTransformRotate180BitKHR")
	}
	if (flags & SurfaceTransformRotate270BitKHR) != 0 {
		out = append(out, "SurfaceTransformRotate270BitKHR")
	}
	if (flags & SurfaceTransformHorizontalMirrorBitKHR) != 0 {
		out = append(out, "SurfaceTransformHorizontalMirrorBitKHR")
	}
	if (flags & SurfaceTransformHorizontalMirrorRotate90BitKHR) != 0 {
		out = append(out, "SurfaceTransformHorizontalMirrorRotate90BitKHR")
	}
	if (flags & SurfaceTransformHorizontalMirrorRotate180BitKHR) != 0 {
		out = append(out, "SurfaceTransformHorizontalMirrorRotate180BitKHR")
	}
	if (flags & SurfaceTransformHorizontalMirrorRotate270BitKHR) != 0 {
		out = append(out, "SurfaceTransformHorizontalMirrorRotate270BitKHR")
	}
	if (flags & SurfaceTransformInheritBitKHR) != 0 {
		out = append(out, "SurfaceTransformInheritBitKHR")
	}
	return strings.Join(out, " | ")
}

func (flags CompositeAlphaFlagsKHR) String() string {
	var out []string
	if (flags & CompositeAlphaOpaqueBitKHR) != 0 {
		out = append(out, "CompositeAlphaOpaqueBitKHR")
	}
	if (flags & CompositeAlphaPreMultipliedBitKHR) != 0 {
		out = append(out, "CompositeAlphaPreMultipliedBitKHR")
	}
	if (flags & CompositeAlphaPostMultipliedBitKHR) != 0 {
		out = append(out, "CompositeAlphaPostMultipliedBitKHR")
	}
	if (flags & CompositeAlphaInheritBitKHR) != 0 {
		out = append(out, "CompositeAlphaInheritBitKHR")
	}
	return strings.Join(out, " | ")
}

func (flags ImageUsageFlags) String() string {
	var out []string
	if (flags & ImageUsageTransferSrcBit) != 0 {
		out = append(out, "ImageUsageTransferSrcBit")
	}
	if (flags & ImageUsageTransferDstBit) != 0 {
		out = append(out, "ImageUsageTransferDstBit")
	}
	if (flags & ImageUsageSampledBit) != 0 {
		out = append(out, "ImageUsageSampledBit")
	}
	if (flags & ImageUsageStorageBit) != 0 {
		out = append(out, "ImageUsageStorageBit")
	}
	if (flags & ImageUsageColorAttachmentBit) != 0 {
		out = append(out, "ImageUsageColorAttachmentBit")
	}
	if (flags & ImageUsageDepthStencilAttachmentBit) != 0 {
		out = append(out, "ImageUsageDepthStencilAttachmentBit")
	}
	if (flags & ImageUsageTransientAttachmentBit) != 0 {
		out = append(out, "ImageUsageTransientAttachmentBit")
	}
	return strings.Join(out, " | ")
}

func (flags CommandPoolCreateFlags) String() string {
	var out []string
	if (flags & CommandPoolCreateTransientBit) != 0 {
		out = append(out, "CommandPoolCreateTransientBit")
	}
	if (flags & CommandPoolCreateResetCommandBufferBit) != 0 {
		out = append(out, "CommandPoolCreateResetCommandBufferBit")
	}
	if (flags & CommandPoolCreateProtectedBit) != 0 {
		out = append(out, "CommandPoolCreateProtectedBit")
	}
	return strings.Join(out, " | ")
}

func (flags CommandPoolTrimFlags) String() string { return "" }

func (flags CommandPoolResetFlags) String() string {
	var out []string
	if (flags & CommandlPoolResetReleaseResourcesBit) != 0 {
		out = append(out, "CommandlPoolResetReleaseResourcesBit")
	}
	return strings.Join(out, " | ")
}

func (flags CommandBufferResetFlags) String() string {
	var out []string
	if (flags & CommandBufferResetReleaseResourcesBit) != 0 {
		out = append(out, "CommandBufferResetReleaseResourcesBit")
	}
	return strings.Join(out, " | ")
}

func (flags CommandBufferUsageFlags) String() string {
	var out []string
	if (flags & CommandBufferUsageOneTimeSubmitBit) != 0 {
		out = append(out, "CommandBufferUsageOneTimeSubmitBit")
	}
	if (flags & CommandBufferUsageRenderPassContinueBit) != 0 {
		out = append(out, "CommandBufferUsageRenderPassContinueBit")
	}
	if (flags & CommandBufferUsageSimultaneousUseBit) != 0 {
		out = append(out, "CommandBufferUsageSimultaneousUseBit")
	}
	return strings.Join(out, " | ")
}

func (flags QueryPipelineStatisticFlags) String() string {
	var out []string
	if (flags & QueryPipelineStatisticInputAssemblyVerticesBit) != 0 {
		out = append(out, "QueryPipelineStatisticInputAssemblyVerticesBit")
	}
	if (flags & QueryPipelineStatisticInputAssemblyPrimitivesBit) != 0 {
		out = append(out, "QueryPipelineStatisticInputAssemblyPrimitivesBit")
	}
	if (flags & QueryPipelineStatisticVertexShaderInvocationsBit) != 0 {
		out = append(out, "QueryPipelineStatisticVertexShaderInvocationsBit")
	}
	if (flags & QueryPipelineStatisticGeometryShaderInvocationsBit) != 0 {
		out = append(out, "QueryPipelineStatisticGeometryShaderInvocationsBit")
	}
	if (flags & QueryPipelineStatisticGeometryShaderPrimitivesBit) != 0 {
		out = append(out, "QueryPipelineStatisticGeometryShaderPrimitivesBit")
	}
	if (flags & QueryPipelineStatisticClippingInvocationsBit) != 0 {
		out = append(out, "QueryPipelineStatisticClippingInvocationsBit")
	}
	if (flags & QueryPipelineStatisticClippingPrimitivesBit) != 0 {
		out = append(out, "QueryPipelineStatisticClippingPrimitivesBit")
	}
	if (flags & QueryPipelineStatisticFragmentShaderInvocationsBit) != 0 {
		out = append(out, "QueryPipelineStatisticFragmentShaderInvocationsBit")
	}
	if (flags & QueryPipelineStatisticTessellationControlShaderPatchesBit) != 0 {
		out = append(out, "QueryPipelineStatisticTessellationControlShaderPatchesBit")
	}
	if (flags & QueryPipelineStatisticTessellationEvaluationShaderInvocationsBit) != 0 {
		out = append(out, "QueryPipelineStatisticTessellationEvaluationShaderInvocationsBit")
	}
	if (flags & QueryPipelineStatisticComputeShaderInvocationsBit) != 0 {
		out = append(out, "QueryPipelineStatisticComputeShaderInvocationsBit")
	}
	return strings.Join(out, " | ")
}

func (flags QueryControlFlags) String() string {
	var out []string
	if (flags & QueryControlPreciseBit) != 0 {
		out = append(out, "QueryControlPreciseBit")
	}
	return strings.Join(out, " | ")
}

func (flags ImageAspectFlags) String() string {
	var out []string
	if (flags & ImageAspectColorBit) != 0 {
		out = append(out, "ImageAspectColorBit")
	}
	if (flags & ImageAspectDepthBit) != 0 {
		out = append(out, "ImageAspectDepthBit")
	}
	if (flags & ImageAspectStencilBit) != 0 {
		out = append(out, "ImageAspectStencilBit")
	}
	if (flags & ImageAspectMetadataBit) != 0 {
		out = append(out, "ImageAspectMetadataBit")
	}
	if (flags & ImageAspectPlane0Bit) != 0 {
		out = append(out, "ImageAspectPlane0Bit")
	}
	if (flags & ImageAspectPlane1Bit) != 0 {
		out = append(out, "ImageAspectPlane1Bit")
	}
	if (flags & ImageAspectPlane2Bit) != 0 {
		out = append(out, "ImageAspectPlane2Bit")
	}
	if (flags & ImageAspectMemoryPlane0BitEXT) != 0 {
		out = append(out, "ImageAspectMemoryPlane0BitEXT")
	}
	if (flags & ImageAspectMemoryPlane1BitEXT) != 0 {
		out = append(out, "ImageAspectMemoryPlane1BitEXT")
	}
	if (flags & ImageAspectMemoryPlane2BitEXT) != 0 {
		out = append(out, "ImageAspectMemoryPlane2BitEXT")
	}
	if (flags & ImageAspectMemoryPlane3BitEXT) != 0 {
		out = append(out, "ImageAspectMemoryPlane3BitEXT")
	}
	if (flags & ImageAspectPlane0BitKHR) != 0 {
		out = append(out, "ImageAspectPlane0BitKHR")
	}
	if (flags & ImageAspectPlane1BitKHR) != 0 {
		out = append(out, "ImageAspectPlane1BitKHR")
	}
	if (flags & ImageAspectPlane2BitKHR) != 0 {
		out = append(out, "ImageAspectPlane2BitKHR")
	}
	return strings.Join(out, " | ")
}

func (flags ShaderStageFlags) String() string {
	var out []string
	if (flags & ShaderStageVertexBit) != 0 {
		out = append(out, "ShaderStageVertexBit")
	}
	if (flags & ShaderStageTessellationControlBit) != 0 {
		out = append(out, "ShaderStageTessellationControlBit")
	}
	if (flags & ShaderStageTessellationEvaluationBit) != 0 {
		out = append(out, "ShaderStageTessellationEvaluationBit")
	}
	if (flags & ShaderStageGeometryBit) != 0 {
		out = append(out, "ShaderStageGeometryBit")
	}
	if (flags & ShaderStageFragmentBit) != 0 {
		out = append(out, "ShaderStageFragmentBit")
	}
	if (flags & ShaderStageComputeBit) != 0 {
		out = append(out, "ShaderStageComputeBit")
	}
	if (flags & ShaderStageAllGraphics) != 0 {
		out = append(out, "ShaderStageAllGraphics")
	}
	if (flags & ShaderStageAll) != 0 {
		out = append(out, "ShaderStageAll")
	}
	if (flags & ShaderStageRaygenBitNVX) != 0 {
		out = append(out, "ShaderStageRaygenBitNVX")
	}
	if (flags & ShaderStageAnyHitBitNVX) != 0 {
		out = append(out, "ShaderStageAnyHitBitNVX")
	}
	if (flags & ShaderStageClosestHitBitNVX) != 0 {
		out = append(out, "ShaderStageClosestHitBitNVX")
	}
	if (flags & ShaderStageMissBitNVX) != 0 {
		out = append(out, "ShaderStageMissBitNVX")
	}
	if (flags & ShaderStageIntersectionBitNVX) != 0 {
		out = append(out, "ShaderStageIntersectionBitNVX")
	}
	if (flags & ShaderStageCallableBitNVX) != 0 {
		out = append(out, "ShaderStageCallableBitNVX")
	}
	if (flags & ShaderStageTaskBitNV) != 0 {
		out = append(out, "ShaderStageTaskBitNV")
	}
	if (flags & ShaderStageMeshBitNV) != 0 {
		out = append(out, "ShaderStageMeshBitNV")
	}
	return strings.Join(out, " | ")
}

func (flags CullModeFlags) String() string {
	var out []string
	if (flags & CullModeNone) != 0 {
		out = append(out, "CullModeNone")
	}
	if (flags & CullModeFrontBit) != 0 {
		out = append(out, "CullModeFrontBit")
	}
	if (flags & CullModeBackBit) != 0 {
		out = append(out, "CullModeBackBit")
	}
	if (flags & CullModeFrontAndBack) != 0 {
		out = append(out, "CullModeFrontAndBack")
	}
	return strings.Join(out, " | ")
}

func (flags SampleCountFlags) String() string {
	var out []string
	if (flags & SampleCount1Bit) != 0 {
		out = append(out, "SampleCount1Bit")
	}
	if (flags & SampleCount2Bit) != 0 {
		out = append(out, "SampleCount2Bit")
	}
	if (flags & SampleCount4Bit) != 0 {
		out = append(out, "SampleCount4Bit")
	}
	if (flags & SampleCount8Bit) != 0 {
		out = append(out, "SampleCount8Bit")
	}
	if (flags & SampleCount16Bit) != 0 {
		out = append(out, "SampleCount16Bit")
	}
	if (flags & SampleCount32Bit) != 0 {
		out = append(out, "SampleCount32Bit")
	}
	if (flags & SampleCount64Bit) != 0 {
		out = append(out, "SampleCount64Bit")
	}
	return strings.Join(out, " | ")
}

func (flags ColorComponentFlags) String() string {
	var out []string
	if (flags & ColorComponentR) != 0 {
		out = append(out, "ColorComponentR")
	}
	if (flags & ColorComponentG) != 0 {
		out = append(out, "ColorComponentG")
	}
	if (flags & ColorComponentB) != 0 {
		out = append(out, "ColorComponentB")
	}
	if (flags & ColorComponentA) != 0 {
		out = append(out, "ColorComponentA")
	}
	return strings.Join(out, " | ")
}

func (flags PipelineCreateFlags) String() string {
	var out []string
	if (flags & PipelineCreateDisableOptimizationBit) != 0 {
		out = append(out, "PipelineCreateDisableOptimizationBit")
	}
	if (flags & PipelineCreateAllowDerivativesBit) != 0 {
		out = append(out, "PipelineCreateAllowDerivativesBit")
	}
	if (flags & PipelineCreateDerivativeBit) != 0 {
		out = append(out, "PipelineCreateDerivativeBit")
	}
	if (flags & PipelineCreateViewIndexFromDeviceIndexBit) != 0 {
		out = append(out, "PipelineCreateViewIndexFromDeviceIndexBit")
	}
	if (flags & PipelineCreateDispatchBase) != 0 {
		out = append(out, "PipelineCreateDispatchBase")
	}
	if (flags & PipelineCreateViewIndexFromDeviceIndexBitKHR) != 0 {
		out = append(out, "PipelineCreateViewIndexFromDeviceIndexBitKHR")
	}
	if (flags & PipelineCreateDispatchBaseKHR) != 0 {
		out = append(out, "PipelineCreateDispatchBaseKHR")
	}
	return strings.Join(out, " | ")
}

func (flags AttachmentDescriptionFlags) String() string {
	var out []string
	if (flags & AttachmentDescriptionMayAliasBit) != 0 {
		out = append(out, "AttachmentDescriptionMayAliasBit")
	}
	return strings.Join(out, " | ")
}

func (flags SubpassDescriptionFlags) String() string {
	var out []string
	if (flags & SubpassDescriptionPerViewAttributesBitNVX) != 0 {
		out = append(out, "SubpassDescriptionPerViewAttributesBitNVX")
	}
	if (flags & SubpassDescriptionPerViewPositionXOnlyBitNVX) != 0 {
		out = append(out, "SubpassDescriptionPerViewPositionXOnlyBitNVX")
	}
	return strings.Join(out, " | ")
}

func (flags PipelineStageFlags) String() string {
	var out []string
	if (flags & PipelineStageTopOfPipeBit) != 0 {
		out = append(out, "PipelineStageTopOfPipeBit")
	}
	if (flags & PipelineStageDrawIndirectBit) != 0 {
		out = append(out, "PipelineStageDrawIndirectBit")
	}
	if (flags & PipelineStageVertexInputBit) != 0 {
		out = append(out, "PipelineStageVertexInputBit")
	}
	if (flags & PipelineStageVertexShaderBit) != 0 {
		out = append(out, "PipelineStageVertexShaderBit")
	}
	if (flags & PipelineStageTessellationControlShaderBit) != 0 {
		out = append(out, "PipelineStageTessellationControlShaderBit")
	}
	if (flags & PipelineStageTessellationEvaluationShaderBit) != 0 {
		out = append(out, "PipelineStageTessellationEvaluationShaderBit")
	}
	if (flags & PipelineStageGeometryShaderBit) != 0 {
		out = append(out, "PipelineStageGeometryShaderBit")
	}
	if (flags & PipelineStageFragmentShaderBit) != 0 {
		out = append(out, "PipelineStageFragmentShaderBit")
	}
	if (flags & PipelineStageEarlyFragmentTestsBit) != 0 {
		out = append(out, "PipelineStageEarlyFragmentTestsBit")
	}
	if (flags & PipelineStageLateFragmentTestsBit) != 0 {
		out = append(out, "PipelineStageLateFragmentTestsBit")
	}
	if (flags & PipelineStageColorAttachmentOutputBit) != 0 {
		out = append(out, "PipelineStageColorAttachmentOutputBit")
	}
	if (flags & PipelineStageComputeShaderBit) != 0 {
		out = append(out, "PipelineStageComputeShaderBit")
	}
	if (flags & PipelineStageTransferBit) != 0 {
		out = append(out, "PipelineStageTransferBit")
	}
	if (flags & PipelineStageBottomOfPipeBit) != 0 {
		out = append(out, "PipelineStageBottomOfPipeBit")
	}
	if (flags & PipelineStageHostBit) != 0 {
		out = append(out, "PipelineStageHostBit")
	}
	if (flags & PipelineStageAllGraphicsBit) != 0 {
		out = append(out, "PipelineStageAllGraphicsBit")
	}
	if (flags & PipelineStageAllCommandsBit) != 0 {
		out = append(out, "PipelineStageAllCommandsBit")
	}
	if (flags & PipelineStageTransformFeedbackBitEXT) != 0 {
		out = append(out, "PipelineStageTransformFeedbackBitEXT")
	}
	if (flags & PipelineStageConditionalRenderingBitEXT) != 0 {
		out = append(out, "PipelineStageConditionalRenderingBitEXT")
	}
	if (flags & PipelineStageCommandProcessBitNVX) != 0 {
		out = append(out, "PipelineStageCommandProcessBitNVX")
	}
	if (flags & PipelineStageShadingRateImageBitNV) != 0 {
		out = append(out, "PipelineStageShadingRateImageBitNV")
	}
	if (flags & PipelineStageRaytracingBitNVX) != 0 {
		out = append(out, "PipelineStageRaytracingBitNVX")
	}
	if (flags & PipelineStageTaskShaderBitNV) != 0 {
		out = append(out, "PipelineStageTaskShaderBitNV")
	}
	if (flags & PipelineStageMeshShaderBitNV) != 0 {
		out = append(out, "PipelineStageMeshShaderBitNV")
	}
	return strings.Join(out, " | ")
}

func (flags AccessFlags) String() string {
	var out []string
	if (flags & AccessIndirectCommandReadBit) != 0 {
		out = append(out, "AccessIndirectCommandReadBit")
	}
	if (flags & AccessIndexReadBit) != 0 {
		out = append(out, "AccessIndexReadBit")
	}
	if (flags & AccessVertexAttributeReadBit) != 0 {
		out = append(out, "AccessVertexAttributeReadBit")
	}
	if (flags & AccessUniformReadBit) != 0 {
		out = append(out, "AccessUniformReadBit")
	}
	if (flags & AccessInputAttachmentReadBit) != 0 {
		out = append(out, "AccessInputAttachmentReadBit")
	}
	if (flags & AccessShaderReadBit) != 0 {
		out = append(out, "AccessShaderReadBit")
	}
	if (flags & AccessShaderWriteBit) != 0 {
		out = append(out, "AccessShaderWriteBit")
	}
	if (flags & AccessColorAttachmentReadBit) != 0 {
		out = append(out, "AccessColorAttachmentReadBit")
	}
	if (flags & AccessColorAttachmentWriteBit) != 0 {
		out = append(out, "AccessColorAttachmentWriteBit")
	}
	if (flags & AccessDepthStencilAttachmentReadBit) != 0 {
		out = append(out, "AccessDepthStencilAttachmentReadBit")
	}
	if (flags & AccessDepthStencilAttachmentWriteBit) != 0 {
		out = append(out, "AccessDepthStencilAttachmentWriteBit")
	}
	if (flags & AccessTransferReadBit) != 0 {
		out = append(out, "AccessTransferReadBit")
	}
	if (flags & AccessTransferWriteBit) != 0 {
		out = append(out, "AccessTransferWriteBit")
	}
	if (flags & AccessHostReadBit) != 0 {
		out = append(out, "AccessHostReadBit")
	}
	if (flags & AccessHostWriteBit) != 0 {
		out = append(out, "AccessHostWriteBit")
	}
	if (flags & AccessMemoryReadBit) != 0 {
		out = append(out, "AccessMemoryReadBit")
	}
	if (flags & AccessMemoryWriteBit) != 0 {
		out = append(out, "AccessMemoryWriteBit")
	}
	if (flags & AccessTransformFeedbackWriteBitEXT) != 0 {
		out = append(out, "AccessTransformFeedbackWriteBitEXT")
	}
	if (flags & AccessTransformFeedbackCounterReadBitEXT) != 0 {
		out = append(out, "AccessTransformFeedbackCounterReadBitEXT")
	}
	if (flags & AccessTransformFeedbackCounterWriteBitEXT) != 0 {
		out = append(out, "AccessTransformFeedbackCounterWriteBitEXT")
	}
	if (flags & AccessConditionalRenderingReadBitEXT) != 0 {
		out = append(out, "AccessConditionalRenderingReadBitEXT")
	}
	if (flags & AccessCommandProcessReadBitNVX) != 0 {
		out = append(out, "AccessCommandProcessReadBitNVX")
	}
	if (flags & AccessCommandProcessWriteBitNVX) != 0 {
		out = append(out, "AccessCommandProcessWriteBitNVX")
	}
	if (flags & AccessColorAttachmentReadNoncoherentBitEXT) != 0 {
		out = append(out, "AccessColorAttachmentReadNoncoherentBitEXT")
	}
	if (flags & AccessShadingRateImageReadBitNV) != 0 {
		out = append(out, "AccessShadingRateImageReadBitNV")
	}
	if (flags & AccessAccelerationStructureReadBitNVX) != 0 {
		out = append(out, "AccessAccelerationStructureReadBitNVX")
	}
	if (flags & AccessAccelerationStructureWriteBitNVX) != 0 {
		out = append(out, "AccessAccelerationStructureWriteBitNVX")
	}
	return strings.Join(out, " | ")
}

func (flags DependencyFlags) String() string {
	var out []string
	if (flags & DependencyByRegionBit) != 0 {
		out = append(out, "DependencyByRegionBit")
	}
	if (flags & DependencyDeviceGroupBit) != 0 {
		out = append(out, "DependencyDeviceGroupBit")
	}
	if (flags & DependencyViewLocalBit) != 0 {
		out = append(out, "DependencyViewLocalBit")
	}
	if (flags & DependencyViewLocalBitKHR) != 0 {
		out = append(out, "DependencyViewLocalBitKHR")
	}
	if (flags & DependencyDeviceGroupBitKHR) != 0 {
		out = append(out, "DependencyDeviceGroupBitKHR")
	}
	return strings.Join(out, " | ")
}

func (flags FenceCreateFlags) String() string {
	var out []string
	if (flags & FenceCreateSignaledBit) != 0 {
		out = append(out, "FenceCreateSignaledBit")
	}
	return strings.Join(out, " | ")
}

func (flags BufferCreateFlags) String() string {
	var out []string
	if (flags & BufferCreateSparseBindingBit) != 0 {
		out = append(out, "BufferCreateSparseBindingBit")
	}
	if (flags & BufferCreateSparseResidencyBit) != 0 {
		out = append(out, "BufferCreateSparseResidencyBit")
	}
	if (flags & BufferCreateSparseAliasedBit) != 0 {
		out = append(out, "BufferCreateSparseAliasedBit")
	}
	if (flags & BufferCreateProtectedBit) != 0 {
		out = append(out, "BufferCreateProtectedBit")
	}
	return strings.Join(out, " | ")
}

func (flags BufferUsageFlags) String() string {
	var out []string
	if (flags & BufferUsageTransferSrcBit) != 0 {
		out = append(out, "BufferUsageTransferSrcBit")
	}
	if (flags & BufferUsageTransferDstBit) != 0 {
		out = append(out, "BufferUsageTransferDstBit")
	}
	if (flags & BufferUsageUniformTexelBufferBit) != 0 {
		out = append(out, "BufferUsageUniformTexelBufferBit")
	}
	if (flags & BufferUsageStorageTexelBufferBit) != 0 {
		out = append(out, "BufferUsageStorageTexelBufferBit")
	}
	if (flags & BufferUsageUniformBufferBit) != 0 {
		out = append(out, "BufferUsageUniformBufferBit")
	}
	if (flags & BufferUsageStorageBufferBit) != 0 {
		out = append(out, "BufferUsageStorageBufferBit")
	}
	if (flags & BufferUsageIndexBufferBit) != 0 {
		out = append(out, "BufferUsageIndexBufferBit")
	}
	if (flags & BufferUsageVertexBufferBit) != 0 {
		out = append(out, "BufferUsageVertexBufferBit")
	}
	if (flags & BufferUsageIndirectBufferBit) != 0 {
		out = append(out, "BufferUsageIndirectBufferBit")
	}
	if (flags & BufferUsageTransformFeedbackBufferBitEXT) != 0 {
		out = append(out, "BufferUsageTransformFeedbackBufferBitEXT")
	}
	if (flags & BufferUsageTransformFeedbackCounterBufferBitEXT) != 0 {
		out = append(out, "BufferUsageTransformFeedbackCounterBufferBitEXT")
	}
	if (flags & BufferUsageConditionalRenderingBitEXT) != 0 {
		out = append(out, "BufferUsageConditionalRenderingBitEXT")
	}
	if (flags & BufferUsageRaytracingBitNVX) != 0 {
		out = append(out, "BufferUsageRaytracingBitNVX")
	}
	return strings.Join(out, " | ")
}

func (flags MemoryPropertyFlags) String() string {
	var out []string
	if (flags & MemoryPropertyDeviceLocalBit) != 0 {
		out = append(out, "MemoryPropertyDeviceLocalBit")
	}
	if (flags & MemoryPropertyHostVisibleBit) != 0 {
		out = append(out, "MemoryPropertyHostVisibleBit")
	}
	if (flags & MemoryPropertyHostCoherentBit) != 0 {
		out = append(out, "MemoryPropertyHostCoherentBit")
	}
	if (flags & MemoryPropertyHostCachedBit) != 0 {
		out = append(out, "MemoryPropertyHostCachedBit")
	}
	if (flags & MemoryPropertyLazilyAllocatedBit) != 0 {
		out = append(out, "MemoryPropertyLazilyAllocatedBit")
	}
	if (flags & MemoryPropertyProtectedBit) != 0 {
		out = append(out, "MemoryPropertyProtectedBit")
	}
	return strings.Join(out, " | ")
}

func (flags MemoryHeapFlags) String() string {
	var out []string
	if (flags & MemoryHeapDeviceLocalBit) != 0 {
		out = append(out, "MemoryHeapDeviceLocalBit")
	}
	if (flags & MemoryHeapMultiInstanceBit) != 0 {
		out = append(out, "MemoryHeapMultiInstanceBit")
	}
	if (flags & MemoryHeapMultiInstanceBitKHR) != 0 {
		out = append(out, "MemoryHeapMultiInstanceBitKHR")
	}
	return strings.Join(out, " | ")
}

func (flags QueryResultFlags) String() string {
	var out []string
	if (flags & QueryResult64Bit) != 0 {
		out = append(out, "QueryResult64Bit")
	}
	if (flags & QueryResultWaitBit) != 0 {
		out = append(out, "QueryResultWaitBit")
	}
	if (flags & QueryResultWithAvailabilityBit) != 0 {
		out = append(out, "QueryResultWithAvailabilityBit")
	}
	if (flags & QueryResultPartialBit) != 0 {
		out = append(out, "QueryResultPartialBit")
	}
	return strings.Join(out, " | ")
}

func (flags StencilFaceFlags) String() string {
	var out []string
	if (flags & StencilFaceFrontBit) != 0 {
		out = append(out, "StencilFaceFrontBit")
	}
	if (flags & StencilFaceBackBit) != 0 {
		out = append(out, "StencilFaceBackBit")
	}
	return strings.Join(out, " | ")
}

func (flags DescriptorPoolCreateFlags) String() string {
	var out []string
	if (flags & DescriptorPoolCreateFreeDescriptorSetBit) != 0 {
		out = append(out, "DescriptorPoolCreateFreeDescriptorSetBit")
	}
	if (flags & DescriptorPoolCreateUpdateAfterBindBitEXT) != 0 {
		out = append(out, "DescriptorPoolCreateUpdateAfterBindBitEXT")
	}
	return strings.Join(out, " | ")
}

func (flags DescriptorSetLayoutCreateFlags) String() string {
	var out []string
	if (flags & DescriptorSetLayoutCreatePushDescriptorBitKHR) != 0 {
		out = append(out, "DescriptorSetLayoutCreatePushDescriptorBitKHR")
	}
	if (flags & DescriptorSetLayoutCreateUpdateAfterBindPoolBitEXT) != 0 {
		out = append(out, "DescriptorSetLayoutCreateUpdateAfterBindPoolBitEXT")
	}
	return strings.Join(out, " | ")
}

func (flags FormatFeatureFlags) String() string {
	var out []string
	if (flags & FormatFeatureSampledImageBit) != 0 {
		out = append(out, "FormatFeatureSampledImageBit")
	}
	if (flags & FormatFeatureStorageImageBit) != 0 {
		out = append(out, "FormatFeatureStorageImageBit")
	}
	if (flags & FormatFeatureStorageImageAtomicBit) != 0 {
		out = append(out, "FormatFeatureStorageImageAtomicBit")
	}
	if (flags & FormatFeatureUniformTexelBufferBit) != 0 {
		out = append(out, "FormatFeatureUniformTexelBufferBit")
	}
	if (flags & FormatFeatureStorageTexelBufferBit) != 0 {
		out = append(out, "FormatFeatureStorageTexelBufferBit")
	}
	if (flags & FormatFeatureStorageTexelBufferAtomicBit) != 0 {
		out = append(out, "FormatFeatureStorageTexelBufferAtomicBit")
	}
	if (flags & FormatFeatureVertexBufferBit) != 0 {
		out = append(out, "FormatFeatureVertexBufferBit")
	}
	if (flags & FormatFeatureColorAttachmentBit) != 0 {
		out = append(out, "FormatFeatureColorAttachmentBit")
	}
	if (flags & FormatFeatureColorAttachmentBlendBit) != 0 {
		out = append(out, "FormatFeatureColorAttachmentBlendBit")
	}
	if (flags & FormatFeatureDepthStencilAttachmentBit) != 0 {
		out = append(out, "FormatFeatureDepthStencilAttachmentBit")
	}
	if (flags & FormatFeatureBlitSrcBit) != 0 {
		out = append(out, "FormatFeatureBlitSrcBit")
	}
	if (flags & FormatFeatureBlitDstBit) != 0 {
		out = append(out, "FormatFeatureBlitDstBit")
	}
	if (flags & FormatFeatureSampledImageFilterLinearBit) != 0 {
		out = append(out, "FormatFeatureSampledImageFilterLinearBit")
	}
	if (flags & FormatFeatureTransferSrcBit) != 0 {
		out = append(out, "FormatFeatureTransferSrcBit")
	}
	if (flags & FormatFeatureTransferDstBit) != 0 {
		out = append(out, "FormatFeatureTransferDstBit")
	}
	if (flags & FormatFeatureMidpointChromaSamplesBit) != 0 {
		out = append(out, "FormatFeatureMidpointChromaSamplesBit")
	}
	if (flags & FormatFeatureSampledImageYcbcrConversionLinearFilterBit) != 0 {
		out = append(out, "FormatFeatureSampledImageYcbcrConversionLinearFilterBit")
	}
	if (flags & FormatFeatureSampledImageYcbcrConversionSeparateReconstructionFilterBit) != 0 {
		out = append(out, "FormatFeatureSampledImageYcbcrConversionSeparateReconstructionFilterBit")
	}
	if (flags & FormatFeatureSampledImageYcbcrConversionChromaReconstructionExplicitBit) != 0 {
		out = append(out, "FormatFeatureSampledImageYcbcrConversionChromaReconstructionExplicitBit")
	}
	if (flags & FormatFeatureSampledImageYcbcrConversionChromaReconstructionExplicitForceableBit) != 0 {
		out = append(out, "FormatFeatureSampledImageYcbcrConversionChromaReconstructionExplicitForceableBit")
	}
	if (flags & FormatFeatureDisjointBit) != 0 {
		out = append(out, "FormatFeatureDisjointBit")
	}
	if (flags & FormatFeatureCositedChromaSamplesBit) != 0 {
		out = append(out, "FormatFeatureCositedChromaSamplesBit")
	}
	if (flags & FormatFeatureSampledImageFilterCubicBitIMG) != 0 {
		out = append(out, "FormatFeatureSampledImageFilterCubicBitIMG")
	}
	if (flags & FormatFeatureSampledImageFilterMinmaxBitEXT) != 0 {
		out = append(out, "FormatFeatureSampledImageFilterMinmaxBitEXT")
	}
	return strings.Join(out, " | ")
}

// Copyright (c) 2015-2018 Khronos Group
// Copyright (c) 2018 Dominik Honnef

package vk

// #cgo pkg-config: vulkan
// #cgo CFLAGS: -DVK_NO_PROTOTYPES
// #include <vulkan/vulkan_core.h>
import "C"
import (
	"strings"
	"unsafe"
)

func init() {
	assertSameSize(unsafe.Sizeof(DeviceQueueCreateFlags(0)), unsafe.Sizeof(C.VkDeviceQueueCreateFlags(0)))
	assertSameSize(unsafe.Sizeof(QueueFlags(0)), unsafe.Sizeof(C.VkQueueFlags(0)))
	assertSameSize(unsafe.Sizeof(PhysicalDeviceType(0)), unsafe.Sizeof(C.VkPhysicalDeviceType(0)))
	assertSameSize(unsafe.Sizeof(Result(0)), unsafe.Sizeof(C.VkResult(0)))
	assertSameSize(unsafe.Sizeof(StructureType(0)), unsafe.Sizeof(C.VkStructureType(0)))
	assertSameSize(unsafe.Sizeof(SurfaceTransformFlagsKHR(0)), unsafe.Sizeof(C.VkSurfaceTransformFlagsKHR(0)))
	assertSameSize(unsafe.Sizeof(CompositeAlphaFlagsKHR(0)), unsafe.Sizeof(C.VkCompositeAlphaFlagsKHR(0)))
	assertSameSize(unsafe.Sizeof(ImageUsageFlags(0)), unsafe.Sizeof(C.VkImageUsageFlags(0)))
	assertSameSize(unsafe.Sizeof(Format(0)), unsafe.Sizeof(C.VkFormat(0)))
	assertSameSize(unsafe.Sizeof(ColorSpaceKHR(0)), unsafe.Sizeof(C.VkColorSpaceKHR(0)))
	assertSameSize(unsafe.Sizeof(PresentModeKHR(0)), unsafe.Sizeof(C.VkPresentModeKHR(0)))
	assertSameSize(unsafe.Sizeof(CommandPoolCreateFlags(0)), unsafe.Sizeof(C.VkCommandPoolCreateFlags(0)))
	assertSameSize(unsafe.Sizeof(CommandPoolTrimFlags(0)), unsafe.Sizeof(C.VkCommandPoolTrimFlags(0)))
	assertSameSize(unsafe.Sizeof(CommandPoolResetFlags(0)), unsafe.Sizeof(C.VkCommandPoolResetFlags(0)))
	assertSameSize(unsafe.Sizeof(CommandBufferLevel(0)), unsafe.Sizeof(C.VkCommandBufferLevel(0)))
	assertSameSize(unsafe.Sizeof(CommandBufferResetFlags(0)), unsafe.Sizeof(C.VkCommandBufferResetFlags(0)))
}

//go:generate stringer -type=PresentModeKHR
//go:generate stringer -type=CommandBufferLevel
//go:generate stringer -type=ColorSpaceKHR
//go:generate stringer -type=Format
//go:generate stringer -type=StructureType
//go:generate stringer -type=Result
//go:generate stringer -type=PhysicalDeviceType

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
	CompositeAlphaOpaqueBitKHR CompositeAlphaFlagsKHR = C.VK_COMPOSITE_ALPHA_OPAQUE_BIT_KHR

	// The alpha channel, if it exists, of the images is respected in the compositing process.
	// The non-alpha channels of the image are expected to already be multiplied by the alpha channel by the application.
	CompositeAlphaPreMultipliedBitKHR CompositeAlphaFlagsKHR = C.VK_COMPOSITE_ALPHA_PRE_MULTIPLIED_BIT_KHR

	// The alpha channel, if it exists, of the images is respected in the compositing process.
	// The non-alpha channels of the image are not expected to already be multiplied by the alpha channel by the application;
	// instead, the compositor will multiply the non-alpha channels of the image by the alpha channel during compositing.
	CompositeAlphaPostMultipliedBitKHR CompositeAlphaFlagsKHR = C.VK_COMPOSITE_ALPHA_POST_MULTIPLIED_BIT_KHR

	// The way in which the presentation engine treats the alpha channel in the images is unknown to the Vulkan API.
	// Instead, the application is responsible for setting the composite alpha blending mode using native window system commands.
	// If the application does not set the blending mode using native window system commands, then a platform-specific default will be used.
	CompositeAlphaInheritBitKHR CompositeAlphaFlagsKHR = C.VK_COMPOSITE_ALPHA_INHERIT_BIT_KHR
)

const (
	// ImageUsageTransferSrcBit specifies that the image can be used as the source of a transfer command.
	ImageUsageTransferSrcBit ImageUsageFlags = 0x00000001

	// ImageUsageTransferDstBit specifies that the image can be used as the destination of a transfer command.
	ImageUsageTransferDstBit ImageUsageFlags = 0x00000002

	// ImageUsageSampledBit specifies that the image can be used to create a VkImageView
	// suitable for occupying a VkDescriptorSet slot either of type VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE
	// or VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER, and be sampled by a shader.
	ImageUsageSampledBit ImageUsageFlags = 0x00000004

	// ImageUsageStorageBit specifies that the image can be used to create a VkImageView
	// suitable for occupying a VkDescriptorSet slot of type VK_DESCRIPTOR_TYPE_STORAGE_IMAGE.
	ImageUsageStorageBit ImageUsageFlags = 0x00000008

	// ImageUsageColorAttachmentBit specifies that the image can be used to create a VkImageView
	// suitable for use as a color or resolve attachment in a VkFramebuffer.
	ImageUsageColorAttachmentBit ImageUsageFlags = 0x00000010

	// ImageUsageDepthStencilAttachmentBit specifies that the image can be used to create a VkImageView
	// suitable for use as a depth/stencil attachment in a VkFramebuffer.
	ImageUsageDepthStencilAttachmentBit ImageUsageFlags = 0x00000020

	// ImageUsageTransientAttachmentBit specifies that the memory bound to this image will have been allocated
	// with the VK_MEMORY_PROPERTY_LAZILY_ALLOCATED_BIT (see Memory Allocation for more detail).
	// This bit can be set for any image that can be used to create a VkImageView
	// suitable for use as a color, resolve, depth/stencil, or input attachment.
	ImageUsageTransientAttachmentBit ImageUsageFlags = 0x00000040

	// ImageUsageInputAttachmentBit specifies that the image can be used to create a VkImageView suitable
	// for occupying VkDescriptorSet slot of type VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT;
	// be read from a shader as an input attachment; and be used as an input attachment in a framebuffer.
	ImageUsageInputAttachmentBit ImageUsageFlags = 0x00000080
)

const (
	FormatUNDEFINED                  Format = 0
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
	// either by calling vkResetCommandBuffer, or via the implicit reset when calling vkBeginCommandBuffer.
	// If this flag is not set on a pool, then vkResetCommandBuffer must not be called for any command buffer allocated from that pool.
	CommandPoolCreateResetCommandBufferBit CommandPoolCreateFlags = 0x00000002

	// CommandPoolCreateProtectedBit specifies that command buffers allocated from the pool are protected command buffers.
	// If the protected memory feature is not enabled, the VK_COMMAND_POOL_CREATE_PROTECTED_BIT bit of flags must not be set.
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

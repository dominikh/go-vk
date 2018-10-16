// Copyright (c) 2015-2018 Khronos Group
// Copyright (c) 2018 Dominik Honnef

package vk

// #cgo pkg-config: vulkan
// #cgo CFLAGS: -DVK_NO_PROTOTYPES
// #include <vulkan/vulkan_core.h>
import "C"
import (
	"strconv"
	"strings"
)

type DeviceQueueCreateFlags C.VkDeviceQueueCreateFlags
type QueueFlags C.VkQueueFlags
type PhysicalDeviceType C.VkPhysicalDeviceType
type Result C.VkResult
type StructureType C.VkStructureType
type SurfaceTransformFlagsKHR C.VkSurfaceTransformFlagsKHR
type CompositeAlphaFlagsKHR C.VkCompositeAlphaFlagsKHR
type ImageUsageFlags C.VkImageUsageFlags
type Format C.VkFormat
type ColorSpaceKHR C.VkColorSpaceKHR
type PresentModeKHR C.VkPresentModeKHR

const (
	DEVICE_QUEUE_CREATE_PROTECTED_BIT DeviceQueueCreateFlags = C.VK_DEVICE_QUEUE_CREATE_PROTECTED_BIT
)

const (
	QueueGraphicsBit      QueueFlags = C.VK_QUEUE_GRAPHICS_BIT
	QueueComputeBit       QueueFlags = C.VK_QUEUE_COMPUTE_BIT
	QueueTransferBit      QueueFlags = C.VK_QUEUE_TRANSFER_BIT
	QueueSparseBindingBit QueueFlags = C.VK_QUEUE_SPARSE_BINDING_BIT
	QueueProtectedBit     QueueFlags = C.VK_QUEUE_PROTECTED_BIT
)

const (
	// The device does not match any other available types.
	PhysicalDeviceTypeOther PhysicalDeviceType = C.VK_PHYSICAL_DEVICE_TYPE_OTHER
	// The device is typically one embedded in or tightly coupled with the host.
	PhysicalDeviceTypeIntegratedGPU PhysicalDeviceType = C.VK_PHYSICAL_DEVICE_TYPE_INTEGRATED_GPU
	// The device is typically a separate processor connected to the host via an interlink.
	PhysicalDeviceTypeDiscreteGPU PhysicalDeviceType = C.VK_PHYSICAL_DEVICE_TYPE_DISCRETE_GPU
	// The device is typically a virtual node in a virtualization environment.
	PhysicalDeviceTypeVirtualGPU PhysicalDeviceType = C.VK_PHYSICAL_DEVICE_TYPE_VIRTUAL_GPU
	// The device is typically running on the same processors as the host.
	PhysicalDeviceTypeCPU PhysicalDeviceType = C.VK_PHYSICAL_DEVICE_TYPE_CPU
)

const (
	// Command successfully completed.
	Success Result = C.VK_SUCCESS

	// A fence or query has not yet completed.
	NotReady Result = C.VK_NOT_READY

	// A wait operation has not completed in the specified time.
	Timeout Result = C.VK_TIMEOUT

	// An event is signaled.
	EventSet Result = C.VK_EVENT_SET

	// An event is unsignaled.
	EventReset Result = C.VK_EVENT_RESET

	// A return array was too small for the result.
	Incomplete Result = C.VK_INCOMPLETE

	// A swapchain no longer matches the surface properties exactly, but can still be used to present to the surface successfully.
	SuboptimalKHR Result = C.VK_SUBOPTIMAL_KHR

	// A host memory allocation has failed.
	ErrOutOfHostMemory Result = C.VK_ERROR_OUT_OF_HOST_MEMORY

	// A device memory allocation has failed.
	ErrOutOfDeviceMemory Result = C.VK_ERROR_OUT_OF_DEVICE_MEMORY

	// Initialization of an object could not be completed for implementation-specific reasons.
	ErrInitializationFailed Result = C.VK_ERROR_INITIALIZATION_FAILED

	// The logical or physical device has been lost.
	ErrDeviceLost Result = C.VK_ERROR_DEVICE_LOST

	// Mapping of a memory object has failed.
	ErrMemoryMapFailed Result = C.VK_ERROR_MEMORY_MAP_FAILED

	// A requested layer is not present or could not be loaded.
	ErrLayerNotPresent Result = C.VK_ERROR_LAYER_NOT_PRESENT

	// A requested extension is not supported.
	ErrExtensionNotPresent Result = C.VK_ERROR_EXTENSION_NOT_PRESENT

	// A requested feature is not supported.
	ErrFeatureNotPresent Result = C.VK_ERROR_FEATURE_NOT_PRESENT

	// The requested version of Vulkan is not supported by the driver or is otherwise incompatible for implementation-specific reasons.
	ErrIncompatibleDriver Result = C.VK_ERROR_INCOMPATIBLE_DRIVER

	// Too many objects of the type have already been created.
	ErrTooManyObjects Result = C.VK_ERROR_TOO_MANY_OBJECTS

	// A requested format is not supported on this device.
	ErrFormatNotSupported Result = C.VK_ERROR_FORMAT_NOT_SUPPORTED

	// A pool allocation has failed due to fragmentation of the pool’s memory.
	ErrFragmentedPool Result = C.VK_ERROR_FRAGMENTED_POOL

	// A surface is no longer available.
	ErrSurfaceLostKHR Result = C.VK_ERROR_SURFACE_LOST_KHR

	// The requested window is already in use by Vulkan or another API in a manner which prevents it from being used again.
	ErrNativeWindowInUseKHR Result = C.VK_ERROR_NATIVE_WINDOW_IN_USE_KHR

	// A surface has changed in such a way that it is no longer compatible with the swapchain,
	// and further presentation requests using the swapchain will fail.
	// Applications must query the new surface properties and recreate their swapchain if they wish to continue presenting to the surface.
	ErrOutOfDateKHR Result = C.VK_ERROR_OUT_OF_DATE_KHR

	// The display used by a swapchain does not use the same presentable image layout,
	// or is incompatible in a way that prevents sharing an image.
	ErrIncompatibleDisplayKHR Result = C.VK_ERROR_INCOMPATIBLE_DISPLAY_KHR

	// A pool memory allocation has failed.
	ErrOutOfPoolMemory Result = C.VK_ERROR_OUT_OF_POOL_MEMORY

	// An external handle is not a valid handle of the specified type.
	ErrInvalidExternalHandle Result = C.VK_ERROR_INVALID_EXTERNAL_HANDLE

	ErrValidationFailedEXT      Result = C.VK_ERROR_VALIDATION_FAILED_EXT
	ErrInvalidShaderNV          Result = C.VK_ERROR_INVALID_SHADER_NV
	ErrFragmentationEXT         Result = C.VK_ERROR_FRAGMENTATION_EXT
	ErrNotPermittedEXT          Result = C.VK_ERROR_NOT_PERMITTED_EXT
	ErrOutOfPoolMemoryKHR       Result = C.VK_ERROR_OUT_OF_POOL_MEMORY_KHR
	ErrInvalidExternalHandleKHR Result = C.VK_ERROR_INVALID_EXTERNAL_HANDLE_KHR
)

const (
	StructureTypeAccelerationStructureCreateInfoNVX                    = C.VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_CREATE_INFO_NVX
	StructureTypeAccelerationStructureMemoryRequirementsInfoNVX        = C.VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_INFO_NVX
	StructureTypeAcquireNextImageInfoKHR                               = C.VK_STRUCTURE_TYPE_ACQUIRE_NEXT_IMAGE_INFO_KHR
	StructureTypeAndroidHardwareBufferFormatPropertiesAndroid          = C.VK_STRUCTURE_TYPE_ANDROID_HARDWARE_BUFFER_FORMAT_PROPERTIES_ANDROID
	StructureTypeAndroidHardwareBufferPropertiesAndroid                = C.VK_STRUCTURE_TYPE_ANDROID_HARDWARE_BUFFER_PROPERTIES_ANDROID
	StructureTypeAndroidHardwareBufferUsageAndroid                     = C.VK_STRUCTURE_TYPE_ANDROID_HARDWARE_BUFFER_USAGE_ANDROID
	StructureTypeAndroidSurfaceCreateInfoKHR                           = C.VK_STRUCTURE_TYPE_ANDROID_SURFACE_CREATE_INFO_KHR
	StructureTypeAttachmentDescription2KHR                             = C.VK_STRUCTURE_TYPE_ATTACHMENT_DESCRIPTION_2_KHR
	StructureTypeAttachmentReference2KHR                               = C.VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_2_KHR
	StructureTypeApplicationInfo                                       = C.VK_STRUCTURE_TYPE_APPLICATION_INFO
	StructureTypeBindAccelerationStructureMemoryInfoNVX                = C.VK_STRUCTURE_TYPE_BIND_ACCELERATION_STRUCTURE_MEMORY_INFO_NVX
	StructureTypeBindBufferMemoryDeviceGroupInfoKHR                    = C.VK_STRUCTURE_TYPE_BIND_BUFFER_MEMORY_DEVICE_GROUP_INFO_KHR
	StructureTypeBindBufferMemoryInfoKHR                               = C.VK_STRUCTURE_TYPE_BIND_BUFFER_MEMORY_INFO_KHR
	StructureTypeBindImageMemoryDeviceGroupInfoKHR                     = C.VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_DEVICE_GROUP_INFO_KHR
	StructureTypeBindImageMemoryInfoKHR                                = C.VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_INFO_KHR
	StructureTypeBindImageMemorySwapchainInfoKHR                       = C.VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_SWAPCHAIN_INFO_KHR
	StructureTypeBindImagePlaneMemoryINFO                              = C.VK_STRUCTURE_TYPE_BIND_IMAGE_PLANE_MEMORY_INFO
	StructureTypeBindImagePlaneMemoryInfoKHR                           = C.VK_STRUCTURE_TYPE_BIND_IMAGE_PLANE_MEMORY_INFO_KHR
	StructureTypeBufferMemoryRequirementsInfo2                         = C.VK_STRUCTURE_TYPE_BUFFER_MEMORY_REQUIREMENTS_INFO_2
	StructureTypeBufferMemoryRequirementsInfo2KHR                      = C.VK_STRUCTURE_TYPE_BUFFER_MEMORY_REQUIREMENTS_INFO_2_KHR
	StructureTypeBindBufferMemoryDeviceGroupInfo                       = C.VK_STRUCTURE_TYPE_BIND_BUFFER_MEMORY_DEVICE_GROUP_INFO
	StructureTypeBindBufferMemoryInfo                                  = C.VK_STRUCTURE_TYPE_BIND_BUFFER_MEMORY_INFO
	StructureTypeBindImageMemoryDeviceGroupInfo                        = C.VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_DEVICE_GROUP_INFO
	StructureTypeBindImageMemoryInfo                                   = C.VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_INFO
	StructureTypeBindSparseInfo                                        = C.VK_STRUCTURE_TYPE_BIND_SPARSE_INFO
	StructureTypeBufferCreateInfo                                      = C.VK_STRUCTURE_TYPE_BUFFER_CREATE_INFO
	StructureTypeBufferMemoryBarrier                                   = C.VK_STRUCTURE_TYPE_BUFFER_MEMORY_BARRIER
	StructureTypeBufferViewCreateInfo                                  = C.VK_STRUCTURE_TYPE_BUFFER_VIEW_CREATE_INFO
	StructureTypeCheckpointDataNV                                      = C.VK_STRUCTURE_TYPE_CHECKPOINT_DATA_NV
	StructureTypeCmdProcessCommandsInfoNVX                             = C.VK_STRUCTURE_TYPE_CMD_PROCESS_COMMANDS_INFO_NVX
	StructureTypeCmdReserveSpaceForCommandsInfoNVX                     = C.VK_STRUCTURE_TYPE_CMD_RESERVE_SPACE_FOR_COMMANDS_INFO_NVX
	StructureTypeCommandBufferInheritanceConditionalRenderingInfoEXT   = C.VK_STRUCTURE_TYPE_COMMAND_BUFFER_INHERITANCE_CONDITIONAL_RENDERING_INFO_EXT
	StructureTypeConditionalRenderingBeginInfoEXT                      = C.VK_STRUCTURE_TYPE_CONDITIONAL_RENDERING_BEGIN_INFO_EXT
	StructureTypeCommandBufferAllocateInfo                             = C.VK_STRUCTURE_TYPE_COMMAND_BUFFER_ALLOCATE_INFO
	StructureTypeCommandBufferBeginInfo                                = C.VK_STRUCTURE_TYPE_COMMAND_BUFFER_BEGIN_INFO
	StructureTypeCommandBufferInheritanceInfo                          = C.VK_STRUCTURE_TYPE_COMMAND_BUFFER_INHERITANCE_INFO
	StructureTypeCommandPoolCreateInfo                                 = C.VK_STRUCTURE_TYPE_COMMAND_POOL_CREATE_INFO
	StructureTypeComputePipelineCreateInfo                             = C.VK_STRUCTURE_TYPE_COMPUTE_PIPELINE_CREATE_INFO
	StructureTypeCopyDescriptorSet                                     = C.VK_STRUCTURE_TYPE_COPY_DESCRIPTOR_SET
	StructureTypeD3D12FenceSubmitInfoKHR                               = C.VK_STRUCTURE_TYPE_D3D12_FENCE_SUBMIT_INFO_KHR
	StructureTypeDebugMarkerMarkerInfoEXT                              = C.VK_STRUCTURE_TYPE_DEBUG_MARKER_MARKER_INFO_EXT
	StructureTypeDebugMarkerObjectNameInfoEXT                          = C.VK_STRUCTURE_TYPE_DEBUG_MARKER_OBJECT_NAME_INFO_EXT
	StructureTypeDebugMarkerObjectTagInfoEXT                           = C.VK_STRUCTURE_TYPE_DEBUG_MARKER_OBJECT_TAG_INFO_EXT
	StructureTypeDebugReportCallbackCreateInfoEXT                      = C.VK_STRUCTURE_TYPE_DEBUG_REPORT_CALLBACK_CREATE_INFO_EXT
	StructureTypeDebugReportCreateInfoEXT                              = C.VK_STRUCTURE_TYPE_DEBUG_REPORT_CREATE_INFO_EXT
	StructureTypeDebugUtilsLabelEXT                                    = C.VK_STRUCTURE_TYPE_DEBUG_UTILS_LABEL_EXT
	StructureTypeDebugUtilsMessengerCallbackDataEXT                    = C.VK_STRUCTURE_TYPE_DEBUG_UTILS_MESSENGER_CALLBACK_DATA_EXT
	StructureTypeDebugUtilsMessengerCreateInfoEXT                      = C.VK_STRUCTURE_TYPE_DEBUG_UTILS_MESSENGER_CREATE_INFO_EXT
	StructureTypeDebugUtilsObjectNameInfoEXT                           = C.VK_STRUCTURE_TYPE_DEBUG_UTILS_OBJECT_NAME_INFO_EXT
	StructureTypeDebugUtilsObjectTagInfoEXT                            = C.VK_STRUCTURE_TYPE_DEBUG_UTILS_OBJECT_TAG_INFO_EXT
	StructureTypeDedicatedAllocationBufferCreateInfoNV                 = C.VK_STRUCTURE_TYPE_DEDICATED_ALLOCATION_BUFFER_CREATE_INFO_NV
	StructureTypeDedicatedAllocationImageCreateInfoNV                  = C.VK_STRUCTURE_TYPE_DEDICATED_ALLOCATION_IMAGE_CREATE_INFO_NV
	StructureTypeDedicatedAllocationMemoryAllocateInfoNV               = C.VK_STRUCTURE_TYPE_DEDICATED_ALLOCATION_MEMORY_ALLOCATE_INFO_NV
	StructureTypeDescriptorAccelerationStructureInfoNVX                = C.VK_STRUCTURE_TYPE_DESCRIPTOR_ACCELERATION_STRUCTURE_INFO_NVX
	StructureTypeDescriptorPoolInlineUniformBlockCreateInfoEXT         = C.VK_STRUCTURE_TYPE_DESCRIPTOR_POOL_INLINE_UNIFORM_BLOCK_CREATE_INFO_EXT
	StructureTypeDescriptorSetLayoutBindingFlagsCreateInfoEXT          = C.VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_BINDING_FLAGS_CREATE_INFO_EXT
	StructureTypeDescriptorSetLayoutSupport                            = C.VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_SUPPORT
	StructureTypeDescriptorSetLayoutSupportKHR                         = C.VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_SUPPORT_KHR
	StructureTypeDescriptorSetVariableDescriptorCountAllocateInfoExt   = C.VK_STRUCTURE_TYPE_DESCRIPTOR_SET_VARIABLE_DESCRIPTOR_COUNT_ALLOCATE_INFO_EXT
	StructureTypeDescriptorSetVariableDescriptorCountLayoutSupportEXT  = C.VK_STRUCTURE_TYPE_DESCRIPTOR_SET_VARIABLE_DESCRIPTOR_COUNT_LAYOUT_SUPPORT_EXT
	StructureTypeDescriptorUpdateTemplateCreateINFO                    = C.VK_STRUCTURE_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_CREATE_INFO
	StructureTypeDescriptorUpdateTemplateCreateInfoKHR                 = C.VK_STRUCTURE_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_CREATE_INFO_KHR
	StructureTypeDeviceEventInfoEXT                                    = C.VK_STRUCTURE_TYPE_DEVICE_EVENT_INFO_EXT
	StructureTypeDeviceGeneratedCommandsFeaturesNVX                    = C.VK_STRUCTURE_TYPE_DEVICE_GENERATED_COMMANDS_FEATURES_NVX
	StructureTypeDeviceGeneratedCommandsLimitsNVX                      = C.VK_STRUCTURE_TYPE_DEVICE_GENERATED_COMMANDS_LIMITS_NVX
	StructureTypeDeviceGroupBindSparseInfoKHR                          = C.VK_STRUCTURE_TYPE_DEVICE_GROUP_BIND_SPARSE_INFO_KHR
	StructureTypeDeviceGroupCommandBufferBeginInfoKHR                  = C.VK_STRUCTURE_TYPE_DEVICE_GROUP_COMMAND_BUFFER_BEGIN_INFO_KHR
	StructureTypeDeviceGroupDeviceCreateInfoKHR                        = C.VK_STRUCTURE_TYPE_DEVICE_GROUP_DEVICE_CREATE_INFO_KHR
	StructureTypeDeviceGroupPresentCapabilitiesKHR                     = C.VK_STRUCTURE_TYPE_DEVICE_GROUP_PRESENT_CAPABILITIES_KHR
	StructureTypeDeviceGroupPresentInfoKHR                             = C.VK_STRUCTURE_TYPE_DEVICE_GROUP_PRESENT_INFO_KHR
	StructureTypeDeviceGroupRenderPassBeginInfoKHR                     = C.VK_STRUCTURE_TYPE_DEVICE_GROUP_RENDER_PASS_BEGIN_INFO_KHR
	StructureTypeDeviceGroupSubmitInfoKHR                              = C.VK_STRUCTURE_TYPE_DEVICE_GROUP_SUBMIT_INFO_KHR
	StructureTypeDeviceGroupSwapchainCreateInfoKHR                     = C.VK_STRUCTURE_TYPE_DEVICE_GROUP_SWAPCHAIN_CREATE_INFO_KHR
	StructureTypeDeviceQueueGlobalPriorityCreateInfoEXT                = C.VK_STRUCTURE_TYPE_DEVICE_QUEUE_GLOBAL_PRIORITY_CREATE_INFO_EXT
	StructureTypeDeviceQueueInfo2                                      = C.VK_STRUCTURE_TYPE_DEVICE_QUEUE_INFO_2
	StructureTypeDisplayEventInfoEXT                                   = C.VK_STRUCTURE_TYPE_DISPLAY_EVENT_INFO_EXT
	StructureTypeDisplayModeCreateInfoKHR                              = C.VK_STRUCTURE_TYPE_DISPLAY_MODE_CREATE_INFO_KHR
	StructureTypeDisplayModeProperties2KHR                             = C.VK_STRUCTURE_TYPE_DISPLAY_MODE_PROPERTIES_2_KHR
	StructureTypeDisplayPlaneCapabilities2KHR                          = C.VK_STRUCTURE_TYPE_DISPLAY_PLANE_CAPABILITIES_2_KHR
	StructureTypeDisplayPlaneInfo2KHR                                  = C.VK_STRUCTURE_TYPE_DISPLAY_PLANE_INFO_2_KHR
	StructureTypeDisplayPlaneProperties2KHR                            = C.VK_STRUCTURE_TYPE_DISPLAY_PLANE_PROPERTIES_2_KHR
	StructureTypeDisplayPowerInfoEXT                                   = C.VK_STRUCTURE_TYPE_DISPLAY_POWER_INFO_EXT
	StructureTypeDisplayPresentInfoKHR                                 = C.VK_STRUCTURE_TYPE_DISPLAY_PRESENT_INFO_KHR
	StructureTypeDisplayProperties2KHR                                 = C.VK_STRUCTURE_TYPE_DISPLAY_PROPERTIES_2_KHR
	StructureTypeDisplaySurfaceCreateInfoKHR                           = C.VK_STRUCTURE_TYPE_DISPLAY_SURFACE_CREATE_INFO_KHR
	StructureTypeDescriptorPoolCreateInfo                              = C.VK_STRUCTURE_TYPE_DESCRIPTOR_POOL_CREATE_INFO
	StructureTypeDescriptorSetAllocateInfo                             = C.VK_STRUCTURE_TYPE_DESCRIPTOR_SET_ALLOCATE_INFO
	StructureTypeDescriptorSetLayoutCreateInfo                         = C.VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_CREATE_INFO
	StructureTypeDeviceCreateInfo                                      = C.VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO
	StructureTypeDeviceGroupBindSparseInfo                             = C.VK_STRUCTURE_TYPE_DEVICE_GROUP_BIND_SPARSE_INFO
	StructureTypeDeviceGroupCommandBufferBeginInfo                     = C.VK_STRUCTURE_TYPE_DEVICE_GROUP_COMMAND_BUFFER_BEGIN_INFO
	StructureTypeDeviceGroupDeviceCreateInfo                           = C.VK_STRUCTURE_TYPE_DEVICE_GROUP_DEVICE_CREATE_INFO
	StructureTypeDeviceGroupRenderPassBeginInfo                        = C.VK_STRUCTURE_TYPE_DEVICE_GROUP_RENDER_PASS_BEGIN_INFO
	StructureTypeDeviceGroupSubmitInfo                                 = C.VK_STRUCTURE_TYPE_DEVICE_GROUP_SUBMIT_INFO
	StructureTypeDeviceQueueCreateInfo                                 = C.VK_STRUCTURE_TYPE_DEVICE_QUEUE_CREATE_INFO
	StructureTypeExportFenceCreateInfo                                 = C.VK_STRUCTURE_TYPE_EXPORT_FENCE_CREATE_INFO
	StructureTypeExportFenceCreateInfoKHR                              = C.VK_STRUCTURE_TYPE_EXPORT_FENCE_CREATE_INFO_KHR
	StructureTypeExportFenceWin32HandleInfoKHR                         = C.VK_STRUCTURE_TYPE_EXPORT_FENCE_WIN32_HANDLE_INFO_KHR
	StructureTypeExportMemoryAllocateInfo                              = C.VK_STRUCTURE_TYPE_EXPORT_MEMORY_ALLOCATE_INFO
	StructureTypeExportMemoryAllocateInfoKHR                           = C.VK_STRUCTURE_TYPE_EXPORT_MEMORY_ALLOCATE_INFO_KHR
	StructureTypeExportMemoryAllocateInfoNV                            = C.VK_STRUCTURE_TYPE_EXPORT_MEMORY_ALLOCATE_INFO_NV
	StructureTypeExportMemoryWin32HandleInfoKHR                        = C.VK_STRUCTURE_TYPE_EXPORT_MEMORY_WIN32_HANDLE_INFO_KHR
	StructureTypeExportMemoryWin32HandleInfoNV                         = C.VK_STRUCTURE_TYPE_EXPORT_MEMORY_WIN32_HANDLE_INFO_NV
	StructureTypeExportSemaphoreCreateInfo                             = C.VK_STRUCTURE_TYPE_EXPORT_SEMAPHORE_CREATE_INFO
	StructureTypeExportSemaphoreCreateInfoKHR                          = C.VK_STRUCTURE_TYPE_EXPORT_SEMAPHORE_CREATE_INFO_KHR
	StructureTypeExportSemaphoreWin32HandleInfoKHR                     = C.VK_STRUCTURE_TYPE_EXPORT_SEMAPHORE_WIN32_HANDLE_INFO_KHR
	StructureTypeExternalBufferProperties                              = C.VK_STRUCTURE_TYPE_EXTERNAL_BUFFER_PROPERTIES
	StructureTypeExternalBufferPropertiesKHR                           = C.VK_STRUCTURE_TYPE_EXTERNAL_BUFFER_PROPERTIES_KHR
	StructureTypeExternalFenceProperties                               = C.VK_STRUCTURE_TYPE_EXTERNAL_FENCE_PROPERTIES
	StructureTypeExternalFencePropertiesKHR                            = C.VK_STRUCTURE_TYPE_EXTERNAL_FENCE_PROPERTIES_KHR
	StructureTypeExternalFormatAndroid                                 = C.VK_STRUCTURE_TYPE_EXTERNAL_FORMAT_ANDROID
	StructureTypeExternalImageFormatProperties                         = C.VK_STRUCTURE_TYPE_EXTERNAL_IMAGE_FORMAT_PROPERTIES
	StructureTypeExternalImageFormatPropertiesKHR                      = C.VK_STRUCTURE_TYPE_EXTERNAL_IMAGE_FORMAT_PROPERTIES_KHR
	StructureTypeExternalMemoryBufferCreateInfo                        = C.VK_STRUCTURE_TYPE_EXTERNAL_MEMORY_BUFFER_CREATE_INFO
	StructureTypeExternalMemoryBufferCreateInfoKHR                     = C.VK_STRUCTURE_TYPE_EXTERNAL_MEMORY_BUFFER_CREATE_INFO_KHR
	StructureTypeExternalMemoryImageCreateInfo                         = C.VK_STRUCTURE_TYPE_EXTERNAL_MEMORY_IMAGE_CREATE_INFO
	StructureTypeExternalMemoryImageCreateInfoKHR                      = C.VK_STRUCTURE_TYPE_EXTERNAL_MEMORY_IMAGE_CREATE_INFO_KHR
	StructureTypeExternalMemoryImageCreateInfoNV                       = C.VK_STRUCTURE_TYPE_EXTERNAL_MEMORY_IMAGE_CREATE_INFO_NV
	StructureTypeExternalSemaphoreProperties                           = C.VK_STRUCTURE_TYPE_EXTERNAL_SEMAPHORE_PROPERTIES
	StructureTypeExternalSemaphorePropertiesKHR                        = C.VK_STRUCTURE_TYPE_EXTERNAL_SEMAPHORE_PROPERTIES_KHR
	StructureTypeEventCreateInfo                                       = C.VK_STRUCTURE_TYPE_EVENT_CREATE_INFO
	StructureTypeFenceGetFdInfoKHR                                     = C.VK_STRUCTURE_TYPE_FENCE_GET_FD_INFO_KHR
	StructureTypeFenceGetWin32HandleInfoKHR                            = C.VK_STRUCTURE_TYPE_FENCE_GET_WIN32_HANDLE_INFO_KHR
	StructureTypeFormatProperties2                                     = C.VK_STRUCTURE_TYPE_FORMAT_PROPERTIES_2
	StructureTypeFormatProperties2KHR                                  = C.VK_STRUCTURE_TYPE_FORMAT_PROPERTIES_2_KHR
	StructureTypeFenceCreateInfo                                       = C.VK_STRUCTURE_TYPE_FENCE_CREATE_INFO
	StructureTypeFramebufferCreateInfo                                 = C.VK_STRUCTURE_TYPE_FRAMEBUFFER_CREATE_INFO
	StructureTypeGeometryAABB_NVX                                      = C.VK_STRUCTURE_TYPE_GEOMETRY_AABB_NVX
	StructureTypeGeometryInstanceNVX                                   = C.VK_STRUCTURE_TYPE_GEOMETRY_INSTANCE_NVX
	StructureTypeGeometryNVX                                           = C.VK_STRUCTURE_TYPE_GEOMETRY_NVX
	StructureTypeGeometryTrianglesNVX                                  = C.VK_STRUCTURE_TYPE_GEOMETRY_TRIANGLES_NVX
	StructureTypeGraphicsPipelineCreateInfo                            = C.VK_STRUCTURE_TYPE_GRAPHICS_PIPELINE_CREATE_INFO
	StructureTypeHdrMetadataEXT                                        = C.VK_STRUCTURE_TYPE_HDR_METADATA_EXT
	StructureTypeHitShaderModuleCreateInfoNVX                          = C.VK_STRUCTURE_TYPE_HIT_SHADER_MODULE_CREATE_INFO_NVX
	StructureTypeImageFormatListCreateInfoKHR                          = C.VK_STRUCTURE_TYPE_IMAGE_FORMAT_LIST_CREATE_INFO_KHR
	StructureTypeImageFormatProperties2                                = C.VK_STRUCTURE_TYPE_IMAGE_FORMAT_PROPERTIES_2
	StructureTypeImageFormatProperties2KHR                             = C.VK_STRUCTURE_TYPE_IMAGE_FORMAT_PROPERTIES_2_KHR
	StructureTypeImageMemoryRequirementsInfo2                          = C.VK_STRUCTURE_TYPE_IMAGE_MEMORY_REQUIREMENTS_INFO_2
	StructureTypeImageMemoryRequirementsInfo2KHR                       = C.VK_STRUCTURE_TYPE_IMAGE_MEMORY_REQUIREMENTS_INFO_2_KHR
	StructureTypeImagePlaneMemoryRequirementsInfo                      = C.VK_STRUCTURE_TYPE_IMAGE_PLANE_MEMORY_REQUIREMENTS_INFO
	StructureTypeImagePlaneMemoryRequirementsInfoKHR                   = C.VK_STRUCTURE_TYPE_IMAGE_PLANE_MEMORY_REQUIREMENTS_INFO_KHR
	StructureTypeImageSparseMemoryRequirementsInfo2                    = C.VK_STRUCTURE_TYPE_IMAGE_SPARSE_MEMORY_REQUIREMENTS_INFO_2
	StructureTypeImageSparseMemoryRequirementsInfo2KHR                 = C.VK_STRUCTURE_TYPE_IMAGE_SPARSE_MEMORY_REQUIREMENTS_INFO_2_KHR
	StructureTypeImageSwapchainCreateInfoKHR                           = C.VK_STRUCTURE_TYPE_IMAGE_SWAPCHAIN_CREATE_INFO_KHR
	StructureTypeImageViewAstcDecodeModeEXT                            = C.VK_STRUCTURE_TYPE_IMAGE_VIEW_ASTC_DECODE_MODE_EXT
	StructureTypeImageViewUsageCreateInfo                              = C.VK_STRUCTURE_TYPE_IMAGE_VIEW_USAGE_CREATE_INFO
	StructureTypeImageViewUsageCreateInfoKHR                           = C.VK_STRUCTURE_TYPE_IMAGE_VIEW_USAGE_CREATE_INFO_KHR
	StructureTypeImportAndroidHardwareBufferInfoAndroid                = C.VK_STRUCTURE_TYPE_IMPORT_ANDROID_HARDWARE_BUFFER_INFO_ANDROID
	StructureTypeImportFenceFdInfoKHR                                  = C.VK_STRUCTURE_TYPE_IMPORT_FENCE_FD_INFO_KHR
	StructureTypeImportFenceWin32HandleInfoKHR                         = C.VK_STRUCTURE_TYPE_IMPORT_FENCE_WIN32_HANDLE_INFO_KHR
	StructureTypeImportMemoryFdInfoKHR                                 = C.VK_STRUCTURE_TYPE_IMPORT_MEMORY_FD_INFO_KHR
	StructureTypeImportMemoryHostPointerInfoEXT                        = C.VK_STRUCTURE_TYPE_IMPORT_MEMORY_HOST_POINTER_INFO_EXT
	StructureTypeImportMemoryWin32HandleInfoKHR                        = C.VK_STRUCTURE_TYPE_IMPORT_MEMORY_WIN32_HANDLE_INFO_KHR
	StructureTypeImportMemoryWin32HandleInfoNV                         = C.VK_STRUCTURE_TYPE_IMPORT_MEMORY_WIN32_HANDLE_INFO_NV
	StructureTypeImportSemaphoreFdInfoKHR                              = C.VK_STRUCTURE_TYPE_IMPORT_SEMAPHORE_FD_INFO_KHR
	StructureTypeImportSemaphoreWin32HandleInfoKHR                     = C.VK_STRUCTURE_TYPE_IMPORT_SEMAPHORE_WIN32_HANDLE_INFO_KHR
	StructureTypeIndirectCommandsLayoutCreateInfoNVX                   = C.VK_STRUCTURE_TYPE_INDIRECT_COMMANDS_LAYOUT_CREATE_INFO_NVX
	StructureTypeIosSurfaceCreateInfoMVK                               = C.VK_STRUCTURE_TYPE_IOS_SURFACE_CREATE_INFO_MVK
	StructureTypeImageCreateInfo                                       = C.VK_STRUCTURE_TYPE_IMAGE_CREATE_INFO
	StructureTypeImageMemoryBarrier                                    = C.VK_STRUCTURE_TYPE_IMAGE_MEMORY_BARRIER
	StructureTypeImageViewCreateInfo                                   = C.VK_STRUCTURE_TYPE_IMAGE_VIEW_CREATE_INFO
	StructureTypeInstanceCreateInfo                                    = C.VK_STRUCTURE_TYPE_INSTANCE_CREATE_INFO
	StructureTypeLoaderDeviceCreateInfo                                = C.VK_STRUCTURE_TYPE_LOADER_DEVICE_CREATE_INFO
	StructureTypeLoaderInstanceCreateInfo                              = C.VK_STRUCTURE_TYPE_LOADER_INSTANCE_CREATE_INFO
	StructureTypeMacosSurfaceCreateInfoMVK                             = C.VK_STRUCTURE_TYPE_MACOS_SURFACE_CREATE_INFO_MVK
	StructureTypeMemoryAllocateFlagsInfoKHR                            = C.VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_FLAGS_INFO_KHR
	StructureTypeMemoryDedicatedAllocateInfoKHR                        = C.VK_STRUCTURE_TYPE_MEMORY_DEDICATED_ALLOCATE_INFO_KHR
	StructureTypeMemoryDedicatedRequirementsKHR                        = C.VK_STRUCTURE_TYPE_MEMORY_DEDICATED_REQUIREMENTS_KHR
	StructureTypeMemoryFdPropertiesKHR                                 = C.VK_STRUCTURE_TYPE_MEMORY_FD_PROPERTIES_KHR
	StructureTypeMemoryGetAndroidHardwareBufferInfoAndroid             = C.VK_STRUCTURE_TYPE_MEMORY_GET_ANDROID_HARDWARE_BUFFER_INFO_ANDROID
	StructureTypeMemoryGetFdInfoKHR                                    = C.VK_STRUCTURE_TYPE_MEMORY_GET_FD_INFO_KHR
	StructureTypeMemoryGetWin32HandleInfoKHR                           = C.VK_STRUCTURE_TYPE_MEMORY_GET_WIN32_HANDLE_INFO_KHR
	StructureTypeMemoryHostPointerPropertiesEXT                        = C.VK_STRUCTURE_TYPE_MEMORY_HOST_POINTER_PROPERTIES_EXT
	StructureTypeMemoryRequirements2                                   = C.VK_STRUCTURE_TYPE_MEMORY_REQUIREMENTS_2
	StructureTypeMemoryRequirements2KHR                                = C.VK_STRUCTURE_TYPE_MEMORY_REQUIREMENTS_2_KHR
	StructureTypeMemoryWin32HandlePropertiesKHR                        = C.VK_STRUCTURE_TYPE_MEMORY_WIN32_HANDLE_PROPERTIES_KHR
	StructureTypeMirSurfaceCreateInfoKHR                               = C.VK_STRUCTURE_TYPE_MIR_SURFACE_CREATE_INFO_KHR
	StructureTypeMultisamplePropertiesEXT                              = C.VK_STRUCTURE_TYPE_MULTISAMPLE_PROPERTIES_EXT
	StructureTypeMappedMemoryRange                                     = C.VK_STRUCTURE_TYPE_MAPPED_MEMORY_RANGE
	StructureTypeMemoryAllocateFlagsInfo                               = C.VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_FLAGS_INFO
	StructureTypeMemoryAllocateInfo                                    = C.VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_INFO
	StructureTypeMemoryBarrier                                         = C.VK_STRUCTURE_TYPE_MEMORY_BARRIER
	StructureTypeMemoryDedicatedAllocateInfo                           = C.VK_STRUCTURE_TYPE_MEMORY_DEDICATED_ALLOCATE_INFO
	StructureTypeMemoryDedicatedRequirements                           = C.VK_STRUCTURE_TYPE_MEMORY_DEDICATED_REQUIREMENTS
	StructureTypeObjectTableCreateInfoNVX                              = C.VK_STRUCTURE_TYPE_OBJECT_TABLE_CREATE_INFO_NVX
	StructureTypePhysicalDevice16BitStorageFeaturesKHR                 = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_16BIT_STORAGE_FEATURES_KHR
	StructureTypePhysicalDevice8BitStorageFeaturesKHR                  = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_8BIT_STORAGE_FEATURES_KHR
	StructureTypePhysicalDeviceAstcDecodeFeaturesEXT                   = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ASTC_DECODE_FEATURES_EXT
	StructureTypePhysicalDeviceBlendOperationAdvancedFeaturesEXT       = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_BLEND_OPERATION_ADVANCED_FEATURES_EXT
	StructureTypePhysicalDeviceBlendOperationAdvancedPropertiesEXT     = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_BLEND_OPERATION_ADVANCED_PROPERTIES_EXT
	StructureTypePhysicalDeviceComputeShaderDerivativesFeatures_NV     = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_COMPUTE_SHADER_DERIVATIVES_FEATURES_NV
	StructureTypePhysicalDeviceConditionalRenderingFeaturesEXT         = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_CONDITIONAL_RENDERING_FEATURES_EXT
	StructureTypePhysicalDeviceConservativeRasterizationPropertiesEXT  = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_CONSERVATIVE_RASTERIZATION_PROPERTIES_EXT
	StructureTypePhysicalDeviceCornerSampledImageFeaturesNV            = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_CORNER_SAMPLED_IMAGE_FEATURES_NV
	StructureTypePhysicalDeviceDescriptorIndexingFeaturesEXT           = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_INDEXING_FEATURES_EXT
	StructureTypePhysicalDeviceDescriptorIndexingPropertiesEXT         = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_INDEXING_PROPERTIES_EXT
	StructureTypePhysicalDeviceDiscardRectanglePropertiesEXT           = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DISCARD_RECTANGLE_PROPERTIES_EXT
	StructureTypePhysicalDeviceExclusiveScissorFeaturesNV              = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXCLUSIVE_SCISSOR_FEATURES_NV
	StructureTypePhysicalDeviceExternalBufferInfo                      = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_BUFFER_INFO
	StructureTypePhysicalDeviceExternalBufferInfoKHR                   = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_BUFFER_INFO_KHR
	StructureTypePhysicalDeviceExternalFenceInfo                       = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_FENCE_INFO
	StructureTypePhysicalDeviceExternalFenceInfoKHR                    = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_FENCE_INFO_KHR
	StructureTypePhysicalDeviceExternalImageFormatInfo                 = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_IMAGE_FORMAT_INFO
	StructureTypePhysicalDeviceExternalImageFormatInfoKHR              = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_IMAGE_FORMAT_INFO_KHR
	StructureTypePhysicalDeviceExternalMemoryHostPropertiesEXT         = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_MEMORY_HOST_PROPERTIES_EXT
	StructureTypePhysicalDeviceExternalSemaphoreInfo                   = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_SEMAPHORE_INFO
	StructureTypePhysicalDeviceExternalSemaphoreInfoKHR                = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_SEMAPHORE_INFO_KHR
	StructureTypePhysicalDeviceFeatures2                               = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2
	StructureTypePhysicalDeviceFeatures2KHR                            = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2_KHR
	StructureTypePhysicalDeviceFragmentShaderBarycentricFeaturesNV     = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_SHADER_BARYCENTRIC_FEATURES_NV
	StructureTypePhysicalDeviceGroupPropertiesKHR                      = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES_KHR
	StructureTypePhysicalDeviceIdProperties                            = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ID_PROPERTIES
	StructureTypePhysicalDeviceIdPropertiesKHR                         = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ID_PROPERTIES_KHR
	StructureTypePhysicalDeviceImageFormatInfo2                        = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_FORMAT_INFO_2
	StructureTypePhysicalDeviceImageFormatInfo2KHR                     = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_FORMAT_INFO_2_KHR
	StructureTypePhysicalDeviceInlineUniformBlockFeaturesEXT           = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_INLINE_UNIFORM_BLOCK_FEATURES_EXT
	StructureTypePhysicalDeviceInlineUniformBlockPropertiesEXT         = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_INLINE_UNIFORM_BLOCK_PROPERTIES_EXT
	StructureTypePhysicalDeviceMaintenance3Properties                  = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_3_PROPERTIES
	StructureTypePhysicalDeviceMaintenance3PropertiesKHR               = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_3_PROPERTIES_KHR
	StructureTypePhysicalDeviceMemoryProperties2                       = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MEMORY_PROPERTIES_2
	StructureTypePhysicalDeviceMemoryProperties2KHR                    = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MEMORY_PROPERTIES_2_KHR
	StructureTypePhysicalDeviceMeshShaderFeaturesNV                    = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MESH_SHADER_FEATURES_NV
	StructureTypePhysicalDeviceMeshShaderPropertiesNV                  = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MESH_SHADER_PROPERTIES_NV
	StructureTypePhysicalDeviceMultiviewFeatures                       = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_FEATURES
	StructureTypePhysicalDeviceMultiviewFeaturesKHR                    = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_FEATURES_KHR
	StructureTypePhysicalDeviceMultiviewPerViewAttributesPropertiesNVX = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_PER_VIEW_ATTRIBUTES_PROPERTIES_NVX
	StructureTypePhysicalDeviceMultiviewProperties                     = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_PROPERTIES
	StructureTypePhysicalDeviceMultiviewPropertiesKHR                  = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_PROPERTIES_KHR
	StructureTypePhysicalDevicePointClippingProperties                 = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_POINT_CLIPPING_PROPERTIES
	StructureTypePhysicalDevicePointClippingPropertiesKHR              = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_POINT_CLIPPING_PROPERTIES_KHR
	StructureTypePhysicalDeviceProperties2                             = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2
	StructureTypePhysicalDeviceProperties2KHR                          = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2_KHR
	StructureTypePhysicalDeviceProtectedMemoryFeatures                 = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROTECTED_MEMORY_FEATURES
	StructureTypePhysicalDeviceProtectedMemoryProperties               = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROTECTED_MEMORY_PROPERTIES
	StructureTypePhysicalDevicePushDescriptorPropertiesKHR             = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PUSH_DESCRIPTOR_PROPERTIES_KHR
	StructureTypePhysicalDeviceRaytracingPropertiesNVX                 = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RAYTRACING_PROPERTIES_NVX
	StructureTypePhysicalDeviceRepresentativeFragmentTestFeaturesNV    = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_REPRESENTATIVE_FRAGMENT_TEST_FEATURES_NV
	StructureTypePhysicalDeviceSamplerFilterMinmaxPropertiesEXT        = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_FILTER_MINMAX_PROPERTIES_EXT
	StructureTypePhysicalDeviceSamplerYCbCrConversionFeatures          = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_YCBCR_CONVERSION_FEATURES
	StructureTypePhysicalDeviceSamplerYCbCrConversionFeaturesKHR       = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_YCBCR_CONVERSION_FEATURES_KHR
	StructureTypePhysicalDeviceSampleLocationsPropertiesEXT            = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLE_LOCATIONS_PROPERTIES_EXT
	StructureTypePhysicalDeviceShaderCorePropertiesAMD                 = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_CORE_PROPERTIES_AMD
	StructureTypePhysicalDeviceShaderDrawParameterFeatures             = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_DRAW_PARAMETER_FEATURES
	StructureTypePhysicalDeviceShaderImageFootprintFeaturesNV          = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_IMAGE_FOOTPRINT_FEATURES_NV
	StructureTypePhysicalDeviceShadingRateImageFeaturesNV              = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADING_RATE_IMAGE_FEATURES_NV
	StructureTypePhysicalDeviceShadingRateImagePropertiesNV            = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADING_RATE_IMAGE_PROPERTIES_NV
	StructureTypePhysicalDeviceSparseImageFormatInfo2                  = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SPARSE_IMAGE_FORMAT_INFO_2
	StructureTypePhysicalDeviceSparseImageFormatInfo2KHR               = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SPARSE_IMAGE_FORMAT_INFO_2_KHR
	StructureTypePhysicalDeviceSurfaceInfo2KHR                         = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SURFACE_INFO_2_KHR
	StructureTypePhysicalDeviceVariablePointerFeatures                 = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VARIABLE_POINTER_FEATURES
	StructureTypePhysicalDeviceVariablePointerFeaturesKHR              = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VARIABLE_POINTER_FEATURES_KHR
	StructureTypePhysicalDeviceVertexAttributeDivisorFeaturesEXT       = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VERTEX_ATTRIBUTE_DIVISOR_FEATURES_EXT
	StructureTypePhysicalDeviceVertexAttributeDivisorPropertiesEXT     = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VERTEX_ATTRIBUTE_DIVISOR_PROPERTIES_EXT
	StructureTypePhysicalDeviceVulkanMemoryModelFeaturesKHR            = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_MEMORY_MODEL_FEATURES_KHR
	StructureTypePipelineColorBlendAdvancedStateCreateInfoEXT          = C.VK_STRUCTURE_TYPE_PIPELINE_COLOR_BLEND_ADVANCED_STATE_CREATE_INFO_EXT
	StructureTypePipelineCoverageModulationStateCreateInfoNV           = C.VK_STRUCTURE_TYPE_PIPELINE_COVERAGE_MODULATION_STATE_CREATE_INFO_NV
	StructureTypePipelineCoverageToColorStateCreateInfoNV              = C.VK_STRUCTURE_TYPE_PIPELINE_COVERAGE_TO_COLOR_STATE_CREATE_INFO_NV
	StructureTypePipelineDiscardRectangleStateCreateInfoEXT            = C.VK_STRUCTURE_TYPE_PIPELINE_DISCARD_RECTANGLE_STATE_CREATE_INFO_EXT
	StructureTypePipelineRasterizationConservativeStateCreateInfoEXT   = C.VK_STRUCTURE_TYPE_PIPELINE_RASTERIZATION_CONSERVATIVE_STATE_CREATE_INFO_EXT
	StructureTypePipelineRasterizationStateRasterizationOrderAMD       = C.VK_STRUCTURE_TYPE_PIPELINE_RASTERIZATION_STATE_RASTERIZATION_ORDER_AMD
	StructureTypePipelineRepresentativeFragmentTestStateCreateInfoNV   = C.VK_STRUCTURE_TYPE_PIPELINE_REPRESENTATIVE_FRAGMENT_TEST_STATE_CREATE_INFO_NV
	StructureTypePipelineSampleLocationsStateCreateInfoEXT             = C.VK_STRUCTURE_TYPE_PIPELINE_SAMPLE_LOCATIONS_STATE_CREATE_INFO_EXT
	StructureTypePipelineTessellationDomainOriginStateCreateInfo       = C.VK_STRUCTURE_TYPE_PIPELINE_TESSELLATION_DOMAIN_ORIGIN_STATE_CREATE_INFO
	StructureTypePipelineTessellationDomainOriginStateCreateInfoKHR    = C.VK_STRUCTURE_TYPE_PIPELINE_TESSELLATION_DOMAIN_ORIGIN_STATE_CREATE_INFO_KHR
	StructureTypePipelineVertexInputDivisorStateCreateInfoEXT          = C.VK_STRUCTURE_TYPE_PIPELINE_VERTEX_INPUT_DIVISOR_STATE_CREATE_INFO_EXT
	StructureTypePipelineViewportCoarseSampleOrderStateCreateInfoNV    = C.VK_STRUCTURE_TYPE_PIPELINE_VIEWPORT_COARSE_SAMPLE_ORDER_STATE_CREATE_INFO_NV
	StructureTypePipelineViewportExclusiveScissorStateCreateInfoNV     = C.VK_STRUCTURE_TYPE_PIPELINE_VIEWPORT_EXCLUSIVE_SCISSOR_STATE_CREATE_INFO_NV
	StructureTypePipelineViewportShadingRateImageStateCreateInfoNV     = C.VK_STRUCTURE_TYPE_PIPELINE_VIEWPORT_SHADING_RATE_IMAGE_STATE_CREATE_INFO_NV
	StructureTypePipelineViewportSwizzleStateCreateInfoNV              = C.VK_STRUCTURE_TYPE_PIPELINE_VIEWPORT_SWIZZLE_STATE_CREATE_INFO_NV
	StructureTypePipelineViewportWScalingStateCreateInfoNV             = C.VK_STRUCTURE_TYPE_PIPELINE_VIEWPORT_W_SCALING_STATE_CREATE_INFO_NV
	StructureTypePresentInfoKHR                                        = C.VK_STRUCTURE_TYPE_PRESENT_INFO_KHR
	StructureTypePresentRegionsKHR                                     = C.VK_STRUCTURE_TYPE_PRESENT_REGIONS_KHR
	StructureTypePresentTimesInfoGoogle                                = C.VK_STRUCTURE_TYPE_PRESENT_TIMES_INFO_GOOGLE
	StructureTypeProtectedSubmitInfo                                   = C.VK_STRUCTURE_TYPE_PROTECTED_SUBMIT_INFO
	StructureTypePhysicalDevice16ItStorageFeatures                     = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_16BIT_STORAGE_FEATURES
	StructureTypePhysicalDeviceGroupProperties                         = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES
	StructureTypePhysicalDeviceSubgroupProperties                      = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SUBGROUP_PROPERTIES
	StructureTypePipelineCacheCreateInfo                               = C.VK_STRUCTURE_TYPE_PIPELINE_CACHE_CREATE_INFO
	StructureTypePipelineColorBlendStateCreateInfo                     = C.VK_STRUCTURE_TYPE_PIPELINE_COLOR_BLEND_STATE_CREATE_INFO
	StructureTypePipelineDepthStencilStateCreateInfo                   = C.VK_STRUCTURE_TYPE_PIPELINE_DEPTH_STENCIL_STATE_CREATE_INFO
	StructureTypePipelineDynamicStateCreateInfo                        = C.VK_STRUCTURE_TYPE_PIPELINE_DYNAMIC_STATE_CREATE_INFO
	StructureTypePipelineInputAssemblyStateCreateInfo                  = C.VK_STRUCTURE_TYPE_PIPELINE_INPUT_ASSEMBLY_STATE_CREATE_INFO
	StructureTypePipelineLayoutCreateInfo                              = C.VK_STRUCTURE_TYPE_PIPELINE_LAYOUT_CREATE_INFO
	StructureTypePipelineMultisampleStateCreateInfo                    = C.VK_STRUCTURE_TYPE_PIPELINE_MULTISAMPLE_STATE_CREATE_INFO
	StructureTypePipelineRasterizationStateCreateInfo                  = C.VK_STRUCTURE_TYPE_PIPELINE_RASTERIZATION_STATE_CREATE_INFO
	StructureTypePipelineShaderStageCreateInfo                         = C.VK_STRUCTURE_TYPE_PIPELINE_SHADER_STAGE_CREATE_INFO
	StructureTypePipelineTessellationStateCreateInfo                   = C.VK_STRUCTURE_TYPE_PIPELINE_TESSELLATION_STATE_CREATE_INFO
	StructureTypePipelineVertexInputStateCreateInfo                    = C.VK_STRUCTURE_TYPE_PIPELINE_VERTEX_INPUT_STATE_CREATE_INFO
	StructureTypePipelineViewportStateCreateInfo                       = C.VK_STRUCTURE_TYPE_PIPELINE_VIEWPORT_STATE_CREATE_INFO
	StructureTypeQueueFamilyCheckpointPropertiesNV                     = C.VK_STRUCTURE_TYPE_QUEUE_FAMILY_CHECKPOINT_PROPERTIES_NV
	StructureTypeQueueFamilyProperties2                                = C.VK_STRUCTURE_TYPE_QUEUE_FAMILY_PROPERTIES_2
	StructureTypeQueueFamilyProperties2KHR                             = C.VK_STRUCTURE_TYPE_QUEUE_FAMILY_PROPERTIES_2_KHR
	StructureTypeQueryPoolCreateInfo                                   = C.VK_STRUCTURE_TYPE_QUERY_POOL_CREATE_INFO
	StructureTypeRaytracingPipelineCreateInfoNVX                       = C.VK_STRUCTURE_TYPE_RAYTRACING_PIPELINE_CREATE_INFO_NVX
	StructureTypeRenderPassCreateInfo2KHR                              = C.VK_STRUCTURE_TYPE_RENDER_PASS_CREATE_INFO_2_KHR
	StructureTypeRenderPassInputAttachmentAspectCreateInfo             = C.VK_STRUCTURE_TYPE_RENDER_PASS_INPUT_ATTACHMENT_ASPECT_CREATE_INFO
	StructureTypeRenderPassInputAttachmentAspectCreateInfoKHR          = C.VK_STRUCTURE_TYPE_RENDER_PASS_INPUT_ATTACHMENT_ASPECT_CREATE_INFO_KHR
	StructureTypeRenderPassMultiviewCreateInfo                         = C.VK_STRUCTURE_TYPE_RENDER_PASS_MULTIVIEW_CREATE_INFO
	StructureTypeRenderPassMultiviewCreateInfoKHR                      = C.VK_STRUCTURE_TYPE_RENDER_PASS_MULTIVIEW_CREATE_INFO_KHR
	StructureTypeRenderPassSampleLocationsBeginInfoEXT                 = C.VK_STRUCTURE_TYPE_RENDER_PASS_SAMPLE_LOCATIONS_BEGIN_INFO_EXT
	StructureTypeRenderPassBeginInfo                                   = C.VK_STRUCTURE_TYPE_RENDER_PASS_BEGIN_INFO
	StructureTypeRenderPassCreateInfo                                  = C.VK_STRUCTURE_TYPE_RENDER_PASS_CREATE_INFO
	StructureTypeSamplerReductionModeCreateInfoEXT                     = C.VK_STRUCTURE_TYPE_SAMPLER_REDUCTION_MODE_CREATE_INFO_EXT
	StructureTypeSamplerYCbCrConversionCreateInfo                      = C.VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_CREATE_INFO
	StructureTypeSamplerYCbCrConversionCreateInfoKHR                   = C.VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_CREATE_INFO_KHR
	StructureTypeSamplerYCbCrConversionImageFormatProperties           = C.VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_IMAGE_FORMAT_PROPERTIES
	StructureTypeSamplerYCbCrConversionImageFormatPropertiesKHR        = C.VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_IMAGE_FORMAT_PROPERTIES_KHR
	StructureTypeSamplerYCbCrConversionINFO                            = C.VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_INFO
	StructureTypeSamplerYCbCrConversionINFO_KHR                        = C.VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_INFO_KHR
	StructureTypeSampleLocationsInfoEXT                                = C.VK_STRUCTURE_TYPE_SAMPLE_LOCATIONS_INFO_EXT
	StructureTypeSemaphoreGetFdInfoKHR                                 = C.VK_STRUCTURE_TYPE_SEMAPHORE_GET_FD_INFO_KHR
	StructureTypeSemaphoreGetWin32HandleInfoKHR                        = C.VK_STRUCTURE_TYPE_SEMAPHORE_GET_WIN32_HANDLE_INFO_KHR
	StructureTypeShaderModuleValidationCacheCreateInfoEXT              = C.VK_STRUCTURE_TYPE_SHADER_MODULE_VALIDATION_CACHE_CREATE_INFO_EXT
	StructureTypeSharedPresentSurfaceCapabilitiesKHR                   = C.VK_STRUCTURE_TYPE_SHARED_PRESENT_SURFACE_CAPABILITIES_KHR
	StructureTypeSparseImageFormatProperties2                          = C.VK_STRUCTURE_TYPE_SPARSE_IMAGE_FORMAT_PROPERTIES_2
	StructureTypeSparseImageFormatProperties2KHR                       = C.VK_STRUCTURE_TYPE_SPARSE_IMAGE_FORMAT_PROPERTIES_2_KHR
	StructureTypeSparseImageMemoryRequirements2                        = C.VK_STRUCTURE_TYPE_SPARSE_IMAGE_MEMORY_REQUIREMENTS_2
	StructureTypeSparseImageMemoryRequirements2KHR                     = C.VK_STRUCTURE_TYPE_SPARSE_IMAGE_MEMORY_REQUIREMENTS_2_KHR
	StructureTypeSubpassBeginInfoKHR                                   = C.VK_STRUCTURE_TYPE_SUBPASS_BEGIN_INFO_KHR
	StructureTypeSubpassDependency2KHR                                 = C.VK_STRUCTURE_TYPE_SUBPASS_DEPENDENCY_2_KHR
	StructureTypeSubpassDescription2KHR                                = C.VK_STRUCTURE_TYPE_SUBPASS_DESCRIPTION_2_KHR
	StructureTypeSubpassEndInfoKHR                                     = C.VK_STRUCTURE_TYPE_SUBPASS_END_INFO_KHR
	StructureTypeSurfaceCapabilities2EXT                               = C.VK_STRUCTURE_TYPE_SURFACE_CAPABILITIES2_EXT
	StructureTypeSurfaceCapabilities2KHR                               = C.VK_STRUCTURE_TYPE_SURFACE_CAPABILITIES_2_KHR
	StructureTypeSurfaceFormat2KHR                                     = C.VK_STRUCTURE_TYPE_SURFACE_FORMAT_2_KHR
	StructureTypeSwapchainCounterCreateInfoEXT                         = C.VK_STRUCTURE_TYPE_SWAPCHAIN_COUNTER_CREATE_INFO_EXT
	StructureTypeSwapchainCreateInfoKHR                                = C.VK_STRUCTURE_TYPE_SWAPCHAIN_CREATE_INFO_KHR
	StructureTypeSamplerCreateInfo                                     = C.VK_STRUCTURE_TYPE_SAMPLER_CREATE_INFO
	StructureTypeSemaphoreCreateInfo                                   = C.VK_STRUCTURE_TYPE_SEMAPHORE_CREATE_INFO
	StructureTypeShaderModuleCreateInfo                                = C.VK_STRUCTURE_TYPE_SHADER_MODULE_CREATE_INFO
	StructureTypeSubmitInfo                                            = C.VK_STRUCTURE_TYPE_SUBMIT_INFO
	StructureTypeTextureLodGatherFormatPropertiesAMD                   = C.VK_STRUCTURE_TYPE_TEXTURE_LOD_GATHER_FORMAT_PROPERTIES_AMD
	StructureTypeValidationCacheCreateInfoEXT                          = C.VK_STRUCTURE_TYPE_VALIDATION_CACHE_CREATE_INFO_EXT
	StructureTypeValidationFlagsEXT                                    = C.VK_STRUCTURE_TYPE_VALIDATION_FLAGS_EXT
	StructureTypeViSurfaceCreateInfoNN                                 = C.VK_STRUCTURE_TYPE_VI_SURFACE_CREATE_INFO_NN
	StructureTypeWaylandSurfaceCreateInfoKHR                           = C.VK_STRUCTURE_TYPE_WAYLAND_SURFACE_CREATE_INFO_KHR
	StructureTypeWin32KeyedMutexAcquireReleaseInfoKHR                  = C.VK_STRUCTURE_TYPE_WIN32_KEYED_MUTEX_ACQUIRE_RELEASE_INFO_KHR
	StructureTypeWin32KeyedMutexAcquireReleaseInfoNV                   = C.VK_STRUCTURE_TYPE_WIN32_KEYED_MUTEX_ACQUIRE_RELEASE_INFO_NV
	StructureTypeWin32SurfaceCreateInfoKHR                             = C.VK_STRUCTURE_TYPE_WIN32_SURFACE_CREATE_INFO_KHR
	StructureTypeWriteDescriptorSetInlineUniformBlockEXT               = C.VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET_INLINE_UNIFORM_BLOCK_EXT
	StructureTypeWriteDescriptorSet                                    = C.VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET
	StructureTypeXcbSurfaceCreateInfoKHR                               = C.VK_STRUCTURE_TYPE_XCB_SURFACE_CREATE_INFO_KHR
	StructureTypeXlibSurfaceCreateInfoKHR                              = C.VK_STRUCTURE_TYPE_XLIB_SURFACE_CREATE_INFO_KHR
)

const (
	// SurfaceTransformIdentityBitKHR specifies that image content is presented without being transformed.
	SurfaceTransformIdentityBitKHR SurfaceTransformFlagsKHR = C.VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR

	// SurfaceTransformRotate90BitKHR  specifies that image content is rotated 90 degrees clockwise.
	SurfaceTransformRotate90BitKHR SurfaceTransformFlagsKHR = C.VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR

	// SurfaceTransformRotate180BitKHR specifies that image content is rotated 180 degrees clockwise.
	SurfaceTransformRotate180BitKHR SurfaceTransformFlagsKHR = C.VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR

	// SurfaceTransformRotate270BitKHR specifies that image content is rotated 270 degrees clockwise.
	SurfaceTransformRotate270BitKHR SurfaceTransformFlagsKHR = C.VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR

	// SurfaceTransformHorizontalMirrorBitKhr specifies that image content is mirrored horizontally.
	SurfaceTransformHorizontalMirrorBitKHR SurfaceTransformFlagsKHR = C.VK_SURFACE_TRANSFORM_HORIZONTAL_MIRROR_BIT_KHR

	// SurfaceTransformHorizontalMirrorRotate90BitKHR specifies that image content is mirrored horizontally, then rotated 90 degrees clockwise.
	SurfaceTransformHorizontalMirrorRotate90BitKHR SurfaceTransformFlagsKHR = C.VK_SURFACE_TRANSFORM_HORIZONTAL_MIRROR_ROTATE_90_BIT_KHR

	// SurfaceTransformHorizontalMirrorRotate180BitKHR specifies that image content is mirrored horizontally, then rotated 180 degrees clockwise.
	SurfaceTransformHorizontalMirrorRotate180BitKHR SurfaceTransformFlagsKHR = C.VK_SURFACE_TRANSFORM_HORIZONTAL_MIRROR_ROTATE_180_BIT_KHR

	// SurfaceTransformHorizontalMirrorRotate270BitKHR specifies that image content is mirrored horizontally, then rotated 270 degrees clockwise.
	SurfaceTransformHorizontalMirrorRotate270BitKHR SurfaceTransformFlagsKHR = C.VK_SURFACE_TRANSFORM_HORIZONTAL_MIRROR_ROTATE_270_BIT_KHR

	// SurfaceTransformInheritBitKHR specifies that the presentation transform is not specified,
	// and is instead determined by platform-specific considerations and mechanisms outside Vulkan.
	SurfaceTransformInheritBitKHR SurfaceTransformFlagsKHR = C.VK_SURFACE_TRANSFORM_INHERIT_BIT_KHR
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
	ImageUsageTransferSrcBit ImageUsageFlags = C.VK_IMAGE_USAGE_TRANSFER_SRC_BIT

	// ImageUsageTransferDstBit specifies that the image can be used as the destination of a transfer command.
	ImageUsageTransferDstBit ImageUsageFlags = C.VK_IMAGE_USAGE_TRANSFER_DST_BIT

	// ImageUsageSampledBit specifies that the image can be used to create a VkImageView
	// suitable for occupying a VkDescriptorSet slot either of type VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE
	// or VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER, and be sampled by a shader.
	ImageUsageSampledBit ImageUsageFlags = C.VK_IMAGE_USAGE_SAMPLED_BIT

	// ImageUsageStorageBit specifies that the image can be used to create a VkImageView
	// suitable for occupying a VkDescriptorSet slot of type VK_DESCRIPTOR_TYPE_STORAGE_IMAGE.
	ImageUsageStorageBit ImageUsageFlags = C.VK_IMAGE_USAGE_STORAGE_BIT

	// ImageUsageColorAttachmentBit specifies that the image can be used to create a VkImageView
	// suitable for use as a color or resolve attachment in a VkFramebuffer.
	ImageUsageColorAttachmentBit ImageUsageFlags = C.VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT

	// ImageUsageDepthStencilAttachmentBit specifies that the image can be used to create a VkImageView
	// suitable for use as a depth/stencil attachment in a VkFramebuffer.
	ImageUsageDepthStencilAttachmentBit ImageUsageFlags = C.VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT

	// ImageUsageTransientAttachmentBit specifies that the memory bound to this image will have been allocated
	// with the VK_MEMORY_PROPERTY_LAZILY_ALLOCATED_BIT (see Memory Allocation for more detail).
	// This bit can be set for any image that can be used to create a VkImageView
	// suitable for use as a color, resolve, depth/stencil, or input attachment.
	ImageUsageTransientAttachmentBit ImageUsageFlags = C.VK_IMAGE_USAGE_TRANSIENT_ATTACHMENT_BIT

	// ImageUsageInputAttachmentBit specifies that the image can be used to create a VkImageView suitable
	// for occupying VkDescriptorSet slot of type VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT;
	// be read from a shader as an input attachment; and be used as an input attachment in a framebuffer.
	ImageUsageInputAttachmentBit ImageUsageFlags = C.VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT
)

const (
	FormatUNDEFINED                  Format = C.VK_FORMAT_UNDEFINED
	FormatR4G4_UNORM_PACK8           Format = C.VK_FORMAT_R4G4_UNORM_PACK8
	FormatR4G4B4A4_UNORM_PACK16      Format = C.VK_FORMAT_R4G4B4A4_UNORM_PACK16
	FormatB4G4R4A4_UNORM_PACK16      Format = C.VK_FORMAT_B4G4R4A4_UNORM_PACK16
	FormatR5G6B5_UNORM_PACK16        Format = C.VK_FORMAT_R5G6B5_UNORM_PACK16
	FormatB5G6R5_UNORM_PACK16        Format = C.VK_FORMAT_B5G6R5_UNORM_PACK16
	FormatR5G5B5A1_UNORM_PACK16      Format = C.VK_FORMAT_R5G5B5A1_UNORM_PACK16
	FormatB5G5R5A1_UNORM_PACK16      Format = C.VK_FORMAT_B5G5R5A1_UNORM_PACK16
	FormatA1R5G5B5_UNORM_PACK16      Format = C.VK_FORMAT_A1R5G5B5_UNORM_PACK16
	FormatR8_UNORM                   Format = C.VK_FORMAT_R8_UNORM
	FormatR8_SNORM                   Format = C.VK_FORMAT_R8_SNORM
	FormatR8_USCALED                 Format = C.VK_FORMAT_R8_USCALED
	FormatR8_SSCALED                 Format = C.VK_FORMAT_R8_SSCALED
	FormatR8_UINT                    Format = C.VK_FORMAT_R8_UINT
	FormatR8_SINT                    Format = C.VK_FORMAT_R8_SINT
	FormatR8_SRGB                    Format = C.VK_FORMAT_R8_SRGB
	FormatR8G8_UNORM                 Format = C.VK_FORMAT_R8G8_UNORM
	FormatR8G8_SNORM                 Format = C.VK_FORMAT_R8G8_SNORM
	FormatR8G8_USCALED               Format = C.VK_FORMAT_R8G8_USCALED
	FormatR8G8_SSCALED               Format = C.VK_FORMAT_R8G8_SSCALED
	FormatR8G8_UINT                  Format = C.VK_FORMAT_R8G8_UINT
	FormatR8G8_SINT                  Format = C.VK_FORMAT_R8G8_SINT
	FormatR8G8_SRGB                  Format = C.VK_FORMAT_R8G8_SRGB
	FormatR8G8B8_UNORM               Format = C.VK_FORMAT_R8G8B8_UNORM
	FormatR8G8B8_SNORM               Format = C.VK_FORMAT_R8G8B8_SNORM
	FormatR8G8B8_USCALED             Format = C.VK_FORMAT_R8G8B8_USCALED
	FormatR8G8B8_SSCALED             Format = C.VK_FORMAT_R8G8B8_SSCALED
	FormatR8G8B8_UINT                Format = C.VK_FORMAT_R8G8B8_UINT
	FormatR8G8B8_SINT                Format = C.VK_FORMAT_R8G8B8_SINT
	FormatR8G8B8_SRGB                Format = C.VK_FORMAT_R8G8B8_SRGB
	FormatB8G8R8_UNORM               Format = C.VK_FORMAT_B8G8R8_UNORM
	FormatB8G8R8_SNORM               Format = C.VK_FORMAT_B8G8R8_SNORM
	FormatB8G8R8_USCALED             Format = C.VK_FORMAT_B8G8R8_USCALED
	FormatB8G8R8_SSCALED             Format = C.VK_FORMAT_B8G8R8_SSCALED
	FormatB8G8R8_UINT                Format = C.VK_FORMAT_B8G8R8_UINT
	FormatB8G8R8_SINT                Format = C.VK_FORMAT_B8G8R8_SINT
	FormatB8G8R8_SRGB                Format = C.VK_FORMAT_B8G8R8_SRGB
	FormatR8G8B8A8_UNORM             Format = C.VK_FORMAT_R8G8B8A8_UNORM
	FormatR8G8B8A8_SNORM             Format = C.VK_FORMAT_R8G8B8A8_SNORM
	FormatR8G8B8A8_USCALED           Format = C.VK_FORMAT_R8G8B8A8_USCALED
	FormatR8G8B8A8_SSCALED           Format = C.VK_FORMAT_R8G8B8A8_SSCALED
	FormatR8G8B8A8_UINT              Format = C.VK_FORMAT_R8G8B8A8_UINT
	FormatR8G8B8A8_SINT              Format = C.VK_FORMAT_R8G8B8A8_SINT
	FormatR8G8B8A8_SRGB              Format = C.VK_FORMAT_R8G8B8A8_SRGB
	FormatB8G8R8A8_UNORM             Format = C.VK_FORMAT_B8G8R8A8_UNORM
	FormatB8G8R8A8_SNORM             Format = C.VK_FORMAT_B8G8R8A8_SNORM
	FormatB8G8R8A8_USCALED           Format = C.VK_FORMAT_B8G8R8A8_USCALED
	FormatB8G8R8A8_SSCALED           Format = C.VK_FORMAT_B8G8R8A8_SSCALED
	FormatB8G8R8A8_UINT              Format = C.VK_FORMAT_B8G8R8A8_UINT
	FormatB8G8R8A8_SINT              Format = C.VK_FORMAT_B8G8R8A8_SINT
	FormatB8G8R8A8_SRGB              Format = C.VK_FORMAT_B8G8R8A8_SRGB
	FormatA8B8G8R8_UNORM_PACK32      Format = C.VK_FORMAT_A8B8G8R8_UNORM_PACK32
	FormatA8B8G8R8_SNORM_PACK32      Format = C.VK_FORMAT_A8B8G8R8_SNORM_PACK32
	FormatA8B8G8R8_USCALED_PACK32    Format = C.VK_FORMAT_A8B8G8R8_USCALED_PACK32
	FormatA8B8G8R8_SSCALED_PACK32    Format = C.VK_FORMAT_A8B8G8R8_SSCALED_PACK32
	FormatA8B8G8R8_UINT_PACK32       Format = C.VK_FORMAT_A8B8G8R8_UINT_PACK32
	FormatA8B8G8R8_SINT_PACK32       Format = C.VK_FORMAT_A8B8G8R8_SINT_PACK32
	FormatA8B8G8R8_SRGB_PACK32       Format = C.VK_FORMAT_A8B8G8R8_SRGB_PACK32
	FormatA2R10G10B10_UNORM_PACK32   Format = C.VK_FORMAT_A2R10G10B10_UNORM_PACK32
	FormatA2R10G10B10_SNORM_PACK32   Format = C.VK_FORMAT_A2R10G10B10_SNORM_PACK32
	FormatA2R10G10B10_USCALED_PACK32 Format = C.VK_FORMAT_A2R10G10B10_USCALED_PACK32
	FormatA2R10G10B10_SSCALED_PACK32 Format = C.VK_FORMAT_A2R10G10B10_SSCALED_PACK32
	FormatA2R10G10B10_UINT_PACK32    Format = C.VK_FORMAT_A2R10G10B10_UINT_PACK32
	FormatA2R10G10B10_SINT_PACK32    Format = C.VK_FORMAT_A2R10G10B10_SINT_PACK32
	FormatA2B10G10R10_UNORM_PACK32   Format = C.VK_FORMAT_A2B10G10R10_UNORM_PACK32
	FormatA2B10G10R10_SNORM_PACK32   Format = C.VK_FORMAT_A2B10G10R10_SNORM_PACK32
	FormatA2B10G10R10_USCALED_PACK32 Format = C.VK_FORMAT_A2B10G10R10_USCALED_PACK32
	FormatA2B10G10R10_SSCALED_PACK32 Format = C.VK_FORMAT_A2B10G10R10_SSCALED_PACK32
	FormatA2B10G10R10_UINT_PACK32    Format = C.VK_FORMAT_A2B10G10R10_UINT_PACK32
	FormatA2B10G10R10_SINT_PACK32    Format = C.VK_FORMAT_A2B10G10R10_SINT_PACK32
	FormatR16_UNORM                  Format = C.VK_FORMAT_R16_UNORM
	FormatR16_SNORM                  Format = C.VK_FORMAT_R16_SNORM
	FormatR16_USCALED                Format = C.VK_FORMAT_R16_USCALED
	FormatR16_SSCALED                Format = C.VK_FORMAT_R16_SSCALED
	FormatR16_UINT                   Format = C.VK_FORMAT_R16_UINT
	FormatR16_SINT                   Format = C.VK_FORMAT_R16_SINT
	FormatR16_SFLOAT                 Format = C.VK_FORMAT_R16_SFLOAT
	FormatR16G16_UNORM               Format = C.VK_FORMAT_R16G16_UNORM
	FormatR16G16_SNORM               Format = C.VK_FORMAT_R16G16_SNORM
	FormatR16G16_USCALED             Format = C.VK_FORMAT_R16G16_USCALED
	FormatR16G16_SSCALED             Format = C.VK_FORMAT_R16G16_SSCALED
	FormatR16G16_UINT                Format = C.VK_FORMAT_R16G16_UINT
	FormatR16G16_SINT                Format = C.VK_FORMAT_R16G16_SINT
	FormatR16G16_SFLOAT              Format = C.VK_FORMAT_R16G16_SFLOAT
	FormatR16G16B16_UNORM            Format = C.VK_FORMAT_R16G16B16_UNORM
	FormatR16G16B16_SNORM            Format = C.VK_FORMAT_R16G16B16_SNORM
	FormatR16G16B16_USCALED          Format = C.VK_FORMAT_R16G16B16_USCALED
	FormatR16G16B16_SSCALED          Format = C.VK_FORMAT_R16G16B16_SSCALED
	FormatR16G16B16_UINT             Format = C.VK_FORMAT_R16G16B16_UINT
	FormatR16G16B16_SINT             Format = C.VK_FORMAT_R16G16B16_SINT
	FormatR16G16B16_SFLOAT           Format = C.VK_FORMAT_R16G16B16_SFLOAT
	FormatR16G16B16A16_UNORM         Format = C.VK_FORMAT_R16G16B16A16_UNORM
	FormatR16G16B16A16_SNORM         Format = C.VK_FORMAT_R16G16B16A16_SNORM
	FormatR16G16B16A16_USCALED       Format = C.VK_FORMAT_R16G16B16A16_USCALED
	FormatR16G16B16A16_SSCALED       Format = C.VK_FORMAT_R16G16B16A16_SSCALED
	FormatR16G16B16A16_UINT          Format = C.VK_FORMAT_R16G16B16A16_UINT
	FormatR16G16B16A16_SINT          Format = C.VK_FORMAT_R16G16B16A16_SINT
	FormatR16G16B16A16_SFLOAT        Format = C.VK_FORMAT_R16G16B16A16_SFLOAT
	FormatR32_UINT                   Format = C.VK_FORMAT_R32_UINT
	FormatR32_SINT                   Format = C.VK_FORMAT_R32_SINT
	FormatR32_SFLOAT                 Format = C.VK_FORMAT_R32_SFLOAT
	FormatR32G32_UINT                Format = C.VK_FORMAT_R32G32_UINT
	FormatR32G32_SINT                Format = C.VK_FORMAT_R32G32_SINT
	FormatR32G32_SFLOAT              Format = C.VK_FORMAT_R32G32_SFLOAT
	FormatR32G32B32_UINT             Format = C.VK_FORMAT_R32G32B32_UINT
	FormatR32G32B32_SINT             Format = C.VK_FORMAT_R32G32B32_SINT
	FormatR32G32B32_SFLOAT           Format = C.VK_FORMAT_R32G32B32_SFLOAT
	FormatR32G32B32A32_UINT          Format = C.VK_FORMAT_R32G32B32A32_UINT
	FormatR32G32B32A32_SINT          Format = C.VK_FORMAT_R32G32B32A32_SINT
	FormatR32G32B32A32_SFLOAT        Format = C.VK_FORMAT_R32G32B32A32_SFLOAT
	FormatR64_UINT                   Format = C.VK_FORMAT_R64_UINT
	FormatR64_SINT                   Format = C.VK_FORMAT_R64_SINT
	FormatR64_SFLOAT                 Format = C.VK_FORMAT_R64_SFLOAT
	FormatR64G64_UINT                Format = C.VK_FORMAT_R64G64_UINT
	FormatR64G64_SINT                Format = C.VK_FORMAT_R64G64_SINT
	FormatR64G64_SFLOAT              Format = C.VK_FORMAT_R64G64_SFLOAT
	FormatR64G64B64_UINT             Format = C.VK_FORMAT_R64G64B64_UINT
	FormatR64G64B64_SINT             Format = C.VK_FORMAT_R64G64B64_SINT
	FormatR64G64B64_SFLOAT           Format = C.VK_FORMAT_R64G64B64_SFLOAT
	FormatR64G64B64A64_UINT          Format = C.VK_FORMAT_R64G64B64A64_UINT
	FormatR64G64B64A64_SINT          Format = C.VK_FORMAT_R64G64B64A64_SINT
	FormatR64G64B64A64_SFLOAT        Format = C.VK_FORMAT_R64G64B64A64_SFLOAT
	FormatB10G11R11_UFLOAT_PACK32    Format = C.VK_FORMAT_B10G11R11_UFLOAT_PACK32
	FormatE5B9G9R9_UFLOAT_PACK32     Format = C.VK_FORMAT_E5B9G9R9_UFLOAT_PACK32
	FormatD16_UNORM                  Format = C.VK_FORMAT_D16_UNORM
	FormatX8_D24_UNORM_PACK32        Format = C.VK_FORMAT_X8_D24_UNORM_PACK32
	FormatD32_SFLOAT                 Format = C.VK_FORMAT_D32_SFLOAT
	FormatS8_UINT                    Format = C.VK_FORMAT_S8_UINT
	FormatD16_UNORM_S8_UINT          Format = C.VK_FORMAT_D16_UNORM_S8_UINT
	FormatD24_UNORM_S8_UINT          Format = C.VK_FORMAT_D24_UNORM_S8_UINT
	FormatD32_SFLOAT_S8_UINT         Format = C.VK_FORMAT_D32_SFLOAT_S8_UINT
	FormatBC1_RGB_UNORM_BLOCK        Format = C.VK_FORMAT_BC1_RGB_UNORM_BLOCK
	FormatBC1_RGB_SRGB_BLOCK         Format = C.VK_FORMAT_BC1_RGB_SRGB_BLOCK
	FormatBC1_RGBA_UNORM_BLOCK       Format = C.VK_FORMAT_BC1_RGBA_UNORM_BLOCK
	FormatBC1_RGBA_SRGB_BLOCK        Format = C.VK_FORMAT_BC1_RGBA_SRGB_BLOCK
	FormatBC2_UNORM_BLOCK            Format = C.VK_FORMAT_BC2_UNORM_BLOCK
	FormatBC2_SRGB_BLOCK             Format = C.VK_FORMAT_BC2_SRGB_BLOCK
	FormatBC3_UNORM_BLOCK            Format = C.VK_FORMAT_BC3_UNORM_BLOCK
	FormatBC3_SRGB_BLOCK             Format = C.VK_FORMAT_BC3_SRGB_BLOCK
	FormatBC4_UNORM_BLOCK            Format = C.VK_FORMAT_BC4_UNORM_BLOCK
	FormatBC4_SNORM_BLOCK            Format = C.VK_FORMAT_BC4_SNORM_BLOCK
	FormatBC5_UNORM_BLOCK            Format = C.VK_FORMAT_BC5_UNORM_BLOCK
	FormatBC5_SNORM_BLOCK            Format = C.VK_FORMAT_BC5_SNORM_BLOCK
	FormatBC6H_UFLOAT_BLOCK          Format = C.VK_FORMAT_BC6H_UFLOAT_BLOCK
	FormatBC6H_SFLOAT_BLOCK          Format = C.VK_FORMAT_BC6H_SFLOAT_BLOCK
	FormatBC7_UNORM_BLOCK            Format = C.VK_FORMAT_BC7_UNORM_BLOCK
	FormatBC7_SRGB_BLOCK             Format = C.VK_FORMAT_BC7_SRGB_BLOCK
	FormatETC2_R8G8B8_UNORM_BLOCK    Format = C.VK_FORMAT_ETC2_R8G8B8_UNORM_BLOCK
	FormatETC2_R8G8B8_SRGB_BLOCK     Format = C.VK_FORMAT_ETC2_R8G8B8_SRGB_BLOCK
	FormatETC2_R8G8B8A1_UNORM_BLOCK  Format = C.VK_FORMAT_ETC2_R8G8B8A1_UNORM_BLOCK
	FormatETC2_R8G8B8A1_SRGB_BLOCK   Format = C.VK_FORMAT_ETC2_R8G8B8A1_SRGB_BLOCK
	FormatETC2_R8G8B8A8_UNORM_BLOCK  Format = C.VK_FORMAT_ETC2_R8G8B8A8_UNORM_BLOCK
	FormatETC2_R8G8B8A8_SRGB_BLOCK   Format = C.VK_FORMAT_ETC2_R8G8B8A8_SRGB_BLOCK
	FormatEAC_R11_UNORM_BLOCK        Format = C.VK_FORMAT_EAC_R11_UNORM_BLOCK
	FormatEAC_R11_SNORM_BLOCK        Format = C.VK_FORMAT_EAC_R11_SNORM_BLOCK
	FormatEAC_R11G11_UNORM_BLOCK     Format = C.VK_FORMAT_EAC_R11G11_UNORM_BLOCK
	FormatEAC_R11G11_SNORM_BLOCK     Format = C.VK_FORMAT_EAC_R11G11_SNORM_BLOCK
	FormatASTC_4x4_UNORM_BLOCK       Format = C.VK_FORMAT_ASTC_4x4_UNORM_BLOCK
	FormatASTC_4x4_SRGB_BLOCK        Format = C.VK_FORMAT_ASTC_4x4_SRGB_BLOCK
	FormatASTC_5x4_UNORM_BLOCK       Format = C.VK_FORMAT_ASTC_5x4_UNORM_BLOCK
	FormatASTC_5x4_SRGB_BLOCK        Format = C.VK_FORMAT_ASTC_5x4_SRGB_BLOCK
	FormatASTC_5x5_UNORM_BLOCK       Format = C.VK_FORMAT_ASTC_5x5_UNORM_BLOCK
	FormatASTC_5x5_SRGB_BLOCK        Format = C.VK_FORMAT_ASTC_5x5_SRGB_BLOCK
	FormatASTC_6x5_UNORM_BLOCK       Format = C.VK_FORMAT_ASTC_6x5_UNORM_BLOCK
	FormatASTC_6x5_SRGB_BLOCK        Format = C.VK_FORMAT_ASTC_6x5_SRGB_BLOCK
	FormatASTC_6x6_UNORM_BLOCK       Format = C.VK_FORMAT_ASTC_6x6_UNORM_BLOCK
	FormatASTC_6x6_SRGB_BLOCK        Format = C.VK_FORMAT_ASTC_6x6_SRGB_BLOCK
	FormatASTC_8x5_UNORM_BLOCK       Format = C.VK_FORMAT_ASTC_8x5_UNORM_BLOCK
	FormatASTC_8x5_SRGB_BLOCK        Format = C.VK_FORMAT_ASTC_8x5_SRGB_BLOCK
	FormatASTC_8x6_UNORM_BLOCK       Format = C.VK_FORMAT_ASTC_8x6_UNORM_BLOCK
	FormatASTC_8x6_SRGB_BLOCK        Format = C.VK_FORMAT_ASTC_8x6_SRGB_BLOCK
	FormatASTC_8x8_UNORM_BLOCK       Format = C.VK_FORMAT_ASTC_8x8_UNORM_BLOCK
	FormatASTC_8x8_SRGB_BLOCK        Format = C.VK_FORMAT_ASTC_8x8_SRGB_BLOCK
	FormatASTC_10x5_UNORM_BLOCK      Format = C.VK_FORMAT_ASTC_10x5_UNORM_BLOCK
	FormatASTC_10x5_SRGB_BLOCK       Format = C.VK_FORMAT_ASTC_10x5_SRGB_BLOCK
	FormatASTC_10x6_UNORM_BLOCK      Format = C.VK_FORMAT_ASTC_10x6_UNORM_BLOCK
	FormatASTC_10x6_SRGB_BLOCK       Format = C.VK_FORMAT_ASTC_10x6_SRGB_BLOCK
	FormatASTC_10x8_UNORM_BLOCK      Format = C.VK_FORMAT_ASTC_10x8_UNORM_BLOCK
	FormatASTC_10x8_SRGB_BLOCK       Format = C.VK_FORMAT_ASTC_10x8_SRGB_BLOCK
	FormatASTC_10x10_UNORM_BLOCK     Format = C.VK_FORMAT_ASTC_10x10_UNORM_BLOCK
	FormatASTC_10x10_SRGB_BLOCK      Format = C.VK_FORMAT_ASTC_10x10_SRGB_BLOCK
	FormatASTC_12x10_UNORM_BLOCK     Format = C.VK_FORMAT_ASTC_12x10_UNORM_BLOCK
	FormatASTC_12x10_SRGB_BLOCK      Format = C.VK_FORMAT_ASTC_12x10_SRGB_BLOCK
	FormatASTC_12x12_UNORM_BLOCK     Format = C.VK_FORMAT_ASTC_12x12_UNORM_BLOCK
	FormatASTC_12x12_SRGB_BLOCK      Format = C.VK_FORMAT_ASTC_12x12_SRGB_BLOCK
)

const (
	COLOR_SPACE_SRGB_NONLINEAR_KHR ColorSpaceKHR = C.VK_COLOR_SPACE_SRGB_NONLINEAR_KHR
)

const (
	// VK_PRESENT_MODE_IMMEDIATE_KHR specifies that the presentation engine does not wait for a vertical blanking period to update the current image,
	// meaning this mode may result in visible tearing.
	// No internal queuing of presentation requests is needed, as the requests are applied immediately.
	PresentModeImmediateKHR PresentModeKHR = C.VK_PRESENT_MODE_IMMEDIATE_KHR

	// VK_PRESENT_MODE_MAILBOX_KHR specifies that the presentation engine waits for the next vertical blanking period to update the current image.
	// Tearing cannot be observed. An internal single-entry queue is used to hold pending presentation requests.
	// If the queue is full when a new presentation request is received, the new request replaces the existing entry,
	// and any images associated with the prior entry become available for re-use by the application.
	// One request is removed from the queue and processed during each vertical blanking period in which the queue is non-empty.
	PresentModeMailboxKHR PresentModeKHR = C.VK_PRESENT_MODE_MAILBOX_KHR

	// VK_PRESENT_MODE_FIFO_KHR specifies that the presentation engine waits for the next vertical blanking period to update the current image.
	// Tearing cannot be observed. An internal queue is used to hold pending presentation requests.
	// New requests are appended to the end of the queue,
	// and one request is removed from the beginning of the queue
	// and processed during each vertical blanking period in which the queue is non-empty.
	//
	// This is the only value of presentMode that is required to be supported.
	PresentModeFifoKHR PresentModeKHR = C.VK_PRESENT_MODE_FIFO_KHR

	// VK_PRESENT_MODE_FIFO_RELAXED_KHR specifies that the presentation engine generally waits for the next vertical blanking period to update the current image.
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
	PresentModeFifoRelaxedKHR PresentModeKHR = C.VK_PRESENT_MODE_FIFO_RELAXED_KHR

	// VK_PRESENT_MODE_SHARED_DEMAND_REFRESH_KHR specifies that the presentation engine and application have concurrent access to a single image,
	// which is referred to as a shared presentable image.
	// The presentation engine is only required to update the current image after a new presentation request is received.
	// Therefore the application must make a presentation request whenever an update is required.
	// However, the presentation engine may update the current image at any point, meaning this mode may result in visible tearing.
	PresentModeSharedDemandRefreshKHR PresentModeKHR = C.VK_PRESENT_MODE_SHARED_DEMAND_REFRESH_KHR

	// VK_PRESENT_MODE_SHARED_CONTINUOUS_REFRESH_KHR specifies that the presentation engine and application have concurrent access to a single image,
	// which is referred to as a shared presentable image.
	// The presentation engine periodically updates the current image on its regular refresh cycle.
	// The application is only required to make one initial presentation request,
	// after which the presentation engine must update the current image without any need for further presentation requests.
	// The application can indicate the image contents have been updated by making a presentation request,
	// but this does not guarantee the timing of when it will be updated.
	// This mode may result in visible tearing if rendering to the image is not timed correctly.
	PresentModeSharedContinuousRefreshKHR PresentModeKHR = C.VK_PRESENT_MODE_SHARED_CONTINUOUS_REFRESH_KHR
)

func (flags DeviceQueueCreateFlags) String() string {
	var props []string
	if (flags & DEVICE_QUEUE_CREATE_PROTECTED_BIT) != 0 {
		props = append(props, "VK_DEVICE_QUEUE_CREATE_PROTECTED_BIT")
	}
	return strings.Join(props, " | ")
}

func (flags QueueFlags) String() string {
	var props []string
	if (flags & QueueGraphicsBit) != 0 {
		props = append(props, "VK_QUEUE_GRAPHICS_BIT")
	}
	if (flags & QueueComputeBit) != 0 {
		props = append(props, "VK_QUEUE_COMPUTE_BIT")
	}
	if (flags & QueueTransferBit) != 0 {
		props = append(props, "VK_QUEUE_TRANSFER_BIT")
	}
	if (flags & QueueSparseBindingBit) != 0 {
		props = append(props, "VK_QUEUE_SPARSE_BINDING_BIT")
	}
	if (flags & QueueProtectedBit) != 0 {
		props = append(props, "VK_QUEUE_PROTECTED_BIT")
	}
	return strings.Join(props, " | ")
}

func (typ PhysicalDeviceType) String() string {
	switch typ {
	case PhysicalDeviceTypeOther:
		return "VK_PHYSICAL_DEVICE_TYPE_OTHER"
	case PhysicalDeviceTypeIntegratedGPU:
		return "VK_PHYSICAL_DEVICE_TYPE_INTEGRATED_GPU"
	case PhysicalDeviceTypeDiscreteGPU:
		return "VK_PHYSICAL_DEVICE_TYPE_DISCRETE_GPU"
	case PhysicalDeviceTypeVirtualGPU:
		return "VK_PHYSICAL_DEVICE_TYPE_VIRTUAL_GPU"
	case PhysicalDeviceTypeCPU:
		return "VK_PHYSICAL_DEVICE_TYPE_CPU"
	default:
		return strconv.Itoa(int(typ))
	}
}

func (res Result) Error() string {
	switch res {
	case Success:
		return "VK_SUCCESS"
	case NotReady:
		return "VK_NOT_READY"
	case Timeout:
		return "VK_TIMEOUT"
	case EventSet:
		return "VK_EVENT_SET"
	case EventReset:
		return "VK_EVENT_RESET"
	case Incomplete:
		return "VK_INCOMPLETE"
	case SuboptimalKHR:
		return "VK_SUBOPTIMAL_KHR"
	case ErrOutOfHostMemory:
		return "VK_ERROR_OUT_OF_HOST_MEMORY"
	case ErrOutOfDeviceMemory:
		return "VK_ERROR_OUT_OF_DEVICE_MEMORY"
	case ErrInitializationFailed:
		return "VK_ERROR_INITIALIZATION_FAILED"
	case ErrDeviceLost:
		return "VK_ERROR_DEVICE_LOST"
	case ErrMemoryMapFailed:
		return "VK_ERROR_MEMORY_MAP_FAILED"
	case ErrLayerNotPresent:
		return "VK_ERROR_LAYER_NOT_PRESENT"
	case ErrExtensionNotPresent:
		return "VK_ERROR_EXTENSION_NOT_PRESENT"
	case ErrFeatureNotPresent:
		return "VK_ERROR_FEATURE_NOT_PRESENT"
	case ErrIncompatibleDriver:
		return "VK_ERROR_INCOMPATIBLE_DRIVER"
	case ErrTooManyObjects:
		return "VK_ERROR_TOO_MANY_OBJECTS"
	case ErrFormatNotSupported:
		return "VK_ERROR_FORMAT_NOT_SUPPORTED"
	case ErrFragmentedPool:
		return "VK_ERROR_FRAGMENTED_POOL"
	case ErrOutOfPoolMemory:
		return "VK_ERROR_OUT_OF_POOL_MEMORY"
	case ErrInvalidExternalHandle:
		return "VK_ERROR_INVALID_EXTERNAL_HANDLE"
	case ErrSurfaceLostKHR:
		return "VK_ERROR_SURFACE_LOST_KHR"
	case ErrNativeWindowInUseKHR:
		return "VK_ERROR_NATIVE_WINDOW_IN_USE_KHR"
	case ErrOutOfDateKHR:
		return "VK_ERROR_OUT_OF_DATE_KHR"
	case ErrIncompatibleDisplayKHR:
		return "VK_ERROR_INCOMPATIBLE_DISPLAY_KHR"
	case ErrValidationFailedEXT:
		return "VK_ERROR_VALIDATION_FAILED_EXT"
	case ErrInvalidShaderNV:
		return "VK_ERROR_INVALID_SHADER_NV"
	case ErrFragmentationEXT:
		return "VK_ERROR_FRAGMENTATION_EXT"
	case ErrNotPermittedEXT:
		return "VK_ERROR_NOT_PERMITTED_EXT"
	default:
		return strconv.Itoa(int(res))
	}
}

func (flags SurfaceTransformFlagsKHR) String() string {
	var out []string
	if (flags & SurfaceTransformIdentityBitKHR) != 0 {
		out = append(out, "VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR")
	}
	if (flags & SurfaceTransformRotate90BitKHR) != 0 {
		out = append(out, "VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR")
	}
	if (flags & SurfaceTransformRotate180BitKHR) != 0 {
		out = append(out, "VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR")
	}
	if (flags & SurfaceTransformRotate270BitKHR) != 0 {
		out = append(out, "VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR")
	}
	if (flags & SurfaceTransformHorizontalMirrorBitKHR) != 0 {
		out = append(out, "VK_SURFACE_TRANSFORM_HORIZONTAL_MIRROR_BIT_KHR")
	}
	if (flags & SurfaceTransformHorizontalMirrorRotate90BitKHR) != 0 {
		out = append(out, "VK_SURFACE_TRANSFORM_HORIZONTAL_MIRROR_ROTATE_90_BIT_KHR")
	}
	if (flags & SurfaceTransformHorizontalMirrorRotate180BitKHR) != 0 {
		out = append(out, "VK_SURFACE_TRANSFORM_HORIZONTAL_MIRROR_ROTATE_180_BIT_KHR")
	}
	if (flags & SurfaceTransformHorizontalMirrorRotate270BitKHR) != 0 {
		out = append(out, "VK_SURFACE_TRANSFORM_HORIZONTAL_MIRROR_ROTATE_270_BIT_KHR")
	}
	if (flags & SurfaceTransformInheritBitKHR) != 0 {
		out = append(out, "VK_SURFACE_TRANSFORM_INHERIT_BIT_KHR")
	}
	return strings.Join(out, " | ")
}

func (flags CompositeAlphaFlagsKHR) String() string {
	var out []string
	if (flags & CompositeAlphaOpaqueBitKHR) != 0 {
		out = append(out, "VK_COMPOSITE_ALPHA_OPAQUE_BIT_KHR")
	}
	if (flags & CompositeAlphaPreMultipliedBitKHR) != 0 {
		out = append(out, "VK_COMPOSITE_ALPHA_PRE_MULTIPLIED_BIT_KHR")
	}
	if (flags & CompositeAlphaPostMultipliedBitKHR) != 0 {
		out = append(out, "VK_COMPOSITE_ALPHA_POST_MULTIPLIED_BIT_KHR")
	}
	if (flags & CompositeAlphaInheritBitKHR) != 0 {
		out = append(out, "VK_COMPOSITE_ALPHA_INHERIT_BIT_KHR")
	}
	return strings.Join(out, " | ")
}

func (flags ImageUsageFlags) String() string {
	var out []string
	if (flags & ImageUsageTransferSrcBit) != 0 {
		out = append(out, "VK_IMAGE_USAGE_TRANSFER_SRC_BIT")
	}
	if (flags & ImageUsageTransferDstBit) != 0 {
		out = append(out, "VK_IMAGE_USAGE_TRANSFER_DST_BIT")
	}
	if (flags & ImageUsageSampledBit) != 0 {
		out = append(out, "VK_IMAGE_USAGE_SAMPLED_BIT")
	}
	if (flags & ImageUsageStorageBit) != 0 {
		out = append(out, "VK_IMAGE_USAGE_STORAGE_BIT")
	}
	if (flags & ImageUsageColorAttachmentBit) != 0 {
		out = append(out, "VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT")
	}
	if (flags & ImageUsageDepthStencilAttachmentBit) != 0 {
		out = append(out, "VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT")
	}
	if (flags & ImageUsageTransientAttachmentBit) != 0 {
		out = append(out, "VK_IMAGE_USAGE_TRANSIENT_ATTACHMENT_BIT")
	}
	return strings.Join(out, " | ")
}

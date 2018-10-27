/*
Package vk provides bindings to the Vulkan graphics API.

Dispatchable handles

Vulkan has two kinds of handles: dispatchable and non-dispatchable.
In Vulkan, virtually all functions take a dispatchable handle as their first argument.
This loosely maps to method receivers.

In the Go bindings, we don't strictly stick to Vulkan's categorization.
Where it makes sense, we define methods on non-dispatchable handles,
for example to provide methods on CommandPool.
In the Vulkan API, VkCommandPool is a non-dispatchable handle,
and functions that act on it take a VkDevice as their first argument.

We prefer pool.AllocateCommandBuffers(...) over dev.AllocateCommandBuffers(pool, ...).

This does mean that our structs have to carry around additional information,
such as the handle of their parent object and function pointers.


Error handling

In Vulkan, functions that can fail, or partially succeed, return a VkResult.
Positive values indicate success, while negative values indicate failure.
An example for a successful result that isn't VK_SUCCESS is VK_SUBOPTIMAL_KHR,
which is returned by operations on a swapchain that no longer fully matches the surface, but still works.

In the Go bindings, we translate VkResult to the error interface.
Any value that isn't VK_SUCCESS gets turned into an error.
This does include values such as VK_SUBOPTIMAL_KHR, which aren't strictly speaking errors.
Most APIs don't actually return any successful values other than VK_SUCCESS, and the remaining values are worth handling.
*/
package vk

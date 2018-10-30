/*
Package vk provides bindings to the Vulkan graphics API.


Documentation

Vulkan is a complex beast, with a very extensive specification.
It wouldn't be feasible to include the entire specification as part of the Go documentation.
Where possible, we have copied parts of the specification (for example to describe constants),
but in order to understand the full API, it is advised to read the specification.


Performance

Using the Go bindings will be slower than using C or C++, for many reasons.

  - With CGo, every call into C involves a context switch.
  - We have to copy a lot of data to be able to pass it to C
  - Some unsafe patterns in the C API (such as pNext extension pointers)
    are wrapped in a safer way in Go, incurring additional overhead.

Some of this cost can be mitigated, however.
Unlike OpenGL, Vulkan allows construction and reuse of command buffers (command lists).
Depending on your scene, you may be able to record hundreds of draw calls once,
then use them many times over, amortizing the cost of the command buffer construction.

Dispatchable handles

Vulkan has two kinds of handles: dispatchable and non-dispatchable.
In Vulkan, virtually all functions take a dispatchable handle as their first argument.
This loosely maps to method receivers.

In the Go bindings, we stick to Vulkan's categorization.
For example, we don't define any methods on VkBuffer.
At first glance, it may make sense to treat VkBindBufferMemory as a method on VkBuffer.
However, Vulkan later added VkBindBufferMemory2, which allows binding multiple buffers to memory at once.
To avoid splitting related methods across multiple receivers, we define both on VkDevice.


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

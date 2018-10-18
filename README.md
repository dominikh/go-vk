# Vulkan Go bindings

Package `vk` provides bindings for the Vulkan graphics API.

**This package is in its very early stages. It is not in a usable state. This repository is a mirror of local development efforts.**

## Supported extensions

- VK_KHR_surface
- VK_KHR_xlib_surface
- VK_KHR_swapchain

## Build tags

We use a number of build tags to guard platform-specific functionality.

| Tag  | Description                             |
|------|-----------------------------------------|
| xlib | Include support for VK_KHR_xlib_surface |

## Comparison to github.com/vulkan-go/vulkan

github.com/vulkan-go/vulkan was generated using an automatic C binding generator and exports a raw C API.
Our bindings, on the other hand, have been written by hand, wrapping the C API in an idiomatic Go API.
We hide Vulkan's sType structure field, populating it automatically;
we fully support slices, making `count` fields unnecessary,
and we support Go strings, not requiring the user to ensure they're null-terminated.

vulkan-go/vulkan uses vkGetInstanceProcAddr for both instance and device commands,
which means all device command calls have to go through the Vulkan loader's dispatch code.
We use vkGetDeviceProcAddr and implement our own dispatch, which should theoretically be slightly faster.

Finally, we provide platform-specific functions, such as WSI, whereas
vulkan-go/vulkan expects users to use GLFW (or similar).

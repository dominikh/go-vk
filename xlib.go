// Copyright (c) 2015-2018 Khronos Group
// Copyright (c) 2018 Dominik Honnef

//go:build xlib
// +build xlib

package vk

// #cgo pkg-config: x11
// #cgo noescape domVkCreateXlibSurfaceKHR
// #cgo nocallback domVkCreateXlibSurfaceKHR
// #include "vk.h"
// #include "xlib.h"
import "C"
import "unsafe"

type XlibDisplay = C.Display
type XlibWindow = C.Window
type XlibVisualID = C.VisualID

type XlibSurfaceCreateInfoKHR struct {
	Next   unsafe.Pointer
	Dpy    *XlibDisplay
	Window XlibWindow
}

func (ins *Instance) CreateXlibSurfaceKHR(info *XlibSurfaceCreateInfoKHR) (SurfaceKHR, error) {
	// TODO(dh): support custom allocator
	var cInfo C.VkXlibSurfaceCreateInfoKHR
	cInfo.sType = C.VkStructureType(StructureTypeXlibSurfaceCreateInfoKHR)
	cInfo.pNext = info.Next
	cInfo.dpy = info.Dpy
	cInfo.window = info.Window
	var out SurfaceKHR
	res := Result(C.domVkCreateXlibSurfaceKHR(ins.fps[vkCreateXlibSurfaceKHR], ins.hnd, &cInfo, nil, &out.hnd))
	return out, result2error(res)
}

func (dev *PhysicalDevice) XlibPresentationSupportKHS(queueFamilyIndex uint32, dpy *XlibDisplay, visualID XlibVisualID) bool {
	return C.domVkGetPhysicalDeviceXlibPresentationSupportKHR(dev.instance.fps[vkGetPhysicalDeviceXlibPresentationSupportKHR], dev.hnd, C.uint32_t(queueFamilyIndex), dpy, visualID) == C.VK_TRUE
}

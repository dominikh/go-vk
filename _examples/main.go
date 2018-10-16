package main

import (
	"fmt"

	"honnef.co/go/vk"
)

func main() {
	instance, _ := vk.CreateInstance(&vk.InstanceCreateInfo{
		ApplicationInfo: &vk.ApplicationInfo{
			ApplicationName: "example",
		},
		EnabledLayerNames:     []string{"VK_LAYER_LUNARG_standard_validation"},
		EnabledExtensionNames: []string{"VK_KHR_surface", "VK_KHR_xlib_surface"},
	})
	pdevs, _ := instance.EnumeratePhysicalDevices()
	pdevs[0].QueueFamilyProperties() // called to make validation happy
	ldev, _ := pdevs[0].CreateDevice(&vk.DeviceCreateInfo{
		QueueCreateInfos: []*vk.DeviceQueueCreateInfo{
			{
				QueueFamilyIndex: 0,
				QueuePriorities:  []float32{1.0},
			},
		},
	})
	queue := ldev.Queue(0, 0)
	fmt.Println(queue)
}

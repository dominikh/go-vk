// Code generated by "stringer -type=CommandBufferLevel"; DO NOT EDIT.

package vk

import "strconv"

const _CommandBufferLevel_name = "CommandBufferLevelPrimaryCommandBufferLevelSecondary"

var _CommandBufferLevel_index = [...]uint8{0, 25, 52}

func (i CommandBufferLevel) String() string {
	if i >= CommandBufferLevel(len(_CommandBufferLevel_index)-1) {
		return "CommandBufferLevel(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _CommandBufferLevel_name[_CommandBufferLevel_index[i]:_CommandBufferLevel_index[i+1]]
}
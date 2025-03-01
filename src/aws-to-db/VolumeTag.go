package main

import "fmt"

type VolumeTag struct {
	Id       int64
	VolumeId int64
	Key      string
	Value    string
}

func NewVolumeTag(id int64, volumeId int64, key string, value string) *VolumeTag {
	return &VolumeTag{
		Id:       id,
		VolumeId: volumeId,
		Key:      key,
		Value:    value}
}

func (d *VolumeTag) String() string {
	return fmt.Sprintf("VolumeTag volumeId(%d) key(%s) value(%s)", d.VolumeId, d.Key, d.Value)
}

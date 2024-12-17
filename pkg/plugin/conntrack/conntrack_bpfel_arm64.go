// Code generated by bpf2go; DO NOT EDIT.
//go:build arm64

package conntrack

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type conntrackCtEntry struct {
	EvictionTime       uint32
	LastReportTxDir    uint32
	LastReportRxDir    uint32
	TrafficDirection   uint8
	FlagsSeenTxDir     uint8
	FlagsSeenRxDir     uint8
	IsDirectionUnknown bool
	ConntrackMetadata  struct {
		BytesForwardCount   uint64
		BytesReplyCount     uint64
		PacketsForwardCount uint32
		PacketsReplyCount   uint32
	}
}

type conntrackCtV4Key struct {
	SrcIp   uint32
	DstIp   uint32
	SrcPort uint16
	DstPort uint16
	Proto   uint8
	_       [3]byte
}

// loadConntrack returns the embedded CollectionSpec for conntrack.
func loadConntrack() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_ConntrackBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load conntrack: %w", err)
	}

	return spec, err
}

// loadConntrackObjects loads conntrack and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*conntrackObjects
//	*conntrackPrograms
//	*conntrackMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadConntrackObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadConntrack()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// conntrackSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type conntrackSpecs struct {
	conntrackProgramSpecs
	conntrackMapSpecs
}

// conntrackSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type conntrackProgramSpecs struct {
}

// conntrackMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type conntrackMapSpecs struct {
	RetinaConntrack *ebpf.MapSpec `ebpf:"retina_conntrack"`
}

// conntrackObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadConntrackObjects or ebpf.CollectionSpec.LoadAndAssign.
type conntrackObjects struct {
	conntrackPrograms
	conntrackMaps
}

func (o *conntrackObjects) Close() error {
	return _ConntrackClose(
		&o.conntrackPrograms,
		&o.conntrackMaps,
	)
}

// conntrackMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadConntrackObjects or ebpf.CollectionSpec.LoadAndAssign.
type conntrackMaps struct {
	RetinaConntrack *ebpf.Map `ebpf:"retina_conntrack"`
}

func (m *conntrackMaps) Close() error {
	return _ConntrackClose(
		m.RetinaConntrack,
	)
}

// conntrackPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadConntrackObjects or ebpf.CollectionSpec.LoadAndAssign.
type conntrackPrograms struct {
}

func (p *conntrackPrograms) Close() error {
	return _ConntrackClose()
}

func _ConntrackClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed conntrack_bpfel_arm64.o
var _ConntrackBytes []byte

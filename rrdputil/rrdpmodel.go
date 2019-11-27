package rrdputil

import (
	xml "encoding/xml"
)

type NotificationModel struct {
	XMLName   xml.Name             `xml:"notification"`
	Xmlns     string               `xml:"xmlns,attr"`
	Version   string               `xml:"version,attr"`
	SessionId string               `xml:"session_id,attr"`
	Serial    uint64               `xml:"serial,attr"`
	Snapshot  NotificationSnapshot `xml:"snapshot"`
	Deltas    []NotificationDelta  `xml:"delta"`

	//map[serial]serial: just save exist serial ,for check
	MapSerialDeltas map[uint64]uint64 `xml:"-"`
	MaxSerail       uint64            `xml:"-"`
	MinSerail       uint64            `xml:"-"`
}

type NotificationSnapshot struct {
	XMLName xml.Name `xml:"snapshot"`
	Uri     string   `xml:"uri,attr"`
	Hash    string   `xml:"hash,attr"`
}
type NotificationDelta struct {
	XMLName xml.Name `xml:"delta"`
	Serial  uint64   `xml:"serial,attr"`
	Uri     string   `xml:"uri,attr"`
	Hash    string   `xml:"hash,attr"`
}

type SnapshotModel struct {
	XMLName          xml.Name          `xml:"snapshot"`
	Xmlns            string            `xml:"xmlns,attr"`
	Version          string            `xml:"version,attr"`
	SessionId        string            `xml:"session_id,attr"`
	Serial           uint64            `xml:"serial,attr"`
	SnapshotPublishs []SnapshotPublish `xml:"publish"`
	// to check
	Hash string `xml:"-"`
}

type SnapshotPublish struct {
	XMLName xml.Name `xml:"publish"`
	Uri     string   `xml:"uri,attr"`
	Base64  string   `xml:",chardata"`
}

type DeltaModel struct {
	XMLName        xml.Name        `xml:"delta"`
	Xmlns          string          `xml:"xmlns,attr"`
	Version        string          `xml:"version,attr"`
	SessionId      string          `xml:"session_id,attr"`
	Serial         uint64          `xml:"serial,attr"`
	DeltaPublishs  []DeltaPublish  `xml:"publish"`
	DeltaWithdraws []DeltaWithdraw `xml:"withdraw"`

	// to check
	Hash string `xml:"-"`
}

type DeltaPublish struct {
	XMLName xml.Name `xml:"publish"`
	Uri     string   `xml:"uri,attr"`
	Hash    string   `xml:"hash,attr"`
	Base64  string   `xml:",chardata"`
}
type DeltaWithdraw struct {
	XMLName xml.Name `xml:"withdraw"`
	Uri     string   `xml:"uri,attr"`
	Hash    string   `xml:"hash,attr"`
}
